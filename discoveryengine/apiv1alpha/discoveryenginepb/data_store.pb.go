// Copyright 2022 Google LLC
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
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.2
// source: google/cloud/discoveryengine/v1alpha/data_store.proto

package discoveryenginepb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Content config of the data store.
type DataStore_ContentConfig int32

const (
	// Default value.
	DataStore_CONTENT_CONFIG_UNSPECIFIED DataStore_ContentConfig = 0
	// Only contains documents without any
	// [Document.content][google.cloud.discoveryengine.v1alpha.Document.content].
	DataStore_NO_CONTENT DataStore_ContentConfig = 1
	// Only contains documents with
	// [Document.content][google.cloud.discoveryengine.v1alpha.Document.content].
	DataStore_CONTENT_REQUIRED DataStore_ContentConfig = 2
	// The data store is used for public website search.
	DataStore_PUBLIC_WEBSITE DataStore_ContentConfig = 3
)

// Enum value maps for DataStore_ContentConfig.
var (
	DataStore_ContentConfig_name = map[int32]string{
		0: "CONTENT_CONFIG_UNSPECIFIED",
		1: "NO_CONTENT",
		2: "CONTENT_REQUIRED",
		3: "PUBLIC_WEBSITE",
	}
	DataStore_ContentConfig_value = map[string]int32{
		"CONTENT_CONFIG_UNSPECIFIED": 0,
		"NO_CONTENT":                 1,
		"CONTENT_REQUIRED":           2,
		"PUBLIC_WEBSITE":             3,
	}
)

func (x DataStore_ContentConfig) Enum() *DataStore_ContentConfig {
	p := new(DataStore_ContentConfig)
	*p = x
	return p
}

func (x DataStore_ContentConfig) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DataStore_ContentConfig) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_discoveryengine_v1alpha_data_store_proto_enumTypes[0].Descriptor()
}

func (DataStore_ContentConfig) Type() protoreflect.EnumType {
	return &file_google_cloud_discoveryengine_v1alpha_data_store_proto_enumTypes[0]
}

func (x DataStore_ContentConfig) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DataStore_ContentConfig.Descriptor instead.
func (DataStore_ContentConfig) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_discoveryengine_v1alpha_data_store_proto_rawDescGZIP(), []int{0, 0}
}

// DataStore captures global settings and configs at the DataStore level.
type DataStore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. The full resource name of the data store.
	// Format:
	// `projects/{project}/locations/{location}/collections/{collection_id}/dataStores/{data_store_id}`.
	//
	// This field must be a UTF-8 encoded string with a length limit of 1024
	// characters.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The data store display name.
	//
	// This field must be a UTF-8 encoded string with a length limit of 128
	// characters. Otherwise, an INVALID_ARGUMENT error is returned.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Immutable. The industry vertical that the data store registers.
	IndustryVertical IndustryVertical `protobuf:"varint,3,opt,name=industry_vertical,json=industryVertical,proto3,enum=google.cloud.discoveryengine.v1alpha.IndustryVertical" json:"industry_vertical,omitempty"`
	// The solutions that the data store enrolls. Available solutions for each
	// [industry_vertical][google.cloud.discoveryengine.v1alpha.DataStore.industry_vertical]:
	//
	// * `MEDIA`: `SOLUTION_TYPE_RECOMMENDATION` and `SOLUTION_TYPE_SEARCH`.
	// * `SITE_SEARCH`: `SOLUTION_TYPE_SEARCH` is automatically enrolled. Other
	//   solutions cannot be enrolled.
	SolutionTypes []SolutionType `protobuf:"varint,5,rep,packed,name=solution_types,json=solutionTypes,proto3,enum=google.cloud.discoveryengine.v1alpha.SolutionType" json:"solution_types,omitempty"`
	// Output only. The id of the default
	// [Schema][google.cloud.discoveryengine.v1alpha.Schema] asscociated to this
	// data store.
	DefaultSchemaId string `protobuf:"bytes,7,opt,name=default_schema_id,json=defaultSchemaId,proto3" json:"default_schema_id,omitempty"`
	// Immutable. The content config of the data store. If this field is unset,
	// the server behavior defaults to
	// [ContentConfig.NO_CONTENT][google.cloud.discoveryengine.v1alpha.DataStore.ContentConfig.NO_CONTENT].
	ContentConfig DataStore_ContentConfig `protobuf:"varint,6,opt,name=content_config,json=contentConfig,proto3,enum=google.cloud.discoveryengine.v1alpha.DataStore_ContentConfig" json:"content_config,omitempty"`
	// Output only. Timestamp the
	// [DataStore][google.cloud.discoveryengine.v1alpha.DataStore] was created at.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
}

