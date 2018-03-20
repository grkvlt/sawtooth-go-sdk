// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/grkvlt/sawtooth-go-sdk/protobuf/block_pb2/block.proto

/*
Package block_pb2 is a generated protocol buffer package.

It is generated from these files:
	github.com/grkvlt/sawtooth-go-sdk/protobuf/block_pb2/block.proto

It has these top-level messages:
	BlockHeader
	Block
*/
package block_pb2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import batch "github.com/grkvlt/sawtooth-go-sdk/protobuf/batch_pb2"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type BlockHeader struct {
	// Block number in the chain
	BlockNum uint64 `protobuf:"varint,1,opt,name=block_num,json=blockNum" json:"block_num,omitempty"`
	// The header_signature of the previous block that was added to the chain.
	PreviousBlockId string `protobuf:"bytes,2,opt,name=previous_block_id,json=previousBlockId" json:"previous_block_id,omitempty"`
	// Public key for the component internal to the validator that
	// signed the BlockHeader
	SignerPublicKey string `protobuf:"bytes,3,opt,name=signer_public_key,json=signerPublicKey" json:"signer_public_key,omitempty"`
	// List of batch.header_signatures that match the order of batches
	// required for the block
	BatchIds []string `protobuf:"bytes,4,rep,name=batch_ids,json=batchIds" json:"batch_ids,omitempty"`
	// Bytes that are set and verified by the consensus algorithm used to
	// create and validate the block
	Consensus []byte `protobuf:"bytes,5,opt,name=consensus,proto3" json:"consensus,omitempty"`
	// The state_root_hash should match the final state_root after all
	// transactions in the batches have been applied, otherwise the block
	// is not valid
	StateRootHash string `protobuf:"bytes,6,opt,name=state_root_hash,json=stateRootHash" json:"state_root_hash,omitempty"`
}

func (m *BlockHeader) Reset()                    { *m = BlockHeader{} }
func (m *BlockHeader) String() string            { return proto.CompactTextString(m) }
func (*BlockHeader) ProtoMessage()               {}
func (*BlockHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BlockHeader) GetBlockNum() uint64 {
	if m != nil {
		return m.BlockNum
	}
	return 0
}

func (m *BlockHeader) GetPreviousBlockId() string {
	if m != nil {
		return m.PreviousBlockId
	}
	return ""
}

func (m *BlockHeader) GetSignerPublicKey() string {
	if m != nil {
		return m.SignerPublicKey
	}
	return ""
}

func (m *BlockHeader) GetBatchIds() []string {
	if m != nil {
		return m.BatchIds
	}
	return nil
}

func (m *BlockHeader) GetConsensus() []byte {
	if m != nil {
		return m.Consensus
	}
	return nil
}

func (m *BlockHeader) GetStateRootHash() string {
	if m != nil {
		return m.StateRootHash
	}
	return ""
}

