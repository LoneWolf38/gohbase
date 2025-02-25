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
// source: RPC.proto

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

// User Information proto.  Included in ConnectionHeader on connection setup
type UserInformation struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	EffectiveUser *string                `protobuf:"bytes,1,req,name=effective_user,json=effectiveUser" json:"effective_user,omitempty"`
	RealUser      *string                `protobuf:"bytes,2,opt,name=real_user,json=realUser" json:"real_user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserInformation) Reset() {
	*x = UserInformation{}
	mi := &file_RPC_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserInformation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInformation) ProtoMessage() {}

func (x *UserInformation) ProtoReflect() protoreflect.Message {
	mi := &file_RPC_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInformation.ProtoReflect.Descriptor instead.
func (*UserInformation) Descriptor() ([]byte, []int) {
	return file_RPC_proto_rawDescGZIP(), []int{0}
}

func (x *UserInformation) GetEffectiveUser() string {
	if x != nil && x.EffectiveUser != nil {
		return *x.EffectiveUser
	}
	return ""
}

func (x *UserInformation) GetRealUser() string {
	if x != nil && x.RealUser != nil {
		return *x.RealUser
	}
	return ""
}

// This is sent on connection setup after the connection preamble is sent.
type ConnectionHeader struct {
	state       protoimpl.MessageState `protogen:"open.v1"`
	UserInfo    *UserInformation       `protobuf:"bytes,1,opt,name=user_info,json=userInfo" json:"user_info,omitempty"`
	ServiceName *string                `protobuf:"bytes,2,opt,name=service_name,json=serviceName" json:"service_name,omitempty"`
	// Cell block codec we will use sending over optional cell blocks.  Server throws exception
	// if cannot deal.  Null means no codec'ing going on so we are pb all the time (SLOW!!!)
	CellBlockCodecClass *string `protobuf:"bytes,3,opt,name=cell_block_codec_class,json=cellBlockCodecClass" json:"cell_block_codec_class,omitempty"`
	// Compressor we will use if cell block is compressed.  Server will throw exception if not supported.
	// Class must implement hadoop's CompressionCodec Interface.  Can't compress if no codec.
	CellBlockCompressorClass *string      `protobuf:"bytes,4,opt,name=cell_block_compressor_class,json=cellBlockCompressorClass" json:"cell_block_compressor_class,omitempty"`
	VersionInfo              *VersionInfo `protobuf:"bytes,5,opt,name=version_info,json=versionInfo" json:"version_info,omitempty"`
	// the transformation for rpc AES encryption with Apache Commons Crypto
	RpcCryptoCipherTransformation *string          `protobuf:"bytes,6,opt,name=rpc_crypto_cipher_transformation,json=rpcCryptoCipherTransformation" json:"rpc_crypto_cipher_transformation,omitempty"`
	Attribute                     []*NameBytesPair `protobuf:"bytes,7,rep,name=attribute" json:"attribute,omitempty"`
	unknownFields                 protoimpl.UnknownFields
	sizeCache                     protoimpl.SizeCache
}

func (x *ConnectionHeader) Reset() {
	*x = ConnectionHeader{}
	mi := &file_RPC_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConnectionHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionHeader) ProtoMessage() {}

func (x *ConnectionHeader) ProtoReflect() protoreflect.Message {
	mi := &file_RPC_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionHeader.ProtoReflect.Descriptor instead.
func (*ConnectionHeader) Descriptor() ([]byte, []int) {
	return file_RPC_proto_rawDescGZIP(), []int{1}
}

func (x *ConnectionHeader) GetUserInfo() *UserInformation {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

func (x *ConnectionHeader) GetServiceName() string {
	if x != nil && x.ServiceName != nil {
		return *x.ServiceName
	}
	return ""
}

func (x *ConnectionHeader) GetCellBlockCodecClass() string {
	if x != nil && x.CellBlockCodecClass != nil {
		return *x.CellBlockCodecClass
	}
	return ""
}

func (x *ConnectionHeader) GetCellBlockCompressorClass() string {
	if x != nil && x.CellBlockCompressorClass != nil {
		return *x.CellBlockCompressorClass
	}
	return ""
}

func (x *ConnectionHeader) GetVersionInfo() *VersionInfo {
	if x != nil {
		return x.VersionInfo
	}
	return nil
}

func (x *ConnectionHeader) GetRpcCryptoCipherTransformation() string {
	if x != nil && x.RpcCryptoCipherTransformation != nil {
		return *x.RpcCryptoCipherTransformation
	}
	return ""
}

func (x *ConnectionHeader) GetAttribute() []*NameBytesPair {
	if x != nil {
		return x.Attribute
	}
	return nil
}

// This is sent by rpc server to negotiate the data if necessary
type ConnectionHeaderResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// To use Apache Commons Crypto, negotiate the metadata
	CryptoCipherMeta *CryptoCipherMeta `protobuf:"bytes,1,opt,name=crypto_cipher_meta,json=cryptoCipherMeta" json:"crypto_cipher_meta,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *ConnectionHeaderResponse) Reset() {
	*x = ConnectionHeaderResponse{}
	mi := &file_RPC_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConnectionHeaderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionHeaderResponse) ProtoMessage() {}

