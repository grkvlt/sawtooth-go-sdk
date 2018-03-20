// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/grkvlt/sawtooth-go-sdk/protobuf/client_batch_submit_pb2/client_batch_submit.proto

/*
Package client_batch_submit_pb2 is a generated protocol buffer package.

It is generated from these files:
	github.com/grkvlt/sawtooth-go-sdk/protobuf/client_batch_submit_pb2/client_batch_submit.proto

It has these top-level messages:
	ClientBatchStatus
	ClientBatchSubmitRequest
	ClientBatchSubmitResponse
	ClientBatchStatusRequest
	ClientBatchStatusResponse
*/
package client_batch_submit_pb2

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

type ClientBatchStatus_Status int32

const (
	ClientBatchStatus_STATUS_UNSET ClientBatchStatus_Status = 0
	ClientBatchStatus_COMMITTED    ClientBatchStatus_Status = 1
	ClientBatchStatus_INVALID      ClientBatchStatus_Status = 2
	ClientBatchStatus_PENDING      ClientBatchStatus_Status = 3
	ClientBatchStatus_UNKNOWN      ClientBatchStatus_Status = 4
)

var ClientBatchStatus_Status_name = map[int32]string{
	0: "STATUS_UNSET",
	1: "COMMITTED",
	2: "INVALID",
	3: "PENDING",
	4: "UNKNOWN",
}
var ClientBatchStatus_Status_value = map[string]int32{
	"STATUS_UNSET": 0,
	"COMMITTED":    1,
	"INVALID":      2,
	"PENDING":      3,
	"UNKNOWN":      4,
}

