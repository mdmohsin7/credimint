// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: credimint/credimint/user.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

type User struct {
	Index               string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	CreditScore         uint64 `protobuf:"varint,2,opt,name=creditScore,proto3" json:"creditScore,omitempty"`
	TimelyPayments      uint64 `protobuf:"varint,3,opt,name=timelyPayments,proto3" json:"timelyPayments,omitempty"`
	DefaultRate         string `protobuf:"bytes,4,opt,name=defaultRate,proto3" json:"defaultRate,omitempty"`
	NumberOfLoans       uint64 `protobuf:"varint,5,opt,name=numberOfLoans,proto3" json:"numberOfLoans,omitempty"`
	LoanDuration        uint64 `protobuf:"varint,6,opt,name=loanDuration,proto3" json:"loanDuration,omitempty"`
	NumberOfLoansFunded uint64 `protobuf:"varint,7,opt,name=numberOfLoansFunded,proto3" json:"numberOfLoansFunded,omitempty"`
	LoanFundedDuration  uint64 `protobuf:"varint,8,opt,name=loanFundedDuration,proto3" json:"loanFundedDuration,omitempty"`
	CollateralPercent   string `protobuf:"bytes,9,opt,name=collateralPercent,proto3" json:"collateralPercent,omitempty"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_de01d8df08bc6659, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_User.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return m.Size()
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *User) GetCreditScore() uint64 {
	if m != nil {
		return m.CreditScore
	}
	return 0
}

func (m *User) GetTimelyPayments() uint64 {
	if m != nil {
		return m.TimelyPayments
	}
	return 0
}

func (m *User) GetDefaultRate() string {
	if m != nil {
		return m.DefaultRate
	}
	return ""
}

func (m *User) GetNumberOfLoans() uint64 {
	if m != nil {
		return m.NumberOfLoans
	}
	return 0
}

func (m *User) GetLoanDuration() uint64 {
	if m != nil {
		return m.LoanDuration
	}
	return 0
}

func (m *User) GetNumberOfLoansFunded() uint64 {
	if m != nil {
		return m.NumberOfLoansFunded
	}
	return 0
}

func (m *User) GetLoanFundedDuration() uint64 {
	if m != nil {
		return m.LoanFundedDuration
	}
	return 0
}

func (m *User) GetCollateralPercent() string {
	if m != nil {
		return m.CollateralPercent
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "credimint.credimint.User")
}

func init() { proto.RegisterFile("credimint/credimint/user.proto", fileDescriptor_de01d8df08bc6659) }

var fileDescriptor_de01d8df08bc6659 = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcd, 0x4a, 0xc3, 0x40,
	0x14, 0x85, 0x3b, 0x35, 0xad, 0x76, 0xfc, 0x01, 0xa7, 0x2e, 0x06, 0x84, 0x21, 0x14, 0x91, 0x2e,
	0x24, 0x0a, 0xe2, 0x0b, 0x88, 0xb8, 0x12, 0x2c, 0x11, 0x37, 0xee, 0xa6, 0xc9, 0x2d, 0x04, 0x26,
	0x33, 0x65, 0x72, 0x07, 0x9a, 0xb7, 0xf0, 0xb1, 0x5c, 0x76, 0xe9, 0x52, 0x92, 0x17, 0x91, 0x4e,
	0x21, 0x4d, 0xb4, 0xbb, 0x7b, 0xbf, 0x73, 0xee, 0x19, 0x86, 0x43, 0x45, 0x62, 0x21, 0xcd, 0xf2,
	0x4c, 0xe3, 0xed, 0x6e, 0x72, 0x05, 0xd8, 0x68, 0x69, 0x0d, 0x1a, 0x36, 0x6e, 0x68, 0xd4, 0x4c,
	0x93, 0xaa, 0x4f, 0x83, 0xf7, 0x02, 0x2c, 0xbb, 0xa0, 0x83, 0x4c, 0xa7, 0xb0, 0xe2, 0x24, 0x24,
	0xd3, 0x51, 0xbc, 0x5d, 0x58, 0x48, 0x8f, 0xbd, 0x17, 0xdf, 0x12, 0x63, 0x81, 0xf7, 0x43, 0x32,
	0x0d, 0xe2, 0x36, 0x62, 0xd7, 0xf4, 0x0c, 0xb3, 0x1c, 0x54, 0x39, 0x93, 0x65, 0x0e, 0x1a, 0x0b,
	0x7e, 0xe0, 0x4d, 0x7f, 0xe8, 0x26, 0x29, 0x85, 0x85, 0x74, 0x0a, 0x63, 0x89, 0xc0, 0x03, 0xff,
	0x4a, 0x1b, 0xb1, 0x2b, 0x7a, 0xaa, 0x5d, 0x3e, 0x07, 0xfb, 0xba, 0x78, 0x31, 0x52, 0x17, 0x7c,
	0xe0, 0x83, 0xba, 0x90, 0x4d, 0xe8, 0x89, 0x32, 0x52, 0x3f, 0x39, 0x2b, 0x31, 0x33, 0x9a, 0x0f,
	0xbd, 0xa9, 0xc3, 0xd8, 0x1d, 0x1d, 0x77, 0x8e, 0x9e, 0x9d, 0x4e, 0x21, 0xe5, 0x87, 0xde, 0xba,
	0x4f, 0x62, 0x11, 0x65, 0x9b, 0x84, 0xed, 0xd6, 0x64, 0x1f, 0xf9, 0x83, 0x3d, 0x0a, 0xbb, 0xa1,
	0xe7, 0x89, 0x51, 0x4a, 0x22, 0x58, 0xa9, 0x66, 0x60, 0x13, 0xd0, 0xc8, 0x47, 0xfe, 0x4f, 0xff,
	0x85, 0xc7, 0x87, 0xaf, 0x4a, 0x90, 0x75, 0x25, 0xc8, 0x4f, 0x25, 0xc8, 0x67, 0x2d, 0x7a, 0xeb,
	0x5a, 0xf4, 0xbe, 0x6b, 0xd1, 0xfb, 0xb8, 0xdc, 0x35, 0xb5, 0x6a, 0xb5, 0x86, 0xe5, 0x12, 0x8a,
	0xf9, 0xd0, 0xf7, 0x76, 0xff, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x2b, 0xcb, 0x82, 0x8e, 0xd9, 0x01,
	0x00, 0x00,
}

func (m *User) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *User) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *User) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CollateralPercent) > 0 {
		i -= len(m.CollateralPercent)
		copy(dAtA[i:], m.CollateralPercent)
		i = encodeVarintUser(dAtA, i, uint64(len(m.CollateralPercent)))
		i--
		dAtA[i] = 0x4a
	}
	if m.LoanFundedDuration != 0 {
		i = encodeVarintUser(dAtA, i, uint64(m.LoanFundedDuration))
		i--
		dAtA[i] = 0x40
	}
	if m.NumberOfLoansFunded != 0 {
		i = encodeVarintUser(dAtA, i, uint64(m.NumberOfLoansFunded))
		i--
		dAtA[i] = 0x38
	}
	if m.LoanDuration != 0 {
		i = encodeVarintUser(dAtA, i, uint64(m.LoanDuration))
		i--
		dAtA[i] = 0x30
	}
	if m.NumberOfLoans != 0 {
		i = encodeVarintUser(dAtA, i, uint64(m.NumberOfLoans))
		i--
		dAtA[i] = 0x28
	}
	if len(m.DefaultRate) > 0 {
		i -= len(m.DefaultRate)
		copy(dAtA[i:], m.DefaultRate)
		i = encodeVarintUser(dAtA, i, uint64(len(m.DefaultRate)))
		i--
		dAtA[i] = 0x22
	}
	if m.TimelyPayments != 0 {
		i = encodeVarintUser(dAtA, i, uint64(m.TimelyPayments))
		i--
		dAtA[i] = 0x18
	}
	if m.CreditScore != 0 {
		i = encodeVarintUser(dAtA, i, uint64(m.CreditScore))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintUser(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintUser(dAtA []byte, offset int, v uint64) int {
	offset -= sovUser(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *User) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	if m.CreditScore != 0 {
		n += 1 + sovUser(uint64(m.CreditScore))
	}
	if m.TimelyPayments != 0 {
		n += 1 + sovUser(uint64(m.TimelyPayments))
	}
	l = len(m.DefaultRate)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	if m.NumberOfLoans != 0 {
		n += 1 + sovUser(uint64(m.NumberOfLoans))
	}
	if m.LoanDuration != 0 {
		n += 1 + sovUser(uint64(m.LoanDuration))
	}
	if m.NumberOfLoansFunded != 0 {
		n += 1 + sovUser(uint64(m.NumberOfLoansFunded))
	}
	if m.LoanFundedDuration != 0 {
		n += 1 + sovUser(uint64(m.LoanFundedDuration))
	}
	l = len(m.CollateralPercent)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	return n
}

func sovUser(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozUser(x uint64) (n int) {
	return sovUser(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *User) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUser
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
			return fmt.Errorf("proto: User: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: User: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreditScore", wireType)
			}
			m.CreditScore = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreditScore |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimelyPayments", wireType)
			}
			m.TimelyPayments = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimelyPayments |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefaultRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DefaultRate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumberOfLoans", wireType)
			}
			m.NumberOfLoans = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumberOfLoans |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LoanDuration", wireType)
			}
			m.LoanDuration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LoanDuration |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumberOfLoansFunded", wireType)
			}
			m.NumberOfLoansFunded = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumberOfLoansFunded |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LoanFundedDuration", wireType)
			}
			m.LoanFundedDuration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LoanFundedDuration |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollateralPercent", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CollateralPercent = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUser(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUser
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
func skipUser(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUser
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
					return 0, ErrIntOverflowUser
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
					return 0, ErrIntOverflowUser
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
				return 0, ErrInvalidLengthUser
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupUser
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthUser
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthUser        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUser          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupUser = fmt.Errorf("proto: unexpected end of group")
)
