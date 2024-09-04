// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.5
// source: proto/ticket.proto

package proto

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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email     string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_proto_ticket_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *User) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type PurchaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *PurchaseRequest) Reset() {
	*x = PurchaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PurchaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurchaseRequest) ProtoMessage() {}

func (x *PurchaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PurchaseRequest.ProtoReflect.Descriptor instead.
func (*PurchaseRequest) Descriptor() ([]byte, []int) {
	return file_proto_ticket_proto_rawDescGZIP(), []int{1}
}

func (x *PurchaseRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type PurchaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReceiptId  string `protobuf:"bytes,1,opt,name=receipt_id,json=receiptId,proto3" json:"receipt_id,omitempty"`
	From       string `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To         string `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	User       *User  `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
	PricePaid  int32  `protobuf:"varint,5,opt,name=price_paid,json=pricePaid,proto3" json:"price_paid,omitempty"`
	SeatNumber string `protobuf:"bytes,6,opt,name=seat_number,json=seatNumber,proto3" json:"seat_number,omitempty"`
	Section    string `protobuf:"bytes,7,opt,name=section,proto3" json:"section,omitempty"`
}

func (x *PurchaseResponse) Reset() {
	*x = PurchaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PurchaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurchaseResponse) ProtoMessage() {}

func (x *PurchaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PurchaseResponse.ProtoReflect.Descriptor instead.
func (*PurchaseResponse) Descriptor() ([]byte, []int) {
	return file_proto_ticket_proto_rawDescGZIP(), []int{2}
}

func (x *PurchaseResponse) GetReceiptId() string {
	if x != nil {
		return x.ReceiptId
	}
	return ""
}

func (x *PurchaseResponse) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *PurchaseResponse) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *PurchaseResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *PurchaseResponse) GetPricePaid() int32 {
	if x != nil {
		return x.PricePaid
	}
	return 0
}

func (x *PurchaseResponse) GetSeatNumber() string {
	if x != nil {
		return x.SeatNumber
	}
	return ""
}

func (x *PurchaseResponse) GetSection() string {
	if x != nil {
		return x.Section
	}
	return ""
}

type ReceiptRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *ReceiptRequest) Reset() {
	*x = ReceiptRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiptRequest) ProtoMessage() {}

func (x *ReceiptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiptRequest.ProtoReflect.Descriptor instead.
func (*ReceiptRequest) Descriptor() ([]byte, []int) {
	return file_proto_ticket_proto_rawDescGZIP(), []int{3}
}

func (x *ReceiptRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type ReceiptResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReceiptId  string `protobuf:"bytes,1,opt,name=receipt_id,json=receiptId,proto3" json:"receipt_id,omitempty"`
	From       string `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To         string `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	User       *User  `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
	PricePaid  int32  `protobuf:"varint,5,opt,name=price_paid,json=pricePaid,proto3" json:"price_paid,omitempty"`
	SeatNumber string `protobuf:"bytes,6,opt,name=seat_number,json=seatNumber,proto3" json:"seat_number,omitempty"`
	Section    string `protobuf:"bytes,7,opt,name=section,proto3" json:"section,omitempty"`
}

func (x *ReceiptResponse) Reset() {
	*x = ReceiptResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiptResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiptResponse) ProtoMessage() {}

func (x *ReceiptResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiptResponse.ProtoReflect.Descriptor instead.
func (*ReceiptResponse) Descriptor() ([]byte, []int) {
	return file_proto_ticket_proto_rawDescGZIP(), []int{4}
}

func (x *ReceiptResponse) GetReceiptId() string {
	if x != nil {
		return x.ReceiptId
	}
	return ""
}

func (x *ReceiptResponse) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *ReceiptResponse) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *ReceiptResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *ReceiptResponse) GetPricePaid() int32 {
	if x != nil {
		return x.PricePaid
	}
	return 0
}

func (x *ReceiptResponse) GetSeatNumber() string {
	if x != nil {
		return x.SeatNumber
	}
	return ""
}

func (x *ReceiptResponse) GetSection() string {
	if x != nil {
		return x.Section
	}
	return ""
}

type SectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Section string `protobuf:"bytes,1,opt,name=section,proto3" json:"section,omitempty"`
}

func (x *SectionRequest) Reset() {
	*x = SectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SectionRequest) ProtoMessage() {}

