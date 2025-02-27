// Copyright (C) 2017  The GoHBase Authors.  All rights reserved.
// This file is part of GoHBase.
// Use of this source code is governed by the Apache License 2.0
// that can be found in the COPYING file.

package hrpc

import (
	"errors"
	"math"
	"time"

	"github.com/LoneWolf38/gohbase/filter"
	"github.com/LoneWolf38/gohbase/pb"
)

// baseQuery bundles common fields that can be provided for quering requests: Scans and Gets
type baseQuery struct {
	families      map[string][]string
	filter        *pb.Filter
	fromTimestamp uint64
	toTimestamp   uint64
	maxVersions   uint32
	storeLimit    uint32
	storeOffset   uint32
	priority      uint32
	cacheBlocks   bool
	consistency   ConsistencyType
}

// ConsistencyType is used to specify the required consistency of data
//
// See https://docs.cloudera.com/HDPDocuments/HDP2/HDP-2.2.9/bk_hadoop-ha/
//
//	content/ha-hbase-timeline-consistency.html
type ConsistencyType int

const (
	// Use HBase's default
	DefaultConsistency ConsistencyType = iota

	// Guarantees that the client receives the latest data.
	StrongConsistency

	// Client might receive stale data (indicated by the Stale field), but
	// the data that is received was valid at a given point of time.
	TimelineConsistency
)

func (c ConsistencyType) toProto() (ret *pb.Consistency) {
	ret = new(pb.Consistency)
	switch c {
	case TimelineConsistency:
		*ret = pb.Consistency_TIMELINE
		return
	case StrongConsistency:
		*ret = pb.Consistency_STRONG
		return
	case DefaultConsistency:
		panic("default consistency depends on context")
	}
	panic("invalid value for ConsistencyType")
}

// newBaseQuery return baseQuery with all default values
func newBaseQuery() baseQuery {
	return baseQuery{
		storeLimit:    DefaultMaxResultsPerColumnFamily,
		fromTimestamp: MinTimestamp,
		toTimestamp:   MaxTimestamp,
		maxVersions:   DefaultMaxVersions,
		cacheBlocks:   DefaultCacheBlocks,
	}
}

func (bq *baseQuery) setFamilies(families map[string][]string) {
	bq.families = families
}

func (bq *baseQuery) setFilter(filter *pb.Filter) {
	bq.filter = filter
}

func (bq *baseQuery) setTimeRangeUint64(from, to uint64) {
	bq.fromTimestamp = from
	bq.toTimestamp = to
}

func (bq *baseQuery) setMaxVersions(versions uint32) {
	bq.maxVersions = versions
}

func (bq *baseQuery) setMaxResultsPerColumnFamily(maxresults uint32) {
	bq.storeLimit = maxresults
}

func (bq *baseQuery) setResultOffset(offset uint32) {
	bq.storeOffset = offset
}

func (bq *baseQuery) setCacheBlocks(cacheBlocks bool) {
	bq.cacheBlocks = cacheBlocks
}

func (bq *baseQuery) setConsistency(consistency ConsistencyType) {
	bq.consistency = consistency
}

func (bq *baseQuery) setPriority(priority uint32) {
	bq.priority = priority
}

func (bq *baseQuery) Priority() uint32 {
	return bq.priority
}

// GetPriority returns the priority of a Call. Returns 0 for calls
// that don't have a priority set or don't support setting a priority.
func GetPriority(c Call) uint32 {
	p, ok := c.(interface{ Priority() uint32 })
	if !ok {
		return 0
	}
	return p.Priority()
}

// Families option adds families constraint to a Scan or Get request.
func Families(f map[string][]string) func(Call) error {
	return func(hc Call) error {
		if c, ok := hc.(hasQueryOptions); ok {
			c.setFamilies(f)
			return nil
		}
		return errors.New("'Families' option can only be used with Get or Scan request")
	}
}

