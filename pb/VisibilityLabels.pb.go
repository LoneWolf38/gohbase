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
// source: VisibilityLabels.proto

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

type VisibilityLabelsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	VisLabel      []*VisibilityLabel     `protobuf:"bytes,1,rep,name=visLabel" json:"visLabel,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VisibilityLabelsRequest) Reset() {
	*x = VisibilityLabelsRequest{}
	mi := &file_VisibilityLabels_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VisibilityLabelsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VisibilityLabelsRequest) ProtoMessage() {}

func (x *VisibilityLabelsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_VisibilityLabels_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VisibilityLabelsRequest.ProtoReflect.Descriptor instead.
func (*VisibilityLabelsRequest) Descriptor() ([]byte, []int) {
	return file_VisibilityLabels_proto_rawDescGZIP(), []int{0}
}

func (x *VisibilityLabelsRequest) GetVisLabel() []*VisibilityLabel {
	if x != nil {
		return x.VisLabel
	}
	return nil
}

type VisibilityLabel struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Label         []byte                 `protobuf:"bytes,1,req,name=label" json:"label,omitempty"`
	Ordinal       *uint32                `protobuf:"varint,2,opt,name=ordinal" json:"ordinal,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VisibilityLabel) Reset() {
	*x = VisibilityLabel{}
	mi := &file_VisibilityLabels_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VisibilityLabel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VisibilityLabel) ProtoMessage() {}

func (x *VisibilityLabel) ProtoReflect() protoreflect.Message {
	mi := &file_VisibilityLabels_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VisibilityLabel.ProtoReflect.Descriptor instead.
func (*VisibilityLabel) Descriptor() ([]byte, []int) {
	return file_VisibilityLabels_proto_rawDescGZIP(), []int{1}
}

func (x *VisibilityLabel) GetLabel() []byte {
	if x != nil {
		return x.Label
	}
	return nil
}

func (x *VisibilityLabel) GetOrdinal() uint32 {
	if x != nil && x.Ordinal != nil {
		return *x.Ordinal
	}
	return 0
}

type VisibilityLabelsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Result        []*RegionActionResult  `protobuf:"bytes,1,rep,name=result" json:"result,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VisibilityLabelsResponse) Reset() {
	*x = VisibilityLabelsResponse{}
	mi := &file_VisibilityLabels_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VisibilityLabelsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VisibilityLabelsResponse) ProtoMessage() {}

func (x *VisibilityLabelsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_VisibilityLabels_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VisibilityLabelsResponse.ProtoReflect.Descriptor instead.
func (*VisibilityLabelsResponse) Descriptor() ([]byte, []int) {
	return file_VisibilityLabels_proto_rawDescGZIP(), []int{2}
}

func (x *VisibilityLabelsResponse) GetResult() []*RegionActionResult {
	if x != nil {
		return x.Result
	}
	return nil
}

type SetAuthsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          []byte                 `protobuf:"bytes,1,req,name=user" json:"user,omitempty"`
	Auth          [][]byte               `protobuf:"bytes,2,rep,name=auth" json:"auth,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetAuthsRequest) Reset() {
	*x = SetAuthsRequest{}
	mi := &file_VisibilityLabels_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetAuthsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetAuthsRequest) ProtoMessage() {}

func (x *SetAuthsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_VisibilityLabels_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetAuthsRequest.ProtoReflect.Descriptor instead.
func (*SetAuthsRequest) Descriptor() ([]byte, []int) {
	return file_VisibilityLabels_proto_rawDescGZIP(), []int{3}
}

func (x *SetAuthsRequest) GetUser() []byte {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *SetAuthsRequest) GetAuth() [][]byte {
	if x != nil {
		return x.Auth
	}
	return nil
}

type UserAuthorizations struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          []byte                 `protobuf:"bytes,1,req,name=user" json:"user,omitempty"`
	Auth          []uint32               `protobuf:"varint,2,rep,name=auth" json:"auth,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserAuthorizations) Reset() {
	*x = UserAuthorizations{}
	mi := &file_VisibilityLabels_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserAuthorizations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAuthorizations) ProtoMessage() {}

func (x *UserAuthorizations) ProtoReflect() protoreflect.Message {
	mi := &file_VisibilityLabels_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAuthorizations.ProtoReflect.Descriptor instead.
func (*UserAuthorizations) Descriptor() ([]byte, []int) {
	return file_VisibilityLabels_proto_rawDescGZIP(), []int{4}
}

func (x *UserAuthorizations) GetUser() []byte {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *UserAuthorizations) GetAuth() []uint32 {
	if x != nil {
		return x.Auth
	}
	return nil
}

type MultiUserAuthorizations struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserAuths     []*UserAuthorizations  `protobuf:"bytes,1,rep,name=userAuths" json:"userAuths,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MultiUserAuthorizations) Reset() {
	*x = MultiUserAuthorizations{}
	mi := &file_VisibilityLabels_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MultiUserAuthorizations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiUserAuthorizations) ProtoMessage() {}

