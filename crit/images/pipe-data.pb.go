// SPDX-License-Identifier: MIT

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: pipe-data.proto

package images

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

type PipeDataEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PipeId *uint32 `protobuf:"varint,1,req,name=pipe_id,json=pipeId" json:"pipe_id,omitempty"`
	Bytes  *uint32 `protobuf:"varint,2,req,name=bytes" json:"bytes,omitempty"`
	Size   *uint32 `protobuf:"varint,3,opt,name=size" json:"size,omitempty"`
}

func (x *PipeDataEntry) Reset() {
	*x = PipeDataEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pipe_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PipeDataEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PipeDataEntry) ProtoMessage() {}

func (x *PipeDataEntry) ProtoReflect() protoreflect.Message {
	mi := &file_pipe_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PipeDataEntry.ProtoReflect.Descriptor instead.
func (*PipeDataEntry) Descriptor() ([]byte, []int) {
	return file_pipe_data_proto_rawDescGZIP(), []int{0}
}

func (x *PipeDataEntry) GetPipeId() uint32 {
	if x != nil && x.PipeId != nil {
		return *x.PipeId
	}
	return 0
}

func (x *PipeDataEntry) GetBytes() uint32 {
	if x != nil && x.Bytes != nil {
		return *x.Bytes
	}
	return 0
}

func (x *PipeDataEntry) GetSize() uint32 {
	if x != nil && x.Size != nil {
		return *x.Size
	}
	return 0
}

var File_pipe_data_proto protoreflect.FileDescriptor

var file_pipe_data_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x69, 0x70, 0x65, 0x2d, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x54, 0x0a, 0x0f, 0x70, 0x69, 0x70, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x65,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x69, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x06, 0x70, 0x69, 0x70, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x05, 0x62, 0x79,
	0x74, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65,
}

var (
	file_pipe_data_proto_rawDescOnce sync.Once
	file_pipe_data_proto_rawDescData = file_pipe_data_proto_rawDesc
)

func file_pipe_data_proto_rawDescGZIP() []byte {
	file_pipe_data_proto_rawDescOnce.Do(func() {
		file_pipe_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_pipe_data_proto_rawDescData)
	})
	return file_pipe_data_proto_rawDescData
}

var file_pipe_data_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pipe_data_proto_goTypes = []interface{}{
	(*PipeDataEntry)(nil), // 0: pipe_data_entry
}
var file_pipe_data_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pipe_data_proto_init() }
func file_pipe_data_proto_init() {
	if File_pipe_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pipe_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PipeDataEntry); i {
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
			RawDescriptor: file_pipe_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pipe_data_proto_goTypes,
		DependencyIndexes: file_pipe_data_proto_depIdxs,
		MessageInfos:      file_pipe_data_proto_msgTypes,
	}.Build()
	File_pipe_data_proto = out.File
	file_pipe_data_proto_rawDesc = nil
	file_pipe_data_proto_goTypes = nil
	file_pipe_data_proto_depIdxs = nil
}
