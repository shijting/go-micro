// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: Courses.proto

package courses

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Timestamp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Timestamp) Reset() {
	*x = Timestamp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Courses_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Timestamp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Timestamp) ProtoMessage() {}

func (x *Timestamp) ProtoReflect() protoreflect.Message {
	mi := &file_Courses_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Timestamp.ProtoReflect.Descriptor instead.
func (*Timestamp) Descriptor() ([]byte, []int) {
	return file_Courses_proto_rawDescGZIP(), []int{0}
}

func (x *Timestamp) GetTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size int32 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Courses_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Courses_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_Courses_proto_rawDescGZIP(), []int{1}
}

func (x *ListRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

type CourseModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CourseId       int32   `protobuf:"varint,1,opt,name=course_id,json=courseId,proto3" json:"course_id,omitempty"`
	CourseName     string  `protobuf:"bytes,2,opt,name=course_name,json=courseName,proto3" json:"course_name,omitempty"`
	CourseDispName string  `protobuf:"bytes,3,opt,name=course_disp_name,json=courseDispName,proto3" json:"course_disp_name,omitempty"`
	CoursePrice    float32 `protobuf:"fixed32,4,opt,name=course_price,json=coursePrice,proto3" json:"course_price,omitempty"`
	CoursePrice2   float32 `protobuf:"fixed32,5,opt,name=course_price2,json=coursePrice2,proto3" json:"course_price2,omitempty"`
	// go get github.com/favadi/protoc-go-inject-tag
	// @inject_tag: gorm:"type:timestamp"
	Addtime *Timestamp `protobuf:"bytes,6,opt,name=addtime,proto3" json:"addtime,omitempty" gorm:"type:timestamp"`
}

func (x *CourseModel) Reset() {
	*x = CourseModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Courses_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CourseModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourseModel) ProtoMessage() {}

