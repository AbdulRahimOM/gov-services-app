// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.27.1
// source: pb/protobuf/appointments.proto

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

type CreateChildOfficeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdminID             int32                `protobuf:"varint,1,opt,name=adminID,proto3" json:"adminID,omitempty"`
	ProposedChildOffice *ProposedChildOffice `protobuf:"bytes,2,opt,name=proposedChildOffice,proto3" json:"proposedChildOffice,omitempty"`
}

func (x *CreateChildOfficeRequest) Reset() {
	*x = CreateChildOfficeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_protobuf_appointments_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateChildOfficeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateChildOfficeRequest) ProtoMessage() {}

func (x *CreateChildOfficeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_protobuf_appointments_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateChildOfficeRequest.ProtoReflect.Descriptor instead.
func (*CreateChildOfficeRequest) Descriptor() ([]byte, []int) {
	return file_pb_protobuf_appointments_proto_rawDescGZIP(), []int{0}
}

func (x *CreateChildOfficeRequest) GetAdminID() int32 {
	if x != nil {
		return x.AdminID
	}
	return 0
}

func (x *CreateChildOfficeRequest) GetProposedChildOffice() *ProposedChildOffice {
	if x != nil {
		return x.ProposedChildOffice
	}
	return nil
}

type CreateChildOfficeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChildOfficeID int32 `protobuf:"varint,1,opt,name=childOfficeID,proto3" json:"childOfficeID,omitempty"`
}

func (x *CreateChildOfficeResponse) Reset() {
	*x = CreateChildOfficeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_protobuf_appointments_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateChildOfficeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateChildOfficeResponse) ProtoMessage() {}

func (x *CreateChildOfficeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_protobuf_appointments_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateChildOfficeResponse.ProtoReflect.Descriptor instead.
func (*CreateChildOfficeResponse) Descriptor() ([]byte, []int) {
	return file_pb_protobuf_appointments_proto_rawDescGZIP(), []int{1}
}

func (x *CreateChildOfficeResponse) GetChildOfficeID() int32 {
	if x != nil {
		return x.ChildOfficeID
	}
	return 0
}

type ProposedChildOffice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *ProposedChildOffice) Reset() {
	*x = ProposedChildOffice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_protobuf_appointments_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProposedChildOffice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProposedChildOffice) ProtoMessage() {}

func (x *ProposedChildOffice) ProtoReflect() protoreflect.Message {
	mi := &file_pb_protobuf_appointments_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProposedChildOffice.ProtoReflect.Descriptor instead.
func (*ProposedChildOffice) Descriptor() ([]byte, []int) {
	return file_pb_protobuf_appointments_proto_rawDescGZIP(), []int{2}
}

func (x *ProposedChildOffice) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProposedChildOffice) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type AttenderAppointmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Appointer *Appointer `protobuf:"bytes,1,opt,name=appointer,proto3" json:"appointer,omitempty"`
	Appointee *Appointee `protobuf:"bytes,2,opt,name=appointee,proto3" json:"appointee,omitempty"`
}

func (x *AttenderAppointmentRequest) Reset() {
	*x = AttenderAppointmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_protobuf_appointments_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttenderAppointmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttenderAppointmentRequest) ProtoMessage() {}

func (x *AttenderAppointmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_protobuf_appointments_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttenderAppointmentRequest.ProtoReflect.Descriptor instead.
func (*AttenderAppointmentRequest) Descriptor() ([]byte, []int) {
	return file_pb_protobuf_appointments_proto_rawDescGZIP(), []int{3}
}

func (x *AttenderAppointmentRequest) GetAppointer() *Appointer {
	if x != nil {
		return x.Appointer
	}
	return nil
}

func (x *AttenderAppointmentRequest) GetAppointee() *Appointee {
	if x != nil {
		return x.Appointee
	}
	return nil
}

type OfficeHeadAppointmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Appointer     *Appointer `protobuf:"bytes,1,opt,name=appointer,proto3" json:"appointer,omitempty"`
	Appointee     *Appointee `protobuf:"bytes,2,opt,name=appointee,proto3" json:"appointee,omitempty"`
	ChildOfficeID int32      `protobuf:"varint,3,opt,name=childOfficeID,proto3" json:"childOfficeID,omitempty"`
}

