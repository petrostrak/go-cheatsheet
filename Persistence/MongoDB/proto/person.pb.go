// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: person.proto

package proto

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

type Person struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                     string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Kind                   string   `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	PersonsName            string   `protobuf:"bytes,3,opt,name=persons_name,json=personsName,proto3" json:"persons_name,omitempty"`
	Origins                string   `protobuf:"bytes,4,opt,name=origins,proto3" json:"origins,omitempty"`
	ProgrammingLanguages   []string `protobuf:"bytes,5,rep,name=programming_languages,json=programmingLanguages,proto3" json:"programming_languages,omitempty"`
	Tools                  []string `protobuf:"bytes,6,rep,name=tools,proto3" json:"tools,omitempty"`
	Linkedin               string   `protobuf:"bytes,7,opt,name=linkedin,proto3" json:"linkedin,omitempty"`
	Github                 string   `protobuf:"bytes,8,opt,name=github,proto3" json:"github,omitempty"`
	Personal               string   `protobuf:"bytes,9,opt,name=personal,proto3" json:"personal,omitempty"`
	ForeignLanguages       []string `protobuf:"bytes,10,rep,name=foreign_languages,json=foreignLanguages,proto3" json:"foreign_languages,omitempty"`
	FavFood                string   `protobuf:"bytes,11,opt,name=fav_food,json=favFood,proto3" json:"fav_food,omitempty"`
	FavDrink               string   `protobuf:"bytes,12,opt,name=fav_drink,json=favDrink,proto3" json:"fav_drink,omitempty"`
	FavProgrammingLanguage string   `protobuf:"bytes,13,opt,name=fav_programming_language,json=favProgrammingLanguage,proto3" json:"fav_programming_language,omitempty"`
	ThinkingAbout          []string `protobuf:"bytes,14,rep,name=thinking_about,json=thinkingAbout,proto3" json:"thinking_about,omitempty"`
	Hobbies                []string `protobuf:"bytes,15,rep,name=hobbies,proto3" json:"hobbies,omitempty"`
}

func (x *Person) Reset() {
	*x = Person{}
	if protoimpl.UnsafeEnabled {
		mi := &file_person_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_person_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person.ProtoReflect.Descriptor instead.
func (*Person) Descriptor() ([]byte, []int) {
	return file_person_proto_rawDescGZIP(), []int{0}
}

func (x *Person) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Person) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *Person) GetPersonsName() string {
	if x != nil {
		return x.PersonsName
	}
	return ""
}

func (x *Person) GetOrigins() string {
	if x != nil {
		return x.Origins
	}
	return ""
}

func (x *Person) GetProgrammingLanguages() []string {
	if x != nil {
		return x.ProgrammingLanguages
	}
	return nil
}

func (x *Person) GetTools() []string {
	if x != nil {
		return x.Tools
	}
	return nil
}

func (x *Person) GetLinkedin() string {
	if x != nil {
		return x.Linkedin
	}
	return ""
}

func (x *Person) GetGithub() string {
	if x != nil {
		return x.Github
	}
	return ""
}

func (x *Person) GetPersonal() string {
	if x != nil {
		return x.Personal
	}
	return ""
}

func (x *Person) GetForeignLanguages() []string {
	if x != nil {
		return x.ForeignLanguages
	}
	return nil
}

func (x *Person) GetFavFood() string {
	if x != nil {
		return x.FavFood
	}
	return ""
}

func (x *Person) GetFavDrink() string {
	if x != nil {
		return x.FavDrink
	}
	return ""
}

func (x *Person) GetFavProgrammingLanguage() string {
	if x != nil {
		return x.FavProgrammingLanguage
	}
	return ""
}

func (x *Person) GetThinkingAbout() []string {
	if x != nil {
		return x.ThinkingAbout
	}
	return nil
}

func (x *Person) GetHobbies() []string {
	if x != nil {
		return x.Hobbies
	}
	return nil
}

type PersonId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PersonId) Reset() {
	*x = PersonId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_person_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonId) ProtoMessage() {}

func (x *PersonId) ProtoReflect() protoreflect.Message {
	mi := &file_person_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonId.ProtoReflect.Descriptor instead.
func (*PersonId) Descriptor() ([]byte, []int) {
	return file_person_proto_rawDescGZIP(), []int{1}
}

func (x *PersonId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_person_proto protoreflect.FileDescriptor

var file_person_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe4, 0x03, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b,
	0x69, 0x6e, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x73,
	0x12, 0x33, 0x0a, 0x15, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x6d, 0x69, 0x6e, 0x67, 0x5f,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x14, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x6d, 0x69, 0x6e, 0x67, 0x4c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6c,
	0x69, 0x6e, 0x6b, 0x65, 0x64, 0x69, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c,
	0x69, 0x6e, 0x6b, 0x65, 0x64, 0x69, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x12, 0x2b, 0x0a, 0x11, 0x66,
	0x6f, 0x72, 0x65, 0x69, 0x67, 0x6e, 0x5f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73,
	0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x66, 0x6f, 0x72, 0x65, 0x69, 0x67, 0x6e, 0x4c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x61, 0x76, 0x5f,
	0x66, 0x6f, 0x6f, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x61, 0x76, 0x46,
	0x6f, 0x6f, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x61, 0x76, 0x5f, 0x64, 0x72, 0x69, 0x6e, 0x6b,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x61, 0x76, 0x44, 0x72, 0x69, 0x6e, 0x6b,
	0x12, 0x38, 0x0a, 0x18, 0x66, 0x61, 0x76, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x6d,
	0x69, 0x6e, 0x67, 0x5f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x16, 0x66, 0x61, 0x76, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x6d, 0x69,
	0x6e, 0x67, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x68,
	0x69, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x18, 0x0e, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0d, 0x74, 0x68, 0x69, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x41, 0x62, 0x6f, 0x75,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x6f, 0x62, 0x62, 0x69, 0x65, 0x73, 0x18, 0x0f, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x07, 0x68, 0x6f, 0x62, 0x62, 0x69, 0x65, 0x73, 0x22, 0x1a, 0x0a, 0x08, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0xa2, 0x02, 0x0a, 0x0d, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x0c, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x0f, 0x2e, 0x6d, 0x6f, 0x6e, 0x67,
	0x6f, 0x64, 0x62, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x1a, 0x11, 0x2e, 0x6d, 0x6f, 0x6e,
	0x67, 0x6f, 0x64, 0x62, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x30, 0x0a,
	0x0a, 0x52, 0x65, 0x61, 0x64, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x11, 0x2e, 0x6d, 0x6f,
	0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x1a, 0x0f,
	0x2e, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12,
	0x37, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12,
	0x0f, 0x2e, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x39, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x11, 0x2e, 0x6d, 0x6f, 0x6e, 0x67, 0x6f,
	0x64, 0x62, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x37, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0f, 0x2e, 0x6d, 0x6f, 0x6e, 0x67,
	0x6f, 0x64, 0x62, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x30, 0x01, 0x42, 0x10, 0x5a, 0x0e,
	0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x2d, 0x64, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_person_proto_rawDescOnce sync.Once
	file_person_proto_rawDescData = file_person_proto_rawDesc
)

func file_person_proto_rawDescGZIP() []byte {
	file_person_proto_rawDescOnce.Do(func() {
		file_person_proto_rawDescData = protoimpl.X.CompressGZIP(file_person_proto_rawDescData)
	})
	return file_person_proto_rawDescData
}

var file_person_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_person_proto_goTypes = []interface{}{
	(*Person)(nil),        // 0: mongodb.Person
	(*PersonId)(nil),      // 1: mongodb.PersonId
	(*emptypb.Empty)(nil), // 2: google.protobuf.Empty
}
var file_person_proto_depIdxs = []int32{
	0, // 0: mongodb.PersonService.CreatePerson:input_type -> mongodb.Person
	1, // 1: mongodb.PersonService.ReadPerson:input_type -> mongodb.PersonId
	0, // 2: mongodb.PersonService.UpdatePerson:input_type -> mongodb.Person
	1, // 3: mongodb.PersonService.DeletePerson:input_type -> mongodb.PersonId
	2, // 4: mongodb.PersonService.ListPerson:input_type -> google.protobuf.Empty
	1, // 5: mongodb.PersonService.CreatePerson:output_type -> mongodb.PersonId
	0, // 6: mongodb.PersonService.ReadPerson:output_type -> mongodb.Person
	2, // 7: mongodb.PersonService.UpdatePerson:output_type -> google.protobuf.Empty
	2, // 8: mongodb.PersonService.DeletePerson:output_type -> google.protobuf.Empty
	0, // 9: mongodb.PersonService.ListPerson:output_type -> mongodb.Person
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_person_proto_init() }
func file_person_proto_init() {
	if File_person_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_person_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Person); i {
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
		file_person_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersonId); i {
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
			RawDescriptor: file_person_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_person_proto_goTypes,
		DependencyIndexes: file_person_proto_depIdxs,
		MessageInfos:      file_person_proto_msgTypes,
	}.Build()
	File_person_proto = out.File
	file_person_proto_rawDesc = nil
	file_person_proto_goTypes = nil
	file_person_proto_depIdxs = nil
}
