// Code generated by protoc-gen-go. DO NOT EDIT.
// source: subRespParams.proto

package protob

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SubRespParams struct {
	Key              *string `protobuf:"bytes,4,req,name=key" json:"key,omitempty"`
	Value            *string `protobuf:"bytes,5,req,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SubRespParams) Reset()                    { *m = SubRespParams{} }
func (m *SubRespParams) String() string            { return proto.CompactTextString(m) }
func (*SubRespParams) ProtoMessage()               {}
func (*SubRespParams) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func (m *SubRespParams) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *SubRespParams) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*SubRespParams)(nil), "protob.SubRespParams")
}

func init() { proto.RegisterFile("subRespParams.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 136 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x2e, 0x4d, 0x0a,
	0x4a, 0x2d, 0x2e, 0x08, 0x48, 0x2c, 0x4a, 0xcc, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x03, 0x53, 0x49, 0x4a, 0xe6, 0x5c, 0xbc, 0xc1, 0xc8, 0xd2, 0x42, 0x02, 0x5c, 0xcc, 0xd9,
	0xa9, 0x95, 0x12, 0x2c, 0x0a, 0x4c, 0x1a, 0x9c, 0x41, 0x20, 0xa6, 0x90, 0x08, 0x17, 0x6b, 0x59,
	0x62, 0x4e, 0x69, 0xaa, 0x04, 0x2b, 0x58, 0x0c, 0xc2, 0x71, 0xb2, 0xe6, 0x92, 0x4d, 0xce, 0xcf,
	0xd5, 0x4b, 0x2a, 0x4f, 0xa9, 0xd2, 0x4b, 0xce, 0x2f, 0x28, 0xca, 0x4f, 0x4e, 0x2d, 0x2e, 0xce,
	0x2f, 0xd2, 0x4b, 0xcd, 0x4b, 0x29, 0xc8, 0xcf, 0xcc, 0x2b, 0x71, 0x12, 0x42, 0x31, 0x37, 0x00,
	0x64, 0x9d, 0x07, 0x63, 0x07, 0x23, 0xe3, 0x02, 0x46, 0x46, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x96, 0xbe, 0x19, 0xe6, 0x93, 0x00, 0x00, 0x00,
}