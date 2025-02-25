//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.21.12
// source: BucketCacheEntry.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BlockType int32

const (
	BlockType_data                     BlockType = 0
	BlockType_encoded_data             BlockType = 1
	BlockType_leaf_index               BlockType = 2
	BlockType_bloom_chunk              BlockType = 3
	BlockType_meta                     BlockType = 4
	BlockType_intermediate_index       BlockType = 5
	BlockType_root_index               BlockType = 6
	BlockType_file_info                BlockType = 7
	BlockType_general_bloom_meta       BlockType = 8
	BlockType_delete_family_bloom_meta BlockType = 9
	BlockType_trailer                  BlockType = 10
	BlockType_index_v1                 BlockType = 11
)

// Enum value maps for BlockType.
var (
	BlockType_name = map[int32]string{
		0:  "data",
		1:  "encoded_data",
		2:  "leaf_index",
		3:  "bloom_chunk",
		4:  "meta",
		5:  "intermediate_index",
		6:  "root_index",
		7:  "file_info",
		8:  "general_bloom_meta",
		9:  "delete_family_bloom_meta",
		10: "trailer",
		11: "index_v1",
	}
	BlockType_value = map[string]int32{
		"data":                     0,
		"encoded_data":             1,
		"leaf_index":               2,
		"bloom_chunk":              3,
		"meta":                     4,
		"intermediate_index":       5,
		"root_index":               6,
		"file_info":                7,
		"general_bloom_meta":       8,
		"delete_family_bloom_meta": 9,
		"trailer":                  10,
		"index_v1":                 11,
	}
)

func (x BlockType) Enum() *BlockType {
	p := new(BlockType)
	*p = x
	return p
}

func (x BlockType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BlockType) Descriptor() protoreflect.EnumDescriptor {
	return file_BucketCacheEntry_proto_enumTypes[0].Descriptor()
}

func (BlockType) Type() protoreflect.EnumType {
	return &file_BucketCacheEntry_proto_enumTypes[0]
}

func (x BlockType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *BlockType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = BlockType(num)
	return nil
}

// Deprecated: Use BlockType.Descriptor instead.
func (BlockType) EnumDescriptor() ([]byte, []int) {
	return file_BucketCacheEntry_proto_rawDescGZIP(), []int{0}
}

type BlockPriority int32

const (
	BlockPriority_single BlockPriority = 0
	BlockPriority_multi  BlockPriority = 1
	BlockPriority_memory BlockPriority = 2
)

// Enum value maps for BlockPriority.
var (
	BlockPriority_name = map[int32]string{
		0: "single",
		1: "multi",
		2: "memory",
	}
	BlockPriority_value = map[string]int32{
		"single": 0,
		"multi":  1,
		"memory": 2,
	}
)

func (x BlockPriority) Enum() *BlockPriority {
	p := new(BlockPriority)
	*p = x
	return p
}

func (x BlockPriority) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BlockPriority) Descriptor() protoreflect.EnumDescriptor {
	return file_BucketCacheEntry_proto_enumTypes[1].Descriptor()
}

func (BlockPriority) Type() protoreflect.EnumType {
	return &file_BucketCacheEntry_proto_enumTypes[1]
}

func (x BlockPriority) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *BlockPriority) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = BlockPriority(num)
	return nil
}

// Deprecated: Use BlockPriority.Descriptor instead.
func (BlockPriority) EnumDescriptor() ([]byte, []int) {
	return file_BucketCacheEntry_proto_rawDescGZIP(), []int{1}
}

type BucketCacheEntry struct {
	state         protoimpl.MessageState        `protogen:"open.v1"`
	CacheCapacity *int64                        `protobuf:"varint,1,req,name=cache_capacity,json=cacheCapacity" json:"cache_capacity,omitempty"`
	IoClass       *string                       `protobuf:"bytes,2,req,name=io_class,json=ioClass" json:"io_class,omitempty"`
	MapClass      *string                       `protobuf:"bytes,3,req,name=map_class,json=mapClass" json:"map_class,omitempty"`
	Deserializers map[int32]string              `protobuf:"bytes,4,rep,name=deserializers" json:"deserializers,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	BackingMap    *BackingMap                   `protobuf:"bytes,5,req,name=backing_map,json=backingMap" json:"backing_map,omitempty"`
	Checksum      []byte                        `protobuf:"bytes,6,opt,name=checksum" json:"checksum,omitempty"`
	CachedFiles   map[string]*RegionFileSizeMap `protobuf:"bytes,7,rep,name=cached_files,json=cachedFiles" json:"cached_files,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BucketCacheEntry) Reset() {
	*x = BucketCacheEntry{}
	mi := &file_BucketCacheEntry_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BucketCacheEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BucketCacheEntry) ProtoMessage() {}

