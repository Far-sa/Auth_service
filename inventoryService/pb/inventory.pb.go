// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: pb/inventory.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UpdateStockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId string `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Quantity  int32  `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *UpdateStockRequest) Reset() {
	*x = UpdateStockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_inventory_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStockRequest) ProtoMessage() {}

func (x *UpdateStockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_inventory_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStockRequest.ProtoReflect.Descriptor instead.
func (*UpdateStockRequest) Descriptor() ([]byte, []int) {
	return file_pb_inventory_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateStockRequest) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *UpdateStockRequest) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type UpdateStockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *UpdateStockResponse) Reset() {
	*x = UpdateStockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_inventory_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStockResponse) ProtoMessage() {}

func (x *UpdateStockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_inventory_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStockResponse.ProtoReflect.Descriptor instead.
func (*UpdateStockResponse) Descriptor() ([]byte, []int) {
	return file_pb_inventory_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateStockResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type CheckStockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId string `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
}

func (x *CheckStockRequest) Reset() {
	*x = CheckStockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_inventory_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckStockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckStockRequest) ProtoMessage() {}

func (x *CheckStockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_inventory_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckStockRequest.ProtoReflect.Descriptor instead.
func (*CheckStockRequest) Descriptor() ([]byte, []int) {
	return file_pb_inventory_proto_rawDescGZIP(), []int{2}
}

func (x *CheckStockRequest) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

type CheckStockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId   string                 `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Quantity    int32                  `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	LastUpdated *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
}

func (x *CheckStockResponse) Reset() {
	*x = CheckStockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_inventory_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckStockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckStockResponse) ProtoMessage() {}

func (x *CheckStockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_inventory_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckStockResponse.ProtoReflect.Descriptor instead.
func (*CheckStockResponse) Descriptor() ([]byte, []int) {
	return file_pb_inventory_proto_rawDescGZIP(), []int{3}
}

func (x *CheckStockResponse) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *CheckStockResponse) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *CheckStockResponse) GetLastUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUpdated
	}
	return nil
}

var File_pb_inventory_proto protoreflect.FileDescriptor

var file_pb_inventory_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x62, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4f, 0x0a, 0x12, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x2f, 0x0a, 0x13, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x32, 0x0a, 0x11, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x22,
	0x8e, 0x01, 0x0a, 0x12, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x12, 0x3d, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x32, 0x8f, 0x01, 0x0a, 0x10, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53,
	0x74, 0x6f, 0x63, 0x6b, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70,
	0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x74,
	0x6f, 0x63, 0x6b, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x74,
	0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x16, 0x5a, 0x14, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_pb_inventory_proto_rawDescOnce sync.Once
	file_pb_inventory_proto_rawDescData = file_pb_inventory_proto_rawDesc
)

func file_pb_inventory_proto_rawDescGZIP() []byte {
	file_pb_inventory_proto_rawDescOnce.Do(func() {
		file_pb_inventory_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_inventory_proto_rawDescData)
	})
	return file_pb_inventory_proto_rawDescData
}

var file_pb_inventory_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pb_inventory_proto_goTypes = []interface{}{
	(*UpdateStockRequest)(nil),    // 0: pb.UpdateStockRequest
	(*UpdateStockResponse)(nil),   // 1: pb.UpdateStockResponse
	(*CheckStockRequest)(nil),     // 2: pb.CheckStockRequest
	(*CheckStockResponse)(nil),    // 3: pb.CheckStockResponse
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_pb_inventory_proto_depIdxs = []int32{
	4, // 0: pb.CheckStockResponse.last_updated:type_name -> google.protobuf.Timestamp
	0, // 1: pb.InventoryService.UpdateStock:input_type -> pb.UpdateStockRequest
	2, // 2: pb.InventoryService.CheckStock:input_type -> pb.CheckStockRequest
	1, // 3: pb.InventoryService.UpdateStock:output_type -> pb.UpdateStockResponse
	3, // 4: pb.InventoryService.CheckStock:output_type -> pb.CheckStockResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pb_inventory_proto_init() }
func file_pb_inventory_proto_init() {
	if File_pb_inventory_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_inventory_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStockRequest); i {
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
		file_pb_inventory_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStockResponse); i {
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
		file_pb_inventory_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckStockRequest); i {
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
		file_pb_inventory_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckStockResponse); i {
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
			RawDescriptor: file_pb_inventory_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_inventory_proto_goTypes,
		DependencyIndexes: file_pb_inventory_proto_depIdxs,
		MessageInfos:      file_pb_inventory_proto_msgTypes,
	}.Build()
	File_pb_inventory_proto = out.File
	file_pb_inventory_proto_rawDesc = nil
	file_pb_inventory_proto_goTypes = nil
	file_pb_inventory_proto_depIdxs = nil
}
