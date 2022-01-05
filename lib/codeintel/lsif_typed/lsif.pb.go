// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: lib/codeintel/lsif_typed/lsif.proto

package lsif_typed

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

type Moniker_Unique int32

const (
	Moniker_UNIQUE_UNSPECIFIED Moniker_Unique = 0
	Moniker_UNIQUE_DOCUMENT    Moniker_Unique = 1
	Moniker_UNIQUE_PROJECT     Moniker_Unique = 2
	Moniker_UNIQUE_WORKSPACE   Moniker_Unique = 3
	Moniker_UNIQUE_SCHEME      Moniker_Unique = 4
	Moniker_UNIQUE_GLOBAL      Moniker_Unique = 5
)

// Enum value maps for Moniker_Unique.
var (
	Moniker_Unique_name = map[int32]string{
		0: "UNIQUE_UNSPECIFIED",
		1: "UNIQUE_DOCUMENT",
		2: "UNIQUE_PROJECT",
		3: "UNIQUE_WORKSPACE",
		4: "UNIQUE_SCHEME",
		5: "UNIQUE_GLOBAL",
	}
	Moniker_Unique_value = map[string]int32{
		"UNIQUE_UNSPECIFIED": 0,
		"UNIQUE_DOCUMENT":    1,
		"UNIQUE_PROJECT":     2,
		"UNIQUE_WORKSPACE":   3,
		"UNIQUE_SCHEME":      4,
		"UNIQUE_GLOBAL":      5,
	}
)

func (x Moniker_Unique) Enum() *Moniker_Unique {
	p := new(Moniker_Unique)
	*p = x
	return p
}

func (x Moniker_Unique) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Moniker_Unique) Descriptor() protoreflect.EnumDescriptor {
	return file_lib_codeintel_lsif_typed_lsif_proto_enumTypes[0].Descriptor()
}

func (Moniker_Unique) Type() protoreflect.EnumType {
	return &file_lib_codeintel_lsif_typed_lsif_proto_enumTypes[0]
}

func (x Moniker_Unique) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Moniker_Unique.Descriptor instead.
func (Moniker_Unique) EnumDescriptor() ([]byte, []int) {
	return file_lib_codeintel_lsif_typed_lsif_proto_rawDescGZIP(), []int{4, 0}
}

type MonikerOccurrence_Role int32

const (
	MonikerOccurrence_ROLE_UNSPECIFIED MonikerOccurrence_Role = 0
	MonikerOccurrence_ROLE_DEFINITION  MonikerOccurrence_Role = 1
	MonikerOccurrence_ROLE_REFERENCE   MonikerOccurrence_Role = 2
)

// Enum value maps for MonikerOccurrence_Role.
var (
	MonikerOccurrence_Role_name = map[int32]string{
		0: "ROLE_UNSPECIFIED",
		1: "ROLE_DEFINITION",
		2: "ROLE_REFERENCE",
	}
	MonikerOccurrence_Role_value = map[string]int32{
		"ROLE_UNSPECIFIED": 0,
		"ROLE_DEFINITION":  1,
		"ROLE_REFERENCE":   2,
	}
)

func (x MonikerOccurrence_Role) Enum() *MonikerOccurrence_Role {
	p := new(MonikerOccurrence_Role)
	*p = x
	return p
}

func (x MonikerOccurrence_Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MonikerOccurrence_Role) Descriptor() protoreflect.EnumDescriptor {
	return file_lib_codeintel_lsif_typed_lsif_proto_enumTypes[1].Descriptor()
}

func (MonikerOccurrence_Role) Type() protoreflect.EnumType {
	return &file_lib_codeintel_lsif_typed_lsif_proto_enumTypes[1]
}