type Block struct {
	// The serialized version of the BlockHeader
	Header []byte `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// The signature derived from signing the header
	HeaderSignature string `protobuf:"bytes,2,opt,name=header_signature,json=headerSignature" json:"header_signature,omitempty"`
	// A list of batches. The batches may contain new batches that other
	// validators may not have received yet, or they will be all batches needed
	// for block validation when passed to the journal
	Batches []*batch.Batch `protobuf:"bytes,3,rep,name=batches" json:"batches,omitempty"`
}

func (m *Block) Reset()                    { *m = Block{} }
func (m *Block) String() string            { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()               {}
func (*Block) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Block) GetHeader() []byte {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Block) GetHeaderSignature() string {
	if m != nil {
		return m.HeaderSignature
	}
	return ""
}

func (m *Block) GetBatches() []*batch.Batch {
	if m != nil {
		return m.Batches
	}
	return nil
}

func init() {
	proto.RegisterType((*BlockHeader)(nil), "BlockHeader")
	proto.RegisterType((*Block)(nil), "Block")
}

func init() { proto.RegisterFile("github.com/grkvlt/sawtooth-go-sdk/protobuf/block_pb2/block.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x91, 0x51, 0x4b, 0xc3, 0x30,
	0x10, 0xc7, 0xa9, 0xdb, 0xea, 0x9a, 0x4d, 0xa6, 0x01, 0xa5, 0xa8, 0x0f, 0x63, 0x88, 0x4c, 0x85,
	0x0e, 0xe6, 0x37, 0xd8, 0xd3, 0x86, 0x20, 0xa3, 0xbe, 0xf9, 0x12, 0xda, 0x26, 0xda, 0xb2, 0xae,
	0x29, 0xb9, 0x44, 0xd9, 0x07, 0xf6, 0x7b, 0x98, 0x5e, 0x8c, 0x3e, 0xf9, 0x76, 0xf9, 0xdd, 0x3f,
	0x77, 0xf7, 0xe7, 0x4f, 0x1e, 0x20, 0xfb, 0xd4, 0x52, 0xea, 0x92, 0x01, 0xdf, 0x2d, 0x5a, 0x25,
	0xb5, 0xcc, 0xcd, 0xdb, 0x22, 0xaf, 0x65, 0xb1, 0x63, 0x6d, 0xbe, 0x74, 0x55, 0x82, 0x8d, 0xcb,
	0xff, 0xc4, 0x99, 0x2e, 0x4a, 0x27, 0xee, 0x2a, 0x27, 0x9e, 0x7d, 0x05, 0x64, 0xb4, 0xea, 0x3e,
	0xaf, 0x45, 0xc6, 0x85, 0xa2, 0x57, 0x24, 0x72, 0x53, 0x1b, 0xb3, 0x8f, 0x83, 0x69, 0x30, 0xef,
	0xa7, 0x43, 0x04, 0xcf, 0x66, 0x4f, 0xef, 0xc9, 0x59, 0xab, 0xc4, 0x47, 0x25, 0x0d, 0x30, 0xa7,
	0xaa, 0x78, 0x7c, 0x64, 0x45, 0x51, 0x3a, 0xf1, 0x0d, 0x1c, 0xb6, 0xe1, 0x9d, 0x16, 0xaa, 0xf7,
	0x46, 0x28, 0xd6, 0x9a, 0xbc, 0xae, 0x0a, 0xb6, 0x13, 0x87, 0xb8, 0xe7, 0xb4, 0xae, 0xb1, 0x45,
	0xfe, 0x24, 0x0e, 0xb8, 0x14, 0xaf, 0xab, 0x38, 0xc4, 0xfd, 0x69, 0xcf, 0x6a, 0x86, 0x08, 0x36,
	0x1c, 0xe8, 0x35, 0x89, 0x0a, 0xd9, 0x80, 0x68, 0xc0, 0x40, 0x3c, 0xb0, 0x03, 0xc6, 0xe9, 0x1f,
	0xa0, 0xb7, 0x64, 0x02, 0x3a, 0xd3, 0x82, 0x29, 0xeb, 0x98, 0x95, 0x19, 0x94, 0x71, 0x88, 0x4b,
	0x4e, 0x10, 0xa7, 0x96, 0xae, 0x2d, 0x9c, 0xd5, 0x64, 0x80, 0x97, 0xd1, 0x0b, 0x12, 0x96, 0x68,
	0x15, 0xdd, 0x8d, 0xd3, 0x9f, 0x17, 0xbd, 0x23, 0xa7, 0xae, 0x62, 0xdd, 0x75, 0x99, 0x36, 0x4a,
	0x78, 0x6b, 0x8e, 0xbf, 0x78, 0x4c, 0xa7, 0xe4, 0x18, 0xaf, 0x13, 0x60, 0x0d, 0xf5, 0xe6, 0xa3,
	0x65, 0x98, 0xac, 0xba, 0x77, 0xea, 0xf1, 0xea, 0x86, 0x9c, 0xfb, 0x10, 0x12, 0x1b, 0x42, 0xe2,
	0x43, 0xd8, 0x06, 0xaf, 0xd1, 0x6f, 0x68, 0x79, 0x88, 0xf8, 0xf1, 0x3b, 0x00, 0x00, 0xff, 0xff,
	0x28, 0xac, 0x53, 0x9f, 0xde, 0x01, 0x00, 0x00,
}
