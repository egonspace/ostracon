// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ostracon/types/validator.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	crypto "github.com/line/ostracon/proto/ostracon/crypto"
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

type ValidatorSet struct {
	Validators []*Validator `protobuf:"bytes,1,rep,name=validators,proto3" json:"validators,omitempty"`
	// Validator          proposer           = 2;   // This field has been removed in Ostracon.
	TotalStakingPower int64 `protobuf:"varint,3,opt,name=total_staking_power,json=totalStakingPower,proto3" json:"total_staking_power,omitempty"`
}

func (m *ValidatorSet) Reset()         { *m = ValidatorSet{} }
func (m *ValidatorSet) String() string { return proto.CompactTextString(m) }
func (*ValidatorSet) ProtoMessage()    {}
func (*ValidatorSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeeaf81579407bf3, []int{0}
}
func (m *ValidatorSet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorSet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorSet.Merge(m, src)
}
func (m *ValidatorSet) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorSet) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorSet.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorSet proto.InternalMessageInfo

func (m *ValidatorSet) GetValidators() []*Validator {
	if m != nil {
		return m.Validators
	}
	return nil
}

func (m *ValidatorSet) GetTotalStakingPower() int64 {
	if m != nil {
		return m.TotalStakingPower
	}
	return 0
}

type Validator struct {
	Address          []byte           `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	PubKey           crypto.PublicKey `protobuf:"bytes,2,opt,name=pub_key,json=pubKey,proto3" json:"pub_key"`
	StakingPower     int64            `protobuf:"varint,3,opt,name=staking_power,json=stakingPower,proto3" json:"staking_power,omitempty"`
	ProposerPriority int64            `protobuf:"varint,4,opt,name=proposer_priority,json=proposerPriority,proto3" json:"proposer_priority,omitempty"`
	// *** Ostracon Extended Fields ***
	VotingPower int64 `protobuf:"varint,1000,opt,name=voting_power,json=votingPower,proto3" json:"voting_power,omitempty"`
}

func (m *Validator) Reset()         { *m = Validator{} }
func (m *Validator) String() string { return proto.CompactTextString(m) }
func (*Validator) ProtoMessage()    {}
func (*Validator) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeeaf81579407bf3, []int{1}
}
func (m *Validator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Validator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Validator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Validator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Validator.Merge(m, src)
}
func (m *Validator) XXX_Size() int {
	return m.Size()
}
func (m *Validator) XXX_DiscardUnknown() {
	xxx_messageInfo_Validator.DiscardUnknown(m)
}

var xxx_messageInfo_Validator proto.InternalMessageInfo

func (m *Validator) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Validator) GetPubKey() crypto.PublicKey {
	if m != nil {
		return m.PubKey
	}
	return crypto.PublicKey{}
}

func (m *Validator) GetStakingPower() int64 {
	if m != nil {
		return m.StakingPower
	}
	return 0
}

func (m *Validator) GetProposerPriority() int64 {
	if m != nil {
		return m.ProposerPriority
	}
	return 0
}

func (m *Validator) GetVotingPower() int64 {
	if m != nil {
		return m.VotingPower
	}
	return 0
}

type SimpleValidator struct {
	PubKey       *crypto.PublicKey `protobuf:"bytes,1,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty"`
	StakingPower int64             `protobuf:"varint,2,opt,name=staking_power,json=stakingPower,proto3" json:"staking_power,omitempty"`
}

func (m *SimpleValidator) Reset()         { *m = SimpleValidator{} }
func (m *SimpleValidator) String() string { return proto.CompactTextString(m) }
func (*SimpleValidator) ProtoMessage()    {}
func (*SimpleValidator) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeeaf81579407bf3, []int{2}
}
func (m *SimpleValidator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SimpleValidator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SimpleValidator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SimpleValidator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleValidator.Merge(m, src)
}
func (m *SimpleValidator) XXX_Size() int {
	return m.Size()
}
func (m *SimpleValidator) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleValidator.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleValidator proto.InternalMessageInfo

func (m *SimpleValidator) GetPubKey() *crypto.PublicKey {
	if m != nil {
		return m.PubKey
	}
	return nil
}

func (m *SimpleValidator) GetStakingPower() int64 {
	if m != nil {
		return m.StakingPower
	}
	return 0
}

func init() {
	proto.RegisterType((*ValidatorSet)(nil), "ostracon.types.ValidatorSet")
	proto.RegisterType((*Validator)(nil), "ostracon.types.Validator")
	proto.RegisterType((*SimpleValidator)(nil), "ostracon.types.SimpleValidator")
}

func init() { proto.RegisterFile("ostracon/types/validator.proto", fileDescriptor_eeeaf81579407bf3) }

var fileDescriptor_eeeaf81579407bf3 = []byte{
	// 369 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xbf, 0x6e, 0xea, 0x30,
	0x18, 0xc5, 0x63, 0x40, 0xa0, 0x6b, 0x72, 0xff, 0xe0, 0xdb, 0x21, 0xcd, 0x90, 0x22, 0xba, 0x44,
	0xaa, 0x9a, 0x48, 0x30, 0xb1, 0xb2, 0x74, 0x60, 0x41, 0x41, 0xea, 0xd0, 0x25, 0x4a, 0x82, 0x95,
	0x5a, 0x09, 0xd8, 0xb2, 0x1d, 0x2a, 0xbf, 0x45, 0x1f, 0x8b, 0x91, 0xa9, 0xea, 0x54, 0x55, 0xb0,
	0xf4, 0x31, 0x2a, 0x92, 0x26, 0x01, 0xb5, 0x55, 0xb7, 0xe4, 0xfb, 0x1d, 0x9f, 0xef, 0x1c, 0xe9,
	0x83, 0x16, 0x15, 0x92, 0x07, 0x11, 0x5d, 0xb9, 0x52, 0x31, 0x2c, 0xdc, 0x75, 0x90, 0x92, 0x45,
	0x20, 0x29, 0x77, 0x18, 0xa7, 0x92, 0xa2, 0x3f, 0x25, 0x77, 0x72, 0x6e, 0x9e, 0xc5, 0x34, 0xa6,
	0x39, 0x72, 0x0f, 0x5f, 0x85, 0xca, 0x34, 0x2b, 0x97, 0x88, 0x2b, 0x26, 0xa9, 0x9b, 0x60, 0x25,
	0x0a, 0x36, 0x50, 0x50, 0xbf, 0x2d, 0x4d, 0xe7, 0x58, 0xa2, 0x31, 0x84, 0xd5, 0x12, 0x61, 0x80,
	0x7e, 0xd3, 0xee, 0x0e, 0xcf, 0x9d, 0xd3, 0x35, 0x4e, 0xf5, 0xc2, 0x3b, 0x12, 0x23, 0x07, 0xfe,
	0x97, 0x54, 0x06, 0xa9, 0x2f, 0x64, 0x90, 0x90, 0x55, 0xec, 0x33, 0xfa, 0x80, 0xb9, 0xd1, 0xec,
	0x03, 0xbb, 0xe9, 0xf5, 0x72, 0x34, 0x2f, 0xc8, 0xec, 0x00, 0x06, 0x4f, 0x00, 0xfe, 0xaa, 0x9c,
	0x90, 0x01, 0x3b, 0xc1, 0x62, 0xc1, 0xb1, 0x38, 0x6c, 0x05, 0xb6, 0xee, 0x95, 0xbf, 0x68, 0x0c,
	0x3b, 0x2c, 0x0b, 0xfd, 0x04, 0x2b, 0xa3, 0xd1, 0x07, 0x76, 0x77, 0x68, 0xd6, 0x79, 0x8a, 0x42,
	0xce, 0x2c, 0x0b, 0x53, 0x12, 0x4d, 0xb1, 0x9a, 0xb4, 0x36, 0x2f, 0x17, 0x9a, 0xd7, 0x66, 0x59,
	0x38, 0xc5, 0x0a, 0x5d, 0xc2, 0xdf, 0x5f, 0x85, 0xd1, 0xc5, 0x51, 0x0e, 0x74, 0x05, 0x7b, 0x8c,
	0x53, 0x46, 0x05, 0xe6, 0x3e, 0xe3, 0x84, 0x72, 0x22, 0x95, 0xd1, 0xca, 0x85, 0xff, 0x4a, 0x30,
	0xfb, 0x98, 0xa3, 0x01, 0xd4, 0xd7, 0x54, 0xd6, 0x86, 0x6f, 0x9d, 0x5c, 0xd8, 0x2d, 0x86, 0x45,
	0xb1, 0x04, 0xfe, 0x9d, 0x93, 0x25, 0x4b, 0x71, 0xdd, 0x6e, 0x54, 0x77, 0x00, 0x3f, 0x75, 0xf8,
	0x3e, 0x7d, 0xe3, 0x73, 0xfa, 0xc9, 0xcd, 0x66, 0x67, 0x81, 0xed, 0xce, 0x02, 0xaf, 0x3b, 0x0b,
	0x3c, 0xee, 0x2d, 0x6d, 0xbb, 0xb7, 0xb4, 0xe7, 0xbd, 0xa5, 0xdd, 0x5d, 0xc7, 0x44, 0xde, 0x67,
	0xa1, 0x13, 0xd1, 0xa5, 0x9b, 0x92, 0x15, 0x76, 0xab, 0x33, 0x28, 0x2e, 0xe4, 0xf4, 0xb6, 0xc2,
	0x76, 0x3e, 0x1d, 0xbd, 0x07, 0x00, 0x00, 0xff, 0xff, 0x14, 0xa3, 0xe5, 0xb9, 0x74, 0x02, 0x00,
	0x00,
}

func (m *ValidatorSet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorSet) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorSet) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TotalStakingPower != 0 {
		i = encodeVarintValidator(dAtA, i, uint64(m.TotalStakingPower))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Validators) > 0 {
		for iNdEx := len(m.Validators) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Validators[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintValidator(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Validator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Validator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Validator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.VotingPower != 0 {
		i = encodeVarintValidator(dAtA, i, uint64(m.VotingPower))
		i--
		dAtA[i] = 0x3e
		i--
		dAtA[i] = 0xc0
	}
	if m.ProposerPriority != 0 {
		i = encodeVarintValidator(dAtA, i, uint64(m.ProposerPriority))
		i--
		dAtA[i] = 0x20
	}
	if m.StakingPower != 0 {
		i = encodeVarintValidator(dAtA, i, uint64(m.StakingPower))
		i--
		dAtA[i] = 0x18
	}
	{
		size, err := m.PubKey.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintValidator(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintValidator(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SimpleValidator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SimpleValidator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SimpleValidator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.StakingPower != 0 {
		i = encodeVarintValidator(dAtA, i, uint64(m.StakingPower))
		i--
		dAtA[i] = 0x10
	}
	if m.PubKey != nil {
		{
			size, err := m.PubKey.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintValidator(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintValidator(dAtA []byte, offset int, v uint64) int {
	offset -= sovValidator(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ValidatorSet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Validators) > 0 {
		for _, e := range m.Validators {
			l = e.Size()
			n += 1 + l + sovValidator(uint64(l))
		}
	}
	if m.TotalStakingPower != 0 {
		n += 1 + sovValidator(uint64(m.TotalStakingPower))
	}
	return n
}

func (m *Validator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovValidator(uint64(l))
	}
	l = m.PubKey.Size()
	n += 1 + l + sovValidator(uint64(l))
	if m.StakingPower != 0 {
		n += 1 + sovValidator(uint64(m.StakingPower))
	}
	if m.ProposerPriority != 0 {
		n += 1 + sovValidator(uint64(m.ProposerPriority))
	}
	if m.VotingPower != 0 {
		n += 2 + sovValidator(uint64(m.VotingPower))
	}
	return n
}

func (m *SimpleValidator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PubKey != nil {
		l = m.PubKey.Size()
		n += 1 + l + sovValidator(uint64(l))
	}
	if m.StakingPower != 0 {
		n += 1 + sovValidator(uint64(m.StakingPower))
	}
	return n
}

func sovValidator(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozValidator(x uint64) (n int) {
	return sovValidator(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ValidatorSet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowValidator
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
			return fmt.Errorf("proto: ValidatorSet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorSet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validators", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validators = append(m.Validators, &Validator{})
			if err := m.Validators[len(m.Validators)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalStakingPower", wireType)
			}
			m.TotalStakingPower = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalStakingPower |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipValidator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthValidator
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
func (m *Validator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowValidator
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
			return fmt.Errorf("proto: Validator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Validator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PubKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakingPower", wireType)
			}
			m.StakingPower = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StakingPower |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposerPriority", wireType)
			}
			m.ProposerPriority = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProposerPriority |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 1000:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotingPower", wireType)
			}
			m.VotingPower = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VotingPower |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipValidator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthValidator
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
func (m *SimpleValidator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowValidator
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
			return fmt.Errorf("proto: SimpleValidator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SimpleValidator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PubKey == nil {
				m.PubKey = &crypto.PublicKey{}
			}
			if err := m.PubKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakingPower", wireType)
			}
			m.StakingPower = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StakingPower |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipValidator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthValidator
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
func skipValidator(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowValidator
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
					return 0, ErrIntOverflowValidator
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
					return 0, ErrIntOverflowValidator
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
				return 0, ErrInvalidLengthValidator
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupValidator
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthValidator
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthValidator        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowValidator          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupValidator = fmt.Errorf("proto: unexpected end of group")
)
