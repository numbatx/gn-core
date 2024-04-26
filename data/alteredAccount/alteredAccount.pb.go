// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: alteredAccount.proto

package alteredAccount

import (
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

type AlteredAccount struct {
	Address string              `protobuf:"bytes,1,opt,name=Address,proto3" json:"Address,omitempty"`
	Nonce   uint64              `protobuf:"varint,2,opt,name=Nonce,proto3" json:"Nonce,omitempty"`
	Balance *math_big.Int       `protobuf:"bytes,3,opt,name=Balance,proto3,casttypewith=math/big.Int;github.com/numbatx/gn-core/data.BigIntCaster" json:"Balance,omitempty"`
	Tokens  []*AccountTokenData `protobuf:"bytes,4,rep,name=Tokens,proto3" json:"Tokens,omitempty"`
}

func (m *AlteredAccount) Reset()      { *m = AlteredAccount{} }
func (*AlteredAccount) ProtoMessage() {}
func (*AlteredAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_804e04a1cde31bca, []int{0}
}
func (m *AlteredAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AlteredAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *AlteredAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlteredAccount.Merge(m, src)
}
func (m *AlteredAccount) XXX_Size() int {
	return m.Size()
}
func (m *AlteredAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_AlteredAccount.DiscardUnknown(m)
}

var xxx_messageInfo_AlteredAccount proto.InternalMessageInfo

func (m *AlteredAccount) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *AlteredAccount) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *AlteredAccount) GetBalance() *math_big.Int {
	if m != nil {
		return m.Balance
	}
	return nil
}

func (m *AlteredAccount) GetTokens() []*AccountTokenData {
	if m != nil {
		return m.Tokens
	}
	return nil
}

type AccountTokenData struct {
	Nonce      uint64        `protobuf:"varint,1,opt,name=Nonce,proto3" json:"Nonce,omitempty"`
	Identifier string        `protobuf:"bytes,2,opt,name=Identifier,proto3" json:"Identifier,omitempty"`
	Balance    *math_big.Int `protobuf:"bytes,3,opt,name=Balance,proto3,casttypewith=math/big.Int;github.com/numbatx/gn-core/data.BigIntCaster" json:"Balance,omitempty"`
	Properties string        `protobuf:"bytes,4,opt,name=Properties,proto3" json:"Properties,omitempty"`
}

func (m *AccountTokenData) Reset()      { *m = AccountTokenData{} }
func (*AccountTokenData) ProtoMessage() {}
func (*AccountTokenData) Descriptor() ([]byte, []int) {
	return fileDescriptor_804e04a1cde31bca, []int{1}
}
func (m *AccountTokenData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AccountTokenData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *AccountTokenData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountTokenData.Merge(m, src)
}
func (m *AccountTokenData) XXX_Size() int {
	return m.Size()
}
func (m *AccountTokenData) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountTokenData.DiscardUnknown(m)
}

var xxx_messageInfo_AccountTokenData proto.InternalMessageInfo

func (m *AccountTokenData) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *AccountTokenData) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *AccountTokenData) GetBalance() *math_big.Int {
	if m != nil {
		return m.Balance
	}
	return nil
}

func (m *AccountTokenData) GetProperties() string {
	if m != nil {
		return m.Properties
	}
	return ""
}

func init() {
	proto.RegisterType((*AlteredAccount)(nil), "proto.AlteredAccount")
	proto.RegisterType((*AccountTokenData)(nil), "proto.AccountTokenData")
}

func init() { proto.RegisterFile("alteredAccount.proto", fileDescriptor_804e04a1cde31bca) }

