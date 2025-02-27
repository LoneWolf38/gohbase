// Copyright (C) 2016  The GoHBase Authors.  All rights reserved.
// This file is part of GoHBase.
// Use of this source code is governed by the Apache License 2.0
// that can be found in the COPYING file.

package gohbase

import (
	"bytes"
	"fmt"
	"io"
	"log/slog"
	"sync"

	"github.com/LoneWolf38/gohbase/hrpc"
	"modernc.org/b/v2"
)

// clientRegionCache is client -> region cache. Used to quickly
// look up all the regioninfos that map to a specific client
type clientRegionCache struct {
	m      sync.RWMutex
	logger *slog.Logger

	regions map[hrpc.RegionClient]map[hrpc.RegionInfo]struct{}
}

// put associates a region with client for provided addrss. It returns the client if it's already
// in cache or otherwise instantiates a new one by calling newClient.
// TODO: obvious place for optimization (use map with address as key to lookup exisiting clients)
func (rcc *clientRegionCache) put(addr string, r hrpc.RegionInfo,
	newClient func() hrpc.RegionClient,
) hrpc.RegionClient {
	rcc.m.Lock()
	for existingClient, regions := range rcc.regions {
		// check if client already exists, checking by host and port
		// because concurrent callers might try to put the same client
		if addr == existingClient.Addr() {
			// check client already knows about the region, checking
			// by pointer is enough because we make sure that there are
			// no regions with the same name around
			if _, ok := regions[r]; !ok {
				regions[r] = struct{}{}
			}
			rcc.m.Unlock()

			rcc.logger.Debug("region client is already in client's cache", "client", existingClient)
			return existingClient
		}
	}

	// no such client yet
	c := newClient()
	rcc.regions[c] = map[hrpc.RegionInfo]struct{}{r: {}}
	rcc.m.Unlock()

	rcc.logger.Info("added new region client", "client", c)
	return c
}

func (rcc *clientRegionCache) del(r hrpc.RegionInfo) {
	rcc.m.Lock()
	c := r.Client()
	if c != nil {
		r.SetClient(nil)
		regions := rcc.regions[c]
		delete(regions, r)
	}
	rcc.m.Unlock()
}

func (rcc *clientRegionCache) closeAll() {
	rcc.m.Lock()
	for client, regions := range rcc.regions {
		for region := range regions {
			region.MarkUnavailable()
			region.SetClient(nil)
		}
		client.Close()
	}
	rcc.m.Unlock()
}

func (rcc *clientRegionCache) clientDown(c hrpc.RegionClient) map[hrpc.RegionInfo]struct{} {
	rcc.m.Lock()
	downregions, ok := rcc.regions[c]
	delete(rcc.regions, c)
	rcc.m.Unlock()

	if ok {
		rcc.logger.Info("removed region client", "client", c)
	}
	return downregions
}

// Collects information about the clientRegion cache and appends them to the two maps to reduce
// duplication of data. We do this in one function to avoid running the iterations twice
func (rcc *clientRegionCache) debugInfo(
	regions map[string]hrpc.RegionInfo,
	clients map[string]hrpc.RegionClient,
) map[string][]string {
	// key = RegionClient memory address , value = List of RegionInfo addresses
	clientRegionCacheMap := map[string][]string{}

	rcc.m.RLock()
	for client, reginfos := range rcc.regions {
		clientRegionInfoMap := make([]string, len(reginfos))
		// put all the region infos in the client into the keyRegionInfosMap b/c its not
		// guaranteed that rcc and krc will have the same infos
		clients[fmt.Sprintf("%p", client)] = client

		i := 0
		for regionInfo := range reginfos {
			clientRegionInfoMap[i] = fmt.Sprintf("%p", regionInfo)
			regions[fmt.Sprintf("%p", regionInfo)] = regionInfo
			i++
		}

		clientRegionCacheMap[fmt.Sprintf("%p", client)] = clientRegionInfoMap
	}
	rcc.m.RUnlock()

	return clientRegionCacheMap
}

// key -> region cache.
type keyRegionCache struct {
	m      sync.RWMutex
	logger *slog.Logger

	// Maps a []byte of a region start key to a hrpc.RegionInfo
	regions *b.Tree[[]byte, hrpc.RegionInfo]
}

func (krc *keyRegionCache) get(key []byte) ([]byte, hrpc.RegionInfo) {
	krc.m.RLock()

	enum, ok := krc.regions.Seek(key)
	if ok {
		krc.m.RUnlock()
		panic(fmt.Errorf("WTF: got exact match for region search key %q", key))
	}
	k, v, err := enum.Prev()
	enum.Close()

	krc.m.RUnlock()

	if err == io.EOF {
		// we are the beginning of the tree
		return nil, nil
	}
	return k, v
}

// reads whole b tree in keyRegionCache and gathers debug info.
// We append that information in the given map
func (krc *keyRegionCache) debugInfo(
	regions map[string]hrpc.RegionInfo,
) map[string]string {
	regionCacheMap := map[string]string{}

	krc.m.RLock()
	enum, err := krc.regions.SeekFirst()
	if err != nil {
		krc.m.RUnlock()
		return regionCacheMap
	}
	krc.m.RUnlock()

	for {
		krc.m.RLock()
		k, v, err := enum.Next()
		// release lock after each iteration to allow other processes a chance to get it
		krc.m.RUnlock()
		if err == io.EOF {
			break
		}
		regions[fmt.Sprintf("%p", v)] = v
		regionCacheMap[string(k)] = fmt.Sprintf("%p", v)
	}

	return regionCacheMap
}

