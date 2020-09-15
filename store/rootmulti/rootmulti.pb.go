// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: store/rootmulti/rootmulti.proto

package rootmulti

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_pokt_network_pocket_core_store_types "github.com/pokt-network/pocket-core/store/types"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// CommitInfo defines commit information used by the multi-store when committing
// a version/height.
type CommitInfo struct {
	Version    int64       `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	StoreInfos []StoreInfo `protobuf:"bytes,2,rep,name=store_infos,json=storeInfos,proto3" json:"store_infos"`
}

func (m *CommitInfo) Reset()         { *m = CommitInfo{} }
func (m *CommitInfo) String() string { return proto.CompactTextString(m) }
func (*CommitInfo) ProtoMessage()    {}
func (*CommitInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_35d8ab40086f9a60, []int{0}
}
func (m *CommitInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CommitInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CommitInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CommitInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommitInfo.Merge(m, src)
}
func (m *CommitInfo) XXX_Size() int {
	return m.Size()
}
func (m *CommitInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CommitInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CommitInfo proto.InternalMessageInfo

func (m *CommitInfo) GetVersion() int64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *CommitInfo) GetStoreInfos() []StoreInfo {
	if m != nil {
		return m.StoreInfos
	}
	return nil
}

// StoreInfo defines store-specific commit information. It contains a reference
// between a store name and the commit ID.
type StoreInfo struct {
	Name string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Core StoreCore `protobuf:"bytes,2,opt,name=core,proto3" json:"core"`
}

func (m *StoreInfo) Reset()         { *m = StoreInfo{} }
func (m *StoreInfo) String() string { return proto.CompactTextString(m) }
func (*StoreInfo) ProtoMessage()    {}
func (*StoreInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_35d8ab40086f9a60, []int{1}
}
func (m *StoreInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StoreInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StoreInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StoreInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreInfo.Merge(m, src)
}
func (m *StoreInfo) XXX_Size() int {
	return m.Size()
}
func (m *StoreInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreInfo.DiscardUnknown(m)
}

var xxx_messageInfo_StoreInfo proto.InternalMessageInfo

func (m *StoreInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StoreInfo) GetCore() StoreCore {
	if m != nil {
		return m.Core
	}
	return StoreCore{}
}

type StoreCore struct {
	CommitID github_com_pokt_network_pocket_core_store_types.CommitID `protobuf:"bytes,2,opt,name=commitID,proto3,casttype=github.com/pokt-network/pocket-core/store/types.CommitID" json:"commitID"`
}

func (m *StoreCore) Reset()         { *m = StoreCore{} }
func (m *StoreCore) String() string { return proto.CompactTextString(m) }
func (*StoreCore) ProtoMessage()    {}
func (*StoreCore) Descriptor() ([]byte, []int) {
	return fileDescriptor_35d8ab40086f9a60, []int{2}
}
func (m *StoreCore) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StoreCore) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StoreCore.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StoreCore) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreCore.Merge(m, src)
}
func (m *StoreCore) XXX_Size() int {
	return m.Size()
}
func (m *StoreCore) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreCore.DiscardUnknown(m)
}

var xxx_messageInfo_StoreCore proto.InternalMessageInfo

func (m *StoreCore) GetCommitID() github_com_pokt_network_pocket_core_store_types.CommitID {
	if m != nil {
		return m.CommitID
	}
	return github_com_pokt_network_pocket_core_store_types.CommitID{}
}

type CommitID struct {
	Version int64  `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Hash    []byte `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (m *CommitID) Reset()         { *m = CommitID{} }
func (m *CommitID) String() string { return proto.CompactTextString(m) }
func (*CommitID) ProtoMessage()    {}
func (*CommitID) Descriptor() ([]byte, []int) {
	return fileDescriptor_35d8ab40086f9a60, []int{3}
}
func (m *CommitID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CommitID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CommitID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CommitID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommitID.Merge(m, src)
}
func (m *CommitID) XXX_Size() int {
	return m.Size()
}
func (m *CommitID) XXX_DiscardUnknown() {
	xxx_messageInfo_CommitID.DiscardUnknown(m)
}

var xxx_messageInfo_CommitID proto.InternalMessageInfo

func (m *CommitID) GetVersion() int64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *CommitID) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type MultiStoreProof struct {
	StoreInfos []StoreInfo `protobuf:"bytes,1,rep,name=storeInfos,proto3" json:"storeInfos"`
}

func (m *MultiStoreProof) Reset()         { *m = MultiStoreProof{} }
func (m *MultiStoreProof) String() string { return proto.CompactTextString(m) }
func (*MultiStoreProof) ProtoMessage()    {}
func (*MultiStoreProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_35d8ab40086f9a60, []int{4}
}
func (m *MultiStoreProof) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MultiStoreProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MultiStoreProof.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MultiStoreProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiStoreProof.Merge(m, src)
}
func (m *MultiStoreProof) XXX_Size() int {
	return m.Size()
}
func (m *MultiStoreProof) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiStoreProof.DiscardUnknown(m)
}

var xxx_messageInfo_MultiStoreProof proto.InternalMessageInfo

func (m *MultiStoreProof) GetStoreInfos() []StoreInfo {
	if m != nil {
		return m.StoreInfos
	}
	return nil
}

type MultiStoreProofOp struct {
	Key   []byte           `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Proof *MultiStoreProof `protobuf:"bytes,2,opt,name=proof,proto3" json:"proof"`
}

