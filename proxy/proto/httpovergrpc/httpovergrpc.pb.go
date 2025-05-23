// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: httpovergrpc.proto

package httpovergrpc

import (
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Represents an HTTP 1.1 header.
type Header struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Values        []string               `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Header) Reset() {
	*x = Header{}
	mi := &file_httpovergrpc_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Header) ProtoMessage() {}

func (x *Header) ProtoReflect() protoreflect.Message {
	mi := &file_httpovergrpc_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Header.ProtoReflect.Descriptor instead.
func (*Header) Descriptor() ([]byte, []int) {
	return file_httpovergrpc_proto_rawDescGZIP(), []int{0}
}

func (x *Header) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Header) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

// Represents an HTTP 1.1 request.
type Request struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Method        string                 `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Url           string                 `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Headers       []*Header              `protobuf:"bytes,3,rep,name=headers,proto3" json:"headers,omitempty"`
	Body          []byte                 `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Request) Reset() {
	*x = Request{}
	mi := &file_httpovergrpc_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_httpovergrpc_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_httpovergrpc_proto_rawDescGZIP(), []int{1}
}

func (x *Request) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *Request) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Request) GetHeaders() []*Header {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *Request) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

// Represents an HTTP 1.1 response.
type Response struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        int32                  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Headers       []*Header              `protobuf:"bytes,2,rep,name=headers,proto3" json:"headers,omitempty"`
	Body          []byte                 `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Response) Reset() {
	*x = Response{}
	mi := &file_httpovergrpc_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_httpovergrpc_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_httpovergrpc_proto_rawDescGZIP(), []int{2}
}

func (x *Response) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Response) GetHeaders() []*Header {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *Response) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

var File_httpovergrpc_proto protoreflect.FileDescriptor

const file_httpovergrpc_proto_rawDesc = "" +
	"\n" +
	"\x12httpovergrpc.proto\x12\fhttpovergrpc\"2\n" +
	"\x06Header\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x16\n" +
	"\x06values\x18\x02 \x03(\tR\x06values\"w\n" +
	"\aRequest\x12\x16\n" +
	"\x06method\x18\x01 \x01(\tR\x06method\x12\x10\n" +
	"\x03url\x18\x02 \x01(\tR\x03url\x12.\n" +
	"\aheaders\x18\x03 \x03(\v2\x14.httpovergrpc.HeaderR\aheaders\x12\x12\n" +
	"\x04body\x18\x04 \x01(\fR\x04body\"f\n" +
	"\bResponse\x12\x16\n" +
	"\x06status\x18\x01 \x01(\x05R\x06status\x12.\n" +
	"\aheaders\x18\x02 \x03(\v2\x14.httpovergrpc.HeaderR\aheaders\x12\x12\n" +
	"\x04body\x18\x03 \x01(\fR\x04body2N\n" +
	"\fHTTPOverGRPC\x12>\n" +
	"\vHTTPRequest\x12\x15.httpovergrpc.Request\x1a\x16.httpovergrpc.Response\"\x00B2Z0github.com/openconfig/ondatra/proto/httpovergrpcb\x06proto3"

var (
	file_httpovergrpc_proto_rawDescOnce sync.Once
	file_httpovergrpc_proto_rawDescData []byte
)

func file_httpovergrpc_proto_rawDescGZIP() []byte {
	file_httpovergrpc_proto_rawDescOnce.Do(func() {
		file_httpovergrpc_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_httpovergrpc_proto_rawDesc), len(file_httpovergrpc_proto_rawDesc)))
	})
	return file_httpovergrpc_proto_rawDescData
}

var file_httpovergrpc_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_httpovergrpc_proto_goTypes = []any{
	(*Header)(nil),   // 0: httpovergrpc.Header
	(*Request)(nil),  // 1: httpovergrpc.Request
	(*Response)(nil), // 2: httpovergrpc.Response
}
var file_httpovergrpc_proto_depIdxs = []int32{
	0, // 0: httpovergrpc.Request.headers:type_name -> httpovergrpc.Header
	0, // 1: httpovergrpc.Response.headers:type_name -> httpovergrpc.Header
	1, // 2: httpovergrpc.HTTPOverGRPC.HTTPRequest:input_type -> httpovergrpc.Request
	2, // 3: httpovergrpc.HTTPOverGRPC.HTTPRequest:output_type -> httpovergrpc.Response
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_httpovergrpc_proto_init() }
func file_httpovergrpc_proto_init() {
	if File_httpovergrpc_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_httpovergrpc_proto_rawDesc), len(file_httpovergrpc_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_httpovergrpc_proto_goTypes,
		DependencyIndexes: file_httpovergrpc_proto_depIdxs,
		MessageInfos:      file_httpovergrpc_proto_msgTypes,
	}.Build()
	File_httpovergrpc_proto = out.File
	file_httpovergrpc_proto_goTypes = nil
	file_httpovergrpc_proto_depIdxs = nil
}
