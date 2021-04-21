// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: encore/engine/runtime/runtime.proto

package runtime

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type RecordTraceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TraceId []byte `protobuf:"bytes,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	Data    []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *RecordTraceRequest) Reset() {
	*x = RecordTraceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_encore_engine_runtime_runtime_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordTraceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordTraceRequest) ProtoMessage() {}

func (x *RecordTraceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_encore_engine_runtime_runtime_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordTraceRequest.ProtoReflect.Descriptor instead.
func (*RecordTraceRequest) Descriptor() ([]byte, []int) {
	return file_encore_engine_runtime_runtime_proto_rawDescGZIP(), []int{0}
}

func (x *RecordTraceRequest) GetTraceId() []byte {
	if x != nil {
		return x.TraceId
	}
	return nil
}

func (x *RecordTraceRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type RecordTraceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RecordTraceResponse) Reset() {
	*x = RecordTraceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_encore_engine_runtime_runtime_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordTraceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordTraceResponse) ProtoMessage() {}

func (x *RecordTraceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_encore_engine_runtime_runtime_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordTraceResponse.ProtoReflect.Descriptor instead.
func (*RecordTraceResponse) Descriptor() ([]byte, []int) {
	return file_encore_engine_runtime_runtime_proto_rawDescGZIP(), []int{1}
}

type SecretsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SecretsRequest) Reset() {
	*x = SecretsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_encore_engine_runtime_runtime_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecretsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecretsRequest) ProtoMessage() {}

func (x *SecretsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_encore_engine_runtime_runtime_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecretsRequest.ProtoReflect.Descriptor instead.
func (*SecretsRequest) Descriptor() ([]byte, []int) {
	return file_encore_engine_runtime_runtime_proto_rawDescGZIP(), []int{2}
}

type SecretsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Secrets map[string]string `protobuf:"bytes,1,rep,name=secrets,proto3" json:"secrets,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SecretsResponse) Reset() {
	*x = SecretsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_encore_engine_runtime_runtime_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecretsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecretsResponse) ProtoMessage() {}

func (x *SecretsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_encore_engine_runtime_runtime_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecretsResponse.ProtoReflect.Descriptor instead.
func (*SecretsResponse) Descriptor() ([]byte, []int) {
	return file_encore_engine_runtime_runtime_proto_rawDescGZIP(), []int{3}
}

func (x *SecretsResponse) GetSecrets() map[string]string {
	if x != nil {
		return x.Secrets
	}
	return nil
}

var File_encore_engine_runtime_runtime_proto protoreflect.FileDescriptor

var file_encore_engine_runtime_runtime_proto_rawDesc = []byte{
	0x0a, 0x23, 0x65, 0x6e, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f,
	0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x65, 0x6e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x65, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x43, 0x0a, 0x12,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x72, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x15, 0x0a, 0x13, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x72, 0x61, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10, 0x0a, 0x0e, 0x53, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x9c, 0x01, 0x0a, 0x0f, 0x53,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d,
	0x0a, 0x07, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x33, 0x2e, 0x65, 0x6e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e,
	0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x1a, 0x3a, 0x0a,
	0x0c, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0xc9, 0x01, 0x0a, 0x07, 0x52, 0x75,
	0x6e, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x64, 0x0a, 0x0b, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54,
	0x72, 0x61, 0x63, 0x65, 0x12, 0x29, 0x2e, 0x65, 0x6e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x65, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x54, 0x72, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2a, 0x2e, 0x65, 0x6e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e,
	0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x72,
	0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x58, 0x0a, 0x07, 0x53,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x12, 0x25, 0x2e, 0x65, 0x6e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x53,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e,
	0x65, 0x6e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x72, 0x75,
	0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x31, 0x5a, 0x2f, 0x65, 0x6e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x64, 0x65, 0x76, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x65, 0x6e, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_encore_engine_runtime_runtime_proto_rawDescOnce sync.Once
	file_encore_engine_runtime_runtime_proto_rawDescData = file_encore_engine_runtime_runtime_proto_rawDesc
)

func file_encore_engine_runtime_runtime_proto_rawDescGZIP() []byte {
	file_encore_engine_runtime_runtime_proto_rawDescOnce.Do(func() {
		file_encore_engine_runtime_runtime_proto_rawDescData = protoimpl.X.CompressGZIP(file_encore_engine_runtime_runtime_proto_rawDescData)
	})
	return file_encore_engine_runtime_runtime_proto_rawDescData
}

var file_encore_engine_runtime_runtime_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_encore_engine_runtime_runtime_proto_goTypes = []interface{}{
	(*RecordTraceRequest)(nil),  // 0: encore.engine.runtime.RecordTraceRequest
	(*RecordTraceResponse)(nil), // 1: encore.engine.runtime.RecordTraceResponse
	(*SecretsRequest)(nil),      // 2: encore.engine.runtime.SecretsRequest
	(*SecretsResponse)(nil),     // 3: encore.engine.runtime.SecretsResponse
	nil,                         // 4: encore.engine.runtime.SecretsResponse.SecretsEntry
}
var file_encore_engine_runtime_runtime_proto_depIdxs = []int32{
	4, // 0: encore.engine.runtime.SecretsResponse.secrets:type_name -> encore.engine.runtime.SecretsResponse.SecretsEntry
	0, // 1: encore.engine.runtime.Runtime.RecordTrace:input_type -> encore.engine.runtime.RecordTraceRequest
	2, // 2: encore.engine.runtime.Runtime.Secrets:input_type -> encore.engine.runtime.SecretsRequest
	1, // 3: encore.engine.runtime.Runtime.RecordTrace:output_type -> encore.engine.runtime.RecordTraceResponse
	3, // 4: encore.engine.runtime.Runtime.Secrets:output_type -> encore.engine.runtime.SecretsResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_encore_engine_runtime_runtime_proto_init() }
func file_encore_engine_runtime_runtime_proto_init() {
	if File_encore_engine_runtime_runtime_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_encore_engine_runtime_runtime_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordTraceRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_encore_engine_runtime_runtime_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordTraceResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_encore_engine_runtime_runtime_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecretsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_encore_engine_runtime_runtime_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecretsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_encore_engine_runtime_runtime_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_encore_engine_runtime_runtime_proto_goTypes,
		DependencyIndexes: file_encore_engine_runtime_runtime_proto_depIdxs,
		MessageInfos:      file_encore_engine_runtime_runtime_proto_msgTypes,
	}.Build()
	File_encore_engine_runtime_runtime_proto = out.File
	file_encore_engine_runtime_runtime_proto_rawDesc = nil
	file_encore_engine_runtime_runtime_proto_goTypes = nil
	file_encore_engine_runtime_runtime_proto_depIdxs = nil
}