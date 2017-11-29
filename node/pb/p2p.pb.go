// Code generated by protoc-gen-go. DO NOT EDIT.
// source: p2p.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	p2p.proto

It has these top-level messages:
	MessageData
	NodeInfo
	PingRequest
	PingResponse
	EchoRequest
	EchoResponse
*/
package pb

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

// designed to be shared between all app protocols
type MessageData struct {
	// shared between all requests
	ClientVersion string `protobuf:"bytes,1,opt,name=clientVersion" json:"clientVersion,omitempty"`
	Timestamp     int64  `protobuf:"varint,2,opt,name=timestamp" json:"timestamp,omitempty"`
	Id            string `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
	Gossip        bool   `protobuf:"varint,4,opt,name=gossip" json:"gossip,omitempty"`
	NodeId        string `protobuf:"bytes,5,opt,name=nodeId" json:"nodeId,omitempty"`
	NodePubKey    []byte `protobuf:"bytes,6,opt,name=nodePubKey,proto3" json:"nodePubKey,omitempty"`
	Sign          string `protobuf:"bytes,7,opt,name=sign" json:"sign,omitempty"`
}

func (m *MessageData) Reset()                    { *m = MessageData{} }
func (m *MessageData) String() string            { return proto.CompactTextString(m) }
func (*MessageData) ProtoMessage()               {}
func (*MessageData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *MessageData) GetClientVersion() string {
	if m != nil {
		return m.ClientVersion
	}
	return ""
}

func (m *MessageData) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *MessageData) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MessageData) GetGossip() bool {
	if m != nil {
		return m.Gossip
	}
	return false
}

func (m *MessageData) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *MessageData) GetNodePubKey() []byte {
	if m != nil {
		return m.NodePubKey
	}
	return nil
}

func (m *MessageData) GetSign() string {
	if m != nil {
		return m.Sign
	}
	return ""
}

// minimal remote node info - sufficient to connect to the node
type NodeInfo struct {
	Id      string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Address []byte `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *NodeInfo) Reset()                    { *m = NodeInfo{} }
func (m *NodeInfo) String() string            { return proto.CompactTextString(m) }
func (*NodeInfo) ProtoMessage()               {}
func (*NodeInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *NodeInfo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *NodeInfo) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