func (x *MultiUserAuthorizations) ProtoReflect() protoreflect.Message {
	mi := &file_VisibilityLabels_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiUserAuthorizations.ProtoReflect.Descriptor instead.
func (*MultiUserAuthorizations) Descriptor() ([]byte, []int) {
	return file_VisibilityLabels_proto_rawDescGZIP(), []int{5}
}

func (x *MultiUserAuthorizations) GetUserAuths() []*UserAuthorizations {
	if x != nil {
		return x.UserAuths
	}
	return nil
}

type GetAuthsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          []byte                 `protobuf:"bytes,1,req,name=user" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAuthsRequest) Reset() {
	*x = GetAuthsRequest{}
	mi := &file_VisibilityLabels_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAuthsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthsRequest) ProtoMessage() {}

func (x *GetAuthsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_VisibilityLabels_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthsRequest.ProtoReflect.Descriptor instead.
func (*GetAuthsRequest) Descriptor() ([]byte, []int) {
	return file_VisibilityLabels_proto_rawDescGZIP(), []int{6}
}

func (x *GetAuthsRequest) GetUser() []byte {
	if x != nil {
		return x.User
	}
	return nil
}

type GetAuthsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          []byte                 `protobuf:"bytes,1,req,name=user" json:"user,omitempty"`
	Auth          [][]byte               `protobuf:"bytes,2,rep,name=auth" json:"auth,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAuthsResponse) Reset() {
	*x = GetAuthsResponse{}
	mi := &file_VisibilityLabels_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAuthsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthsResponse) ProtoMessage() {}

func (x *GetAuthsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_VisibilityLabels_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthsResponse.ProtoReflect.Descriptor instead.
func (*GetAuthsResponse) Descriptor() ([]byte, []int) {
	return file_VisibilityLabels_proto_rawDescGZIP(), []int{7}
}

func (x *GetAuthsResponse) GetUser() []byte {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *GetAuthsResponse) GetAuth() [][]byte {
	if x != nil {
		return x.Auth
	}
	return nil
}

type ListLabelsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Regex         *string                `protobuf:"bytes,1,opt,name=regex" json:"regex,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListLabelsRequest) Reset() {
	*x = ListLabelsRequest{}
	mi := &file_VisibilityLabels_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListLabelsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLabelsRequest) ProtoMessage() {}

func (x *ListLabelsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_VisibilityLabels_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLabelsRequest.ProtoReflect.Descriptor instead.
func (*ListLabelsRequest) Descriptor() ([]byte, []int) {
	return file_VisibilityLabels_proto_rawDescGZIP(), []int{8}
}

func (x *ListLabelsRequest) GetRegex() string {
	if x != nil && x.Regex != nil {
		return *x.Regex
	}
	return ""
}

type ListLabelsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Label         [][]byte               `protobuf:"bytes,1,rep,name=label" json:"label,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListLabelsResponse) Reset() {
	*x = ListLabelsResponse{}
	mi := &file_VisibilityLabels_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListLabelsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLabelsResponse) ProtoMessage() {}

func (x *ListLabelsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_VisibilityLabels_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLabelsResponse.ProtoReflect.Descriptor instead.
func (*ListLabelsResponse) Descriptor() ([]byte, []int) {
	return file_VisibilityLabels_proto_rawDescGZIP(), []int{9}
}

func (x *ListLabelsResponse) GetLabel() [][]byte {
	if x != nil {
		return x.Label
	}
	return nil
}

var File_VisibilityLabels_proto protoreflect.FileDescriptor

var file_VisibilityLabels_proto_rawDesc = string([]byte{
	0x0a, 0x16, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x0c, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4a, 0x0a, 0x17, 0x56, 0x69,
	0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x08, 0x76, 0x69, 0x73, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x69, 0x73,
	0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x08, 0x76, 0x69,
	0x73, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0x41, 0x0a, 0x0f, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0c, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12,
	0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x07, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x6c, 0x22, 0x4a, 0x0a, 0x18, 0x56, 0x69, 0x73,
	0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x39, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x02, 0x28, 0x0c, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x61, 0x75, 0x74, 0x68, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68,
	0x22, 0x3c, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x02, 0x28, 0x0c, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x75,
	0x74, 0x68, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x22, 0x4f,
	0x0a, 0x17, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x55, 0x73, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x34, 0x0a, 0x09, 0x75, 0x73, 0x65,
	0x72, 0x41, 0x75, 0x74, 0x68, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70,
	0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x73, 0x22,
	0x25, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0c,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x3a, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74,
	0x68, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0c, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x04, 0x61, 0x75,
	0x74, 0x68, 0x22, 0x29, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x67, 0x65, 0x78,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x67, 0x65, 0x78, 0x22, 0x2a, 0x0a,
	0x12, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0c, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x32, 0xd5, 0x02, 0x0a, 0x17, 0x56, 0x69,
	0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x09, 0x61, 0x64, 0x64, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x12, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a,
	0x08, 0x73, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x73, 0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x53,
	0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x70, 0x62, 0x2e, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0a,
	0x63, 0x6c, 0x65, 0x61, 0x72, 0x41, 0x75, 0x74, 0x68, 0x73, 0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e,
	0x53, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a,
	0x08, 0x67, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x73, 0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14,
	0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x6c, 0x69, 0x73, 0x74, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x5a, 0x0a, 0x31, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e,
	0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x68, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x68, 0x61,
	0x64, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x42, 0x16, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x48, 0x01,
	0x5a, 0x05, 0x2e, 0x2e, 0x2f, 0x70, 0x62, 0x88, 0x01, 0x01, 0xa0, 0x01, 0x01,
})

var (
	file_VisibilityLabels_proto_rawDescOnce sync.Once
	file_VisibilityLabels_proto_rawDescData []byte
)

func file_VisibilityLabels_proto_rawDescGZIP() []byte {
	file_VisibilityLabels_proto_rawDescOnce.Do(func() {
		file_VisibilityLabels_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_VisibilityLabels_proto_rawDesc), len(file_VisibilityLabels_proto_rawDesc)))
	})
	return file_VisibilityLabels_proto_rawDescData
}

