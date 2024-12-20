// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.28.2
// source: api/rpc.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type KeyValue struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Key            []byte                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	CreateRevision int64                  `protobuf:"varint,2,opt,name=create_revision,json=createRevision,proto3" json:"create_revision,omitempty"`
	ModRevision    int64                  `protobuf:"varint,3,opt,name=mod_revision,json=modRevision,proto3" json:"mod_revision,omitempty"`
	Version        int64                  `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`
	Value          []byte                 `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"` //  int64 lease = 6;
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *KeyValue) Reset() {
	*x = KeyValue{}
	mi := &file_api_rpc_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KeyValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyValue) ProtoMessage() {}

func (x *KeyValue) ProtoReflect() protoreflect.Message {
	mi := &file_api_rpc_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyValue.ProtoReflect.Descriptor instead.
func (*KeyValue) Descriptor() ([]byte, []int) {
	return file_api_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *KeyValue) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *KeyValue) GetCreateRevision() int64 {
	if x != nil {
		return x.CreateRevision
	}
	return 0
}

func (x *KeyValue) GetModRevision() int64 {
	if x != nil {
		return x.ModRevision
	}
	return 0
}

func (x *KeyValue) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *KeyValue) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type RangeRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Key   []byte                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// bytes range_end = 2;
	// int64 limit = 3;
	Revision      int64 `protobuf:"varint,4,opt,name=revision,proto3" json:"revision,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RangeRequest) Reset() {
	*x = RangeRequest{}
	mi := &file_api_rpc_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RangeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RangeRequest) ProtoMessage() {}

func (x *RangeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_rpc_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RangeRequest.ProtoReflect.Descriptor instead.
func (*RangeRequest) Descriptor() ([]byte, []int) {
	return file_api_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *RangeRequest) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *RangeRequest) GetRevision() int64 {
	if x != nil {
		return x.Revision
	}
	return 0
}

type RangeResponse struct {
	state  protoimpl.MessageState `protogen:"open.v1"`
	Header *ResponseHeader        `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Kvs    []*KeyValue            `protobuf:"bytes,2,rep,name=kvs,proto3" json:"kvs,omitempty"`
	// bool more = 3;
	Count         int64 `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RangeResponse) Reset() {
	*x = RangeResponse{}
	mi := &file_api_rpc_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RangeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RangeResponse) ProtoMessage() {}

func (x *RangeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_rpc_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RangeResponse.ProtoReflect.Descriptor instead.
func (*RangeResponse) Descriptor() ([]byte, []int) {
	return file_api_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *RangeResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *RangeResponse) GetKvs() []*KeyValue {
	if x != nil {
		return x.Kvs
	}
	return nil
}

func (x *RangeResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type ResponseHeader struct {
	state     protoimpl.MessageState `protogen:"open.v1"`
	ClusterId uint64                 `protobuf:"varint,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	MemberId  uint64                 `protobuf:"varint,2,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
	// int64 revision = 3;
	RaftTerm      uint64 `protobuf:"varint,4,opt,name=raft_term,json=raftTerm,proto3" json:"raft_term,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ResponseHeader) Reset() {
	*x = ResponseHeader{}
	mi := &file_api_rpc_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResponseHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseHeader) ProtoMessage() {}

func (x *ResponseHeader) ProtoReflect() protoreflect.Message {
	mi := &file_api_rpc_proto_msgTypes[3]
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
	return file_api_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *ResponseHeader) GetClusterId() uint64 {
	if x != nil {
		return x.ClusterId
	}
	return 0
}

func (x *ResponseHeader) GetMemberId() uint64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *ResponseHeader) GetRaftTerm() uint64 {
	if x != nil {
		return x.RaftTerm
	}
	return 0
}

var File_api_rpc_proto protoreflect.FileDescriptor

var file_api_rpc_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x61, 0x70, 0x69, 0x70, 0x62, 0x22, 0x98, 0x01, 0x0a, 0x08, 0x4b, 0x65, 0x79, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f,
	0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x21,
	0x0a, 0x0c, 0x6d, 0x6f, 0x64, 0x5f, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6d, 0x6f, 0x64, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x3c, 0x0a, 0x0c, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x22,
	0x77, 0x0a, 0x0d, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2d, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x21, 0x0a, 0x03, 0x6b, 0x76, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61,
	0x70, 0x69, 0x70, 0x62, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x03, 0x6b,
	0x76, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x69, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x61, 0x66, 0x74, 0x5f, 0x74,
	0x65, 0x72, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x72, 0x61, 0x66, 0x74, 0x54,
	0x65, 0x72, 0x6d, 0x42, 0x08, 0x5a, 0x06, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_rpc_proto_rawDescOnce sync.Once
	file_api_rpc_proto_rawDescData = file_api_rpc_proto_rawDesc
)

func file_api_rpc_proto_rawDescGZIP() []byte {
	file_api_rpc_proto_rawDescOnce.Do(func() {
		file_api_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_rpc_proto_rawDescData)
	})
	return file_api_rpc_proto_rawDescData
}

var file_api_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_rpc_proto_goTypes = []any{
	(*KeyValue)(nil),       // 0: apipb.KeyValue
	(*RangeRequest)(nil),   // 1: apipb.RangeRequest
	(*RangeResponse)(nil),  // 2: apipb.RangeResponse
	(*ResponseHeader)(nil), // 3: apipb.ResponseHeader
}
var file_api_rpc_proto_depIdxs = []int32{
	3, // 0: apipb.RangeResponse.header:type_name -> apipb.ResponseHeader
	0, // 1: apipb.RangeResponse.kvs:type_name -> apipb.KeyValue
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_rpc_proto_init() }
func file_api_rpc_proto_init() {
	if File_api_rpc_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_rpc_proto_goTypes,
		DependencyIndexes: file_api_rpc_proto_depIdxs,
		MessageInfos:      file_api_rpc_proto_msgTypes,
	}.Build()
	File_api_rpc_proto = out.File
	file_api_rpc_proto_rawDesc = nil
	file_api_rpc_proto_goTypes = nil
	file_api_rpc_proto_depIdxs = nil
}
