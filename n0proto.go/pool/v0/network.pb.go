// Code generated by protoc-gen-go. DO NOT EDIT.
// source: n0stack/pool/v0/network.proto

package ppool

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	v0 "github.com/onokatio/terraform-provider-n0stack/n0proto.go/budget/v0"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Network_NetworkState int32

const (
	// falied state because failed some process on API.
	Network_NETWORK_UNSPECIFIED Network_NetworkState = 0
	// unknown state because failed to connect for scheduled node after RUNNING.
	Network_UNKNOWN   Network_NetworkState = 1
	Network_AVAILABLE Network_NetworkState = 2
)

var Network_NetworkState_name = map[int32]string{
	0: "NETWORK_UNSPECIFIED",
	1: "UNKNOWN",
	2: "AVAILABLE",
}

var Network_NetworkState_value = map[string]int32{
	"NETWORK_UNSPECIFIED": 0,
	"UNKNOWN":             1,
	"AVAILABLE":           2,
}

func (x Network_NetworkState) String() string {
	return proto.EnumName(Network_NetworkState_name, int32(x))
}

func (Network_NetworkState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_462186787d56a6a3, []int{0, 0}
}

type Network struct {
	// Name is a unique field.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Annotations can store metadata used by the system for control.
	// In particular, implementation-dependent fields that can not be set as protobuf fields are targeted.
	// The control specified by n0stack may delete metadata specified by the user.
	Annotations map[string]string `protobuf:"bytes,3,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Labels stores user-defined metadata.
	// The n0stack system must not rewrite this value.
	Labels                    map[string]string               `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Ipv4Cidr                  string                          `protobuf:"bytes,10,opt,name=ipv4_cidr,json=ipv4Cidr,proto3" json:"ipv4_cidr,omitempty"`
	Ipv6Cidr                  string                          `protobuf:"bytes,11,opt,name=ipv6_cidr,json=ipv6Cidr,proto3" json:"ipv6_cidr,omitempty"`
	Domain                    string                          `protobuf:"bytes,12,opt,name=domain,proto3" json:"domain,omitempty"`
	State                     Network_NetworkState            `protobuf:"varint,50,opt,name=state,proto3,enum=n0stack.pool.v0.Network_NetworkState" json:"state,omitempty"`
	ReservedNetworkInterfaces map[string]*v0.NetworkInterface `protobuf:"bytes,51,rep,name=reserved_network_interfaces,json=reservedNetworkInterfaces,proto3" json:"reserved_network_interfaces,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral      struct{}                        `json:"-"`
	XXX_unrecognized          []byte                          `json:"-"`
	XXX_sizecache             int32                           `json:"-"`
}

func (m *Network) Reset()         { *m = Network{} }
func (m *Network) String() string { return proto.CompactTextString(m) }
func (*Network) ProtoMessage()    {}
func (*Network) Descriptor() ([]byte, []int) {
	return fileDescriptor_462186787d56a6a3, []int{0}
}

func (m *Network) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Network.Unmarshal(m, b)
}
func (m *Network) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Network.Marshal(b, m, deterministic)
}
func (m *Network) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Network.Merge(m, src)
}
func (m *Network) XXX_Size() int {
	return xxx_messageInfo_Network.Size(m)
}
func (m *Network) XXX_DiscardUnknown() {
	xxx_messageInfo_Network.DiscardUnknown(m)
}

var xxx_messageInfo_Network proto.InternalMessageInfo

func (m *Network) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Network) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func (m *Network) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Network) GetIpv4Cidr() string {
	if m != nil {
		return m.Ipv4Cidr
	}
	return ""
}

func (m *Network) GetIpv6Cidr() string {
	if m != nil {
		return m.Ipv6Cidr
	}
	return ""
}

func (m *Network) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *Network) GetState() Network_NetworkState {
	if m != nil {
		return m.State
	}
	return Network_NETWORK_UNSPECIFIED
}

func (m *Network) GetReservedNetworkInterfaces() map[string]*v0.NetworkInterface {
	if m != nil {
		return m.ReservedNetworkInterfaces
	}
	return nil
}

type ListNetworksRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListNetworksRequest) Reset()         { *m = ListNetworksRequest{} }
func (m *ListNetworksRequest) String() string { return proto.CompactTextString(m) }
func (*ListNetworksRequest) ProtoMessage()    {}
func (*ListNetworksRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_462186787d56a6a3, []int{1}
}

func (m *ListNetworksRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListNetworksRequest.Unmarshal(m, b)
}
func (m *ListNetworksRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListNetworksRequest.Marshal(b, m, deterministic)
}
func (m *ListNetworksRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListNetworksRequest.Merge(m, src)
}
func (m *ListNetworksRequest) XXX_Size() int {
	return xxx_messageInfo_ListNetworksRequest.Size(m)
}
func (m *ListNetworksRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListNetworksRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListNetworksRequest proto.InternalMessageInfo

type ListNetworksResponse struct {
	Networks             []*Network `protobuf:"bytes,1,rep,name=networks,proto3" json:"networks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListNetworksResponse) Reset()         { *m = ListNetworksResponse{} }
func (m *ListNetworksResponse) String() string { return proto.CompactTextString(m) }
func (*ListNetworksResponse) ProtoMessage()    {}
func (*ListNetworksResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_462186787d56a6a3, []int{2}
}

func (m *ListNetworksResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListNetworksResponse.Unmarshal(m, b)
}
func (m *ListNetworksResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListNetworksResponse.Marshal(b, m, deterministic)
}
func (m *ListNetworksResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListNetworksResponse.Merge(m, src)
}
func (m *ListNetworksResponse) XXX_Size() int {
	return xxx_messageInfo_ListNetworksResponse.Size(m)
}
func (m *ListNetworksResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListNetworksResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListNetworksResponse proto.InternalMessageInfo

func (m *ListNetworksResponse) GetNetworks() []*Network {
	if m != nil {
		return m.Networks
	}
	return nil
}

type GetNetworkRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetNetworkRequest) Reset()         { *m = GetNetworkRequest{} }
func (m *GetNetworkRequest) String() string { return proto.CompactTextString(m) }
func (*GetNetworkRequest) ProtoMessage()    {}
func (*GetNetworkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_462186787d56a6a3, []int{3}
}

func (m *GetNetworkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNetworkRequest.Unmarshal(m, b)
}
func (m *GetNetworkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNetworkRequest.Marshal(b, m, deterministic)
}
func (m *GetNetworkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNetworkRequest.Merge(m, src)
}
func (m *GetNetworkRequest) XXX_Size() int {
	return xxx_messageInfo_GetNetworkRequest.Size(m)
}
func (m *GetNetworkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNetworkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetNetworkRequest proto.InternalMessageInfo

func (m *GetNetworkRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ApplyNetworkRequest struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Annotations          map[string]string `protobuf:"bytes,3,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Labels               map[string]string `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Ipv4Cidr             string            `protobuf:"bytes,10,opt,name=ipv4_cidr,json=ipv4Cidr,proto3" json:"ipv4_cidr,omitempty"`
	Ipv6Cidr             string            `protobuf:"bytes,11,opt,name=ipv6_cidr,json=ipv6Cidr,proto3" json:"ipv6_cidr,omitempty"`
	Domain               string            `protobuf:"bytes,12,opt,name=domain,proto3" json:"domain,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ApplyNetworkRequest) Reset()         { *m = ApplyNetworkRequest{} }
func (m *ApplyNetworkRequest) String() string { return proto.CompactTextString(m) }
func (*ApplyNetworkRequest) ProtoMessage()    {}
func (*ApplyNetworkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_462186787d56a6a3, []int{4}
}

func (m *ApplyNetworkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplyNetworkRequest.Unmarshal(m, b)
}
func (m *ApplyNetworkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplyNetworkRequest.Marshal(b, m, deterministic)
}
func (m *ApplyNetworkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplyNetworkRequest.Merge(m, src)
}
func (m *ApplyNetworkRequest) XXX_Size() int {
	return xxx_messageInfo_ApplyNetworkRequest.Size(m)
}
func (m *ApplyNetworkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplyNetworkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ApplyNetworkRequest proto.InternalMessageInfo

func (m *ApplyNetworkRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ApplyNetworkRequest) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func (m *ApplyNetworkRequest) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *ApplyNetworkRequest) GetIpv4Cidr() string {
	if m != nil {
		return m.Ipv4Cidr
	}
	return ""
}

func (m *ApplyNetworkRequest) GetIpv6Cidr() string {
	if m != nil {
		return m.Ipv6Cidr
	}
	return ""
}

func (m *ApplyNetworkRequest) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

type DeleteNetworkRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteNetworkRequest) Reset()         { *m = DeleteNetworkRequest{} }
func (m *DeleteNetworkRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteNetworkRequest) ProtoMessage()    {}
func (*DeleteNetworkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_462186787d56a6a3, []int{5}
}

func (m *DeleteNetworkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteNetworkRequest.Unmarshal(m, b)
}
func (m *DeleteNetworkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteNetworkRequest.Marshal(b, m, deterministic)
}
func (m *DeleteNetworkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteNetworkRequest.Merge(m, src)
}
func (m *DeleteNetworkRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteNetworkRequest.Size(m)
}
func (m *DeleteNetworkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteNetworkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteNetworkRequest proto.InternalMessageInfo

func (m *DeleteNetworkRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ReserveNetworkInterfaceRequest struct {
	NetworkName          string            `protobuf:"bytes,1,opt,name=network_name,json=networkName,proto3" json:"network_name,omitempty"`
	NetworkInterfaceName string            `protobuf:"bytes,2,opt,name=network_interface_name,json=networkInterfaceName,proto3" json:"network_interface_name,omitempty"`
	Annotations          map[string]string `protobuf:"bytes,3,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	HardwareAddress      string            `protobuf:"bytes,4,opt,name=hardware_address,json=hardwareAddress,proto3" json:"hardware_address,omitempty"`
	Ipv4Address          string            `protobuf:"bytes,5,opt,name=ipv4_address,json=ipv4Address,proto3" json:"ipv4_address,omitempty"`
	Ipv6Address          string            `protobuf:"bytes,6,opt,name=ipv6_address,json=ipv6Address,proto3" json:"ipv6_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ReserveNetworkInterfaceRequest) Reset()         { *m = ReserveNetworkInterfaceRequest{} }
func (m *ReserveNetworkInterfaceRequest) String() string { return proto.CompactTextString(m) }
func (*ReserveNetworkInterfaceRequest) ProtoMessage()    {}
func (*ReserveNetworkInterfaceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_462186787d56a6a3, []int{6}
}

func (m *ReserveNetworkInterfaceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReserveNetworkInterfaceRequest.Unmarshal(m, b)
}
func (m *ReserveNetworkInterfaceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReserveNetworkInterfaceRequest.Marshal(b, m, deterministic)
}
func (m *ReserveNetworkInterfaceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReserveNetworkInterfaceRequest.Merge(m, src)
}
func (m *ReserveNetworkInterfaceRequest) XXX_Size() int {
	return xxx_messageInfo_ReserveNetworkInterfaceRequest.Size(m)
}
func (m *ReserveNetworkInterfaceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReserveNetworkInterfaceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReserveNetworkInterfaceRequest proto.InternalMessageInfo

func (m *ReserveNetworkInterfaceRequest) GetNetworkName() string {
	if m != nil {
		return m.NetworkName
	}
	return ""
}

func (m *ReserveNetworkInterfaceRequest) GetNetworkInterfaceName() string {
	if m != nil {
		return m.NetworkInterfaceName
	}
	return ""
}

func (m *ReserveNetworkInterfaceRequest) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func (m *ReserveNetworkInterfaceRequest) GetHardwareAddress() string {
	if m != nil {
		return m.HardwareAddress
	}
	return ""
}

func (m *ReserveNetworkInterfaceRequest) GetIpv4Address() string {
	if m != nil {
		return m.Ipv4Address
	}
	return ""
}

func (m *ReserveNetworkInterfaceRequest) GetIpv6Address() string {
	if m != nil {
		return m.Ipv6Address
	}
	return ""
}

type ReleaseNetworkInterfaceRequest struct {
	NetworkName          string   `protobuf:"bytes,1,opt,name=network_name,json=networkName,proto3" json:"network_name,omitempty"`
	NetworkInterfaceName string   `protobuf:"bytes,2,opt,name=network_interface_name,json=networkInterfaceName,proto3" json:"network_interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReleaseNetworkInterfaceRequest) Reset()         { *m = ReleaseNetworkInterfaceRequest{} }
func (m *ReleaseNetworkInterfaceRequest) String() string { return proto.CompactTextString(m) }
func (*ReleaseNetworkInterfaceRequest) ProtoMessage()    {}
func (*ReleaseNetworkInterfaceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_462186787d56a6a3, []int{7}
}

func (m *ReleaseNetworkInterfaceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReleaseNetworkInterfaceRequest.Unmarshal(m, b)
}
func (m *ReleaseNetworkInterfaceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReleaseNetworkInterfaceRequest.Marshal(b, m, deterministic)
}
func (m *ReleaseNetworkInterfaceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReleaseNetworkInterfaceRequest.Merge(m, src)
}
func (m *ReleaseNetworkInterfaceRequest) XXX_Size() int {
	return xxx_messageInfo_ReleaseNetworkInterfaceRequest.Size(m)
}
func (m *ReleaseNetworkInterfaceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReleaseNetworkInterfaceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReleaseNetworkInterfaceRequest proto.InternalMessageInfo

func (m *ReleaseNetworkInterfaceRequest) GetNetworkName() string {
	if m != nil {
		return m.NetworkName
	}
	return ""
}

func (m *ReleaseNetworkInterfaceRequest) GetNetworkInterfaceName() string {
	if m != nil {
		return m.NetworkInterfaceName
	}
	return ""
}

func init() {
	proto.RegisterEnum("n0stack.pool.v0.Network_NetworkState", Network_NetworkState_name, Network_NetworkState_value)
	proto.RegisterType((*Network)(nil), "n0stack.pool.v0.Network")
	proto.RegisterMapType((map[string]string)(nil), "n0stack.pool.v0.Network.AnnotationsEntry")
	proto.RegisterMapType((map[string]string)(nil), "n0stack.pool.v0.Network.LabelsEntry")
	proto.RegisterMapType((map[string]*v0.NetworkInterface)(nil), "n0stack.pool.v0.Network.ReservedNetworkInterfacesEntry")
	proto.RegisterType((*ListNetworksRequest)(nil), "n0stack.pool.v0.ListNetworksRequest")
	proto.RegisterType((*ListNetworksResponse)(nil), "n0stack.pool.v0.ListNetworksResponse")
	proto.RegisterType((*GetNetworkRequest)(nil), "n0stack.pool.v0.GetNetworkRequest")
	proto.RegisterType((*ApplyNetworkRequest)(nil), "n0stack.pool.v0.ApplyNetworkRequest")
	proto.RegisterMapType((map[string]string)(nil), "n0stack.pool.v0.ApplyNetworkRequest.AnnotationsEntry")
	proto.RegisterMapType((map[string]string)(nil), "n0stack.pool.v0.ApplyNetworkRequest.LabelsEntry")
	proto.RegisterType((*DeleteNetworkRequest)(nil), "n0stack.pool.v0.DeleteNetworkRequest")
	proto.RegisterType((*ReserveNetworkInterfaceRequest)(nil), "n0stack.pool.v0.ReserveNetworkInterfaceRequest")
	proto.RegisterMapType((map[string]string)(nil), "n0stack.pool.v0.ReserveNetworkInterfaceRequest.AnnotationsEntry")
	proto.RegisterType((*ReleaseNetworkInterfaceRequest)(nil), "n0stack.pool.v0.ReleaseNetworkInterfaceRequest")
}

func init() { proto.RegisterFile("n0stack/pool/v0/network.proto", fileDescriptor_462186787d56a6a3) }

var fileDescriptor_462186787d56a6a3 = []byte{
	// 916 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x56, 0x5f, 0x6f, 0xdb, 0x54,
	0x14, 0xaf, 0x93, 0xa5, 0x5d, 0x4f, 0xb2, 0x35, 0xbb, 0x2d, 0xad, 0x71, 0x61, 0xea, 0x0c, 0x13,
	0x6d, 0x45, 0xed, 0x28, 0x2d, 0x65, 0xeb, 0x10, 0x22, 0xeb, 0x02, 0x54, 0x8d, 0x32, 0x94, 0x31,
	0x2a, 0xf1, 0x12, 0xdd, 0xc4, 0xb7, 0xae, 0x57, 0xc7, 0x76, 0x7d, 0x6f, 0x12, 0x45, 0x88, 0x17,
	0x3e, 0x00, 0x0f, 0xf0, 0x04, 0xaf, 0x7c, 0x24, 0x3e, 0x02, 0x7c, 0x10, 0xe4, 0xeb, 0x6b, 0xd7,
	0x8d, 0xed, 0x64, 0x13, 0xd2, 0xf6, 0xe4, 0x7b, 0xcf, 0x9f, 0xdf, 0x39, 0xe7, 0x9e, 0x3f, 0x3e,
	0xf0, 0xa1, 0x53, 0xa3, 0x0c, 0xf7, 0x2f, 0x75, 0xcf, 0x75, 0x6d, 0x7d, 0x54, 0xd3, 0x1d, 0xc2,
	0xc6, 0xae, 0x7f, 0xa9, 0x79, 0xbe, 0xcb, 0x5c, 0xb4, 0x22, 0xd8, 0x5a, 0xc0, 0xd6, 0x46, 0x35,
	0xe5, 0x03, 0xd3, 0x75, 0x4d, 0x9b, 0xe8, 0xd8, 0xb3, 0x74, 0xec, 0x38, 0x2e, 0xc3, 0xcc, 0x72,
	0x1d, 0x1a, 0x8a, 0x2b, 0x9b, 0x82, 0xcb, 0x6f, 0xbd, 0xe1, 0xb9, 0x4e, 0x06, 0x1e, 0x9b, 0x08,
	0xe6, 0x4e, 0x64, 0xaa, 0x37, 0x34, 0x4c, 0xc2, 0x12, 0xc6, 0xba, 0x96, 0xc3, 0x88, 0x7f, 0x8e,
	0xfb, 0x44, 0x88, 0x7e, 0xca, 0x3f, 0xfd, 0x3d, 0x93, 0x38, 0x7b, 0x74, 0x8c, 0x4d, 0x93, 0xf8,
	0xba, 0xeb, 0x71, 0x4b, 0x69, 0xab, 0xea, 0x3f, 0x25, 0x58, 0x6a, 0x87, 0x48, 0x08, 0xc1, 0x2d,
	0x07, 0x0f, 0x88, 0x2c, 0x6d, 0x49, 0xdb, 0xcb, 0x1d, 0x7e, 0x46, 0xa7, 0x50, 0x4e, 0x28, 0xc9,
	0xc5, 0xad, 0xe2, 0x76, 0xb9, 0xbe, 0xa3, 0x4d, 0x85, 0xa6, 0x09, 0x08, 0xad, 0x71, 0x2d, 0xdb,
	0x74, 0x98, 0x3f, 0xe9, 0x24, 0xb5, 0xd1, 0x17, 0xb0, 0x68, 0xe3, 0x1e, 0xb1, 0xa9, 0x7c, 0x8b,
	0xe3, 0x7c, 0x9c, 0x8b, 0xd3, 0xe2, 0x62, 0x21, 0x84, 0xd0, 0x41, 0x9b, 0xb0, 0x6c, 0x79, 0xa3,
	0x83, 0x6e, 0xdf, 0x32, 0x7c, 0x19, 0xb8, 0x8f, 0xb7, 0x03, 0xc2, 0xb1, 0x65, 0xf8, 0x82, 0x79,
	0x18, 0x32, 0xcb, 0x31, 0xf3, 0x90, 0x33, 0xd7, 0x61, 0xd1, 0x70, 0x07, 0xd8, 0x72, 0xe4, 0x0a,
	0xe7, 0x88, 0x1b, 0x7a, 0x02, 0x25, 0xca, 0x30, 0x23, 0x72, 0x7d, 0x4b, 0xda, 0xbe, 0x5b, 0x7f,
	0x98, 0xeb, 0x8e, 0xf8, 0xbe, 0x08, 0x84, 0x3b, 0xa1, 0x0e, 0x1a, 0xc3, 0xa6, 0x4f, 0x28, 0xf1,
	0x47, 0xc4, 0xe8, 0xa6, 0x72, 0x41, 0xe5, 0x7d, 0x1e, 0xe1, 0xe7, 0xb9, 0x90, 0x1d, 0xa1, 0x2b,
	0xee, 0x27, 0xb1, 0x66, 0x18, 0xf4, 0xfb, 0x7e, 0x1e, 0x5f, 0xf9, 0x12, 0xaa, 0xd3, 0xcf, 0x8c,
	0xaa, 0x50, 0xbc, 0x24, 0x13, 0x91, 0xb9, 0xe0, 0x88, 0xd6, 0xa0, 0x34, 0xc2, 0xf6, 0x90, 0xc8,
	0x05, 0x4e, 0x0b, 0x2f, 0x47, 0x85, 0x47, 0x92, 0xf2, 0x18, 0xca, 0x89, 0xe7, 0x7d, 0x23, 0xd5,
	0x2b, 0xb8, 0x3f, 0xdb, 0xef, 0x0c, 0xb4, 0xc7, 0x49, 0xb4, 0x72, 0xfd, 0xa3, 0xf8, 0x45, 0xc2,
	0x52, 0x4e, 0xbc, 0x49, 0x8c, 0x95, 0x30, 0xa9, 0x1e, 0x43, 0x25, 0xf9, 0xfa, 0x68, 0x03, 0x56,
	0xdb, 0xcd, 0xef, 0xcf, 0x9e, 0x77, 0x4e, 0xbb, 0x2f, 0xdb, 0x2f, 0xbe, 0x6b, 0x1e, 0x9f, 0x7c,
	0x7d, 0xd2, 0x7c, 0x56, 0x5d, 0x40, 0x65, 0x58, 0x7a, 0xd9, 0x3e, 0x6d, 0x3f, 0x3f, 0x6b, 0x57,
	0x25, 0x74, 0x07, 0x96, 0x1b, 0x3f, 0x34, 0x4e, 0x5a, 0x8d, 0xa7, 0xad, 0x66, 0xb5, 0xa0, 0xbe,
	0x07, 0xab, 0x2d, 0x8b, 0x32, 0x01, 0x44, 0x3b, 0xe4, 0x6a, 0x48, 0x28, 0x53, 0x5b, 0xb0, 0x76,
	0x93, 0x4c, 0x3d, 0xd7, 0xa1, 0x04, 0x1d, 0xc0, 0x6d, 0x91, 0x51, 0x2a, 0x4b, 0x3c, 0x8f, 0x72,
	0x5e, 0x1e, 0x3b, 0xb1, 0xa4, 0xfa, 0x09, 0xdc, 0xfb, 0x86, 0x44, 0x60, 0xc2, 0x44, 0x56, 0x4f,
	0xa9, 0x7f, 0x14, 0x61, 0xb5, 0xe1, 0x79, 0xf6, 0x64, 0xbe, 0x2c, 0x3a, 0xcb, 0xea, 0xbf, 0xcf,
	0x52, 0xde, 0x64, 0xc0, 0xcd, 0xe9, 0xc5, 0x6f, 0xa7, 0x7a, 0xb1, 0xf6, 0x5a, 0x98, 0x6f, 0xa5,
	0x2f, 0xdf, 0x61, 0x85, 0xab, 0xbb, 0xb0, 0xf6, 0x8c, 0xd8, 0x84, 0x91, 0xd7, 0xc8, 0xe3, 0xaf,
	0xc5, 0xb8, 0x1d, 0x52, 0x15, 0x2c, 0xd4, 0x1e, 0x40, 0x25, 0x9a, 0x0d, 0x09, 0xf5, 0xb2, 0xa0,
	0xb5, 0x83, 0x0c, 0x1f, 0xc0, 0x7a, 0x6a, 0x7c, 0x84, 0xc2, 0xa1, 0x73, 0x6b, 0xce, 0x14, 0x36,
	0xd7, 0xea, 0x65, 0xd5, 0xc5, 0x57, 0xa9, 0x1c, 0xce, 0x76, 0x6f, 0x4e, 0x89, 0xec, 0x40, 0xf5,
	0x02, 0xfb, 0xc6, 0x18, 0xfb, 0xa4, 0x8b, 0x0d, 0xc3, 0x27, 0x34, 0x28, 0x96, 0xc0, 0xa7, 0x95,
	0x88, 0xde, 0x08, 0xc9, 0x41, 0x9c, 0xbc, 0x06, 0x22, 0xb1, 0x52, 0x18, 0x67, 0x40, 0xbb, 0x29,
	0x72, 0x18, 0x8b, 0x2c, 0xc6, 0x22, 0x87, 0x42, 0xe4, 0xff, 0xe6, 0x5d, 0x9d, 0x04, 0xf9, 0xb0,
	0x09, 0xa6, 0x6f, 0x3d, 0x1f, 0xf5, 0xbf, 0x4a, 0x70, 0x37, 0x9a, 0x53, 0xc4, 0x1f, 0x59, 0x7d,
	0x82, 0x18, 0x54, 0x92, 0xd3, 0x05, 0xa5, 0xff, 0x76, 0x19, 0x33, 0x49, 0x79, 0x38, 0x47, 0x2a,
	0x1c, 0x51, 0xea, 0xc6, 0x2f, 0x7f, 0xff, 0xfb, 0x7b, 0xe1, 0x1e, 0x5a, 0xe1, 0xdb, 0xc4, 0xf5,
	0x3a, 0x80, 0x5e, 0x01, 0x5c, 0x4f, 0x21, 0xa4, 0xa6, 0xd0, 0x52, 0x23, 0x4a, 0xc9, 0x9d, 0x6d,
	0xea, 0x7d, 0x6e, 0x44, 0x46, 0xeb, 0x53, 0x46, 0xf4, 0x9f, 0x82, 0xf7, 0xf8, 0x19, 0x5d, 0x41,
	0x25, 0x39, 0x24, 0x32, 0x22, 0xcc, 0x98, 0x21, 0x33, 0xec, 0x3d, 0xe0, 0xf6, 0x36, 0x95, 0x1c,
	0x7b, 0x47, 0xd2, 0x2e, 0x72, 0xe0, 0xce, 0x8d, 0xfe, 0x44, 0xe9, 0xf7, 0xca, 0xea, 0x5f, 0x65,
	0x5d, 0x0b, 0xd7, 0x2b, 0x2d, 0x5a, 0xaf, 0xb4, 0x66, 0xb0, 0x5e, 0x45, 0x21, 0xee, 0xe6, 0x85,
	0x78, 0x0e, 0x1b, 0x39, 0x3d, 0x84, 0xf4, 0x37, 0xec, 0xb6, 0x19, 0x81, 0x2f, 0x20, 0x23, 0xb0,
	0x93, 0x59, 0xba, 0x99, 0x76, 0x66, 0x15, 0x79, 0x6e, 0xac, 0x0b, 0x4f, 0xff, 0x94, 0x7e, 0x6b,
	0xf4, 0xd0, 0x23, 0x58, 0x12, 0x88, 0xea, 0x5e, 0x7c, 0x44, 0xea, 0x05, 0x63, 0x1e, 0x3d, 0xd2,
	0x75, 0xd3, 0x62, 0x17, 0xc3, 0x9e, 0xd6, 0x77, 0x07, 0x7a, 0xb4, 0x75, 0x8a, 0xef, 0x6e, 0x41,
	0x2a, 0xd4, 0xab, 0xd8, 0xf3, 0x6c, 0xab, 0xcf, 0x7b, 0x54, 0x7f, 0x45, 0x5d, 0xe7, 0x28, 0x45,
	0xf9, 0x71, 0x3f, 0x1f, 0x43, 0x77, 0x6a, 0xdc, 0x2d, 0xcd, 0x74, 0xa3, 0xbd, 0xf9, 0x89, 0x17,
	0x1c, 0x7a, 0x8b, 0x9c, 0xbe, 0xff, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x87, 0xdb, 0xb9, 0x33,
	0x57, 0x0b, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NetworkServiceClient is the client API for NetworkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NetworkServiceClient interface {
	ListNetworks(ctx context.Context, in *ListNetworksRequest, opts ...grpc.CallOption) (*ListNetworksResponse, error)
	GetNetwork(ctx context.Context, in *GetNetworkRequest, opts ...grpc.CallOption) (*Network, error)
	ApplyNetwork(ctx context.Context, in *ApplyNetworkRequest, opts ...grpc.CallOption) (*Network, error)
	DeleteNetwork(ctx context.Context, in *DeleteNetworkRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	ReserveNetworkInterface(ctx context.Context, in *ReserveNetworkInterfaceRequest, opts ...grpc.CallOption) (*Network, error)
	ReleaseNetworkInterface(ctx context.Context, in *ReleaseNetworkInterfaceRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type networkServiceClient struct {
	cc *grpc.ClientConn
}

func NewNetworkServiceClient(cc *grpc.ClientConn) NetworkServiceClient {
	return &networkServiceClient{cc}
}

func (c *networkServiceClient) ListNetworks(ctx context.Context, in *ListNetworksRequest, opts ...grpc.CallOption) (*ListNetworksResponse, error) {
	out := new(ListNetworksResponse)
	err := c.cc.Invoke(ctx, "/n0stack.pool.v0.NetworkService/ListNetworks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServiceClient) GetNetwork(ctx context.Context, in *GetNetworkRequest, opts ...grpc.CallOption) (*Network, error) {
	out := new(Network)
	err := c.cc.Invoke(ctx, "/n0stack.pool.v0.NetworkService/GetNetwork", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServiceClient) ApplyNetwork(ctx context.Context, in *ApplyNetworkRequest, opts ...grpc.CallOption) (*Network, error) {
	out := new(Network)
	err := c.cc.Invoke(ctx, "/n0stack.pool.v0.NetworkService/ApplyNetwork", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServiceClient) DeleteNetwork(ctx context.Context, in *DeleteNetworkRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/n0stack.pool.v0.NetworkService/DeleteNetwork", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServiceClient) ReserveNetworkInterface(ctx context.Context, in *ReserveNetworkInterfaceRequest, opts ...grpc.CallOption) (*Network, error) {
	out := new(Network)
	err := c.cc.Invoke(ctx, "/n0stack.pool.v0.NetworkService/ReserveNetworkInterface", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServiceClient) ReleaseNetworkInterface(ctx context.Context, in *ReleaseNetworkInterfaceRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/n0stack.pool.v0.NetworkService/ReleaseNetworkInterface", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetworkServiceServer is the server API for NetworkService service.
type NetworkServiceServer interface {
	ListNetworks(context.Context, *ListNetworksRequest) (*ListNetworksResponse, error)
	GetNetwork(context.Context, *GetNetworkRequest) (*Network, error)
	ApplyNetwork(context.Context, *ApplyNetworkRequest) (*Network, error)
	DeleteNetwork(context.Context, *DeleteNetworkRequest) (*empty.Empty, error)
	ReserveNetworkInterface(context.Context, *ReserveNetworkInterfaceRequest) (*Network, error)
	ReleaseNetworkInterface(context.Context, *ReleaseNetworkInterfaceRequest) (*empty.Empty, error)
}

// UnimplementedNetworkServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNetworkServiceServer struct {
}

func (*UnimplementedNetworkServiceServer) ListNetworks(ctx context.Context, req *ListNetworksRequest) (*ListNetworksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNetworks not implemented")
}
func (*UnimplementedNetworkServiceServer) GetNetwork(ctx context.Context, req *GetNetworkRequest) (*Network, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNetwork not implemented")
}
func (*UnimplementedNetworkServiceServer) ApplyNetwork(ctx context.Context, req *ApplyNetworkRequest) (*Network, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyNetwork not implemented")
}
func (*UnimplementedNetworkServiceServer) DeleteNetwork(ctx context.Context, req *DeleteNetworkRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNetwork not implemented")
}
func (*UnimplementedNetworkServiceServer) ReserveNetworkInterface(ctx context.Context, req *ReserveNetworkInterfaceRequest) (*Network, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReserveNetworkInterface not implemented")
}
func (*UnimplementedNetworkServiceServer) ReleaseNetworkInterface(ctx context.Context, req *ReleaseNetworkInterfaceRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReleaseNetworkInterface not implemented")
}

func RegisterNetworkServiceServer(s *grpc.Server, srv NetworkServiceServer) {
	s.RegisterService(&_NetworkService_serviceDesc, srv)
}

func _NetworkService_ListNetworks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNetworksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).ListNetworks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.pool.v0.NetworkService/ListNetworks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).ListNetworks(ctx, req.(*ListNetworksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkService_GetNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).GetNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.pool.v0.NetworkService/GetNetwork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).GetNetwork(ctx, req.(*GetNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkService_ApplyNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).ApplyNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.pool.v0.NetworkService/ApplyNetwork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).ApplyNetwork(ctx, req.(*ApplyNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkService_DeleteNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).DeleteNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.pool.v0.NetworkService/DeleteNetwork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).DeleteNetwork(ctx, req.(*DeleteNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkService_ReserveNetworkInterface_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReserveNetworkInterfaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).ReserveNetworkInterface(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.pool.v0.NetworkService/ReserveNetworkInterface",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).ReserveNetworkInterface(ctx, req.(*ReserveNetworkInterfaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkService_ReleaseNetworkInterface_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReleaseNetworkInterfaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).ReleaseNetworkInterface(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.pool.v0.NetworkService/ReleaseNetworkInterface",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).ReleaseNetworkInterface(ctx, req.(*ReleaseNetworkInterfaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NetworkService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "n0stack.pool.v0.NetworkService",
	HandlerType: (*NetworkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListNetworks",
			Handler:    _NetworkService_ListNetworks_Handler,
		},
		{
			MethodName: "GetNetwork",
			Handler:    _NetworkService_GetNetwork_Handler,
		},
		{
			MethodName: "ApplyNetwork",
			Handler:    _NetworkService_ApplyNetwork_Handler,
		},
		{
			MethodName: "DeleteNetwork",
			Handler:    _NetworkService_DeleteNetwork_Handler,
		},
		{
			MethodName: "ReserveNetworkInterface",
			Handler:    _NetworkService_ReserveNetworkInterface_Handler,
		},
		{
			MethodName: "ReleaseNetworkInterface",
			Handler:    _NetworkService_ReleaseNetworkInterface_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "n0stack/pool/v0/network.proto",
}
