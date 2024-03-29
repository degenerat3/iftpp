// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: iftpp.proto

package pbuf

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type IFTPP_Flag int32

const (
	IFTPP_SESSION_INIT IFTPP_Flag = 0 // used by client to propose session ID
	IFTPP_ACK          IFTPP_Flag = 1 // generic acknowledge, multiple uses
	IFTPP_CLIENT_KEY   IFTPP_Flag = 2 // client proposed key
	IFTPP_SERVER_KEY   IFTPP_Flag = 3 // server proposed key
	IFTPP_FILE_REQ     IFTPP_Flag = 4 // client requesting file
	IFTPP_FILE_DATA    IFTPP_Flag = 5 // requested file data
	IFTPP_FIN          IFTPP_Flag = 6 // transfer is complete
	IFTPP_RETRANS      IFTPP_Flag = 7 // request retransmission of prev packet
)

// Enum value maps for IFTPP_Flag.
var (
	IFTPP_Flag_name = map[int32]string{
		0: "SESSION_INIT",
		1: "ACK",
		2: "CLIENT_KEY",
		3: "SERVER_KEY",
		4: "FILE_REQ",
		5: "FILE_DATA",
		6: "FIN",
		7: "RETRANS",
	}
	IFTPP_Flag_value = map[string]int32{
		"SESSION_INIT": 0,
		"ACK":          1,
		"CLIENT_KEY":   2,
		"SERVER_KEY":   3,
		"FILE_REQ":     4,
		"FILE_DATA":    5,
		"FIN":          6,
		"RETRANS":      7,
	}
)

func (x IFTPP_Flag) Enum() *IFTPP_Flag {
	p := new(IFTPP_Flag)
	*p = x
	return p
}

func (x IFTPP_Flag) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IFTPP_Flag) Descriptor() protoreflect.EnumDescriptor {
	return file_iftpp_proto_enumTypes[0].Descriptor()
}

func (IFTPP_Flag) Type() protoreflect.EnumType {
	return &file_iftpp_proto_enumTypes[0]
}

func (x IFTPP_Flag) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IFTPP_Flag.Descriptor instead.
func (IFTPP_Flag) EnumDescriptor() ([]byte, []int) {
	return file_iftpp_proto_rawDescGZIP(), []int{0, 0}
}

//Proto for Insecure File Transfer Protocol over Ping
type IFTPP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId int32      `protobuf:"varint,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`              // unique int will be generated for each session
	Payload   []byte     `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`                                    // the actual data being sent will go here
	Checksum  []byte     `protobuf:"bytes,3,opt,name=checksum,proto3" json:"checksum,omitempty"`                                  // last 8 digits of payload SHA1 will go here
	TypeFlag  IFTPP_Flag `protobuf:"varint,4,opt,name=type_flag,json=typeFlag,proto3,enum=IFTPP_Flag" json:"type_flag,omitempty"` // flag to say what type of data is in the payload
}

func (x *IFTPP) Reset() {
	*x = IFTPP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iftpp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IFTPP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IFTPP) ProtoMessage() {}

func (x *IFTPP) ProtoReflect() protoreflect.Message {
	mi := &file_iftpp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IFTPP.ProtoReflect.Descriptor instead.
func (*IFTPP) Descriptor() ([]byte, []int) {
	return file_iftpp_proto_rawDescGZIP(), []int{0}
}

func (x *IFTPP) GetSessionId() int32 {
	if x != nil {
		return x.SessionId
	}
	return 0
}

func (x *IFTPP) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *IFTPP) GetChecksum() []byte {
	if x != nil {
		return x.Checksum
	}
	return nil
}

func (x *IFTPP) GetTypeFlag() IFTPP_Flag {
	if x != nil {
		return x.TypeFlag
	}
	return IFTPP_SESSION_INIT
}

var File_iftpp_proto protoreflect.FileDescriptor

var file_iftpp_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x69, 0x66, 0x74, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfc, 0x01,
	0x0a, 0x05, 0x49, 0x46, 0x54, 0x50, 0x50, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x12, 0x28, 0x0a, 0x09,
	0x74, 0x79, 0x70, 0x65, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0b, 0x2e, 0x49, 0x46, 0x54, 0x50, 0x50, 0x2e, 0x46, 0x6c, 0x61, 0x67, 0x52, 0x08, 0x74, 0x79,
	0x70, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x22, 0x74, 0x0a, 0x04, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x10,
	0x0a, 0x0c, 0x53, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x4e, 0x49, 0x54, 0x10, 0x00,
	0x12, 0x07, 0x0a, 0x03, 0x41, 0x43, 0x4b, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x4c, 0x49,
	0x45, 0x4e, 0x54, 0x5f, 0x4b, 0x45, 0x59, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x45, 0x52,
	0x56, 0x45, 0x52, 0x5f, 0x4b, 0x45, 0x59, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x49, 0x4c,
	0x45, 0x5f, 0x52, 0x45, 0x51, 0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x46, 0x49, 0x4c, 0x45, 0x5f,
	0x44, 0x41, 0x54, 0x41, 0x10, 0x05, 0x12, 0x07, 0x0a, 0x03, 0x46, 0x49, 0x4e, 0x10, 0x06, 0x12,
	0x0b, 0x0a, 0x07, 0x52, 0x45, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x10, 0x07, 0x42, 0x22, 0x5a, 0x20,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x33, 0x2f, 0x69, 0x66, 0x74, 0x70, 0x70, 0x2f, 0x70, 0x62, 0x75, 0x66,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_iftpp_proto_rawDescOnce sync.Once
	file_iftpp_proto_rawDescData = file_iftpp_proto_rawDesc
)

func file_iftpp_proto_rawDescGZIP() []byte {
	file_iftpp_proto_rawDescOnce.Do(func() {
		file_iftpp_proto_rawDescData = protoimpl.X.CompressGZIP(file_iftpp_proto_rawDescData)
	})
	return file_iftpp_proto_rawDescData
}

var file_iftpp_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_iftpp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_iftpp_proto_goTypes = []interface{}{
	(IFTPP_Flag)(0), // 0: IFTPP.Flag
	(*IFTPP)(nil),   // 1: IFTPP
}
var file_iftpp_proto_depIdxs = []int32{
	0, // 0: IFTPP.type_flag:type_name -> IFTPP.Flag
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_iftpp_proto_init() }
func file_iftpp_proto_init() {
	if File_iftpp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_iftpp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IFTPP); i {
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
			RawDescriptor: file_iftpp_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_iftpp_proto_goTypes,
		DependencyIndexes: file_iftpp_proto_depIdxs,
		EnumInfos:         file_iftpp_proto_enumTypes,
		MessageInfos:      file_iftpp_proto_msgTypes,
	}.Build()
	File_iftpp_proto = out.File
	file_iftpp_proto_rawDesc = nil
	file_iftpp_proto_goTypes = nil
	file_iftpp_proto_depIdxs = nil
}
