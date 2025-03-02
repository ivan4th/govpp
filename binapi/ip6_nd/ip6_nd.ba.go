// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.4.0
//  VPP:              21.06-release
// source: /usr/share/vpp/api/core/ip6_nd.api.json

// Package ip6_nd contains generated bindings for API file ip6_nd.api.
//
// Contents:
//   1 struct
//  13 messages
//
package ip6_nd

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
	APIFile    = "ip6_nd"
	APIVersion = "1.0.0"
	VersionCrc = 0x90f5895b
)

// IP6RaPrefixInfo defines type 'ip6_ra_prefix_info'.
type IP6RaPrefixInfo struct {
	Prefix        ip_types.Prefix `binapi:"prefix,name=prefix" json:"prefix,omitempty"`
	Flags         uint8           `binapi:"u8,name=flags" json:"flags,omitempty"`
	ValidTime     uint32          `binapi:"u32,name=valid_time" json:"valid_time,omitempty"`
	PreferredTime uint32          `binapi:"u32,name=preferred_time" json:"preferred_time,omitempty"`
}

// IP6RaEvent defines message 'ip6_ra_event'.
type IP6RaEvent struct {
	PID                                                 uint32                         `binapi:"u32,name=pid" json:"pid,omitempty"`
	SwIfIndex                                           interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	RouterAddr                                          ip_types.IP6Address            `binapi:"ip6_address,name=router_addr" json:"router_addr,omitempty"`
	CurrentHopLimit                                     uint8                          `binapi:"u8,name=current_hop_limit" json:"current_hop_limit,omitempty"`
	Flags                                               uint8                          `binapi:"u8,name=flags" json:"flags,omitempty"`
	RouterLifetimeInSec                                 uint16                         `binapi:"u16,name=router_lifetime_in_sec" json:"router_lifetime_in_sec,omitempty"`
	NeighborReachableTimeInMsec                         uint32                         `binapi:"u32,name=neighbor_reachable_time_in_msec" json:"neighbor_reachable_time_in_msec,omitempty"`
	TimeInMsecBetweenRetransmittedNeighborSolicitations uint32                         `binapi:"u32,name=time_in_msec_between_retransmitted_neighbor_solicitations" json:"time_in_msec_between_retransmitted_neighbor_solicitations,omitempty"`
	NPrefixes                                           uint32                         `binapi:"u32,name=n_prefixes" json:"-"`
	Prefixes                                            []IP6RaPrefixInfo              `binapi:"ip6_ra_prefix_info[n_prefixes],name=prefixes" json:"prefixes,omitempty"`
}

func (m *IP6RaEvent) Reset()               { *m = IP6RaEvent{} }
func (*IP6RaEvent) GetMessageName() string { return "ip6_ra_event" }
func (*IP6RaEvent) GetCrcString() string   { return "0364c1c5" }
func (*IP6RaEvent) GetMessageType() api.MessageType {
	return api.EventMessage
}