// Filters option adds filters constraint to a Scan or Get request.
func Filters(f filter.Filter) func(Call) error {
	return func(hc Call) error {
		if c, ok := hc.(hasQueryOptions); ok {
			pbF, err := f.ConstructPBFilter()
			if err != nil {
				return err
			}
			c.setFilter(pbF)
			return nil
		}
		return errors.New("'Filters' option can only be used with Get or Scan request")
	}
}

// TimeRange is used as a parameter for request creation. Adds TimeRange constraint to a request.
// It will get values in range [from, to[ ('to' is exclusive).
func TimeRange(from, to time.Time) func(Call) error {
	return TimeRangeUint64(uint64(from.UnixNano()/1e6), uint64(to.UnixNano()/1e6))
}

// TimeRangeUint64 is used as a parameter for request creation.
// Adds TimeRange constraint to a request.
// from and to should be in milliseconds
// // It will get values in range [from, to[ ('to' is exclusive).
func TimeRangeUint64(from, to uint64) func(Call) error {
	return func(hc Call) error {
		if c, ok := hc.(hasQueryOptions); ok {
			if from >= to {
				// or equal is becuase 'to' is exclusive
				return errors.New("'from' timestamp is greater or equal to 'to' timestamp")
			}
			c.setTimeRangeUint64(from, to)
			return nil
		}
		return errors.New("'TimeRange' option can only be used with Get or Scan request")
	}
}

// MaxVersions is used as a parameter for request creation.
// Adds MaxVersions constraint to a request.
func MaxVersions(versions uint32) func(Call) error {
	return func(hc Call) error {
		if c, ok := hc.(hasQueryOptions); ok {
			if versions > math.MaxInt32 {
				return errors.New("'MaxVersions' exceeds supported number of versions")
			}
			c.setMaxVersions(versions)
			return nil
		}
		return errors.New("'MaxVersions' option can only be used with Get or Scan request")
	}
}

// MaxResultsPerColumnFamily is an option for Get or Scan requests that sets the maximum
// number of cells returned per column family in a row.
func MaxResultsPerColumnFamily(maxresults uint32) func(Call) error {
	return func(hc Call) error {
		if c, ok := hc.(hasQueryOptions); ok {
			if maxresults > math.MaxInt32 {
				return errors.New(
					"'MaxResultsPerColumnFamily' exceeds supported number of value results")
			}
			c.setMaxResultsPerColumnFamily(maxresults)
			return nil
		}
		return errors.New(
			"'MaxResultsPerColumnFamily' option can only be used with Get or Scan request")
	}
}

// ResultOffset is a option for Scan or Get requests that sets the offset for cells
// within a column family.
func ResultOffset(offset uint32) func(Call) error {
	return func(hc Call) error {
		if c, ok := hc.(hasQueryOptions); ok {
			if offset > math.MaxInt32 {
				return errors.New("'ResultOffset' exceeds supported offset value")
			}
			c.setResultOffset(offset)
			return nil
		}
		return errors.New("'ResultOffset' option can only be used with Get or Scan request")
	}
}

// CacheBlocks is an option for Scan or Get requests to enable/disable the block cache
// for the request
func CacheBlocks(cacheBlocks bool) func(Call) error {
	return func(hc Call) error {
		if c, ok := hc.(hasQueryOptions); ok {
			c.setCacheBlocks(cacheBlocks)
			return nil
		}
		return errors.New("'CacheBlocks' option can only be used with Get or Scan request")
	}
}

// Consistency is a Scan or Get option that requests the given
// consistency of data.
func Consistency(consistency ConsistencyType) func(Call) error {
	return func(g Call) error {
		if c, ok := g.(hasQueryOptions); ok {
			c.setConsistency(consistency)
			return nil
		}
		return errors.New("'Consistency' option can only be used with Get or Scan requests")
	}
}

func Priority(priority uint32) func(Call) error {
	return func(hc Call) error {
		if c, ok := hc.(hasQueryOptions); ok {
			c.setPriority(priority)
			return nil
		}
		return errors.New("'Priority' option can only be used with Get or Scan requests")
	}
}
