// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: proto/rbacAccess.proto

package rbacAccess

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

type AccessModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ModuleName  string         `protobuf:"bytes,2,opt,name=moduleName,proto3" json:"moduleName,omitempty"`
	ActionName  string         `protobuf:"bytes,3,opt,name=actionName,proto3" json:"actionName,omitempty"`
	Type        int64          `protobuf:"varint,4,opt,name=type,proto3" json:"type,omitempty"`
	Url         string         `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
	ModuleId    int64          `protobuf:"varint,6,opt,name=moduleId,proto3" json:"moduleId,omitempty"`
	Sort        int64          `protobuf:"varint,7,opt,name=sort,proto3" json:"sort,omitempty"`
	Description string         `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	Status      int64          `protobuf:"varint,9,opt,name=status,proto3" json:"status,omitempty"`
	AddTime     int64          `protobuf:"varint,10,opt,name=addTime,proto3" json:"addTime,omitempty"`
	AccessItem  []*AccessModel `protobuf:"bytes,11,rep,name=accessItem,proto3" json:"accessItem,omitempty"`
}

func (x *AccessModel) Reset() {
	*x = AccessModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rbacAccess_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessModel) ProtoMessage() {}

func (x *AccessModel) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rbacAccess_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessModel.ProtoReflect.Descriptor instead.
func (*AccessModel) Descriptor() ([]byte, []int) {
	return file_proto_rbacAccess_proto_rawDescGZIP(), []int{0}
}

func (x *AccessModel) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AccessModel) GetModuleName() string {
	if x != nil {
		return x.ModuleName
	}
	return ""
}

func (x *AccessModel) GetActionName() string {
	if x != nil {
		return x.ActionName
	}
	return ""
}

func (x *AccessModel) GetType() int64 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *AccessModel) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *AccessModel) GetModuleId() int64 {
	if x != nil {
		return x.ModuleId
	}
	return 0
}

func (x *AccessModel) GetSort() int64 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *AccessModel) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AccessModel) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *AccessModel) GetAddTime() int64 {
	if x != nil {
		return x.AddTime
	}
	return 0
}

func (x *AccessModel) GetAccessItem() []*AccessModel {
	if x != nil {
		return x.AccessItem
	}
	return nil
}

type AccessGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AccessGetRequest) Reset() {
	*x = AccessGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rbacAccess_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessGetRequest) ProtoMessage() {}