func (m *IP6RaEvent) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4      // m.PID
	size += 4      // m.SwIfIndex
	size += 1 * 16 // m.RouterAddr
	size += 1      // m.CurrentHopLimit
	size += 1      // m.Flags
	size += 2      // m.RouterLifetimeInSec
	size += 4      // m.NeighborReachableTimeInMsec
	size += 4      // m.TimeInMsecBetweenRetransmittedNeighborSolicitations
	size += 4      // m.NPrefixes
	for j1 := 0; j1 < len(m.Prefixes); j1++ {
		var s1 IP6RaPrefixInfo
		_ = s1
		if j1 < len(m.Prefixes) {
			s1 = m.Prefixes[j1]
		}
		size += 1      // s1.Prefix.Address.Af
		size += 1 * 16 // s1.Prefix.Address.Un
		size += 1      // s1.Prefix.Len
		size += 1      // s1.Flags
		size += 4      // s1.ValidTime
		size += 4      // s1.PreferredTime
	}
	return size
}
func (m *IP6RaEvent) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.PID)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeBytes(m.RouterAddr[:], 16)
	buf.EncodeUint8(m.CurrentHopLimit)
	buf.EncodeUint8(m.Flags)
	buf.EncodeUint16(m.RouterLifetimeInSec)
	buf.EncodeUint32(m.NeighborReachableTimeInMsec)
	buf.EncodeUint32(m.TimeInMsecBetweenRetransmittedNeighborSolicitations)
	buf.EncodeUint32(uint32(len(m.Prefixes)))
	for j0 := 0; j0 < len(m.Prefixes); j0++ {
		var v0 IP6RaPrefixInfo // Prefixes
		if j0 < len(m.Prefixes) {
			v0 = m.Prefixes[j0]
		}
		buf.EncodeUint8(uint8(v0.Prefix.Address.Af))
		buf.EncodeBytes(v0.Prefix.Address.Un.XXX_UnionData[:], 16)
		buf.EncodeUint8(v0.Prefix.Len)
		buf.EncodeUint8(v0.Flags)
		buf.EncodeUint32(v0.ValidTime)
		buf.EncodeUint32(v0.PreferredTime)
	}
	return buf.Bytes(), nil
}
func (m *IP6RaEvent) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.PID = buf.DecodeUint32()
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	copy(m.RouterAddr[:], buf.DecodeBytes(16))
	m.CurrentHopLimit = buf.DecodeUint8()
	m.Flags = buf.DecodeUint8()
	m.RouterLifetimeInSec = buf.DecodeUint16()
	m.NeighborReachableTimeInMsec = buf.DecodeUint32()
	m.TimeInMsecBetweenRetransmittedNeighborSolicitations = buf.DecodeUint32()
	m.NPrefixes = buf.DecodeUint32()
	m.Prefixes = make([]IP6RaPrefixInfo, m.NPrefixes)
	for j0 := 0; j0 < len(m.Prefixes); j0++ {
		m.Prefixes[j0].Prefix.Address.Af = ip_types.AddressFamily(buf.DecodeUint8())
		copy(m.Prefixes[j0].Prefix.Address.Un.XXX_UnionData[:], buf.DecodeBytes(16))
		m.Prefixes[j0].Prefix.Len = buf.DecodeUint8()
		m.Prefixes[j0].Flags = buf.DecodeUint8()
		m.Prefixes[j0].ValidTime = buf.DecodeUint32()
		m.Prefixes[j0].PreferredTime = buf.DecodeUint32()
	}
	return nil
}

// IP6ndProxyAddDel defines message 'ip6nd_proxy_add_del'.
type IP6ndProxyAddDel struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	IsAdd     bool                           `binapi:"bool,name=is_add" json:"is_add,omitempty"`
	IP        ip_types.IP6Address            `binapi:"ip6_address,name=ip" json:"ip,omitempty"`
}

func (m *IP6ndProxyAddDel) Reset()               { *m = IP6ndProxyAddDel{} }
func (*IP6ndProxyAddDel) GetMessageName() string { return "ip6nd_proxy_add_del" }
func (*IP6ndProxyAddDel) GetCrcString() string   { return "c2e4a686" }
func (*IP6ndProxyAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *IP6ndProxyAddDel) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4      // m.SwIfIndex
	size += 1      // m.IsAdd
	size += 1 * 16 // m.IP
	return size
}
func (m *IP6ndProxyAddDel) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeBool(m.IsAdd)
	buf.EncodeBytes(m.IP[:], 16)
	return buf.Bytes(), nil
}
func (m *IP6ndProxyAddDel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.IsAdd = buf.DecodeBool()
	copy(m.IP[:], buf.DecodeBytes(16))
	return nil
}

// IP6ndProxyAddDelReply defines message 'ip6nd_proxy_add_del_reply'.
type IP6ndProxyAddDelReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *IP6ndProxyAddDelReply) Reset()               { *m = IP6ndProxyAddDelReply{} }
func (*IP6ndProxyAddDelReply) GetMessageName() string { return "ip6nd_proxy_add_del_reply" }
func (*IP6ndProxyAddDelReply) GetCrcString() string   { return "e8d4e804" }
func (*IP6ndProxyAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *IP6ndProxyAddDelReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *IP6ndProxyAddDelReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *IP6ndProxyAddDelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// IP6ndProxyDetails defines message 'ip6nd_proxy_details'.
type IP6ndProxyDetails struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	IP        ip_types.IP6Address            `binapi:"ip6_address,name=ip" json:"ip,omitempty"`
}

