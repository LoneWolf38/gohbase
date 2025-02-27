//*
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
// source: SnapshotCleanup.proto

// This file contains protocol buffers to represent the state of the snapshot auto cleanup based on TTL

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

type SnapshotCleanupState struct {
	state                  protoimpl.MessageState `protogen:"open.v1"`
	SnapshotCleanupEnabled *bool                  `protobuf:"varint,1,req,name=snapshot_cleanup_enabled,json=snapshotCleanupEnabled" json:"snapshot_cleanup_enabled,omitempty"`
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *SnapshotCleanupState) Reset() {
	*x = SnapshotCleanupState{}
	mi := &file_SnapshotCleanup_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SnapshotCleanupState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapshotCleanupState) ProtoMessage() {}

func (x *SnapshotCleanupState) ProtoReflect() protoreflect.Message {
	mi := &file_SnapshotCleanup_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapshotCleanupState.ProtoReflect.Descriptor instead.
func (*SnapshotCleanupState) Descriptor() ([]byte, []int) {
	return file_SnapshotCleanup_proto_rawDescGZIP(), []int{0}
}

func (x *SnapshotCleanupState) GetSnapshotCleanupEnabled() bool {
	if x != nil && x.SnapshotCleanupEnabled != nil {
		return *x.SnapshotCleanupEnabled
	}
	return false
}

var File_SnapshotCleanup_proto protoreflect.FileDescriptor

var file_SnapshotCleanup_proto_rawDesc = string([]byte{
	0x0a, 0x15, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x43, 0x6c, 0x65, 0x61, 0x6e, 0x75,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x50, 0x0a, 0x14, 0x53,
	0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x43, 0x6c, 0x65, 0x61, 0x6e, 0x75, 0x70, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x38, 0x0a, 0x18, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x5f,
	0x63, 0x6c, 0x65, 0x61, 0x6e, 0x75, 0x70, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x02, 0x28, 0x08, 0x52, 0x16, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x43,
	0x6c, 0x65, 0x61, 0x6e, 0x75, 0x70, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x42, 0x56, 0x0a,
	0x31, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x68, 0x61, 0x64, 0x6f,
	0x6f, 0x70, 0x2e, 0x68, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x68, 0x61, 0x64, 0x65, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x64, 0x42, 0x15, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x43, 0x6c, 0x65, 0x61,
	0x6e, 0x75, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x48, 0x01, 0x5a, 0x05, 0x2e, 0x2e, 0x2f,
	0x70, 0x62, 0xa0, 0x01, 0x01,
})

var (
	file_SnapshotCleanup_proto_rawDescOnce sync.Once
	file_SnapshotCleanup_proto_rawDescData []byte
)

func file_SnapshotCleanup_proto_rawDescGZIP() []byte {
	file_SnapshotCleanup_proto_rawDescOnce.Do(func() {
		file_SnapshotCleanup_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_SnapshotCleanup_proto_rawDesc), len(file_SnapshotCleanup_proto_rawDesc)))
	})
	return file_SnapshotCleanup_proto_rawDescData
}

var file_SnapshotCleanup_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SnapshotCleanup_proto_goTypes = []any{
	(*SnapshotCleanupState)(nil), // 0: pb.SnapshotCleanupState
}
var file_SnapshotCleanup_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_SnapshotCleanup_proto_init() }
func file_SnapshotCleanup_proto_init() {
	if File_SnapshotCleanup_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_SnapshotCleanup_proto_rawDesc), len(file_SnapshotCleanup_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SnapshotCleanup_proto_goTypes,
		DependencyIndexes: file_SnapshotCleanup_proto_depIdxs,
		MessageInfos:      file_SnapshotCleanup_proto_msgTypes,
	}.Build()
	File_SnapshotCleanup_proto = out.File
	file_SnapshotCleanup_proto_goTypes = nil
	file_SnapshotCleanup_proto_depIdxs = nil
}
