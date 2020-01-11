// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: test/abc.proto

package test

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
	github_com_regen_network_cosmos_proto_test_iface "github.com/regen-network/cosmos-proto/test/iface"
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

type ABC struct {
	// Types that are valid to be assigned to Sum:
	//	*ABC_A
	//	*ABC_B
	//	*ABC_C
	Sum isABC_Sum `protobuf_oneof:"sum"`
}

func (m *ABC) Reset()         { *m = ABC{} }
func (m *ABC) String() string { return proto.CompactTextString(m) }
func (*ABC) ProtoMessage()    {}
func (*ABC) Descriptor() ([]byte, []int) {
	return fileDescriptor_edbe9e0eab89a1a7, []int{0}
}
func (m *ABC) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ABC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ABC.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ABC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ABC.Merge(m, src)
}
func (m *ABC) XXX_Size() int {
	return m.Size()
}
func (m *ABC) XXX_DiscardUnknown() {
	xxx_messageInfo_ABC.DiscardUnknown(m)
}

var xxx_messageInfo_ABC proto.InternalMessageInfo

type isABC_Sum interface {
	isABC_Sum()
	MarshalTo([]byte) (int, error)
	Size() int
}

type ABC_A struct {
	A *A `protobuf:"bytes,1,opt,name=a,proto3,oneof" json:"a,omitempty"`
}
type ABC_B struct {
	B *B `protobuf:"bytes,2,opt,name=b,proto3,oneof" json:"b,omitempty"`
}
type ABC_C struct {
	C *C `protobuf:"bytes,3,opt,name=c,proto3,oneof" json:"c,omitempty"`
}

func (*ABC_A) isABC_Sum() {}
func (*ABC_B) isABC_Sum() {}
func (*ABC_C) isABC_Sum() {}

func (m *ABC) GetSum() isABC_Sum {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *ABC) GetA() *A {
	if x, ok := m.GetSum().(*ABC_A); ok {
		return x.A
	}
	return nil
}

func (m *ABC) GetB() *B {
	if x, ok := m.GetSum().(*ABC_B); ok {
		return x.B
	}
	return nil
}

func (m *ABC) GetC() *C {
	if x, ok := m.GetSum().(*ABC_C); ok {
		return x.C
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ABC) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ABC_A)(nil),
		(*ABC_B)(nil),
		(*ABC_C)(nil),
	}
}

type A struct {
	X int32 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
}

func (m *A) Reset()         { *m = A{} }
func (m *A) String() string { return proto.CompactTextString(m) }
func (*A) ProtoMessage()    {}
func (*A) Descriptor() ([]byte, []int) {
	return fileDescriptor_edbe9e0eab89a1a7, []int{1}
}
func (m *A) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *A) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_A.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *A) XXX_Merge(src proto.Message) {
	xxx_messageInfo_A.Merge(m, src)
}
func (m *A) XXX_Size() int {
	return m.Size()
}
func (m *A) XXX_DiscardUnknown() {
	xxx_messageInfo_A.DiscardUnknown(m)
}

var xxx_messageInfo_A proto.InternalMessageInfo

func (m *A) GetX() int32 {
	if m != nil {
		return m.X
	}
	return 0
}

type B struct {
	Y uint32 `protobuf:"varint,1,opt,name=y,proto3" json:"y,omitempty"`
}

func (m *B) Reset()         { *m = B{} }
func (m *B) String() string { return proto.CompactTextString(m) }
func (*B) ProtoMessage()    {}
func (*B) Descriptor() ([]byte, []int) {
	return fileDescriptor_edbe9e0eab89a1a7, []int{2}
}
func (m *B) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *B) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_B.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *B) XXX_Merge(src proto.Message) {
	xxx_messageInfo_B.Merge(m, src)
}
func (m *B) XXX_Size() int {
	return m.Size()
}
func (m *B) XXX_DiscardUnknown() {
	xxx_messageInfo_B.DiscardUnknown(m)
}

var xxx_messageInfo_B proto.InternalMessageInfo