func (m *IP6ndProxyDetails) Reset()               { *m = IP6ndProxyDetails{} }
func (*IP6ndProxyDetails) GetMessageName() string { return "ip6nd_proxy_details" }
func (*IP6ndProxyDetails) GetCrcString() string   { return "30b9ff4a" }
func (*IP6ndProxyDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *IP6ndProxyDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4      // m.SwIfIndex
	size += 1 * 16 // m.IP
	return size
}
func (m *IP6ndProxyDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeBytes(m.IP[:], 16)
	return buf.Bytes(), nil
}
func (m *IP6ndProxyDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	copy(m.IP[:], buf.DecodeBytes(16))
	return nil
}

// IP6ndProxyDump defines message 'ip6nd_proxy_dump'.
type IP6ndProxyDump struct{}

func (m *IP6ndProxyDump) Reset()               { *m = IP6ndProxyDump{} }
func (*IP6ndProxyDump) GetMessageName() string { return "ip6nd_proxy_dump" }
func (*IP6ndProxyDump) GetCrcString() string   { return "51077d14" }
func (*IP6ndProxyDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *IP6ndProxyDump) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *IP6ndProxyDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *IP6ndProxyDump) Unmarshal(b []byte) error {
	return nil
}

// IP6ndSendRouterSolicitation defines message 'ip6nd_send_router_solicitation'.
type IP6ndSendRouterSolicitation struct {
	Irt       uint32                         `binapi:"u32,name=irt" json:"irt,omitempty"`
	Mrt       uint32                         `binapi:"u32,name=mrt" json:"mrt,omitempty"`
	Mrc       uint32                         `binapi:"u32,name=mrc" json:"mrc,omitempty"`
	Mrd       uint32                         `binapi:"u32,name=mrd" json:"mrd,omitempty"`
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	Stop      bool                           `binapi:"bool,name=stop" json:"stop,omitempty"`
}

func (m *IP6ndSendRouterSolicitation) Reset()               { *m = IP6ndSendRouterSolicitation{} }
func (*IP6ndSendRouterSolicitation) GetMessageName() string { return "ip6nd_send_router_solicitation" }
func (*IP6ndSendRouterSolicitation) GetCrcString() string   { return "e5de609c" }
func (*IP6ndSendRouterSolicitation) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *IP6ndSendRouterSolicitation) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Irt
	size += 4 // m.Mrt
	size += 4 // m.Mrc
	size += 4 // m.Mrd
	size += 4 // m.SwIfIndex
	size += 1 // m.Stop
	return size
}
func (m *IP6ndSendRouterSolicitation) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.Irt)
	buf.EncodeUint32(m.Mrt)
	buf.EncodeUint32(m.Mrc)
	buf.EncodeUint32(m.Mrd)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeBool(m.Stop)
	return buf.Bytes(), nil
}
func (m *IP6ndSendRouterSolicitation) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Irt = buf.DecodeUint32()
	m.Mrt = buf.DecodeUint32()
	m.Mrc = buf.DecodeUint32()
	m.Mrd = buf.DecodeUint32()
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.Stop = buf.DecodeBool()
	return nil
}

