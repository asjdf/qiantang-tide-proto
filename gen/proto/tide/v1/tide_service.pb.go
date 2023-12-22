// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: tide/v1/tide_service.proto

package tideV1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetDayPredictTideRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The date of the predict data.
	Date *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *GetDayPredictTideRequest) Reset() {
	*x = GetDayPredictTideRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tide_v1_tide_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDayPredictTideRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDayPredictTideRequest) ProtoMessage() {}

func (x *GetDayPredictTideRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tide_v1_tide_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDayPredictTideRequest.ProtoReflect.Descriptor instead.
func (*GetDayPredictTideRequest) Descriptor() ([]byte, []int) {
	return file_tide_v1_tide_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetDayPredictTideRequest) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

type GetDayPredictTideResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The predicted tide of the given day.
	TideList []*Tide `protobuf:"bytes,1,rep,name=tideList,proto3" json:"tideList,omitempty"`
}

func (x *GetDayPredictTideResponse) Reset() {
	*x = GetDayPredictTideResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tide_v1_tide_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDayPredictTideResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDayPredictTideResponse) ProtoMessage() {}

func (x *GetDayPredictTideResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tide_v1_tide_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDayPredictTideResponse.ProtoReflect.Descriptor instead.
func (*GetDayPredictTideResponse) Descriptor() ([]byte, []int) {
	return file_tide_v1_tide_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetDayPredictTideResponse) GetTideList() []*Tide {
	if x != nil {
		return x.TideList
	}
	return nil
}

type GetLocationPredictTideRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The locationId of the tide.
	LocationId string `protobuf:"bytes,1,opt,name=locationId,proto3" json:"locationId,omitempty"`
	// The date of the tide.
	Date *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *GetLocationPredictTideRequest) Reset() {
	*x = GetLocationPredictTideRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tide_v1_tide_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLocationPredictTideRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLocationPredictTideRequest) ProtoMessage() {}

func (x *GetLocationPredictTideRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tide_v1_tide_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLocationPredictTideRequest.ProtoReflect.Descriptor instead.
func (*GetLocationPredictTideRequest) Descriptor() ([]byte, []int) {
	return file_tide_v1_tide_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetLocationPredictTideRequest) GetLocationId() string {
	if x != nil {
		return x.LocationId
	}
	return ""
}

func (x *GetLocationPredictTideRequest) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

type GetLocationPredictTideResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The predicted tide of the given location.
	TideList []*Tide `protobuf:"bytes,1,rep,name=tideList,proto3" json:"tideList,omitempty"`
}

func (x *GetLocationPredictTideResponse) Reset() {
	*x = GetLocationPredictTideResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tide_v1_tide_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLocationPredictTideResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLocationPredictTideResponse) ProtoMessage() {}

func (x *GetLocationPredictTideResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tide_v1_tide_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLocationPredictTideResponse.ProtoReflect.Descriptor instead.
func (*GetLocationPredictTideResponse) Descriptor() ([]byte, []int) {
	return file_tide_v1_tide_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetLocationPredictTideResponse) GetTideList() []*Tide {
	if x != nil {
		return x.TideList
	}
	return nil
}

type GetLocationListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetLocationListRequest) Reset() {
	*x = GetLocationListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tide_v1_tide_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLocationListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLocationListRequest) ProtoMessage() {}

func (x *GetLocationListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tide_v1_tide_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLocationListRequest.ProtoReflect.Descriptor instead.
func (*GetLocationListRequest) Descriptor() ([]byte, []int) {
	return file_tide_v1_tide_service_proto_rawDescGZIP(), []int{4}
}

type GetLocationListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The location list.
	LocationList []*Location `protobuf:"bytes,1,rep,name=locationList,proto3" json:"locationList,omitempty"`
}

func (x *GetLocationListResponse) Reset() {
	*x = GetLocationListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tide_v1_tide_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLocationListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLocationListResponse) ProtoMessage() {}

