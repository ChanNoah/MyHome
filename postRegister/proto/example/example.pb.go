// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/example/example.proto

package go_micro_srv_postRegister

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Request struct {
	Mobile               string   `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Pin                  string   `protobuf:"bytes,3,opt,name=pin,proto3" json:"pin,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *Request) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *Request) GetPin() string {
	if m != nil {
		return m.Pin
	}
	return ""
}

type Response struct {
	Errno                string   `protobuf:"bytes,1,opt,name=errno,proto3" json:"errno,omitempty"`
	Errmsg               string   `protobuf:"bytes,2,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
	SessionID            string   `protobuf:"bytes,3,opt,name=sessionID,proto3" json:"sessionID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetErrno() string {
	if m != nil {
		return m.Errno
	}
	return ""
}

func (m *Response) GetErrmsg() string {
	if m != nil {
		return m.Errmsg
	}
	return ""
}

func (m *Response) GetSessionID() string {
	if m != nil {
		return m.SessionID
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "go.micro.srv.postRegister.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.postRegister.Response")
}

func init() { proto.RegisterFile("proto/example/example.proto", fileDescriptor_097b3f5db5cf5789) }

var fileDescriptor_097b3f5db5cf5789 = []byte{
	// 217 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x3d, 0x4b, 0x04, 0x31,
	0x10, 0x86, 0x3d, 0x0f, 0xef, 0x63, 0xb0, 0x90, 0x41, 0x64, 0x3d, 0x2d, 0x24, 0x36, 0x56, 0x11,
	0xf4, 0x2f, 0x68, 0x61, 0xa5, 0xa4, 0x10, 0x2c, 0x6f, 0xdd, 0x61, 0x09, 0x6c, 0x32, 0x71, 0x26,
	0x7e, 0xfc, 0x7c, 0xd9, 0x6c, 0xfc, 0x68, 0xb4, 0x4a, 0x9e, 0x37, 0xe1, 0x4d, 0x9e, 0x81, 0x93,
	0x24, 0x9c, 0xf9, 0x92, 0x3e, 0xb6, 0x21, 0x0d, 0xf4, 0xb5, 0xda, 0x92, 0xe2, 0x71, 0xcf, 0x36,
	0xf8, 0x67, 0x61, 0xab, 0xf2, 0x66, 0x13, 0x6b, 0x76, 0xd4, 0x7b, 0xcd, 0x24, 0xe6, 0x1e, 0x96,
	0x8e, 0x5e, 0x5e, 0x49, 0x33, 0x1e, 0xc1, 0x22, 0x70, 0xeb, 0x07, 0x6a, 0x66, 0x67, 0xb3, 0x8b,
	0xb5, 0xab, 0x84, 0x1b, 0x58, 0xa5, 0xad, 0xea, 0x3b, 0x4b, 0xd7, 0xec, 0x96, 0x93, 0x6f, 0xc6,
	0x03, 0x98, 0x27, 0x1f, 0x9b, 0x79, 0x89, 0xc7, 0xad, 0x79, 0x84, 0x95, 0x23, 0x4d, 0x1c, 0x95,
	0xf0, 0x10, 0xf6, 0x48, 0x24, 0x72, 0x2d, 0x9c, 0x60, 0x7c, 0x87, 0x44, 0x82, 0xf6, 0xb5, 0xad,
	0x12, 0x9e, 0xc2, 0x5a, 0x49, 0xd5, 0x73, 0xbc, 0xbb, 0xa9, 0x8d, 0x3f, 0xc1, 0x55, 0x07, 0xcb,
	0xdb, 0x49, 0x0a, 0x9f, 0x60, 0xff, 0xe1, 0x97, 0x03, 0x1a, 0xfb, 0xa7, 0x9f, 0xad, 0x72, 0x9b,
	0xf3, 0x7f, 0xef, 0x4c, 0xff, 0x35, 0x3b, 0xed, 0xa2, 0x0c, 0xec, 0xfa, 0x33, 0x00, 0x00, 0xff,
	0xff, 0x3d, 0xdb, 0xfc, 0x24, 0x4f, 0x01, 0x00, 0x00,
}