func (x MonikerOccurrence_Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MonikerOccurrence_Role.Descriptor instead.
func (MonikerOccurrence_Role) EnumDescriptor() ([]byte, []int) {
	return file_lib_codeintel_lsif_typed_lsif_proto_rawDescGZIP(), []int{5, 0}
}

// LsifValue defines a single LSIF value that. An LSIF index is a stream of values.
type LsifValues struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []*LsifValue `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *LsifValues) Reset() {
	*x = LsifValues{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_codeintel_lsif_typed_lsif_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LsifValues) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LsifValues) ProtoMessage() {}

func (x *LsifValues) ProtoReflect() protoreflect.Message {
	mi := &file_lib_codeintel_lsif_typed_lsif_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LsifValues.ProtoReflect.Descriptor instead.
func (*LsifValues) Descriptor() ([]byte, []int) {
	return file_lib_codeintel_lsif_typed_lsif_proto_rawDescGZIP(), []int{0}
}

func (x *LsifValues) GetValues() []*LsifValue {
	if x != nil {
		return x.Values
	}
	return nil
}

type LsifValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ParentTypes that are assignable to Value:
	//	*LsifValue_Package
	//	*LsifValue_Document
	//	*LsifValue_Moniker
	Value isLsifValue_Value `protobuf_oneof:"value"`
}

func (x *LsifValue) Reset() {
	*x = LsifValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_codeintel_lsif_typed_lsif_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LsifValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LsifValue) ProtoMessage() {}

func (x *LsifValue) ProtoReflect() protoreflect.Message {
	mi := &file_lib_codeintel_lsif_typed_lsif_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LsifValue.ProtoReflect.Descriptor instead.
func (*LsifValue) Descriptor() ([]byte, []int) {
	return file_lib_codeintel_lsif_typed_lsif_proto_rawDescGZIP(), []int{1}
}

func (m *LsifValue) GetValue() isLsifValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *LsifValue) GetPackage() *Package {
	if x, ok := x.GetValue().(*LsifValue_Package); ok {
		return x.Package
	}
	return nil
}

func (x *LsifValue) GetDocument() *Document {
	if x, ok := x.GetValue().(*LsifValue_Document); ok {
		return x.Document
	}
	return nil
}

func (x *LsifValue) GetMoniker() *Moniker {
	if x, ok := x.GetValue().(*LsifValue_Moniker); ok {
		return x.Moniker
	}
	return nil
}

type isLsifValue_Value interface {
	isLsifValue_Value()
}

type LsifValue_Package struct {
	Package *Package `protobuf:"bytes,1,opt,name=package,proto3,oneof"`
}

type LsifValue_Document struct {
	Document *Document `protobuf:"bytes,2,opt,name=document,proto3,oneof"`
}

type LsifValue_Moniker struct {
	Moniker *Moniker `protobuf:"bytes,3,opt,name=moniker,proto3,oneof"`
}

func (*LsifValue_Package) isLsifValue_Value() {}

func (*LsifValue_Document) isLsifValue_Value() {}

func (*LsifValue_Moniker) isLsifValue_Value() {}

// Position defines a single offset in a source file.
type Position struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 0-based line number in the source file. Editors typically use 1-based line
	// numbers so make sure to increment this number before displaying it in a UI.
	// References the last line in the source if this number exceeds the actual number
	// of lines in the source file.
	Line int32 `protobuf:"varint,1,opt,name=line,proto3" json:"line,omitempty"`
	// 0-based column number of the line assuming UTF-8 string encoding.
	// References the last character in this line if this number exceeds the
	// actual number of characters in line.
	Character int32 `protobuf:"varint,2,opt,name=character,proto3" json:"character,omitempty"`
}

func (x *Position) Reset() {
	*x = Position{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_codeintel_lsif_typed_lsif_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Position) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Position) ProtoMessage() {}

func (x *Position) ProtoReflect() protoreflect.Message {
	mi := &file_lib_codeintel_lsif_typed_lsif_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Position.ProtoReflect.Descriptor instead.
func (*Position) Descriptor() ([]byte, []int) {
	return file_lib_codeintel_lsif_typed_lsif_proto_rawDescGZIP(), []int{2}
}

func (x *Position) GetLine() int32 {
	if x != nil {
		return x.Line
	}
	return 0
}

func (x *Position) GetCharacter() int32 {
	if x != nil {
		return x.Character
	}
	return 0
}

// Range defines the slice of a source file between two offset positions.
type Range struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The beginning of this range.
	Start *Position `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	// The end of this range. Must appear later in the file than the start position.
	End *Position `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *Range) Reset() {
	*x = Range{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_codeintel_lsif_typed_lsif_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Range) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Range) ProtoMessage() {}

func (x *Range) ProtoReflect() protoreflect.Message {
	mi := &file_lib_codeintel_lsif_typed_lsif_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Range.ProtoReflect.Descriptor instead.
func (*Range) Descriptor() ([]byte, []int) {
	return file_lib_codeintel_lsif_typed_lsif_proto_rawDescGZIP(), []int{3}
}

func (x *Range) GetStart() *Position {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *Range) GetEnd() *Position {
	if x != nil {
		return x.End
	}
	return nil
}

// Moniker defines a symbol, such as a function or an interface.
type Moniker struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The identifier of this moniker, which can be referenced from
	// MonikerOccurence. An empty id means this moniker can be ignored.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// (optional) the syntax that is used for the scheme. Only relevant when
	// the `unique` field has value `Unique.SCHEME`.
	Scheme string `protobuf:"bytes,2,opt,name=scheme,proto3" json:"scheme,omitempty"`
	// Determines how unique the id field is. For example, it is allowed for two
	// distinct monikers to have the same ID as long as they origin from different
	// document/project/workspace/scheme and have the appropriate `unique` field
	// value. Monikers with different `unique` values must never be compared.
	Unique Moniker_Unique `protobuf:"varint,3,opt,name=unique,proto3,enum=lib.codeintel.lsif_typed.Moniker_Unique" json:"unique,omitempty"`
	// (optional, but strongly recommended) The markdown-formatted documentation
	// for this moniker. This field is repeated to allow different kinds of
	// documentation.  For example, it's nice to include both the signature of a
	// method (parameters and return type) along with the accompanying docstring.
	MarkdownHover []string `protobuf:"bytes,4,rep,name=markdown_hover,json=markdownHover,proto3" json:"markdown_hover,omitempty"`
	// (optional) Links to other monikers that this moniker "implements" in some form. For
	// example, if this moniker represents a function that implements another
	// function from an interface, then this field would link to the other
	// moniker.
	ImplementationMonikers []string `protobuf:"bytes,5,rep,name=implementation_monikers,json=implementationMonikers,proto3" json:"implementation_monikers,omitempty"`
	// (optional) Links to the original package that defines this moniker to
	// enable navigation across different LSIF indexes (whether they come from
	// different projects or git repositories).
	PackageId string `protobuf:"bytes,6,opt,name=package_id,json=packageId,proto3" json:"package_id,omitempty"`
	// "import", "export", or "implementation"
	Kind string `protobuf:"bytes,7,opt,name=kind,proto3" json:"kind,omitempty"`
}

func (x *Moniker) Reset() {
	*x = Moniker{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_codeintel_lsif_typed_lsif_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Moniker) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Moniker) ProtoMessage() {}

func (x *Moniker) ProtoReflect() protoreflect.Message {
	mi := &file_lib_codeintel_lsif_typed_lsif_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Moniker.ProtoReflect.Descriptor instead.
func (*Moniker) Descriptor() ([]byte, []int) {
	return file_lib_codeintel_lsif_typed_lsif_proto_rawDescGZIP(), []int{4}
}

func (x *Moniker) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Moniker) GetScheme() string {
	if x != nil {
		return x.Scheme
	}
	return ""
}

func (x *Moniker) GetUnique() Moniker_Unique {
	if x != nil {
		return x.Unique
	}
	return Moniker_UNIQUE_UNSPECIFIED
}

func (x *Moniker) GetMarkdownHover() []string {
	if x != nil {
		return x.MarkdownHover
	}
	return nil
}

func (x *Moniker) GetImplementationMonikers() []string {
	if x != nil {
		return x.ImplementationMonikers
	}
	return nil
}

func (x *Moniker) GetPackageId() string {
	if x != nil {
		return x.PackageId
	}
	return ""
}

func (x *Moniker) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

// MonikerOccurrence describes the appearance of a moniker at a given range in a
// source file.
type MonikerOccurrence struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Reference to the `Moniker.id` field.
	MonikerId string `protobuf:"bytes,1,opt,name=moniker_id,json=monikerId,proto3" json:"moniker_id,omitempty"`
	// The start/end range where this moniker appears.
	Range *Range `protobuf:"bytes,2,opt,name=range,proto3" json:"range,omitempty"`
	// Is this the location where the moniker is defined or referenced?
	Role MonikerOccurrence_Role `protobuf:"varint,3,opt,name=role,proto3,enum=lib.codeintel.lsif_typed.MonikerOccurrence_Role" json:"role,omitempty"`
	// (optional) Markdown-formatted documentation for this specific range.  If
	// empty, the `Moniker.markdown_hover` field is used instead. One example
	// where this field might be useful is when the moniker represents a generic
	// function (with abstract type parameters such as `List<T>`) and at this
	// occurrence we know the exact values (such as `List<String>`).
	MarkdownHover []string `protobuf:"bytes,4,rep,name=markdown_hover,json=markdownHover,proto3" json:"markdown_hover,omitempty"`
}

func (x *MonikerOccurrence) Reset() {
	*x = MonikerOccurrence{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_codeintel_lsif_typed_lsif_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MonikerOccurrence) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MonikerOccurrence) ProtoMessage() {}

func (x *MonikerOccurrence) ProtoReflect() protoreflect.Message {
	mi := &file_lib_codeintel_lsif_typed_lsif_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MonikerOccurrence.ProtoReflect.Descriptor instead.
func (*MonikerOccurrence) Descriptor() ([]byte, []int) {
	return file_lib_codeintel_lsif_typed_lsif_proto_rawDescGZIP(), []int{5}
}

func (x *MonikerOccurrence) GetMonikerId() string {
	if x != nil {
		return x.MonikerId
	}
	return ""
}

func (x *MonikerOccurrence) GetRange() *Range {
	if x != nil {
		return x.Range
	}
	return nil
}

func (x *MonikerOccurrence) GetRole() MonikerOccurrence_Role {
	if x != nil {
		return x.Role
	}
	return MonikerOccurrence_ROLE_UNSPECIFIED
}

func (x *MonikerOccurrence) GetMarkdownHover() []string {
	if x != nil {
		return x.MarkdownHover
	}
	return nil
}

// Document defines information about a particular source file.
type Document struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// URI-formatted absolute path of the source file on disk (example "file:///path/to/some/file/on/disk.ts").
	Uri string `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	// The moniker occurrences that appear in this file.
	Occurrences []*MonikerOccurrence `protobuf:"bytes,2,rep,name=occurrences,proto3" json:"occurrences,omitempty"`
}