func (x ClientBatchStatus_Status) String() string {
	return proto.EnumName(ClientBatchStatus_Status_name, int32(x))
}
func (ClientBatchStatus_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type ClientBatchSubmitResponse_Status int32

const (
	ClientBatchSubmitResponse_STATUS_UNSET   ClientBatchSubmitResponse_Status = 0
	ClientBatchSubmitResponse_OK             ClientBatchSubmitResponse_Status = 1
	ClientBatchSubmitResponse_INTERNAL_ERROR ClientBatchSubmitResponse_Status = 2
	ClientBatchSubmitResponse_INVALID_BATCH  ClientBatchSubmitResponse_Status = 3
	ClientBatchSubmitResponse_QUEUE_FULL     ClientBatchSubmitResponse_Status = 4
)

var ClientBatchSubmitResponse_Status_name = map[int32]string{
	0: "STATUS_UNSET",
	1: "OK",
	2: "INTERNAL_ERROR",
	3: "INVALID_BATCH",
	4: "QUEUE_FULL",
}
var ClientBatchSubmitResponse_Status_value = map[string]int32{
	"STATUS_UNSET":   0,
	"OK":             1,
	"INTERNAL_ERROR": 2,
	"INVALID_BATCH":  3,
	"QUEUE_FULL":     4,
}

func (x ClientBatchSubmitResponse_Status) String() string {
	return proto.EnumName(ClientBatchSubmitResponse_Status_name, int32(x))
}
func (ClientBatchSubmitResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{2, 0}
}

type ClientBatchStatusResponse_Status int32

const (
	ClientBatchStatusResponse_STATUS_UNSET   ClientBatchStatusResponse_Status = 0
	ClientBatchStatusResponse_OK             ClientBatchStatusResponse_Status = 1
	ClientBatchStatusResponse_INTERNAL_ERROR ClientBatchStatusResponse_Status = 2
	ClientBatchStatusResponse_NO_RESOURCE    ClientBatchStatusResponse_Status = 5
	ClientBatchStatusResponse_INVALID_ID     ClientBatchStatusResponse_Status = 8
)

var ClientBatchStatusResponse_Status_name = map[int32]string{
	0: "STATUS_UNSET",
	1: "OK",
	2: "INTERNAL_ERROR",
	5: "NO_RESOURCE",
	8: "INVALID_ID",
}
var ClientBatchStatusResponse_Status_value = map[string]int32{
	"STATUS_UNSET":   0,
	"OK":             1,
	"INTERNAL_ERROR": 2,
	"NO_RESOURCE":    5,
	"INVALID_ID":     8,
}

func (x ClientBatchStatusResponse_Status) String() string {
	return proto.EnumName(ClientBatchStatusResponse_Status_name, int32(x))
}
func (ClientBatchStatusResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{4, 0}
}

// Information about the status of a batch submitted to the validator.
//
// Attributes:
//     batch_id: The id (header_signature) of the batch
//     status: The committed status of the batch
//     invalid_transactions: Info for transactions that failed, if any
//
// Statuses:
//     COMMITTED - the batch was accepted and has been committed to the chain
//     INVALID - the batch failed validation, it should be resubmitted
//     PENDING - the batch is still being processed
//     UNKNOWN - no status for the batch could be found (possibly invalid)
type ClientBatchStatus struct {
	BatchId             string                                  `protobuf:"bytes,1,opt,name=batch_id,json=batchId" json:"batch_id,omitempty"`
	Status              ClientBatchStatus_Status                `protobuf:"varint,2,opt,name=status,enum=ClientBatchStatus_Status" json:"status,omitempty"`
	InvalidTransactions []*ClientBatchStatus_InvalidTransaction `protobuf:"bytes,3,rep,name=invalid_transactions,json=invalidTransactions" json:"invalid_transactions,omitempty"`
}

func (m *ClientBatchStatus) Reset()                    { *m = ClientBatchStatus{} }
func (m *ClientBatchStatus) String() string            { return proto.CompactTextString(m) }
func (*ClientBatchStatus) ProtoMessage()               {}
func (*ClientBatchStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ClientBatchStatus) GetBatchId() string {
	if m != nil {
		return m.BatchId
	}
	return ""
}

func (m *ClientBatchStatus) GetStatus() ClientBatchStatus_Status {
	if m != nil {
		return m.Status
	}
	return ClientBatchStatus_STATUS_UNSET
}

func (m *ClientBatchStatus) GetInvalidTransactions() []*ClientBatchStatus_InvalidTransaction {
	if m != nil {
		return m.InvalidTransactions
	}
	return nil
}

type ClientBatchStatus_InvalidTransaction struct {
	TransactionId string `protobuf:"bytes,1,opt,name=transaction_id,json=transactionId" json:"transaction_id,omitempty"`
	Message       string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	ExtendedData  []byte `protobuf:"bytes,3,opt,name=extended_data,json=extendedData,proto3" json:"extended_data,omitempty"`
}

func (m *ClientBatchStatus_InvalidTransaction) Reset()         { *m = ClientBatchStatus_InvalidTransaction{} }
func (m *ClientBatchStatus_InvalidTransaction) String() string { return proto.CompactTextString(m) }
func (*ClientBatchStatus_InvalidTransaction) ProtoMessage()    {}
func (*ClientBatchStatus_InvalidTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0}
}

func (m *ClientBatchStatus_InvalidTransaction) GetTransactionId() string {
	if m != nil {
		return m.TransactionId
	}
	return ""
}

func (m *ClientBatchStatus_InvalidTransaction) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ClientBatchStatus_InvalidTransaction) GetExtendedData() []byte {
	if m != nil {
		return m.ExtendedData
	}
	return nil
}

// Submits a list of Batches to be added to the blockchain.
type ClientBatchSubmitRequest struct {
	Batches []*batch.Batch `protobuf:"bytes,1,rep,name=batches" json:"batches,omitempty"`
}

func (m *ClientBatchSubmitRequest) Reset()                    { *m = ClientBatchSubmitRequest{} }
func (m *ClientBatchSubmitRequest) String() string            { return proto.CompactTextString(m) }
func (*ClientBatchSubmitRequest) ProtoMessage()               {}
func (*ClientBatchSubmitRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ClientBatchSubmitRequest) GetBatches() []*batch.Batch {
	if m != nil {
		return m.Batches
	}
	return nil
}

