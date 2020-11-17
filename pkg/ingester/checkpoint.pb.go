// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pkg/ingester/checkpoint.proto

package ingester

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/cortexproject/cortex/pkg/ingester/client"
	github_com_cortexproject_cortex_pkg_ingester_client "github.com/cortexproject/cortex/pkg/ingester/client"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	io "io"
	math "math"
	reflect "reflect"
	strings "strings"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Chunk is a {de,}serializable intermediate type for chunkDesc which allows
// efficient loading/unloading to disk during WAL checkpoint recovery.
type Chunk struct {
	From      time.Time `protobuf:"bytes,1,opt,name=from,proto3,stdtime" json:"from"`
	To        time.Time `protobuf:"bytes,2,opt,name=to,proto3,stdtime" json:"to"`
	FlushedAt time.Time `protobuf:"bytes,3,opt,name=flushedAt,proto3,stdtime" json:"flushedAt"`
	Closed    bool      `protobuf:"varint,4,opt,name=closed,proto3" json:"closed,omitempty"`
	Data      []byte    `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Chunk) Reset()      { *m = Chunk{} }
func (*Chunk) ProtoMessage() {}
func (*Chunk) Descriptor() ([]byte, []int) {
	return fileDescriptor_00f4b7152db9bdb5, []int{0}
}
func (m *Chunk) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Chunk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Chunk.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Chunk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Chunk.Merge(m, src)
}
func (m *Chunk) XXX_Size() int {
	return m.Size()
}
func (m *Chunk) XXX_DiscardUnknown() {
	xxx_messageInfo_Chunk.DiscardUnknown(m)
}

var xxx_messageInfo_Chunk proto.InternalMessageInfo

func (m *Chunk) GetFrom() time.Time {
	if m != nil {
		return m.From
	}
	return time.Time{}
}

func (m *Chunk) GetTo() time.Time {
	if m != nil {
		return m.To
	}
	return time.Time{}
}

func (m *Chunk) GetFlushedAt() time.Time {
	if m != nil {
		return m.FlushedAt
	}
	return time.Time{}
}

func (m *Chunk) GetClosed() bool {
	if m != nil {
		return m.Closed
	}
	return false
}

func (m *Chunk) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// Series is a {de,}serializable intermediate type for Series.
type Series struct {
	UserID      string                                                             `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Fingerprint uint64                                                             `protobuf:"varint,2,opt,name=fingerprint,proto3" json:"fingerprint,omitempty"`
	Labels      []github_com_cortexproject_cortex_pkg_ingester_client.LabelAdapter `protobuf:"bytes,3,rep,name=labels,proto3,customtype=github.com/cortexproject/cortex/pkg/ingester/client.LabelAdapter" json:"labels"`
	Chunks      []Chunk                                                            `protobuf:"bytes,4,rep,name=chunks,proto3" json:"chunks"`
}

func (m *Series) Reset()      { *m = Series{} }
func (*Series) ProtoMessage() {}
func (*Series) Descriptor() ([]byte, []int) {
	return fileDescriptor_00f4b7152db9bdb5, []int{1}
}
func (m *Series) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Series) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Series.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Series) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Series.Merge(m, src)
}
func (m *Series) XXX_Size() int {
	return m.Size()
}
func (m *Series) XXX_DiscardUnknown() {
	xxx_messageInfo_Series.DiscardUnknown(m)
}

var xxx_messageInfo_Series proto.InternalMessageInfo

func (m *Series) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *Series) GetFingerprint() uint64 {
	if m != nil {
		return m.Fingerprint
	}
	return 0
}

func (m *Series) GetChunks() []Chunk {
	if m != nil {
		return m.Chunks
	}
	return nil
}

func init() {
	proto.RegisterType((*Chunk)(nil), "ingester.Chunk")
	proto.RegisterType((*Series)(nil), "ingester.Series")
}

func init() { proto.RegisterFile("pkg/ingester/checkpoint.proto", fileDescriptor_00f4b7152db9bdb5) }

