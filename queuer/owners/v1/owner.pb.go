// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0-devel
// 	protoc        (unknown)
// source: queuer/owners/v1/owner.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gsols/goproto/queuer/entities/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_queuer_owners_v1_owner_proto protoreflect.FileDescriptor

var file_queuer_owners_v1_owner_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x71, 0x75, 0x65, 0x75, 0x65, 0x72, 0x2f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x71, 0x75, 0x65, 0x75, 0x65, 0x72, 0x2e, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x72, 0x2f, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x72, 0x2f, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x75, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x72, 0x2f, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xda, 0x02, 0x0a, 0x0c, 0x4f, 0x77, 0x6e, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x71, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x24, 0x2e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x72,
	0x2e, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e,
	0x71, 0x75, 0x65, 0x75, 0x65, 0x72, 0x2e, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x3a, 0x01, 0x2a, 0x22,
	0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x6a, 0x0a, 0x08, 0x47,
	0x65, 0x74, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x21, 0x2e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x72,
	0x2e, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x77,
	0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x71, 0x75, 0x65,
	0x75, 0x65, 0x72, 0x2e, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x6b, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x51,
	0x75, 0x65, 0x75, 0x65, 0x73, 0x12, 0x23, 0x2e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x72, 0x2e, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x51, 0x75, 0x65,
	0x75, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x71, 0x75, 0x65,
	0x75, 0x65, 0x72, 0x2e, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x51, 0x75, 0x65, 0x75, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75,
	0x65, 0x75, 0x65, 0x73, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x67, 0x73, 0x6f, 0x6c, 0x73, 0x2f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x72, 0x2f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x2f, 0x76,
	0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_queuer_owners_v1_owner_proto_goTypes = []interface{}{
	(*CreateOwnerRequest)(nil),  // 0: queuer.owners.v1.CreateOwnerRequest
	(*GetOwnerRequest)(nil),     // 1: queuer.owners.v1.GetOwnerRequest
	(*ListQueuesRequest)(nil),   // 2: queuer.owners.v1.ListQueuesRequest
	(*CreateOwnerResponse)(nil), // 3: queuer.owners.v1.CreateOwnerResponse
	(*GetOwnerResponse)(nil),    // 4: queuer.owners.v1.GetOwnerResponse
	(*ListQueuesResponse)(nil),  // 5: queuer.owners.v1.ListQueuesResponse
}
var file_queuer_owners_v1_owner_proto_depIdxs = []int32{
	0, // 0: queuer.owners.v1.OwnerService.CreateOwner:input_type -> queuer.owners.v1.CreateOwnerRequest
	1, // 1: queuer.owners.v1.OwnerService.GetOwner:input_type -> queuer.owners.v1.GetOwnerRequest
	2, // 2: queuer.owners.v1.OwnerService.ListQueues:input_type -> queuer.owners.v1.ListQueuesRequest
	3, // 3: queuer.owners.v1.OwnerService.CreateOwner:output_type -> queuer.owners.v1.CreateOwnerResponse
	4, // 4: queuer.owners.v1.OwnerService.GetOwner:output_type -> queuer.owners.v1.GetOwnerResponse
	5, // 5: queuer.owners.v1.OwnerService.ListQueues:output_type -> queuer.owners.v1.ListQueuesResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_queuer_owners_v1_owner_proto_init() }
func file_queuer_owners_v1_owner_proto_init() {
	if File_queuer_owners_v1_owner_proto != nil {
		return
	}
	file_queuer_owners_v1_messages_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_queuer_owners_v1_owner_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_queuer_owners_v1_owner_proto_goTypes,
		DependencyIndexes: file_queuer_owners_v1_owner_proto_depIdxs,
	}.Build()
	File_queuer_owners_v1_owner_proto = out.File
	file_queuer_owners_v1_owner_proto_rawDesc = nil
	file_queuer_owners_v1_owner_proto_goTypes = nil
	file_queuer_owners_v1_owner_proto_depIdxs = nil
}
