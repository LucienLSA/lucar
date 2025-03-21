// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.15.5
// source: rental.proto

package rentalpb

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

type TripStatus int32

const (
	TripStatus_TS_NOT_SPECIFIED TripStatus = 0
	TripStatus_IN_PROGRESS      TripStatus = 1
	TripStatus_FINISHED         TripStatus = 2
)

// Enum value maps for TripStatus.
var (
	TripStatus_name = map[int32]string{
		0: "TS_NOT_SPECIFIED",
		1: "IN_PROGRESS",
		2: "FINISHED",
	}
	TripStatus_value = map[string]int32{
		"TS_NOT_SPECIFIED": 0,
		"IN_PROGRESS":      1,
		"FINISHED":         2,
	}
)

func (x TripStatus) Enum() *TripStatus {
	p := new(TripStatus)
	*p = x
	return p
}

func (x TripStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TripStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_rental_proto_enumTypes[0].Descriptor()
}

func (TripStatus) Type() protoreflect.EnumType {
	return &file_rental_proto_enumTypes[0]
}

func (x TripStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TripStatus.Descriptor instead.
func (TripStatus) EnumDescriptor() ([]byte, []int) {
	return file_rental_proto_rawDescGZIP(), []int{0}
}

type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude  float64 `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	mi := &file_rental_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_rental_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_rental_proto_rawDescGZIP(), []int{0}
}

func (x *Location) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Location) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

type LocationStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location *Location `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	FeeCent  int32     `protobuf:"varint,2,opt,name=fee_cent,json=feeCent,proto3" json:"fee_cent,omitempty"`
	KmDriven float64   `protobuf:"fixed64,3,opt,name=km_driven,json=kmDriven,proto3" json:"km_driven,omitempty"`
	PoiName  string    `protobuf:"bytes,4,opt,name=poi_name,json=poiName,proto3" json:"poi_name,omitempty"`
}

func (x *LocationStatus) Reset() {
	*x = LocationStatus{}
	mi := &file_rental_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LocationStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocationStatus) ProtoMessage() {}

func (x *LocationStatus) ProtoReflect() protoreflect.Message {
	mi := &file_rental_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocationStatus.ProtoReflect.Descriptor instead.
func (*LocationStatus) Descriptor() ([]byte, []int) {
	return file_rental_proto_rawDescGZIP(), []int{1}
}

func (x *LocationStatus) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *LocationStatus) GetFeeCent() int32 {
	if x != nil {
		return x.FeeCent
	}
	return 0
}

func (x *LocationStatus) GetKmDriven() float64 {
	if x != nil {
		return x.KmDriven
	}
	return 0
}

func (x *LocationStatus) GetPoiName() string {
	if x != nil {
		return x.PoiName
	}
	return ""
}

type TripEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Trip *Trip  `protobuf:"bytes,2,opt,name=trip,proto3" json:"trip,omitempty"`
}

func (x *TripEntity) Reset() {
	*x = TripEntity{}
	mi := &file_rental_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TripEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TripEntity) ProtoMessage() {}

func (x *TripEntity) ProtoReflect() protoreflect.Message {
	mi := &file_rental_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TripEntity.ProtoReflect.Descriptor instead.
func (*TripEntity) Descriptor() ([]byte, []int) {
	return file_rental_proto_rawDescGZIP(), []int{2}
}

func (x *TripEntity) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TripEntity) GetTrip() *Trip {
	if x != nil {
		return x.Trip
	}
	return nil
}

