// Code generated by protoc-gen-go.
// source: v1.proto
// DO NOT EDIT!

/*
Package v1 is a generated protocol buffer package.

It is generated from these files:
	v1.proto

It has these top-level messages:
	Type
	FunctionRef
	MultiEdge
*/
package v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Kind is mostly identical to reflect.TypeKind, expect we handle certain
// types specially, such as "error".
type Type_Kind int32

const (
	Type_INVALID Type_Kind = 0
	// Primitive.
	Type_BOOL   Type_Kind = 1
	Type_INT    Type_Kind = 2
	Type_INT8   Type_Kind = 3
	Type_INT16  Type_Kind = 4
	Type_INT32  Type_Kind = 5
	Type_INT64  Type_Kind = 6
	Type_UINT   Type_Kind = 7
	Type_UINT8  Type_Kind = 8
	Type_UINT16 Type_Kind = 9
	Type_UINT32 Type_Kind = 10
	Type_UINT64 Type_Kind = 11
	Type_STRING Type_Kind = 12
	// Aggregate.
	Type_SLICE  Type_Kind = 20
	Type_STRUCT Type_Kind = 21
	// Special.
	Type_FUNC  Type_Kind = 30
	Type_CHAN  Type_Kind = 31
	Type_ERROR Type_Kind = 32
)

var Type_Kind_name = map[int32]string{
	0:  "INVALID",
	1:  "BOOL",
	2:  "INT",
	3:  "INT8",
	4:  "INT16",
	5:  "INT32",
	6:  "INT64",
	7:  "UINT",
	8:  "UINT8",
	9:  "UINT16",
	10: "UINT32",
	11: "UINT64",
	12: "STRING",
	20: "SLICE",
	21: "STRUCT",
	30: "FUNC",
	31: "CHAN",
	32: "ERROR",
}
var Type_Kind_value = map[string]int32{
	"INVALID": 0,
	"BOOL":    1,
	"INT":     2,
	"INT8":    3,
	"INT16":   4,
	"INT32":   5,
	"INT64":   6,
	"UINT":    7,
	"UINT8":   8,
	"UINT16":  9,
	"UINT32":  10,
	"UINT64":  11,
	"STRING":  12,
	"SLICE":   20,
	"STRUCT":  21,
	"FUNC":    30,
	"CHAN":    31,
	"ERROR":   32,
}

func (x Type_Kind) String() string {
	return proto.EnumName(Type_Kind_name, int32(x))
}
func (Type_Kind) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

// ChanDir matches reflect.ChanDir.
type Type_ChanDir int32

const (
	Type_RECV Type_ChanDir = 0
	Type_SEND Type_ChanDir = 1
	Type_BOTH Type_ChanDir = 2
)

var Type_ChanDir_name = map[int32]string{
	0: "RECV",
	1: "SEND",
	2: "BOTH",
}
var Type_ChanDir_value = map[string]int32{
	"RECV": 0,
	"SEND": 1,
	"BOTH": 2,
}

func (x Type_ChanDir) String() string {
	return proto.EnumName(Type_ChanDir_name, int32(x))
}
func (Type_ChanDir) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

// Type represents a serializable reflect.Type.
type Type struct {
	// (Required) Type kind.
	Kind Type_Kind `protobuf:"varint,1,opt,name=kind,enum=v1.Type_Kind" json:"kind,omitempty"`
	// (Optional) Element type (if SLICE or CHAN)
	Element *Type `protobuf:"bytes,2,opt,name=element" json:"element,omitempty"`
	// (Optional) Fields (if STRUCT).
	Fields []*Type_StructField `protobuf:"bytes,3,rep,name=fields" json:"fields,omitempty"`
	// (Optional) Parameter types (if FUNC).
	ParameterTypes []*Type `protobuf:"bytes,4,rep,name=parameter_types,json=parameterTypes" json:"parameter_types,omitempty"`
	// (Optional) Return types (if FUNC).
	ReturnTypes []*Type `protobuf:"bytes,5,rep,name=return_types,json=returnTypes" json:"return_types,omitempty"`
	// (Optional) Is variadic (if FUNC).
	IsVariadic bool `protobuf:"varint,6,opt,name=is_variadic,json=isVariadic" json:"is_variadic,omitempty"`
	// (Optional) Channel direction (if CHAN).
	ChanDir Type_ChanDir `protobuf:"varint,7,opt,name=chan_dir,json=chanDir,enum=v1.Type_ChanDir" json:"chan_dir,omitempty"`
}

func (m *Type) Reset()                    { *m = Type{} }
func (m *Type) String() string            { return proto.CompactTextString(m) }
func (*Type) ProtoMessage()               {}
func (*Type) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Type) GetKind() Type_Kind {
	if m != nil {
		return m.Kind
	}
	return Type_INVALID
}

