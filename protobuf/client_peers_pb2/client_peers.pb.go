// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sawtooth-go-sdk/protobuf/client_peers_pb2/client_peers.proto

/*
Package client_peer is a generated protocol buffer package.

It is generated from these files:
	sawtooth-go-sdk/protobuf/client_peers_pb2/client_peers.proto

It has these top-level messages:
	ClientPeersGetRequest
	ClientPeersGetResponse
*/
package client_peer

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

type ClientPeersGetResponse_Status int32

const (
	ClientPeersGetResponse_STATUS_UNSET ClientPeersGetResponse_Status = 0
	ClientPeersGetResponse_OK           ClientPeersGetResponse_Status = 1
	ClientPeersGetResponse_ERROR        ClientPeersGetResponse_Status = 2
)

var ClientPeersGetResponse_Status_name = map[int32]string{
	0: "STATUS_UNSET",
	1: "OK",
	2: "ERROR",
}
var ClientPeersGetResponse_Status_value = map[string]int32{
	"STATUS_UNSET": 0,
	"OK":           1,
	"ERROR":        2,
}

func (x ClientPeersGetResponse_Status) String() string {
	return proto.EnumName(ClientPeersGetResponse_Status_name, int32(x))
}
func (ClientPeersGetResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 0}
}

type ClientPeersGetRequest struct {
}

func (m *ClientPeersGetRequest) Reset()                    { *m = ClientPeersGetRequest{} }
func (m *ClientPeersGetRequest) String() string            { return proto.CompactTextString(m) }
func (*ClientPeersGetRequest) ProtoMessage()               {}
func (*ClientPeersGetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ClientPeersGetResponse struct {
	Status ClientPeersGetResponse_Status `protobuf:"varint,1,opt,name=status,enum=ClientPeersGetResponse_Status" json:"status,omitempty"`
	Peers  []string                      `protobuf:"bytes,2,rep,name=peers" json:"peers,omitempty"`
}

func (m *ClientPeersGetResponse) Reset()                    { *m = ClientPeersGetResponse{} }
func (m *ClientPeersGetResponse) String() string            { return proto.CompactTextString(m) }
func (*ClientPeersGetResponse) ProtoMessage()               {}
func (*ClientPeersGetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ClientPeersGetResponse) GetStatus() ClientPeersGetResponse_Status {
	if m != nil {
		return m.Status
	}
	return ClientPeersGetResponse_STATUS_UNSET
}

func (m *ClientPeersGetResponse) GetPeers() []string {
	if m != nil {
		return m.Peers
	}
	return nil
}

func init() {
	proto.RegisterType((*ClientPeersGetRequest)(nil), "ClientPeersGetRequest")
	proto.RegisterType((*ClientPeersGetResponse)(nil), "ClientPeersGetResponse")
	proto.RegisterEnum("ClientPeersGetResponse_Status", ClientPeersGetResponse_Status_name, ClientPeersGetResponse_Status_value)
}

func init() {
	proto.RegisterFile("sawtooth-go-sdk/protobuf/client_peers_pb2/client_peers.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xb2, 0x2c, 0x4e, 0x2c, 0x2f,
	0xc9, 0xcf, 0x2f, 0xc9, 0x88, 0x2f, 0x4e, 0xc9, 0xd6, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x2a,
	0x4d, 0xd3, 0x4f, 0xce, 0xc9, 0x4c, 0xcd, 0x2b, 0x89, 0x2f, 0x48, 0x4d, 0x2d, 0x2a, 0x8e, 0x2f,
	0x48, 0x32, 0x42, 0x11, 0xd0, 0x03, 0x2b, 0x53, 0x12, 0xe7, 0x12, 0x75, 0x06, 0x8b, 0x06, 0x80,
	0x04, 0xdd, 0x53, 0x4b, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x94, 0xa6, 0x32, 0x72, 0x89,
	0xa1, 0xcb, 0x14, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0x99, 0x71, 0xb1, 0x15, 0x97, 0x24, 0x96,
	0x94, 0x16, 0x4b, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x19, 0xc9, 0xe9, 0x61, 0x57, 0xa8, 0x17, 0x0c,
	0x56, 0x15, 0x04, 0x55, 0x2d, 0x24, 0xc2, 0xc5, 0x0a, 0xb6, 0x5a, 0x82, 0x49, 0x81, 0x59, 0x83,
	0x33, 0x08, 0xc2, 0x51, 0xd2, 0xe5, 0x62, 0x83, 0xa8, 0x13, 0x12, 0xe0, 0xe2, 0x09, 0x0e, 0x71,
	0x0c, 0x09, 0x0d, 0x8e, 0x0f, 0xf5, 0x0b, 0x76, 0x0d, 0x11, 0x60, 0x10, 0x62, 0xe3, 0x62, 0xf2,
	0xf7, 0x16, 0x60, 0x14, 0xe2, 0xe4, 0x62, 0x75, 0x0d, 0x0a, 0xf2, 0x0f, 0x12, 0x60, 0x72, 0x52,
	0xe3, 0x12, 0x85, 0xf9, 0x56, 0x0f, 0xe8, 0x5b, 0x3d, 0x98, 0x6f, 0x03, 0x18, 0xa3, 0xb8, 0x91,
	0xfc, 0x97, 0xc4, 0x06, 0x96, 0x30, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x7e, 0x44, 0x31, 0xda,
	0x1c, 0x01, 0x00, 0x00,
}