type Trip struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId string          `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	CarId     string          `protobuf:"bytes,2,opt,name=car_id,json=carId,proto3" json:"car_id,omitempty"`
	Start     *LocationStatus `protobuf:"bytes,3,opt,name=start,proto3" json:"start,omitempty"`
	Current   *LocationStatus `protobuf:"bytes,4,opt,name=current,proto3" json:"current,omitempty"`
	End       *LocationStatus `protobuf:"bytes,5,opt,name=end,proto3" json:"end,omitempty"`
	Status    TripStatus      `protobuf:"varint,6,opt,name=status,proto3,enum=rental.v1.TripStatus" json:"status,omitempty"`
}

func (x *Trip) Reset() {
	*x = Trip{}
	mi := &file_rental_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Trip) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trip) ProtoMessage() {}

func (x *Trip) ProtoReflect() protoreflect.Message {
	mi := &file_rental_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trip.ProtoReflect.Descriptor instead.
func (*Trip) Descriptor() ([]byte, []int) {
	return file_rental_proto_rawDescGZIP(), []int{3}
}

func (x *Trip) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *Trip) GetCarId() string {
	if x != nil {
		return x.CarId
	}
	return ""
}

func (x *Trip) GetStart() *LocationStatus {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *Trip) GetCurrent() *LocationStatus {
	if x != nil {
		return x.Current
	}
	return nil
}

func (x *Trip) GetEnd() *LocationStatus {
	if x != nil {
		return x.End
	}
	return nil
}

func (x *Trip) GetStatus() TripStatus {
	if x != nil {
		return x.Status
	}
	return TripStatus_TS_NOT_SPECIFIED
}

type CreateTripRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start *Location `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	CarId string    `protobuf:"bytes,2,opt,name=car_id,json=carId,proto3" json:"car_id,omitempty"`
}

func (x *CreateTripRequest) Reset() {
	*x = CreateTripRequest{}
	mi := &file_rental_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTripRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTripRequest) ProtoMessage() {}

func (x *CreateTripRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rental_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTripRequest.ProtoReflect.Descriptor instead.
func (*CreateTripRequest) Descriptor() ([]byte, []int) {
	return file_rental_proto_rawDescGZIP(), []int{4}
}

func (x *CreateTripRequest) GetStart() *Location {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *CreateTripRequest) GetCarId() string {
	if x != nil {
		return x.CarId
	}
	return ""
}

type GetTripRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetTripRequest) Reset() {
	*x = GetTripRequest{}
	mi := &file_rental_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTripRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTripRequest) ProtoMessage() {}

