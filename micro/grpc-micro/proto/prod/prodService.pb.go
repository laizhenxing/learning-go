// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: prodService.proto

package prod

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

type ProdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"size",form:"size"
	Size int32 `protobuf:"varint,1,opt,name=size,proto3" json:"size" form:"size"`
	// @inject_tag: uri:"pid"
	ProdId int32 `protobuf:"varint,2,opt,name=prod_id,json=prodId,proto3" json:"prod_id,omitempty" uri:"pid"`
}

func (x *ProdRequest) Reset() {
	*x = ProdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProdRequest) ProtoMessage() {}

func (x *ProdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_prodService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProdRequest.ProtoReflect.Descriptor instead.
func (*ProdRequest) Descriptor() ([]byte, []int) {
	return file_prodService_proto_rawDescGZIP(), []int{0}
}

func (x *ProdRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ProdRequest) GetProdId() int32 {
	if x != nil {
		return x.ProdId
	}
	return 0
}

type ProdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*ProdModel `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ProdResponse) Reset() {
	*x = ProdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProdResponse) ProtoMessage() {}

func (x *ProdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_prodService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProdResponse.ProtoReflect.Descriptor instead.
func (*ProdResponse) Descriptor() ([]byte, []int) {
	return file_prodService_proto_rawDescGZIP(), []int{1}
}

func (x *ProdResponse) GetData() []*ProdModel {
	if x != nil {
		return x.Data
	}
	return nil
}

type ProdDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *ProdModel `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ProdDetailResponse) Reset() {
	*x = ProdDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProdDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProdDetailResponse) ProtoMessage() {}

func (x *ProdDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_prodService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProdDetailResponse.ProtoReflect.Descriptor instead.
func (*ProdDetailResponse) Descriptor() ([]byte, []int) {
	return file_prodService_proto_rawDescGZIP(), []int{2}
}

func (x *ProdDetailResponse) GetData() *ProdModel {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_prodService_proto protoreflect.FileDescriptor

var file_prodService_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x70, 0x72, 0x6f, 0x64, 0x1a, 0x0c, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3a, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x72,
	0x6f, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x72, 0x6f,
	0x64, 0x49, 0x64, 0x22, 0x33, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x39, 0x0a, 0x12, 0x50, 0x72, 0x6f, 0x64,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70,
	0x72, 0x6f, 0x64, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x32, 0x81, 0x01, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x2e, 0x50, 0x72, 0x6f,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0d, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x64, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f,
	0x64, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x70, 0x72, 0x6f, 0x64, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_prodService_proto_rawDescOnce sync.Once
	file_prodService_proto_rawDescData = file_prodService_proto_rawDesc
)

func file_prodService_proto_rawDescGZIP() []byte {
	file_prodService_proto_rawDescOnce.Do(func() {
		file_prodService_proto_rawDescData = protoimpl.X.CompressGZIP(file_prodService_proto_rawDescData)
	})
	return file_prodService_proto_rawDescData
}

var file_prodService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_prodService_proto_goTypes = []interface{}{
	(*ProdRequest)(nil),        // 0: prod.ProdRequest
	(*ProdResponse)(nil),       // 1: prod.ProdResponse
	(*ProdDetailResponse)(nil), // 2: prod.ProdDetailResponse
	(*ProdModel)(nil),          // 3: prod.ProdModel
}
var file_prodService_proto_depIdxs = []int32{
	3, // 0: prod.ProdResponse.data:type_name -> prod.ProdModel
	3, // 1: prod.ProdDetailResponse.data:type_name -> prod.ProdModel
	0, // 2: prod.ProdService.GetProdList:input_type -> prod.ProdRequest
	0, // 3: prod.ProdService.GetProdDetail:input_type -> prod.ProdRequest
	1, // 4: prod.ProdService.GetProdList:output_type -> prod.ProdResponse
	2, // 5: prod.ProdService.GetProdDetail:output_type -> prod.ProdDetailResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_prodService_proto_init() }
func file_prodService_proto_init() {
	if File_prodService_proto != nil {
		return
	}
	file_models_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_prodService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProdRequest); i {
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
		file_prodService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProdResponse); i {
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
		file_prodService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProdDetailResponse); i {
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
			RawDescriptor: file_prodService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_prodService_proto_goTypes,
		DependencyIndexes: file_prodService_proto_depIdxs,
		MessageInfos:      file_prodService_proto_msgTypes,
	}.Build()
	File_prodService_proto = out.File
	file_prodService_proto_rawDesc = nil
	file_prodService_proto_goTypes = nil
	file_prodService_proto_depIdxs = nil
}