// This is a response to a submission of one or more Batches.
// Statuses:
//   * OK - everything with the request worked as expected
//   * INTERNAL_ERROR - general error, such as protobuf failing to deserialize
//   * INVALID_BATCH - the batch failed validation, likely due to a bad signature
//   * QUEUE_FULL - the batch is unable to be queued for processing, due to
//        a full processing queue.  The batch may be submitted again.
type ClientBatchSubmitResponse struct {
	Status ClientBatchSubmitResponse_Status `protobuf:"varint,1,opt,name=status,enum=ClientBatchSubmitResponse_Status" json:"status,omitempty"`
}

func (m *ClientBatchSubmitResponse) Reset()                    { *m = ClientBatchSubmitResponse{} }
func (m *ClientBatchSubmitResponse) String() string            { return proto.CompactTextString(m) }
func (*ClientBatchSubmitResponse) ProtoMessage()               {}
func (*ClientBatchSubmitResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ClientBatchSubmitResponse) GetStatus() ClientBatchSubmitResponse_Status {
	if m != nil {
		return m.Status
	}
	return ClientBatchSubmitResponse_STATUS_UNSET
}

// A request for the status of one or more batches, specified by id.
// If `wait` is set to true, the validator will wait to respond until all
// batches are committed, or until the specified `timeout` in seconds has
// elapsed. Defaults to 300.
type ClientBatchStatusRequest struct {
	BatchIds []string `protobuf:"bytes,1,rep,name=batch_ids,json=batchIds" json:"batch_ids,omitempty"`
	Wait     bool     `protobuf:"varint,2,opt,name=wait" json:"wait,omitempty"`
	Timeout  uint32   `protobuf:"varint,3,opt,name=timeout" json:"timeout,omitempty"`
}

func (m *ClientBatchStatusRequest) Reset()                    { *m = ClientBatchStatusRequest{} }
func (m *ClientBatchStatusRequest) String() string            { return proto.CompactTextString(m) }
func (*ClientBatchStatusRequest) ProtoMessage()               {}
func (*ClientBatchStatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ClientBatchStatusRequest) GetBatchIds() []string {
	if m != nil {
		return m.BatchIds
	}
	return nil
}

func (m *ClientBatchStatusRequest) GetWait() bool {
	if m != nil {
		return m.Wait
	}
	return false
}

func (m *ClientBatchStatusRequest) GetTimeout() uint32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

// This is a response to a request for the status of specific batches.
// Statuses:
//   * OK - everything with the request worked as expected
//   * INTERNAL_ERROR - general error, such as protobuf failing to deserialize
//   * NO_RESOURCE - the response contains no data, likely because
//     no ids were specified in the request
type ClientBatchStatusResponse struct {
	Status        ClientBatchStatusResponse_Status `protobuf:"varint,1,opt,name=status,enum=ClientBatchStatusResponse_Status" json:"status,omitempty"`
	BatchStatuses []*ClientBatchStatus             `protobuf:"bytes,2,rep,name=batch_statuses,json=batchStatuses" json:"batch_statuses,omitempty"`
}

func (m *ClientBatchStatusResponse) Reset()                    { *m = ClientBatchStatusResponse{} }
func (m *ClientBatchStatusResponse) String() string            { return proto.CompactTextString(m) }
func (*ClientBatchStatusResponse) ProtoMessage()               {}
func (*ClientBatchStatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ClientBatchStatusResponse) GetStatus() ClientBatchStatusResponse_Status {
	if m != nil {
		return m.Status
	}
	return ClientBatchStatusResponse_STATUS_UNSET
}

func (m *ClientBatchStatusResponse) GetBatchStatuses() []*ClientBatchStatus {
	if m != nil {
		return m.BatchStatuses
	}
	return nil
}

func init() {
	proto.RegisterType((*ClientBatchStatus)(nil), "ClientBatchStatus")
	proto.RegisterType((*ClientBatchStatus_InvalidTransaction)(nil), "ClientBatchStatus.InvalidTransaction")
	proto.RegisterType((*ClientBatchSubmitRequest)(nil), "ClientBatchSubmitRequest")
	proto.RegisterType((*ClientBatchSubmitResponse)(nil), "ClientBatchSubmitResponse")
	proto.RegisterType((*ClientBatchStatusRequest)(nil), "ClientBatchStatusRequest")
	proto.RegisterType((*ClientBatchStatusResponse)(nil), "ClientBatchStatusResponse")
	proto.RegisterEnum("ClientBatchStatus_Status", ClientBatchStatus_Status_name, ClientBatchStatus_Status_value)
	proto.RegisterEnum("ClientBatchSubmitResponse_Status", ClientBatchSubmitResponse_Status_name, ClientBatchSubmitResponse_Status_value)
	proto.RegisterEnum("ClientBatchStatusResponse_Status", ClientBatchStatusResponse_Status_name, ClientBatchStatusResponse_Status_value)
}

