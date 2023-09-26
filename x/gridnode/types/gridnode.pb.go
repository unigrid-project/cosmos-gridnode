// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: proto/cosmossdkgridnode/gridnode/gridnode.proto

package types

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// QueryDelegatedAmountRequest is the request type for the Query/DelegatedAmount RPC method.
type QueryDelegatedAmountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DelegatorAddress string `protobuf:"bytes,1,opt,name=delegator_address,json=delegatorAddress,proto3" json:"delegator_address,omitempty"`
}

func (x *QueryDelegatedAmountRequest) Reset() {
	*x = QueryDelegatedAmountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cosmossdkgridnode_gridnode_gridnode_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryDelegatedAmountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryDelegatedAmountRequest) ProtoMessage() {}

func (x *QueryDelegatedAmountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cosmossdkgridnode_gridnode_gridnode_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryDelegatedAmountRequest.ProtoReflect.Descriptor instead.
func (*QueryDelegatedAmountRequest) Descriptor() ([]byte, []int) {
	return file_proto_cosmossdkgridnode_gridnode_gridnode_proto_rawDescGZIP(), []int{0}
}

func (x *QueryDelegatedAmountRequest) GetDelegatorAddress() string {
	if x != nil {
		return x.DelegatorAddress
	}
	return ""
}

// QueryDelegatedAmountResponse is the response type for the Query/DelegatedAmount RPC method.
type QueryDelegatedAmountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount int64 `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *QueryDelegatedAmountResponse) Reset() {
	*x = QueryDelegatedAmountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cosmossdkgridnode_gridnode_gridnode_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryDelegatedAmountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryDelegatedAmountResponse) ProtoMessage() {}

func (x *QueryDelegatedAmountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cosmossdkgridnode_gridnode_gridnode_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryDelegatedAmountResponse.ProtoReflect.Descriptor instead.
func (*QueryDelegatedAmountResponse) Descriptor() ([]byte, []int) {
	return file_proto_cosmossdkgridnode_gridnode_gridnode_proto_rawDescGZIP(), []int{1}
}

func (x *QueryDelegatedAmountResponse) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type Delegation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DelegatorAddress string `protobuf:"bytes,1,opt,name=delegator_address,json=delegatorAddress,proto3" json:"delegator_address,omitempty"`
	Amount           int64  `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *Delegation) Reset() {
	*x = Delegation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cosmossdkgridnode_gridnode_gridnode_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Delegation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Delegation) ProtoMessage() {}

func (x *Delegation) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cosmossdkgridnode_gridnode_gridnode_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Delegation.ProtoReflect.Descriptor instead.
func (*Delegation) Descriptor() ([]byte, []int) {
	return file_proto_cosmossdkgridnode_gridnode_gridnode_proto_rawDescGZIP(), []int{2}
}

func (x *Delegation) GetDelegatorAddress() string {
	if x != nil {
		return x.DelegatorAddress
	}
	return ""
}

func (x *Delegation) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

// MsgGridnodeDelegate is the request type for the Msg/DelegateTokens RPC method.
type MsgGridnodeDelegate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DelegatorAddress string `protobuf:"bytes,1,opt,name=delegator_address,json=delegatorAddress,proto3" json:"delegator_address,omitempty"`
	Amount           int64  `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *MsgGridnodeDelegate) Reset() {
	*x = MsgGridnodeDelegate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cosmossdkgridnode_gridnode_gridnode_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgGridnodeDelegate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgGridnodeDelegate) ProtoMessage() {}

func (x *MsgGridnodeDelegate) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cosmossdkgridnode_gridnode_gridnode_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgGridnodeDelegate.ProtoReflect.Descriptor instead.
func (*MsgGridnodeDelegate) Descriptor() ([]byte, []int) {
	return file_proto_cosmossdkgridnode_gridnode_gridnode_proto_rawDescGZIP(), []int{3}
}

func (x *MsgGridnodeDelegate) GetDelegatorAddress() string {
	if x != nil {
		return x.DelegatorAddress
	}
	return ""
}

func (x *MsgGridnodeDelegate) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

// MsgGridnodeDelegateResponse is the response type for the Msg/DelegateTokens RPC method.
type MsgGridnodeDelegateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxHash string `protobuf:"bytes,1,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"` // Transaction hash of the delegated tokens operation
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`               // Status of the operation, e.g., "success" or "failure"
}

