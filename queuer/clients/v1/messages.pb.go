// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0-devel
// 	protoc        (unknown)
// source: queuer/clients/v1/messages.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	v1 "github.com/gsols/goproto/queuer/entities/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RegisterClientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Client *v1.Client `protobuf:"bytes,1,opt,name=client,proto3" json:"client,omitempty"`
}

func (x *RegisterClientRequest) Reset() {
	*x = RegisterClientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queuer_clients_v1_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterClientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterClientRequest) ProtoMessage() {}

func (x *RegisterClientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_queuer_clients_v1_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterClientRequest.ProtoReflect.Descriptor instead.
func (*RegisterClientRequest) Descriptor() ([]byte, []int) {
	return file_queuer_clients_v1_messages_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterClientRequest) GetClient() *v1.Client {
	if x != nil {
		return x.Client
	}
	return nil
}

type RegisterClientResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *RegisterClientResponse) Reset() {
	*x = RegisterClientResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queuer_clients_v1_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterClientResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterClientResponse) ProtoMessage() {}

func (x *RegisterClientResponse) ProtoReflect() protoreflect.Message {
	mi := &file_queuer_clients_v1_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterClientResponse.ProtoReflect.Descriptor instead.
func (*RegisterClientResponse) Descriptor() ([]byte, []int) {
	return file_queuer_clients_v1_messages_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterClientResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type PublishClientStatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId string    `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Stats    *v1.Stats `protobuf:"bytes,2,opt,name=stats,proto3" json:"stats,omitempty"`
}

func (x *PublishClientStatsRequest) Reset() {
	*x = PublishClientStatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queuer_clients_v1_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishClientStatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishClientStatsRequest) ProtoMessage() {}

func (x *PublishClientStatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_queuer_clients_v1_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishClientStatsRequest.ProtoReflect.Descriptor instead.
func (*PublishClientStatsRequest) Descriptor() ([]byte, []int) {
	return file_queuer_clients_v1_messages_proto_rawDescGZIP(), []int{2}
}

func (x *PublishClientStatsRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *PublishClientStatsRequest) GetStats() *v1.Stats {
	if x != nil {
		return x.Stats
	}
	return nil
}

type PublishClientStatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *PublishClientStatsResponse) Reset() {
	*x = PublishClientStatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queuer_clients_v1_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishClientStatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishClientStatsResponse) ProtoMessage() {}

func (x *PublishClientStatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_queuer_clients_v1_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishClientStatsResponse.ProtoReflect.Descriptor instead.
func (*PublishClientStatsResponse) Descriptor() ([]byte, []int) {
	return file_queuer_clients_v1_messages_proto_rawDescGZIP(), []int{3}
}

func (x *PublishClientStatsResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type SubscribeToCommandsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
}

func (x *SubscribeToCommandsRequest) Reset() {
	*x = SubscribeToCommandsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queuer_clients_v1_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeToCommandsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeToCommandsRequest) ProtoMessage() {}

func (x *SubscribeToCommandsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_queuer_clients_v1_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeToCommandsRequest.ProtoReflect.Descriptor instead.
func (*SubscribeToCommandsRequest) Descriptor() ([]byte, []int) {
	return file_queuer_clients_v1_messages_proto_rawDescGZIP(), []int{4}
}

func (x *SubscribeToCommandsRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

type SubscribeToCommandsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Command *v1.Command `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
}

func (x *SubscribeToCommandsResponse) Reset() {
	*x = SubscribeToCommandsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queuer_clients_v1_messages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeToCommandsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeToCommandsResponse) ProtoMessage() {}

func (x *SubscribeToCommandsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_queuer_clients_v1_messages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeToCommandsResponse.ProtoReflect.Descriptor instead.
func (*SubscribeToCommandsResponse) Descriptor() ([]byte, []int) {
	return file_queuer_clients_v1_messages_proto_rawDescGZIP(), []int{5}
}

func (x *SubscribeToCommandsResponse) GetCommand() *v1.Command {
	if x != nil {
		return x.Command
	}
	return nil
}

type AckCommandRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId  string           `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	CommandId string           `protobuf:"bytes,2,opt,name=command_id,json=commandId,proto3" json:"command_id,omitempty"`
	Status    v1.CommandStatus `protobuf:"varint,3,opt,name=status,proto3,enum=queuer.entities.v1.CommandStatus" json:"status,omitempty"`
}

func (x *AckCommandRequest) Reset() {
	*x = AckCommandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queuer_clients_v1_messages_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AckCommandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AckCommandRequest) ProtoMessage() {}

func (x *AckCommandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_queuer_clients_v1_messages_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AckCommandRequest.ProtoReflect.Descriptor instead.
func (*AckCommandRequest) Descriptor() ([]byte, []int) {
	return file_queuer_clients_v1_messages_proto_rawDescGZIP(), []int{6}
}

func (x *AckCommandRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *AckCommandRequest) GetCommandId() string {
	if x != nil {
		return x.CommandId
	}
	return ""
}

func (x *AckCommandRequest) GetStatus() v1.CommandStatus {
	if x != nil {
		return x.Status
	}
	return v1.CommandStatus(0)
}

type AckCommandResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *AckCommandResponse) Reset() {
	*x = AckCommandResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queuer_clients_v1_messages_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AckCommandResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AckCommandResponse) ProtoMessage() {}

func (x *AckCommandResponse) ProtoReflect() protoreflect.Message {
	mi := &file_queuer_clients_v1_messages_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AckCommandResponse.ProtoReflect.Descriptor instead.
func (*AckCommandResponse) Descriptor() ([]byte, []int) {
	return file_queuer_clients_v1_messages_proto_rawDescGZIP(), []int{7}
}

func (x *AckCommandResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_queuer_clients_v1_messages_proto protoreflect.FileDescriptor

var file_queuer_clients_v1_messages_proto_rawDesc = []byte{
	0x0a, 0x20, 0x71, 0x75, 0x65, 0x75, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x11, 0x71, 0x75, 0x65, 0x75, 0x65, 0x72, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x71, 0x75, 0x65,
	0x75, 0x65, 0x72, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x71, 0x75,
	0x65, 0x75, 0x65, 0x72, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x71, 0x75,
	0x65, 0x75, 0x65, 0x72, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20,
	0x71, 0x75, 0x65, 0x75, 0x65, 0x72, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x72, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x4b, 0x0a, 0x15, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x06, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x71, 0x75, 0x65, 0x75,
	0x65, 0x72, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x32, 0x0a,
	0x16, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x22, 0x69, 0x0a, 0x19, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x71, 0x75, 0x65,
	0x75, 0x65, 0x72, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x22, 0x36, 0x0a, 0x1a,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x22, 0x39, 0x0a, 0x1a, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x54, 0x6f, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22,
	0x54, 0x0a, 0x1b, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x6f, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35,
	0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x72, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x07, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x22, 0x8a, 0x01, 0x0a, 0x11, 0x41, 0x63, 0x6b, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x72,
	0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x2e, 0x0a, 0x12, 0x41, 0x63, 0x6b, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x73, 0x6f, 0x6c, 0x73, 0x2f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x71,
	0x75, 0x65, 0x75, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31,
	0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_queuer_clients_v1_messages_proto_rawDescOnce sync.Once
	file_queuer_clients_v1_messages_proto_rawDescData = file_queuer_clients_v1_messages_proto_rawDesc
)

func file_queuer_clients_v1_messages_proto_rawDescGZIP() []byte {
	file_queuer_clients_v1_messages_proto_rawDescOnce.Do(func() {
		file_queuer_clients_v1_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_queuer_clients_v1_messages_proto_rawDescData)
	})
	return file_queuer_clients_v1_messages_proto_rawDescData
}

var file_queuer_clients_v1_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_queuer_clients_v1_messages_proto_goTypes = []interface{}{
	(*RegisterClientRequest)(nil),       // 0: queuer.clients.v1.RegisterClientRequest
	(*RegisterClientResponse)(nil),      // 1: queuer.clients.v1.RegisterClientResponse
	(*PublishClientStatsRequest)(nil),   // 2: queuer.clients.v1.PublishClientStatsRequest
	(*PublishClientStatsResponse)(nil),  // 3: queuer.clients.v1.PublishClientStatsResponse
	(*SubscribeToCommandsRequest)(nil),  // 4: queuer.clients.v1.SubscribeToCommandsRequest
	(*SubscribeToCommandsResponse)(nil), // 5: queuer.clients.v1.SubscribeToCommandsResponse
	(*AckCommandRequest)(nil),           // 6: queuer.clients.v1.AckCommandRequest
	(*AckCommandResponse)(nil),          // 7: queuer.clients.v1.AckCommandResponse
	(*v1.Client)(nil),                   // 8: queuer.entities.v1.Client
	(*v1.Stats)(nil),                    // 9: queuer.entities.v1.Stats
	(*v1.Command)(nil),                  // 10: queuer.entities.v1.Command
	(v1.CommandStatus)(0),               // 11: queuer.entities.v1.CommandStatus
}
var file_queuer_clients_v1_messages_proto_depIdxs = []int32{
	8,  // 0: queuer.clients.v1.RegisterClientRequest.client:type_name -> queuer.entities.v1.Client
	9,  // 1: queuer.clients.v1.PublishClientStatsRequest.stats:type_name -> queuer.entities.v1.Stats
	10, // 2: queuer.clients.v1.SubscribeToCommandsResponse.command:type_name -> queuer.entities.v1.Command
	11, // 3: queuer.clients.v1.AckCommandRequest.status:type_name -> queuer.entities.v1.CommandStatus
	4,  // [4:4] is the sub-list for method output_type
	4,  // [4:4] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_queuer_clients_v1_messages_proto_init() }
func file_queuer_clients_v1_messages_proto_init() {
	if File_queuer_clients_v1_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_queuer_clients_v1_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterClientRequest); i {
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
		file_queuer_clients_v1_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterClientResponse); i {
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
		file_queuer_clients_v1_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishClientStatsRequest); i {
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
		file_queuer_clients_v1_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishClientStatsResponse); i {
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
		file_queuer_clients_v1_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeToCommandsRequest); i {
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
		file_queuer_clients_v1_messages_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeToCommandsResponse); i {
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
		file_queuer_clients_v1_messages_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AckCommandRequest); i {
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
		file_queuer_clients_v1_messages_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AckCommandResponse); i {
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
			RawDescriptor: file_queuer_clients_v1_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_queuer_clients_v1_messages_proto_goTypes,
		DependencyIndexes: file_queuer_clients_v1_messages_proto_depIdxs,
		MessageInfos:      file_queuer_clients_v1_messages_proto_msgTypes,
	}.Build()
	File_queuer_clients_v1_messages_proto = out.File
	file_queuer_clients_v1_messages_proto_rawDesc = nil
	file_queuer_clients_v1_messages_proto_goTypes = nil
	file_queuer_clients_v1_messages_proto_depIdxs = nil
}
