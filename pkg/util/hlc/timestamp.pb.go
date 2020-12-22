// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: util/hlc/timestamp.proto

package hlc

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// TimestampFlag is used to provide extra classification for Timestamps.
type TimestampFlag int32

const (
	TimestampFlag_UNKNOWN TimestampFlag = 0
	// A synthetic timestamp is defined as a timestamp that makes no claim
	// about the value of clocks in the system. While standard timestamps
	// are pulled from HLC clocks and indicate that some node in the system
	// has a clock with a reading equal to or above its value, a synthetic
	// timestamp makes no such indication.
	//
	// Synthetic timestamps are central to non-blocking transactions, which
	// write at "future timestamps". They are also used to disconnect some
	// committed versions from observed timestamps, where they indicate that
	// versions were moved from the timestamp at which they were originally
	// written. Only synthetic timestamps require observing the full
	// uncertainty interval, whereas readings off the leaseholders's clock
	// can tighten it for non-synthetic versions.
	TimestampFlag_SYNTHETIC TimestampFlag = 1
)

var TimestampFlag_name = map[int32]string{
	0: "UNKNOWN",
	1: "SYNTHETIC",
}
var TimestampFlag_value = map[string]int32{
	"UNKNOWN":   0,
	"SYNTHETIC": 1,
}

func (x TimestampFlag) String() string {
	return proto.EnumName(TimestampFlag_name, int32(x))
}
func (TimestampFlag) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_timestamp_2a42725d81a17263, []int{0}
}

// Timestamp represents a state of the hybrid logical clock.
type Timestamp struct {
	// Holds a wall time, typically a unix epoch time expressed in
	// nanoseconds.
	//
	// It is not safe to mutate this field directly. Instead, use one of the
	// methods on Timestamp, which ensure that the synthetic flag is updated
	// appropriately.
	WallTime int64 `protobuf:"varint,1,opt,name=wall_time,json=wallTime,proto3" json:"wall_time,omitempty"`
	// The logical component captures causality for events whose wall times
	// are equal. It is effectively bounded by (maximum clock skew)/(minimal
	// ns between events) and nearly impossible to overflow.
	//
	// It is not safe to mutate this field directly. Instead, use one of the
	// methods on Timestamp, which ensure that the synthetic flag is updated
	// appropriately.
	Logical int32 `protobuf:"varint,2,opt,name=logical,proto3" json:"logical,omitempty"`
	// A collection of bit flags that provide details about the timestamp
	// and its meaning. The data type is a uint32, but the number of flags
	// is limited to 8 so that the flags can be encoded into a single byte.
	//
	// Flags do not affect the sort order of Timestamps. However, they are
	// considered when performing structural equality checks (e.g. using the
	// == operator). Consider use of the EqOrdering method when testing for
	// equality.
	//
	// TODO(nvanbenschoten): use a bool to shave off a
	// byte when set. This will allow the flag to serve as the dynamically
	// typed version of ClockTimestamp. See TryToClockTimestamp.
	//
	// Should look like:
	//   bool synthetic = 3;
	//
	Flags uint32 `protobuf:"varint,3,opt,name=flags,proto3" json:"flags,omitempty"`
}

func (m *Timestamp) Reset()      { *m = Timestamp{} }
func (*Timestamp) ProtoMessage() {}
func (*Timestamp) Descriptor() ([]byte, []int) {
	return fileDescriptor_timestamp_2a42725d81a17263, []int{0}
}
func (m *Timestamp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Timestamp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *Timestamp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Timestamp.Merge(dst, src)
}
func (m *Timestamp) XXX_Size() int {
	return m.Size()
}
func (m *Timestamp) XXX_DiscardUnknown() {
	xxx_messageInfo_Timestamp.DiscardUnknown(m)
}

var xxx_messageInfo_Timestamp proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Timestamp)(nil), "cockroach.util.hlc.Timestamp")
	proto.RegisterEnum("cockroach.util.hlc.TimestampFlag", TimestampFlag_name, TimestampFlag_value)
}
func (this *Timestamp) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Timestamp)
	if !ok {
		that2, ok := that.(Timestamp)
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
	if this.WallTime != that1.WallTime {
		return false
	}
	if this.Logical != that1.Logical {
		return false
	}
	if this.Flags != that1.Flags {
		return false
	}
	return true
}
func (m *Timestamp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Timestamp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.WallTime != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintTimestamp(dAtA, i, uint64(m.WallTime))
	}
	if m.Logical != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintTimestamp(dAtA, i, uint64(m.Logical))
	}
	if m.Flags != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintTimestamp(dAtA, i, uint64(m.Flags))
	}
	return i, nil
}