var fileDescriptor_00f4b7152db9bdb5 = []byte{
	// 427 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0xbd, 0x8e, 0xd4, 0x30,
	0x10, 0x8e, 0x77, 0x73, 0xd1, 0x9e, 0x17, 0x09, 0xe1, 0x02, 0x45, 0x2b, 0xe1, 0x44, 0x57, 0xa5,
	0x39, 0x47, 0x3a, 0x28, 0xa8, 0x10, 0x17, 0x28, 0x40, 0xa2, 0x40, 0x81, 0x8a, 0x2e, 0x3f, 0x4e,
	0x62, 0x36, 0x89, 0x23, 0xdb, 0x91, 0x28, 0x79, 0x84, 0x7b, 0x0c, 0x1e, 0xe5, 0xca, 0x2d, 0x4f,
	0x14, 0x07, 0x9b, 0x95, 0x80, 0xf2, 0x1e, 0x01, 0xc5, 0x49, 0xd8, 0xa5, 0xdc, 0xeb, 0xfc, 0xcd,
	0x7c, 0xdf, 0xcc, 0x37, 0x33, 0x86, 0x4f, 0x9a, 0x75, 0xee, 0xb3, 0x3a, 0xa7, 0x52, 0x51, 0xe1,
	0x27, 0x05, 0x4d, 0xd6, 0x0d, 0x67, 0xb5, 0x22, 0x8d, 0xe0, 0x8a, 0xa3, 0xc5, 0x94, 0x5a, 0x39,
	0x39, 0xe7, 0x79, 0x49, 0x7d, 0x1d, 0x8f, 0xdb, 0xcc, 0x57, 0xac, 0xa2, 0x52, 0x45, 0x55, 0x33,
	0x50, 0x57, 0xe7, 0x39, 0x53, 0x45, 0x1b, 0x93, 0x84, 0x57, 0x7e, 0xce, 0x73, 0xbe, 0x67, 0xf6,
	0x48, 0x03, 0xfd, 0x1a, 0xe9, 0x2f, 0x0f, 0xe8, 0x09, 0x17, 0x8a, 0x7e, 0x69, 0x04, 0xff, 0x4c,
	0x13, 0x35, 0x22, 0xff, 0x7f, 0x63, 0x25, 0xa3, 0xf5, 0x94, 0x1a, 0x2a, 0x9c, 0xfd, 0x06, 0xf0,
	0xe4, 0x55, 0xd1, 0xd6, 0x6b, 0xf4, 0x1c, 0x9a, 0x99, 0xe0, 0x95, 0x0d, 0x5c, 0xe0, 0x2d, 0x2f,
	0x56, 0x64, 0xb0, 0x4a, 0x26, 0x03, 0xe4, 0xe3, 0x64, 0x35, 0x58, 0x5c, 0xdf, 0x3a, 0xc6, 0xd5,
	0x0f, 0x07, 0x84, 0x5a, 0x81, 0x9e, 0xc1, 0x99, 0xe2, 0xf6, 0xec, 0x08, 0xdd, 0x4c, 0x71, 0x14,
	0xc0, 0xd3, 0xac, 0x6c, 0x65, 0x41, 0xd3, 0x4b, 0x65, 0xcf, 0x8f, 0x10, 0xef, 0x65, 0xe8, 0x31,
	0xb4, 0x92, 0x92, 0x4b, 0x9a, 0xda, 0xa6, 0x0b, 0xbc, 0x45, 0x38, 0x22, 0x84, 0xa0, 0x99, 0x46,
	0x2a, 0xb2, 0x4f, 0x5c, 0xe0, 0x3d, 0x08, 0xf5, 0xfb, 0xec, 0x17, 0x80, 0xd6, 0x07, 0x2a, 0x18,
	0x95, 0xbd, 0xac, 0x95, 0x54, 0xbc, 0x7d, 0xad, 0x87, 0x3d, 0x0d, 0x47, 0x84, 0x5c, 0xb8, 0xcc,
	0xfa, 0x6d, 0x89, 0x46, 0xb0, 0x5a, 0xe9, 0x89, 0xcc, 0xf0, 0x30, 0x84, 0x24, 0xb4, 0xca, 0x28,
	0xa6, 0xa5, 0xb4, 0xe7, 0xee, 0xdc, 0x5b, 0x5e, 0x3c, 0x22, 0xe3, 0x36, 0xdf, 0xf5, 0xd1, 0xf7,
	0x11, 0x13, 0xc1, 0x9b, 0xde, 0xe8, 0xf7, 0x5b, 0xe7, 0x3e, 0xb7, 0x19, 0xca, 0x5c, 0xa6, 0x51,
	0xa3, 0xa8, 0x08, 0xc7, 0x56, 0xe8, 0x1c, 0x5a, 0x49, 0x7f, 0x22, 0x69, 0x9b, 0xba, 0xe9, 0x43,
	0x32, 0xc9, 0x88, 0x3e, 0x5d, 0x60, 0xf6, 0x2d, 0xc3, 0x91, 0x14, 0xbc, 0xd8, 0x6c, 0xb1, 0x71,
	0xb3, 0xc5, 0xc6, 0xdd, 0x16, 0x83, 0xaf, 0x1d, 0x06, 0xdf, 0x3a, 0x0c, 0xae, 0x3b, 0x0c, 0x36,
	0x1d, 0x06, 0x3f, 0x3b, 0x0c, 0xfe, 0x74, 0xd8, 0xb8, 0xeb, 0x30, 0xb8, 0xda, 0x61, 0x63, 0xb3,
	0xc3, 0xc6, 0xcd, 0x0e, 0x1b, 0x9f, 0xfe, 0x7d, 0xd2, 0xd8, 0xd2, 0xdb, 0x7f, 0xfa, 0x37, 0x00,
	0x00, 0xff, 0xff, 0xaf, 0xc1, 0x8c, 0xad, 0xd6, 0x02, 0x00, 0x00,
}