func (m *B) GetY() uint32 {
	if m != nil {
		return m.Y
	}
	return 0
}

type C struct {
	Z bool `protobuf:"varint,1,opt,name=z,proto3" json:"z,omitempty"`
}

func (m *C) Reset()         { *m = C{} }
func (m *C) String() string { return proto.CompactTextString(m) }
func (*C) ProtoMessage()    {}
func (*C) Descriptor() ([]byte, []int) {
	return fileDescriptor_edbe9e0eab89a1a7, []int{3}
}
func (m *C) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *C) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_C.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *C) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C.Merge(m, src)
}
func (m *C) XXX_Size() int {
	return m.Size()
}
func (m *C) XXX_DiscardUnknown() {
	xxx_messageInfo_C.DiscardUnknown(m)
}

var xxx_messageInfo_C proto.InternalMessageInfo

func (m *C) GetZ() bool {
	if m != nil {
		return m.Z
	}
	return false
}

func init() {
	proto.RegisterType((*ABC)(nil), "test.abc.ABC")
	proto.RegisterType((*A)(nil), "test.abc.A")
	proto.RegisterType((*B)(nil), "test.abc.B")
	proto.RegisterType((*C)(nil), "test.abc.C")
}

func init() { proto.RegisterFile("test/abc.proto", fileDescriptor_edbe9e0eab89a1a7) }

var fileDescriptor_edbe9e0eab89a1a7 = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x49, 0x2d, 0x2e,
	0xd1, 0x4f, 0x4c, 0x4a, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0xf1, 0xf5, 0x12,
	0x93, 0x92, 0xa5, 0x78, 0x92, 0xf3, 0x8b, 0x73, 0xf3, 0x8b, 0x21, 0xe2, 0x4a, 0x73, 0x19, 0xb9,
	0x98, 0x1d, 0x9d, 0x9c, 0x85, 0xa4, 0xb9, 0x18, 0x13, 0x25, 0x18, 0x15, 0x18, 0x35, 0xb8, 0x8d,
	0xb8, 0xf5, 0x60, 0x6a, 0xf5, 0x1c, 0x3d, 0x18, 0x82, 0x18, 0x13, 0x41, 0x92, 0x49, 0x12, 0x4c,
	0xe8, 0x92, 0x4e, 0x20, 0xc9, 0x24, 0x90, 0x64, 0xb2, 0x04, 0x33, 0xba, 0xa4, 0x33, 0x48, 0x32,
	0xd9, 0xca, 0xe2, 0xd4, 0x16, 0x5d, 0x93, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc,
	0x5c, 0xfd, 0xa2, 0xd4, 0xf4, 0xd4, 0x3c, 0xdd, 0xbc, 0xd4, 0x92, 0xf2, 0xfc, 0xa2, 0x6c, 0x7d,
	0x88, 0x4b, 0x74, 0xc1, 0x2e, 0xd1, 0x07, 0x3b, 0x38, 0x33, 0x2d, 0x31, 0x39, 0x55, 0xcf, 0xb7,
	0x38, 0xdd, 0x89, 0x95, 0x8b, 0xb9, 0xb8, 0x34, 0x57, 0x49, 0x90, 0x8b, 0xd1, 0x51, 0x88, 0x87,
	0x8b, 0xb1, 0x02, 0xec, 0x38, 0xd6, 0x20, 0xc6, 0x0a, 0x90, 0x90, 0x13, 0x48, 0xa8, 0x12, 0x2c,
	0xc4, 0x1b, 0xc4, 0x58, 0x09, 0x12, 0x72, 0x06, 0x09, 0x55, 0x81, 0x85, 0x38, 0x82, 0x18, 0xab,
	0x9c, 0xe4, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09,
	0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x8a, 0x05, 0x64, 0x53,
	0x12, 0x1b, 0xd8, 0x56, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb1, 0x4d, 0x04, 0x1f, 0x29,
	0x01, 0x00, 0x00,
}