func (m *MultiStoreProofOp) Reset()         { *m = MultiStoreProofOp{} }
func (m *MultiStoreProofOp) String() string { return proto.CompactTextString(m) }
func (*MultiStoreProofOp) ProtoMessage()    {}
func (*MultiStoreProofOp) Descriptor() ([]byte, []int) {
	return fileDescriptor_35d8ab40086f9a60, []int{5}
}
func (m *MultiStoreProofOp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MultiStoreProofOp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MultiStoreProofOp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MultiStoreProofOp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiStoreProofOp.Merge(m, src)
}
func (m *MultiStoreProofOp) XXX_Size() int {
	return m.Size()
}
func (m *MultiStoreProofOp) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiStoreProofOp.DiscardUnknown(m)
}

var xxx_messageInfo_MultiStoreProofOp proto.InternalMessageInfo

func (m *MultiStoreProofOp) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *MultiStoreProofOp) GetProof() *MultiStoreProof {
	if m != nil {
		return m.Proof
	}
	return nil
}

type Version struct {
	Version int64 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
}

func (m *Version) Reset()         { *m = Version{} }
func (m *Version) String() string { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()    {}
func (*Version) Descriptor() ([]byte, []int) {
	return fileDescriptor_35d8ab40086f9a60, []int{6}
}
func (m *Version) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Version) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Version.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Version) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Version.Merge(m, src)
}
func (m *Version) XXX_Size() int {
	return m.Size()
}
func (m *Version) XXX_DiscardUnknown() {
	xxx_messageInfo_Version.DiscardUnknown(m)
}

var xxx_messageInfo_Version proto.InternalMessageInfo

func (m *Version) GetVersion() int64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func init() {
	proto.RegisterType((*CommitInfo)(nil), "store.rootmulti.CommitInfo")
	proto.RegisterType((*StoreInfo)(nil), "store.rootmulti.StoreInfo")
	proto.RegisterType((*StoreCore)(nil), "store.rootmulti.StoreCore")
	proto.RegisterType((*CommitID)(nil), "store.rootmulti.CommitID")
	proto.RegisterType((*MultiStoreProof)(nil), "store.rootmulti.MultiStoreProof")
	proto.RegisterType((*MultiStoreProofOp)(nil), "store.rootmulti.MultiStoreProofOp")
	proto.RegisterType((*Version)(nil), "store.rootmulti.Version")
}

func init() { proto.RegisterFile("store/rootmulti/rootmulti.proto", fileDescriptor_35d8ab40086f9a60) }