// IP6ndSendRouterSolicitationReply defines message 'ip6nd_send_router_solicitation_reply'.
type IP6ndSendRouterSolicitationReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *IP6ndSendRouterSolicitationReply) Reset() { *m = IP6ndSendRouterSolicitationReply{} }
func (*IP6ndSendRouterSolicitationReply) GetMessageName() string {
	return "ip6nd_send_router_solicitation_reply"
}
func (*IP6ndSendRouterSolicitationReply) GetCrcString() string { return "e8d4e804" }
func (*IP6ndSendRouterSolicitationReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *IP6ndSendRouterSolicitationReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *IP6ndSendRouterSolicitationReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *IP6ndSendRouterSolicitationReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// SwInterfaceIP6ndRaConfig defines message 'sw_interface_ip6nd_ra_config'.
type SwInterfaceIP6ndRaConfig struct {
	SwIfIndex       interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	Suppress        uint8                          `binapi:"u8,name=suppress" json:"suppress,omitempty"`
	Managed         uint8                          `binapi:"u8,name=managed" json:"managed,omitempty"`
	Other           uint8                          `binapi:"u8,name=other" json:"other,omitempty"`
	LlOption        uint8                          `binapi:"u8,name=ll_option" json:"ll_option,omitempty"`
	SendUnicast     uint8                          `binapi:"u8,name=send_unicast" json:"send_unicast,omitempty"`
	Cease           uint8                          `binapi:"u8,name=cease" json:"cease,omitempty"`
	IsNo            bool                           `binapi:"bool,name=is_no" json:"is_no,omitempty"`
	DefaultRouter   uint8                          `binapi:"u8,name=default_router" json:"default_router,omitempty"`
	MaxInterval     uint32                         `binapi:"u32,name=max_interval" json:"max_interval,omitempty"`
	MinInterval     uint32                         `binapi:"u32,name=min_interval" json:"min_interval,omitempty"`
	Lifetime        uint32                         `binapi:"u32,name=lifetime" json:"lifetime,omitempty"`
	InitialCount    uint32                         `binapi:"u32,name=initial_count" json:"initial_count,omitempty"`
	InitialInterval uint32                         `binapi:"u32,name=initial_interval" json:"initial_interval,omitempty"`
}

func (m *SwInterfaceIP6ndRaConfig) Reset()               { *m = SwInterfaceIP6ndRaConfig{} }
func (*SwInterfaceIP6ndRaConfig) GetMessageName() string { return "sw_interface_ip6nd_ra_config" }
func (*SwInterfaceIP6ndRaConfig) GetCrcString() string   { return "3eb00b1c" }
func (*SwInterfaceIP6ndRaConfig) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SwInterfaceIP6ndRaConfig) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	size += 1 // m.Suppress
	size += 1 // m.Managed
	size += 1 // m.Other
	size += 1 // m.LlOption
	size += 1 // m.SendUnicast
	size += 1 // m.Cease
	size += 1 // m.IsNo
	size += 1 // m.DefaultRouter
	size += 4 // m.MaxInterval
	size += 4 // m.MinInterval
	size += 4 // m.Lifetime
	size += 4 // m.InitialCount
	size += 4 // m.InitialInterval
	return size
}
func (m *SwInterfaceIP6ndRaConfig) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeUint8(m.Suppress)
	buf.EncodeUint8(m.Managed)
	buf.EncodeUint8(m.Other)
	buf.EncodeUint8(m.LlOption)
	buf.EncodeUint8(m.SendUnicast)
	buf.EncodeUint8(m.Cease)
	buf.EncodeBool(m.IsNo)
	buf.EncodeUint8(m.DefaultRouter)
	buf.EncodeUint32(m.MaxInterval)
	buf.EncodeUint32(m.MinInterval)
	buf.EncodeUint32(m.Lifetime)
	buf.EncodeUint32(m.InitialCount)
	buf.EncodeUint32(m.InitialInterval)
	return buf.Bytes(), nil
}
func (m *SwInterfaceIP6ndRaConfig) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.Suppress = buf.DecodeUint8()
	m.Managed = buf.DecodeUint8()
	m.Other = buf.DecodeUint8()
	m.LlOption = buf.DecodeUint8()
	m.SendUnicast = buf.DecodeUint8()
	m.Cease = buf.DecodeUint8()
	m.IsNo = buf.DecodeBool()
	m.DefaultRouter = buf.DecodeUint8()
	m.MaxInterval = buf.DecodeUint32()
	m.MinInterval = buf.DecodeUint32()
	m.Lifetime = buf.DecodeUint32()
	m.InitialCount = buf.DecodeUint32()
	m.InitialInterval = buf.DecodeUint32()
	return nil
}

