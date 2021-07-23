// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.15.7
// source: proto/account/query.proto

package account

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

// 查询列表的请求
type QueryListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int64 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"` // 偏移值
	Count  int64 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`   // 数量
}

func (x *QueryListRequest) Reset() {
	*x = QueryListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_query_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryListRequest) ProtoMessage() {}

func (x *QueryListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_query_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryListRequest.ProtoReflect.Descriptor instead.
func (*QueryListRequest) Descriptor() ([]byte, []int) {
	return file_proto_account_query_proto_rawDescGZIP(), []int{0}
}

func (x *QueryListRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *QueryListRequest) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

// 查询列表的回复
type QueryListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  *Status          `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`   // 状态
	Total   int64            `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`    // 总数
	Account []*AccountEntity `protobuf:"bytes,3,rep,name=account,proto3" json:"account,omitempty"` // 账号列表
}

func (x *QueryListResponse) Reset() {
	*x = QueryListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_query_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryListResponse) ProtoMessage() {}

func (x *QueryListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_query_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryListResponse.ProtoReflect.Descriptor instead.
func (*QueryListResponse) Descriptor() ([]byte, []int) {
	return file_proto_account_query_proto_rawDescGZIP(), []int{1}
}

func (x *QueryListResponse) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *QueryListResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *QueryListResponse) GetAccount() []*AccountEntity {
	if x != nil {
		return x.Account
	}
	return nil
}

// 查询个体的请求
type QuerySingleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field QueryField `protobuf:"varint,1,opt,name=field,proto3,enum=QueryField" json:"field,omitempty"` // 字段
	Value string     `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`                  // 值
}

func (x *QuerySingleRequest) Reset() {
	*x = QuerySingleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_query_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuerySingleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuerySingleRequest) ProtoMessage() {}

func (x *QuerySingleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_query_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuerySingleRequest.ProtoReflect.Descriptor instead.
func (*QuerySingleRequest) Descriptor() ([]byte, []int) {
	return file_proto_account_query_proto_rawDescGZIP(), []int{2}
}

func (x *QuerySingleRequest) GetField() QueryField {
	if x != nil {
		return x.Field
	}
	return QueryField_QUERY_FIELD_NONE
}

func (x *QuerySingleRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// 查询列表的回复
type QuerySingleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  *Status        `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`   // 状态
	Account *AccountEntity `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"` // 账号
}

func (x *QuerySingleResponse) Reset() {
	*x = QuerySingleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_query_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuerySingleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuerySingleResponse) ProtoMessage() {}

func (x *QuerySingleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_query_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuerySingleResponse.ProtoReflect.Descriptor instead.
func (*QuerySingleResponse) Descriptor() ([]byte, []int) {
	return file_proto_account_query_proto_rawDescGZIP(), []int{3}
}

func (x *QuerySingleResponse) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *QuerySingleResponse) GetAccount() *AccountEntity {
	if x != nil {
		return x.Account
	}
	return nil
}

var File_proto_account_query_proto protoreflect.FileDescriptor

var file_proto_account_query_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x40, 0x0a, 0x10, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x74, 0x0a, 0x11, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x28, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x4d, 0x0a, 0x12, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x60,
	0x0a, 0x13, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x28, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x32, 0x6f, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x2f, 0x0a, 0x04, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x11, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x06, 0x53, 0x69,
	0x6e, 0x67, 0x6c, 0x65, 0x12, 0x13, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x69, 0x6e, 0x67,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x17, 0x5a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x3b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_account_query_proto_rawDescOnce sync.Once
	file_proto_account_query_proto_rawDescData = file_proto_account_query_proto_rawDesc
)

func file_proto_account_query_proto_rawDescGZIP() []byte {
	file_proto_account_query_proto_rawDescOnce.Do(func() {
		file_proto_account_query_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_account_query_proto_rawDescData)
	})
	return file_proto_account_query_proto_rawDescData
}

var file_proto_account_query_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_account_query_proto_goTypes = []interface{}{
	(*QueryListRequest)(nil),    // 0: QueryListRequest
	(*QueryListResponse)(nil),   // 1: QueryListResponse
	(*QuerySingleRequest)(nil),  // 2: QuerySingleRequest
	(*QuerySingleResponse)(nil), // 3: QuerySingleResponse
	(*Status)(nil),              // 4: Status
	(*AccountEntity)(nil),       // 5: AccountEntity
	(QueryField)(0),             // 6: QueryField
}
var file_proto_account_query_proto_depIdxs = []int32{
	4, // 0: QueryListResponse.status:type_name -> Status
	5, // 1: QueryListResponse.account:type_name -> AccountEntity
	6, // 2: QuerySingleRequest.field:type_name -> QueryField
	4, // 3: QuerySingleResponse.status:type_name -> Status
	5, // 4: QuerySingleResponse.account:type_name -> AccountEntity
	0, // 5: Query.List:input_type -> QueryListRequest
	2, // 6: Query.Single:input_type -> QuerySingleRequest
	1, // 7: Query.List:output_type -> QueryListResponse
	3, // 8: Query.Single:output_type -> QuerySingleResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_account_query_proto_init() }
func file_proto_account_query_proto_init() {
	if File_proto_account_query_proto != nil {
		return
	}
	file_proto_account_shared_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_account_query_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryListRequest); i {
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
		file_proto_account_query_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryListResponse); i {
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
		file_proto_account_query_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuerySingleRequest); i {
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
		file_proto_account_query_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuerySingleResponse); i {
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
			RawDescriptor: file_proto_account_query_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_account_query_proto_goTypes,
		DependencyIndexes: file_proto_account_query_proto_depIdxs,
		MessageInfos:      file_proto_account_query_proto_msgTypes,
	}.Build()
	File_proto_account_query_proto = out.File
	file_proto_account_query_proto_rawDesc = nil
	file_proto_account_query_proto_goTypes = nil
	file_proto_account_query_proto_depIdxs = nil
}