func (x *GetTripRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rental_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTripRequest.ProtoReflect.Descriptor instead.
func (*GetTripRequest) Descriptor() ([]byte, []int) {
	return file_rental_proto_rawDescGZIP(), []int{5}
}

func (x *GetTripRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetTripsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status TripStatus `protobuf:"varint,1,opt,name=status,proto3,enum=rental.v1.TripStatus" json:"status,omitempty"`
}

func (x *GetTripsRequest) Reset() {
	*x = GetTripsRequest{}
	mi := &file_rental_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTripsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTripsRequest) ProtoMessage() {}

func (x *GetTripsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rental_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTripsRequest.ProtoReflect.Descriptor instead.
func (*GetTripsRequest) Descriptor() ([]byte, []int) {
	return file_rental_proto_rawDescGZIP(), []int{6}
}

func (x *GetTripsRequest) GetStatus() TripStatus {
	if x != nil {
		return x.Status
	}
	return TripStatus_TS_NOT_SPECIFIED
}

type GetTripResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Trips []*TripEntity `protobuf:"bytes,1,rep,name=trips,proto3" json:"trips,omitempty"`
}

func (x *GetTripResponse) Reset() {
	*x = GetTripResponse{}
	mi := &file_rental_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTripResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTripResponse) ProtoMessage() {}

func (x *GetTripResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rental_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTripResponse.ProtoReflect.Descriptor instead.
func (*GetTripResponse) Descriptor() ([]byte, []int) {
	return file_rental_proto_rawDescGZIP(), []int{7}
}

func (x *GetTripResponse) GetTrips() []*TripEntity {
	if x != nil {
		return x.Trips
	}
	return nil
}

type UpdateTripRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Current *Location `protobuf:"bytes,2,opt,name=current,proto3" json:"current,omitempty"`
	EndTrip bool      `protobuf:"varint,3,opt,name=end_trip,json=endTrip,proto3" json:"end_trip,omitempty"`
}

func (x *UpdateTripRequest) Reset() {
	*x = UpdateTripRequest{}
	mi := &file_rental_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateTripRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTripRequest) ProtoMessage() {}

func (x *UpdateTripRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rental_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTripRequest.ProtoReflect.Descriptor instead.
func (*UpdateTripRequest) Descriptor() ([]byte, []int) {
	return file_rental_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateTripRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateTripRequest) GetCurrent() *Location {
	if x != nil {
		return x.Current
	}
	return nil
}

func (x *UpdateTripRequest) GetEndTrip() bool {
	if x != nil {
		return x.EndTrip
	}
	return false
}

var File_rental_proto protoreflect.FileDescriptor

var file_rental_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x72, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x22, 0x44, 0x0a, 0x08, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x22,
	0x94, 0x01, 0x0a, 0x0e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x2f, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x65, 0x65, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x66, 0x65, 0x65, 0x43, 0x65, 0x6e, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x6b, 0x6d, 0x5f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x08, 0x6b, 0x6d, 0x44, 0x72, 0x69, 0x76, 0x65, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x70,
	0x6f, 0x69, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70,
	0x6f, 0x69, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x41, 0x0a, 0x0a, 0x54, 0x72, 0x69, 0x70, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x04, 0x74, 0x72, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x72, 0x69, 0x70, 0x52, 0x04, 0x74, 0x72, 0x69, 0x70, 0x22, 0xfe, 0x01, 0x0a, 0x04, 0x54, 0x72,
	0x69, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x15, 0x0a, 0x06, 0x63, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x63, 0x61, 0x72, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x6c,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x33, 0x0a, 0x07, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x72, 0x65, 0x6e,
	0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x2b,
	0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x72, 0x65,
	0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x2d, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x72, 0x65,
	0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x69, 0x70, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x55, 0x0a, 0x11, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x29, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x63, 0x61,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x61, 0x72, 0x49,
	0x64, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x40, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54, 0x72, 0x69, 0x70, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x72, 0x69, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x3e, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54, 0x72, 0x69, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x05, 0x74, 0x72, 0x69, 0x70,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x6c,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x69, 0x70, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x05,
	0x74, 0x72, 0x69, 0x70, 0x73, 0x22, 0x6d, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x72, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2d, 0x0a, 0x07, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x72, 0x65,
	0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64,
	0x5f, 0x74, 0x72, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x64,
	0x54, 0x72, 0x69, 0x70, 0x2a, 0x41, 0x0a, 0x0a, 0x54, 0x72, 0x69, 0x70, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x4e, 0x5f, 0x50,
	0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x49, 0x4e,
	0x49, 0x53, 0x48, 0x45, 0x44, 0x10, 0x02, 0x32, 0x88, 0x02, 0x0a, 0x0b, 0x54, 0x72, 0x69, 0x70,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x72, 0x69, 0x70, 0x12, 0x1c, 0x2e, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x72, 0x69, 0x70, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x35, 0x0a, 0x07, 0x47, 0x65,
	0x74, 0x54, 0x72, 0x69, 0x70, 0x12, 0x19, 0x2e, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0f, 0x2e, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x69,
	0x70, 0x12, 0x42, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x54, 0x72, 0x69, 0x70, 0x73, 0x12, 0x1a, 0x2e,
	0x72, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x69,
	0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x72, 0x65, 0x6e, 0x74,
	0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x69, 0x70, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x72, 0x69, 0x70, 0x12, 0x1c, 0x2e, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x72, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0f, 0x2e, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72,
	0x69, 0x70, 0x42, 0x22, 0x5a, 0x20, 0x6c, 0x75, 0x63, 0x61, 0x72, 0x2f, 0x72, 0x65, 0x6e, 0x74,
	0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x72, 0x65,
	0x6e, 0x74, 0x61, 0x6c, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rental_proto_rawDescOnce sync.Once
	file_rental_proto_rawDescData = file_rental_proto_rawDesc
)

func file_rental_proto_rawDescGZIP() []byte {
	file_rental_proto_rawDescOnce.Do(func() {
		file_rental_proto_rawDescData = protoimpl.X.CompressGZIP(file_rental_proto_rawDescData)
	})
	return file_rental_proto_rawDescData
}

var file_rental_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_rental_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_rental_proto_goTypes = []any{
	(TripStatus)(0),           // 0: rental.v1.TripStatus
	(*Location)(nil),          // 1: rental.v1.Location
	(*LocationStatus)(nil),    // 2: rental.v1.LocationStatus
	(*TripEntity)(nil),        // 3: rental.v1.TripEntity
	(*Trip)(nil),              // 4: rental.v1.Trip
	(*CreateTripRequest)(nil), // 5: rental.v1.CreateTripRequest
	(*GetTripRequest)(nil),    // 6: rental.v1.GetTripRequest
	(*GetTripsRequest)(nil),   // 7: rental.v1.GetTripsRequest
	(*GetTripResponse)(nil),   // 8: rental.v1.GetTripResponse
	(*UpdateTripRequest)(nil), // 9: rental.v1.UpdateTripRequest
}
var file_rental_proto_depIdxs = []int32{
	1,  // 0: rental.v1.LocationStatus.location:type_name -> rental.v1.Location
	4,  // 1: rental.v1.TripEntity.trip:type_name -> rental.v1.Trip
	2,  // 2: rental.v1.Trip.start:type_name -> rental.v1.LocationStatus
	2,  // 3: rental.v1.Trip.current:type_name -> rental.v1.LocationStatus
	2,  // 4: rental.v1.Trip.end:type_name -> rental.v1.LocationStatus
	0,  // 5: rental.v1.Trip.status:type_name -> rental.v1.TripStatus
	1,  // 6: rental.v1.CreateTripRequest.start:type_name -> rental.v1.Location
	0,  // 7: rental.v1.GetTripsRequest.status:type_name -> rental.v1.TripStatus
	3,  // 8: rental.v1.GetTripResponse.trips:type_name -> rental.v1.TripEntity
	1,  // 9: rental.v1.UpdateTripRequest.current:type_name -> rental.v1.Location
	5,  // 10: rental.v1.TripService.CreateTrip:input_type -> rental.v1.CreateTripRequest
	6,  // 11: rental.v1.TripService.GetTrip:input_type -> rental.v1.GetTripRequest
	7,  // 12: rental.v1.TripService.GetTrips:input_type -> rental.v1.GetTripsRequest
	9,  // 13: rental.v1.TripService.UpdateTrip:input_type -> rental.v1.UpdateTripRequest
	3,  // 14: rental.v1.TripService.CreateTrip:output_type -> rental.v1.TripEntity
	4,  // 15: rental.v1.TripService.GetTrip:output_type -> rental.v1.Trip
	8,  // 16: rental.v1.TripService.GetTrips:output_type -> rental.v1.GetTripResponse
	4,  // 17: rental.v1.TripService.UpdateTrip:output_type -> rental.v1.Trip
	14, // [14:18] is the sub-list for method output_type
	10, // [10:14] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_rental_proto_init() }
func file_rental_proto_init() {
	if File_rental_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rental_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rental_proto_goTypes,
		DependencyIndexes: file_rental_proto_depIdxs,
		EnumInfos:         file_rental_proto_enumTypes,
		MessageInfos:      file_rental_proto_msgTypes,
	}.Build()
	File_rental_proto = out.File
	file_rental_proto_rawDesc = nil
	file_rental_proto_goTypes = nil
	file_rental_proto_depIdxs = nil
}