func (x *ConnectionHeaderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_RPC_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionHeaderResponse.ProtoReflect.Descriptor instead.
func (*ConnectionHeaderResponse) Descriptor() ([]byte, []int) {
	return file_RPC_proto_rawDescGZIP(), []int{2}
}

func (x *ConnectionHeaderResponse) GetCryptoCipherMeta() *CryptoCipherMeta {
	if x != nil {
		return x.CryptoCipherMeta
	}
	return nil
}

// Optional Cell block Message.  Included in client RequestHeader
type CellBlockMeta struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Length of the following cell block.  Could calculate it but convenient having it too hand.
	Length        *uint32 `protobuf:"varint,1,opt,name=length" json:"length,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CellBlockMeta) Reset() {
	*x = CellBlockMeta{}
	mi := &file_RPC_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CellBlockMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CellBlockMeta) ProtoMessage() {}

func (x *CellBlockMeta) ProtoReflect() protoreflect.Message {
	mi := &file_RPC_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CellBlockMeta.ProtoReflect.Descriptor instead.
func (*CellBlockMeta) Descriptor() ([]byte, []int) {
	return file_RPC_proto_rawDescGZIP(), []int{3}
}

func (x *CellBlockMeta) GetLength() uint32 {
	if x != nil && x.Length != nil {
		return *x.Length
	}
	return 0
}

// At the RPC layer, this message is used to carry
// the server side exception to the RPC client.
type ExceptionResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Class name of the exception thrown from the server
	ExceptionClassName *string `protobuf:"bytes,1,opt,name=exception_class_name,json=exceptionClassName" json:"exception_class_name,omitempty"`
	// Exception stack trace from the server side
	StackTrace *string `protobuf:"bytes,2,opt,name=stack_trace,json=stackTrace" json:"stack_trace,omitempty"`
	// Optional hostname.  Filled in for some exceptions such as region moved
	// where exception gives clue on where the region may have moved.
	Hostname *string `protobuf:"bytes,3,opt,name=hostname" json:"hostname,omitempty"`
	Port     *int32  `protobuf:"varint,4,opt,name=port" json:"port,omitempty"`
	// Set if we are NOT to retry on receipt of this exception
	DoNotRetry *bool `protobuf:"varint,5,opt,name=do_not_retry,json=doNotRetry" json:"do_not_retry,omitempty"`
	// Set true if the server was considered to be overloaded when exception was thrown
	ServerOverloaded *bool `protobuf:"varint,6,opt,name=server_overloaded,json=serverOverloaded" json:"server_overloaded,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *ExceptionResponse) Reset() {
	*x = ExceptionResponse{}
	mi := &file_RPC_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExceptionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExceptionResponse) ProtoMessage() {}