var fileDescriptor_35d8ab40086f9a60 = []byte{
	// 417 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x31, 0xaf, 0xd2, 0x50,
	0x14, 0xc7, 0x7b, 0x1f, 0xd5, 0x07, 0xe7, 0xbd, 0xe4, 0xe9, 0x8d, 0x43, 0x7d, 0x43, 0xdb, 0x74,
	0x7a, 0x0b, 0x6d, 0x22, 0x0e, 0x86, 0xc1, 0x40, 0x61, 0x71, 0x20, 0x9a, 0x12, 0x1d, 0x5c, 0x0c,
	0x34, 0x17, 0x68, 0x6a, 0x7b, 0x9a, 0xdb, 0x8b, 0x84, 0xdd, 0x0f, 0xe0, 0xe8, 0xc8, 0xc7, 0x61,
	0x64, 0x74, 0x22, 0x06, 0x16, 0xe3, 0x47, 0x70, 0x32, 0xf7, 0xb6, 0x54, 0x03, 0xc1, 0x98, 0x37,
	0xf5, 0xdc, 0x7b, 0x4e, 0x7f, 0xff, 0xf3, 0xff, 0xb7, 0x60, 0xe5, 0x02, 0x39, 0xf3, 0x38, 0xa2,
	0x48, 0xe6, 0x1f, 0x45, 0xf4, 0xa7, 0x72, 0x33, 0x8e, 0x02, 0xe9, 0x8d, 0x1a, 0x70, 0xab, 0xeb,
	0xdb, 0x27, 0x53, 0x9c, 0xa2, 0xea, 0x79, 0xb2, 0x2a, 0xc6, 0x9c, 0x08, 0xa0, 0x87, 0x49, 0x12,
	0x89, 0x57, 0xe9, 0x04, 0xa9, 0x01, 0x97, 0x9f, 0x18, 0xcf, 0x23, 0x4c, 0x0d, 0x62, 0x93, 0xbb,
	0x5a, 0x70, 0x38, 0xd2, 0x2e, 0x5c, 0x29, 0xe0, 0x87, 0x28, 0x9d, 0x60, 0x6e, 0x5c, 0xd8, 0xb5,
	0xbb, 0xab, 0x67, 0xb7, 0xee, 0x91, 0x88, 0x3b, 0x94, 0x67, 0x89, 0xf2, 0xf5, 0xf5, 0xd6, 0xd2,
	0x02, 0xc8, 0x0f, 0x17, 0xb9, 0xf3, 0x16, 0x1a, 0x55, 0x9b, 0x52, 0xd0, 0xd3, 0x51, 0xc2, 0x94,
	0x4c, 0x23, 0x50, 0x35, 0x7d, 0x0e, 0x7a, 0x88, 0x9c, 0x19, 0x17, 0x36, 0x39, 0x0f, 0xef, 0x21,
	0x67, 0x25, 0x5c, 0x4d, 0x3b, 0x9f, 0x49, 0xc9, 0x95, 0x1d, 0xba, 0x80, 0x7a, 0x58, 0xf8, 0xe9,
	0x97, 0x9c, 0xa7, 0x27, 0x9c, 0xd2, 0x70, 0xdf, 0xef, 0x48, 0xcc, 0xaf, 0xad, 0xf5, 0x62, 0x1a,
	0x89, 0xd9, 0x7c, 0xec, 0x86, 0x98, 0x78, 0x19, 0xc6, 0xa2, 0x99, 0x32, 0xb1, 0x40, 0x1e, 0x7b,
	0x19, 0x86, 0x31, 0x13, 0x4d, 0xa9, 0xe3, 0x15, 0x81, 0x8b, 0x65, 0xc6, 0xf2, 0x8a, 0x10, 0x54,
	0x62, 0xce, 0x4b, 0xa8, 0x1f, 0x6e, 0xff, 0x11, 0x23, 0x05, 0x7d, 0x36, 0xca, 0x67, 0x6a, 0xb5,
	0xeb, 0x40, 0xd5, 0x6d, 0xfd, 0xeb, 0xca, 0x22, 0xce, 0x10, 0x6e, 0x06, 0x72, 0x3b, 0x65, 0xe5,
	0x0d, 0x47, 0x9c, 0xd0, 0x0e, 0xfc, 0x15, 0x9f, 0x41, 0xee, 0x11, 0x79, 0x0a, 0x8f, 0x8f, 0xa0,
	0xaf, 0x33, 0xfa, 0x08, 0x6a, 0x31, 0x5b, 0xaa, 0xcd, 0xae, 0x03, 0x59, 0xd2, 0x2e, 0x3c, 0xc8,
	0x64, 0xb3, 0x4c, 0xcc, 0x3e, 0xd1, 0x38, 0x82, 0xf8, 0x8d, 0x9f, 0x5b, 0xab, 0x78, 0x25, 0x28,
	0x1e, 0xa5, 0x89, 0x26, 0x5c, 0xbe, 0x2b, 0x9d, 0x9e, 0xcd, 0xa0, 0x5d, 0x97, 0xa3, 0x3f, 0x56,
	0x16, 0xf1, 0x07, 0xeb, 0x9d, 0x49, 0x36, 0x3b, 0x93, 0x7c, 0xdf, 0x99, 0xe4, 0xcb, 0xde, 0xd4,
	0x36, 0x7b, 0x53, 0xfb, 0xb6, 0x37, 0xb5, 0xf7, 0xad, 0xff, 0xff, 0x22, 0xd5, 0x96, 0xe3, 0x87,
	0xea, 0x97, 0x6e, 0xfd, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x30, 0xce, 0x4b, 0x70, 0x1c, 0x03, 0x00,
	0x00,
}

