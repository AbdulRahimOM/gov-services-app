// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.27.1
// source: pb/protobuf/adminAccount.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AdminUpdatePasswordUsingOldPwRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdminId     int32  `protobuf:"varint,1,opt,name=adminId,proto3" json:"adminId,omitempty"`
	OldPassword string `protobuf:"bytes,2,opt,name=oldPassword,proto3" json:"oldPassword,omitempty"`
	NewPassword string `protobuf:"bytes,3,opt,name=newPassword,proto3" json:"newPassword,omitempty"`
}

func (x *AdminUpdatePasswordUsingOldPwRequest) Reset() {
	*x = AdminUpdatePasswordUsingOldPwRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_protobuf_adminAccount_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminUpdatePasswordUsingOldPwRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminUpdatePasswordUsingOldPwRequest) ProtoMessage() {}

func (x *AdminUpdatePasswordUsingOldPwRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_protobuf_adminAccount_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminUpdatePasswordUsingOldPwRequest.ProtoReflect.Descriptor instead.
func (*AdminUpdatePasswordUsingOldPwRequest) Descriptor() ([]byte, []int) {
	return file_pb_protobuf_adminAccount_proto_rawDescGZIP(), []int{0}
}

func (x *AdminUpdatePasswordUsingOldPwRequest) GetAdminId() int32 {
	if x != nil {
		return x.AdminId
	}
	return 0
}

func (x *AdminUpdatePasswordUsingOldPwRequest) GetOldPassword() string {
	if x != nil {
		return x.OldPassword
	}
	return ""
}

func (x *AdminUpdatePasswordUsingOldPwRequest) GetNewPassword() string {
	if x != nil {
		return x.NewPassword
	}
	return ""
}

type AdminGetProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdminId int32 `protobuf:"varint,1,opt,name=adminId,proto3" json:"adminId,omitempty"`
}

func (x *AdminGetProfileRequest) Reset() {
	*x = AdminGetProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_protobuf_adminAccount_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminGetProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminGetProfileRequest) ProtoMessage() {}

func (x *AdminGetProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_protobuf_adminAccount_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminGetProfileRequest.ProtoReflect.Descriptor instead.
func (*AdminGetProfileRequest) Descriptor() ([]byte, []int) {
	return file_pb_protobuf_adminAccount_proto_rawDescGZIP(), []int{1}
}

func (x *AdminGetProfileRequest) GetAdminId() int32 {
	if x != nil {
		return x.AdminId
	}
	return 0
}

type AdminGetProfileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName   string `protobuf:"bytes,1,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName    string `protobuf:"bytes,2,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Email       string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Address     string `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	Pincode     string `protobuf:"bytes,5,opt,name=pincode,proto3" json:"pincode,omitempty"`
	PhoneNumber string `protobuf:"bytes,6,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	Designation string `protobuf:"bytes,7,opt,name=designation,proto3" json:"designation,omitempty"`
	DeptId      int32  `protobuf:"varint,8,opt,name=deptId,proto3" json:"deptId,omitempty"`
	RankId      int32  `protobuf:"varint,9,opt,name=rankId,proto3" json:"rankId,omitempty"`
	Username    string `protobuf:"bytes,10,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *AdminGetProfileResponse) Reset() {
	*x = AdminGetProfileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_protobuf_adminAccount_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminGetProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminGetProfileResponse) ProtoMessage() {}

func (x *AdminGetProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_protobuf_adminAccount_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminGetProfileResponse.ProtoReflect.Descriptor instead.
func (*AdminGetProfileResponse) Descriptor() ([]byte, []int) {
	return file_pb_protobuf_adminAccount_proto_rawDescGZIP(), []int{2}
}

func (x *AdminGetProfileResponse) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *AdminGetProfileResponse) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *AdminGetProfileResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *AdminGetProfileResponse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *AdminGetProfileResponse) GetPincode() string {
	if x != nil {
		return x.Pincode
	}
	return ""
}

func (x *AdminGetProfileResponse) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *AdminGetProfileResponse) GetDesignation() string {
	if x != nil {
		return x.Designation
	}
	return ""
}

func (x *AdminGetProfileResponse) GetDeptId() int32 {
	if x != nil {
		return x.DeptId
	}
	return 0
}

func (x *AdminGetProfileResponse) GetRankId() int32 {
	if x != nil {
		return x.RankId
	}
	return 0
}