func encodeVarintTimestamp(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedTimestamp(r randyTimestamp, easy bool) *Timestamp {
	this := &Timestamp{}
	this.WallTime = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.WallTime *= -1
	}
	this.Logical = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.Logical *= -1
	}
	this.Flags = uint32(r.Uint32())
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyTimestamp interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneTimestamp(r randyTimestamp) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringTimestamp(r randyTimestamp) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneTimestamp(r)
	}
	return string(tmps)
}
func randUnrecognizedTimestamp(r randyTimestamp, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldTimestamp(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldTimestamp(dAtA []byte, r randyTimestamp, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateTimestamp(dAtA, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		dAtA = encodeVarintPopulateTimestamp(dAtA, uint64(v2))
	case 1:
		dAtA = encodeVarintPopulateTimestamp(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateTimestamp(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateTimestamp(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateTimestamp(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateTimestamp(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Timestamp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.WallTime != 0 {
		n += 1 + sovTimestamp(uint64(m.WallTime))
	}
	if m.Logical != 0 {
		n += 1 + sovTimestamp(uint64(m.Logical))
	}
	if m.Flags != 0 {
		n += 1 + sovTimestamp(uint64(m.Flags))
	}
	return n
}

func sovTimestamp(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTimestamp(x uint64) (n int) {
	return sovTimestamp(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Timestamp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTimestamp
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Timestamp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Timestamp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WallTime", wireType)
			}
			m.WallTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimestamp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WallTime |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Logical", wireType)
			}
			m.Logical = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimestamp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Logical |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Flags", wireType)
			}
			m.Flags = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimestamp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Flags |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTimestamp(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTimestamp
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
func skipTimestamp(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTimestamp
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
					return 0, ErrIntOverflowTimestamp
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTimestamp
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthTimestamp
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTimestamp
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipTimestamp(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthTimestamp = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTimestamp   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("util/hlc/timestamp.proto", fileDescriptor_timestamp_2a42725d81a17263)
}

var fileDescriptor_timestamp_2a42725d81a17263 = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x28, 0x2d, 0xc9, 0xcc,
	0xd1, 0xcf, 0xc8, 0x49, 0xd6, 0x2f, 0xc9, 0xcc, 0x4d, 0x2d, 0x2e, 0x49, 0xcc, 0x2d, 0xd0, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4a, 0xce, 0x4f, 0xce, 0x2e, 0xca, 0x4f, 0x4c, 0xce, 0xd0,
	0x03, 0xa9, 0xd1, 0xcb, 0xc8, 0x49, 0x96, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x4b, 0xeb, 0x83,
	0x58, 0x10, 0x95, 0x4a, 0x69, 0x5c, 0x9c, 0x21, 0x30, 0xcd, 0x42, 0xd2, 0x5c, 0x9c, 0xe5, 0x89,
	0x39, 0x39, 0xf1, 0x20, 0xe3, 0x24, 0x18, 0x15, 0x18, 0x35, 0x98, 0x83, 0x38, 0x40, 0x02, 0x20,
	0x15, 0x42, 0x12, 0x5c, 0xec, 0x39, 0xf9, 0xe9, 0x99, 0xc9, 0x89, 0x39, 0x12, 0x4c, 0x0a, 0x8c,
	0x1a, 0xac, 0x41, 0x30, 0xae, 0x90, 0x08, 0x17, 0x6b, 0x5a, 0x4e, 0x62, 0x7a, 0xb1, 0x04, 0xb3,
	0x02, 0xa3, 0x06, 0x6f, 0x10, 0x84, 0x63, 0xc5, 0x33, 0x63, 0x81, 0x3c, 0xc3, 0x8e, 0x05, 0xf2,
	0x8c, 0x2f, 0x16, 0xc8, 0x33, 0x6a, 0x69, 0x73, 0xf1, 0xc2, 0xed, 0x71, 0xcb, 0x49, 0x4c, 0x17,
	0xe2, 0xe6, 0x62, 0x0f, 0xf5, 0xf3, 0xf6, 0xf3, 0x0f, 0xf7, 0x13, 0x60, 0x10, 0xe2, 0xe5, 0xe2,
	0x0c, 0x8e, 0xf4, 0x0b, 0xf1, 0x70, 0x0d, 0xf1, 0x74, 0x16, 0x60, 0x74, 0x52, 0x3d, 0xf1, 0x50,
	0x8e, 0xe1, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x6f, 0x3c, 0x92, 0x63, 0x7c, 0xf0,
	0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88,
	0x62, 0xce, 0xc8, 0x49, 0x4e, 0x62, 0x03, 0x7b, 0xc1, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xf4,
	0x8d, 0x21, 0xb8, 0x08, 0x01, 0x00, 0x00,
}