func (this *Version) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Version)
	if !ok {
		that2, ok := that.(Version)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Version != that1.Version {
		return false
	}
	return true
}
func (m *CommitInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CommitInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CommitInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.StoreInfos) > 0 {
		for iNdEx := len(m.StoreInfos) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.StoreInfos[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRootmulti(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Version != 0 {
		i = encodeVarintRootmulti(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *StoreInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StoreInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StoreInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Core.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintRootmulti(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintRootmulti(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *StoreCore) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StoreCore) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StoreCore) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.CommitID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintRootmulti(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	return len(dAtA) - i, nil
}

func (m *CommitID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CommitID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CommitID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintRootmulti(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0x12
	}
	if m.Version != 0 {
		i = encodeVarintRootmulti(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MultiStoreProof) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MultiStoreProof) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MultiStoreProof) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.StoreInfos) > 0 {
		for iNdEx := len(m.StoreInfos) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.StoreInfos[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRootmulti(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *MultiStoreProofOp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MultiStoreProofOp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MultiStoreProofOp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Proof != nil {
		{
			size, err := m.Proof.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRootmulti(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintRootmulti(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Version) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Version) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Version) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Version != 0 {
		i = encodeVarintRootmulti(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintRootmulti(dAtA []byte, offset int, v uint64) int {
	offset -= sovRootmulti(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CommitInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Version != 0 {
		n += 1 + sovRootmulti(uint64(m.Version))
	}
	if len(m.StoreInfos) > 0 {
		for _, e := range m.StoreInfos {
			l = e.Size()
			n += 1 + l + sovRootmulti(uint64(l))
		}
	}
	return n
}

func (m *StoreInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovRootmulti(uint64(l))
	}
	l = m.Core.Size()
	n += 1 + l + sovRootmulti(uint64(l))
	return n
}

func (m *StoreCore) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.CommitID.Size()
	n += 1 + l + sovRootmulti(uint64(l))
	return n
}

func (m *CommitID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Version != 0 {
		n += 1 + sovRootmulti(uint64(m.Version))
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovRootmulti(uint64(l))
	}
	return n
}

func (m *MultiStoreProof) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.StoreInfos) > 0 {
		for _, e := range m.StoreInfos {
			l = e.Size()
			n += 1 + l + sovRootmulti(uint64(l))
		}
	}
	return n
}

func (m *MultiStoreProofOp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovRootmulti(uint64(l))
	}
	if m.Proof != nil {
		l = m.Proof.Size()
		n += 1 + l + sovRootmulti(uint64(l))
	}
	return n
}

func (m *Version) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Version != 0 {
		n += 1 + sovRootmulti(uint64(m.Version))
	}
	return n
}

func sovRootmulti(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRootmulti(x uint64) (n int) {
	return sovRootmulti(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CommitInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRootmulti
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CommitInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CommitInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRootmulti
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StoreInfos", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRootmulti
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRootmulti
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRootmulti
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StoreInfos = append(m.StoreInfos, StoreInfo{})
			if err := m.StoreInfos[len(m.StoreInfos)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRootmulti(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRootmulti
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRootmulti
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StoreInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRootmulti
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StoreInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StoreInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRootmulti
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRootmulti
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRootmulti
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Core", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRootmulti
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRootmulti
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRootmulti
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Core.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRootmulti(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRootmulti
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRootmulti
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StoreCore) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRootmulti
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StoreCore: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StoreCore: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommitID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRootmulti
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRootmulti
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRootmulti
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CommitID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRootmulti(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRootmulti
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRootmulti
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CommitID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRootmulti
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CommitID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CommitID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRootmulti
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRootmulti
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRootmulti
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRootmulti
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = append(m.Hash[:0], dAtA[iNdEx:postIndex]...)
			if m.Hash == nil {
				m.Hash = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRootmulti(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRootmulti
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRootmulti
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MultiStoreProof) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRootmulti
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MultiStoreProof: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MultiStoreProof: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StoreInfos", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRootmulti
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRootmulti
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRootmulti
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StoreInfos = append(m.StoreInfos, StoreInfo{})
			if err := m.StoreInfos[len(m.StoreInfos)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRootmulti(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRootmulti
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRootmulti
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MultiStoreProofOp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRootmulti
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MultiStoreProofOp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MultiStoreProofOp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRootmulti
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRootmulti
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRootmulti
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proof", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRootmulti
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRootmulti
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRootmulti
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Proof == nil {
				m.Proof = &MultiStoreProof{}
			}
			if err := m.Proof.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRootmulti(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRootmulti
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRootmulti
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Version) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRootmulti
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Version: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Version: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRootmulti
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRootmulti(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRootmulti
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRootmulti
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipRootmulti(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRootmulti
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRootmulti
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRootmulti
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthRootmulti
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRootmulti
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRootmulti
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRootmulti        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRootmulti          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRootmulti = fmt.Errorf("proto: unexpected end of group")
)