func (m *Type) GetElement() *Type {
	if m != nil {
		return m.Element
	}
	return nil
}

func (m *Type) GetFields() []*Type_StructField {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *Type) GetParameterTypes() []*Type {
	if m != nil {
		return m.ParameterTypes
	}
	return nil
}

func (m *Type) GetReturnTypes() []*Type {
	if m != nil {
		return m.ReturnTypes
	}
	return nil
}

func (m *Type) GetIsVariadic() bool {
	if m != nil {
		return m.IsVariadic
	}
	return false
}

func (m *Type) GetChanDir() Type_ChanDir {
	if m != nil {
		return m.ChanDir
	}
	return Type_RECV
}

// StructField matches reflect.StructField.
type Type_StructField struct {
	Name      string  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	PkgPath   string  `protobuf:"bytes,2,opt,name=pkg_path,json=pkgPath" json:"pkg_path,omitempty"`
	Type      *Type   `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	Tag       string  `protobuf:"bytes,4,opt,name=tag" json:"tag,omitempty"`
	Offset    int64   `protobuf:"varint,5,opt,name=offset" json:"offset,omitempty"`
	Index     []int32 `protobuf:"varint,6,rep,packed,name=index" json:"index,omitempty"`
	Anonymous bool    `protobuf:"varint,7,opt,name=anonymous" json:"anonymous,omitempty"`
}

func (m *Type_StructField) Reset()                    { *m = Type_StructField{} }
func (m *Type_StructField) String() string            { return proto.CompactTextString(m) }
func (*Type_StructField) ProtoMessage()               {}
func (*Type_StructField) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *Type_StructField) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Type_StructField) GetPkgPath() string {
	if m != nil {
		return m.PkgPath
	}
	return ""
}

func (m *Type_StructField) GetType() *Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *Type_StructField) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *Type_StructField) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *Type_StructField) GetIndex() []int32 {
	if m != nil {
		return m.Index
	}
	return nil
}

func (m *Type_StructField) GetAnonymous() bool {
	if m != nil {
		return m.Anonymous
	}
	return false
}

// FunctionRef represents a serialized function reference. The
// implementation is notably not serialized and must be present (and
// somehow discoverable from the symbol name) on the decoding side.
type FunctionRef struct {
	// (Required) Symbol name of function.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// (Required) Function type.
	Type *Type `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
}

func (m *FunctionRef) Reset()                    { *m = FunctionRef{} }
func (m *FunctionRef) String() string            { return proto.CompactTextString(m) }
func (*FunctionRef) ProtoMessage()               {}
func (*FunctionRef) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *FunctionRef) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FunctionRef) GetType() *Type {
	if m != nil {
		return m.Type
	}
	return nil
}

// MultiEdge prepresents a serialized MultiEdge.
type MultiEdge struct {
	// UserFn
	UserFn *FunctionRef `protobuf:"bytes,1,opt,name=user_fn,json=userFn" json:"user_fn,omitempty"`
	// (Optional) JSON-serialized data.
	Data     string                `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
	Inbound  []*MultiEdge_Inbound  `protobuf:"bytes,3,rep,name=inbound" json:"inbound,omitempty"`
	Outbound []*MultiEdge_Outbound `protobuf:"bytes,4,rep,name=outbound" json:"outbound,omitempty"`
}

func (m *MultiEdge) Reset()                    { *m = MultiEdge{} }
func (m *MultiEdge) String() string            { return proto.CompactTextString(m) }
func (*MultiEdge) ProtoMessage()               {}
func (*MultiEdge) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *MultiEdge) GetUserFn() *FunctionRef {
	if m != nil {
		return m.UserFn
	}
	return nil
}

func (m *MultiEdge) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *MultiEdge) GetInbound() []*MultiEdge_Inbound {
	if m != nil {
		return m.Inbound
	}
	return nil
}

func (m *MultiEdge) GetOutbound() []*MultiEdge_Outbound {
	if m != nil {
		return m.Outbound
	}
	return nil
}

type MultiEdge_Inbound struct {
	Type *Type `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
}

func (m *MultiEdge_Inbound) Reset()                    { *m = MultiEdge_Inbound{} }
func (m *MultiEdge_Inbound) String() string            { return proto.CompactTextString(m) }
func (*MultiEdge_Inbound) ProtoMessage()               {}
func (*MultiEdge_Inbound) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

func (m *MultiEdge_Inbound) GetType() *Type {
	if m != nil {
		return m.Type
	}
	return nil
}

