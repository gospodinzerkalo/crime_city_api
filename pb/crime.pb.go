// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: pb/crime.proto

package pb

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

type GetCrimesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetCrimesRequest) Reset() {
	*x = GetCrimesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_crime_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCrimesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCrimesRequest) ProtoMessage() {}

func (x *GetCrimesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_crime_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCrimesRequest.ProtoReflect.Descriptor instead.
func (*GetCrimesRequest) Descriptor() ([]byte, []int) {
	return file_pb_crime_proto_rawDescGZIP(), []int{0}
}

type GetCrimesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Crimes []*Crime `protobuf:"bytes,1,rep,name=crimes,proto3" json:"crimes,omitempty"`
}

func (x *GetCrimesResponse) Reset() {
	*x = GetCrimesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_crime_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCrimesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCrimesResponse) ProtoMessage() {}

func (x *GetCrimesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_crime_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCrimesResponse.ProtoReflect.Descriptor instead.
func (*GetCrimesResponse) Descriptor() ([]byte, []int) {
	return file_pb_crime_proto_rawDescGZIP(), []int{1}
}

func (x *GetCrimesResponse) GetCrimes() []*Crime {
	if x != nil {
		return x.Crimes
	}
	return nil
}

type Crime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	LocationName string  `protobuf:"bytes,2,opt,name=location_name,json=locationName,proto3" json:"location_name,omitempty"`
	Longitude    float64 `protobuf:"fixed64,3,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Latitude     float64 `protobuf:"fixed64,4,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Description  string  `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Image        string  `protobuf:"bytes,6,opt,name=image,proto3" json:"image,omitempty"`
	Date         string  `protobuf:"bytes,7,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *Crime) Reset() {
	*x = Crime{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_crime_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Crime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Crime) ProtoMessage() {}

func (x *Crime) ProtoReflect() protoreflect.Message {
	mi := &file_pb_crime_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Crime.ProtoReflect.Descriptor instead.
func (*Crime) Descriptor() ([]byte, []int) {
	return file_pb_crime_proto_rawDescGZIP(), []int{2}
}

func (x *Crime) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Crime) GetLocationName() string {
	if x != nil {
		return x.LocationName
	}
	return ""
}

func (x *Crime) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *Crime) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Crime) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Crime) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Crime) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type GetCrimeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCrimeRequest) Reset() {
	*x = GetCrimeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_crime_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCrimeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCrimeRequest) ProtoMessage() {}

func (x *GetCrimeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_crime_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCrimeRequest.ProtoReflect.Descriptor instead.
func (*GetCrimeRequest) Descriptor() ([]byte, []int) {
	return file_pb_crime_proto_rawDescGZIP(), []int{3}
}

func (x *GetCrimeRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetCrimeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Crime *Crime `protobuf:"bytes,1,opt,name=crime,proto3" json:"crime,omitempty"`
}

func (x *GetCrimeResponse) Reset() {
	*x = GetCrimeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_crime_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCrimeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCrimeResponse) ProtoMessage() {}

func (x *GetCrimeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_crime_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCrimeResponse.ProtoReflect.Descriptor instead.
func (*GetCrimeResponse) Descriptor() ([]byte, []int) {
	return file_pb_crime_proto_rawDescGZIP(), []int{4}
}

func (x *GetCrimeResponse) GetCrime() *Crime {
	if x != nil {
		return x.Crime
	}
	return nil
}

type CreateHomeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Home *Home `protobuf:"bytes,1,opt,name=home,proto3" json:"home,omitempty"`
}

func (x *CreateHomeRequest) Reset() {
	*x = CreateHomeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_crime_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateHomeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHomeRequest) ProtoMessage() {}

func (x *CreateHomeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_crime_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHomeRequest.ProtoReflect.Descriptor instead.
func (*CreateHomeRequest) Descriptor() ([]byte, []int) {
	return file_pb_crime_proto_rawDescGZIP(), []int{5}
}

func (x *CreateHomeRequest) GetHome() *Home {
	if x != nil {
		return x.Home
	}
	return nil
}

type CreateHomeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Home *Home `protobuf:"bytes,1,opt,name=home,proto3" json:"home,omitempty"`
}

func (x *CreateHomeResponse) Reset() {
	*x = CreateHomeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_crime_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateHomeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHomeResponse) ProtoMessage() {}

func (x *CreateHomeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_crime_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHomeResponse.ProtoReflect.Descriptor instead.
func (*CreateHomeResponse) Descriptor() ([]byte, []int) {
	return file_pb_crime_proto_rawDescGZIP(), []int{6}
}

func (x *CreateHomeResponse) GetHome() *Home {
	if x != nil {
		return x.Home
	}
	return nil
}

type Home struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName string  `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string  `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	UserName  string  `protobuf:"bytes,4,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	Longitude float64 `protobuf:"fixed64,5,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Latitude  float64 `protobuf:"fixed64,6,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Image     string  `protobuf:"bytes,7,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *Home) Reset() {
	*x = Home{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_crime_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Home) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Home) ProtoMessage() {}

