// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: worker.proto

package worker

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

// Registers the executor with the service.
// Once an executor registers, it can run a workflow.
type ExecutorMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExecutorId string `protobuf:"bytes,1,opt,name=executor_id,json=executorId,proto3" json:"executor_id,omitempty"`
}

func (x *ExecutorMessage) Reset() {
	*x = ExecutorMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecutorMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutorMessage) ProtoMessage() {}

func (x *ExecutorMessage) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutorMessage.ProtoReflect.Descriptor instead.
func (*ExecutorMessage) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{0}
}

func (x *ExecutorMessage) GetExecutorId() string {
	if x != nil {
		return x.ExecutorId
	}
	return ""
}

// Respond to the request with the desired workflow xml.
type WorkerMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Command       string `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	WorkflowXml   []byte `protobuf:"bytes,2,opt,name=workflow_xml,json=workflowXml,proto3" json:"workflow_xml,omitempty"`
	ParametersXml []byte `protobuf:"bytes,3,opt,name=parameters_xml,json=parametersXml,proto3" json:"parameters_xml,omitempty"`
}

func (x *WorkerMessage) Reset() {
	*x = WorkerMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkerMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerMessage) ProtoMessage() {}

func (x *WorkerMessage) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerMessage.ProtoReflect.Descriptor instead.
func (*WorkerMessage) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{1}
}

func (x *WorkerMessage) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *WorkerMessage) GetWorkflowXml() []byte {
	if x != nil {
		return x.WorkflowXml
	}
	return nil
}

func (x *WorkerMessage) GetParametersXml() []byte {
	if x != nil {
		return x.ParametersXml
	}
	return nil
}

type ResultMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Done string `protobuf:"bytes,1,opt,name=done,proto3" json:"done,omitempty"`
}

func (x *ResultMessageRequest) Reset() {
	*x = ResultMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultMessageRequest) ProtoMessage() {}

func (x *ResultMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultMessageRequest.ProtoReflect.Descriptor instead.
func (*ResultMessageRequest) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{2}
}

func (x *ResultMessageRequest) GetDone() string {
	if x != nil {
		return x.Done
	}
	return ""
}

type ResultMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ack string `protobuf:"bytes,1,opt,name=ack,proto3" json:"ack,omitempty"`
}

func (x *ResultMessageResponse) Reset() {
	*x = ResultMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultMessageResponse) ProtoMessage() {}

func (x *ResultMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultMessageResponse.ProtoReflect.Descriptor instead.
func (*ResultMessageResponse) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{3}
}

func (x *ResultMessageResponse) GetAck() string {
	if x != nil {
		return x.Ack
	}
	return ""
}

var File_worker_proto protoreflect.FileDescriptor

var file_worker_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x22, 0x32, 0x0a, 0x0f, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x22, 0x73, 0x0a, 0x0d, 0x57, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x5f, 0x78, 0x6d, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x77, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x58, 0x6d, 0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x5f, 0x78, 0x6d, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0d, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x58, 0x6d, 0x6c, 0x22,
	0x2a, 0x0a, 0x14, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x6f, 0x6e, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x6f, 0x6e, 0x65, 0x22, 0x29, 0x0a, 0x15, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x61, 0x63, 0x6b, 0x32, 0x9a, 0x01, 0x0a, 0x06, 0x57, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x12, 0x49, 0x0a, 0x13, 0x43, 0x6f, 0x6d, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x17, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x1a, 0x15, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x28, 0x01, 0x30, 0x01, 0x12, 0x45, 0x0a, 0x06,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1c, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x42, 0x72, 0x6f, 0x77, 0x64, 0x75, 0x65, 0x73, 0x4d, 0x61, 0x6e, 0x38, 0x35, 0x2f,
	0x6f, 0x75, 0x74, 0x2d, 0x6f, 0x66, 0x2d, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x74,
	0x61, 0x6b, 0x65, 0x2d, 0x74, 0x77, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x77, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_worker_proto_rawDescOnce sync.Once
	file_worker_proto_rawDescData = file_worker_proto_rawDesc
)

func file_worker_proto_rawDescGZIP() []byte {
	file_worker_proto_rawDescOnce.Do(func() {
		file_worker_proto_rawDescData = protoimpl.X.CompressGZIP(file_worker_proto_rawDescData)
	})
	return file_worker_proto_rawDescData
}

var file_worker_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_worker_proto_goTypes = []interface{}{
	(*ExecutorMessage)(nil),       // 0: worker.ExecutorMessage
	(*WorkerMessage)(nil),         // 1: worker.WorkerMessage
	(*ResultMessageRequest)(nil),  // 2: worker.ResultMessageRequest
	(*ResultMessageResponse)(nil), // 3: worker.ResultMessageResponse
}
var file_worker_proto_depIdxs = []int32{
	0, // 0: worker.Worker.CommmandMessageLink:input_type -> worker.ExecutorMessage
	2, // 1: worker.Worker.Result:input_type -> worker.ResultMessageRequest
	1, // 2: worker.Worker.CommmandMessageLink:output_type -> worker.WorkerMessage
	3, // 3: worker.Worker.Result:output_type -> worker.ResultMessageResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_worker_proto_init() }
func file_worker_proto_init() {
	if File_worker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_worker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecutorMessage); i {
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
		file_worker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkerMessage); i {
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
		file_worker_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultMessageRequest); i {
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
		file_worker_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultMessageResponse); i {
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
			RawDescriptor: file_worker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_worker_proto_goTypes,
		DependencyIndexes: file_worker_proto_depIdxs,
		MessageInfos:      file_worker_proto_msgTypes,
	}.Build()
	File_worker_proto = out.File
	file_worker_proto_rawDesc = nil
	file_worker_proto_goTypes = nil
	file_worker_proto_depIdxs = nil
}
