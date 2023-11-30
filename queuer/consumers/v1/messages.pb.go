// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0-devel
// 	protoc        (unknown)
// source: queuer/consumers/v1/messages.proto

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

type RegisterConsumerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Consumer *v1.Consumer `protobuf:"bytes,1,opt,name=consumer,proto3" json:"consumer,omitempty"`
	OwnerId  string       `protobuf:"bytes,2,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
}

func (x *RegisterConsumerRequest) Reset() {
	*x = RegisterConsumerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queuer_consumers_v1_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterConsumerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterConsumerRequest) ProtoMessage() {}

func (x *RegisterConsumerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_queuer_consumers_v1_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterConsumerRequest.ProtoReflect.Descriptor instead.
func (*RegisterConsumerRequest) Descriptor() ([]byte, []int) {
	return file_queuer_consumers_v1_messages_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterConsumerRequest) GetConsumer() *v1.Consumer {
	if x != nil {
		return x.Consumer
	}
	return nil
}

func (x *RegisterConsumerRequest) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

type RegisterConsumerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *RegisterConsumerResponse) Reset() {
	*x = RegisterConsumerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queuer_consumers_v1_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterConsumerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterConsumerResponse) ProtoMessage() {}

func (x *RegisterConsumerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_queuer_consumers_v1_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterConsumerResponse.ProtoReflect.Descriptor instead.
func (*RegisterConsumerResponse) Descriptor() ([]byte, []int) {
	return file_queuer_consumers_v1_messages_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterConsumerResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type GetConsumersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetConsumersRequest) Reset() {
	*x = GetConsumersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queuer_consumers_v1_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConsumersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConsumersRequest) ProtoMessage() {}

func (x *GetConsumersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_queuer_consumers_v1_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConsumersRequest.ProtoReflect.Descriptor instead.
func (*GetConsumersRequest) Descriptor() ([]byte, []int) {
	return file_queuer_consumers_v1_messages_proto_rawDescGZIP(), []int{2}
}

type GetConsumersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Consumers []*v1.Consumer `protobuf:"bytes,1,rep,name=consumers,proto3" json:"consumers,omitempty"`
}

func (x *GetConsumersResponse) Reset() {
	*x = GetConsumersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queuer_consumers_v1_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConsumersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConsumersResponse) ProtoMessage() {}

func (x *GetConsumersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_queuer_consumers_v1_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConsumersResponse.ProtoReflect.Descriptor instead.
func (*GetConsumersResponse) Descriptor() ([]byte, []int) {
	return file_queuer_consumers_v1_messages_proto_rawDescGZIP(), []int{3}
}

func (x *GetConsumersResponse) GetConsumers() []*v1.Consumer {
	if x != nil {
		return x.Consumers
	}
	return nil
}

type PublishConsumerStatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsumerId string    `protobuf:"bytes,1,opt,name=consumer_id,json=consumerId,proto3" json:"consumer_id,omitempty"`
	Stats      *v1.Stats `protobuf:"bytes,2,opt,name=stats,proto3" json:"stats,omitempty"`
}

func (x *PublishConsumerStatsRequest) Reset() {
	*x = PublishConsumerStatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queuer_consumers_v1_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishConsumerStatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishConsumerStatsRequest) ProtoMessage() {}

func (x *PublishConsumerStatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_queuer_consumers_v1_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishConsumerStatsRequest.ProtoReflect.Descriptor instead.
func (*PublishConsumerStatsRequest) Descriptor() ([]byte, []int) {
	return file_queuer_consumers_v1_messages_proto_rawDescGZIP(), []int{4}
}

func (x *PublishConsumerStatsRequest) GetConsumerId() string {
	if x != nil {
		return x.ConsumerId
	}
	return ""
}

func (x *PublishConsumerStatsRequest) GetStats() *v1.Stats {
	if x != nil {
		return x.Stats
	}
	return nil
}

type PublishConsumerStatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *PublishConsumerStatsResponse) Reset() {
	*x = PublishConsumerStatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queuer_consumers_v1_messages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishConsumerStatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishConsumerStatsResponse) ProtoMessage() {}

func (x *PublishConsumerStatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_queuer_consumers_v1_messages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishConsumerStatsResponse.ProtoReflect.Descriptor instead.
func (*PublishConsumerStatsResponse) Descriptor() ([]byte, []int) {
	return file_queuer_consumers_v1_messages_proto_rawDescGZIP(), []int{5}
}

func (x *PublishConsumerStatsResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type SubscribeToCommandsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsumerId string `protobuf:"bytes,1,opt,name=consumer_id,json=consumerId,proto3" json:"consumer_id,omitempty"`
}

func (x *SubscribeToCommandsRequest) Reset() {
	*x = SubscribeToCommandsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queuer_consumers_v1_messages_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeToCommandsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeToCommandsRequest) ProtoMessage() {}

func (x *SubscribeToCommandsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_queuer_consumers_v1_messages_proto_msgTypes[6]
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
	return file_queuer_consumers_v1_messages_proto_rawDescGZIP(), []int{6}
}

func (x *SubscribeToCommandsRequest) GetConsumerId() string {
	if x != nil {
		return x.ConsumerId
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
		mi := &file_queuer_consumers_v1_messages_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeToCommandsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeToCommandsResponse) ProtoMessage() {}

func (x *SubscribeToCommandsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_queuer_consumers_v1_messages_proto_msgTypes[7]
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
	return file_queuer_consumers_v1_messages_proto_rawDescGZIP(), []int{7}
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

	ConsumerId string           `protobuf:"bytes,1,opt,name=consumer_id,json=consumerId,proto3" json:"consumer_id,omitempty"`
	CommandId  string           `protobuf:"bytes,2,opt,name=command_id,json=commandId,proto3" json:"command_id,omitempty"`
	Status     v1.CommandStatus `protobuf:"varint,3,opt,name=status,proto3,enum=queuer.entities.v1.CommandStatus" json:"status,omitempty"`
}

func (x *AckCommandRequest) Reset() {
	*x = AckCommandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queuer_consumers_v1_messages_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AckCommandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AckCommandRequest) ProtoMessage() {}

func (x *AckCommandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_queuer_consumers_v1_messages_proto_msgTypes[8]
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
	return file_queuer_consumers_v1_messages_proto_rawDescGZIP(), []int{8}
}

func (x *AckCommandRequest) GetConsumerId() string {
	if x != nil {
		return x.ConsumerId
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
		mi := &file_queuer_consumers_v1_messages_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AckCommandResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AckCommandResponse) ProtoMessage() {}

func (x *AckCommandResponse) ProtoReflect() protoreflect.Message {
	mi := &file_queuer_consumers_v1_messages_proto_msgTypes[9]
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
	return file_queuer_consumers_v1_messages_proto_rawDescGZIP(), []int{9}
}

func (x *AckCommandResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type GetSubscribedStreamsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsumerId string `protobuf:"bytes,1,opt,name=consumer_id,json=consumerId,proto3" json:"consumer_id,omitempty"`
}

func (x *GetSubscribedStreamsRequest) Reset() {
	*x = GetSubscribedStreamsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queuer_consumers_v1_messages_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSubscribedStreamsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSubscribedStreamsRequest) ProtoMessage() {}

func (x *GetSubscribedStreamsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_queuer_consumers_v1_messages_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSubscribedStreamsRequest.ProtoReflect.Descriptor instead.
func (*GetSubscribedStreamsRequest) Descriptor() ([]byte, []int) {
	return file_queuer_consumers_v1_messages_proto_rawDescGZIP(), []int{10}
}

func (x *GetSubscribedStreamsRequest) GetConsumerId() string {
	if x != nil {
		return x.ConsumerId
	}
	return ""
}

type GetSubscribedStreamsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Streams []*v1.Stream `protobuf:"bytes,1,rep,name=streams,proto3" json:"streams,omitempty"`
}

func (x *GetSubscribedStreamsResponse) Reset() {
	*x = GetSubscribedStreamsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queuer_consumers_v1_messages_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSubscribedStreamsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSubscribedStreamsResponse) ProtoMessage() {}

func (x *GetSubscribedStreamsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_queuer_consumers_v1_messages_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSubscribedStreamsResponse.ProtoReflect.Descriptor instead.
func (*GetSubscribedStreamsResponse) Descriptor() ([]byte, []int) {
	return file_queuer_consumers_v1_messages_proto_rawDescGZIP(), []int{11}
}

func (x *GetSubscribedStreamsResponse) GetStreams() []*v1.Stream {
	if x != nil {
		return x.Streams
	}
	return nil
}

var File_queuer_consumers_v1_messages_proto protoreflect.FileDescriptor

var file_queuer_consumers_v1_messages_proto_rawDesc = []byte{
	0x0a, 0x22, 0x71, 0x75, 0x65, 0x75, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x72, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x71, 0x75, 0x65, 0x75, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6e,
	0x73, 0x75, 0x6d, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x21, 0x71, 0x75, 0x65, 0x75, 0x65, 0x72, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x72, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x71, 0x75, 0x65, 0x75, 0x65, 0x72, 0x2f, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x71, 0x75, 0x65, 0x75, 0x65, 0x72, 0x2f, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x72, 0x2f,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x78, 0x0a, 0x17, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x38, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x72, 0x2e, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d,
	0x65, 0x72, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x08,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08,
	0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x34, 0x0a, 0x18, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e,
	0x73, 0x75, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x15, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x52,
	0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d,
	0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x71, 0x75, 0x65, 0x75,
	0x65, 0x72, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x72, 0x73, 0x22, 0x79, 0x0a, 0x1b, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x43, 0x6f, 0x6e,
	0x73, 0x75, 0x6d, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x29, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01,
	0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x71, 0x75,
	0x65, 0x75, 0x65, 0x72, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x22, 0x38, 0x0a,
	0x1c, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x47, 0x0a, 0x1a, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x54, 0x6f, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72,
	0x03, 0xb0, 0x01, 0x01, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x54, 0x0a, 0x1b, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x6f, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x35, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x72, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x22, 0xa2, 0x01, 0x0a, 0x11, 0x41, 0x63, 0x6b, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x0b,
	0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x0a, 0x63, 0x6f, 0x6e,
	0x73, 0x75, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05,
	0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x49, 0x64,
	0x12, 0x39, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x21, 0x2e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x72, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x2e, 0x0a, 0x12, 0x41,
	0x63, 0x6b, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x48, 0x0a, 0x1b, 0x47,
	0x65, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x64, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x0b, 0x63, 0x6f,
	0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x73, 0x75,
	0x6d, 0x65, 0x72, 0x49, 0x64, 0x22, 0x54, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x72, 0x2e,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x52, 0x07, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x42, 0x31, 0x5a, 0x2f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x73, 0x6f, 0x6c, 0x73, 0x2f,
	0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x72, 0x2f, 0x63,
	0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_queuer_consumers_v1_messages_proto_rawDescOnce sync.Once
	file_queuer_consumers_v1_messages_proto_rawDescData = file_queuer_consumers_v1_messages_proto_rawDesc
)

func file_queuer_consumers_v1_messages_proto_rawDescGZIP() []byte {
	file_queuer_consumers_v1_messages_proto_rawDescOnce.Do(func() {
		file_queuer_consumers_v1_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_queuer_consumers_v1_messages_proto_rawDescData)
	})
	return file_queuer_consumers_v1_messages_proto_rawDescData
}

var file_queuer_consumers_v1_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_queuer_consumers_v1_messages_proto_goTypes = []interface{}{
	(*RegisterConsumerRequest)(nil),      // 0: queuer.consumers.v1.RegisterConsumerRequest
	(*RegisterConsumerResponse)(nil),     // 1: queuer.consumers.v1.RegisterConsumerResponse
	(*GetConsumersRequest)(nil),          // 2: queuer.consumers.v1.GetConsumersRequest
	(*GetConsumersResponse)(nil),         // 3: queuer.consumers.v1.GetConsumersResponse
	(*PublishConsumerStatsRequest)(nil),  // 4: queuer.consumers.v1.PublishConsumerStatsRequest
	(*PublishConsumerStatsResponse)(nil), // 5: queuer.consumers.v1.PublishConsumerStatsResponse
	(*SubscribeToCommandsRequest)(nil),   // 6: queuer.consumers.v1.SubscribeToCommandsRequest
	(*SubscribeToCommandsResponse)(nil),  // 7: queuer.consumers.v1.SubscribeToCommandsResponse
	(*AckCommandRequest)(nil),            // 8: queuer.consumers.v1.AckCommandRequest
	(*AckCommandResponse)(nil),           // 9: queuer.consumers.v1.AckCommandResponse
	(*GetSubscribedStreamsRequest)(nil),  // 10: queuer.consumers.v1.GetSubscribedStreamsRequest
	(*GetSubscribedStreamsResponse)(nil), // 11: queuer.consumers.v1.GetSubscribedStreamsResponse
	(*v1.Consumer)(nil),                  // 12: queuer.entities.v1.Consumer
	(*v1.Stats)(nil),                     // 13: queuer.entities.v1.Stats
	(*v1.Command)(nil),                   // 14: queuer.entities.v1.Command
	(v1.CommandStatus)(0),                // 15: queuer.entities.v1.CommandStatus
	(*v1.Stream)(nil),                    // 16: queuer.entities.v1.Stream
}
var file_queuer_consumers_v1_messages_proto_depIdxs = []int32{
	12, // 0: queuer.consumers.v1.RegisterConsumerRequest.consumer:type_name -> queuer.entities.v1.Consumer
	12, // 1: queuer.consumers.v1.GetConsumersResponse.consumers:type_name -> queuer.entities.v1.Consumer
	13, // 2: queuer.consumers.v1.PublishConsumerStatsRequest.stats:type_name -> queuer.entities.v1.Stats
	14, // 3: queuer.consumers.v1.SubscribeToCommandsResponse.command:type_name -> queuer.entities.v1.Command
	15, // 4: queuer.consumers.v1.AckCommandRequest.status:type_name -> queuer.entities.v1.CommandStatus
	16, // 5: queuer.consumers.v1.GetSubscribedStreamsResponse.streams:type_name -> queuer.entities.v1.Stream
	6,  // [6:6] is the sub-list for method output_type
	6,  // [6:6] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_queuer_consumers_v1_messages_proto_init() }
func file_queuer_consumers_v1_messages_proto_init() {
	if File_queuer_consumers_v1_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_queuer_consumers_v1_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterConsumerRequest); i {
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
		file_queuer_consumers_v1_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterConsumerResponse); i {
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
		file_queuer_consumers_v1_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConsumersRequest); i {
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
		file_queuer_consumers_v1_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConsumersResponse); i {
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
		file_queuer_consumers_v1_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishConsumerStatsRequest); i {
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
		file_queuer_consumers_v1_messages_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishConsumerStatsResponse); i {
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
		file_queuer_consumers_v1_messages_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_queuer_consumers_v1_messages_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
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
		file_queuer_consumers_v1_messages_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
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
		file_queuer_consumers_v1_messages_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
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
		file_queuer_consumers_v1_messages_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSubscribedStreamsRequest); i {
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
		file_queuer_consumers_v1_messages_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSubscribedStreamsResponse); i {
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
			RawDescriptor: file_queuer_consumers_v1_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_queuer_consumers_v1_messages_proto_goTypes,
		DependencyIndexes: file_queuer_consumers_v1_messages_proto_depIdxs,
		MessageInfos:      file_queuer_consumers_v1_messages_proto_msgTypes,
	}.Build()
	File_queuer_consumers_v1_messages_proto = out.File
	file_queuer_consumers_v1_messages_proto_rawDesc = nil
	file_queuer_consumers_v1_messages_proto_goTypes = nil
	file_queuer_consumers_v1_messages_proto_depIdxs = nil
}
