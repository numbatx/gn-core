// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: blockV2.proto

package block

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_numbatx_gn_core_data "github.com/numbatx/gn-core/data"
	io "io"
	math "math"
	math_big "math/big"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

// HeaderV2 extends the Header structure with extra fields for version 2
type HeaderV2 struct {
	Header                   *Header       `protobuf:"bytes,1,opt,name=Header,proto3" json:"Header,omitempty"`
	ScheduledRootHash        []byte        `protobuf:"bytes,2,opt,name=ScheduledRootHash,proto3" json:"ScheduledRootHash,omitempty"`
	ScheduledAccumulatedFees *math_big.Int `protobuf:"bytes,3,opt,name=ScheduledAccumulatedFees,proto3,casttypewith=math/big.Int;github.com/numbatx/gn-core/data.BigIntCaster" json:"ScheduledAccumulatedFees,omitempty"`
	ScheduledDeveloperFees   *math_big.Int `protobuf:"bytes,4,opt,name=ScheduledDeveloperFees,proto3,casttypewith=math/big.Int;github.com/numbatx/gn-core/data.BigIntCaster" json:"ScheduledDeveloperFees,omitempty"`
	ScheduledGasProvided     uint64        `protobuf:"varint,5,opt,name=ScheduledGasProvided,proto3" json:"ScheduledGasProvided,omitempty"`
	ScheduledGasPenalized    uint64        `protobuf:"varint,6,opt,name=ScheduledGasPenalized,proto3" json:"ScheduledGasPenalized,omitempty"`
	ScheduledGasRefunded     uint64        `protobuf:"varint,7,opt,name=ScheduledGasRefunded,proto3" json:"ScheduledGasRefunded,omitempty"`
}

func (m *HeaderV2) Reset()      { *m = HeaderV2{} }
func (*HeaderV2) ProtoMessage() {}
func (*HeaderV2) Descriptor() ([]byte, []int) {
	return fileDescriptor_17a3844aa051366e, []int{0}
}
func (m *HeaderV2) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HeaderV2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *HeaderV2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeaderV2.Merge(m, src)
}
func (m *HeaderV2) XXX_Size() int {
	return m.Size()
}
func (m *HeaderV2) XXX_DiscardUnknown() {
	xxx_messageInfo_HeaderV2.DiscardUnknown(m)
}

var xxx_messageInfo_HeaderV2 proto.InternalMessageInfo

func (m *HeaderV2) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HeaderV2) GetScheduledRootHash() []byte {
	if m != nil {
		return m.ScheduledRootHash
	}
	return nil
}

func (m *HeaderV2) GetScheduledAccumulatedFees() *math_big.Int {
	if m != nil {
		return m.ScheduledAccumulatedFees
	}
	return nil
}

func (m *HeaderV2) GetScheduledDeveloperFees() *math_big.Int {
	if m != nil {
		return m.ScheduledDeveloperFees
	}
	return nil
}

func (m *HeaderV2) GetScheduledGasProvided() uint64 {
	if m != nil {
		return m.ScheduledGasProvided
	}
	return 0
}

func (m *HeaderV2) GetScheduledGasPenalized() uint64 {
	if m != nil {
		return m.ScheduledGasPenalized
	}
	return 0
}

func (m *HeaderV2) GetScheduledGasRefunded() uint64 {
	if m != nil {
		return m.ScheduledGasRefunded
	}
	return 0
}

type MiniBlockReserved struct {
	ExecutionType    ProcessingType `protobuf:"varint,1,opt,name=ExecutionType,proto3,enum=proto.ProcessingType" json:"ExecutionType,omitempty"`
	TransactionsType []byte         `protobuf:"bytes,2,opt,name=TransactionsType,proto3" json:"TransactionsType,omitempty"`
}

func (m *MiniBlockReserved) Reset()      { *m = MiniBlockReserved{} }
func (*MiniBlockReserved) ProtoMessage() {}
func (*MiniBlockReserved) Descriptor() ([]byte, []int) {
	return fileDescriptor_17a3844aa051366e, []int{1}
}
func (m *MiniBlockReserved) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MiniBlockReserved) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *MiniBlockReserved) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MiniBlockReserved.Merge(m, src)
}
func (m *MiniBlockReserved) XXX_Size() int {
	return m.Size()
}
func (m *MiniBlockReserved) XXX_DiscardUnknown() {
	xxx_messageInfo_MiniBlockReserved.DiscardUnknown(m)
}