var file_VisibilityLabels_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_VisibilityLabels_proto_goTypes = []any{
	(*VisibilityLabelsRequest)(nil),  // 0: pb.VisibilityLabelsRequest
	(*VisibilityLabel)(nil),          // 1: pb.VisibilityLabel
	(*VisibilityLabelsResponse)(nil), // 2: pb.VisibilityLabelsResponse
	(*SetAuthsRequest)(nil),          // 3: pb.SetAuthsRequest
	(*UserAuthorizations)(nil),       // 4: pb.UserAuthorizations
	(*MultiUserAuthorizations)(nil),  // 5: pb.MultiUserAuthorizations
	(*GetAuthsRequest)(nil),          // 6: pb.GetAuthsRequest
	(*GetAuthsResponse)(nil),         // 7: pb.GetAuthsResponse
	(*ListLabelsRequest)(nil),        // 8: pb.ListLabelsRequest
	(*ListLabelsResponse)(nil),       // 9: pb.ListLabelsResponse
	(*RegionActionResult)(nil),       // 10: pb.RegionActionResult
}
var file_VisibilityLabels_proto_depIdxs = []int32{
	1,  // 0: pb.VisibilityLabelsRequest.visLabel:type_name -> pb.VisibilityLabel
	10, // 1: pb.VisibilityLabelsResponse.result:type_name -> pb.RegionActionResult
	4,  // 2: pb.MultiUserAuthorizations.userAuths:type_name -> pb.UserAuthorizations
	0,  // 3: pb.VisibilityLabelsService.addLabels:input_type -> pb.VisibilityLabelsRequest
	3,  // 4: pb.VisibilityLabelsService.setAuths:input_type -> pb.SetAuthsRequest
	3,  // 5: pb.VisibilityLabelsService.clearAuths:input_type -> pb.SetAuthsRequest
	6,  // 6: pb.VisibilityLabelsService.getAuths:input_type -> pb.GetAuthsRequest
	8,  // 7: pb.VisibilityLabelsService.listLabels:input_type -> pb.ListLabelsRequest
	2,  // 8: pb.VisibilityLabelsService.addLabels:output_type -> pb.VisibilityLabelsResponse
	2,  // 9: pb.VisibilityLabelsService.setAuths:output_type -> pb.VisibilityLabelsResponse
	2,  // 10: pb.VisibilityLabelsService.clearAuths:output_type -> pb.VisibilityLabelsResponse
	7,  // 11: pb.VisibilityLabelsService.getAuths:output_type -> pb.GetAuthsResponse
	9,  // 12: pb.VisibilityLabelsService.listLabels:output_type -> pb.ListLabelsResponse
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_VisibilityLabels_proto_init() }
func file_VisibilityLabels_proto_init() {
	if File_VisibilityLabels_proto != nil {
		return
	}
	file_Client_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_VisibilityLabels_proto_rawDesc), len(file_VisibilityLabels_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_VisibilityLabels_proto_goTypes,
		DependencyIndexes: file_VisibilityLabels_proto_depIdxs,
		MessageInfos:      file_VisibilityLabels_proto_msgTypes,
	}.Build()
	File_VisibilityLabels_proto = out.File
	file_VisibilityLabels_proto_goTypes = nil
	file_VisibilityLabels_proto_depIdxs = nil
}