// SwInterfaceIP6ndRaConfigReply defines message 'sw_interface_ip6nd_ra_config_reply'.
type SwInterfaceIP6ndRaConfigReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *SwInterfaceIP6ndRaConfigReply) Reset() { *m = SwInterfaceIP6ndRaConfigReply{} }
func (*SwInterfaceIP6ndRaConfigReply) GetMessageName() string {
	return "sw_interface_ip6nd_ra_config_reply"
}
func (*SwInterfaceIP6ndRaConfigReply) GetCrcString() string { return "e8d4e804" }
func (*SwInterfaceIP6ndRaConfigReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *SwInterfaceIP6ndRaConfigReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *SwInterfaceIP6ndRaConfigReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *SwInterfaceIP6ndRaConfigReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// SwInterfaceIP6ndRaPrefix defines message 'sw_interface_ip6nd_ra_prefix'.
type SwInterfaceIP6ndRaPrefix struct {
	SwIfIndex    interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	Prefix       ip_types.Prefix                `binapi:"prefix,name=prefix" json:"prefix,omitempty"`
	UseDefault   bool                           `binapi:"bool,name=use_default" json:"use_default,omitempty"`
	NoAdvertise  bool                           `binapi:"bool,name=no_advertise" json:"no_advertise,omitempty"`
	OffLink      bool                           `binapi:"bool,name=off_link" json:"off_link,omitempty"`
	NoAutoconfig bool                           `binapi:"bool,name=no_autoconfig" json:"no_autoconfig,omitempty"`
	NoOnlink     bool                           `binapi:"bool,name=no_onlink" json:"no_onlink,omitempty"`
	IsNo         bool                           `binapi:"bool,name=is_no" json:"is_no,omitempty"`
	ValLifetime  uint32                         `binapi:"u32,name=val_lifetime" json:"val_lifetime,omitempty"`
	PrefLifetime uint32                         `binapi:"u32,name=pref_lifetime" json:"pref_lifetime,omitempty"`
}

func (m *SwInterfaceIP6ndRaPrefix) Reset()               { *m = SwInterfaceIP6ndRaPrefix{} }
func (*SwInterfaceIP6ndRaPrefix) GetMessageName() string { return "sw_interface_ip6nd_ra_prefix" }
func (*SwInterfaceIP6ndRaPrefix) GetCrcString() string   { return "82cc1b28" }
func (*SwInterfaceIP6ndRaPrefix) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SwInterfaceIP6ndRaPrefix) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4      // m.SwIfIndex
	size += 1      // m.Prefix.Address.Af
	size += 1 * 16 // m.Prefix.Address.Un
	size += 1      // m.Prefix.Len
	size += 1      // m.UseDefault
	size += 1      // m.NoAdvertise
	size += 1      // m.OffLink
	size += 1      // m.NoAutoconfig
	size += 1      // m.NoOnlink
	size += 1      // m.IsNo
	size += 4      // m.ValLifetime
	size += 4      // m.PrefLifetime
	return size
}
func (m *SwInterfaceIP6ndRaPrefix) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeUint8(uint8(m.Prefix.Address.Af))
	buf.EncodeBytes(m.Prefix.Address.Un.XXX_UnionData[:], 16)
	buf.EncodeUint8(m.Prefix.Len)
	buf.EncodeBool(m.UseDefault)
	buf.EncodeBool(m.NoAdvertise)
	buf.EncodeBool(m.OffLink)
	buf.EncodeBool(m.NoAutoconfig)
	buf.EncodeBool(m.NoOnlink)
	buf.EncodeBool(m.IsNo)
	buf.EncodeUint32(m.ValLifetime)
	buf.EncodeUint32(m.PrefLifetime)
	return buf.Bytes(), nil
}
func (m *SwInterfaceIP6ndRaPrefix) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.Prefix.Address.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.Prefix.Address.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.Prefix.Len = buf.DecodeUint8()
	m.UseDefault = buf.DecodeBool()
	m.NoAdvertise = buf.DecodeBool()
	m.OffLink = buf.DecodeBool()
	m.NoAutoconfig = buf.DecodeBool()
	m.NoOnlink = buf.DecodeBool()
	m.IsNo = buf.DecodeBool()
	m.ValLifetime = buf.DecodeUint32()
	m.PrefLifetime = buf.DecodeUint32()
	return nil
}