func (this *Chunk) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Chunk)
	if !ok {
		that2, ok := that.(Chunk)
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
	if !this.From.Equal(that1.From) {
		return false
	}
	if !this.To.Equal(that1.To) {
		return false
	}
	if !this.FlushedAt.Equal(that1.FlushedAt) {
		return false
	}
	if this.Closed != that1.Closed {
		return false
	}
	if !bytes.Equal(this.Data, that1.Data) {
		return false
	}
	return true
}
func (this *Series) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Series)
	if !ok {
		that2, ok := that.(Series)
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
	if this.UserID != that1.UserID {
		return false
	}
	if this.Fingerprint != that1.Fingerprint {
		return false
	}
	if len(this.Labels) != len(that1.Labels) {
		return false
	}
	for i := range this.Labels {
		if !this.Labels[i].Equal(that1.Labels[i]) {
			return false
		}
	}
	if len(this.Chunks) != len(that1.Chunks) {
		return false
	}
	for i := range this.Chunks {
		if !this.Chunks[i].Equal(&that1.Chunks[i]) {
			return false
		}
	}
	return true
}
func (this *Chunk) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 9)
	s = append(s, "&ingester.Chunk{")
	s = append(s, "From: "+fmt.Sprintf("%#v", this.From)+",\n")
	s = append(s, "To: "+fmt.Sprintf("%#v", this.To)+",\n")
	s = append(s, "FlushedAt: "+fmt.Sprintf("%#v", this.FlushedAt)+",\n")
	s = append(s, "Closed: "+fmt.Sprintf("%#v", this.Closed)+",\n")
	s = append(s, "Data: "+fmt.Sprintf("%#v", this.Data)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Series) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&ingester.Series{")
	s = append(s, "UserID: "+fmt.Sprintf("%#v", this.UserID)+",\n")
	s = append(s, "Fingerprint: "+fmt.Sprintf("%#v", this.Fingerprint)+",\n")
	s = append(s, "Labels: "+fmt.Sprintf("%#v", this.Labels)+",\n")
	if this.Chunks != nil {
		vs := make([]*Chunk, len(this.Chunks))
		for i := range vs {
			vs[i] = &this.Chunks[i]
		}
		s = append(s, "Chunks: "+fmt.Sprintf("%#v", vs)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringCheckpoint(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Chunk) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Chunk) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintCheckpoint(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(m.From)))
	n1, err := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.From, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x12
	i++
	i = encodeVarintCheckpoint(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(m.To)))
	n2, err := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.To, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	dAtA[i] = 0x1a
	i++
	i = encodeVarintCheckpoint(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(m.FlushedAt)))
	n3, err := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.FlushedAt, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	if m.Closed {
		dAtA[i] = 0x20
		i++
		if m.Closed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.Data) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintCheckpoint(dAtA, i, uint64(len(m.Data)))
		i += copy(dAtA[i:], m.Data)
	}
	return i, nil
}