func isRegionOverlap(regA, regB hrpc.RegionInfo) bool {
	// if region's stop key is empty, it's assumed to be the greatest key
	return bytes.Equal(regA.Namespace(), regB.Namespace()) &&
		bytes.Equal(regA.Table(), regB.Table()) &&
		(len(regB.StopKey()) == 0 || bytes.Compare(regA.StartKey(), regB.StopKey()) < 0) &&
		(len(regA.StopKey()) == 0 || bytes.Compare(regA.StopKey(), regB.StartKey()) > 0)
}

func (krc *keyRegionCache) getOverlaps(reg hrpc.RegionInfo) []hrpc.RegionInfo {
	var overlaps []hrpc.RegionInfo
	var v hrpc.RegionInfo
	var err error

	// deal with empty tree in the beginning so that we don't have to check
	// EOF errors for enum later
	if krc.regions.Len() == 0 {
		return overlaps
	}

	// check if key created from new region falls into any cached regions
	key := createRegionSearchKey(fullyQualifiedTable(reg), reg.StartKey())
	enum, ok := krc.regions.Seek(key)
	if ok {
		panic(fmt.Errorf("WTF: found a region with exact name as the search key %q", key))
	}

	// case 1: landed before the first region in cache
	// enum.Prev() returns io.EOF
	// enum.Next() returns io.EOF
	// SeekFirst() + enum.Next() returns the first region, which has larger start key

	// case 2: landed before the second region in cache
	// enum.Prev() returns the first region X and moves pointer to -infinity
	// enum.Next() returns io.EOF
	// SeekFirst() + enum.Next() returns first region X, which has smaller start key

	// case 3: landed anywhere after the second region
	// enum.Prev() returns the region X before it landed, moves pointer to the region X - 1
	// enum.Next() returns X - 1 and move pointer to X, which has smaller start key

	_, _, _ = enum.Prev()
	_, _, err = enum.Next()
	if err == io.EOF {
		// we are in the beginning of tree, get new enum starting
		// from first region
		enum.Close()
		enum, err = krc.regions.SeekFirst()
		if err != nil {
			panic(fmt.Errorf(
				"error seeking first region when getting overlaps for region %v: %v", reg, err))
		}
	}

	_, v, err = enum.Next()
	if err != nil {
		panic(fmt.Errorf(
			"error accessing first region when getting overlaps for region %v: %v", reg, err))
	}
	if isRegionOverlap(v, reg) {
		overlaps = append(overlaps, v)
	}
	_, v, err = enum.Next()

	// now append all regions that overlap until the end of the tree
	// or until they don't overlap
	for err != io.EOF && isRegionOverlap(v, reg) {
		overlaps = append(overlaps, v)
		_, v, err = enum.Next()
	}
	enum.Close()
	return overlaps
}

// put looks up if there's already region with this name in regions cache
// and if there's, returns it in overlaps and doesn't modify the cache.
// Otherwise, it puts the region and removes all overlaps in case all of
// them are older. Returns a slice of overlapping regions and whether
// passed region was put in the cache.
func (krc *keyRegionCache) put(reg hrpc.RegionInfo) (overlaps []hrpc.RegionInfo, replaced bool) {
	krc.m.Lock()
	defer krc.m.Unlock()

	// Update region cache metric
	beforeLen := krc.regions.Len()
	defer func() {
		afterLen := krc.regions.Len()
		cachedRegionTotal.Add(float64(afterLen - beforeLen))
	}()

	krc.regions.Put(reg.Name(), func(v hrpc.RegionInfo, exists bool) (hrpc.RegionInfo, bool) {
		if exists {
			// region is already in cache,
			// note: regions with the same name have the same age
			overlaps = []hrpc.RegionInfo{v}
			return nil, false
		}
		// find all entries that are overlapping with the range of the new region.
		overlaps = krc.getOverlaps(reg)
		for _, o := range overlaps {
			if o.ID() > reg.ID() {
				// overlapping region is younger,
				// don't replace any regions
				// TODO: figure out if there can a case where we might
				// have both older and younger overlapping regions, for
				// now we only replace if all overlaps are older
				return nil, false
			}
		}
		// all overlaps are older, put the new region
		replaced = true
		return reg, true
	})
	if !replaced {
		krc.logger.Debug("region is already in cache",
			"region", reg, "overlaps", overlaps, "replaced", replaced)
		return
	}
	// delete overlapping regions
	// TODO: in case overlaps are always either younger or older,
	// we can just greedily remove them in Put function
	for _, o := range overlaps {
		krc.regions.Delete(o.Name())
		// let region establishers know that they can give up
		o.MarkDead()
	}

	krc.logger.Info("added new region",
		"region", reg, "overlaps", overlaps, "replaced", replaced)
	return
}

func (krc *keyRegionCache) del(reg hrpc.RegionInfo) bool {
	krc.m.Lock()
	success := krc.regions.Delete(reg.Name())
	krc.m.Unlock()
	// let region establishers know that they can give up
	reg.MarkDead()

	if success {
		cachedRegionTotal.Dec()
	}
	krc.logger.Debug("removed region", "region", reg)
	return success
}
