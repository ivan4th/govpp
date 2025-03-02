// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.4.0
//  VPP:              21.06-release
// source: /usr/share/vpp/api/plugins/arping.api.json

// Package arping contains generated bindings for API file arping.api.
//
// Contents:
//   2 messages
//
package arping

import (
	api "git.fd.io/govpp.git/api"
	interface_types "git.fd.io/govpp.git/binapi/interface_types"
	ip_types "git.fd.io/govpp.git/binapi/ip_types"
	codec "git.fd.io/govpp.git/codec"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "arping"
	APIVersion = "1.0.0"
	VersionCrc = 0x666f91cc
)

// Arping defines message 'arping'.
type Arping struct {
	Address   ip_types.Address               `binapi:"address,name=address" json:"address,omitempty"`
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	IsGarp    bool                           `binapi:"bool,name=is_garp" json:"is_garp,omitempty"`
	Repeat    uint32                         `binapi:"u32,name=repeat,default=1" json:"repeat,omitempty"`
	Interval  float64                        `binapi:"f64,name=interval,default=1" json:"interval,omitempty"`
}

func (m *Arping) Reset()               { *m = Arping{} }
func (*Arping) GetMessageName() string { return "arping" }
func (*Arping) GetCrcString() string   { return "48817482" }
func (*Arping) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *Arping) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1      // m.Address.Af
	size += 1 * 16 // m.Address.Un
	size += 4      // m.SwIfIndex
	size += 1      // m.IsGarp
	size += 4      // m.Repeat
	size += 8      // m.Interval
	return size
}
func (m *Arping) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint8(uint8(m.Address.Af))
	buf.EncodeBytes(m.Address.Un.XXX_UnionData[:], 16)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeBool(m.IsGarp)
	buf.EncodeUint32(m.Repeat)
	buf.EncodeFloat64(m.Interval)
	return buf.Bytes(), nil
}
func (m *Arping) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Address.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.Address.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.IsGarp = buf.DecodeBool()
	m.Repeat = buf.DecodeUint32()
	m.Interval = buf.DecodeFloat64()
	return nil
}

// ArpingReply defines message 'arping_reply'.
type ArpingReply struct {
	Retval     int32  `binapi:"i32,name=retval" json:"retval,omitempty"`
	ReplyCount uint32 `binapi:"u32,name=reply_count" json:"reply_count,omitempty"`
}

func (m *ArpingReply) Reset()               { *m = ArpingReply{} }
func (*ArpingReply) GetMessageName() string { return "arping_reply" }
func (*ArpingReply) GetCrcString() string   { return "bb9d1cbd" }
func (*ArpingReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *ArpingReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.ReplyCount
	return size
}
func (m *ArpingReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(m.ReplyCount)
	return buf.Bytes(), nil
}
func (m *ArpingReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.ReplyCount = buf.DecodeUint32()
	return nil
}

func init() { file_arping_binapi_init() }
func file_arping_binapi_init() {
	api.RegisterMessage((*Arping)(nil), "arping_48817482")
	api.RegisterMessage((*ArpingReply)(nil), "arping_reply_bb9d1cbd")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*Arping)(nil),
		(*ArpingReply)(nil),
	}
}