func (x *CourseModel) ProtoReflect() protoreflect.Message {
	mi := &file_Courses_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourseModel.ProtoReflect.Descriptor instead.
func (*CourseModel) Descriptor() ([]byte, []int) {
	return file_Courses_proto_rawDescGZIP(), []int{2}
}

func (x *CourseModel) GetCourseId() int32 {
	if x != nil {
		return x.CourseId
	}
	return 0
}

func (x *CourseModel) GetCourseName() string {
	if x != nil {
		return x.CourseName
	}
	return ""
}

func (x *CourseModel) GetCourseDispName() string {
	if x != nil {
		return x.CourseDispName
	}
	return ""
}

func (x *CourseModel) GetCoursePrice() float32 {
	if x != nil {
		return x.CoursePrice
	}
	return 0
}

func (x *CourseModel) GetCoursePrice2() float32 {
	if x != nil {
		return x.CoursePrice2
	}
	return 0
}

func (x *CourseModel) GetAddtime() *Timestamp {
	if x != nil {
		return x.Addtime
	}
	return nil
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result []*CourseModel `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Courses_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Courses_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_Courses_proto_rawDescGZIP(), []int{3}
}

func (x *ListResponse) GetResult() []*CourseModel {
	if x != nil {
		return x.Result
	}
	return nil
}

type DetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: uri:"course_id"
	CourseId int32 `protobuf:"varint,1,opt,name=course_id,json=courseId,proto3" json:"course_id,omitempty" uri:"course_id"`
	// fetch_type: 1 只取课程详情， 2 取课程统计， 3两者都取
	// @inject_tag: header:"fetch_type"
	FetchType int32 `protobuf:"varint,2,opt,name=fetch_type,json=fetchType,proto3" json:"fetch_type,omitempty" header:"fetch_type"`
}

func (x *DetailRequest) Reset() {
	*x = DetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Courses_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetailRequest) ProtoMessage() {}

func (x *DetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Courses_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetailRequest.ProtoReflect.Descriptor instead.
func (*DetailRequest) Descriptor() ([]byte, []int) {
	return file_Courses_proto_rawDescGZIP(), []int{4}
}

func (x *DetailRequest) GetCourseId() int32 {
	if x != nil {
		return x.CourseId
	}
	return 0
}

func (x *DetailRequest) GetFetchType() int32 {
	if x != nil {
		return x.FetchType
	}
	return 0
}

// 课程统计表
type CourseCounts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"-"
	CountId int32 `protobuf:"varint,1,opt,name=count_id,json=countId,proto3" json:"-"`
	// @inject_tag: json:"-"
	CourseId   int32  `protobuf:"varint,2,opt,name=course_id,json=courseId,proto3" json:"-"`
	CountKey   string `protobuf:"bytes,3,opt,name=count_key,json=countKey,proto3" json:"count_key,omitempty"`
	CountValue int32  `protobuf:"varint,4,opt,name=count_value,json=countValue,proto3" json:"count_value,omitempty"`
}

func (x *CourseCounts) Reset() {
	*x = CourseCounts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Courses_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CourseCounts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourseCounts) ProtoMessage() {}

func (x *CourseCounts) ProtoReflect() protoreflect.Message {
	mi := &file_Courses_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourseCounts.ProtoReflect.Descriptor instead.
func (*CourseCounts) Descriptor() ([]byte, []int) {
	return file_Courses_proto_rawDescGZIP(), []int{5}
}

func (x *CourseCounts) GetCountId() int32 {
	if x != nil {
		return x.CountId
	}
	return 0
}

func (x *CourseCounts) GetCourseId() int32 {
	if x != nil {
		return x.CourseId
	}
	return 0
}

func (x *CourseCounts) GetCountKey() string {
	if x != nil {
		return x.CountKey
	}
	return ""
}

func (x *CourseCounts) GetCountValue() int32 {
	if x != nil {
		return x.CountValue
	}
	return 0
}

type DetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Course *CourseModel    `protobuf:"bytes,1,opt,name=course,proto3" json:"course,omitempty"`
	Counts []*CourseCounts `protobuf:"bytes,2,rep,name=counts,proto3" json:"counts,omitempty"`
}

func (x *DetailResponse) Reset() {
	*x = DetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Courses_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetailResponse) ProtoMessage() {}

func (x *DetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Courses_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetailResponse.ProtoReflect.Descriptor instead.
func (*DetailResponse) Descriptor() ([]byte, []int) {
	return file_Courses_proto_rawDescGZIP(), []int{6}
}

func (x *DetailResponse) GetCourse() *CourseModel {
	if x != nil {
		return x.Course
	}
	return nil
}

func (x *DetailResponse) GetCounts() []*CourseCounts {
	if x != nil {
		return x.Counts
	}
	return nil
}

var File_Courses_proto protoreflect.FileDescriptor

var file_Courses_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a, 0x09, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x22, 0x21, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x22, 0xeb, 0x01, 0x0a, 0x0b, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x28, 0x0a, 0x10, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x64, 0x69, 0x73, 0x70,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x44, 0x69, 0x73, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x23,
	0x0a, 0x0d, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x32, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x32, 0x12, 0x2c, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x61, 0x64, 0x64, 0x74, 0x69, 0x6d,
	0x65, 0x22, 0x3c, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2c, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0x4b, 0x0a, 0x0d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x66, 0x65, 0x74, 0x63, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x66, 0x65, 0x74, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x22, 0x84, 0x01, 0x0a,
	0x0c, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x19, 0x0a,
	0x08, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4b,
	0x65, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x6d, 0x0a, 0x0e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x2e,
	0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x06, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x2e, 0x43, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x06, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x32, 0x8d, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x72,
	0x54, 0x6f, 0x70, 0x12, 0x14, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x3e, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12,
	0x16, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x73, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Courses_proto_rawDescOnce sync.Once
	file_Courses_proto_rawDescData = file_Courses_proto_rawDesc
)

func file_Courses_proto_rawDescGZIP() []byte {
	file_Courses_proto_rawDescOnce.Do(func() {
		file_Courses_proto_rawDescData = protoimpl.X.CompressGZIP(file_Courses_proto_rawDescData)
	})
	return file_Courses_proto_rawDescData
}

var file_Courses_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_Courses_proto_goTypes = []interface{}{
	(*Timestamp)(nil),           // 0: courses.Timestamp
	(*ListRequest)(nil),         // 1: courses.ListRequest
	(*CourseModel)(nil),         // 2: courses.CourseModel
	(*ListResponse)(nil),        // 3: courses.ListResponse
	(*DetailRequest)(nil),       // 4: courses.DetailRequest
	(*CourseCounts)(nil),        // 5: courses.CourseCounts
	(*DetailResponse)(nil),      // 6: courses.DetailResponse
	(*timestamp.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_Courses_proto_depIdxs = []int32{
	7, // 0: courses.Timestamp.timestamp:type_name -> google.protobuf.Timestamp
	0, // 1: courses.CourseModel.addtime:type_name -> courses.Timestamp
	2, // 2: courses.ListResponse.result:type_name -> courses.CourseModel
	2, // 3: courses.DetailResponse.course:type_name -> courses.CourseModel
	5, // 4: courses.DetailResponse.counts:type_name -> courses.CourseCounts
	1, // 5: courses.CoursesService.ListForTop:input_type -> courses.ListRequest
	4, // 6: courses.CoursesService.GetDetail:input_type -> courses.DetailRequest
	3, // 7: courses.CoursesService.ListForTop:output_type -> courses.ListResponse
	6, // 8: courses.CoursesService.GetDetail:output_type -> courses.DetailResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_Courses_proto_init() }
func file_Courses_proto_init() {
	if File_Courses_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Courses_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Timestamp); i {
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
		file_Courses_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_Courses_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CourseModel); i {
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
		file_Courses_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
		file_Courses_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetailRequest); i {
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
		file_Courses_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CourseCounts); i {
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
		file_Courses_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetailResponse); i {
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
			RawDescriptor: file_Courses_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Courses_proto_goTypes,
		DependencyIndexes: file_Courses_proto_depIdxs,
		MessageInfos:      file_Courses_proto_msgTypes,
	}.Build()
	File_Courses_proto = out.File
	file_Courses_proto_rawDesc = nil
	file_Courses_proto_goTypes = nil
	file_Courses_proto_depIdxs = nil
}