func (x *AdminGetProfileResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type AdminLoginViaPasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *AdminLoginViaPasswordRequest) Reset() {
	*x = AdminLoginViaPasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_protobuf_adminAccount_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminLoginViaPasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminLoginViaPasswordRequest) ProtoMessage() {}

func (x *AdminLoginViaPasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_protobuf_adminAccount_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminLoginViaPasswordRequest.ProtoReflect.Descriptor instead.
func (*AdminLoginViaPasswordRequest) Descriptor() ([]byte, []int) {
	return file_pb_protobuf_adminAccount_proto_rawDescGZIP(), []int{3}
}

func (x *AdminLoginViaPasswordRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AdminLoginViaPasswordRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type AdminLoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message      string               `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Token        string               `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	AdminDetails *LoggingAdminDetails `protobuf:"bytes,3,opt,name=adminDetails,proto3" json:"adminDetails,omitempty"`
}

func (x *AdminLoginResponse) Reset() {
	*x = AdminLoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_protobuf_adminAccount_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminLoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminLoginResponse) ProtoMessage() {}

func (x *AdminLoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_protobuf_adminAccount_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminLoginResponse.ProtoReflect.Descriptor instead.
func (*AdminLoginResponse) Descriptor() ([]byte, []int) {
	return file_pb_protobuf_adminAccount_proto_rawDescGZIP(), []int{4}
}

func (x *AdminLoginResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *AdminLoginResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *AdminLoginResponse) GetAdminDetails() *LoggingAdminDetails {
	if x != nil {
		return x.AdminDetails
	}
	return nil
}

type LoggingAdminDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName   string `protobuf:"bytes,2,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName    string `protobuf:"bytes,3,opt,name=lastName,proto3" json:"lastName,omitempty"`
	PhoneNumber string `protobuf:"bytes,4,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	DeptId      int32  `protobuf:"varint,5,opt,name=deptId,proto3" json:"deptId,omitempty"`
	RankId      int32  `protobuf:"varint,6,opt,name=rankId,proto3" json:"rankId,omitempty"`
	Designation string `protobuf:"bytes,7,opt,name=designation,proto3" json:"designation,omitempty"`
}

func (x *LoggingAdminDetails) Reset() {
	*x = LoggingAdminDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_protobuf_adminAccount_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoggingAdminDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoggingAdminDetails) ProtoMessage() {}

func (x *LoggingAdminDetails) ProtoReflect() protoreflect.Message {
	mi := &file_pb_protobuf_adminAccount_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoggingAdminDetails.ProtoReflect.Descriptor instead.
func (*LoggingAdminDetails) Descriptor() ([]byte, []int) {
	return file_pb_protobuf_adminAccount_proto_rawDescGZIP(), []int{5}
}

func (x *LoggingAdminDetails) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LoggingAdminDetails) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *LoggingAdminDetails) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *LoggingAdminDetails) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *LoggingAdminDetails) GetDeptId() int32 {
	if x != nil {
		return x.DeptId
	}
	return 0
}

func (x *LoggingAdminDetails) GetRankId() int32 {
	if x != nil {
		return x.RankId
	}
	return 0
}

func (x *LoggingAdminDetails) GetDesignation() string {
	if x != nil {
		return x.Designation
	}
	return ""
}

var File_pb_protobuf_adminAccount_proto protoreflect.FileDescriptor