func init() {
	proto.RegisterFile("github.com/grkvlt/sawtooth-go-sdk/protobuf/client_batch_submit_pb2/client_batch_submit.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 560 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0xdb, 0x6e, 0xd3, 0x40,
	0x10, 0xc5, 0x49, 0xc9, 0x65, 0x12, 0x07, 0x77, 0x01, 0x91, 0x94, 0x97, 0x60, 0x54, 0x29, 0x12,
	0x92, 0x2b, 0xc2, 0x53, 0x25, 0x5e, 0x72, 0x31, 0xc5, 0x6a, 0x6a, 0x87, 0x8d, 0x4d, 0x81, 0x17,
	0x6b, 0x1d, 0x2f, 0x60, 0xd1, 0xc4, 0xa1, 0xbb, 0xa1, 0x88, 0x3f, 0xe0, 0x73, 0xf8, 0x2d, 0xbe,
	0x82, 0xf5, 0xda, 0x26, 0x69, 0x92, 0x22, 0xd4, 0xa7, 0x78, 0xc6, 0x67, 0x66, 0x72, 0xce, 0x19,
	0x0f, 0x9c, 0x30, 0x72, 0xc5, 0xe3, 0x98, 0x7f, 0xf6, 0x59, 0xf8, 0xe5, 0x68, 0x71, 0x19, 0xf3,
	0x38, 0x58, 0x7e, 0x3c, 0x9a, 0x5e, 0x44, 0x74, 0xce, 0xfd, 0x80, 0xf0, 0xa9, 0x78, 0xb3, 0x0c,
	0x66, 0x11, 0xf7, 0x17, 0x41, 0x77, 0x57, 0xde, 0x90, 0x45, 0x07, 0xcf, 0x76, 0x37, 0x4a, 0x91,
	0x49, 0xa9, 0x7c, 0x4a, 0xc1, 0xfa, 0xcf, 0x22, 0xec, 0x0f, 0x64, 0xab, 0x7e, 0x92, 0x9d, 0x70,
	0xc2, 0x97, 0x0c, 0xb5, 0xa0, 0x92, 0xc2, 0xa3, 0xb0, 0xa9, 0xb4, 0x95, 0x4e, 0x15, 0x97, 0x65,
	0x6c, 0x85, 0xe8, 0x39, 0x94, 0x98, 0x04, 0x35, 0x0b, 0xe2, 0x45, 0xa3, 0xdb, 0x32, 0xb6, 0xca,
	0x8d, 0xf4, 0x07, 0x67, 0x40, 0xf4, 0x0e, 0x1e, 0x44, 0xf3, 0x6f, 0xe4, 0x22, 0x0a, 0x7d, 0x7e,
	0x49, 0xe6, 0x8c, 0x4c, 0x79, 0x14, 0xcf, 0x59, 0xb3, 0xd8, 0x2e, 0x76, 0x6a, 0xdd, 0xc3, 0x1d,
	0x0d, 0xac, 0x14, 0xee, 0xae, 0xd0, 0xf8, 0x7e, 0xb4, 0x95, 0x63, 0x07, 0x3f, 0x00, 0x6d, 0x43,
	0xd1, 0x21, 0x34, 0xd6, 0xe6, 0xac, 0x38, 0xa8, 0x6b, 0x59, 0xc1, 0xa4, 0x09, 0xe5, 0x19, 0x65,
	0x8c, 0x7c, 0xa2, 0x92, 0x8a, 0xe0, 0x98, 0x85, 0xe8, 0x29, 0xa8, 0xf4, 0x3b, 0xa7, 0xf3, 0x90,
	0x86, 0x7e, 0x48, 0x38, 0x11, 0xff, 0x54, 0xe9, 0xd4, 0x71, 0x3d, 0x4f, 0x0e, 0x45, 0x4e, 0x1f,
	0x43, 0x29, 0x53, 0x4b, 0x83, 0xfa, 0xc4, 0xed, 0xb9, 0xde, 0xc4, 0xf7, 0xec, 0x89, 0xe9, 0x6a,
	0x77, 0x90, 0x0a, 0xd5, 0x81, 0x73, 0x76, 0x66, 0xb9, 0xae, 0x39, 0xd4, 0x14, 0x54, 0x83, 0xb2,
	0x65, 0xbf, 0xed, 0x8d, 0xac, 0xa1, 0x56, 0x48, 0x82, 0xb1, 0x69, 0x0f, 0x2d, 0xfb, 0x44, 0x2b,
	0x26, 0x81, 0x67, 0x9f, 0xda, 0xce, 0xb9, 0xad, 0xed, 0xe9, 0x2f, 0xa1, 0xb9, 0x2e, 0x85, 0xf4,
	0x14, 0xd3, 0xaf, 0x4b, 0xca, 0x38, 0x6a, 0x43, 0xea, 0x00, 0x65, 0x82, 0x4c, 0x22, 0x5b, 0xc9,
	0x90, 0x28, 0x9c, 0xa7, 0xf5, 0x5f, 0x0a, 0xb4, 0x76, 0x94, 0xb3, 0x85, 0x10, 0x8a, 0xa2, 0xe3,
	0xbf, 0xb6, 0x29, 0xd2, 0xb6, 0x27, 0xc6, 0x8d, 0xd8, 0x0d, 0xfb, 0xf4, 0xf7, 0xff, 0x20, 0x5a,
	0x82, 0x82, 0x73, 0x2a, 0x18, 0x22, 0x68, 0x58, 0xb6, 0x6b, 0x62, 0xbb, 0x37, 0xf2, 0x4d, 0x8c,
	0x1d, 0x2c, 0x88, 0xee, 0x83, 0x9a, 0xb1, 0xf6, 0xfb, 0x3d, 0x77, 0xf0, 0x5a, 0xd0, 0x6d, 0x00,
	0xbc, 0xf1, 0x4c, 0xcf, 0xf4, 0x5f, 0x79, 0xa3, 0x91, 0x60, 0x4c, 0xaf, 0x33, 0x4e, 0xe7, 0x66,
	0x8c, 0x1f, 0x43, 0x35, 0xdf, 0xc1, 0x94, 0x73, 0x15, 0x57, 0xb2, 0x25, 0x64, 0x62, 0xde, 0xde,
	0x15, 0x89, 0xb8, 0x34, 0xae, 0x82, 0xe5, 0x73, 0xe2, 0x27, 0x8f, 0x66, 0x34, 0x5e, 0x72, 0xe9,
	0x97, 0x8a, 0xf3, 0x50, 0xff, 0xbd, 0x21, 0x4d, 0x36, 0xe7, 0xbf, 0xa4, 0xb9, 0x86, 0xdd, 0xdc,
	0xec, 0x63, 0x68, 0x64, 0x1f, 0xa0, 0x8c, 0x69, 0xf2, 0x51, 0x24, 0xe6, 0xa0, 0x1d, 0x2d, 0xd4,
	0x60, 0x15, 0x08, 0xbb, 0xce, 0x6f, 0xa9, 0xea, 0x3d, 0xa8, 0xd9, 0x8e, 0x8f, 0xcd, 0x89, 0xe3,
	0xe1, 0x81, 0xa9, 0xdd, 0x4d, 0x34, 0xcd, 0x65, 0x16, 0xfb, 0x55, 0xe9, 0x77, 0xe1, 0x61, 0x7e,
	0x00, 0x0c, 0x71, 0x00, 0x8c, 0xfc, 0x00, 0x8c, 0x95, 0x0f, 0x8f, 0x6e, 0x38, 0x26, 0x41, 0x49,
	0x82, 0x5e, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0xa7, 0xe5, 0x26, 0x9e, 0x84, 0x04, 0x00, 0x00,
}
