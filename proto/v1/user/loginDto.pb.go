// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: proto/files/v1/user/dto/loginDto.proto

package user

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

type LoginRequestModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"` // ชื่อผู้ใช้
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"` // รหัสผ่าน
}

func (x *LoginRequestModel) Reset() {
	*x = LoginRequestModel{}
	mi := &file_proto_files_v1_user_dto_loginDto_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginRequestModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequestModel) ProtoMessage() {}

func (x *LoginRequestModel) ProtoReflect() protoreflect.Message {
	mi := &file_proto_files_v1_user_dto_loginDto_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequestModel.ProtoReflect.Descriptor instead.
func (*LoginRequestModel) Descriptor() ([]byte, []int) {
	return file_proto_files_v1_user_dto_loginDto_proto_rawDescGZIP(), []int{0}
}

func (x *LoginRequestModel) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginRequestModel) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

// LoginDataListResponseModel สำหรับข้อมูลผู้ใช้ที่ล็อกอิน
type LoginDataListResponseModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid        string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`                                  // UUID ของผู้ใช้
	AccessToken string `protobuf:"bytes,2,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"` // Token สำหรับเข้าถึง
	Username    string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`                          // ชื่อผู้ใช้
	Nickname    string `protobuf:"bytes,4,opt,name=nickname,proto3" json:"nickname,omitempty"`                          // ชื่อเล่น
	PrefixName  string `protobuf:"bytes,5,opt,name=prefix_name,json=prefixName,proto3" json:"prefix_name,omitempty"`    // คำนำหน้าชื่อ
	Firstname   string `protobuf:"bytes,6,opt,name=firstname,proto3" json:"firstname,omitempty"`                        // ชื่อจริง
	Lastname    string `protobuf:"bytes,7,opt,name=lastname,proto3" json:"lastname,omitempty"`                          // นามสกุล
	Birthday    string `protobuf:"bytes,8,opt,name=birthday,proto3" json:"birthday,omitempty"`                          // วันเกิด (ใช้ string แทน time.Time)
}

func (x *LoginDataListResponseModel) Reset() {
	*x = LoginDataListResponseModel{}
	mi := &file_proto_files_v1_user_dto_loginDto_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginDataListResponseModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginDataListResponseModel) ProtoMessage() {}

func (x *LoginDataListResponseModel) ProtoReflect() protoreflect.Message {
	mi := &file_proto_files_v1_user_dto_loginDto_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginDataListResponseModel.ProtoReflect.Descriptor instead.
func (*LoginDataListResponseModel) Descriptor() ([]byte, []int) {
	return file_proto_files_v1_user_dto_loginDto_proto_rawDescGZIP(), []int{1}
}

func (x *LoginDataListResponseModel) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *LoginDataListResponseModel) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *LoginDataListResponseModel) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginDataListResponseModel) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *LoginDataListResponseModel) GetPrefixName() string {
	if x != nil {
		return x.PrefixName
	}
	return ""
}

func (x *LoginDataListResponseModel) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *LoginDataListResponseModel) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *LoginDataListResponseModel) GetBirthday() string {
	if x != nil {
		return x.Birthday
	}
	return ""
}

// LoginResponseModel สำหรับผลลัพธ์ของการล็อกอิน
type LoginResponseModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data        *LoginDataListResponseModel `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`                                  // ข้อมูลผู้ใช้
	StatusCode  int32                       `protobuf:"varint,2,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`   // รหัสสถานะ
	MessageDesc string                      `protobuf:"bytes,3,opt,name=message_desc,json=messageDesc,proto3" json:"message_desc,omitempty"` // คำอธิบายข้อความ
}

func (x *LoginResponseModel) Reset() {
	*x = LoginResponseModel{}
	mi := &file_proto_files_v1_user_dto_loginDto_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginResponseModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponseModel) ProtoMessage() {}

func (x *LoginResponseModel) ProtoReflect() protoreflect.Message {
	mi := &file_proto_files_v1_user_dto_loginDto_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponseModel.ProtoReflect.Descriptor instead.
func (*LoginResponseModel) Descriptor() ([]byte, []int) {
	return file_proto_files_v1_user_dto_loginDto_proto_rawDescGZIP(), []int{2}
}

func (x *LoginResponseModel) GetData() *LoginDataListResponseModel {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *LoginResponseModel) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *LoginResponseModel) GetMessageDesc() string {
	if x != nil {
		return x.MessageDesc
	}
	return ""
}

var File_proto_files_v1_user_dto_loginDto_proto protoreflect.FileDescriptor

var file_proto_files_v1_user_dto_loginDto_proto_rawDesc = []byte{
	0x0a, 0x26, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x64, 0x74, 0x6f, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x44,
	0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x22, 0x4b, 0x0a, 0x11, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x22, 0x82, 0x02, 0x0a, 0x1a, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x69,
	0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x69,
	0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x22, 0x94, 0x01, 0x0a, 0x12, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x3a, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x76, 0x31,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x44, 0x61,
	0x74, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x44, 0x65, 0x73, 0x63, 0x42, 0x16, 0x5a,
	0x14, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x3b, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_files_v1_user_dto_loginDto_proto_rawDescOnce sync.Once
	file_proto_files_v1_user_dto_loginDto_proto_rawDescData = file_proto_files_v1_user_dto_loginDto_proto_rawDesc
)

func file_proto_files_v1_user_dto_loginDto_proto_rawDescGZIP() []byte {
	file_proto_files_v1_user_dto_loginDto_proto_rawDescOnce.Do(func() {
		file_proto_files_v1_user_dto_loginDto_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_files_v1_user_dto_loginDto_proto_rawDescData)
	})
	return file_proto_files_v1_user_dto_loginDto_proto_rawDescData
}

var file_proto_files_v1_user_dto_loginDto_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_files_v1_user_dto_loginDto_proto_goTypes = []any{
	(*LoginRequestModel)(nil),          // 0: v1.service.LoginRequestModel
	(*LoginDataListResponseModel)(nil), // 1: v1.service.LoginDataListResponseModel
	(*LoginResponseModel)(nil),         // 2: v1.service.LoginResponseModel
}
var file_proto_files_v1_user_dto_loginDto_proto_depIdxs = []int32{
	1, // 0: v1.service.LoginResponseModel.data:type_name -> v1.service.LoginDataListResponseModel
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_files_v1_user_dto_loginDto_proto_init() }
func file_proto_files_v1_user_dto_loginDto_proto_init() {
	if File_proto_files_v1_user_dto_loginDto_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_files_v1_user_dto_loginDto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_files_v1_user_dto_loginDto_proto_goTypes,
		DependencyIndexes: file_proto_files_v1_user_dto_loginDto_proto_depIdxs,
		MessageInfos:      file_proto_files_v1_user_dto_loginDto_proto_msgTypes,
	}.Build()
	File_proto_files_v1_user_dto_loginDto_proto = out.File
	file_proto_files_v1_user_dto_loginDto_proto_rawDesc = nil
	file_proto_files_v1_user_dto_loginDto_proto_goTypes = nil
	file_proto_files_v1_user_dto_loginDto_proto_depIdxs = nil
}