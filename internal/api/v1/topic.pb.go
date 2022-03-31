// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: api/topic/topic.proto

package v1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TopicEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Specversion     string          `protobuf:"bytes,2,opt,name=specversion,proto3" json:"specversion,omitempty"`
	Type            string          `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Source          string          `protobuf:"bytes,4,opt,name=source,proto3" json:"source,omitempty"`
	Datacontenttype string          `protobuf:"bytes,5,opt,name=datacontenttype,proto3" json:"datacontenttype,omitempty"`
	Data            *structpb.Value `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
	DataBase64      string          `protobuf:"bytes,7,opt,name=data_base64,json=dataBase64,proto3" json:"data_base64,omitempty"`
	Subject         string          `protobuf:"bytes,8,opt,name=subject,proto3" json:"subject,omitempty"`
	Topic           string          `protobuf:"bytes,9,opt,name=topic,proto3" json:"topic,omitempty"`
	Pubsubname      string          `protobuf:"bytes,10,opt,name=pubsubname,proto3" json:"pubsubname,omitempty"`
}

func (x *TopicEventRequest) Reset() {
	*x = TopicEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_topic_topic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopicEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopicEventRequest) ProtoMessage() {}

func (x *TopicEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_topic_topic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopicEventRequest.ProtoReflect.Descriptor instead.
func (*TopicEventRequest) Descriptor() ([]byte, []int) {
	return file_api_topic_topic_proto_rawDescGZIP(), []int{0}
}

func (x *TopicEventRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TopicEventRequest) GetSpecversion() string {
	if x != nil {
		return x.Specversion
	}
	return ""
}

func (x *TopicEventRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *TopicEventRequest) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *TopicEventRequest) GetDatacontenttype() string {
	if x != nil {
		return x.Datacontenttype
	}
	return ""
}

func (x *TopicEventRequest) GetData() *structpb.Value {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *TopicEventRequest) GetDataBase64() string {
	if x != nil {
		return x.DataBase64
	}
	return ""
}

func (x *TopicEventRequest) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *TopicEventRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *TopicEventRequest) GetPubsubname() string {
	if x != nil {
		return x.Pubsubname
	}
	return ""
}

type TopicEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *TopicEventResponse) Reset() {
	*x = TopicEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_topic_topic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopicEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopicEventResponse) ProtoMessage() {}

func (x *TopicEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_topic_topic_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopicEventResponse.ProtoReflect.Descriptor instead.
func (*TopicEventResponse) Descriptor() ([]byte, []int) {
	return file_api_topic_topic_proto_rawDescGZIP(), []int{1}
}

func (x *TopicEventResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_api_topic_topic_proto protoreflect.FileDescriptor

var file_api_topic_topic_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x2f, 0x74, 0x6f, 0x70, 0x69,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xb8, 0x02, 0x0a, 0x11, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x70, 0x65, 0x63, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x70, 0x65,
	0x63, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x64,
	0x61, 0x74, 0x61, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2a,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x36, 0x34, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x64, 0x61, 0x74, 0x61, 0x42, 0x61, 0x73, 0x65, 0x36, 0x34, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x1e, 0x0a, 0x0a, 0x70,
	0x75, 0x62, 0x73, 0x75, 0x62, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2c, 0x0a, 0x12, 0x54,
	0x6f, 0x70, 0x69, 0x63, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xb4, 0x01, 0x0a, 0x05, 0x54, 0x6f,
	0x70, 0x69, 0x63, 0x12, 0xaa, 0x01, 0x0a, 0x11, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x54, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0b, 0x22, 0x06, 0x2f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x40,
	0x0a, 0x0a, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x54, 0x6f,
	0x70, 0x69, 0x63, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x20, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72,
	0x2a, 0x11, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x72, 0x4a, 0x0b, 0x0a, 0x03, 0x32, 0x30, 0x30, 0x12, 0x04, 0x0a, 0x02, 0x4f, 0x4b,
	0x42, 0x42, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x50,
	0x01, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x6b,
	0x65, 0x65, 0x6c, 0x2d, 0x69, 0x6f, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2d, 0x62, 0x72,
	0x6f, 0x6b, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x2f, 0x76,
	0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_topic_topic_proto_rawDescOnce sync.Once
	file_api_topic_topic_proto_rawDescData = file_api_topic_topic_proto_rawDesc
)

func file_api_topic_topic_proto_rawDescGZIP() []byte {
	file_api_topic_topic_proto_rawDescOnce.Do(func() {
		file_api_topic_topic_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_topic_topic_proto_rawDescData)
	})
	return file_api_topic_topic_proto_rawDescData
}

var file_api_topic_topic_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_topic_topic_proto_goTypes = []interface{}{
	(*TopicEventRequest)(nil),  // 0: api.core.v1.TopicEventRequest
	(*TopicEventResponse)(nil), // 1: api.core.v1.TopicEventResponse
	(*structpb.Value)(nil),     // 2: google.protobuf.Value
}
var file_api_topic_topic_proto_depIdxs = []int32{
	2, // 0: api.core.v1.TopicEventRequest.data:type_name -> google.protobuf.Value
	0, // 1: api.core.v1.Topic.TopicEventHandler:input_type -> api.core.v1.TopicEventRequest
	1, // 2: api.core.v1.Topic.TopicEventHandler:output_type -> api.core.v1.TopicEventResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_topic_topic_proto_init() }
func file_api_topic_topic_proto_init() {
	if File_api_topic_topic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_topic_topic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopicEventRequest); i {
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
		file_api_topic_topic_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopicEventResponse); i {
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
			RawDescriptor: file_api_topic_topic_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_topic_topic_proto_goTypes,
		DependencyIndexes: file_api_topic_topic_proto_depIdxs,
		MessageInfos:      file_api_topic_topic_proto_msgTypes,
	}.Build()
	File_api_topic_topic_proto = out.File
	file_api_topic_topic_proto_rawDesc = nil
	file_api_topic_topic_proto_goTypes = nil
	file_api_topic_topic_proto_depIdxs = nil
}