func (x *OfficeHeadAppointmentRequest) Reset() {
	*x = OfficeHeadAppointmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_protobuf_appointments_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OfficeHeadAppointmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OfficeHeadAppointmentRequest) ProtoMessage() {}

func (x *OfficeHeadAppointmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_protobuf_appointments_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OfficeHeadAppointmentRequest.ProtoReflect.Descriptor instead.
func (*OfficeHeadAppointmentRequest) Descriptor() ([]byte, []int) {
	return file_pb_protobuf_appointments_proto_rawDescGZIP(), []int{4}
}

func (x *OfficeHeadAppointmentRequest) GetAppointer() *Appointer {
	if x != nil {
		return x.Appointer
	}
	return nil
}

func (x *OfficeHeadAppointmentRequest) GetAppointee() *Appointee {
	if x != nil {
		return x.Appointee
	}
	return nil
}

func (x *OfficeHeadAppointmentRequest) GetChildOfficeID() int32 {
	if x != nil {
		return x.ChildOfficeID
	}
	return 0
}

type Appointer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Appointer) Reset() {
	*x = Appointer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_protobuf_appointments_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Appointer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Appointer) ProtoMessage() {}

func (x *Appointer) ProtoReflect() protoreflect.Message {
	mi := &file_pb_protobuf_appointments_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Appointer.ProtoReflect.Descriptor instead.
func (*Appointer) Descriptor() ([]byte, []int) {
	return file_pb_protobuf_appointments_proto_rawDescGZIP(), []int{5}
}

func (x *Appointer) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Appointee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName   string `protobuf:"bytes,1,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName    string `protobuf:"bytes,2,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Email       string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber string `protobuf:"bytes,4,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
}

func (x *Appointee) Reset() {
	*x = Appointee{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_protobuf_appointments_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Appointee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Appointee) ProtoMessage() {}

func (x *Appointee) ProtoReflect() protoreflect.Message {
	mi := &file_pb_protobuf_appointments_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Appointee.ProtoReflect.Descriptor instead.
func (*Appointee) Descriptor() ([]byte, []int) {
	return file_pb_protobuf_appointments_proto_rawDescGZIP(), []int{6}
}

func (x *Appointee) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Appointee) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Appointee) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Appointee) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

var File_pb_protobuf_appointments_proto protoreflect.FileDescriptor

var file_pb_protobuf_appointments_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x70,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7c, 0x0a,
	0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x4f, 0x66, 0x66, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x49, 0x44, 0x12, 0x46, 0x0a, 0x13, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x64, 0x43,
	0x68, 0x69, 0x6c, 0x64, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x64, 0x43, 0x68, 0x69, 0x6c, 0x64,
	0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x52, 0x13, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x64,
	0x43, 0x68, 0x69, 0x6c, 0x64, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x22, 0x41, 0x0a, 0x19, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x68, 0x69, 0x6c,
	0x64, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0d, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x49, 0x44, 0x22, 0x43,
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x64, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x4f,
	0x66, 0x66, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x22, 0x70, 0x0a, 0x1a, 0x41, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x41,
	0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x28, 0x0a, 0x09, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x52, 0x09, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x09, 0x61,
	0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x65, 0x52, 0x09, 0x61, 0x70, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x65, 0x65, 0x22, 0x98, 0x01, 0x0a, 0x1c, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65,
	0x48, 0x65, 0x61, 0x64, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x09, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x70, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x09, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x12, 0x28, 0x0a, 0x09, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x65, 0x52,
	0x09, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x68,
	0x69, 0x6c, 0x64, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0d, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x49, 0x44,
	0x22, 0x1b, 0x0a, 0x09, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x7d, 0x0a,
	0x09, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x32, 0xd0, 0x02, 0x0a,
	0x12, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x0f, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x41, 0x74,
	0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x41, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x4a, 0x0a, 0x11, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65,
	0x12, 0x19, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x4f, 0x66,
	0x66, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x16, 0x41, 0x70, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x48, 0x65, 0x61,
	0x64, 0x12, 0x1d, 0x2e, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x48, 0x65, 0x61, 0x64, 0x41, 0x70,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x55, 0x0a, 0x1c, 0x41, 0x70, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x44, 0x65,
	0x70, 0x75, 0x74, 0x79, 0x48, 0x65, 0x61, 0x64, 0x12, 0x1d, 0x2e, 0x4f, 0x66, 0x66, 0x69, 0x63,
	0x65, 0x48, 0x65, 0x61, 0x64, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42,
	0x13, 0x5a, 0x11, 0x2e, 0x2f, 0x70, 0x62, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x64, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_protobuf_appointments_proto_rawDescOnce sync.Once
	file_pb_protobuf_appointments_proto_rawDescData = file_pb_protobuf_appointments_proto_rawDesc
)