func (x *ExceptionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_RPC_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExceptionResponse.ProtoReflect.Descriptor instead.
func (*ExceptionResponse) Descriptor() ([]byte, []int) {
	return file_RPC_proto_rawDescGZIP(), []int{4}
}

func (x *ExceptionResponse) GetExceptionClassName() string {
	if x != nil && x.ExceptionClassName != nil {
		return *x.ExceptionClassName
	}
	return ""
}

func (x *ExceptionResponse) GetStackTrace() string {
	if x != nil && x.StackTrace != nil {
		return *x.StackTrace
	}
	return ""
}

func (x *ExceptionResponse) GetHostname() string {
	if x != nil && x.Hostname != nil {
		return *x.Hostname
	}
	return ""
}

func (x *ExceptionResponse) GetPort() int32 {
	if x != nil && x.Port != nil {
		return *x.Port
	}
	return 0
}

func (x *ExceptionResponse) GetDoNotRetry() bool {
	if x != nil && x.DoNotRetry != nil {
		return *x.DoNotRetry
	}
	return false
}

func (x *ExceptionResponse) GetServerOverloaded() bool {
	if x != nil && x.ServerOverloaded != nil {
		return *x.ServerOverloaded
	}
	return false
}

// *
// Cipher meta for Crypto
type CryptoCipherMeta struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Transformation *string                `protobuf:"bytes,1,req,name=transformation" json:"transformation,omitempty"`
	InKey          []byte                 `protobuf:"bytes,2,opt,name=inKey" json:"inKey,omitempty"`
	InIv           []byte                 `protobuf:"bytes,3,opt,name=inIv" json:"inIv,omitempty"`
	OutKey         []byte                 `protobuf:"bytes,4,opt,name=outKey" json:"outKey,omitempty"`
	OutIv          []byte                 `protobuf:"bytes,5,opt,name=outIv" json:"outIv,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *CryptoCipherMeta) Reset() {
	*x = CryptoCipherMeta{}
	mi := &file_RPC_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CryptoCipherMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CryptoCipherMeta) ProtoMessage() {}

func (x *CryptoCipherMeta) ProtoReflect() protoreflect.Message {
	mi := &file_RPC_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CryptoCipherMeta.ProtoReflect.Descriptor instead.
func (*CryptoCipherMeta) Descriptor() ([]byte, []int) {
	return file_RPC_proto_rawDescGZIP(), []int{5}
}

func (x *CryptoCipherMeta) GetTransformation() string {
	if x != nil && x.Transformation != nil {
		return *x.Transformation
	}
	return ""
}

func (x *CryptoCipherMeta) GetInKey() []byte {
	if x != nil {
		return x.InKey
	}
	return nil
}

func (x *CryptoCipherMeta) GetInIv() []byte {
	if x != nil {
		return x.InIv
	}
	return nil
}

func (x *CryptoCipherMeta) GetOutKey() []byte {
	if x != nil {
		return x.OutKey
	}
	return nil
}

func (x *CryptoCipherMeta) GetOutIv() []byte {
	if x != nil {
		return x.OutIv
	}
	return nil
}

// Header sent making a request.
type RequestHeader struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Monotonically increasing call_id to keep track of RPC requests and their response
	CallId     *uint32   `protobuf:"varint,1,opt,name=call_id,json=callId" json:"call_id,omitempty"`
	TraceInfo  *RPCTInfo `protobuf:"bytes,2,opt,name=trace_info,json=traceInfo" json:"trace_info,omitempty"`
	MethodName *string   `protobuf:"bytes,3,opt,name=method_name,json=methodName" json:"method_name,omitempty"`
	// If true, then a pb Message param follows.
	RequestParam *bool `protobuf:"varint,4,opt,name=request_param,json=requestParam" json:"request_param,omitempty"`
	// If present, then an encoded data block follows.
	CellBlockMeta *CellBlockMeta `protobuf:"bytes,5,opt,name=cell_block_meta,json=cellBlockMeta" json:"cell_block_meta,omitempty"`
	// 0 is NORMAL priority.  200 is HIGH.  If no priority, treat it as NORMAL.
	// See HConstants.
	Priority      *uint32          `protobuf:"varint,6,opt,name=priority" json:"priority,omitempty"`
	Timeout       *uint32          `protobuf:"varint,7,opt,name=timeout" json:"timeout,omitempty"`
	Attribute     []*NameBytesPair `protobuf:"bytes,8,rep,name=attribute" json:"attribute,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RequestHeader) Reset() {
	*x = RequestHeader{}
	mi := &file_RPC_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RequestHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestHeader) ProtoMessage() {}