func (x *Home) ProtoReflect() protoreflect.Message {
	mi := &file_pb_crime_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Home.ProtoReflect.Descriptor instead.
func (*Home) Descriptor() ([]byte, []int) {
	return file_pb_crime_proto_rawDescGZIP(), []int{7}
}

func (x *Home) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Home) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Home) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Home) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *Home) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *Home) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Home) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

type GetHomeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetHomeRequest) Reset() {
	*x = GetHomeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_crime_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHomeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHomeRequest) ProtoMessage() {}

func (x *GetHomeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_crime_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHomeRequest.ProtoReflect.Descriptor instead.
func (*GetHomeRequest) Descriptor() ([]byte, []int) {
	return file_pb_crime_proto_rawDescGZIP(), []int{8}
}

func (x *GetHomeRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetHomeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Home *Home `protobuf:"bytes,1,opt,name=home,proto3" json:"home,omitempty"`
}

func (x *GetHomeResponse) Reset() {
	*x = GetHomeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_crime_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHomeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHomeResponse) ProtoMessage() {}

func (x *GetHomeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_crime_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHomeResponse.ProtoReflect.Descriptor instead.
func (*GetHomeResponse) Descriptor() ([]byte, []int) {
	return file_pb_crime_proto_rawDescGZIP(), []int{9}
}

func (x *GetHomeResponse) GetHome() *Home {
	if x != nil {
		return x.Home
	}
	return nil
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Code    int32  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_crime_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_pb_crime_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_pb_crime_proto_rawDescGZIP(), []int{10}
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Error) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type DeleteHomeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteHomeRequest) Reset() {
	*x = DeleteHomeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_crime_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteHomeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteHomeRequest) ProtoMessage() {}

func (x *DeleteHomeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_crime_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteHomeRequest.ProtoReflect.Descriptor instead.
func (*DeleteHomeRequest) Descriptor() ([]byte, []int) {
	return file_pb_crime_proto_rawDescGZIP(), []int{11}
}

func (x *DeleteHomeRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteHomeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteHomeResponse) Reset() {
	*x = DeleteHomeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_crime_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteHomeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteHomeResponse) ProtoMessage() {}

func (x *DeleteHomeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_crime_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteHomeResponse.ProtoReflect.Descriptor instead.
func (*DeleteHomeResponse) Descriptor() ([]byte, []int) {
	return file_pb_crime_proto_rawDescGZIP(), []int{12}
}

type CheckHomeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CheckHomeRequest) Reset() {
	*x = CheckHomeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_crime_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckHomeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckHomeRequest) ProtoMessage() {}

func (x *CheckHomeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_crime_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckHomeRequest.ProtoReflect.Descriptor instead.
func (*CheckHomeRequest) Descriptor() ([]byte, []int) {
	return file_pb_crime_proto_rawDescGZIP(), []int{13}
}

func (x *CheckHomeRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CheckHomeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LocationName string  `protobuf:"bytes,1,opt,name=location_name,json=locationName,proto3" json:"location_name,omitempty"`
	Description  string  `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Url          string  `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	Distance     float64 `protobuf:"fixed64,4,opt,name=distance,proto3" json:"distance,omitempty"`
}

func (x *CheckHomeResponse) Reset() {
	*x = CheckHomeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_crime_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckHomeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckHomeResponse) ProtoMessage() {}

func (x *CheckHomeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_crime_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckHomeResponse.ProtoReflect.Descriptor instead.
func (*CheckHomeResponse) Descriptor() ([]byte, []int) {
	return file_pb_crime_proto_rawDescGZIP(), []int{14}
}

func (x *CheckHomeResponse) GetLocationName() string {
	if x != nil {
		return x.LocationName
	}
	return ""
}

func (x *CheckHomeResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CheckHomeResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *CheckHomeResponse) GetDistance() float64 {
	if x != nil {
		return x.Distance
	}
	return 0
}

var File_pb_crime_proto protoreflect.FileDescriptor

var file_pb_crime_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x62, 0x2f, 0x63, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x12, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x72, 0x69, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x33, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x72, 0x69, 0x6d, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x06, 0x63, 0x72, 0x69,
	0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x43, 0x72, 0x69, 0x6d,
	0x65, 0x52, 0x06, 0x63, 0x72, 0x69, 0x6d, 0x65, 0x73, 0x22, 0xc2, 0x01, 0x0a, 0x05, 0x43, 0x72,
	0x69, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e,
	0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0x21,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x72, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x30, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x72, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x05, 0x63, 0x72, 0x69, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x43, 0x72, 0x69, 0x6d, 0x65, 0x52, 0x05, 0x63, 0x72,
	0x69, 0x6d, 0x65, 0x22, 0x2e, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x6d,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x04, 0x68, 0x6f, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x52, 0x04, 0x68,
	0x6f, 0x6d, 0x65, 0x22, 0x2f, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x6d,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x04, 0x68, 0x6f, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x52, 0x04,
	0x68, 0x6f, 0x6d, 0x65, 0x22, 0xbf, 0x01, 0x0a, 0x04, 0x48, 0x6f, 0x6d, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x6d,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2c, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x48,
	0x6f, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x04, 0x68,
	0x6f, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x48, 0x6f, 0x6d, 0x65,
	0x52, 0x04, 0x68, 0x6f, 0x6d, 0x65, 0x22, 0x35, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x23, 0x0a,
	0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x6f, 0x6d, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x0a, 0x10, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x48, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x88, 0x01, 0x0a,
	0x11, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x48, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x64,
	0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x64,
	0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x32, 0xcf, 0x02, 0x0a, 0x0c, 0x43, 0x72, 0x69, 0x6d,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x43,
	0x72, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x11, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x72, 0x69, 0x6d, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x72,
	0x69, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x31,
	0x0a, 0x08, 0x47, 0x65, 0x74, 0x43, 0x72, 0x69, 0x6d, 0x65, 0x12, 0x10, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x72, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x72, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x37, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x6d, 0x65, 0x12,
	0x12, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x6d, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x07, 0x47, 0x65,
	0x74, 0x48, 0x6f, 0x6d, 0x65, 0x12, 0x0f, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x6d, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x6d, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0a, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x48, 0x6f, 0x6d, 0x65, 0x12, 0x12, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x48, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x09, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x48, 0x6f, 0x6d, 0x65,
	0x12, 0x11, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x48, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x48, 0x6f, 0x6d, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x73, 0x70, 0x6f, 0x64, 0x69, 0x6e,
	0x7a, 0x65, 0x72, 0x6b, 0x61, 0x6c, 0x6f, 0x2f, 0x63, 0x72, 0x69, 0x6d, 0x65, 0x5f, 0x63, 0x69,
	0x74, 0x79, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_pb_crime_proto_rawDescOnce sync.Once
	file_pb_crime_proto_rawDescData = file_pb_crime_proto_rawDesc
)

func file_pb_crime_proto_rawDescGZIP() []byte {
	file_pb_crime_proto_rawDescOnce.Do(func() {
		file_pb_crime_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_crime_proto_rawDescData)
	})
	return file_pb_crime_proto_rawDescData
}

var file_pb_crime_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_pb_crime_proto_goTypes = []interface{}{
	(*GetCrimesRequest)(nil),   // 0: GetCrimesRequest
	(*GetCrimesResponse)(nil),  // 1: GetCrimesResponse
	(*Crime)(nil),              // 2: Crime
	(*GetCrimeRequest)(nil),    // 3: GetCrimeRequest
	(*GetCrimeResponse)(nil),   // 4: GetCrimeResponse
	(*CreateHomeRequest)(nil),  // 5: CreateHomeRequest
	(*CreateHomeResponse)(nil), // 6: CreateHomeResponse
	(*Home)(nil),               // 7: Home
	(*GetHomeRequest)(nil),     // 8: GetHomeRequest
	(*GetHomeResponse)(nil),    // 9: GetHomeResponse
	(*Error)(nil),              // 10: Error
	(*DeleteHomeRequest)(nil),  // 11: DeleteHomeRequest
	(*DeleteHomeResponse)(nil), // 12: DeleteHomeResponse
	(*CheckHomeRequest)(nil),   // 13: CheckHomeRequest
	(*CheckHomeResponse)(nil),  // 14: CheckHomeResponse
}
var file_pb_crime_proto_depIdxs = []int32{
	2,  // 0: GetCrimesResponse.crimes:type_name -> Crime
	2,  // 1: GetCrimeResponse.crime:type_name -> Crime
	7,  // 2: CreateHomeRequest.home:type_name -> Home
	7,  // 3: CreateHomeResponse.home:type_name -> Home
	7,  // 4: GetHomeResponse.home:type_name -> Home
	0,  // 5: CrimeService.GetCrimes:input_type -> GetCrimesRequest
	3,  // 6: CrimeService.GetCrime:input_type -> GetCrimeRequest
	5,  // 7: CrimeService.CreateHome:input_type -> CreateHomeRequest
	8,  // 8: CrimeService.GetHome:input_type -> GetHomeRequest
	11, // 9: CrimeService.DeleteHome:input_type -> DeleteHomeRequest
	13, // 10: CrimeService.CheckHome:input_type -> CheckHomeRequest
	1,  // 11: CrimeService.GetCrimes:output_type -> GetCrimesResponse
	4,  // 12: CrimeService.GetCrime:output_type -> GetCrimeResponse
	6,  // 13: CrimeService.CreateHome:output_type -> CreateHomeResponse
	9,  // 14: CrimeService.GetHome:output_type -> GetHomeResponse
	12, // 15: CrimeService.DeleteHome:output_type -> DeleteHomeResponse
	14, // 16: CrimeService.CheckHome:output_type -> CheckHomeResponse
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_pb_crime_proto_init() }
func file_pb_crime_proto_init() {
	if File_pb_crime_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_crime_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCrimesRequest); i {
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
		file_pb_crime_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCrimesResponse); i {
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
		file_pb_crime_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Crime); i {
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
		file_pb_crime_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCrimeRequest); i {
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
		file_pb_crime_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCrimeResponse); i {
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
		file_pb_crime_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateHomeRequest); i {
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
		file_pb_crime_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateHomeResponse); i {
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
		file_pb_crime_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Home); i {
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
		file_pb_crime_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHomeRequest); i {
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
		file_pb_crime_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHomeResponse); i {
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
		file_pb_crime_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
		file_pb_crime_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteHomeRequest); i {
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
		file_pb_crime_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteHomeResponse); i {
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
		file_pb_crime_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckHomeRequest); i {
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
		file_pb_crime_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckHomeResponse); i {
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
			RawDescriptor: file_pb_crime_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_crime_proto_goTypes,
		DependencyIndexes: file_pb_crime_proto_depIdxs,
		MessageInfos:      file_pb_crime_proto_msgTypes,
	}.Build()
	File_pb_crime_proto = out.File
	file_pb_crime_proto_rawDesc = nil
	file_pb_crime_proto_goTypes = nil
	file_pb_crime_proto_depIdxs = nil
}