func (x *SectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SectionRequest.ProtoReflect.Descriptor instead.
func (*SectionRequest) Descriptor() ([]byte, []int) {
	return file_proto_ticket_proto_rawDescGZIP(), []int{5}
}

func (x *SectionRequest) GetSection() string {
	if x != nil {
		return x.Section
	}
	return ""
}

type SectionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*SeatAllocation `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *SectionResponse) Reset() {
	*x = SectionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SectionResponse) ProtoMessage() {}

func (x *SectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SectionResponse.ProtoReflect.Descriptor instead.
func (*SectionResponse) Descriptor() ([]byte, []int) {
	return file_proto_ticket_proto_rawDescGZIP(), []int{6}
}

func (x *SectionResponse) GetUsers() []*SeatAllocation {
	if x != nil {
		return x.Users
	}
	return nil
}

type SeatAllocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User       *User  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	SeatNumber string `protobuf:"bytes,2,opt,name=seat_number,json=seatNumber,proto3" json:"seat_number,omitempty"`
}

func (x *SeatAllocation) Reset() {
	*x = SeatAllocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeatAllocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeatAllocation) ProtoMessage() {}

func (x *SeatAllocation) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeatAllocation.ProtoReflect.Descriptor instead.
func (*SeatAllocation) Descriptor() ([]byte, []int) {
	return file_proto_ticket_proto_rawDescGZIP(), []int{7}
}

func (x *SeatAllocation) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *SeatAllocation) GetSeatNumber() string {
	if x != nil {
		return x.SeatNumber
	}
	return ""
}

type RemoveUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *RemoveUserRequest) Reset() {
	*x = RemoveUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveUserRequest) ProtoMessage() {}

func (x *RemoveUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveUserRequest.ProtoReflect.Descriptor instead.
func (*RemoveUserRequest) Descriptor() ([]byte, []int) {
	return file_proto_ticket_proto_rawDescGZIP(), []int{8}
}

func (x *RemoveUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type RemoveUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *RemoveUserResponse) Reset() {
	*x = RemoveUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveUserResponse) ProtoMessage() {}

func (x *RemoveUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveUserResponse.ProtoReflect.Descriptor instead.
func (*RemoveUserResponse) Descriptor() ([]byte, []int) {
	return file_proto_ticket_proto_rawDescGZIP(), []int{9}
}

func (x *RemoveUserResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type ModifySeatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email         string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	NewSeatNumber string `protobuf:"bytes,2,opt,name=new_seat_number,json=newSeatNumber,proto3" json:"new_seat_number,omitempty"`
}

func (x *ModifySeatRequest) Reset() {
	*x = ModifySeatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifySeatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifySeatRequest) ProtoMessage() {}

func (x *ModifySeatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifySeatRequest.ProtoReflect.Descriptor instead.
func (*ModifySeatRequest) Descriptor() ([]byte, []int) {
	return file_proto_ticket_proto_rawDescGZIP(), []int{10}
}

func (x *ModifySeatRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ModifySeatRequest) GetNewSeatNumber() string {
	if x != nil {
		return x.NewSeatNumber
	}
	return ""
}

type ModifySeatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *ModifySeatResponse) Reset() {
	*x = ModifySeatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifySeatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifySeatResponse) ProtoMessage() {}

func (x *ModifySeatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifySeatResponse.ProtoReflect.Descriptor instead.
func (*ModifySeatResponse) Descriptor() ([]byte, []int) {
	return file_proto_ticket_proto_rawDescGZIP(), []int{11}
}

func (x *ModifySeatResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_proto_ticket_proto protoreflect.FileDescriptor

var file_proto_ticket_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x22, 0x58, 0x0a, 0x04,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x33, 0x0a, 0x0f, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61,
	0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0xd1, 0x01, 0x0a, 0x10,
	0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66,
	0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x74, 0x6f, 0x12, 0x20, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x70,
	0x61, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x50, 0x61, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x61, 0x74, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x61, 0x74, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x26, 0x0a, 0x0e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0xd0, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x63, 0x65,
	0x69, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72,
	0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e,
	0x0a, 0x02, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x20,
	0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x74,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x61, 0x69, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x72, 0x69, 0x63, 0x65, 0x50, 0x61, 0x69, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x61, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x61, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x2a, 0x0a, 0x0e, 0x53, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3f, 0x0a, 0x0f, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x05, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x2e, 0x53, 0x65, 0x61, 0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0x53, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x74, 0x41,
	0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x65, 0x61, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x73, 0x65, 0x61, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x29, 0x0a, 0x11,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x2e, 0x0a, 0x12, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x51, 0x0a, 0x11, 0x4d, 0x6f, 0x64, 0x69, 0x66,
	0x79, 0x53, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x77, 0x5f, 0x73, 0x65, 0x61, 0x74, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x77,
	0x53, 0x65, 0x61, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x2e, 0x0a, 0x12, 0x4d, 0x6f,
	0x64, 0x69, 0x66, 0x79, 0x53, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0xe3, 0x02, 0x0a, 0x0d, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x0e,
	0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x17,
	0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x2e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3d, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x12,
	0x16, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x44, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x42, 0x79, 0x53, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x53,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0a, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0a, 0x4d,
	0x6f, 0x64, 0x69, 0x66, 0x79, 0x53, 0x65, 0x61, 0x74, 0x12, 0x19, 0x2e, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x53, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x4d, 0x6f,
	0x64, 0x69, 0x66, 0x79, 0x53, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c,
	0x69, 0x6e, 0x74, 0x6f, 0x70, 0x61, 0x75, 0x6c, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_ticket_proto_rawDescOnce sync.Once
	file_proto_ticket_proto_rawDescData = file_proto_ticket_proto_rawDesc
)

func file_proto_ticket_proto_rawDescGZIP() []byte {
	file_proto_ticket_proto_rawDescOnce.Do(func() {
		file_proto_ticket_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_ticket_proto_rawDescData)
	})
	return file_proto_ticket_proto_rawDescData
}

var file_proto_ticket_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_proto_ticket_proto_goTypes = []any{
	(*User)(nil),               // 0: ticket.User
	(*PurchaseRequest)(nil),    // 1: ticket.PurchaseRequest
	(*PurchaseResponse)(nil),   // 2: ticket.PurchaseResponse
	(*ReceiptRequest)(nil),     // 3: ticket.ReceiptRequest
	(*ReceiptResponse)(nil),    // 4: ticket.ReceiptResponse
	(*SectionRequest)(nil),     // 5: ticket.SectionRequest
	(*SectionResponse)(nil),    // 6: ticket.SectionResponse
	(*SeatAllocation)(nil),     // 7: ticket.SeatAllocation
	(*RemoveUserRequest)(nil),  // 8: ticket.RemoveUserRequest
	(*RemoveUserResponse)(nil), // 9: ticket.RemoveUserResponse
	(*ModifySeatRequest)(nil),  // 10: ticket.ModifySeatRequest
	(*ModifySeatResponse)(nil), // 11: ticket.ModifySeatResponse
}
var file_proto_ticket_proto_depIdxs = []int32{
	0,  // 0: ticket.PurchaseRequest.user:type_name -> ticket.User
	0,  // 1: ticket.PurchaseResponse.user:type_name -> ticket.User
	0,  // 2: ticket.ReceiptResponse.user:type_name -> ticket.User
	7,  // 3: ticket.SectionResponse.users:type_name -> ticket.SeatAllocation
	0,  // 4: ticket.SeatAllocation.user:type_name -> ticket.User
	1,  // 5: ticket.TicketService.PurchaseTicket:input_type -> ticket.PurchaseRequest
	3,  // 6: ticket.TicketService.GetReceipt:input_type -> ticket.ReceiptRequest
	5,  // 7: ticket.TicketService.GetUsersBySection:input_type -> ticket.SectionRequest
	8,  // 8: ticket.TicketService.RemoveUser:input_type -> ticket.RemoveUserRequest
	10, // 9: ticket.TicketService.ModifySeat:input_type -> ticket.ModifySeatRequest
	2,  // 10: ticket.TicketService.PurchaseTicket:output_type -> ticket.PurchaseResponse
	4,  // 11: ticket.TicketService.GetReceipt:output_type -> ticket.ReceiptResponse
	6,  // 12: ticket.TicketService.GetUsersBySection:output_type -> ticket.SectionResponse
	9,  // 13: ticket.TicketService.RemoveUser:output_type -> ticket.RemoveUserResponse
	11, // 14: ticket.TicketService.ModifySeat:output_type -> ticket.ModifySeatResponse
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_proto_ticket_proto_init() }
func file_proto_ticket_proto_init() {
	if File_proto_ticket_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_ticket_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*User); i {
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
		file_proto_ticket_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*PurchaseRequest); i {
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
		file_proto_ticket_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*PurchaseResponse); i {
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
		file_proto_ticket_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ReceiptRequest); i {
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
		file_proto_ticket_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ReceiptResponse); i {
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
		file_proto_ticket_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*SectionRequest); i {
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
		file_proto_ticket_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*SectionResponse); i {
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
		file_proto_ticket_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*SeatAllocation); i {
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
		file_proto_ticket_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*RemoveUserRequest); i {
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
		file_proto_ticket_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*RemoveUserResponse); i {
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
		file_proto_ticket_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*ModifySeatRequest); i {
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
		file_proto_ticket_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*ModifySeatResponse); i {
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
			RawDescriptor: file_proto_ticket_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_ticket_proto_goTypes,
		DependencyIndexes: file_proto_ticket_proto_depIdxs,
		MessageInfos:      file_proto_ticket_proto_msgTypes,
	}.Build()
	File_proto_ticket_proto = out.File
	file_proto_ticket_proto_rawDesc = nil
	file_proto_ticket_proto_goTypes = nil
	file_proto_ticket_proto_depIdxs = nil
}
