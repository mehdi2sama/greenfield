// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: greenfield/payment/delayed_withdrawal_record.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

type DelayedWithdrawalRecord struct {
	// the withdrawal address
	Addr string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	// the withdrawal amount
	Amount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount"`
	// the withdrawal from payment account address
	From string `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	// unlock timestamp is the unix timestamp to unlock the withdrawal
	UnlockTimestamp int64 `protobuf:"varint,4,opt,name=unlock_timestamp,json=unlockTimestamp,proto3" json:"unlock_timestamp,omitempty"`
}

func (m *DelayedWithdrawalRecord) Reset()         { *m = DelayedWithdrawalRecord{} }
func (m *DelayedWithdrawalRecord) String() string { return proto.CompactTextString(m) }
func (*DelayedWithdrawalRecord) ProtoMessage()    {}
func (*DelayedWithdrawalRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_237dd2860d399f1a, []int{0}
}
func (m *DelayedWithdrawalRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DelayedWithdrawalRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DelayedWithdrawalRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DelayedWithdrawalRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelayedWithdrawalRecord.Merge(m, src)
}
func (m *DelayedWithdrawalRecord) XXX_Size() int {
	return m.Size()
}
func (m *DelayedWithdrawalRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_DelayedWithdrawalRecord.DiscardUnknown(m)
}

var xxx_messageInfo_DelayedWithdrawalRecord proto.InternalMessageInfo

func (m *DelayedWithdrawalRecord) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *DelayedWithdrawalRecord) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *DelayedWithdrawalRecord) GetUnlockTimestamp() int64 {
	if m != nil {
		return m.UnlockTimestamp
	}
	return 0
}

func init() {
	proto.RegisterType((*DelayedWithdrawalRecord)(nil), "greenfield.payment.DelayedWithdrawalRecord")
}

func init() {
	proto.RegisterFile("greenfield/payment/delayed_withdrawal_record.proto", fileDescriptor_237dd2860d399f1a)
}

var fileDescriptor_237dd2860d399f1a = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xc1, 0x4a, 0xc3, 0x30,
	0x1c, 0xc6, 0x1b, 0x37, 0x06, 0xf6, 0xa2, 0x94, 0x81, 0x75, 0x87, 0x6e, 0x78, 0x90, 0x09, 0xb6,
	0x05, 0xbd, 0x7a, 0x71, 0x78, 0xd9, 0xb5, 0x0e, 0x04, 0x2f, 0x25, 0x6d, 0xb2, 0x2e, 0xac, 0x49,
	0x4a, 0x92, 0x31, 0xf7, 0x16, 0x3e, 0xcc, 0x1e, 0x62, 0xc7, 0xb1, 0x93, 0x78, 0x18, 0xb2, 0xbd,
	0x83, 0x67, 0x69, 0x12, 0xe7, 0x6e, 0x9e, 0x92, 0x7c, 0x7c, 0xff, 0xef, 0xfb, 0x91, 0xbf, 0x7b,
	0x57, 0x08, 0x8c, 0xd9, 0x98, 0xe0, 0x12, 0xc5, 0x15, 0x5c, 0x50, 0xcc, 0x54, 0x8c, 0x70, 0x09,
	0x17, 0x18, 0xa5, 0x73, 0xa2, 0x26, 0x48, 0xc0, 0x39, 0x2c, 0x53, 0x81, 0x73, 0x2e, 0x50, 0x54,
	0x09, 0xae, 0xb8, 0xe7, 0xfd, 0xcd, 0x44, 0x76, 0xa6, 0x73, 0x99, 0x73, 0x49, 0xb9, 0x4c, 0xb5,
	0x23, 0x36, 0x0f, 0x63, 0xef, 0xb4, 0x0b, 0x5e, 0x70, 0xa3, 0xd7, 0x37, 0xa3, 0x5e, 0x7d, 0x03,
	0xf7, 0xe2, 0xc9, 0x14, 0xbd, 0x1c, 0x7a, 0x12, 0x5d, 0xe3, 0xdd, 0xba, 0x4d, 0x88, 0x90, 0xf0,
	0x41, 0x0f, 0xf4, 0x4f, 0x07, 0xfe, 0x66, 0x19, 0xb6, 0x6d, 0xe2, 0x23, 0x42, 0x02, 0x4b, 0xf9,
	0xac, 0x04, 0x61, 0x45, 0xa2, 0x5d, 0xde, 0xc8, 0x6d, 0x41, 0xca, 0x67, 0x4c, 0xf9, 0x27, 0xda,
	0xff, 0xb0, 0xda, 0x76, 0x9d, 0xcf, 0x6d, 0xf7, 0xba, 0x20, 0x6a, 0x32, 0xcb, 0xa2, 0x9c, 0x53,
	0x0b, 0x64, 0x8f, 0x50, 0xa2, 0x69, 0xac, 0x16, 0x15, 0x96, 0xd1, 0x90, 0xa9, 0xcd, 0x32, 0x74,
	0x6d, 0xfa, 0x90, 0xa9, 0xc4, 0x66, 0xd5, 0x0c, 0x63, 0xc1, 0xa9, 0xdf, 0xf8, 0x8f, 0xa1, 0x76,
	0x79, 0x37, 0xee, 0xf9, 0x8c, 0x95, 0x3c, 0x9f, 0xa6, 0x8a, 0x50, 0x2c, 0x15, 0xa4, 0x95, 0xdf,
	0xec, 0x81, 0x7e, 0x23, 0x39, 0x33, 0xfa, 0xe8, 0x57, 0x1e, 0x0c, 0x57, 0xbb, 0x00, 0xac, 0x77,
	0x01, 0xf8, 0xda, 0x05, 0xe0, 0x7d, 0x1f, 0x38, 0xeb, 0x7d, 0xe0, 0x7c, 0xec, 0x03, 0xe7, 0x35,
	0x3e, 0x02, 0xce, 0x58, 0x16, 0xe6, 0x13, 0x48, 0x58, 0x7c, 0xb4, 0xa0, 0xb7, 0xc3, 0x8a, 0x34,
	0x7d, 0xd6, 0xd2, 0x5f, 0x79, 0xff, 0x13, 0x00, 0x00, 0xff, 0xff, 0x6f, 0xd9, 0xf0, 0x28, 0xc5,
	0x01, 0x00, 0x00,
}

func (m *DelayedWithdrawalRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DelayedWithdrawalRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DelayedWithdrawalRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.UnlockTimestamp != 0 {
		i = encodeVarintDelayedWithdrawalRecord(dAtA, i, uint64(m.UnlockTimestamp))
		i--
		dAtA[i] = 0x20
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintDelayedWithdrawalRecord(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintDelayedWithdrawalRecord(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Addr) > 0 {
		i -= len(m.Addr)
		copy(dAtA[i:], m.Addr)
		i = encodeVarintDelayedWithdrawalRecord(dAtA, i, uint64(len(m.Addr)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDelayedWithdrawalRecord(dAtA []byte, offset int, v uint64) int {
	offset -= sovDelayedWithdrawalRecord(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DelayedWithdrawalRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Addr)
	if l > 0 {
		n += 1 + l + sovDelayedWithdrawalRecord(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovDelayedWithdrawalRecord(uint64(l))
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovDelayedWithdrawalRecord(uint64(l))
	}
	if m.UnlockTimestamp != 0 {
		n += 1 + sovDelayedWithdrawalRecord(uint64(m.UnlockTimestamp))
	}
	return n
}

func sovDelayedWithdrawalRecord(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDelayedWithdrawalRecord(x uint64) (n int) {
	return sovDelayedWithdrawalRecord(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DelayedWithdrawalRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDelayedWithdrawalRecord
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
			return fmt.Errorf("proto: DelayedWithdrawalRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DelayedWithdrawalRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelayedWithdrawalRecord
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
				return ErrInvalidLengthDelayedWithdrawalRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDelayedWithdrawalRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelayedWithdrawalRecord
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
				return ErrInvalidLengthDelayedWithdrawalRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDelayedWithdrawalRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelayedWithdrawalRecord
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
				return ErrInvalidLengthDelayedWithdrawalRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDelayedWithdrawalRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnlockTimestamp", wireType)
			}
			m.UnlockTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelayedWithdrawalRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UnlockTimestamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDelayedWithdrawalRecord(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDelayedWithdrawalRecord
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
func skipDelayedWithdrawalRecord(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDelayedWithdrawalRecord
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
					return 0, ErrIntOverflowDelayedWithdrawalRecord
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
					return 0, ErrIntOverflowDelayedWithdrawalRecord
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
				return 0, ErrInvalidLengthDelayedWithdrawalRecord
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDelayedWithdrawalRecord
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDelayedWithdrawalRecord
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDelayedWithdrawalRecord        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDelayedWithdrawalRecord          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDelayedWithdrawalRecord = fmt.Errorf("proto: unexpected end of group")
)