func (x *GetLocationListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tide_v1_tide_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLocationListResponse.ProtoReflect.Descriptor instead.
func (*GetLocationListResponse) Descriptor() ([]byte, []int) {
	return file_tide_v1_tide_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetLocationListResponse) GetLocationList() []*Location {
	if x != nil {
		return x.LocationList
	}
	return nil
}

var File_tide_v1_tide_service_proto protoreflect.FileDescriptor

var file_tide_v1_tide_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x74, 0x69, 0x64, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x69, 0x64, 0x65, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x74, 0x69,
	0x64, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x74, 0x69, 0x64, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x69,
	0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4a, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x44,
	0x61, 0x79, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x54, 0x69, 0x64, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x22, 0x46, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x44, 0x61, 0x79, 0x50, 0x72,
	0x65, 0x64, 0x69, 0x63, 0x74, 0x54, 0x69, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x29, 0x0a, 0x08, 0x74, 0x69, 0x64, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x74, 0x69, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69,
	0x64, 0x65, 0x52, 0x08, 0x74, 0x69, 0x64, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x6f, 0x0a, 0x1d,
	0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x65, 0x64, 0x69,
	0x63, 0x74, 0x54, 0x69, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a,
	0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x2e, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0x4b, 0x0a,
	0x1e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x65, 0x64,
	0x69, 0x63, 0x74, 0x54, 0x69, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x29, 0x0a, 0x08, 0x74, 0x69, 0x64, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x74, 0x69, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x64, 0x65,
	0x52, 0x08, 0x74, 0x69, 0x64, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x18, 0x0a, 0x16, 0x47, 0x65,
	0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x50, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x35, 0x0a, 0x0c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x69, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x32, 0x9b, 0x03, 0x0a, 0x0b, 0x54, 0x69, 0x64, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7d, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x44, 0x61, 0x79,
	0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x54, 0x69, 0x64, 0x65, 0x12, 0x21, 0x2e, 0x74, 0x69,
	0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x79, 0x50, 0x72, 0x65, 0x64,
	0x69, 0x63, 0x74, 0x54, 0x69, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22,
	0x2e, 0x74, 0x69, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x79, 0x50,
	0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x54, 0x69, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x67, 0x61, 0x70,
	0x69, 0x2f, 0x74, 0x69, 0x64, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63,
	0x74, 0x2f, 0x64, 0x61, 0x79, 0x12, 0x91, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x54, 0x69, 0x64, 0x65,
	0x12, 0x26, 0x2e, 0x74, 0x69, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x54, 0x69, 0x64,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x74, 0x69, 0x64, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72,
	0x65, 0x64, 0x69, 0x63, 0x74, 0x54, 0x69, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x12, 0x1e, 0x2f, 0x67, 0x61, 0x70, 0x69,
	0x2f, 0x74, 0x69, 0x64, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74,
	0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x79, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1f, 0x2e, 0x74,
	0x69, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e,
	0x74, 0x69, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12, 0x1b, 0x2f, 0x67, 0x61, 0x70, 0x69, 0x2f, 0x74,
	0x69, 0x64, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x6c, 0x69, 0x73, 0x74, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x61, 0x73, 0x6a, 0x64, 0x66, 0x2f, 0x71, 0x69, 0x61, 0x6e, 0x74, 0x61, 0x6e,
	0x67, 0x2d, 0x74, 0x69, 0x64, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x69, 0x64, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x74,
	0x69, 0x64, 0x65, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tide_v1_tide_service_proto_rawDescOnce sync.Once
	file_tide_v1_tide_service_proto_rawDescData = file_tide_v1_tide_service_proto_rawDesc
)

func file_tide_v1_tide_service_proto_rawDescGZIP() []byte {
	file_tide_v1_tide_service_proto_rawDescOnce.Do(func() {
		file_tide_v1_tide_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_tide_v1_tide_service_proto_rawDescData)
	})
	return file_tide_v1_tide_service_proto_rawDescData
}

var file_tide_v1_tide_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_tide_v1_tide_service_proto_goTypes = []interface{}{
	(*GetDayPredictTideRequest)(nil),       // 0: tide.v1.GetDayPredictTideRequest
	(*GetDayPredictTideResponse)(nil),      // 1: tide.v1.GetDayPredictTideResponse
	(*GetLocationPredictTideRequest)(nil),  // 2: tide.v1.GetLocationPredictTideRequest
	(*GetLocationPredictTideResponse)(nil), // 3: tide.v1.GetLocationPredictTideResponse
	(*GetLocationListRequest)(nil),         // 4: tide.v1.GetLocationListRequest
	(*GetLocationListResponse)(nil),        // 5: tide.v1.GetLocationListResponse
	(*timestamppb.Timestamp)(nil),          // 6: google.protobuf.Timestamp
	(*Tide)(nil),                           // 7: tide.v1.Tide
	(*Location)(nil),                       // 8: tide.v1.Location
}
var file_tide_v1_tide_service_proto_depIdxs = []int32{
	6, // 0: tide.v1.GetDayPredictTideRequest.date:type_name -> google.protobuf.Timestamp
	7, // 1: tide.v1.GetDayPredictTideResponse.tideList:type_name -> tide.v1.Tide
	6, // 2: tide.v1.GetLocationPredictTideRequest.date:type_name -> google.protobuf.Timestamp
	7, // 3: tide.v1.GetLocationPredictTideResponse.tideList:type_name -> tide.v1.Tide
	8, // 4: tide.v1.GetLocationListResponse.locationList:type_name -> tide.v1.Location
	0, // 5: tide.v1.TideService.GetDayPredictTide:input_type -> tide.v1.GetDayPredictTideRequest
	2, // 6: tide.v1.TideService.GetLocationPredictTide:input_type -> tide.v1.GetLocationPredictTideRequest
	4, // 7: tide.v1.TideService.GetLocationList:input_type -> tide.v1.GetLocationListRequest
	1, // 8: tide.v1.TideService.GetDayPredictTide:output_type -> tide.v1.GetDayPredictTideResponse
	3, // 9: tide.v1.TideService.GetLocationPredictTide:output_type -> tide.v1.GetLocationPredictTideResponse
	5, // 10: tide.v1.TideService.GetLocationList:output_type -> tide.v1.GetLocationListResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_tide_v1_tide_service_proto_init() }
func file_tide_v1_tide_service_proto_init() {
	if File_tide_v1_tide_service_proto != nil {
		return
	}
	file_tide_v1_tide_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tide_v1_tide_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDayPredictTideRequest); i {
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
		file_tide_v1_tide_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDayPredictTideResponse); i {
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
		file_tide_v1_tide_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLocationPredictTideRequest); i {
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
		file_tide_v1_tide_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLocationPredictTideResponse); i {
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
		file_tide_v1_tide_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLocationListRequest); i {
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
		file_tide_v1_tide_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLocationListResponse); i {
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
			RawDescriptor: file_tide_v1_tide_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tide_v1_tide_service_proto_goTypes,
		DependencyIndexes: file_tide_v1_tide_service_proto_depIdxs,
		MessageInfos:      file_tide_v1_tide_service_proto_msgTypes,
	}.Build()
	File_tide_v1_tide_service_proto = out.File
	file_tide_v1_tide_service_proto_rawDesc = nil
	file_tide_v1_tide_service_proto_goTypes = nil
	file_tide_v1_tide_service_proto_depIdxs = nil
}