var fileDescriptor_804e04a1cde31bca = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0x3f, 0x6b, 0xe3, 0x30,
	0x18, 0xc6, 0xfd, 0x5e, 0xfe, 0x11, 0xdd, 0x11, 0x0e, 0x13, 0x38, 0x73, 0xc3, 0x7b, 0x26, 0x93,
	0x97, 0xd8, 0x70, 0x37, 0x1d, 0xa5, 0x43, 0xdc, 0x0e, 0xcd, 0x52, 0x8a, 0x29, 0x14, 0xba, 0xc9,
	0xb6, 0xe2, 0x98, 0x26, 0x52, 0x90, 0x65, 0xe8, 0xd8, 0x8f, 0xd0, 0x8f, 0x51, 0xfa, 0x2d, 0xba,
	0x95, 0x4e, 0x19, 0xb3, 0xb5, 0x51, 0x96, 0x8e, 0xf9, 0x08, 0x25, 0x4a, 0x42, 0x92, 0xee, 0x9d,
	0xa4, 0xe7, 0x79, 0xa5, 0x57, 0xbf, 0x47, 0xbc, 0xa4, 0x4d, 0x47, 0x8a, 0x49, 0x96, 0xf6, 0x92,
	0x44, 0x94, 0x5c, 0xf9, 0x13, 0x29, 0x94, 0xb0, 0x6b, 0x66, 0xf9, 0xdd, 0xcd, 0x72, 0x35, 0x2c,
	0x63, 0x3f, 0x11, 0xe3, 0x20, 0x13, 0x99, 0x08, 0x8c, 0x1d, 0x97, 0x03, 0xa3, 0x8c, 0x30, 0xbb,
	0xf5, 0xad, 0xce, 0x0b, 0x90, 0x56, 0xef, 0xa0, 0x9d, 0xed, 0x90, 0x46, 0x2f, 0x4d, 0x25, 0x2b,
	0x0a, 0x07, 0x5c, 0xf0, 0x9a, 0xd1, 0x56, 0xda, 0x6d, 0x52, 0x3b, 0x17, 0x3c, 0x61, 0xce, 0x37,
	0x17, 0xbc, 0x6a, 0xb4, 0x16, 0xf6, 0x15, 0x69, 0x84, 0x74, 0x44, 0x57, 0x7e, 0xc5, 0x05, 0xef,
	0x47, 0x78, 0xfc, 0xf8, 0xfa, 0xe7, 0xff, 0x98, 0xaa, 0x61, 0x10, 0xe7, 0x99, 0xdf, 0xe7, 0xea,
	0x68, 0x8f, 0x89, 0x97, 0xe3, 0x98, 0xaa, 0xdb, 0x20, 0xe3, 0xdd, 0x44, 0x48, 0x16, 0xa4, 0x54,
	0x51, 0x3f, 0xcc, 0xb3, 0x3e, 0x57, 0x27, 0xb4, 0x50, 0x4c, 0x46, 0xdb, 0x6e, 0x76, 0x40, 0xea,
	0x97, 0xe2, 0x86, 0xf1, 0xc2, 0xa9, 0xba, 0x15, 0xef, 0xfb, 0xdf, 0x5f, 0x6b, 0x66, 0x7f, 0x03,
	0x6a, 0x6a, 0xa7, 0x54, 0xd1, 0x68, 0x73, 0xac, 0xf3, 0x04, 0xe4, 0xe7, 0xe7, 0xe2, 0x0e, 0x1a,
	0xf6, 0xa1, 0x91, 0x90, 0x7e, 0xca, 0xb8, 0xca, 0x07, 0x39, 0x93, 0x26, 0x4f, 0x33, 0xda, 0x73,
	0xbe, 0x2e, 0x14, 0x12, 0x72, 0x21, 0xc5, 0x84, 0x49, 0x95, 0xb3, 0x55, 0x30, 0xf3, 0xf0, 0xce,
	0x09, 0xcf, 0xa6, 0x73, 0xb4, 0x66, 0x73, 0xb4, 0x96, 0x73, 0x84, 0x3b, 0x8d, 0xf0, 0xa0, 0x11,
	0x9e, 0x35, 0xc2, 0x54, 0x23, 0xcc, 0x34, 0xc2, 0x9b, 0x46, 0x78, 0xd7, 0x68, 0x2d, 0x35, 0xc2,
	0xfd, 0x02, 0xad, 0xe9, 0x02, 0xad, 0xd9, 0x02, 0xad, 0xeb, 0xd6, 0xe1, 0x58, 0xc4, 0x75, 0xf3,
	0x5b, 0xff, 0x3e, 0x02, 0x00, 0x00, 0xff, 0xff, 0xae, 0xcc, 0x3d, 0x43, 0x2f, 0x02, 0x00, 0x00,
}