func (m *Series) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Series) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.UserID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCheckpoint(dAtA, i, uint64(len(m.UserID)))
		i += copy(dAtA[i:], m.UserID)
	}
	if m.Fingerprint != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintCheckpoint(dAtA, i, uint64(m.Fingerprint))
	}
	if len(m.Labels) > 0 {
		for _, msg := range m.Labels {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintCheckpoint(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Chunks) > 0 {
		for _, msg := range m.Chunks {
			dAtA[i] = 0x22
			i++
			i = encodeVarintCheckpoint(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintCheckpoint(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Chunk) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.From)
	n += 1 + l + sovCheckpoint(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.To)
	n += 1 + l + sovCheckpoint(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.FlushedAt)
	n += 1 + l + sovCheckpoint(uint64(l))
	if m.Closed {
		n += 2
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovCheckpoint(uint64(l))
	}
	return n
}

func (m *Series) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UserID)
	if l > 0 {
		n += 1 + l + sovCheckpoint(uint64(l))
	}
	if m.Fingerprint != 0 {
		n += 1 + sovCheckpoint(uint64(m.Fingerprint))
	}
	if len(m.Labels) > 0 {
		for _, e := range m.Labels {
			l = e.Size()
			n += 1 + l + sovCheckpoint(uint64(l))
		}
	}
	if len(m.Chunks) > 0 {
		for _, e := range m.Chunks {
			l = e.Size()
			n += 1 + l + sovCheckpoint(uint64(l))
		}
	}
	return n
}

func sovCheckpoint(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCheckpoint(x uint64) (n int) {
	return sovCheckpoint(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Chunk) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Chunk{`,
		`From:` + strings.Replace(strings.Replace(this.From.String(), "Timestamp", "types.Timestamp", 1), `&`, ``, 1) + `,`,
		`To:` + strings.Replace(strings.Replace(this.To.String(), "Timestamp", "types.Timestamp", 1), `&`, ``, 1) + `,`,
		`FlushedAt:` + strings.Replace(strings.Replace(this.FlushedAt.String(), "Timestamp", "types.Timestamp", 1), `&`, ``, 1) + `,`,
		`Closed:` + fmt.Sprintf("%v", this.Closed) + `,`,
		`Data:` + fmt.Sprintf("%v", this.Data) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Series) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Series{`,
		`UserID:` + fmt.Sprintf("%v", this.UserID) + `,`,
		`Fingerprint:` + fmt.Sprintf("%v", this.Fingerprint) + `,`,
		`Labels:` + fmt.Sprintf("%v", this.Labels) + `,`,
		`Chunks:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Chunks), "Chunk", "Chunk", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringCheckpoint(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Chunk) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCheckpoint
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
			return fmt.Errorf("proto: Chunk: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Chunk: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCheckpoint
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
				return ErrInvalidLengthCheckpoint
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCheckpoint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.From, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field To", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCheckpoint
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
				return ErrInvalidLengthCheckpoint
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCheckpoint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.To, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FlushedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCheckpoint
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
				return ErrInvalidLengthCheckpoint
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCheckpoint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.FlushedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Closed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCheckpoint
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
			m.Closed = bool(v != 0)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCheckpoint
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
				return ErrInvalidLengthCheckpoint
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCheckpoint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCheckpoint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCheckpoint
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCheckpoint
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
func (m *Series) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCheckpoint
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
			return fmt.Errorf("proto: Series: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Series: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCheckpoint
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
				return ErrInvalidLengthCheckpoint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCheckpoint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fingerprint", wireType)
			}
			m.Fingerprint = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCheckpoint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Fingerprint |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Labels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCheckpoint
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
				return ErrInvalidLengthCheckpoint
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCheckpoint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Labels = append(m.Labels, github_com_cortexproject_cortex_pkg_ingester_client.LabelAdapter{})
			if err := m.Labels[len(m.Labels)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chunks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCheckpoint
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
				return ErrInvalidLengthCheckpoint
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCheckpoint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chunks = append(m.Chunks, Chunk{})
			if err := m.Chunks[len(m.Chunks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCheckpoint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCheckpoint
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCheckpoint
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
func skipCheckpoint(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCheckpoint
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
					return 0, ErrIntOverflowCheckpoint
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
					return 0, ErrIntOverflowCheckpoint
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
				return 0, ErrInvalidLengthCheckpoint
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthCheckpoint
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCheckpoint
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
				next, err := skipCheckpoint(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthCheckpoint
				}
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
	ErrInvalidLengthCheckpoint = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCheckpoint   = fmt.Errorf("proto: integer overflow")
)