func (x *Document) Reset() {
	*x = Document{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_codeintel_lsif_typed_lsif_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Document) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Document) ProtoMessage() {}

func (x *Document) ProtoReflect() protoreflect.Message {
	mi := &file_lib_codeintel_lsif_typed_lsif_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Document.ProtoReflect.Descriptor instead.
func (*Document) Descriptor() ([]byte, []int) {
	return file_lib_codeintel_lsif_typed_lsif_proto_rawDescGZIP(), []int{6}
}

func (x *Document) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *Document) GetOccurrences() []*MonikerOccurrence {
	if x != nil {
		return x.Occurrences
	}
	return nil
}

// Package defines a publishable artifact such as an npm package, Docker
// container, JVM dependency, or a Cargo crate.
type Package struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique identifier of this package that can be referenced from
	// `Moniker.package_id`.  This ID is not intended to be displayed to humans,
	// but it's recommended to use a human-readable format to aid with debugging.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The name of this package, for example "@types/react" or "com.google.guava:guava".
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// The version of this package, for example "0.1.0" or "2.1.5".
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	// The package manager, for example "npm", "maven" or "cargo".
	Manager string `protobuf:"bytes,4,opt,name=manager,proto3" json:"manager,omitempty"`
}

func (x *Package) Reset() {
	*x = Package{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_codeintel_lsif_typed_lsif_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Package) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Package) ProtoMessage() {}