type MultiEdge_Outbound struct {
	Type *Type `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
}

func (m *MultiEdge_Outbound) Reset()                    { *m = MultiEdge_Outbound{} }
func (m *MultiEdge_Outbound) String() string            { return proto.CompactTextString(m) }
func (*MultiEdge_Outbound) ProtoMessage()               {}
func (*MultiEdge_Outbound) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 1} }

func (m *MultiEdge_Outbound) GetType() *Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func init() {
	proto.RegisterType((*Type)(nil), "v1.Type")
	proto.RegisterType((*Type_StructField)(nil), "v1.Type.StructField")
	proto.RegisterType((*FunctionRef)(nil), "v1.FunctionRef")
	proto.RegisterType((*MultiEdge)(nil), "v1.MultiEdge")
	proto.RegisterType((*MultiEdge_Inbound)(nil), "v1.MultiEdge.Inbound")
	proto.RegisterType((*MultiEdge_Outbound)(nil), "v1.MultiEdge.Outbound")
	proto.RegisterEnum("v1.Type_Kind", Type_Kind_name, Type_Kind_value)
	proto.RegisterEnum("v1.Type_ChanDir", Type_ChanDir_name, Type_ChanDir_value)
}

func init() { proto.RegisterFile("v1.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 623 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x54, 0xdd, 0x4e, 0xdb, 0x30,
	0x18, 0x5d, 0x9a, 0xdf, 0x7e, 0x61, 0x60, 0x59, 0x80, 0x32, 0x84, 0x46, 0xd6, 0x1b, 0x22, 0x31,
	0x75, 0x6a, 0x40, 0x68, 0x77, 0x13, 0x2b, 0xed, 0x88, 0xc6, 0xd2, 0xc9, 0x4d, 0xb9, 0xad, 0x42,
	0xe3, 0xb6, 0x11, 0xd4, 0x89, 0x12, 0xa7, 0x1a, 0x97, 0x7b, 0xab, 0x3d, 0xc4, 0x1e, 0x61, 0x0f,
	0x33, 0x39, 0x3f, 0x85, 0x21, 0xb4, 0xbb, 0xe3, 0xef, 0x9c, 0xf3, 0xd9, 0xc7, 0x5f, 0x1c, 0x30,
	0xd6, 0xbd, 0x6e, 0x9a, 0x25, 0x3c, 0xc1, 0xad, 0x75, 0xaf, 0xf3, 0x47, 0x05, 0x25, 0x78, 0x48,
	0x29, 0x7e, 0x07, 0xca, 0x5d, 0xcc, 0x22, 0x4b, 0xb2, 0x25, 0x67, 0xdb, 0x7d, 0xdd, 0x5d, 0xf7,
	0xba, 0xa2, 0xde, 0xfd, 0x1a, 0xb3, 0x88, 0x94, 0x14, 0xee, 0x80, 0x4e, 0xef, 0xe9, 0x8a, 0x32,
	0x6e, 0xb5, 0x6c, 0xc9, 0x31, 0x5d, 0xa3, 0x51, 0x91, 0x86, 0xc0, 0xef, 0x41, 0x9b, 0xc7, 0xf4,
	0x3e, 0xca, 0x2d, 0xd9, 0x96, 0x1d, 0xd3, 0xdd, 0xdd, 0x34, 0x1a, 0xf3, 0xac, 0x98, 0xf1, 0xa1,
	0x20, 0x49, 0xad, 0xc1, 0x3d, 0xd8, 0x49, 0xc3, 0x2c, 0x5c, 0x51, 0x4e, 0xb3, 0x29, 0x7f, 0x48,
	0x69, 0x6e, 0x29, 0xa5, 0xed, 0xb1, 0xf3, 0xf6, 0x46, 0x20, 0x96, 0x39, 0x3e, 0x81, 0xad, 0x8c,
	0xf2, 0x22, 0x63, 0xb5, 0x5e, 0x7d, 0xa6, 0x37, 0x2b, 0xb6, 0x12, 0x1f, 0x81, 0x19, 0xe7, 0xd3,
	0x75, 0x98, 0xc5, 0x61, 0x14, 0xcf, 0x2c, 0xcd, 0x96, 0x1c, 0x83, 0x40, 0x9c, 0xdf, 0xd4, 0x15,
	0x7c, 0x02, 0xc6, 0x6c, 0x19, 0xb2, 0x69, 0x14, 0x67, 0x96, 0x5e, 0x26, 0x47, 0x9b, 0x03, 0xf7,
	0x97, 0x21, 0xbb, 0x8c, 0x33, 0xa2, 0xcf, 0x2a, 0x70, 0xf0, 0x4b, 0x02, 0xf3, 0x49, 0x0a, 0x8c,
	0x41, 0x61, 0xe1, 0x8a, 0x96, 0x57, 0xd6, 0x26, 0x25, 0xc6, 0x6f, 0xc0, 0x48, 0xef, 0x16, 0xd3,
	0x34, 0xe4, 0xcb, 0xf2, 0x92, 0xda, 0x44, 0x4f, 0xef, 0x16, 0xdf, 0x43, 0xbe, 0xc4, 0x87, 0xa0,
	0x88, 0x23, 0x5b, 0xf2, 0xb3, 0xbb, 0x2b, 0xab, 0x18, 0x81, 0xcc, 0xc3, 0x85, 0xa5, 0x94, 0x1e,
	0x01, 0xf1, 0x3e, 0x68, 0xc9, 0x7c, 0x9e, 0x53, 0x6e, 0xa9, 0xb6, 0xe4, 0xc8, 0xa4, 0x5e, 0xe1,
	0x5d, 0x50, 0x63, 0x16, 0xd1, 0x1f, 0x96, 0x66, 0xcb, 0x8e, 0x4a, 0xaa, 0x05, 0x3e, 0x84, 0x76,
	0xc8, 0x12, 0xf6, 0xb0, 0x4a, 0x8a, 0xbc, 0x8c, 0x62, 0x90, 0xc7, 0x42, 0xe7, 0xb7, 0x04, 0x8a,
	0x98, 0x24, 0x36, 0x41, 0xf7, 0xfc, 0x9b, 0x8b, 0x6b, 0xef, 0x12, 0xbd, 0xc2, 0x06, 0x28, 0x9f,
	0x47, 0xa3, 0x6b, 0x24, 0x61, 0x1d, 0x64, 0xcf, 0x0f, 0x50, 0x4b, 0x94, 0x3c, 0x3f, 0xf8, 0x88,
	0x64, 0xdc, 0x06, 0xd5, 0xf3, 0x83, 0xde, 0x39, 0x52, 0x6a, 0x78, 0xea, 0x22, 0xb5, 0x86, 0xe7,
	0x67, 0x48, 0x13, 0xd2, 0x89, 0x30, 0xe9, 0xa2, 0x38, 0x29, 0x5d, 0x06, 0x06, 0xd0, 0x26, 0x95,
	0xad, 0xdd, 0xe0, 0x53, 0x17, 0x41, 0x83, 0xcf, 0xcf, 0x90, 0x29, 0xf0, 0x38, 0x20, 0x9e, 0xff,
	0x05, 0x6d, 0x09, 0xeb, 0xf8, 0xda, 0xeb, 0x0f, 0xd0, 0x6e, 0x5d, 0x9e, 0xf4, 0x03, 0xb4, 0x27,
	0x7a, 0x0f, 0x27, 0x7e, 0x1f, 0xbd, 0x15, 0xa8, 0x7f, 0x75, 0xe1, 0xa3, 0x23, 0x21, 0x1d, 0x10,
	0x32, 0x22, 0xc8, 0xee, 0x1c, 0x83, 0x5e, 0x4f, 0x47, 0xf0, 0x64, 0xd0, 0xbf, 0xa9, 0xd2, 0x8c,
	0x07, 0xfe, 0x25, 0x92, 0xaa, 0x5c, 0xc1, 0x15, 0x6a, 0x75, 0x3e, 0x81, 0x39, 0x2c, 0xd8, 0x8c,
	0xc7, 0x09, 0x23, 0x74, 0xfe, 0xe2, 0xc4, 0x9a, 0xb1, 0xb4, 0x5e, 0x1a, 0x4b, 0xe7, 0x67, 0x0b,
	0xda, 0xdf, 0x8a, 0x7b, 0x1e, 0x0f, 0xa2, 0x05, 0xc5, 0x0e, 0xe8, 0x45, 0x4e, 0xb3, 0xe9, 0x9c,
	0x95, 0x2d, 0x4c, 0x77, 0x47, 0xc8, 0x9f, 0xec, 0x40, 0x34, 0xc1, 0x0f, 0x99, 0xd8, 0x29, 0x0a,
	0x79, 0x58, 0x7f, 0x03, 0x25, 0xc6, 0x1f, 0x40, 0x8f, 0xd9, 0x6d, 0x52, 0xb0, 0xa8, 0x7e, 0x1c,
	0x7b, 0xc2, 0xbd, 0xe9, 0xde, 0xf5, 0x2a, 0x92, 0x34, 0x2a, 0xec, 0x82, 0x91, 0x14, 0xbc, 0x72,
	0x54, 0xef, 0x62, 0xff, 0x5f, 0xc7, 0xa8, 0x66, 0xc9, 0x46, 0x77, 0x70, 0x0c, 0x7a, 0xdd, 0x67,
	0x93, 0x4c, 0x7a, 0x29, 0xd9, 0x81, 0x03, 0x46, 0x63, 0xff, 0xbf, 0xf2, 0x56, 0x2b, 0x7f, 0x17,
	0xa7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x5a, 0xb0, 0x28, 0x8c, 0x3a, 0x04, 0x00, 0x00,
}