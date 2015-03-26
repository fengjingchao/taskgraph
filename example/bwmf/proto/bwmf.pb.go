// Code generated by protoc-gen-go.
// source: bwmf.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	bwmf.proto

It has these top-level messages:
	Request
	Response
	DenseMatrixShard
	SparseMatrixShard
*/
package proto

import proto1 "github.com/golang/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal

// Request for block matrix data at a given epoch.
type Request struct {
	Epoch uint64 `protobuf:"varint,1,opt,name=epoch" json:"epoch,omitempty"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto1.CompactTextString(m) }
func (*Request) ProtoMessage()    {}

// Response of the block matrix data, with the associated block id.
type Response struct {
	Id    int32             `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Shard *DenseMatrixShard `protobuf:"bytes,2,opt,name=shard" json:"shard,omitempty"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto1.CompactTextString(m) }
func (*Response) ProtoMessage()    {}

func (m *Response) GetShard() *DenseMatrixShard {
	if m != nil {
		return m.Shard
	}
	return nil
}

// sharded dense matrix data
// In bwmf, this message is for full D/T and sharded D/T data blocks.
type DenseMatrixShard struct {
	Row []*DenseMatrixShard_DenseRow `protobuf:"bytes,1,rep,name=row" json:"row,omitempty"`
}

func (m *DenseMatrixShard) Reset()         { *m = DenseMatrixShard{} }
func (m *DenseMatrixShard) String() string { return proto1.CompactTextString(m) }
func (*DenseMatrixShard) ProtoMessage()    {}

func (m *DenseMatrixShard) GetRow() []*DenseMatrixShard_DenseRow {
	if m != nil {
		return m.Row
	}
	return nil
}

type DenseMatrixShard_DenseRow struct {
	At []float32 `protobuf:"fixed32,1,rep,name=at" json:"at,omitempty"`
}

func (m *DenseMatrixShard_DenseRow) Reset()         { *m = DenseMatrixShard_DenseRow{} }
func (m *DenseMatrixShard_DenseRow) String() string { return proto1.CompactTextString(m) }
func (*DenseMatrixShard_DenseRow) ProtoMessage()    {}

// sharded matrix data
// In bwmf, this message is for sharded row/column data.
type SparseMatrixShard struct {
	Row []*SparseMatrixShard_SparseRow `protobuf:"bytes,1,rep,name=row" json:"row,omitempty"`
}

func (m *SparseMatrixShard) Reset()         { *m = SparseMatrixShard{} }
func (m *SparseMatrixShard) String() string { return proto1.CompactTextString(m) }
func (*SparseMatrixShard) ProtoMessage()    {}

func (m *SparseMatrixShard) GetRow() []*SparseMatrixShard_SparseRow {
	if m != nil {
		return m.Row
	}
	return nil
}

type SparseMatrixShard_SparseRow struct {
	At map[uint32]float32 `protobuf:"bytes,1,rep,name=at" json:"at,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"fixed32,2,opt,name=value"`
}

func (m *SparseMatrixShard_SparseRow) Reset()         { *m = SparseMatrixShard_SparseRow{} }
func (m *SparseMatrixShard_SparseRow) String() string { return proto1.CompactTextString(m) }
func (*SparseMatrixShard_SparseRow) ProtoMessage()    {}

func (m *SparseMatrixShard_SparseRow) GetAt() map[uint32]float32 {
	if m != nil {
		return m.At
	}
	return nil
}

func init() {
}