func (x *DataStore) Reset() {
	*x = DataStore{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_discoveryengine_v1alpha_data_store_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataStore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataStore) ProtoMessage() {}

func (x *DataStore) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_discoveryengine_v1alpha_data_store_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataStore.ProtoReflect.Descriptor instead.
func (*DataStore) Descriptor() ([]byte, []int) {
	return file_google_cloud_discoveryengine_v1alpha_data_store_proto_rawDescGZIP(), []int{0}
}

func (x *DataStore) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DataStore) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *DataStore) GetIndustryVertical() IndustryVertical {
	if x != nil {
		return x.IndustryVertical
	}
	return IndustryVertical_INDUSTRY_VERTICAL_UNSPECIFIED
}

func (x *DataStore) GetSolutionTypes() []SolutionType {
	if x != nil {
		return x.SolutionTypes
	}
	return nil
}

func (x *DataStore) GetDefaultSchemaId() string {
	if x != nil {
		return x.DefaultSchemaId
	}
	return ""
}

func (x *DataStore) GetContentConfig() DataStore_ContentConfig {
	if x != nil {
		return x.ContentConfig
	}
	return DataStore_CONTENT_CONFIG_UNSPECIFIED
}

func (x *DataStore) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

var File_google_cloud_discoveryengine_v1alpha_data_store_proto protoreflect.FileDescriptor

var file_google_cloud_discoveryengine_v1alpha_data_store_proto_rawDesc = []byte{
	0x0a, 0x35, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x06,
	0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x17, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52,
	0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x68, 0x0a, 0x11,
	0x69, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x5f, 0x76, 0x65, 0x72, 0x74, 0x69, 0x63, 0x61,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x36, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x49,
	0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x56, 0x65, 0x72, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x42,
	0x03, 0xe0, 0x41, 0x05, 0x52, 0x10, 0x69, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x56, 0x65,
	0x72, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x12, 0x59, 0x0a, 0x0e, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x32,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x0d, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x73, 0x12, 0x2f, 0x0a, 0x11, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x52, 0x0f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x49, 0x64, 0x12, 0x69, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x0d,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x40, 0x0a,
	0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03,
	0xe0, 0x41, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22,
	0x69, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x1e, 0x0a, 0x1a, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x46,
	0x49, 0x47, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x0e, 0x0a, 0x0a, 0x4e, 0x4f, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x10, 0x01,
	0x12, 0x14, 0x0a, 0x10, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x5f, 0x52, 0x45, 0x51, 0x55,
	0x49, 0x52, 0x45, 0x44, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x43,
	0x5f, 0x57, 0x45, 0x42, 0x53, 0x49, 0x54, 0x45, 0x10, 0x03, 0x3a, 0xc9, 0x01, 0xea, 0x41, 0xc5,
	0x01, 0x0a, 0x28, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x3f, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x2f,
	0x7b, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x7d, 0x12, 0x58, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d,
	0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x2f, 0x7b, 0x64, 0x61, 0x74, 0x61, 0x5f,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x7d, 0x42, 0x9a, 0x02, 0x0a, 0x28, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x42, 0x0e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x52, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x70, 0x62, 0x3b, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x70, 0x62, 0xa2, 0x02, 0x0f, 0x44, 0x49, 0x53, 0x43,
	0x4f, 0x56, 0x45, 0x52, 0x59, 0x45, 0x4e, 0x47, 0x49, 0x4e, 0x45, 0xaa, 0x02, 0x24, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x56, 0x31, 0x41, 0x6c, 0x70,
	0x68, 0x61, 0xca, 0x02, 0x24, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x5c, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x45, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xea, 0x02, 0x27, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x44, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_discoveryengine_v1alpha_data_store_proto_rawDescOnce sync.Once
	file_google_cloud_discoveryengine_v1alpha_data_store_proto_rawDescData = file_google_cloud_discoveryengine_v1alpha_data_store_proto_rawDesc
)

