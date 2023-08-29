// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.2
// source: pb_gen/Room.proto

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

type GetRoomsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GetRoomResponse []*GetRoomResponse `protobuf:"bytes,1,rep,name=get_room_response,json=getRoomResponse,proto3" json:"get_room_response,omitempty"`
}

func (x *GetRoomsResponse) Reset() {
	*x = GetRoomsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_gen_Room_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRoomsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoomsResponse) ProtoMessage() {}

func (x *GetRoomsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_gen_Room_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoomsResponse.ProtoReflect.Descriptor instead.
func (*GetRoomsResponse) Descriptor() ([]byte, []int) {
	return file_pb_gen_Room_proto_rawDescGZIP(), []int{0}
}

func (x *GetRoomsResponse) GetGetRoomResponse() []*GetRoomResponse {
	if x != nil {
		return x.GetRoomResponse
	}
	return nil
}

type GetRoomResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	RoomName      string `protobuf:"bytes,2,opt,name=room_name,json=roomName,proto3" json:"room_name,omitempty"`
	MaximumPerson int64  `protobuf:"varint,3,opt,name=maximum_person,json=maximumPerson,proto3" json:"maximum_person,omitempty"`
}

func (x *GetRoomResponse) Reset() {
	*x = GetRoomResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_gen_Room_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRoomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoomResponse) ProtoMessage() {}

func (x *GetRoomResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_gen_Room_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoomResponse.ProtoReflect.Descriptor instead.
func (*GetRoomResponse) Descriptor() ([]byte, []int) {
	return file_pb_gen_Room_proto_rawDescGZIP(), []int{1}
}

func (x *GetRoomResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetRoomResponse) GetRoomName() string {
	if x != nil {
		return x.RoomName
	}
	return ""
}

func (x *GetRoomResponse) GetMaximumPerson() int64 {
	if x != nil {
		return x.MaximumPerson
	}
	return 0
}

type GetRoomsWithBookingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rooms []*GetRoomWithBookingResponse `protobuf:"bytes,1,rep,name=rooms,proto3" json:"rooms,omitempty"`
}

func (x *GetRoomsWithBookingsResponse) Reset() {
	*x = GetRoomsWithBookingsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_gen_Room_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRoomsWithBookingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoomsWithBookingsResponse) ProtoMessage() {}

func (x *GetRoomsWithBookingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_gen_Room_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoomsWithBookingsResponse.ProtoReflect.Descriptor instead.
func (*GetRoomsWithBookingsResponse) Descriptor() ([]byte, []int) {
	return file_pb_gen_Room_proto_rawDescGZIP(), []int{2}
}

func (x *GetRoomsWithBookingsResponse) GetRooms() []*GetRoomWithBookingResponse {
	if x != nil {
		return x.Rooms
	}
	return nil
}

type GetRoomWithBookingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	RoomName      string                `protobuf:"bytes,2,opt,name=room_name,json=roomName,proto3" json:"room_name,omitempty"`
	MaximumPerson int64                 `protobuf:"varint,3,opt,name=maximum_person,json=maximumPerson,proto3" json:"maximum_person,omitempty"`
	Bookings      []*GetBookingResponse `protobuf:"bytes,4,rep,name=bookings,proto3" json:"bookings,omitempty"`
}

func (x *GetRoomWithBookingResponse) Reset() {
	*x = GetRoomWithBookingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_gen_Room_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRoomWithBookingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoomWithBookingResponse) ProtoMessage() {}

func (x *GetRoomWithBookingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_gen_Room_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoomWithBookingResponse.ProtoReflect.Descriptor instead.
func (*GetRoomWithBookingResponse) Descriptor() ([]byte, []int) {
	return file_pb_gen_Room_proto_rawDescGZIP(), []int{3}
}

func (x *GetRoomWithBookingResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetRoomWithBookingResponse) GetRoomName() string {
	if x != nil {
		return x.RoomName
	}
	return ""
}

func (x *GetRoomWithBookingResponse) GetMaximumPerson() int64 {
	if x != nil {
		return x.MaximumPerson
	}
	return 0
}

func (x *GetRoomWithBookingResponse) GetBookings() []*GetBookingResponse {
	if x != nil {
		return x.Bookings
	}
	return nil
}

type GetBookingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	RoomId       int64  `protobuf:"varint,2,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	UserId       int64  `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	AmountPerson int64  `protobuf:"varint,4,opt,name=amount_person,json=amountPerson,proto3" json:"amount_person,omitempty"`
	StartDate    string `protobuf:"bytes,5,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate      string `protobuf:"bytes,6,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
}

func (x *GetBookingResponse) Reset() {
	*x = GetBookingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_gen_Room_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBookingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookingResponse) ProtoMessage() {}