var xxx_messageInfo_MiniBlockReserved proto.InternalMessageInfo

func (m *MiniBlockReserved) GetExecutionType() ProcessingType {
	if m != nil {
		return m.ExecutionType
	}
	return Normal
}

func (m *MiniBlockReserved) GetTransactionsType() []byte {
	if m != nil {
		return m.TransactionsType
	}
	return nil
}

type MiniBlockHeaderReserved struct {
	ExecutionType ProcessingType `protobuf:"varint,1,opt,name=ExecutionType,proto3,enum=proto.ProcessingType" json:"ExecutionType,omitempty"`
	State         MiniBlockState `protobuf:"varint,2,opt,name=State,proto3,enum=proto.MiniBlockState" json:"State,omitempty"`
}

func (m *MiniBlockHeaderReserved) Reset()      { *m = MiniBlockHeaderReserved{} }
func (*MiniBlockHeaderReserved) ProtoMessage() {}
func (*MiniBlockHeaderReserved) Descriptor() ([]byte, []int) {
	return fileDescriptor_17a3844aa051366e, []int{2}
}
func (m *MiniBlockHeaderReserved) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MiniBlockHeaderReserved) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *MiniBlockHeaderReserved) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MiniBlockHeaderReserved.Merge(m, src)
}
func (m *MiniBlockHeaderReserved) XXX_Size() int {
	return m.Size()
}
func (m *MiniBlockHeaderReserved) XXX_DiscardUnknown() {
	xxx_messageInfo_MiniBlockHeaderReserved.DiscardUnknown(m)
}

var xxx_messageInfo_MiniBlockHeaderReserved proto.InternalMessageInfo

func (m *MiniBlockHeaderReserved) GetExecutionType() ProcessingType {
	if m != nil {
		return m.ExecutionType
	}
	return Normal
}

func (m *MiniBlockHeaderReserved) GetState() MiniBlockState {
	if m != nil {
		return m.State
	}
	return Final
}

func init() {
	proto.RegisterType((*HeaderV2)(nil), "proto.HeaderV2")
	proto.RegisterType((*MiniBlockReserved)(nil), "proto.MiniBlockReserved")
	proto.RegisterType((*MiniBlockHeaderReserved)(nil), "proto.MiniBlockHeaderReserved")
}

func init() { proto.RegisterFile("blockV2.proto", fileDescriptor_17a3844aa051366e) }

