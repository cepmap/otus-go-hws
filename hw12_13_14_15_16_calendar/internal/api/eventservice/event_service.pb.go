// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.2
// source: event_service.proto

package eventservice

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Event struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	Id               string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title            string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	DateStart        *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=date_start,json=dateStart,proto3" json:"date_start,omitempty"`
	DateEnd          *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=date_end,json=dateEnd,proto3" json:"date_end,omitempty"`
	Description      string                 `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	UserId           string                 `protobuf:"bytes,6,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	DateNotification *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=date_notification,json=dateNotification,proto3" json:"date_notification,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *Event) Reset() {
	*x = Event{}
	mi := &file_event_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_event_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_event_service_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Event) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Event) GetDateStart() *timestamppb.Timestamp {
	if x != nil {
		return x.DateStart
	}
	return nil
}

func (x *Event) GetDateEnd() *timestamppb.Timestamp {
	if x != nil {
		return x.DateEnd
	}
	return nil
}

func (x *Event) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Event) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Event) GetDateNotification() *timestamppb.Timestamp {
	if x != nil {
		return x.DateNotification
	}
	return nil
}

type AddEventRequest struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	Title            string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	DateStart        *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=date_start,json=dateStart,proto3" json:"date_start,omitempty"`
	DateEnd          *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=date_end,json=dateEnd,proto3" json:"date_end,omitempty"`
	Description      string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	UserId           string                 `protobuf:"bytes,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	DateNotification *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=date_notification,json=dateNotification,proto3" json:"date_notification,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *AddEventRequest) Reset() {
	*x = AddEventRequest{}
	mi := &file_event_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddEventRequest) ProtoMessage() {}

func (x *AddEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_event_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddEventRequest.ProtoReflect.Descriptor instead.
func (*AddEventRequest) Descriptor() ([]byte, []int) {
	return file_event_service_proto_rawDescGZIP(), []int{1}
}

func (x *AddEventRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *AddEventRequest) GetDateStart() *timestamppb.Timestamp {
	if x != nil {
		return x.DateStart
	}
	return nil
}

func (x *AddEventRequest) GetDateEnd() *timestamppb.Timestamp {
	if x != nil {
		return x.DateEnd
	}
	return nil
}

func (x *AddEventRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AddEventRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AddEventRequest) GetDateNotification() *timestamppb.Timestamp {
	if x != nil {
		return x.DateNotification
	}
	return nil
}

type AddEventResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	EventId       string                 `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddEventResponse) Reset() {
	*x = AddEventResponse{}
	mi := &file_event_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddEventResponse) ProtoMessage() {}

func (x *AddEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_event_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddEventResponse.ProtoReflect.Descriptor instead.
func (*AddEventResponse) Descriptor() ([]byte, []int) {
	return file_event_service_proto_rawDescGZIP(), []int{2}
}

func (x *AddEventResponse) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

type UpdateEventRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Event         *Event                 `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateEventRequest) Reset() {
	*x = UpdateEventRequest{}
	mi := &file_event_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEventRequest) ProtoMessage() {}

func (x *UpdateEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_event_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEventRequest.ProtoReflect.Descriptor instead.
func (*UpdateEventRequest) Descriptor() ([]byte, []int) {
	return file_event_service_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateEventRequest) GetEvent() *Event {
	if x != nil {
		return x.Event
	}
	return nil
}

type UpdateEventResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	X             *emptypb.Empty         `protobuf:"bytes,1,opt,name=_,proto3" json:"_,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateEventResponse) Reset() {
	*x = UpdateEventResponse{}
	mi := &file_event_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEventResponse) ProtoMessage() {}

func (x *UpdateEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_event_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEventResponse.ProtoReflect.Descriptor instead.
func (*UpdateEventResponse) Descriptor() ([]byte, []int) {
	return file_event_service_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateEventResponse) GetX() *emptypb.Empty {
	if x != nil {
		return x.X
	}
	return nil
}

type DeleteEventRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	EventId       string                 `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteEventRequest) Reset() {
	*x = DeleteEventRequest{}
	mi := &file_event_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEventRequest) ProtoMessage() {}

func (x *DeleteEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_event_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEventRequest.ProtoReflect.Descriptor instead.
func (*DeleteEventRequest) Descriptor() ([]byte, []int) {
	return file_event_service_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteEventRequest) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

type DeleteEventResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	X             *emptypb.Empty         `protobuf:"bytes,1,opt,name=_,proto3" json:"_,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteEventResponse) Reset() {
	*x = DeleteEventResponse{}
	mi := &file_event_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEventResponse) ProtoMessage() {}