func (this *ABC) GetMsg() github_com_regen_network_cosmos_proto_test_iface.Msg {
	if x := this.GetA(); x != nil {
		return x
	}
	if x := this.GetB(); x != nil {
		return x
	}
	if x := this.GetC(); x != nil {
		return x
	}
	return nil
}

func (this *ABC) SetMsg(value github_com_regen_network_cosmos_proto_test_iface.Msg) error {
	switch vt := value.(type) {
	case *A:
		this.Sum = &ABC_A{vt}
		return nil
	case A:
		x := new(A)
		*x = vt
		this.Sum = &ABC_A{x}
		return nil
	case *B:
		this.Sum = &ABC_B{vt}
		return nil
	case B:
		x := new(B)
		*x = vt
		this.Sum = &ABC_B{x}
		return nil
	case *C:
		this.Sum = &ABC_C{vt}
		return nil
	case C:
		x := new(C)
		*x = vt
		this.Sum = &ABC_C{x}
		return nil
	}
	return fmt.Errorf("can't encode value of type %T as message ABC", value)
}

func (m *ABC) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ABC) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ABC) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Sum != nil {
		{
			size := m.Sum.Size()
			i -= size
			if _, err := m.Sum.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *ABC_A) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ABC_A) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.A != nil {
		{
			size, err := m.A.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAbc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *ABC_B) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ABC_B) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.B != nil {
		{
			size, err := m.B.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAbc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *ABC_C) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ABC_C) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.C != nil {
		{
			size, err := m.C.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAbc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *A) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *A) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *A) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.X != 0 {
		i = encodeVarintAbc(dAtA, i, uint64(m.X))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *B) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *B) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *B) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Y != 0 {
		i = encodeVarintAbc(dAtA, i, uint64(m.Y))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *C) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *C) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *C) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Z {
		i--
		if m.Z {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintAbc(dAtA []byte, offset int, v uint64) int {
	offset -= sovAbc(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ABC) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sum != nil {
		n += m.Sum.Size()
	}
	return n
}

func (m *ABC_A) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.A != nil {
		l = m.A.Size()
		n += 1 + l + sovAbc(uint64(l))
	}
	return n
}
func (m *ABC_B) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.B != nil {
		l = m.B.Size()
		n += 1 + l + sovAbc(uint64(l))
	}
	return n
}
func (m *ABC_C) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.C != nil {
		l = m.C.Size()
		n += 1 + l + sovAbc(uint64(l))
	}
	return n
}
func (m *A) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.X != 0 {
		n += 1 + sovAbc(uint64(m.X))
	}
	return n
}

func (m *B) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Y != 0 {
		n += 1 + sovAbc(uint64(m.Y))
	}
	return n
}

func (m *C) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Z {
		n += 2
	}
	return n
}

func sovAbc(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAbc(x uint64) (n int) {
	return sovAbc(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ABC) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAbc
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
			return fmt.Errorf("proto: ABC: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ABC: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field A", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAbc
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
				return ErrInvalidLengthAbc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAbc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &A{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &ABC_A{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field B", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAbc
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
				return ErrInvalidLengthAbc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAbc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &B{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &ABC_B{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field C", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAbc
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
				return ErrInvalidLengthAbc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAbc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &C{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &ABC_C{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAbc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAbc
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAbc
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
func (m *A) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAbc
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
			return fmt.Errorf("proto: A: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: A: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field X", wireType)
			}
			m.X = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAbc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.X |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAbc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAbc
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAbc
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
func (m *B) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAbc
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
			return fmt.Errorf("proto: B: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: B: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Y", wireType)
			}
			m.Y = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAbc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Y |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAbc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAbc
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAbc
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
func (m *C) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAbc
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
			return fmt.Errorf("proto: C: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: C: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Z", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAbc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Z = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipAbc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAbc
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAbc
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
func skipAbc(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAbc
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
					return 0, ErrIntOverflowAbc
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
					return 0, ErrIntOverflowAbc
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
				return 0, ErrInvalidLengthAbc
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAbc
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAbc
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAbc        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAbc          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAbc = fmt.Errorf("proto: unexpected end of group")
)