func (this *AlteredAccount) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AlteredAccount)
	if !ok {
		that2, ok := that.(AlteredAccount)
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
	if this.Address != that1.Address {
		return false
	}
	if this.Nonce != that1.Nonce {
		return false
	}
	{
		__caster := &github_com_numbatx_gn_core_data.BigIntCaster{}
		if !__caster.Equal(this.Balance, that1.Balance) {
			return false
		}
	}
	if len(this.Tokens) != len(that1.Tokens) {
		return false
	}
	for i := range this.Tokens {
		if !this.Tokens[i].Equal(that1.Tokens[i]) {
			return false
		}
	}
	return true
}
func (this *AccountTokenData) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AccountTokenData)
	if !ok {
		that2, ok := that.(AccountTokenData)
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
	if this.Nonce != that1.Nonce {
		return false
	}
	if this.Identifier != that1.Identifier {
		return false
	}
	{
		__caster := &github_com_numbatx_gn_core_data.BigIntCaster{}
		if !__caster.Equal(this.Balance, that1.Balance) {
			return false
		}
	}
	if this.Properties != that1.Properties {
		return false
	}
	return true
}
func (this *AlteredAccount) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&alteredAccount.AlteredAccount{")
	s = append(s, "Address: "+fmt.Sprintf("%#v", this.Address)+",\n")
	s = append(s, "Nonce: "+fmt.Sprintf("%#v", this.Nonce)+",\n")
	s = append(s, "Balance: "+fmt.Sprintf("%#v", this.Balance)+",\n")
	if this.Tokens != nil {
		s = append(s, "Tokens: "+fmt.Sprintf("%#v", this.Tokens)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *AccountTokenData) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&alteredAccount.AccountTokenData{")
	s = append(s, "Nonce: "+fmt.Sprintf("%#v", this.Nonce)+",\n")
	s = append(s, "Identifier: "+fmt.Sprintf("%#v", this.Identifier)+",\n")
	s = append(s, "Balance: "+fmt.Sprintf("%#v", this.Balance)+",\n")
	s = append(s, "Properties: "+fmt.Sprintf("%#v", this.Properties)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringAlteredAccount(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *AlteredAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AlteredAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AlteredAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Tokens) > 0 {
		for iNdEx := len(m.Tokens) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Tokens[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAlteredAccount(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	{
		__caster := &github_com_numbatx_gn_core_data.BigIntCaster{}
		size := __caster.Size(m.Balance)
		i -= size
		if _, err := __caster.MarshalTo(m.Balance, dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAlteredAccount(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.Nonce != 0 {
		i = encodeVarintAlteredAccount(dAtA, i, uint64(m.Nonce))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintAlteredAccount(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AccountTokenData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AccountTokenData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AccountTokenData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Properties) > 0 {
		i -= len(m.Properties)
		copy(dAtA[i:], m.Properties)
		i = encodeVarintAlteredAccount(dAtA, i, uint64(len(m.Properties)))
		i--
		dAtA[i] = 0x22
	}
	{
		__caster := &github_com_numbatx_gn_core_data.BigIntCaster{}
		size := __caster.Size(m.Balance)
		i -= size
		if _, err := __caster.MarshalTo(m.Balance, dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAlteredAccount(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Identifier) > 0 {
		i -= len(m.Identifier)
		copy(dAtA[i:], m.Identifier)
		i = encodeVarintAlteredAccount(dAtA, i, uint64(len(m.Identifier)))
		i--
		dAtA[i] = 0x12
	}
	if m.Nonce != 0 {
		i = encodeVarintAlteredAccount(dAtA, i, uint64(m.Nonce))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintAlteredAccount(dAtA []byte, offset int, v uint64) int {
	offset -= sovAlteredAccount(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AlteredAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovAlteredAccount(uint64(l))
	}
	if m.Nonce != 0 {
		n += 1 + sovAlteredAccount(uint64(m.Nonce))
	}
	{
		__caster := &github_com_numbatx_gn_core_data.BigIntCaster{}
		l = __caster.Size(m.Balance)
		n += 1 + l + sovAlteredAccount(uint64(l))
	}
	if len(m.Tokens) > 0 {
		for _, e := range m.Tokens {
			l = e.Size()
			n += 1 + l + sovAlteredAccount(uint64(l))
		}
	}
	return n
}

func (m *AccountTokenData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Nonce != 0 {
		n += 1 + sovAlteredAccount(uint64(m.Nonce))
	}
	l = len(m.Identifier)
	if l > 0 {
		n += 1 + l + sovAlteredAccount(uint64(l))
	}
	{
		__caster := &github_com_numbatx_gn_core_data.BigIntCaster{}
		l = __caster.Size(m.Balance)
		n += 1 + l + sovAlteredAccount(uint64(l))
	}
	l = len(m.Properties)
	if l > 0 {
		n += 1 + l + sovAlteredAccount(uint64(l))
	}
	return n
}

func sovAlteredAccount(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAlteredAccount(x uint64) (n int) {
	return sovAlteredAccount(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *AlteredAccount) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForTokens := "[]*AccountTokenData{"
	for _, f := range this.Tokens {
		repeatedStringForTokens += strings.Replace(f.String(), "AccountTokenData", "AccountTokenData", 1) + ","
	}
	repeatedStringForTokens += "}"
	s := strings.Join([]string{`&AlteredAccount{`,
		`Address:` + fmt.Sprintf("%v", this.Address) + `,`,
		`Nonce:` + fmt.Sprintf("%v", this.Nonce) + `,`,
		`Balance:` + fmt.Sprintf("%v", this.Balance) + `,`,
		`Tokens:` + repeatedStringForTokens + `,`,
		`}`,
	}, "")
	return s
}
func (this *AccountTokenData) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&AccountTokenData{`,
		`Nonce:` + fmt.Sprintf("%v", this.Nonce) + `,`,
		`Identifier:` + fmt.Sprintf("%v", this.Identifier) + `,`,
		`Balance:` + fmt.Sprintf("%v", this.Balance) + `,`,
		`Properties:` + fmt.Sprintf("%v", this.Properties) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringAlteredAccount(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *AlteredAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAlteredAccount
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
			return fmt.Errorf("proto: AlteredAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AlteredAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAlteredAccount
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
				return ErrInvalidLengthAlteredAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAlteredAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAlteredAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balance", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAlteredAccount
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
				return ErrInvalidLengthAlteredAccount
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAlteredAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			{
				__caster := &github_com_numbatx_gn_core_data.BigIntCaster{}
				if tmp, err := __caster.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
					return err
				} else {
					m.Balance = tmp
				}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tokens", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAlteredAccount
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
				return ErrInvalidLengthAlteredAccount
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAlteredAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tokens = append(m.Tokens, &AccountTokenData{})
			if err := m.Tokens[len(m.Tokens)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAlteredAccount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAlteredAccount
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
func (m *AccountTokenData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAlteredAccount
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
			return fmt.Errorf("proto: AccountTokenData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AccountTokenData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAlteredAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Identifier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAlteredAccount
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
				return ErrInvalidLengthAlteredAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAlteredAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Identifier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balance", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAlteredAccount
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
				return ErrInvalidLengthAlteredAccount
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAlteredAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			{
				__caster := &github_com_numbatx_gn_core_data.BigIntCaster{}
				if tmp, err := __caster.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
					return err
				} else {
					m.Balance = tmp
				}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Properties", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAlteredAccount
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
				return ErrInvalidLengthAlteredAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAlteredAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Properties = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAlteredAccount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAlteredAccount
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
func skipAlteredAccount(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAlteredAccount
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
					return 0, ErrIntOverflowAlteredAccount
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
					return 0, ErrIntOverflowAlteredAccount
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
				return 0, ErrInvalidLengthAlteredAccount
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAlteredAccount
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAlteredAccount
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAlteredAccount        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAlteredAccount          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAlteredAccount = fmt.Errorf("proto: unexpected end of group")
)
