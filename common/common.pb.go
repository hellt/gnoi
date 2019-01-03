// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/common.proto

package gnoi_common

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	types "github.com/openconfig/gnoi/types"
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

type RemoteDownload_Protocol int32

const (
	RemoteDownload_UNKNOWN RemoteDownload_Protocol = 0
	RemoteDownload_SFTP    RemoteDownload_Protocol = 1
	RemoteDownload_HTTP    RemoteDownload_Protocol = 2
	RemoteDownload_HTTPS   RemoteDownload_Protocol = 3
	RemoteDownload_SCP     RemoteDownload_Protocol = 4
)

var RemoteDownload_Protocol_name = map[int32]string{
	0: "UNKNOWN",
	1: "SFTP",
	2: "HTTP",
	3: "HTTPS",
	4: "SCP",
}

var RemoteDownload_Protocol_value = map[string]int32{
	"UNKNOWN": 0,
	"SFTP":    1,
	"HTTP":    2,
	"HTTPS":   3,
	"SCP":     4,
}

func (x RemoteDownload_Protocol) String() string {
	return proto.EnumName(RemoteDownload_Protocol_name, int32(x))
}

func (RemoteDownload_Protocol) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{0, 0}
}

// RemoteDownload defines the details for a device to initiate a file transfer
// from or to a remote location.
type RemoteDownload struct {
	// The path information containing where to download the data from or to.
	// For HTTP(S), this will be the URL (i.e. foo.com/file.tbz2).
	// For SFTP and SCP, this will be the address:/path/to/file
	// (i.e. host.foo.com:/bar/baz).
	Path                 string                  `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Protocol             RemoteDownload_Protocol `protobuf:"varint,2,opt,name=protocol,proto3,enum=gnoi.common.RemoteDownload_Protocol" json:"protocol,omitempty"`
	Credentials          *types.Credentials      `protobuf:"bytes,3,opt,name=credentials,proto3" json:"credentials,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *RemoteDownload) Reset()         { *m = RemoteDownload{} }
func (m *RemoteDownload) String() string { return proto.CompactTextString(m) }
func (*RemoteDownload) ProtoMessage()    {}
func (*RemoteDownload) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{0}
}

func (m *RemoteDownload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoteDownload.Unmarshal(m, b)
}
func (m *RemoteDownload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoteDownload.Marshal(b, m, deterministic)
}
func (m *RemoteDownload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoteDownload.Merge(m, src)
}
func (m *RemoteDownload) XXX_Size() int {
	return xxx_messageInfo_RemoteDownload.Size(m)
}
func (m *RemoteDownload) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoteDownload.DiscardUnknown(m)
}

var xxx_messageInfo_RemoteDownload proto.InternalMessageInfo

func (m *RemoteDownload) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *RemoteDownload) GetProtocol() RemoteDownload_Protocol {
	if m != nil {
		return m.Protocol
	}
	return RemoteDownload_UNKNOWN
}

func (m *RemoteDownload) GetCredentials() *types.Credentials {
	if m != nil {
		return m.Credentials
	}
	return nil
}

func init() {
	proto.RegisterEnum("gnoi.common.RemoteDownload_Protocol", RemoteDownload_Protocol_name, RemoteDownload_Protocol_value)
	proto.RegisterType((*RemoteDownload)(nil), "gnoi.common.RemoteDownload")
}

func init() { proto.RegisterFile("common/common.proto", fileDescriptor_8f954d82c0b891f6) }

var fileDescriptor_8f954d82c0b891f6 = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0xd3, 0x87, 0x50, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xdc, 0xe9, 0x79, 0xf9,
	0x99, 0x7a, 0x10, 0x21, 0x29, 0x9d, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0x10, 0x57, 0x3f, 0xbf,
	0x20, 0x35, 0x2f, 0x39, 0x3f, 0x2f, 0x2d, 0x33, 0x5d, 0x1f, 0xa4, 0x44, 0xbf, 0xa4, 0xb2, 0x20,
	0xb5, 0x18, 0x42, 0x42, 0xb4, 0x2a, 0x3d, 0x62, 0xe4, 0xe2, 0x0b, 0x4a, 0xcd, 0xcd, 0x2f, 0x49,
	0x75, 0xc9, 0x2f, 0xcf, 0xcb, 0xc9, 0x4f, 0x4c, 0x11, 0x12, 0xe2, 0x62, 0x29, 0x48, 0x2c, 0xc9,
	0x90, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x85, 0x1c, 0xb8, 0x38, 0xc0, 0xea, 0x93,
	0xf3, 0x73, 0x24, 0x98, 0x14, 0x18, 0x35, 0xf8, 0x8c, 0x54, 0xf4, 0x90, 0x2c, 0xd5, 0x43, 0x35,
	0x42, 0x2f, 0x00, 0xaa, 0x36, 0x08, 0xae, 0x4b, 0xc8, 0x92, 0x8b, 0x3b, 0xb9, 0x28, 0x35, 0x25,
	0x35, 0xaf, 0x24, 0x33, 0x31, 0xa7, 0x58, 0x82, 0x59, 0x81, 0x51, 0x83, 0xdb, 0x48, 0x1c, 0x62,
	0x08, 0xc4, 0x41, 0xce, 0x08, 0xe9, 0x20, 0x64, 0xb5, 0x4a, 0xf6, 0x5c, 0x1c, 0x30, 0x03, 0x85,
	0xb8, 0xb9, 0xd8, 0x43, 0xfd, 0xbc, 0xfd, 0xfc, 0xc3, 0xfd, 0x04, 0x18, 0x84, 0x38, 0xb8, 0x58,
	0x82, 0xdd, 0x42, 0x02, 0x04, 0x18, 0x41, 0x2c, 0x8f, 0x90, 0x90, 0x00, 0x01, 0x26, 0x21, 0x4e,
	0x2e, 0x56, 0x10, 0x2b, 0x58, 0x80, 0x59, 0x88, 0x9d, 0x8b, 0x39, 0xd8, 0x39, 0x40, 0x80, 0x25,
	0x89, 0x0d, 0xec, 0x0a, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6f, 0xdb, 0xfe, 0xf8, 0x3d,
	0x01, 0x00, 0x00,
}