func (x *Package) ProtoReflect() protoreflect.Message {
	mi := &file_lib_codeintel_lsif_typed_lsif_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Package.ProtoReflect.Descriptor instead.
func (*Package) Descriptor() ([]byte, []int) {
	return file_lib_codeintel_lsif_typed_lsif_proto_rawDescGZIP(), []int{7}
}

func (x *Package) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Package) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Package) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Package) GetManager() string {
	if x != nil {
		return x.Manager
	}
	return ""
}

var File_lib_codeintel_lsif_typed_lsif_proto protoreflect.FileDescriptor

var file_lib_codeintel_lsif_typed_lsif_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6c, 0x69, 0x62, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x69, 0x6e, 0x74, 0x65, 0x6c, 0x2f,
	0x6c, 0x73, 0x69, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x64, 0x2f, 0x6c, 0x73, 0x69, 0x66, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x6c, 0x69, 0x62, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x69,
	0x6e, 0x74, 0x65, 0x6c, 0x2e, 0x6c, 0x73, 0x69, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x64, 0x22,
	0x49, 0x0a, 0x0a, 0x4c, 0x73, 0x69, 0x66, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x3b, 0x0a,
	0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e,
	0x6c, 0x69, 0x62, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x69, 0x6e, 0x74, 0x65, 0x6c, 0x2e, 0x6c, 0x73,
	0x69, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x64, 0x2e, 0x4c, 0x73, 0x69, 0x66, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0xd4, 0x01, 0x0a, 0x09, 0x4c,
	0x73, 0x69, 0x66, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3d, 0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6c, 0x69, 0x62, 0x2e,
	0x63, 0x6f, 0x64, 0x65, 0x69, 0x6e, 0x74, 0x65, 0x6c, 0x2e, 0x6c, 0x73, 0x69, 0x66, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x64, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x07,
	0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x40, 0x0a, 0x08, 0x64, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6c, 0x69, 0x62, 0x2e,
	0x63, 0x6f, 0x64, 0x65, 0x69, 0x6e, 0x74, 0x65, 0x6c, 0x2e, 0x6c, 0x73, 0x69, 0x66, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x64, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52,
	0x08, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x3d, 0x0a, 0x07, 0x6d, 0x6f, 0x6e,
	0x69, 0x6b, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6c, 0x69, 0x62,
	0x2e, 0x63, 0x6f, 0x64, 0x65, 0x69, 0x6e, 0x74, 0x65, 0x6c, 0x2e, 0x6c, 0x73, 0x69, 0x66, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x64, 0x2e, 0x4d, 0x6f, 0x6e, 0x69, 0x6b, 0x65, 0x72, 0x48, 0x00, 0x52,
	0x07, 0x6d, 0x6f, 0x6e, 0x69, 0x6b, 0x65, 0x72, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x3c, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6c, 0x69, 0x6e,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x22,
	0x77, 0x0a, 0x05, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x38, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6c, 0x69, 0x62, 0x2e, 0x63, 0x6f,
	0x64, 0x65, 0x69, 0x6e, 0x74, 0x65, 0x6c, 0x2e, 0x6c, 0x73, 0x69, 0x66, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x64, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x12, 0x34, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x6c, 0x69, 0x62, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x69, 0x6e, 0x74, 0x65, 0x6c, 0x2e,
	0x6c, 0x73, 0x69, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x64, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x22, 0x8e, 0x03, 0x0a, 0x07, 0x4d, 0x6f, 0x6e,
	0x69, 0x6b, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x06,
	0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x6c,
	0x69, 0x62, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x69, 0x6e, 0x74, 0x65, 0x6c, 0x2e, 0x6c, 0x73, 0x69,
	0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x64, 0x2e, 0x4d, 0x6f, 0x6e, 0x69, 0x6b, 0x65, 0x72, 0x2e,
	0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x52, 0x06, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x12, 0x25,
	0x0a, 0x0e, 0x6d, 0x61, 0x72, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x5f, 0x68, 0x6f, 0x76, 0x65, 0x72,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x6d, 0x61, 0x72, 0x6b, 0x64, 0x6f, 0x77, 0x6e,
	0x48, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x37, 0x0a, 0x17, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x6b, 0x65, 0x72, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x16, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x6e, 0x69, 0x6b, 0x65, 0x72, 0x73, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e,
	0x64, 0x22, 0x85, 0x01, 0x0a, 0x06, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x12,
	0x55, 0x4e, 0x49, 0x51, 0x55, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x55, 0x4e, 0x49, 0x51, 0x55, 0x45, 0x5f, 0x44,
	0x4f, 0x43, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x4e, 0x49,
	0x51, 0x55, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x10, 0x02, 0x12, 0x14, 0x0a,
	0x10, 0x55, 0x4e, 0x49, 0x51, 0x55, 0x45, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x53, 0x50, 0x41, 0x43,
	0x45, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x49, 0x51, 0x55, 0x45, 0x5f, 0x53, 0x43,
	0x48, 0x45, 0x4d, 0x45, 0x10, 0x04, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x49, 0x51, 0x55, 0x45,
	0x5f, 0x47, 0x4c, 0x4f, 0x42, 0x41, 0x4c, 0x10, 0x05, 0x22, 0x9d, 0x02, 0x0a, 0x11, 0x4d, 0x6f,
	0x6e, 0x69, 0x6b, 0x65, 0x72, 0x4f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x6d, 0x6f, 0x6e, 0x69, 0x6b, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x6f, 0x6e, 0x69, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x12, 0x35,
	0x0a, 0x05, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x6c, 0x69, 0x62, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x69, 0x6e, 0x74, 0x65, 0x6c, 0x2e, 0x6c, 0x73,
	0x69, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x64, 0x2e, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x05,
	0x72, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x44, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e, 0x6c, 0x69, 0x62, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x69, 0x6e,
	0x74, 0x65, 0x6c, 0x2e, 0x6c, 0x73, 0x69, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x64, 0x2e, 0x4d,
	0x6f, 0x6e, 0x69, 0x6b, 0x65, 0x72, 0x4f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x6d,
	0x61, 0x72, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x5f, 0x68, 0x6f, 0x76, 0x65, 0x72, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0d, 0x6d, 0x61, 0x72, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x48, 0x6f, 0x76,
	0x65, 0x72, 0x22, 0x45, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x52, 0x4f,
	0x4c, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x13, 0x0a, 0x0f, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x49, 0x54,
	0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x52, 0x45,
	0x46, 0x45, 0x52, 0x45, 0x4e, 0x43, 0x45, 0x10, 0x02, 0x22, 0x6b, 0x0a, 0x08, 0x44, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x4d, 0x0a, 0x0b, 0x6f, 0x63, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x6c,
	0x69, 0x62, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x69, 0x6e, 0x74, 0x65, 0x6c, 0x2e, 0x6c, 0x73, 0x69,
	0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x64, 0x2e, 0x4d, 0x6f, 0x6e, 0x69, 0x6b, 0x65, 0x72, 0x4f,
	0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x0b, 0x6f, 0x63, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x22, 0x61, 0x0a, 0x07, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x67, 0x72,
	0x61, 0x70, 0x68, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f,
	0x6c, 0x69, 0x62, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x69, 0x6e, 0x74, 0x65, 0x6c, 0x2f, 0x6c, 0x73,
	0x69, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x64, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_lib_codeintel_lsif_typed_lsif_proto_rawDescOnce sync.Once
	file_lib_codeintel_lsif_typed_lsif_proto_rawDescData = file_lib_codeintel_lsif_typed_lsif_proto_rawDesc
)