func (x *RequestHeader) ProtoReflect() protoreflect.Message {
	mi := &file_RPC_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestHeader.ProtoReflect.Descriptor instead.
func (*RequestHeader) Descriptor() ([]byte, []int) {
	return file_RPC_proto_rawDescGZIP(), []int{6}
}

func (x *RequestHeader) GetCallId() uint32 {
	if x != nil && x.CallId != nil {
		return *x.CallId
	}
	return 0
}

func (x *RequestHeader) GetTraceInfo() *RPCTInfo {
	if x != nil {
		return x.TraceInfo
	}
	return nil
}

func (x *RequestHeader) GetMethodName() string {
	if x != nil && x.MethodName != nil {
		return *x.MethodName
	}
	return ""
}

func (x *RequestHeader) GetRequestParam() bool {
	if x != nil && x.RequestParam != nil {
		return *x.RequestParam
	}
	return false
}

func (x *RequestHeader) GetCellBlockMeta() *CellBlockMeta {
	if x != nil {
		return x.CellBlockMeta
	}
	return nil
}

func (x *RequestHeader) GetPriority() uint32 {
	if x != nil && x.Priority != nil {
		return *x.Priority
	}
	return 0
}

func (x *RequestHeader) GetTimeout() uint32 {
	if x != nil && x.Timeout != nil {
		return *x.Timeout
	}
	return 0
}

func (x *RequestHeader) GetAttribute() []*NameBytesPair {
	if x != nil {
		return x.Attribute
	}
	return nil
}

type ResponseHeader struct {
	state  protoimpl.MessageState `protogen:"open.v1"`
	CallId *uint32                `protobuf:"varint,1,opt,name=call_id,json=callId" json:"call_id,omitempty"`
	// If present, then request threw an exception and no response message (else we presume one)
	Exception *ExceptionResponse `protobuf:"bytes,2,opt,name=exception" json:"exception,omitempty"`
	// If present, then an encoded data block follows.
	CellBlockMeta *CellBlockMeta `protobuf:"bytes,3,opt,name=cell_block_meta,json=cellBlockMeta" json:"cell_block_meta,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ResponseHeader) Reset() {
	*x = ResponseHeader{}
	mi := &file_RPC_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResponseHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseHeader) ProtoMessage() {}

func (x *ResponseHeader) ProtoReflect() protoreflect.Message {
	mi := &file_RPC_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseHeader.ProtoReflect.Descriptor instead.
func (*ResponseHeader) Descriptor() ([]byte, []int) {
	return file_RPC_proto_rawDescGZIP(), []int{7}
}

func (x *ResponseHeader) GetCallId() uint32 {
	if x != nil && x.CallId != nil {
		return *x.CallId
	}
	return 0
}

func (x *ResponseHeader) GetException() *ExceptionResponse {
	if x != nil {
		return x.Exception
	}
	return nil
}

func (x *ResponseHeader) GetCellBlockMeta() *CellBlockMeta {
	if x != nil {
		return x.CellBlockMeta
	}
	return nil
}

type SecurityPreamableResponse struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	ServerPrincipal *string                `protobuf:"bytes,1,req,name=server_principal,json=serverPrincipal" json:"server_principal,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *SecurityPreamableResponse) Reset() {
	*x = SecurityPreamableResponse{}
	mi := &file_RPC_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SecurityPreamableResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecurityPreamableResponse) ProtoMessage() {}

func (x *SecurityPreamableResponse) ProtoReflect() protoreflect.Message {
	mi := &file_RPC_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecurityPreamableResponse.ProtoReflect.Descriptor instead.
func (*SecurityPreamableResponse) Descriptor() ([]byte, []int) {
	return file_RPC_proto_rawDescGZIP(), []int{8}
}

func (x *SecurityPreamableResponse) GetServerPrincipal() string {
	if x != nil && x.ServerPrincipal != nil {
		return *x.ServerPrincipal
	}
	return ""
}

var File_RPC_proto protoreflect.FileDescriptor

var file_RPC_proto_rawDesc = string([]byte{
	0x0a, 0x09, 0x52, 0x50, 0x43, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a,
	0x0b, 0x48, 0x42, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x54, 0x72,
	0x61, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x55, 0x0a, 0x0f, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25,
	0x0a, 0x0e, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x6c, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x61, 0x6c, 0x55, 0x73,
	0x65, 0x72, 0x22, 0x89, 0x03, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x62, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x16,
	0x63, 0x65, 0x6c, 0x6c, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x63,
	0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x63, 0x65,
	0x6c, 0x6c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x64, 0x65, 0x63, 0x43, 0x6c, 0x61, 0x73,
	0x73, 0x12, 0x3d, 0x0a, 0x1b, 0x63, 0x65, 0x6c, 0x6c, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f,
	0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x18, 0x63, 0x65, 0x6c, 0x6c, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x43, 0x6c, 0x61, 0x73, 0x73,
	0x12, 0x32, 0x0a, 0x0c, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x47, 0x0a, 0x20, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x5f, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1d,
	0x72, 0x70, 0x63, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a,
	0x09, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x50,
	0x61, 0x69, 0x72, 0x52, 0x09, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x22, 0x5e,
	0x0a, 0x18, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x12, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x5f, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x10, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x22, 0x27,
	0x0a, 0x0d, 0x43, 0x65, 0x6c, 0x6c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4d, 0x65, 0x74, 0x61, 0x12,
	0x16, 0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x22, 0xe5, 0x01, 0x0a, 0x11, 0x45, 0x78, 0x63, 0x65,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a,
	0x14, 0x65, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x65, 0x78, 0x63,
	0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x63, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74,
	0x12, 0x20, 0x0a, 0x0c, 0x64, 0x6f, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x72, 0x65, 0x74, 0x72, 0x79,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x64, 0x6f, 0x4e, 0x6f, 0x74, 0x52, 0x65, 0x74,
	0x72, 0x79, 0x12, 0x2b, 0x0a, 0x11, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6f, 0x76, 0x65,
	0x72, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x4f, 0x76, 0x65, 0x72, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x22,
	0x92, 0x01, 0x0a, 0x10, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72,
	0x4d, 0x65, 0x74, 0x61, 0x12, 0x26, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x69, 0x6e, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x69, 0x6e, 0x4b,
	0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x49, 0x76, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x69, 0x6e, 0x49, 0x76, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x4b, 0x65, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x6f, 0x75, 0x74, 0x49, 0x76, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x6f,
	0x75, 0x74, 0x49, 0x76, 0x22, 0xbd, 0x02, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x63, 0x61, 0x6c, 0x6c, 0x49, 0x64, 0x12,
	0x2b, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x54, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x09, 0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0b,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a,
	0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x12, 0x39, 0x0a, 0x0f, 0x63, 0x65, 0x6c, 0x6c, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x5f, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x65, 0x6c, 0x6c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x0d,
	0x63, 0x65, 0x6c, 0x6c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x12, 0x2f, 0x0a, 0x09, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x61, 0x6d, 0x65,
	0x42, 0x79, 0x74, 0x65, 0x73, 0x50, 0x61, 0x69, 0x72, 0x52, 0x09, 0x61, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x22, 0x99, 0x01, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x61, 0x6c, 0x6c, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x63, 0x61, 0x6c, 0x6c, 0x49, 0x64,
	0x12, 0x33, 0x0a, 0x09, 0x65, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x09, 0x65, 0x78, 0x63, 0x65,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0f, 0x63, 0x65, 0x6c, 0x6c, 0x5f, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x70, 0x62, 0x2e, 0x43, 0x65, 0x6c, 0x6c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4d, 0x65, 0x74,
	0x61, 0x52, 0x0d, 0x63, 0x65, 0x6c, 0x6c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4d, 0x65, 0x74, 0x61,
	0x22, 0x46, 0x0a, 0x19, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x50, 0x72, 0x65, 0x61,
	0x6d, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a,
	0x10, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61,
	0x6c, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x50,
	0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x42, 0x4a, 0x0a, 0x31, 0x6f, 0x72, 0x67, 0x2e,
	0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x68, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x73, 0x68, 0x61, 0x64, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x42, 0x09, 0x52,
	0x50, 0x43, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x48, 0x01, 0x5a, 0x05, 0x2e, 0x2e, 0x2f, 0x70,
	0x62, 0xa0, 0x01, 0x01,
})

var (
	file_RPC_proto_rawDescOnce sync.Once
	file_RPC_proto_rawDescData []byte
)

func file_RPC_proto_rawDescGZIP() []byte {
	file_RPC_proto_rawDescOnce.Do(func() {
		file_RPC_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_RPC_proto_rawDesc), len(file_RPC_proto_rawDesc)))
	})
	return file_RPC_proto_rawDescData
}