// SwInterfaceIP6ndRaPrefixReply defines message 'sw_interface_ip6nd_ra_prefix_reply'.
type SwInterfaceIP6ndRaPrefixReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *SwInterfaceIP6ndRaPrefixReply) Reset() { *m = SwInterfaceIP6ndRaPrefixReply{} }
func (*SwInterfaceIP6ndRaPrefixReply) GetMessageName() string {
	return "sw_interface_ip6nd_ra_prefix_reply"
}
func (*SwInterfaceIP6ndRaPrefixReply) GetCrcString() string { return "e8d4e804" }
func (*SwInterfaceIP6ndRaPrefixReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *SwInterfaceIP6ndRaPrefixReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *SwInterfaceIP6ndRaPrefixReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *SwInterfaceIP6ndRaPrefixReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// WantIP6RaEvents defines message 'want_ip6_ra_events'.
type WantIP6RaEvents struct {
	Enable bool   `binapi:"bool,name=enable" json:"enable,omitempty"`
	PID    uint32 `binapi:"u32,name=pid" json:"pid,omitempty"`
}

func (m *WantIP6RaEvents) Reset()               { *m = WantIP6RaEvents{} }
func (*WantIP6RaEvents) GetMessageName() string { return "want_ip6_ra_events" }
func (*WantIP6RaEvents) GetCrcString() string   { return "3ec6d6c2" }
func (*WantIP6RaEvents) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *WantIP6RaEvents) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.Enable
	size += 4 // m.PID
	return size
}
func (m *WantIP6RaEvents) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.Enable)
	buf.EncodeUint32(m.PID)
	return buf.Bytes(), nil
}
func (m *WantIP6RaEvents) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Enable = buf.DecodeBool()
	m.PID = buf.DecodeUint32()
	return nil
}

// WantIP6RaEventsReply defines message 'want_ip6_ra_events_reply'.
type WantIP6RaEventsReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *WantIP6RaEventsReply) Reset()               { *m = WantIP6RaEventsReply{} }
func (*WantIP6RaEventsReply) GetMessageName() string { return "want_ip6_ra_events_reply" }
func (*WantIP6RaEventsReply) GetCrcString() string   { return "e8d4e804" }
func (*WantIP6RaEventsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *WantIP6RaEventsReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *WantIP6RaEventsReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *WantIP6RaEventsReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_ip6_nd_binapi_init() }
func file_ip6_nd_binapi_init() {
	api.RegisterMessage((*IP6RaEvent)(nil), "ip6_ra_event_0364c1c5")
	api.RegisterMessage((*IP6ndProxyAddDel)(nil), "ip6nd_proxy_add_del_c2e4a686")
	api.RegisterMessage((*IP6ndProxyAddDelReply)(nil), "ip6nd_proxy_add_del_reply_e8d4e804")
	api.RegisterMessage((*IP6ndProxyDetails)(nil), "ip6nd_proxy_details_30b9ff4a")
	api.RegisterMessage((*IP6ndProxyDump)(nil), "ip6nd_proxy_dump_51077d14")
	api.RegisterMessage((*IP6ndSendRouterSolicitation)(nil), "ip6nd_send_router_solicitation_e5de609c")
	api.RegisterMessage((*IP6ndSendRouterSolicitationReply)(nil), "ip6nd_send_router_solicitation_reply_e8d4e804")
	api.RegisterMessage((*SwInterfaceIP6ndRaConfig)(nil), "sw_interface_ip6nd_ra_config_3eb00b1c")
	api.RegisterMessage((*SwInterfaceIP6ndRaConfigReply)(nil), "sw_interface_ip6nd_ra_config_reply_e8d4e804")
	api.RegisterMessage((*SwInterfaceIP6ndRaPrefix)(nil), "sw_interface_ip6nd_ra_prefix_82cc1b28")
	api.RegisterMessage((*SwInterfaceIP6ndRaPrefixReply)(nil), "sw_interface_ip6nd_ra_prefix_reply_e8d4e804")
	api.RegisterMessage((*WantIP6RaEvents)(nil), "want_ip6_ra_events_3ec6d6c2")
	api.RegisterMessage((*WantIP6RaEventsReply)(nil), "want_ip6_ra_events_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*IP6RaEvent)(nil),
		(*IP6ndProxyAddDel)(nil),
		(*IP6ndProxyAddDelReply)(nil),
		(*IP6ndProxyDetails)(nil),
		(*IP6ndProxyDump)(nil),
		(*IP6ndSendRouterSolicitation)(nil),
		(*IP6ndSendRouterSolicitationReply)(nil),
		(*SwInterfaceIP6ndRaConfig)(nil),
		(*SwInterfaceIP6ndRaConfigReply)(nil),
		(*SwInterfaceIP6ndRaPrefix)(nil),
		(*SwInterfaceIP6ndRaPrefixReply)(nil),
		(*WantIP6RaEvents)(nil),
		(*WantIP6RaEventsReply)(nil),
	}
}