// a protocol define a set of reuqest and responses
type PingRequest struct {
	MessageData *MessageData `protobuf:"bytes,1,opt,name=messageData" json:"messageData,omitempty"`
	// method specific data
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *PingRequest) Reset()                    { *m = PingRequest{} }
func (m *PingRequest) String() string            { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()               {}
func (*PingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PingRequest) GetMessageData() *MessageData {
	if m != nil {
		return m.MessageData
	}
	return nil
}

func (m *PingRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type PingResponse struct {
	MessageData *MessageData `protobuf:"bytes,1,opt,name=messageData" json:"messageData,omitempty"`
	// response specific data
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *PingResponse) Reset()                    { *m = PingResponse{} }
func (m *PingResponse) String() string            { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()               {}
func (*PingResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PingResponse) GetMessageData() *MessageData {
	if m != nil {
		return m.MessageData
	}
	return nil
}

func (m *PingResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// a protocol define a set of reuqest and responses
type EchoRequest struct {
	MessageData *MessageData `protobuf:"bytes,1,opt,name=messageData" json:"messageData,omitempty"`
	// method specific data
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *EchoRequest) Reset()                    { *m = EchoRequest{} }
func (m *EchoRequest) String() string            { return proto.CompactTextString(m) }
func (*EchoRequest) ProtoMessage()               {}
func (*EchoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *EchoRequest) GetMessageData() *MessageData {
	if m != nil {
		return m.MessageData
	}
	return nil
}

func (m *EchoRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type EchoResponse struct {
	MessageData *MessageData `protobuf:"bytes,1,opt,name=messageData" json:"messageData,omitempty"`
	// response specific data
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *EchoResponse) Reset()                    { *m = EchoResponse{} }
func (m *EchoResponse) String() string            { return proto.CompactTextString(m) }
func (*EchoResponse) ProtoMessage()               {}
func (*EchoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *EchoResponse) GetMessageData() *MessageData {
	if m != nil {
		return m.MessageData
	}
	return nil
}

func (m *EchoResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*MessageData)(nil), "pb.MessageData")
	proto.RegisterType((*NodeInfo)(nil), "pb.NodeInfo")
	proto.RegisterType((*PingRequest)(nil), "pb.PingRequest")
	proto.RegisterType((*PingResponse)(nil), "pb.PingResponse")
	proto.RegisterType((*EchoRequest)(nil), "pb.EchoRequest")
	proto.RegisterType((*EchoResponse)(nil), "pb.EchoResponse")
}

func init() { proto.RegisterFile("p2p.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0xe5, 0x34, 0xa4, 0xcd, 0x25, 0x80, 0xe4, 0x01, 0x79, 0x40, 0x28, 0x8a, 0x18, 0x32,
	0x45, 0xa2, 0xf0, 0x0b, 0x10, 0x0c, 0x08, 0x81, 0x2a, 0x0f, 0x0c, 0x61, 0x4a, 0x6a, 0x13, 0x2c,
	0x11, 0xdb, 0xe4, 0xdc, 0x81, 0x1f, 0xc8, 0xff, 0x42, 0x71, 0x53, 0x1a, 0x76, 0xba, 0xf9, 0x3d,
	0xdd, 0xfb, 0xee, 0xe9, 0x64, 0x88, 0xed, 0xd2, 0x96, 0xb6, 0x37, 0xce, 0xd0, 0xc0, 0x36, 0xf9,
	0x37, 0x81, 0xe4, 0x49, 0x22, 0xd6, 0xad, 0xbc, 0xab, 0x5d, 0x4d, 0x2f, 0xe1, 0x78, 0xfd, 0xa1,
	0xa4, 0x76, 0x2f, 0xb2, 0x47, 0x65, 0x34, 0x23, 0x19, 0x29, 0x62, 0xfe, 0xd7, 0xa4, 0xe7, 0x10,
	0x3b, 0xd5, 0x49, 0x74, 0x75, 0x67, 0x59, 0x90, 0x91, 0x62, 0xc6, 0xf7, 0x06, 0x3d, 0x81, 0x40,
	0x09, 0x36, 0xf3, 0xc1, 0x40, 0x09, 0x7a, 0x06, 0x51, 0x6b, 0x10, 0x95, 0x65, 0x61, 0x46, 0x8a,
	0x05, 0x1f, 0xd5, 0xe0, 0x6b, 0x23, 0xe4, 0x83, 0x60, 0x47, 0x7e, 0x76, 0x54, 0xf4, 0x02, 0x60,
	0x78, 0xad, 0x36, 0xcd, 0xa3, 0xfc, 0x62, 0x51, 0x46, 0x8a, 0x94, 0x4f, 0x1c, 0x4a, 0x21, 0x44,
	0xd5, 0x6a, 0x36, 0xf7, 0x29, 0xff, 0xce, 0x6f, 0x60, 0xf1, 0x3c, 0xa4, 0xf5, 0x9b, 0x19, 0xf7,
	0x93, 0xdf, 0xfd, 0x0c, 0xe6, 0xb5, 0x10, 0xbd, 0x44, 0xf4, 0x5d, 0x53, 0xbe, 0x93, 0x79, 0x05,
	0xc9, 0x4a, 0xe9, 0x96, 0xcb, 0xcf, 0x8d, 0x44, 0x47, 0xaf, 0x20, 0xe9, 0xf6, 0xb7, 0xf0, 0x84,
	0x64, 0x79, 0x5a, 0xda, 0xa6, 0x9c, 0x9c, 0x88, 0x4f, 0x67, 0x06, 0xf6, 0x28, 0x3d, 0x3b, 0xe6,
	0x3b, 0x99, 0xbf, 0x42, 0xba, 0x65, 0xa3, 0x35, 0x1a, 0xe5, 0xff, 0xc2, 0x2b, 0x48, 0xee, 0xd7,
	0xef, 0xe6, 0x50, 0xc5, 0xb7, 0xec, 0x03, 0x14, 0xbf, 0x0d, 0xab, 0xc0, 0x36, 0x4d, 0xe4, 0x3f,
	0xe0, 0xf5, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcb, 0x02, 0x9d, 0x62, 0x8d, 0x02, 0x00, 0x00,
}