func (x *BucketCacheEntry) ProtoReflect() protoreflect.Message {
	mi := &file_BucketCacheEntry_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BucketCacheEntry.ProtoReflect.Descriptor instead.
func (*BucketCacheEntry) Descriptor() ([]byte, []int) {
	return file_BucketCacheEntry_proto_rawDescGZIP(), []int{0}
}

func (x *BucketCacheEntry) GetCacheCapacity() int64 {
	if x != nil && x.CacheCapacity != nil {
		return *x.CacheCapacity
	}
	return 0
}

func (x *BucketCacheEntry) GetIoClass() string {
	if x != nil && x.IoClass != nil {
		return *x.IoClass
	}
	return ""
}

func (x *BucketCacheEntry) GetMapClass() string {
	if x != nil && x.MapClass != nil {
		return *x.MapClass
	}
	return ""
}

func (x *BucketCacheEntry) GetDeserializers() map[int32]string {
	if x != nil {
		return x.Deserializers
	}
	return nil
}

func (x *BucketCacheEntry) GetBackingMap() *BackingMap {
	if x != nil {
		return x.BackingMap
	}
	return nil
}

func (x *BucketCacheEntry) GetChecksum() []byte {
	if x != nil {
		return x.Checksum
	}
	return nil
}

func (x *BucketCacheEntry) GetCachedFiles() map[string]*RegionFileSizeMap {
	if x != nil {
		return x.CachedFiles
	}
	return nil
}

type BackingMap struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Entry         []*BackingMapEntry     `protobuf:"bytes,1,rep,name=entry" json:"entry,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BackingMap) Reset() {
	*x = BackingMap{}
	mi := &file_BucketCacheEntry_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BackingMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackingMap) ProtoMessage() {}

func (x *BackingMap) ProtoReflect() protoreflect.Message {
	mi := &file_BucketCacheEntry_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackingMap.ProtoReflect.Descriptor instead.
func (*BackingMap) Descriptor() ([]byte, []int) {
	return file_BucketCacheEntry_proto_rawDescGZIP(), []int{1}
}

func (x *BackingMap) GetEntry() []*BackingMapEntry {
	if x != nil {
		return x.Entry
	}
	return nil
}

type BackingMapEntry struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           *BlockCacheKey         `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	Value         *BucketEntry           `protobuf:"bytes,2,req,name=value" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BackingMapEntry) Reset() {
	*x = BackingMapEntry{}
	mi := &file_BucketCacheEntry_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BackingMapEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackingMapEntry) ProtoMessage() {}

func (x *BackingMapEntry) ProtoReflect() protoreflect.Message {
	mi := &file_BucketCacheEntry_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackingMapEntry.ProtoReflect.Descriptor instead.
func (*BackingMapEntry) Descriptor() ([]byte, []int) {
	return file_BucketCacheEntry_proto_rawDescGZIP(), []int{2}
}

func (x *BackingMapEntry) GetKey() *BlockCacheKey {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *BackingMapEntry) GetValue() *BucketEntry {
	if x != nil {
		return x.Value
	}
	return nil
}

type BlockCacheKey struct {
	state               protoimpl.MessageState `protogen:"open.v1"`
	Hfilename           *string                `protobuf:"bytes,1,req,name=hfilename" json:"hfilename,omitempty"`
	Offset              *int64                 `protobuf:"varint,2,req,name=offset" json:"offset,omitempty"`
	BlockType           *BlockType             `protobuf:"varint,3,req,name=block_type,json=blockType,enum=pb.BlockType" json:"block_type,omitempty"`
	PrimaryReplicaBlock *bool                  `protobuf:"varint,4,req,name=primary_replica_block,json=primaryReplicaBlock" json:"primary_replica_block,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *BlockCacheKey) Reset() {
	*x = BlockCacheKey{}
	mi := &file_BucketCacheEntry_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BlockCacheKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockCacheKey) ProtoMessage() {}