func (x *DeleteEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_event_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEventResponse.ProtoReflect.Descriptor instead.
func (*DeleteEventResponse) Descriptor() ([]byte, []int) {
	return file_event_service_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteEventResponse) GetX() *emptypb.Empty {
	if x != nil {
		return x.X
	}
	return nil
}

type GetEventRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	EventId       string                 `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetEventRequest) Reset() {
	*x = GetEventRequest{}
	mi := &file_event_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventRequest) ProtoMessage() {}

func (x *GetEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_event_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventRequest.ProtoReflect.Descriptor instead.
func (*GetEventRequest) Descriptor() ([]byte, []int) {
	return file_event_service_proto_rawDescGZIP(), []int{7}
}

func (x *GetEventRequest) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

type GetEventsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Since         *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=since,proto3" json:"since,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetEventsRequest) Reset() {
	*x = GetEventsRequest{}
	mi := &file_event_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventsRequest) ProtoMessage() {}

func (x *GetEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_event_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventsRequest.ProtoReflect.Descriptor instead.
func (*GetEventsRequest) Descriptor() ([]byte, []int) {
	return file_event_service_proto_rawDescGZIP(), []int{8}
}

func (x *GetEventsRequest) GetSince() *timestamppb.Timestamp {
	if x != nil {
		return x.Since
	}
	return nil
}

type GetEventResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Event         *Event                 `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetEventResponse) Reset() {
	*x = GetEventResponse{}
	mi := &file_event_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventResponse) ProtoMessage() {}