func file_google_cloud_discoveryengine_v1alpha_data_store_proto_rawDescGZIP() []byte {
	file_google_cloud_discoveryengine_v1alpha_data_store_proto_rawDescOnce.Do(func() {
		file_google_cloud_discoveryengine_v1alpha_data_store_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_discoveryengine_v1alpha_data_store_proto_rawDescData)
	})
	return file_google_cloud_discoveryengine_v1alpha_data_store_proto_rawDescData
}

var file_google_cloud_discoveryengine_v1alpha_data_store_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_discoveryengine_v1alpha_data_store_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_discoveryengine_v1alpha_data_store_proto_goTypes = []interface{}{
	(DataStore_ContentConfig)(0),  // 0: google.cloud.discoveryengine.v1alpha.DataStore.ContentConfig
	(*DataStore)(nil),             // 1: google.cloud.discoveryengine.v1alpha.DataStore
	(IndustryVertical)(0),         // 2: google.cloud.discoveryengine.v1alpha.IndustryVertical
	(SolutionType)(0),             // 3: google.cloud.discoveryengine.v1alpha.SolutionType
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_google_cloud_discoveryengine_v1alpha_data_store_proto_depIdxs = []int32{
	2, // 0: google.cloud.discoveryengine.v1alpha.DataStore.industry_vertical:type_name -> google.cloud.discoveryengine.v1alpha.IndustryVertical
	3, // 1: google.cloud.discoveryengine.v1alpha.DataStore.solution_types:type_name -> google.cloud.discoveryengine.v1alpha.SolutionType
	0, // 2: google.cloud.discoveryengine.v1alpha.DataStore.content_config:type_name -> google.cloud.discoveryengine.v1alpha.DataStore.ContentConfig
	4, // 3: google.cloud.discoveryengine.v1alpha.DataStore.create_time:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_cloud_discoveryengine_v1alpha_data_store_proto_init() }
func file_google_cloud_discoveryengine_v1alpha_data_store_proto_init() {
	if File_google_cloud_discoveryengine_v1alpha_data_store_proto != nil {
		return
	}
	file_google_cloud_discoveryengine_v1alpha_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_discoveryengine_v1alpha_data_store_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataStore); i {
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
			RawDescriptor: file_google_cloud_discoveryengine_v1alpha_data_store_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_discoveryengine_v1alpha_data_store_proto_goTypes,
		DependencyIndexes: file_google_cloud_discoveryengine_v1alpha_data_store_proto_depIdxs,
		EnumInfos:         file_google_cloud_discoveryengine_v1alpha_data_store_proto_enumTypes,
		MessageInfos:      file_google_cloud_discoveryengine_v1alpha_data_store_proto_msgTypes,
	}.Build()
	File_google_cloud_discoveryengine_v1alpha_data_store_proto = out.File
	file_google_cloud_discoveryengine_v1alpha_data_store_proto_rawDesc = nil
	file_google_cloud_discoveryengine_v1alpha_data_store_proto_goTypes = nil
	file_google_cloud_discoveryengine_v1alpha_data_store_proto_depIdxs = nil
}