func (x *BlockCacheKey) ProtoReflect() protoreflect.Message {
	mi := &file_BucketCacheEntry_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockCacheKey.ProtoReflect.Descriptor instead.
func (*BlockCacheKey) Descriptor() ([]byte, []int) {
	return file_BucketCacheEntry_proto_rawDescGZIP(), []int{3}
}

func (x *BlockCacheKey) GetHfilename() string {
	if x != nil && x.Hfilename != nil {
		return *x.Hfilename
	}
	return ""
}

func (x *BlockCacheKey) GetOffset() int64 {
	if x != nil && x.Offset != nil {
		return *x.Offset
	}
	return 0
}

func (x *BlockCacheKey) GetBlockType() BlockType {
	if x != nil && x.BlockType != nil {
		return *x.BlockType
	}
	return BlockType_data
}

func (x *BlockCacheKey) GetPrimaryReplicaBlock() bool {
	if x != nil && x.PrimaryReplicaBlock != nil {
		return *x.PrimaryReplicaBlock
	}
	return false
}

type BucketEntry struct {
	state              protoimpl.MessageState `protogen:"open.v1"`
	Offset             *int64                 `protobuf:"varint,1,req,name=offset" json:"offset,omitempty"`
	Length             *int32                 `protobuf:"varint,2,req,name=length" json:"length,omitempty"`
	AccessCounter      *int64                 `protobuf:"varint,3,req,name=access_counter,json=accessCounter" json:"access_counter,omitempty"`
	DeserialiserIndex  *int32                 `protobuf:"varint,4,req,name=deserialiser_index,json=deserialiserIndex" json:"deserialiser_index,omitempty"`
	Priority           *BlockPriority         `protobuf:"varint,5,req,name=priority,enum=pb.BlockPriority" json:"priority,omitempty"`
	CachedTime         *int64                 `protobuf:"varint,6,req,name=cachedTime" json:"cachedTime,omitempty"`
	DiskSizeWithHeader *int32                 `protobuf:"varint,7,opt,name=disk_size_with_header,json=diskSizeWithHeader" json:"disk_size_with_header,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *BucketEntry) Reset() {
	*x = BucketEntry{}
	mi := &file_BucketCacheEntry_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BucketEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BucketEntry) ProtoMessage() {}

func (x *BucketEntry) ProtoReflect() protoreflect.Message {
	mi := &file_BucketCacheEntry_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BucketEntry.ProtoReflect.Descriptor instead.
func (*BucketEntry) Descriptor() ([]byte, []int) {
	return file_BucketCacheEntry_proto_rawDescGZIP(), []int{4}
}

func (x *BucketEntry) GetOffset() int64 {
	if x != nil && x.Offset != nil {
		return *x.Offset
	}
	return 0
}

func (x *BucketEntry) GetLength() int32 {
	if x != nil && x.Length != nil {
		return *x.Length
	}
	return 0
}

func (x *BucketEntry) GetAccessCounter() int64 {
	if x != nil && x.AccessCounter != nil {
		return *x.AccessCounter
	}
	return 0
}

func (x *BucketEntry) GetDeserialiserIndex() int32 {
	if x != nil && x.DeserialiserIndex != nil {
		return *x.DeserialiserIndex
	}
	return 0
}

func (x *BucketEntry) GetPriority() BlockPriority {
	if x != nil && x.Priority != nil {
		return *x.Priority
	}
	return BlockPriority_single
}

func (x *BucketEntry) GetCachedTime() int64 {
	if x != nil && x.CachedTime != nil {
		return *x.CachedTime
	}
	return 0
}

func (x *BucketEntry) GetDiskSizeWithHeader() int32 {
	if x != nil && x.DiskSizeWithHeader != nil {
		return *x.DiskSizeWithHeader
	}
	return 0
}

type RegionFileSizeMap struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	RegionName       *string                `protobuf:"bytes,1,req,name=region_name,json=regionName" json:"region_name,omitempty"`
	RegionCachedSize *uint64                `protobuf:"varint,2,req,name=region_cached_size,json=regionCachedSize" json:"region_cached_size,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *RegionFileSizeMap) Reset() {
	*x = RegionFileSizeMap{}
	mi := &file_BucketCacheEntry_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegionFileSizeMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegionFileSizeMap) ProtoMessage() {}

func (x *RegionFileSizeMap) ProtoReflect() protoreflect.Message {
	mi := &file_BucketCacheEntry_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegionFileSizeMap.ProtoReflect.Descriptor instead.
func (*RegionFileSizeMap) Descriptor() ([]byte, []int) {
	return file_BucketCacheEntry_proto_rawDescGZIP(), []int{5}
}

func (x *RegionFileSizeMap) GetRegionName() string {
	if x != nil && x.RegionName != nil {
		return *x.RegionName
	}
	return ""
}

func (x *RegionFileSizeMap) GetRegionCachedSize() uint64 {
	if x != nil && x.RegionCachedSize != nil {
		return *x.RegionCachedSize
	}
	return 0
}

var File_BucketCacheEntry_proto protoreflect.FileDescriptor

var file_BucketCacheEntry_proto_rawDesc = string([]byte{
	0x0a, 0x16, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x61, 0x63, 0x68, 0x65, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0xf0, 0x03, 0x0a,
	0x10, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x61, 0x63, 0x68, 0x65, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x63, 0x61, 0x70, 0x61, 0x63,
	0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x02, 0x28, 0x03, 0x52, 0x0d, 0x63, 0x61, 0x63, 0x68, 0x65,
	0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x6f, 0x5f, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x07, 0x69, 0x6f, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x70, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x18, 0x03, 0x20, 0x02, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x61, 0x70, 0x43, 0x6c, 0x61, 0x73, 0x73,
	0x12, 0x4d, 0x0a, 0x0d, 0x64, 0x65, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x43, 0x61, 0x63, 0x68, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x44, 0x65,
	0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0d, 0x64, 0x65, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x73, 0x12,
	0x2f, 0x0a, 0x0b, 0x62, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x05,
	0x20, 0x02, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x69, 0x6e,
	0x67, 0x4d, 0x61, 0x70, 0x52, 0x0a, 0x62, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x70,
	0x12, 0x1a, 0x0a, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x12, 0x48, 0x0a, 0x0c,
	0x63, 0x61, 0x63, 0x68, 0x65, 0x64, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x61,
	0x63, 0x68, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x64, 0x46,
	0x69, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x63, 0x61, 0x63, 0x68, 0x65,
	0x64, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x1a, 0x40, 0x0a, 0x12, 0x44, 0x65, 0x73, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x55, 0x0a, 0x10, 0x43, 0x61, 0x63, 0x68,
	0x65, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2b,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x70, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x4d, 0x61, 0x70, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x37, 0x0a, 0x0a, 0x42, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x70, 0x12, 0x29, 0x0a,
	0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70,
	0x62, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x5d, 0x0a, 0x0f, 0x42, 0x61, 0x63, 0x6b,
	0x69, 0x6e, 0x67, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x23, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x43, 0x61, 0x63, 0x68, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x25, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xa7, 0x01, 0x0a, 0x0d, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x43, 0x61, 0x63, 0x68, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x68, 0x66, 0x69,
	0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x09, 0x68, 0x66,
	0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x18, 0x02, 0x20, 0x02, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12,
	0x2c, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x02, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x32, 0x0a,
	0x15, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x04, 0x20, 0x02, 0x28, 0x08, 0x52, 0x13, 0x70, 0x72,
	0x69, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x22, 0x95, 0x02, 0x0a, 0x0b, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x02, 0x28,
	0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x6e,
	0x67, 0x74, 0x68, 0x18, 0x02, 0x20, 0x02, 0x28, 0x05, 0x52, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74,
	0x68, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x02, 0x28, 0x03, 0x52, 0x0d, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x12, 0x64, 0x65, 0x73, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x69, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04,
	0x20, 0x02, 0x28, 0x05, 0x52, 0x11, 0x64, 0x65, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x2d, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72,
	0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x02, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x52, 0x08, 0x70, 0x72,
	0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x63, 0x68, 0x65, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x02, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x61, 0x63, 0x68,
	0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x15, 0x64, 0x69, 0x73, 0x6b, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x64, 0x69, 0x73, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x57,
	0x69, 0x74, 0x68, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x22, 0x62, 0x0a, 0x11, 0x52, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x4d, 0x61, 0x70, 0x12, 0x1f,
	0x0a, 0x0b, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x02, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x2c, 0x0a, 0x12, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x64,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x04, 0x52, 0x10, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x64, 0x53, 0x69, 0x7a, 0x65, 0x2a, 0xda, 0x01,
	0x0a, 0x09, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x6c, 0x65, 0x61, 0x66, 0x5f,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x6f, 0x6d,
	0x5f, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61,
	0x10, 0x04, 0x12, 0x16, 0x0a, 0x12, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6d, 0x65, 0x64, 0x69, 0x61,
	0x74, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x10, 0x05, 0x12, 0x0e, 0x0a, 0x0a, 0x72, 0x6f,
	0x6f, 0x74, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x10, 0x06, 0x12, 0x0d, 0x0a, 0x09, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x10, 0x07, 0x12, 0x16, 0x0a, 0x12, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x6c, 0x5f, 0x62, 0x6c, 0x6f, 0x6f, 0x6d, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x10,
	0x08, 0x12, 0x1c, 0x0a, 0x18, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x66, 0x61, 0x6d, 0x69,
	0x6c, 0x79, 0x5f, 0x62, 0x6c, 0x6f, 0x6f, 0x6d, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x10, 0x09, 0x12,
	0x0b, 0x0a, 0x07, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x10, 0x0a, 0x12, 0x0c, 0x0a, 0x08,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x76, 0x31, 0x10, 0x0b, 0x2a, 0x32, 0x0a, 0x0d, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x0a, 0x0a, 0x06, 0x73,
	0x69, 0x6e, 0x67, 0x6c, 0x65, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x6d, 0x75, 0x6c, 0x74, 0x69,
	0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x10, 0x02, 0x42, 0x55,
	0x0a, 0x31, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x68, 0x61, 0x64,
	0x6f, 0x6f, 0x70, 0x2e, 0x68, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x68, 0x61, 0x64, 0x65, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x64, 0x42, 0x11, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x61, 0x63, 0x68, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x48, 0x01, 0x5a, 0x05, 0x2e, 0x2e, 0x2f, 0x70, 0x62, 0x88,
	0x01, 0x01, 0xa0, 0x01, 0x01,
})

var (
	file_BucketCacheEntry_proto_rawDescOnce sync.Once
	file_BucketCacheEntry_proto_rawDescData []byte
)

func file_BucketCacheEntry_proto_rawDescGZIP() []byte {
	file_BucketCacheEntry_proto_rawDescOnce.Do(func() {
		file_BucketCacheEntry_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_BucketCacheEntry_proto_rawDesc), len(file_BucketCacheEntry_proto_rawDesc)))
	})
	return file_BucketCacheEntry_proto_rawDescData
}