var file_RPC_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_RPC_proto_goTypes = []any{
	(*UserInformation)(nil),           // 0: pb.UserInformation
	(*ConnectionHeader)(nil),          // 1: pb.ConnectionHeader
	(*ConnectionHeaderResponse)(nil),  // 2: pb.ConnectionHeaderResponse
	(*CellBlockMeta)(nil),             // 3: pb.CellBlockMeta
	(*ExceptionResponse)(nil),         // 4: pb.ExceptionResponse
	(*CryptoCipherMeta)(nil),          // 5: pb.CryptoCipherMeta
	(*RequestHeader)(nil),             // 6: pb.RequestHeader
	(*ResponseHeader)(nil),            // 7: pb.ResponseHeader
	(*SecurityPreamableResponse)(nil), // 8: pb.SecurityPreamableResponse
	(*VersionInfo)(nil),               // 9: pb.VersionInfo
	(*NameBytesPair)(nil),             // 10: pb.NameBytesPair
	(*RPCTInfo)(nil),                  // 11: pb.RPCTInfo
}
var file_RPC_proto_depIdxs = []int32{
	0,  // 0: pb.ConnectionHeader.user_info:type_name -> pb.UserInformation
	9,  // 1: pb.ConnectionHeader.version_info:type_name -> pb.VersionInfo
	10, // 2: pb.ConnectionHeader.attribute:type_name -> pb.NameBytesPair
	5,  // 3: pb.ConnectionHeaderResponse.crypto_cipher_meta:type_name -> pb.CryptoCipherMeta
	11, // 4: pb.RequestHeader.trace_info:type_name -> pb.RPCTInfo
	3,  // 5: pb.RequestHeader.cell_block_meta:type_name -> pb.CellBlockMeta
	10, // 6: pb.RequestHeader.attribute:type_name -> pb.NameBytesPair
	4,  // 7: pb.ResponseHeader.exception:type_name -> pb.ExceptionResponse
	3,  // 8: pb.ResponseHeader.cell_block_meta:type_name -> pb.CellBlockMeta
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_RPC_proto_init() }
func file_RPC_proto_init() {
	if File_RPC_proto != nil {
		return
	}
	file_HBase_proto_init()
	file_Tracing_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_RPC_proto_rawDesc), len(file_RPC_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RPC_proto_goTypes,
		DependencyIndexes: file_RPC_proto_depIdxs,
		MessageInfos:      file_RPC_proto_msgTypes,
	}.Build()
	File_RPC_proto = out.File
	file_RPC_proto_goTypes = nil
	file_RPC_proto_depIdxs = nil
}