func file_lib_codeintel_lsif_typed_lsif_proto_rawDescGZIP() []byte {
	file_lib_codeintel_lsif_typed_lsif_proto_rawDescOnce.Do(func() {
		file_lib_codeintel_lsif_typed_lsif_proto_rawDescData = protoimpl.X.CompressGZIP(file_lib_codeintel_lsif_typed_lsif_proto_rawDescData)
	})
	return file_lib_codeintel_lsif_typed_lsif_proto_rawDescData
}

var file_lib_codeintel_lsif_typed_lsif_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_lib_codeintel_lsif_typed_lsif_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_lib_codeintel_lsif_typed_lsif_proto_goTypes = []interface{}{
	(Moniker_Unique)(0),         // 0: lib.codeintel.lsif_typed.Moniker.Unique
	(MonikerOccurrence_Role)(0), // 1: lib.codeintel.lsif_typed.MonikerOccurrence.Role
	(*LsifValues)(nil),          // 2: lib.codeintel.lsif_typed.LsifValues
	(*LsifValue)(nil),           // 3: lib.codeintel.lsif_typed.LsifValue
	(*Position)(nil),            // 4: lib.codeintel.lsif_typed.Position
	(*Range)(nil),               // 5: lib.codeintel.lsif_typed.Range
	(*Moniker)(nil),             // 6: lib.codeintel.lsif_typed.Moniker
	(*MonikerOccurrence)(nil),   // 7: lib.codeintel.lsif_typed.MonikerOccurrence
	(*Document)(nil),            // 8: lib.codeintel.lsif_typed.Document
	(*Package)(nil),             // 9: lib.codeintel.lsif_typed.Package
}
var file_lib_codeintel_lsif_typed_lsif_proto_depIdxs = []int32{
	3,  // 0: lib.codeintel.lsif_typed.LsifValues.values:type_name -> lib.codeintel.lsif_typed.LsifValue
	9,  // 1: lib.codeintel.lsif_typed.LsifValue.package:type_name -> lib.codeintel.lsif_typed.Package
	8,  // 2: lib.codeintel.lsif_typed.LsifValue.document:type_name -> lib.codeintel.lsif_typed.Document
	6,  // 3: lib.codeintel.lsif_typed.LsifValue.moniker:type_name -> lib.codeintel.lsif_typed.Moniker
	4,  // 4: lib.codeintel.lsif_typed.Range.start:type_name -> lib.codeintel.lsif_typed.Position
	4,  // 5: lib.codeintel.lsif_typed.Range.end:type_name -> lib.codeintel.lsif_typed.Position
	0,  // 6: lib.codeintel.lsif_typed.Moniker.unique:type_name -> lib.codeintel.lsif_typed.Moniker.Unique
	5,  // 7: lib.codeintel.lsif_typed.MonikerOccurrence.range:type_name -> lib.codeintel.lsif_typed.Range
	1,  // 8: lib.codeintel.lsif_typed.MonikerOccurrence.role:type_name -> lib.codeintel.lsif_typed.MonikerOccurrence.Role
	7,  // 9: lib.codeintel.lsif_typed.Document.occurrences:type_name -> lib.codeintel.lsif_typed.MonikerOccurrence
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_lib_codeintel_lsif_typed_lsif_proto_init() }
func file_lib_codeintel_lsif_typed_lsif_proto_init() {
	if File_lib_codeintel_lsif_typed_lsif_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lib_codeintel_lsif_typed_lsif_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LsifValues); i {
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
		file_lib_codeintel_lsif_typed_lsif_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LsifValue); i {
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
		file_lib_codeintel_lsif_typed_lsif_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Position); i {
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
		file_lib_codeintel_lsif_typed_lsif_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Range); i {
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
		file_lib_codeintel_lsif_typed_lsif_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Moniker); i {
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
		file_lib_codeintel_lsif_typed_lsif_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MonikerOccurrence); i {
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
		file_lib_codeintel_lsif_typed_lsif_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Document); i {
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
		file_lib_codeintel_lsif_typed_lsif_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Package); i {
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
	file_lib_codeintel_lsif_typed_lsif_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*LsifValue_Package)(nil),
		(*LsifValue_Document)(nil),
		(*LsifValue_Moniker)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_lib_codeintel_lsif_typed_lsif_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_lib_codeintel_lsif_typed_lsif_proto_goTypes,
		DependencyIndexes: file_lib_codeintel_lsif_typed_lsif_proto_depIdxs,
		EnumInfos:         file_lib_codeintel_lsif_typed_lsif_proto_enumTypes,
		MessageInfos:      file_lib_codeintel_lsif_typed_lsif_proto_msgTypes,
	}.Build()
	File_lib_codeintel_lsif_typed_lsif_proto = out.File
	file_lib_codeintel_lsif_typed_lsif_proto_rawDesc = nil
	file_lib_codeintel_lsif_typed_lsif_proto_goTypes = nil
	file_lib_codeintel_lsif_typed_lsif_proto_depIdxs = nil
}