func (x *MsgGridnodeDelegateResponse) Reset() {
	*x = MsgGridnodeDelegateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cosmossdkgridnode_gridnode_gridnode_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgGridnodeDelegateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgGridnodeDelegateResponse) ProtoMessage() {}

func (x *MsgGridnodeDelegateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cosmossdkgridnode_gridnode_gridnode_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgGridnodeDelegateResponse.ProtoReflect.Descriptor instead.
func (*MsgGridnodeDelegateResponse) Descriptor() ([]byte, []int) {
	return file_proto_cosmossdkgridnode_gridnode_gridnode_proto_rawDescGZIP(), []int{4}
}

func (x *MsgGridnodeDelegateResponse) GetTxHash() string {
	if x != nil {
		return x.TxHash
	}
	return ""
}

func (x *MsgGridnodeDelegateResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type MsgGridnodeUndelegate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DelegatorAddress string `protobuf:"bytes,1,opt,name=delegator_address,json=delegatorAddress,proto3" json:"delegator_address,omitempty"`
	Amount           int64  `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *MsgGridnodeUndelegate) Reset() {
	*x = MsgGridnodeUndelegate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cosmossdkgridnode_gridnode_gridnode_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgGridnodeUndelegate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgGridnodeUndelegate) ProtoMessage() {}

func (x *MsgGridnodeUndelegate) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cosmossdkgridnode_gridnode_gridnode_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgGridnodeUndelegate.ProtoReflect.Descriptor instead.
func (*MsgGridnodeUndelegate) Descriptor() ([]byte, []int) {
	return file_proto_cosmossdkgridnode_gridnode_gridnode_proto_rawDescGZIP(), []int{5}
}

func (x *MsgGridnodeUndelegate) GetDelegatorAddress() string {
	if x != nil {
		return x.DelegatorAddress
	}
	return ""
}

func (x *MsgGridnodeUndelegate) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type MsgGridnodeUndelegateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxHash string `protobuf:"bytes,1,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"` // Transaction hash of the delegated tokens operation
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`               // Status of the operation, e.g., "success" or "failure"
}

func (x *MsgGridnodeUndelegateResponse) Reset() {
	*x = MsgGridnodeUndelegateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cosmossdkgridnode_gridnode_gridnode_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgGridnodeUndelegateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgGridnodeUndelegateResponse) ProtoMessage() {}

func (x *MsgGridnodeUndelegateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cosmossdkgridnode_gridnode_gridnode_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgGridnodeUndelegateResponse.ProtoReflect.Descriptor instead.
func (*MsgGridnodeUndelegateResponse) Descriptor() ([]byte, []int) {
	return file_proto_cosmossdkgridnode_gridnode_gridnode_proto_rawDescGZIP(), []int{6}
}

func (x *MsgGridnodeUndelegateResponse) GetTxHash() string {
	if x != nil {
		return x.TxHash
	}
	return ""
}

func (x *MsgGridnodeUndelegateResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_proto_cosmossdkgridnode_gridnode_gridnode_proto protoreflect.FileDescriptor

var file_proto_cosmossdkgridnode_gridnode_gridnode_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x73, 0x64,
	0x6b, 0x67, 0x72, 0x69, 0x64, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x67, 0x72, 0x69, 0x64, 0x6e, 0x6f,
	0x64, 0x65, 0x2f, 0x67, 0x72, 0x69, 0x64, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1a, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x73, 0x64, 0x6b, 0x67, 0x72, 0x69, 0x64,
	0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x67, 0x72, 0x69, 0x64, 0x6e, 0x6f, 0x64, 0x65, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4a, 0x0a, 0x1b, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x11, 0x64, 0x65,
	0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x36, 0x0a, 0x1c, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x51, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a,
	0x11, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61,
	0x74, 0x6f, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x5a, 0x0a, 0x13, 0x4d, 0x73, 0x67, 0x47, 0x72, 0x69, 0x64, 0x6e, 0x6f, 0x64,
	0x65, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x64, 0x65, 0x6c,
	0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x4e,
	0x0a, 0x1b, 0x4d, 0x73, 0x67, 0x47, 0x72, 0x69, 0x64, 0x6e, 0x6f, 0x64, 0x65, 0x44, 0x65, 0x6c,
	0x65, 0x67, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a,
	0x07, 0x74, 0x78, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x74, 0x78, 0x48, 0x61, 0x73, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x5c,
	0x0a, 0x15, 0x4d, 0x73, 0x67, 0x47, 0x72, 0x69, 0x64, 0x6e, 0x6f, 0x64, 0x65, 0x55, 0x6e, 0x64,
	0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x64, 0x65, 0x6c, 0x65, 0x67,
	0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x10, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x50, 0x0a, 0x1d,
	0x4d, 0x73, 0x67, 0x47, 0x72, 0x69, 0x64, 0x6e, 0x6f, 0x64, 0x65, 0x55, 0x6e, 0x64, 0x65, 0x6c,
	0x65, 0x67, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a,
	0x07, 0x74, 0x78, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x74, 0x78, 0x48, 0x61, 0x73, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xf5,
	0x02, 0x0a, 0x0d, 0x47, 0x72, 0x69, 0x64, 0x6e, 0x6f, 0x64, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x12, 0xe0, 0x01, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x37, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x73, 0x64, 0x6b,
	0x67, 0x72, 0x69, 0x64, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x67, 0x72, 0x69, 0x64, 0x6e, 0x6f, 0x64,
	0x65, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x73, 0x64, 0x6b, 0x67, 0x72, 0x69, 0x64, 0x6e, 0x6f, 0x64,
	0x65, 0x2e, 0x67, 0x72, 0x69, 0x64, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x54, 0x12,
	0x52, 0x2f, 0x75, 0x6e, 0x69, 0x67, 0x72, 0x69, 0x64, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x72, 0x69,
	0x64, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x67, 0x72, 0x69, 0x64, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x64,
	0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x2d, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x2f,
	0x7b, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x7d, 0x12, 0x80, 0x01, 0x0a, 0x10, 0x55, 0x6e, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x31, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x73, 0x64, 0x6b, 0x67, 0x72, 0x69, 0x64, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x67, 0x72, 0x69,
	0x64, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x4d, 0x73, 0x67, 0x47, 0x72, 0x69, 0x64, 0x6e, 0x6f, 0x64,
	0x65, 0x55, 0x6e, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x1a, 0x39, 0x2e, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x73, 0x64, 0x6b, 0x67, 0x72, 0x69, 0x64, 0x6e, 0x6f, 0x64, 0x65, 0x2e,
	0x67, 0x72, 0x69, 0x64, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x4d, 0x73, 0x67, 0x47, 0x72, 0x69, 0x64,
	0x6e, 0x6f, 0x64, 0x65, 0x55, 0x6e, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x89, 0x01, 0x0a, 0x0b, 0x47, 0x72, 0x69, 0x64, 0x6e,
	0x6f, 0x64, 0x65, 0x4d, 0x73, 0x67, 0x12, 0x7a, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x2f, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x73, 0x64, 0x6b, 0x67, 0x72, 0x69, 0x64, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x67, 0x72, 0x69,
	0x64, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x4d, 0x73, 0x67, 0x47, 0x72, 0x69, 0x64, 0x6e, 0x6f, 0x64,
	0x65, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x1a, 0x37, 0x2e, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x73, 0x64, 0x6b, 0x67, 0x72, 0x69, 0x64, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x67, 0x72,
	0x69, 0x64, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x4d, 0x73, 0x67, 0x47, 0x72, 0x69, 0x64, 0x6e, 0x6f,
	0x64, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x75, 0x6e, 0x69, 0x67, 0x72, 0x69, 0x64, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x72, 0x69, 0x64,
	0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x78, 0x2f, 0x67, 0x72, 0x69, 0x64, 0x6e, 0x6f, 0x64, 0x65, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_cosmossdkgridnode_gridnode_gridnode_proto_rawDescOnce sync.Once
	file_proto_cosmossdkgridnode_gridnode_gridnode_proto_rawDescData = file_proto_cosmossdkgridnode_gridnode_gridnode_proto_rawDesc
)

func file_proto_cosmossdkgridnode_gridnode_gridnode_proto_rawDescGZIP() []byte {
	file_proto_cosmossdkgridnode_gridnode_gridnode_proto_rawDescOnce.Do(func() {
		file_proto_cosmossdkgridnode_gridnode_gridnode_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_cosmossdkgridnode_gridnode_gridnode_proto_rawDescData)
	})
	return file_proto_cosmossdkgridnode_gridnode_gridnode_proto_rawDescData
}

var file_proto_cosmossdkgridnode_gridnode_gridnode_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_cosmossdkgridnode_gridnode_gridnode_proto_goTypes = []interface{}{
	(*QueryDelegatedAmountRequest)(nil),   // 0: cosmossdkgridnode.gridnode.QueryDelegatedAmountRequest
	(*QueryDelegatedAmountResponse)(nil),  // 1: cosmossdkgridnode.gridnode.QueryDelegatedAmountResponse
	(*Delegation)(nil),                    // 2: cosmossdkgridnode.gridnode.Delegation
	(*MsgGridnodeDelegate)(nil),           // 3: cosmossdkgridnode.gridnode.MsgGridnodeDelegate
	(*MsgGridnodeDelegateResponse)(nil),   // 4: cosmossdkgridnode.gridnode.MsgGridnodeDelegateResponse
	(*MsgGridnodeUndelegate)(nil),         // 5: cosmossdkgridnode.gridnode.MsgGridnodeUndelegate
	(*MsgGridnodeUndelegateResponse)(nil), // 6: cosmossdkgridnode.gridnode.MsgGridnodeUndelegateResponse
}
var file_proto_cosmossdkgridnode_gridnode_gridnode_proto_depIdxs = []int32{
	0, // 0: cosmossdkgridnode.gridnode.GridnodeQuery.DelegatedAmount:input_type -> cosmossdkgridnode.gridnode.QueryDelegatedAmountRequest
	5, // 1: cosmossdkgridnode.gridnode.GridnodeQuery.UndelegateTokens:input_type -> cosmossdkgridnode.gridnode.MsgGridnodeUndelegate
	3, // 2: cosmossdkgridnode.gridnode.GridnodeMsg.DelegateTokens:input_type -> cosmossdkgridnode.gridnode.MsgGridnodeDelegate
	1, // 3: cosmossdkgridnode.gridnode.GridnodeQuery.DelegatedAmount:output_type -> cosmossdkgridnode.gridnode.QueryDelegatedAmountResponse
	6, // 4: cosmossdkgridnode.gridnode.GridnodeQuery.UndelegateTokens:output_type -> cosmossdkgridnode.gridnode.MsgGridnodeUndelegateResponse
	4, // 5: cosmossdkgridnode.gridnode.GridnodeMsg.DelegateTokens:output_type -> cosmossdkgridnode.gridnode.MsgGridnodeDelegateResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_cosmossdkgridnode_gridnode_gridnode_proto_init() }
func file_proto_cosmossdkgridnode_gridnode_gridnode_proto_init() {
	if File_proto_cosmossdkgridnode_gridnode_gridnode_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_cosmossdkgridnode_gridnode_gridnode_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryDelegatedAmountRequest); i {
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
		file_proto_cosmossdkgridnode_gridnode_gridnode_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryDelegatedAmountResponse); i {
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
		file_proto_cosmossdkgridnode_gridnode_gridnode_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Delegation); i {
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
		file_proto_cosmossdkgridnode_gridnode_gridnode_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgGridnodeDelegate); i {
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
		file_proto_cosmossdkgridnode_gridnode_gridnode_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgGridnodeDelegateResponse); i {
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
		file_proto_cosmossdkgridnode_gridnode_gridnode_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgGridnodeUndelegate); i {
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
		file_proto_cosmossdkgridnode_gridnode_gridnode_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgGridnodeUndelegateResponse); i {
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
			RawDescriptor: file_proto_cosmossdkgridnode_gridnode_gridnode_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_proto_cosmossdkgridnode_gridnode_gridnode_proto_goTypes,
		DependencyIndexes: file_proto_cosmossdkgridnode_gridnode_gridnode_proto_depIdxs,
		MessageInfos:      file_proto_cosmossdkgridnode_gridnode_gridnode_proto_msgTypes,
	}.Build()
	File_proto_cosmossdkgridnode_gridnode_gridnode_proto = out.File
	file_proto_cosmossdkgridnode_gridnode_gridnode_proto_rawDesc = nil
	file_proto_cosmossdkgridnode_gridnode_gridnode_proto_goTypes = nil
	file_proto_cosmossdkgridnode_gridnode_gridnode_proto_depIdxs = nil
}