var fileDescriptor_17a3844aa051366e = []byte{
	// 474 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xcd, 0x6e, 0xd3, 0x40,
	0x14, 0x85, 0x3d, 0xb4, 0x09, 0x68, 0x4a, 0x10, 0xb5, 0x28, 0x58, 0x5d, 0x0c, 0x51, 0x24, 0xa4,
	0x08, 0xa8, 0x2d, 0x19, 0x36, 0xa8, 0x42, 0x88, 0xf0, 0xd7, 0x2e, 0x90, 0x2a, 0xb7, 0xea, 0x82,
	0xdd, 0xd8, 0xbe, 0x75, 0x46, 0x38, 0x33, 0xd1, 0xfc, 0x44, 0x2d, 0x62, 0x81, 0xc4, 0x0b, 0xf0,
	0x18, 0x88, 0x27, 0x61, 0x99, 0x65, 0x76, 0x10, 0x67, 0xc3, 0x8e, 0x3e, 0x02, 0xf2, 0x38, 0x18,
	0xda, 0xb4, 0xbb, 0xac, 0x3c, 0x73, 0xbe, 0x7b, 0xee, 0xb1, 0xa5, 0x63, 0xdc, 0x8a, 0x73, 0x91,
	0xbc, 0x3f, 0x0c, 0xfd, 0xa1, 0x14, 0x5a, 0xb8, 0x0d, 0xfb, 0xd8, 0xdc, 0xca, 0x98, 0xee, 0x9b,
	0xd8, 0x4f, 0xc4, 0x20, 0xc8, 0x44, 0x26, 0x02, 0x2b, 0xc7, 0xe6, 0xc8, 0xde, 0xec, 0xc5, 0x9e,
	0x2a, 0xd7, 0xe6, 0x9a, 0x5d, 0x52, 0x5d, 0x3a, 0xbf, 0x57, 0xf0, 0xb5, 0x1d, 0xa0, 0x29, 0xc8,
	0xc3, 0xd0, 0xbd, 0x87, 0x9b, 0xd5, 0xd9, 0x43, 0x6d, 0xd4, 0x5d, 0x0b, 0x5b, 0xd5, 0x90, 0x5f,
	0x89, 0xd1, 0x1c, 0xba, 0x0f, 0xf1, 0xfa, 0x7e, 0xd2, 0x87, 0xd4, 0xe4, 0x90, 0x46, 0x42, 0xe8,
	0x1d, 0xaa, 0xfa, 0xde, 0x95, 0x36, 0xea, 0x5e, 0x8f, 0x16, 0x81, 0x7b, 0x82, 0xbd, 0x5a, 0x7c,
	0x9e, 0x24, 0x66, 0x60, 0x72, 0xaa, 0x21, 0x7d, 0x0d, 0xa0, 0xbc, 0x95, 0xd2, 0xd4, 0x7b, 0xfa,
	0xed, 0xc7, 0xdd, 0x27, 0x03, 0xaa, 0xfb, 0x41, 0xcc, 0x32, 0x7f, 0x97, 0xeb, 0xed, 0xff, 0x3e,
	0x88, 0x9b, 0x41, 0x4c, 0xf5, 0x71, 0x90, 0xf1, 0xad, 0x44, 0x48, 0x08, 0x52, 0xaa, 0xa9, 0xdf,
	0x63, 0xd9, 0x2e, 0xd7, 0x2f, 0xa8, 0xd2, 0x20, 0xa3, 0x4b, 0xd7, 0xbb, 0x06, 0xdf, 0xae, 0xd9,
	0x4b, 0x18, 0x41, 0x2e, 0x86, 0x20, 0x6d, 0xf0, 0xea, 0x32, 0x82, 0x2f, 0x59, 0xee, 0x86, 0xf8,
	0x56, 0x4d, 0xde, 0x50, 0xb5, 0x27, 0xc5, 0x88, 0xa5, 0x90, 0x7a, 0x8d, 0x36, 0xea, 0xae, 0x46,
	0x17, 0x32, 0xf7, 0x31, 0xde, 0x38, 0xa3, 0x03, 0xa7, 0x39, 0xfb, 0x00, 0xa9, 0xd7, 0xb4, 0xa6,
	0x8b, 0xe1, 0xf9, 0xa4, 0x08, 0x8e, 0x0c, 0x2f, 0x93, 0xae, 0x2e, 0x26, 0xfd, 0x65, 0x9d, 0x8f,
	0x78, 0xfd, 0x2d, 0xe3, 0xac, 0x57, 0x96, 0x20, 0x02, 0x05, 0x72, 0x04, 0xa9, 0xbb, 0x8d, 0x5b,
	0xaf, 0x8e, 0x21, 0x31, 0x9a, 0x09, 0x7e, 0x70, 0x32, 0x04, 0x5b, 0x80, 0x1b, 0xe1, 0xc6, 0xbc,
	0x00, 0x7b, 0x52, 0x24, 0xa0, 0x14, 0xe3, 0x59, 0x09, 0xa3, 0xb3, 0xb3, 0xee, 0x7d, 0x7c, 0xf3,
	0x40, 0x52, 0xae, 0x68, 0x52, 0x4a, 0xca, 0xfa, 0xab, 0x3a, 0x2c, 0xe8, 0x9d, 0xcf, 0x08, 0xdf,
	0xa9, 0xe3, 0xe7, 0xbd, 0x5a, 0xca, 0x4b, 0x3c, 0xc0, 0x8d, 0x7d, 0x4d, 0x75, 0x95, 0xfc, 0xcf,
	0x54, 0x67, 0x59, 0x18, 0x55, 0x33, 0xbd, 0x67, 0xe3, 0x29, 0x71, 0x26, 0x53, 0xe2, 0x9c, 0x4e,
	0x09, 0xfa, 0x54, 0x10, 0xf4, 0xb5, 0x20, 0xe8, 0x7b, 0x41, 0xd0, 0xb8, 0x20, 0x68, 0x52, 0x10,
	0xf4, 0xb3, 0x20, 0xe8, 0x57, 0x41, 0x9c, 0xd3, 0x82, 0xa0, 0x2f, 0x33, 0xe2, 0x8c, 0x67, 0xc4,
	0x99, 0xcc, 0x88, 0xf3, 0xae, 0x61, 0x7f, 0x9e, 0xb8, 0x69, 0xb7, 0x3f, 0xfa, 0x13, 0x00, 0x00,
	0xff, 0xff, 0x70, 0xaf, 0x8b, 0x94, 0x91, 0x03, 0x00, 0x00,
}

