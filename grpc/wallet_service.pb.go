// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: wallet_service.proto

package grpc

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

type GetBalanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetBalanceRequest) Reset() {
	*x = GetBalanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBalanceRequest) ProtoMessage() {}

func (x *GetBalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBalanceRequest.ProtoReflect.Descriptor instead.
func (*GetBalanceRequest) Descriptor() ([]byte, []int) {
	return file_wallet_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetBalanceRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetBalanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Balance int64 `protobuf:"varint,1,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *GetBalanceResponse) Reset() {
	*x = GetBalanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBalanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBalanceResponse) ProtoMessage() {}

func (x *GetBalanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBalanceResponse.ProtoReflect.Descriptor instead.
func (*GetBalanceResponse) Descriptor() ([]byte, []int) {
	return file_wallet_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetBalanceResponse) GetBalance() int64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

type AddMoneyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Amount int64 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *AddMoneyRequest) Reset() {
	*x = AddMoneyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddMoneyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddMoneyRequest) ProtoMessage() {}

func (x *AddMoneyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddMoneyRequest.ProtoReflect.Descriptor instead.
func (*AddMoneyRequest) Descriptor() ([]byte, []int) {
	return file_wallet_service_proto_rawDescGZIP(), []int{2}
}

func (x *AddMoneyRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AddMoneyRequest) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type AddMoneyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReferenceId int64 `protobuf:"varint,1,opt,name=reference_id,json=referenceId,proto3" json:"reference_id,omitempty"`
}

func (x *AddMoneyResponse) Reset() {
	*x = AddMoneyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddMoneyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddMoneyResponse) ProtoMessage() {}

func (x *AddMoneyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddMoneyResponse.ProtoReflect.Descriptor instead.
func (*AddMoneyResponse) Descriptor() ([]byte, []int) {
	return file_wallet_service_proto_rawDescGZIP(), []int{3}
}

func (x *AddMoneyResponse) GetReferenceId() int64 {
	if x != nil {
		return x.ReferenceId
	}
	return 0
}

var File_wallet_service_proto protoreflect.FileDescriptor

var file_wallet_service_proto_rawDesc = []byte{
	0x0a, 0x14, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x64, 0x6f, 0x63, 0x6b, 0x69, 0x2e, 0x77, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x22, 0x2c, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x2e, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x22, 0x42, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x35, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x4d, 0x6f, 0x6e,
	0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0b, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x32, 0xaf, 0x01,
	0x0a, 0x0d, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x51, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1f, 0x2e,
	0x64, 0x6f, 0x63, 0x6b, 0x69, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x47, 0x65, 0x74,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20,
	0x2e, 0x64, 0x6f, 0x63, 0x6b, 0x69, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x47, 0x65,
	0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x4b, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x12, 0x1d,
	0x2e, 0x64, 0x6f, 0x63, 0x6b, 0x69, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x41, 0x64,
	0x64, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x64, 0x6f, 0x63, 0x6b, 0x69, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x41, 0x64, 0x64,
	0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_wallet_service_proto_rawDescOnce sync.Once
	file_wallet_service_proto_rawDescData = file_wallet_service_proto_rawDesc
)

func file_wallet_service_proto_rawDescGZIP() []byte {
	file_wallet_service_proto_rawDescOnce.Do(func() {
		file_wallet_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_wallet_service_proto_rawDescData)
	})
	return file_wallet_service_proto_rawDescData
}

var file_wallet_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_wallet_service_proto_goTypes = []interface{}{
	(*GetBalanceRequest)(nil),  // 0: docki.wallet.GetBalanceRequest
	(*GetBalanceResponse)(nil), // 1: docki.wallet.GetBalanceResponse
	(*AddMoneyRequest)(nil),    // 2: docki.wallet.AddMoneyRequest
	(*AddMoneyResponse)(nil),   // 3: docki.wallet.AddMoneyResponse
}
var file_wallet_service_proto_depIdxs = []int32{
	0, // 0: docki.wallet.WalletService.GetBalance:input_type -> docki.wallet.GetBalanceRequest
	2, // 1: docki.wallet.WalletService.AddMoney:input_type -> docki.wallet.AddMoneyRequest
	1, // 2: docki.wallet.WalletService.GetBalance:output_type -> docki.wallet.GetBalanceResponse
	3, // 3: docki.wallet.WalletService.AddMoney:output_type -> docki.wallet.AddMoneyResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_wallet_service_proto_init() }
func file_wallet_service_proto_init() {
	if File_wallet_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wallet_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBalanceRequest); i {
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
		file_wallet_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBalanceResponse); i {
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
		file_wallet_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddMoneyRequest); i {
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
		file_wallet_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddMoneyResponse); i {
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
			RawDescriptor: file_wallet_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wallet_service_proto_goTypes,
		DependencyIndexes: file_wallet_service_proto_depIdxs,
		MessageInfos:      file_wallet_service_proto_msgTypes,
	}.Build()
	File_wallet_service_proto = out.File
	file_wallet_service_proto_rawDesc = nil
	file_wallet_service_proto_goTypes = nil
	file_wallet_service_proto_depIdxs = nil
}