func (x *GetEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_event_service_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventResponse.ProtoReflect.Descriptor instead.
func (*GetEventResponse) Descriptor() ([]byte, []int) {
	return file_event_service_proto_rawDescGZIP(), []int{9}
}

func (x *GetEventResponse) GetEvent() *Event {
	if x != nil {
		return x.Event
	}
	return nil
}

type GetEventsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Events        []*Event               `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetEventsResponse) Reset() {
	*x = GetEventsResponse{}
	mi := &file_event_service_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetEventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventsResponse) ProtoMessage() {}

func (x *GetEventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_event_service_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventsResponse.ProtoReflect.Descriptor instead.
func (*GetEventsResponse) Descriptor() ([]byte, []int) {
	return file_event_service_proto_rawDescGZIP(), []int{10}
}

func (x *GetEventsResponse) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

var File_event_service_proto protoreflect.FileDescriptor

var file_event_service_proto_rawDesc = string([]byte{
	0x0a, 0x13, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xa3, 0x02, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x35, 0x0a,
	0x08, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x64, 0x61, 0x74,
	0x65, 0x45, 0x6e, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x47, 0x0a, 0x11, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x10, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x9d, 0x02, 0x0a, 0x0f, 0x41, 0x64, 0x64,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x35, 0x0a,
	0x08, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x64, 0x61, 0x74,
	0x65, 0x45, 0x6e, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x47, 0x0a, 0x11, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x10, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x2d, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x40, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a,
	0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x3a, 0x0a, 0x13, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x23, 0x0a, 0x01, 0x5f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x52, 0x00, 0x22, 0x2f, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x3a, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a,
	0x01, 0x5f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x52, 0x00, 0x22, 0x2c, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x22, 0x44, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x05, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x05, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x22, 0x3e, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52,
	0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x41, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x32, 0xe0, 0x04, 0x0a, 0x08, 0x43, 0x61,
	0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x12, 0x4d, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x1e, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x56, 0x0a,
	0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x22, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x1e, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x4f, 0x66, 0x44, 0x61, 0x79, 0x12, 0x1f, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x4f, 0x66, 0x57, 0x65, 0x65, 0x6b, 0x12, 0x1f,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47,
	0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x20, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x4f, 0x66, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x1f, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x1d, 0x5a, 0x1b,
	0x2e, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3b, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

var (
	file_event_service_proto_rawDescOnce sync.Once
	file_event_service_proto_rawDescData []byte
)

func file_event_service_proto_rawDescGZIP() []byte {
	file_event_service_proto_rawDescOnce.Do(func() {
		file_event_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_event_service_proto_rawDesc), len(file_event_service_proto_rawDesc)))
	})
	return file_event_service_proto_rawDescData
}

var file_event_service_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_event_service_proto_goTypes = []any{
	(*Event)(nil),                 // 0: event_service.Event
	(*AddEventRequest)(nil),       // 1: event_service.AddEventRequest
	(*AddEventResponse)(nil),      // 2: event_service.AddEventResponse
	(*UpdateEventRequest)(nil),    // 3: event_service.UpdateEventRequest
	(*UpdateEventResponse)(nil),   // 4: event_service.UpdateEventResponse
	(*DeleteEventRequest)(nil),    // 5: event_service.DeleteEventRequest
	(*DeleteEventResponse)(nil),   // 6: event_service.DeleteEventResponse
	(*GetEventRequest)(nil),       // 7: event_service.GetEventRequest
	(*GetEventsRequest)(nil),      // 8: event_service.GetEventsRequest
	(*GetEventResponse)(nil),      // 9: event_service.GetEventResponse
	(*GetEventsResponse)(nil),     // 10: event_service.GetEventsResponse
	(*timestamppb.Timestamp)(nil), // 11: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 12: google.protobuf.Empty
}
var file_event_service_proto_depIdxs = []int32{
	11, // 0: event_service.Event.date_start:type_name -> google.protobuf.Timestamp
	11, // 1: event_service.Event.date_end:type_name -> google.protobuf.Timestamp
	11, // 2: event_service.Event.date_notification:type_name -> google.protobuf.Timestamp
	11, // 3: event_service.AddEventRequest.date_start:type_name -> google.protobuf.Timestamp
	11, // 4: event_service.AddEventRequest.date_end:type_name -> google.protobuf.Timestamp
	11, // 5: event_service.AddEventRequest.date_notification:type_name -> google.protobuf.Timestamp
	0,  // 6: event_service.UpdateEventRequest.event:type_name -> event_service.Event
	12, // 7: event_service.UpdateEventResponse._:type_name -> google.protobuf.Empty
	12, // 8: event_service.DeleteEventResponse._:type_name -> google.protobuf.Empty
	11, // 9: event_service.GetEventsRequest.since:type_name -> google.protobuf.Timestamp
	0,  // 10: event_service.GetEventResponse.event:type_name -> event_service.Event
	0,  // 11: event_service.GetEventsResponse.events:type_name -> event_service.Event
	1,  // 12: event_service.Calendar.AddEvent:input_type -> event_service.AddEventRequest
	3,  // 13: event_service.Calendar.UpdateEvent:input_type -> event_service.UpdateEventRequest
	5,  // 14: event_service.Calendar.DeleteEvent:input_type -> event_service.DeleteEventRequest
	7,  // 15: event_service.Calendar.GetEvent:input_type -> event_service.GetEventRequest
	8,  // 16: event_service.Calendar.GetEventsOfDay:input_type -> event_service.GetEventsRequest
	8,  // 17: event_service.Calendar.GetEventsOfWeek:input_type -> event_service.GetEventsRequest
	8,  // 18: event_service.Calendar.GetEventsOfMonth:input_type -> event_service.GetEventsRequest
	2,  // 19: event_service.Calendar.AddEvent:output_type -> event_service.AddEventResponse
	4,  // 20: event_service.Calendar.UpdateEvent:output_type -> event_service.UpdateEventResponse
	6,  // 21: event_service.Calendar.DeleteEvent:output_type -> event_service.DeleteEventResponse
	9,  // 22: event_service.Calendar.GetEvent:output_type -> event_service.GetEventResponse
	10, // 23: event_service.Calendar.GetEventsOfDay:output_type -> event_service.GetEventsResponse
	10, // 24: event_service.Calendar.GetEventsOfWeek:output_type -> event_service.GetEventsResponse
	10, // 25: event_service.Calendar.GetEventsOfMonth:output_type -> event_service.GetEventsResponse
	19, // [19:26] is the sub-list for method output_type
	12, // [12:19] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_event_service_proto_init() }
func file_event_service_proto_init() {
	if File_event_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_event_service_proto_rawDesc), len(file_event_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_event_service_proto_goTypes,
		DependencyIndexes: file_event_service_proto_depIdxs,
		MessageInfos:      file_event_service_proto_msgTypes,
	}.Build()
	File_event_service_proto = out.File
	file_event_service_proto_goTypes = nil
	file_event_service_proto_depIdxs = nil
}