var file_BucketCacheEntry_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_BucketCacheEntry_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_BucketCacheEntry_proto_goTypes = []any{
	(BlockType)(0),            // 0: pb.BlockType
	(BlockPriority)(0),        // 1: pb.BlockPriority
	(*BucketCacheEntry)(nil),  // 2: pb.BucketCacheEntry
	(*BackingMap)(nil),        // 3: pb.BackingMap
	(*BackingMapEntry)(nil),   // 4: pb.BackingMapEntry
	(*BlockCacheKey)(nil),     // 5: pb.BlockCacheKey
	(*BucketEntry)(nil),       // 6: pb.BucketEntry
	(*RegionFileSizeMap)(nil), // 7: pb.RegionFileSizeMap
	nil,                       // 8: pb.BucketCacheEntry.DeserializersEntry
	nil,                       // 9: pb.BucketCacheEntry.CachedFilesEntry
}
var file_BucketCacheEntry_proto_depIdxs = []int32{
	8, // 0: pb.BucketCacheEntry.deserializers:type_name -> pb.BucketCacheEntry.DeserializersEntry
	3, // 1: pb.BucketCacheEntry.backing_map:type_name -> pb.BackingMap
	9, // 2: pb.BucketCacheEntry.cached_files:type_name -> pb.BucketCacheEntry.CachedFilesEntry
	4, // 3: pb.BackingMap.entry:type_name -> pb.BackingMapEntry
	5, // 4: pb.BackingMapEntry.key:type_name -> pb.BlockCacheKey
	6, // 5: pb.BackingMapEntry.value:type_name -> pb.BucketEntry
	0, // 6: pb.BlockCacheKey.block_type:type_name -> pb.BlockType
	1, // 7: pb.BucketEntry.priority:type_name -> pb.BlockPriority
	7, // 8: pb.BucketCacheEntry.CachedFilesEntry.value:type_name -> pb.RegionFileSizeMap
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_BucketCacheEntry_proto_init() }
func file_BucketCacheEntry_proto_init() {
	if File_BucketCacheEntry_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_BucketCacheEntry_proto_rawDesc), len(file_BucketCacheEntry_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_BucketCacheEntry_proto_goTypes,
		DependencyIndexes: file_BucketCacheEntry_proto_depIdxs,
		EnumInfos:         file_BucketCacheEntry_proto_enumTypes,
		MessageInfos:      file_BucketCacheEntry_proto_msgTypes,
	}.Build()
	File_BucketCacheEntry_proto = out.File
	file_BucketCacheEntry_proto_goTypes = nil
	file_BucketCacheEntry_proto_depIdxs = nil
}