func (x *GetBookingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_gen_Room_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookingResponse.ProtoReflect.Descriptor instead.
func (*GetBookingResponse) Descriptor() ([]byte, []int) {
	return file_pb_gen_Room_proto_rawDescGZIP(), []int{4}
}

func (x *GetBookingResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetBookingResponse) GetRoomId() int64 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

func (x *GetBookingResponse) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetBookingResponse) GetAmountPerson() int64 {
	if x != nil {
		return x.AmountPerson
	}
	return 0
}

func (x *GetBookingResponse) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *GetBookingResponse) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

var File_pb_gen_Room_proto protoreflect.FileDescriptor

var file_pb_gen_Room_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x62, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x52, 0x6f, 0x6f, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x53, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x11, 0x67, 0x65, 0x74, 0x5f,
	0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0f, 0x67, 0x65, 0x74, 0x52, 0x6f, 0x6f,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x65, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x6d, 0x61, 0x78,
	0x69, 0x6d, 0x75, 0x6d, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0d, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x22, 0x54, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x57, 0x69, 0x74, 0x68,
	0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x34, 0x0a, 0x05, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x57, 0x69, 0x74, 0x68,
	0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52,
	0x05, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x22, 0xa4, 0x01, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x52, 0x6f,
	0x6f, 0x6d, 0x57, 0x69, 0x74, 0x68, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x5f, 0x70, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x6d, 0x61, 0x78, 0x69,
	0x6d, 0x75, 0x6d, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x08, 0x62, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x62,
	0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x22, 0xb5, 0x01,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e,
	0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e,
	0x64, 0x44, 0x61, 0x74, 0x65, 0x32, 0x93, 0x01, 0x0a, 0x05, 0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x12,
	0x38, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x6d,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x14, 0x47, 0x65, 0x74,
	0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x57, 0x69, 0x74, 0x68, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x20, 0x2e, 0x70, 0x62, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x57, 0x69, 0x74, 0x68, 0x42, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_gen_Room_proto_rawDescOnce sync.Once
	file_pb_gen_Room_proto_rawDescData = file_pb_gen_Room_proto_rawDesc
)

func file_pb_gen_Room_proto_rawDescGZIP() []byte {
	file_pb_gen_Room_proto_rawDescOnce.Do(func() {
		file_pb_gen_Room_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_gen_Room_proto_rawDescData)
	})
	return file_pb_gen_Room_proto_rawDescData
}

var file_pb_gen_Room_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pb_gen_Room_proto_goTypes = []interface{}{
	(*GetRoomsResponse)(nil),             // 0: pb.GetRoomsResponse
	(*GetRoomResponse)(nil),              // 1: pb.GetRoomResponse
	(*GetRoomsWithBookingsResponse)(nil), // 2: pb.GetRoomsWithBookingsResponse
	(*GetRoomWithBookingResponse)(nil),   // 3: pb.GetRoomWithBookingResponse
	(*GetBookingResponse)(nil),           // 4: pb.GetBookingResponse
	(*emptypb.Empty)(nil),                // 5: google.protobuf.Empty
}
var file_pb_gen_Room_proto_depIdxs = []int32{
	1, // 0: pb.GetRoomsResponse.get_room_response:type_name -> pb.GetRoomResponse
	3, // 1: pb.GetRoomsWithBookingsResponse.rooms:type_name -> pb.GetRoomWithBookingResponse
	4, // 2: pb.GetRoomWithBookingResponse.bookings:type_name -> pb.GetBookingResponse
	5, // 3: pb.Rooms.GetRooms:input_type -> google.protobuf.Empty
	5, // 4: pb.Rooms.GetRoomsWithBookings:input_type -> google.protobuf.Empty
	0, // 5: pb.Rooms.GetRooms:output_type -> pb.GetRoomsResponse
	2, // 6: pb.Rooms.GetRoomsWithBookings:output_type -> pb.GetRoomsWithBookingsResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pb_gen_Room_proto_init() }
func file_pb_gen_Room_proto_init() {
	if File_pb_gen_Room_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_gen_Room_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRoomsResponse); i {
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
		file_pb_gen_Room_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRoomResponse); i {
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
		file_pb_gen_Room_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRoomsWithBookingsResponse); i {
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
		file_pb_gen_Room_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRoomWithBookingResponse); i {
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
		file_pb_gen_Room_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBookingResponse); i {
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
			RawDescriptor: file_pb_gen_Room_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_gen_Room_proto_goTypes,
		DependencyIndexes: file_pb_gen_Room_proto_depIdxs,
		MessageInfos:      file_pb_gen_Room_proto_msgTypes,
	}.Build()
	File_pb_gen_Room_proto = out.File
	file_pb_gen_Room_proto_rawDesc = nil
	file_pb_gen_Room_proto_goTypes = nil
	file_pb_gen_Room_proto_depIdxs = nil
}
