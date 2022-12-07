// Copyright 2016-2018, Pulumi Corporation.
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
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: pulumi/plugin.proto

package pulumirpc

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

// PluginInfo is meta-information about a plugin that is used by the system.
type PluginInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"` // the semver for this plugin.
}

func (x *PluginInfo) Reset() {
	*x = PluginInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pulumi_plugin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginInfo) ProtoMessage() {}

func (x *PluginInfo) ProtoReflect() protoreflect.Message {
	mi := &file_pulumi_plugin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginInfo.ProtoReflect.Descriptor instead.
func (*PluginInfo) Descriptor() ([]byte, []int) {
	return file_pulumi_plugin_proto_rawDescGZIP(), []int{0}
}

func (x *PluginInfo) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

// PluginDependency is information about a plugin that a program may depend upon.
type PluginDependency struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`       // the name of the plugin.
	Kind    string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`       // the kind of plugin (e.g., language, etc).
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"` // the semver for this plugin.
	Server  string `protobuf:"bytes,4,opt,name=server,proto3" json:"server,omitempty"`   // the URL of a server that can be used to download this plugin, if needed.
}

func (x *PluginDependency) Reset() {
	*x = PluginDependency{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pulumi_plugin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginDependency) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginDependency) ProtoMessage() {}

func (x *PluginDependency) ProtoReflect() protoreflect.Message {
	mi := &file_pulumi_plugin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginDependency.ProtoReflect.Descriptor instead.
func (*PluginDependency) Descriptor() ([]byte, []int) {
	return file_pulumi_plugin_proto_rawDescGZIP(), []int{1}
}

func (x *PluginDependency) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PluginDependency) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *PluginDependency) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *PluginDependency) GetServer() string {
	if x != nil {
		return x.Server
	}
	return ""
}

// PluginAttach is used to attach an already running plugin to the engine.
//
// Normally the engine starts the plugin process itself and passes the engine address as the first argumnent.
// But when debugging it can be useful to have an already running provider that the engine instead attaches
// to, this message is used so the provider can still be passed the engine address to communicate with.
type PluginAttach struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"` // the grpc address for the engine
}

func (x *PluginAttach) Reset() {
	*x = PluginAttach{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pulumi_plugin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginAttach) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginAttach) ProtoMessage() {}

func (x *PluginAttach) ProtoReflect() protoreflect.Message {
	mi := &file_pulumi_plugin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginAttach.ProtoReflect.Descriptor instead.
func (*PluginAttach) Descriptor() ([]byte, []int) {
	return file_pulumi_plugin_proto_rawDescGZIP(), []int{2}
}

func (x *PluginAttach) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

var File_pulumi_plugin_proto protoreflect.FileDescriptor

var file_pulumi_plugin_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x72, 0x70, 0x63,
	0x22, 0x26, 0x0a, 0x0a, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18,
	0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x6c, 0x0a, 0x10, 0x50, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6b, 0x69, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x28, 0x0a, 0x0c, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70,
	0x75, 0x6c, 0x75, 0x6d, 0x69, 0x2f, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x2f, 0x73, 0x64, 0x6b,
	0x2f, 0x76, 0x33, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x3b, 0x70, 0x75, 0x6c,
	0x75, 0x6d, 0x69, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pulumi_plugin_proto_rawDescOnce sync.Once
	file_pulumi_plugin_proto_rawDescData = file_pulumi_plugin_proto_rawDesc
)

func file_pulumi_plugin_proto_rawDescGZIP() []byte {
	file_pulumi_plugin_proto_rawDescOnce.Do(func() {
		file_pulumi_plugin_proto_rawDescData = protoimpl.X.CompressGZIP(file_pulumi_plugin_proto_rawDescData)
	})
	return file_pulumi_plugin_proto_rawDescData
}

var file_pulumi_plugin_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pulumi_plugin_proto_goTypes = []interface{}{
	(*PluginInfo)(nil),       // 0: pulumirpc.PluginInfo
	(*PluginDependency)(nil), // 1: pulumirpc.PluginDependency
	(*PluginAttach)(nil),     // 2: pulumirpc.PluginAttach
}
var file_pulumi_plugin_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pulumi_plugin_proto_init() }
func file_pulumi_plugin_proto_init() {
	if File_pulumi_plugin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pulumi_plugin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginInfo); i {
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
		file_pulumi_plugin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginDependency); i {
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
		file_pulumi_plugin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginAttach); i {
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
			RawDescriptor: file_pulumi_plugin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pulumi_plugin_proto_goTypes,
		DependencyIndexes: file_pulumi_plugin_proto_depIdxs,
		MessageInfos:      file_pulumi_plugin_proto_msgTypes,
	}.Build()
	File_pulumi_plugin_proto = out.File
	file_pulumi_plugin_proto_rawDesc = nil
	file_pulumi_plugin_proto_goTypes = nil
	file_pulumi_plugin_proto_depIdxs = nil
}
