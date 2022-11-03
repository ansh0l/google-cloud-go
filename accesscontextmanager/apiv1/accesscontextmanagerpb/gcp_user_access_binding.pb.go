// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.5
// source: google/identity/accesscontextmanager/v1/gcp_user_access_binding.proto

package accesscontextmanagerpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Restricts access to Cloud Console and Google Cloud APIs for a set of users
// using Context-Aware Access.
type GcpUserAccessBinding struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. Assigned by the server during creation. The last segment has an arbitrary
	// length and has only URI unreserved characters (as defined by
	// [RFC 3986 Section 2.3](https://tools.ietf.org/html/rfc3986#section-2.3)).
	// Should not be specified by the client during creation.
	// Example: "organizations/256/gcpUserAccessBindings/b3-BhcX_Ud5N"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. Immutable. Google Group id whose members are subject to this binding's restrictions.
	// See "id" in the [G Suite Directory API's Groups resource]
	// (https://developers.google.com/admin-sdk/directory/v1/reference/groups#resource).
	// If a group's email address/alias is changed, this resource will continue
	// to point at the changed group. This field does not accept group email
	// addresses or aliases.
	// Example: "01d520gv4vjcrht"
	GroupKey string `protobuf:"bytes,2,opt,name=group_key,json=groupKey,proto3" json:"group_key,omitempty"`
	// Required. Access level that a user must have to be granted access. Only one access
	// level is supported, not multiple. This repeated field must have exactly
	// one element.
	// Example: "accessPolicies/9522/accessLevels/device_trusted"
	AccessLevels []string `protobuf:"bytes,3,rep,name=access_levels,json=accessLevels,proto3" json:"access_levels,omitempty"`
}

func (x *GcpUserAccessBinding) Reset() {
	*x = GcpUserAccessBinding{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_identity_accesscontextmanager_v1_gcp_user_access_binding_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GcpUserAccessBinding) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcpUserAccessBinding) ProtoMessage() {}

func (x *GcpUserAccessBinding) ProtoReflect() protoreflect.Message {
	mi := &file_google_identity_accesscontextmanager_v1_gcp_user_access_binding_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcpUserAccessBinding.ProtoReflect.Descriptor instead.
func (*GcpUserAccessBinding) Descriptor() ([]byte, []int) {
	return file_google_identity_accesscontextmanager_v1_gcp_user_access_binding_proto_rawDescGZIP(), []int{0}
}

func (x *GcpUserAccessBinding) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GcpUserAccessBinding) GetGroupKey() string {
	if x != nil {
		return x.GroupKey
	}
	return ""
}

func (x *GcpUserAccessBinding) GetAccessLevels() []string {
	if x != nil {
		return x.AccessLevels
	}
	return nil
}

var File_google_identity_accesscontextmanager_v1_gcp_user_access_binding_proto protoreflect.FileDescriptor

var file_google_identity_accesscontextmanager_v1_gcp_user_access_binding_proto_rawDesc = []byte{
	0x0a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x63, 0x70, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x27, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc1, 0x02, 0x0a,
	0x14, 0x47, 0x63, 0x70, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x69,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23,
	0x0a, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x06, 0xe0, 0x41, 0x02, 0xe0, 0x41, 0x05, 0x52, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x4b, 0x65, 0x79, 0x12, 0x5c, 0x0a, 0x0d, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x42, 0x37, 0xe0, 0x41, 0x02, 0xfa,
	0x41, 0x31, 0x0a, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x52, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x73, 0x3a, 0x8c, 0x01, 0xea, 0x41, 0x88, 0x01, 0x0a, 0x38, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x63,
	0x70, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x69, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x12, 0x4c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x7b, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d,
	0x2f, 0x67, 0x63, 0x70, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x69,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x7b, 0x67, 0x63, 0x70, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x7d,
	0x42, 0xaf, 0x02, 0x0a, 0x2b, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x42, 0x19, 0x47, 0x63, 0x70, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42,
	0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0xa2, 0x02, 0x04, 0x47, 0x41, 0x43,
	0x4d, 0xaa, 0x02, 0x27, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x27, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5c, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x2a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x3a, 0x3a, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_identity_accesscontextmanager_v1_gcp_user_access_binding_proto_rawDescOnce sync.Once
	file_google_identity_accesscontextmanager_v1_gcp_user_access_binding_proto_rawDescData = file_google_identity_accesscontextmanager_v1_gcp_user_access_binding_proto_rawDesc
)

func file_google_identity_accesscontextmanager_v1_gcp_user_access_binding_proto_rawDescGZIP() []byte {
	file_google_identity_accesscontextmanager_v1_gcp_user_access_binding_proto_rawDescOnce.Do(func() {
		file_google_identity_accesscontextmanager_v1_gcp_user_access_binding_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_identity_accesscontextmanager_v1_gcp_user_access_binding_proto_rawDescData)
	})
	return file_google_identity_accesscontextmanager_v1_gcp_user_access_binding_proto_rawDescData
}

var file_google_identity_accesscontextmanager_v1_gcp_user_access_binding_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_identity_accesscontextmanager_v1_gcp_user_access_binding_proto_goTypes = []interface{}{
	(*GcpUserAccessBinding)(nil), // 0: google.identity.accesscontextmanager.v1.GcpUserAccessBinding
}
var file_google_identity_accesscontextmanager_v1_gcp_user_access_binding_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_identity_accesscontextmanager_v1_gcp_user_access_binding_proto_init() }
func file_google_identity_accesscontextmanager_v1_gcp_user_access_binding_proto_init() {
	if File_google_identity_accesscontextmanager_v1_gcp_user_access_binding_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_identity_accesscontextmanager_v1_gcp_user_access_binding_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GcpUserAccessBinding); i {
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
			RawDescriptor: file_google_identity_accesscontextmanager_v1_gcp_user_access_binding_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_identity_accesscontextmanager_v1_gcp_user_access_binding_proto_goTypes,
		DependencyIndexes: file_google_identity_accesscontextmanager_v1_gcp_user_access_binding_proto_depIdxs,
		MessageInfos:      file_google_identity_accesscontextmanager_v1_gcp_user_access_binding_proto_msgTypes,
	}.Build()
	File_google_identity_accesscontextmanager_v1_gcp_user_access_binding_proto = out.File
	file_google_identity_accesscontextmanager_v1_gcp_user_access_binding_proto_rawDesc = nil
	file_google_identity_accesscontextmanager_v1_gcp_user_access_binding_proto_goTypes = nil
	file_google_identity_accesscontextmanager_v1_gcp_user_access_binding_proto_depIdxs = nil
}