func (x *AccessGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rbacAccess_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessGetRequest.ProtoReflect.Descriptor instead.
func (*AccessGetRequest) Descriptor() ([]byte, []int) {
	return file_proto_rbacAccess_proto_rawDescGZIP(), []int{1}
}

func (x *AccessGetRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type AccessGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessList []*AccessModel `protobuf:"bytes,1,rep,name=accessList,proto3" json:"accessList,omitempty"`
}

func (x *AccessGetResponse) Reset() {
	*x = AccessGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rbacAccess_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessGetResponse) ProtoMessage() {}

func (x *AccessGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rbacAccess_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessGetResponse.ProtoReflect.Descriptor instead.
func (*AccessGetResponse) Descriptor() ([]byte, []int) {
	return file_proto_rbacAccess_proto_rawDescGZIP(), []int{2}
}

func (x *AccessGetResponse) GetAccessList() []*AccessModel {
	if x != nil {
		return x.AccessList
	}
	return nil
}

type AccessAddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModuleName  string `protobuf:"bytes,1,opt,name=moduleName,proto3" json:"moduleName,omitempty"`
	Type        int64  `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	ActionName  string `protobuf:"bytes,3,opt,name=actionName,proto3" json:"actionName,omitempty"`
	Url         string `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	ModuleId    int64  `protobuf:"varint,5,opt,name=moduleId,proto3" json:"moduleId,omitempty"`
	Sort        int64  `protobuf:"varint,6,opt,name=sort,proto3" json:"sort,omitempty"`
	Description string `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	Status      int64  `protobuf:"varint,8,opt,name=status,proto3" json:"status,omitempty"`
	AddTime     int64  `protobuf:"varint,9,opt,name=addTime,proto3" json:"addTime,omitempty"`
}

func (x *AccessAddRequest) Reset() {
	*x = AccessAddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rbacAccess_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessAddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessAddRequest) ProtoMessage() {}

func (x *AccessAddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rbacAccess_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessAddRequest.ProtoReflect.Descriptor instead.
func (*AccessAddRequest) Descriptor() ([]byte, []int) {
	return file_proto_rbacAccess_proto_rawDescGZIP(), []int{3}
}

func (x *AccessAddRequest) GetModuleName() string {
	if x != nil {
		return x.ModuleName
	}
	return ""
}

func (x *AccessAddRequest) GetType() int64 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *AccessAddRequest) GetActionName() string {
	if x != nil {
		return x.ActionName
	}
	return ""
}

func (x *AccessAddRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *AccessAddRequest) GetModuleId() int64 {
	if x != nil {
		return x.ModuleId
	}
	return 0
}

func (x *AccessAddRequest) GetSort() int64 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *AccessAddRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AccessAddRequest) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *AccessAddRequest) GetAddTime() int64 {
	if x != nil {
		return x.AddTime
	}
	return 0
}

type AccessAddResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *AccessAddResponse) Reset() {
	*x = AccessAddResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rbacAccess_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessAddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessAddResponse) ProtoMessage() {}

func (x *AccessAddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rbacAccess_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessAddResponse.ProtoReflect.Descriptor instead.
func (*AccessAddResponse) Descriptor() ([]byte, []int) {
	return file_proto_rbacAccess_proto_rawDescGZIP(), []int{4}
}

func (x *AccessAddResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *AccessAddResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type AccessEditRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ModuleName  string `protobuf:"bytes,2,opt,name=moduleName,proto3" json:"moduleName,omitempty"`
	ActionName  string `protobuf:"bytes,3,opt,name=actionName,proto3" json:"actionName,omitempty"`
	Type        int64  `protobuf:"varint,4,opt,name=type,proto3" json:"type,omitempty"`
	Url         string `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
	ModuleId    int64  `protobuf:"varint,6,opt,name=moduleId,proto3" json:"moduleId,omitempty"`
	Sort        int64  `protobuf:"varint,7,opt,name=sort,proto3" json:"sort,omitempty"`
	Description string `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	Status      int64  `protobuf:"varint,9,opt,name=status,proto3" json:"status,omitempty"`
	AddTime     int64  `protobuf:"varint,10,opt,name=addTime,proto3" json:"addTime,omitempty"`
}

func (x *AccessEditRequest) Reset() {
	*x = AccessEditRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rbacAccess_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessEditRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessEditRequest) ProtoMessage() {}

func (x *AccessEditRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rbacAccess_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessEditRequest.ProtoReflect.Descriptor instead.
func (*AccessEditRequest) Descriptor() ([]byte, []int) {
	return file_proto_rbacAccess_proto_rawDescGZIP(), []int{5}
}

func (x *AccessEditRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AccessEditRequest) GetModuleName() string {
	if x != nil {
		return x.ModuleName
	}
	return ""
}

func (x *AccessEditRequest) GetActionName() string {
	if x != nil {
		return x.ActionName
	}
	return ""
}

func (x *AccessEditRequest) GetType() int64 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *AccessEditRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *AccessEditRequest) GetModuleId() int64 {
	if x != nil {
		return x.ModuleId
	}
	return 0
}

func (x *AccessEditRequest) GetSort() int64 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *AccessEditRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AccessEditRequest) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *AccessEditRequest) GetAddTime() int64 {
	if x != nil {
		return x.AddTime
	}
	return 0
}

type AccessEditResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *AccessEditResponse) Reset() {
	*x = AccessEditResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rbacAccess_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessEditResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessEditResponse) ProtoMessage() {}

func (x *AccessEditResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rbacAccess_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessEditResponse.ProtoReflect.Descriptor instead.
func (*AccessEditResponse) Descriptor() ([]byte, []int) {
	return file_proto_rbacAccess_proto_rawDescGZIP(), []int{6}
}

func (x *AccessEditResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *AccessEditResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type AccessDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AccessDeleteRequest) Reset() {
	*x = AccessDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rbacAccess_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessDeleteRequest) ProtoMessage() {}

func (x *AccessDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rbacAccess_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessDeleteRequest.ProtoReflect.Descriptor instead.
func (*AccessDeleteRequest) Descriptor() ([]byte, []int) {
	return file_proto_rbacAccess_proto_rawDescGZIP(), []int{7}
}

func (x *AccessDeleteRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type AccessDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *AccessDeleteResponse) Reset() {
	*x = AccessDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rbacAccess_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessDeleteResponse) ProtoMessage() {}

func (x *AccessDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rbacAccess_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessDeleteResponse.ProtoReflect.Descriptor instead.
func (*AccessDeleteResponse) Descriptor() ([]byte, []int) {
	return file_proto_rbacAccess_proto_rawDescGZIP(), []int{8}
}

func (x *AccessDeleteResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *AccessDeleteResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_rbacAccess_proto protoreflect.FileDescriptor

var file_proto_rbacAccess_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x62, 0x61, 0x63, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x22, 0xbc, 0x02, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x61, 0x64, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x0a, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x52, 0x0a, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x22,
	0x22, 0x0a, 0x10, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x48, 0x0a, 0x11, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x52, 0x0a, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x22, 0xfc, 0x01,
	0x0a, 0x10, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x61, 0x64, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x47, 0x0a, 0x11,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x8d, 0x02, 0x0a, 0x11, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x45, 0x64, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x6f, 0x72,
	0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x48, 0x0a, 0x12, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x45,
	0x64, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x25, 0x0a, 0x13, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4a, 0x0a, 0x14, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x32, 0xa8, 0x02, 0x0a, 0x0a, 0x52, 0x62, 0x61, 0x63, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x12, 0x42, 0x0a, 0x09, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x47, 0x65, 0x74, 0x12, 0x18,
	0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x09, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x41,
	0x64, 0x64, 0x12, 0x18, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x41, 0x64, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0a, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x45, 0x64, 0x69, 0x74, 0x12, 0x19, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x45, 0x64, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x45, 0x64, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x4b, 0x0a, 0x0c, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x1b, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x14, 0x5a,
	0x12, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x62, 0x61, 0x63, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_rbacAccess_proto_rawDescOnce sync.Once
	file_proto_rbacAccess_proto_rawDescData = file_proto_rbacAccess_proto_rawDesc
)

func file_proto_rbacAccess_proto_rawDescGZIP() []byte {
	file_proto_rbacAccess_proto_rawDescOnce.Do(func() {
		file_proto_rbacAccess_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_rbacAccess_proto_rawDescData)
	})
	return file_proto_rbacAccess_proto_rawDescData
}

var file_proto_rbacAccess_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_rbacAccess_proto_goTypes = []any{
	(*AccessModel)(nil),          // 0: access.AccessModel
	(*AccessGetRequest)(nil),     // 1: access.AccessGetRequest
	(*AccessGetResponse)(nil),    // 2: access.AccessGetResponse
	(*AccessAddRequest)(nil),     // 3: access.AccessAddRequest
	(*AccessAddResponse)(nil),    // 4: access.AccessAddResponse
	(*AccessEditRequest)(nil),    // 5: access.AccessEditRequest
	(*AccessEditResponse)(nil),   // 6: access.AccessEditResponse
	(*AccessDeleteRequest)(nil),  // 7: access.AccessDeleteRequest
	(*AccessDeleteResponse)(nil), // 8: access.AccessDeleteResponse
}
var file_proto_rbacAccess_proto_depIdxs = []int32{
	0, // 0: access.AccessModel.accessItem:type_name -> access.AccessModel
	0, // 1: access.AccessGetResponse.accessList:type_name -> access.AccessModel
	1, // 2: access.RbacAccess.AccessGet:input_type -> access.AccessGetRequest
	3, // 3: access.RbacAccess.AccessAdd:input_type -> access.AccessAddRequest
	5, // 4: access.RbacAccess.AccessEdit:input_type -> access.AccessEditRequest
	7, // 5: access.RbacAccess.AccessDelete:input_type -> access.AccessDeleteRequest
	2, // 6: access.RbacAccess.AccessGet:output_type -> access.AccessGetResponse
	4, // 7: access.RbacAccess.AccessAdd:output_type -> access.AccessAddResponse
	6, // 8: access.RbacAccess.AccessEdit:output_type -> access.AccessEditResponse
	8, // 9: access.RbacAccess.AccessDelete:output_type -> access.AccessDeleteResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_rbacAccess_proto_init() }
func file_proto_rbacAccess_proto_init() {
	if File_proto_rbacAccess_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_rbacAccess_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AccessModel); i {
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
		file_proto_rbacAccess_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AccessGetRequest); i {
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
		file_proto_rbacAccess_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*AccessGetResponse); i {
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
		file_proto_rbacAccess_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*AccessAddRequest); i {
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
		file_proto_rbacAccess_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*AccessAddResponse); i {
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
		file_proto_rbacAccess_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*AccessEditRequest); i {
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
		file_proto_rbacAccess_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*AccessEditResponse); i {
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
		file_proto_rbacAccess_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*AccessDeleteRequest); i {
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
		file_proto_rbacAccess_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*AccessDeleteResponse); i {
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
			RawDescriptor: file_proto_rbacAccess_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_rbacAccess_proto_goTypes,
		DependencyIndexes: file_proto_rbacAccess_proto_depIdxs,
		MessageInfos:      file_proto_rbacAccess_proto_msgTypes,
	}.Build()
	File_proto_rbacAccess_proto = out.File
	file_proto_rbacAccess_proto_rawDesc = nil
	file_proto_rbacAccess_proto_goTypes = nil
	file_proto_rbacAccess_proto_depIdxs = nil
}