func file_pb_protobuf_appointments_proto_rawDescGZIP() []byte {
	file_pb_protobuf_appointments_proto_rawDescOnce.Do(func() {
		file_pb_protobuf_appointments_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_protobuf_appointments_proto_rawDescData)
	})
	return file_pb_protobuf_appointments_proto_rawDescData
}

var file_pb_protobuf_appointments_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pb_protobuf_appointments_proto_goTypes = []interface{}{
	(*CreateChildOfficeRequest)(nil),     // 0: CreateChildOfficeRequest
	(*CreateChildOfficeResponse)(nil),    // 1: CreateChildOfficeResponse
	(*ProposedChildOffice)(nil),          // 2: proposedChildOffice
	(*AttenderAppointmentRequest)(nil),   // 3: AttenderAppointmentRequest
	(*OfficeHeadAppointmentRequest)(nil), // 4: OfficeHeadAppointmentRequest
	(*Appointer)(nil),                    // 5: appointer
	(*Appointee)(nil),                    // 6: appointee
	(*emptypb.Empty)(nil),                // 7: google.protobuf.Empty
}
var file_pb_protobuf_appointments_proto_depIdxs = []int32{
	2, // 0: CreateChildOfficeRequest.proposedChildOffice:type_name -> proposedChildOffice
	5, // 1: AttenderAppointmentRequest.appointer:type_name -> appointer
	6, // 2: AttenderAppointmentRequest.appointee:type_name -> appointee
	5, // 3: OfficeHeadAppointmentRequest.appointer:type_name -> appointer
	6, // 4: OfficeHeadAppointmentRequest.appointee:type_name -> appointee
	3, // 5: AppointmentService.AppointAttender:input_type -> AttenderAppointmentRequest
	0, // 6: AppointmentService.CreateChildOffice:input_type -> CreateChildOfficeRequest
	4, // 7: AppointmentService.AppointChildOfficeHead:input_type -> OfficeHeadAppointmentRequest
	4, // 8: AppointmentService.AppointChildOfficeDeputyHead:input_type -> OfficeHeadAppointmentRequest
	7, // 9: AppointmentService.AppointAttender:output_type -> google.protobuf.Empty
	1, // 10: AppointmentService.CreateChildOffice:output_type -> CreateChildOfficeResponse
	7, // 11: AppointmentService.AppointChildOfficeHead:output_type -> google.protobuf.Empty
	7, // 12: AppointmentService.AppointChildOfficeDeputyHead:output_type -> google.protobuf.Empty
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_pb_protobuf_appointments_proto_init() }
func file_pb_protobuf_appointments_proto_init() {
	if File_pb_protobuf_appointments_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_protobuf_appointments_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateChildOfficeRequest); i {
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
		file_pb_protobuf_appointments_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateChildOfficeResponse); i {
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
		file_pb_protobuf_appointments_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProposedChildOffice); i {
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
		file_pb_protobuf_appointments_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttenderAppointmentRequest); i {
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
		file_pb_protobuf_appointments_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OfficeHeadAppointmentRequest); i {
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
		file_pb_protobuf_appointments_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Appointer); i {
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
		file_pb_protobuf_appointments_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Appointee); i {
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
			RawDescriptor: file_pb_protobuf_appointments_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_protobuf_appointments_proto_goTypes,
		DependencyIndexes: file_pb_protobuf_appointments_proto_depIdxs,
		MessageInfos:      file_pb_protobuf_appointments_proto_msgTypes,
	}.Build()
	File_pb_protobuf_appointments_proto = out.File
	file_pb_protobuf_appointments_proto_rawDesc = nil
	file_pb_protobuf_appointments_proto_goTypes = nil
	file_pb_protobuf_appointments_proto_depIdxs = nil
}