var file_pb_protobuf_adminAccount_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x01,
	0x0a, 0x24, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x55, 0x73, 0x69, 0x6e, 0x67, 0x4f, 0x6c, 0x64, 0x50, 0x77, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x49, 0x64,
	0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x6c, 0x64, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x6c, 0x64, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x22, 0x32, 0x0a, 0x16, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x49, 0x64, 0x22, 0xad, 0x02, 0x0a, 0x17, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x70, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x70, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06,
	0x64, 0x65, 0x70, 0x74, 0x49, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x64, 0x65,
	0x70, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x6e, 0x6b, 0x49, 0x64, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x61, 0x6e, 0x6b, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x56, 0x0a, 0x1c, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x56, 0x69, 0x61, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x22, 0x7e, 0x0a, 0x12, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x38, 0x0a, 0x0c, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x4c,
	0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x52, 0x0c, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x22, 0xd3, 0x01, 0x0a, 0x13, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x70, 0x74, 0x49, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x64, 0x65, 0x70, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x61, 0x6e, 0x6b, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x61,
	0x6e, 0x6b, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x88, 0x02, 0x0a, 0x13, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b,
	0x0a, 0x15, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x56, 0x69, 0x61, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1d, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x56, 0x69, 0x61, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0f, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x17,
	0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x47,
	0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x5e, 0x0a, 0x1d, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x55, 0x73, 0x69, 0x6e, 0x67, 0x4f, 0x6c, 0x64,
	0x50, 0x77, 0x12, 0x25, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x55, 0x73, 0x69, 0x6e, 0x67, 0x4f, 0x6c, 0x64,
	0x50, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x42, 0x13, 0x5a, 0x11, 0x2e, 0x2f, 0x70, 0x62, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x64, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_protobuf_adminAccount_proto_rawDescOnce sync.Once
	file_pb_protobuf_adminAccount_proto_rawDescData = file_pb_protobuf_adminAccount_proto_rawDesc
)

func file_pb_protobuf_adminAccount_proto_rawDescGZIP() []byte {
	file_pb_protobuf_adminAccount_proto_rawDescOnce.Do(func() {
		file_pb_protobuf_adminAccount_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_protobuf_adminAccount_proto_rawDescData)
	})
	return file_pb_protobuf_adminAccount_proto_rawDescData
}

var file_pb_protobuf_adminAccount_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pb_protobuf_adminAccount_proto_goTypes = []interface{}{
	(*AdminUpdatePasswordUsingOldPwRequest)(nil), // 0: AdminUpdatePasswordUsingOldPwRequest
	(*AdminGetProfileRequest)(nil),               // 1: AdminGetProfileRequest
	(*AdminGetProfileResponse)(nil),              // 2: AdminGetProfileResponse
	(*AdminLoginViaPasswordRequest)(nil),         // 3: AdminLoginViaPasswordRequest
	(*AdminLoginResponse)(nil),                   // 4: AdminLoginResponse
	(*LoggingAdminDetails)(nil),                  // 5: LoggingAdminDetails
	(*emptypb.Empty)(nil),                        // 6: google.protobuf.Empty
}
var file_pb_protobuf_adminAccount_proto_depIdxs = []int32{
	5, // 0: AdminLoginResponse.adminDetails:type_name -> LoggingAdminDetails
	3, // 1: AdminAccountService.AdminLoginViaPassword:input_type -> AdminLoginViaPasswordRequest
	1, // 2: AdminAccountService.AdminGetProfile:input_type -> AdminGetProfileRequest
	0, // 3: AdminAccountService.AdminUpdatePasswordUsingOldPw:input_type -> AdminUpdatePasswordUsingOldPwRequest
	4, // 4: AdminAccountService.AdminLoginViaPassword:output_type -> AdminLoginResponse
	2, // 5: AdminAccountService.AdminGetProfile:output_type -> AdminGetProfileResponse
	6, // 6: AdminAccountService.AdminUpdatePasswordUsingOldPw:output_type -> google.protobuf.Empty
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pb_protobuf_adminAccount_proto_init() }
func file_pb_protobuf_adminAccount_proto_init() {
	if File_pb_protobuf_adminAccount_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_protobuf_adminAccount_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminUpdatePasswordUsingOldPwRequest); i {
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
		file_pb_protobuf_adminAccount_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminGetProfileRequest); i {
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
		file_pb_protobuf_adminAccount_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminGetProfileResponse); i {
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
		file_pb_protobuf_adminAccount_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminLoginViaPasswordRequest); i {
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
		file_pb_protobuf_adminAccount_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminLoginResponse); i {
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
		file_pb_protobuf_adminAccount_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoggingAdminDetails); i {
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
			RawDescriptor: file_pb_protobuf_adminAccount_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_protobuf_adminAccount_proto_goTypes,
		DependencyIndexes: file_pb_protobuf_adminAccount_proto_depIdxs,
		MessageInfos:      file_pb_protobuf_adminAccount_proto_msgTypes,
	}.Build()
	File_pb_protobuf_adminAccount_proto = out.File
	file_pb_protobuf_adminAccount_proto_rawDesc = nil
	file_pb_protobuf_adminAccount_proto_goTypes = nil
	file_pb_protobuf_adminAccount_proto_depIdxs = nil
}
