// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: bookstore.proto

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

type NewBook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Author      string `protobuf:"bytes,2,opt,name=Author,proto3" json:"Author,omitempty"`
	Publication string `protobuf:"bytes,3,opt,name=Publication,proto3" json:"Publication,omitempty"`
}

func (x *NewBook) Reset() {
	*x = NewBook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookstore_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewBook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewBook) ProtoMessage() {}

func (x *NewBook) ProtoReflect() protoreflect.Message {
	mi := &file_bookstore_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewBook.ProtoReflect.Descriptor instead.
func (*NewBook) Descriptor() ([]byte, []int) {
	return file_bookstore_proto_rawDescGZIP(), []int{0}
}

func (x *NewBook) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NewBook) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *NewBook) GetPublication() string {
	if x != nil {
		return x.Publication
	}
	return ""
}

type Book struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Author      string `protobuf:"bytes,2,opt,name=Author,proto3" json:"Author,omitempty"`
	Publication string `protobuf:"bytes,3,opt,name=Publication,proto3" json:"Publication,omitempty"`
}

func (x *Book) Reset() {
	*x = Book{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookstore_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_bookstore_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_bookstore_proto_rawDescGZIP(), []int{1}
}

func (x *Book) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Book) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Book) GetPublication() string {
	if x != nil {
		return x.Publication
	}
	return ""
}

type GetBooksParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetBooksParams) Reset() {
	*x = GetBooksParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookstore_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBooksParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBooksParams) ProtoMessage() {}

func (x *GetBooksParams) ProtoReflect() protoreflect.Message {
	mi := &file_bookstore_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBooksParams.ProtoReflect.Descriptor instead.
func (*GetBooksParams) Descriptor() ([]byte, []int) {
	return file_bookstore_proto_rawDescGZIP(), []int{2}
}

type BooksList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Books []*Book `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
}

func (x *BooksList) Reset() {
	*x = BooksList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookstore_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BooksList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BooksList) ProtoMessage() {}

func (x *BooksList) ProtoReflect() protoreflect.Message {
	mi := &file_bookstore_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BooksList.ProtoReflect.Descriptor instead.
func (*BooksList) Descriptor() ([]byte, []int) {
	return file_bookstore_proto_rawDescGZIP(), []int{3}
}

func (x *BooksList) GetBooks() []*Book {
	if x != nil {
		return x.Books
	}
	return nil
}

var File_bookstore_proto protoreflect.FileDescriptor

var file_bookstore_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x22, 0x57, 0x0a, 0x07, 0x4e, 0x65, 0x77, 0x42, 0x6f,
	0x6f, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x20,
	0x0a, 0x0b, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x54, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x10, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f,
	0x6b, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x2d, 0x0a, 0x09, 0x42, 0x6f, 0x6f, 0x6b,
	0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x42, 0x6f, 0x6f, 0x6b,
	0x52, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x32, 0x70, 0x0a, 0x0e, 0x42, 0x6f, 0x6f, 0x6b, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x0a, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x0d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4e,
	0x65, 0x77, 0x42, 0x6f, 0x6f, 0x6b, 0x1a, 0x0a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x42, 0x6f,
	0x6f, 0x6b, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x08, 0x67, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73,
	0x12, 0x14, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x42, 0x6f,
	0x6f, 0x6b, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bookstore_proto_rawDescOnce sync.Once
	file_bookstore_proto_rawDescData = file_bookstore_proto_rawDesc
)

func file_bookstore_proto_rawDescGZIP() []byte {
	file_bookstore_proto_rawDescOnce.Do(func() {
		file_bookstore_proto_rawDescData = protoimpl.X.CompressGZIP(file_bookstore_proto_rawDescData)
	})
	return file_bookstore_proto_rawDescData
}

var file_bookstore_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_bookstore_proto_goTypes = []interface{}{
	(*NewBook)(nil),        // 0: grpc.NewBook
	(*Book)(nil),           // 1: grpc.Book
	(*GetBooksParams)(nil), // 2: grpc.GetBooksParams
	(*BooksList)(nil),      // 3: grpc.BooksList
}
var file_bookstore_proto_depIdxs = []int32{
	1, // 0: grpc.BooksList.books:type_name -> grpc.Book
	0, // 1: grpc.BookManagement.CreateBook:input_type -> grpc.NewBook
	2, // 2: grpc.BookManagement.getBooks:input_type -> grpc.GetBooksParams
	1, // 3: grpc.BookManagement.CreateBook:output_type -> grpc.Book
	3, // 4: grpc.BookManagement.getBooks:output_type -> grpc.BooksList
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_bookstore_proto_init() }
func file_bookstore_proto_init() {
	if File_bookstore_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bookstore_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewBook); i {
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
		file_bookstore_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Book); i {
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
		file_bookstore_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBooksParams); i {
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
		file_bookstore_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BooksList); i {
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
			RawDescriptor: file_bookstore_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bookstore_proto_goTypes,
		DependencyIndexes: file_bookstore_proto_depIdxs,
		MessageInfos:      file_bookstore_proto_msgTypes,
	}.Build()
	File_bookstore_proto = out.File
	file_bookstore_proto_rawDesc = nil
	file_bookstore_proto_goTypes = nil
	file_bookstore_proto_depIdxs = nil
}