func (this *HeaderV2) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HeaderV2)
	if !ok {
		that2, ok := that.(HeaderV2)
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
	if !this.Header.Equal(that1.Header) {
		return false
	}
	if !bytes.Equal(this.ScheduledRootHash, that1.ScheduledRootHash) {
		return false
	}
	{
		__caster := &github_com_numbatx_gn_core_data.BigIntCaster{}
		if !__caster.Equal(this.ScheduledAccumulatedFees, that1.ScheduledAccumulatedFees) {
			return false
		}
	}
	{
		__caster := &github_com_numbatx_gn_core_data.BigIntCaster{}
		if !__caster.Equal(this.ScheduledDeveloperFees, that1.ScheduledDeveloperFees) {
			return false
		}
	}
	if this.ScheduledGasProvided != that1.ScheduledGasProvided {
		return false
	}
	if this.ScheduledGasPenalized != that1.ScheduledGasPenalized {
		return false
	}
	if this.ScheduledGasRefunded != that1.ScheduledGasRefunded {
		return false
	}
	return true
}
func (this *MiniBlockReserved) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MiniBlockReserved)
	if !ok {
		that2, ok := that.(MiniBlockReserved)
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
	if this.ExecutionType != that1.ExecutionType {
		return false
	}
	if !bytes.Equal(this.TransactionsType, that1.TransactionsType) {
		return false
	}
	return true
}
func (this *MiniBlockHeaderReserved) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MiniBlockHeaderReserved)
	if !ok {
		that2, ok := that.(MiniBlockHeaderReserved)
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
	if this.ExecutionType != that1.ExecutionType {
		return false
	}
	if this.State != that1.State {
		return false
	}
	return true
}
func (this *HeaderV2) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 11)
	s = append(s, "&block.HeaderV2{")
	if this.Header != nil {
		s = append(s, "Header: "+fmt.Sprintf("%#v", this.Header)+",\n")
	}
	s = append(s, "ScheduledRootHash: "+fmt.Sprintf("%#v", this.ScheduledRootHash)+",\n")
	s = append(s, "ScheduledAccumulatedFees: "+fmt.Sprintf("%#v", this.ScheduledAccumulatedFees)+",\n")
	s = append(s, "ScheduledDeveloperFees: "+fmt.Sprintf("%#v", this.ScheduledDeveloperFees)+",\n")
	s = append(s, "ScheduledGasProvided: "+fmt.Sprintf("%#v", this.ScheduledGasProvided)+",\n")
	s = append(s, "ScheduledGasPenalized: "+fmt.Sprintf("%#v", this.ScheduledGasPenalized)+",\n")
	s = append(s, "ScheduledGasRefunded: "+fmt.Sprintf("%#v", this.ScheduledGasRefunded)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *MiniBlockReserved) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&block.MiniBlockReserved{")
	s = append(s, "ExecutionType: "+fmt.Sprintf("%#v", this.ExecutionType)+",\n")
	s = append(s, "TransactionsType: "+fmt.Sprintf("%#v", this.TransactionsType)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *MiniBlockHeaderReserved) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&block.MiniBlockHeaderReserved{")
	s = append(s, "ExecutionType: "+fmt.Sprintf("%#v", this.ExecutionType)+",\n")
	s = append(s, "State: "+fmt.Sprintf("%#v", this.State)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringBlockV2(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *HeaderV2) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HeaderV2) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HeaderV2) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ScheduledGasRefunded != 0 {
		i = encodeVarintBlockV2(dAtA, i, uint64(m.ScheduledGasRefunded))
		i--
		dAtA[i] = 0x38
	}
	if m.ScheduledGasPenalized != 0 {
		i = encodeVarintBlockV2(dAtA, i, uint64(m.ScheduledGasPenalized))
		i--
		dAtA[i] = 0x30
	}
	if m.ScheduledGasProvided != 0 {
		i = encodeVarintBlockV2(dAtA, i, uint64(m.ScheduledGasProvided))
		i--
		dAtA[i] = 0x28
	}
	{
		__caster := &github_com_numbatx_gn_core_data.BigIntCaster{}
		size := __caster.Size(m.ScheduledDeveloperFees)
		i -= size
		if _, err := __caster.MarshalTo(m.ScheduledDeveloperFees, dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintBlockV2(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		__caster := &github_com_numbatx_gn_core_data.BigIntCaster{}
		size := __caster.Size(m.ScheduledAccumulatedFees)
		i -= size
		if _, err := __caster.MarshalTo(m.ScheduledAccumulatedFees, dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintBlockV2(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.ScheduledRootHash) > 0 {
		i -= len(m.ScheduledRootHash)
		copy(dAtA[i:], m.ScheduledRootHash)
		i = encodeVarintBlockV2(dAtA, i, uint64(len(m.ScheduledRootHash)))
		i--
		dAtA[i] = 0x12
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBlockV2(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MiniBlockReserved) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MiniBlockReserved) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MiniBlockReserved) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TransactionsType) > 0 {
		i -= len(m.TransactionsType)
		copy(dAtA[i:], m.TransactionsType)
		i = encodeVarintBlockV2(dAtA, i, uint64(len(m.TransactionsType)))
		i--
		dAtA[i] = 0x12
	}
	if m.ExecutionType != 0 {
		i = encodeVarintBlockV2(dAtA, i, uint64(m.ExecutionType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MiniBlockHeaderReserved) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MiniBlockHeaderReserved) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MiniBlockHeaderReserved) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.State != 0 {
		i = encodeVarintBlockV2(dAtA, i, uint64(m.State))
		i--
		dAtA[i] = 0x10
	}
	if m.ExecutionType != 0 {
		i = encodeVarintBlockV2(dAtA, i, uint64(m.ExecutionType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintBlockV2(dAtA []byte, offset int, v uint64) int {
	offset -= sovBlockV2(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *HeaderV2) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovBlockV2(uint64(l))
	}
	l = len(m.ScheduledRootHash)
	if l > 0 {
		n += 1 + l + sovBlockV2(uint64(l))
	}
	{
		__caster := &github_com_numbatx_gn_core_data.BigIntCaster{}
		l = __caster.Size(m.ScheduledAccumulatedFees)
		n += 1 + l + sovBlockV2(uint64(l))
	}
	{
		__caster := &github_com_numbatx_gn_core_data.BigIntCaster{}
		l = __caster.Size(m.ScheduledDeveloperFees)
		n += 1 + l + sovBlockV2(uint64(l))
	}
	if m.ScheduledGasProvided != 0 {
		n += 1 + sovBlockV2(uint64(m.ScheduledGasProvided))
	}
	if m.ScheduledGasPenalized != 0 {
		n += 1 + sovBlockV2(uint64(m.ScheduledGasPenalized))
	}
	if m.ScheduledGasRefunded != 0 {
		n += 1 + sovBlockV2(uint64(m.ScheduledGasRefunded))
	}
	return n
}

func (m *MiniBlockReserved) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ExecutionType != 0 {
		n += 1 + sovBlockV2(uint64(m.ExecutionType))
	}
	l = len(m.TransactionsType)
	if l > 0 {
		n += 1 + l + sovBlockV2(uint64(l))
	}
	return n
}

func (m *MiniBlockHeaderReserved) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ExecutionType != 0 {
		n += 1 + sovBlockV2(uint64(m.ExecutionType))
	}
	if m.State != 0 {
		n += 1 + sovBlockV2(uint64(m.State))
	}
	return n
}

func sovBlockV2(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBlockV2(x uint64) (n int) {
	return sovBlockV2(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *HeaderV2) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&HeaderV2{`,
		`Header:` + strings.Replace(fmt.Sprintf("%v", this.Header), "Header", "Header", 1) + `,`,
		`ScheduledRootHash:` + fmt.Sprintf("%v", this.ScheduledRootHash) + `,`,
		`ScheduledAccumulatedFees:` + fmt.Sprintf("%v", this.ScheduledAccumulatedFees) + `,`,
		`ScheduledDeveloperFees:` + fmt.Sprintf("%v", this.ScheduledDeveloperFees) + `,`,
		`ScheduledGasProvided:` + fmt.Sprintf("%v", this.ScheduledGasProvided) + `,`,
		`ScheduledGasPenalized:` + fmt.Sprintf("%v", this.ScheduledGasPenalized) + `,`,
		`ScheduledGasRefunded:` + fmt.Sprintf("%v", this.ScheduledGasRefunded) + `,`,
		`}`,
	}, "")
	return s
}
func (this *MiniBlockReserved) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&MiniBlockReserved{`,
		`ExecutionType:` + fmt.Sprintf("%v", this.ExecutionType) + `,`,
		`TransactionsType:` + fmt.Sprintf("%v", this.TransactionsType) + `,`,
		`}`,
	}, "")
	return s
}
func (this *MiniBlockHeaderReserved) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&MiniBlockHeaderReserved{`,
		`ExecutionType:` + fmt.Sprintf("%v", this.ExecutionType) + `,`,
		`State:` + fmt.Sprintf("%v", this.State) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringBlockV2(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *HeaderV2) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlockV2
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
			return fmt.Errorf("proto: HeaderV2: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HeaderV2: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockV2
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
				return ErrInvalidLengthBlockV2
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBlockV2
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &Header{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ScheduledRootHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockV2
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
				return ErrInvalidLengthBlockV2
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlockV2
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ScheduledRootHash = append(m.ScheduledRootHash[:0], dAtA[iNdEx:postIndex]...)
			if m.ScheduledRootHash == nil {
				m.ScheduledRootHash = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ScheduledAccumulatedFees", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockV2
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
				return ErrInvalidLengthBlockV2
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlockV2
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			{
				__caster := &github_com_numbatx_gn_core_data.BigIntCaster{}
				if tmp, err := __caster.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
					return err
				} else {
					m.ScheduledAccumulatedFees = tmp
				}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ScheduledDeveloperFees", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockV2
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
				return ErrInvalidLengthBlockV2
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlockV2
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			{
				__caster := &github_com_numbatx_gn_core_data.BigIntCaster{}
				if tmp, err := __caster.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
					return err
				} else {
					m.ScheduledDeveloperFees = tmp
				}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ScheduledGasProvided", wireType)
			}
			m.ScheduledGasProvided = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockV2
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ScheduledGasProvided |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ScheduledGasPenalized", wireType)
			}
			m.ScheduledGasPenalized = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockV2
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ScheduledGasPenalized |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ScheduledGasRefunded", wireType)
			}
			m.ScheduledGasRefunded = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockV2
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ScheduledGasRefunded |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBlockV2(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBlockV2
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
func (m *MiniBlockReserved) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlockV2
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
			return fmt.Errorf("proto: MiniBlockReserved: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MiniBlockReserved: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExecutionType", wireType)
			}
			m.ExecutionType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockV2
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExecutionType |= ProcessingType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TransactionsType", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockV2
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
				return ErrInvalidLengthBlockV2
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlockV2
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TransactionsType = append(m.TransactionsType[:0], dAtA[iNdEx:postIndex]...)
			if m.TransactionsType == nil {
				m.TransactionsType = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBlockV2(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBlockV2
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
func (m *MiniBlockHeaderReserved) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlockV2
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
			return fmt.Errorf("proto: MiniBlockHeaderReserved: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MiniBlockHeaderReserved: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExecutionType", wireType)
			}
			m.ExecutionType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockV2
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExecutionType |= ProcessingType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockV2
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= MiniBlockState(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBlockV2(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBlockV2
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
func skipBlockV2(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBlockV2
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
					return 0, ErrIntOverflowBlockV2
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
					return 0, ErrIntOverflowBlockV2
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
				return 0, ErrInvalidLengthBlockV2
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBlockV2
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBlockV2
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBlockV2        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBlockV2          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBlockV2 = fmt.Errorf("proto: unexpected end of group")
)
