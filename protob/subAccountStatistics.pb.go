// Code generated by protoc-gen-go. DO NOT EDIT.
// source: subAccountStatistics.proto

package protob

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SubAccountStatisticsRequest struct {
	StartKey         *string   `protobuf:"bytes,1,req,name=startKey" json:"startKey,omitempty"`
	EndKey           *string   `protobuf:"bytes,2,req,name=endKey" json:"endKey,omitempty"`
	Params           []*Params `protobuf:"bytes,3,rep,name=params" json:"params,omitempty"`
	Column           *string   `protobuf:"bytes,4,req,name=column" json:"column,omitempty"`
	DefaultQualifier *string   `protobuf:"bytes,5,opt,name=defaultQualifier" json:"defaultQualifier,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *SubAccountStatisticsRequest) Reset()                    { *m = SubAccountStatisticsRequest{} }
func (m *SubAccountStatisticsRequest) String() string            { return proto.CompactTextString(m) }
func (*SubAccountStatisticsRequest) ProtoMessage()               {}
func (*SubAccountStatisticsRequest) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func (m *SubAccountStatisticsRequest) GetStartKey() string {
	if m != nil && m.StartKey != nil {
		return *m.StartKey
	}
	return ""
}

func (m *SubAccountStatisticsRequest) GetEndKey() string {
	if m != nil && m.EndKey != nil {
		return *m.EndKey
	}
	return ""
}

func (m *SubAccountStatisticsRequest) GetParams() []*Params {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *SubAccountStatisticsRequest) GetColumn() string {
	if m != nil && m.Column != nil {
		return *m.Column
	}
	return ""
}

func (m *SubAccountStatisticsRequest) GetDefaultQualifier() string {
	if m != nil && m.DefaultQualifier != nil {
		return *m.DefaultQualifier
	}
	return ""
}

type SubAccountStatisticsResponse struct {
	RespParams       []*RespParams `protobuf:"bytes,1,rep,name=respParams" json:"respParams,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *SubAccountStatisticsResponse) Reset()                    { *m = SubAccountStatisticsResponse{} }
func (m *SubAccountStatisticsResponse) String() string            { return proto.CompactTextString(m) }
func (*SubAccountStatisticsResponse) ProtoMessage()               {}
func (*SubAccountStatisticsResponse) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{1} }

func (m *SubAccountStatisticsResponse) GetRespParams() []*RespParams {
	if m != nil {
		return m.RespParams
	}
	return nil
}

func init() {
	proto.RegisterType((*SubAccountStatisticsRequest)(nil), "protob.SubAccountStatisticsRequest")
	proto.RegisterType((*SubAccountStatisticsResponse)(nil), "protob.SubAccountStatisticsResponse")
}

func init() { proto.RegisterFile("subAccountStatistics.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x50, 0xdf, 0x4a, 0xf3, 0x30,
	0x14, 0x27, 0xdb, 0xf7, 0x0d, 0x77, 0x14, 0x19, 0x41, 0xa4, 0x54, 0xc1, 0x31, 0x45, 0x8a, 0x17,
	0xb9, 0xd8, 0x1b, 0xb8, 0x2b, 0x41, 0x10, 0xcd, 0x9e, 0x20, 0x4d, 0xcf, 0xa0, 0xd0, 0x25, 0x31,
	0x27, 0x51, 0xf4, 0xda, 0x0b, 0x1f, 0xc3, 0x17, 0xf1, 0xdd, 0x64, 0xcb, 0x36, 0x07, 0x96, 0x5e,
	0x95, 0xf3, 0xfb, 0xdf, 0x40, 0x4e, 0xb1, 0xbc, 0xd5, 0xda, 0x46, 0x13, 0xe6, 0x41, 0x85, 0x9a,
	0x42, 0xad, 0x49, 0x38, 0x6f, 0x83, 0xe5, 0x83, 0xf5, 0xa7, 0xcc, 0x8f, 0x9c, 0xf2, 0x6a, 0xb9,
	0x41, 0xf3, 0x91, 0x47, 0x72, 0x8f, 0x7b, 0xc8, 0xe4, 0x9b, 0xc1, 0xd9, 0xbc, 0x25, 0x46, 0xe2,
	0x73, 0x44, 0x0a, 0x3c, 0x87, 0x03, 0x0a, 0xca, 0x87, 0x7b, 0x7c, 0xcb, 0xd8, 0xb8, 0x57, 0x0c,
	0xe5, 0xee, 0xe6, 0xa7, 0x30, 0x40, 0x53, 0xad, 0x98, 0xde, 0x9a, 0xd9, 0x5c, 0xfc, 0x1a, 0x06,
	0xa9, 0x35, 0xeb, 0x8f, 0xfb, 0xc5, 0xe1, 0xf4, 0x38, 0x75, 0x95, 0x22, 0x35, 0xcb, 0x0d, 0xbb,
	0xf2, 0x6b, 0xdb, 0xc4, 0xa5, 0xc9, 0xfe, 0x25, 0x7f, 0xba, 0xf8, 0x0d, 0x8c, 0x2a, 0x5c, 0xa8,
	0xd8, 0x84, 0xa7, 0xa8, 0x9a, 0x7a, 0x51, 0xa3, 0xcf, 0xfe, 0x8f, 0x59, 0x31, 0x94, 0x7f, 0xf0,
	0x89, 0x84, 0xf3, 0xf6, 0xf9, 0xe4, 0xac, 0x21, 0xe4, 0x53, 0x80, 0xdf, 0x7f, 0xce, 0xd8, 0x7a,
	0x0f, 0xdf, 0xee, 0x91, 0x3b, 0x46, 0xee, 0xa9, 0xa6, 0x1f, 0x0c, 0x4e, 0xda, 0x42, 0x79, 0x03,
	0x17, 0x84, 0xa6, 0xea, 0x7a, 0xaf, 0xcb, 0x6d, 0x76, 0x87, 0x28, 0xbf, 0xea, 0x16, 0xa5, 0xe9,
	0xb3, 0x07, 0x28, 0xb4, 0x5d, 0x8a, 0xf2, 0xb5, 0x7a, 0x17, 0xda, 0x3a, 0x6f, 0x35, 0x12, 0x59,
	0x2f, 0xd0, 0x54, 0xce, 0xd6, 0x26, 0x08, 0x8a, 0xa5, 0x4a, 0xee, 0x59, 0xde, 0x96, 0x34, 0x47,
	0xff, 0x82, 0xfe, 0x8e, 0x7d, 0x32, 0xf6, 0xc5, 0xd8, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x08,
	0x4f, 0x64, 0x1f, 0x2f, 0x02, 0x00, 0x00,
}
