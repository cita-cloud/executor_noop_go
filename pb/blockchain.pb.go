// Copyright Rivtower Technologies LLC.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.21.12
// source: blockchain.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BlockHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// hash of previous BlockHeader
	Prevhash         []byte `protobuf:"bytes,1,opt,name=prevhash,proto3" json:"prevhash,omitempty"`
	Timestamp        uint64 `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Height           uint64 `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	TransactionsRoot []byte `protobuf:"bytes,4,opt,name=transactions_root,json=transactionsRoot,proto3" json:"transactions_root,omitempty"`
	Proposer         []byte `protobuf:"bytes,5,opt,name=proposer,proto3" json:"proposer,omitempty"`
}

func (x *BlockHeader) Reset() {
	*x = BlockHeader{}
	mi := &file_blockchain_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BlockHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockHeader) ProtoMessage() {}

func (x *BlockHeader) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockHeader.ProtoReflect.Descriptor instead.
func (*BlockHeader) Descriptor() ([]byte, []int) {
	return file_blockchain_proto_rawDescGZIP(), []int{0}
}

func (x *BlockHeader) GetPrevhash() []byte {
	if x != nil {
		return x.Prevhash
	}
	return nil
}

func (x *BlockHeader) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *BlockHeader) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *BlockHeader) GetTransactionsRoot() []byte {
	if x != nil {
		return x.TransactionsRoot
	}
	return nil
}

func (x *BlockHeader) GetProposer() []byte {
	if x != nil {
		return x.Proposer
	}
	return nil
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version uint32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	// 1. length is 20 bytes for evm.
	// 2. if executor is multi-vm, it will be a path.
	To []byte `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	// length is less than 128
	Nonce           string `protobuf:"bytes,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Quota           uint64 `protobuf:"varint,4,opt,name=quota,proto3" json:"quota,omitempty"`
	ValidUntilBlock uint64 `protobuf:"varint,5,opt,name=valid_until_block,json=validUntilBlock,proto3" json:"valid_until_block,omitempty"`
	Data            []byte `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
	// length is 32 bytes.
	Value []byte `protobuf:"bytes,7,opt,name=value,proto3" json:"value,omitempty"`
	// length is 32 bytes.
	ChainId []byte `protobuf:"bytes,8,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	mi := &file_blockchain_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_blockchain_proto_rawDescGZIP(), []int{1}
}

func (x *Transaction) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Transaction) GetTo() []byte {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *Transaction) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

func (x *Transaction) GetQuota() uint64 {
	if x != nil {
		return x.Quota
	}
	return 0
}

func (x *Transaction) GetValidUntilBlock() uint64 {
	if x != nil {
		return x.ValidUntilBlock
	}
	return 0
}

func (x *Transaction) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Transaction) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *Transaction) GetChainId() []byte {
	if x != nil {
		return x.ChainId
	}
	return nil
}

type Witness struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signature []byte `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	// add to support multi-address, or we don't know which address algorithm to use
	Sender []byte `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
}

func (x *Witness) Reset() {
	*x = Witness{}
	mi := &file_blockchain_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Witness) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Witness) ProtoMessage() {}

func (x *Witness) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Witness.ProtoReflect.Descriptor instead.
func (*Witness) Descriptor() ([]byte, []int) {
	return file_blockchain_proto_rawDescGZIP(), []int{2}
}

func (x *Witness) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *Witness) GetSender() []byte {
	if x != nil {
		return x.Sender
	}
	return nil
}

type UnverifiedTransaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *Transaction `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	// add to support multi-hash, or we don't know which hash algorithm to use
	TransactionHash []byte   `protobuf:"bytes,2,opt,name=transaction_hash,json=transactionHash,proto3" json:"transaction_hash,omitempty"`
	Witness         *Witness `protobuf:"bytes,3,opt,name=witness,proto3" json:"witness,omitempty"`
}

func (x *UnverifiedTransaction) Reset() {
	*x = UnverifiedTransaction{}
	mi := &file_blockchain_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UnverifiedTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnverifiedTransaction) ProtoMessage() {}

func (x *UnverifiedTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnverifiedTransaction.ProtoReflect.Descriptor instead.
func (*UnverifiedTransaction) Descriptor() ([]byte, []int) {
	return file_blockchain_proto_rawDescGZIP(), []int{3}
}

func (x *UnverifiedTransaction) GetTransaction() *Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *UnverifiedTransaction) GetTransactionHash() []byte {
	if x != nil {
		return x.TransactionHash
	}
	return nil
}

func (x *UnverifiedTransaction) GetWitness() *Witness {
	if x != nil {
		return x.Witness
	}
	return nil
}

type UtxoTransaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version   uint32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	PreTxHash []byte `protobuf:"bytes,2,opt,name=pre_tx_hash,json=preTxHash,proto3" json:"pre_tx_hash,omitempty"`
	Output    []byte `protobuf:"bytes,3,opt,name=output,proto3" json:"output,omitempty"`
	LockId    uint64 `protobuf:"varint,4,opt,name=lock_id,json=lockId,proto3" json:"lock_id,omitempty"`
}

func (x *UtxoTransaction) Reset() {
	*x = UtxoTransaction{}
	mi := &file_blockchain_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UtxoTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UtxoTransaction) ProtoMessage() {}

func (x *UtxoTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UtxoTransaction.ProtoReflect.Descriptor instead.
func (*UtxoTransaction) Descriptor() ([]byte, []int) {
	return file_blockchain_proto_rawDescGZIP(), []int{4}
}

func (x *UtxoTransaction) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *UtxoTransaction) GetPreTxHash() []byte {
	if x != nil {
		return x.PreTxHash
	}
	return nil
}

func (x *UtxoTransaction) GetOutput() []byte {
	if x != nil {
		return x.Output
	}
	return nil
}

func (x *UtxoTransaction) GetLockId() uint64 {
	if x != nil {
		return x.LockId
	}
	return 0
}

type UnverifiedUtxoTransaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *UtxoTransaction `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	// add to support multi-hash, or we don't know which hash algorithm to use
	TransactionHash []byte     `protobuf:"bytes,2,opt,name=transaction_hash,json=transactionHash,proto3" json:"transaction_hash,omitempty"`
	Witnesses       []*Witness `protobuf:"bytes,3,rep,name=witnesses,proto3" json:"witnesses,omitempty"`
}

func (x *UnverifiedUtxoTransaction) Reset() {
	*x = UnverifiedUtxoTransaction{}
	mi := &file_blockchain_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UnverifiedUtxoTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnverifiedUtxoTransaction) ProtoMessage() {}

func (x *UnverifiedUtxoTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnverifiedUtxoTransaction.ProtoReflect.Descriptor instead.
func (*UnverifiedUtxoTransaction) Descriptor() ([]byte, []int) {
	return file_blockchain_proto_rawDescGZIP(), []int{5}
}

func (x *UnverifiedUtxoTransaction) GetTransaction() *UtxoTransaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *UnverifiedUtxoTransaction) GetTransactionHash() []byte {
	if x != nil {
		return x.TransactionHash
	}
	return nil
}

func (x *UnverifiedUtxoTransaction) GetWitnesses() []*Witness {
	if x != nil {
		return x.Witnesses
	}
	return nil
}

type RawTransactions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body []*RawTransaction `protobuf:"bytes,1,rep,name=body,proto3" json:"body,omitempty"`
}

func (x *RawTransactions) Reset() {
	*x = RawTransactions{}
	mi := &file_blockchain_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RawTransactions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RawTransactions) ProtoMessage() {}

func (x *RawTransactions) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RawTransactions.ProtoReflect.Descriptor instead.
func (*RawTransactions) Descriptor() ([]byte, []int) {
	return file_blockchain_proto_rawDescGZIP(), []int{6}
}

func (x *RawTransactions) GetBody() []*RawTransaction {
	if x != nil {
		return x.Body
	}
	return nil
}

type RawTransaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Tx:
	//
	//	*RawTransaction_NormalTx
	//	*RawTransaction_UtxoTx
	Tx isRawTransaction_Tx `protobuf_oneof:"tx"`
}

func (x *RawTransaction) Reset() {
	*x = RawTransaction{}
	mi := &file_blockchain_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RawTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RawTransaction) ProtoMessage() {}

func (x *RawTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RawTransaction.ProtoReflect.Descriptor instead.
func (*RawTransaction) Descriptor() ([]byte, []int) {
	return file_blockchain_proto_rawDescGZIP(), []int{7}
}

func (m *RawTransaction) GetTx() isRawTransaction_Tx {
	if m != nil {
		return m.Tx
	}
	return nil
}

func (x *RawTransaction) GetNormalTx() *UnverifiedTransaction {
	if x, ok := x.GetTx().(*RawTransaction_NormalTx); ok {
		return x.NormalTx
	}
	return nil
}

func (x *RawTransaction) GetUtxoTx() *UnverifiedUtxoTransaction {
	if x, ok := x.GetTx().(*RawTransaction_UtxoTx); ok {
		return x.UtxoTx
	}
	return nil
}

type isRawTransaction_Tx interface {
	isRawTransaction_Tx()
}

type RawTransaction_NormalTx struct {
	NormalTx *UnverifiedTransaction `protobuf:"bytes,1,opt,name=normal_tx,json=normalTx,proto3,oneof"`
}

type RawTransaction_UtxoTx struct {
	UtxoTx *UnverifiedUtxoTransaction `protobuf:"bytes,2,opt,name=utxo_tx,json=utxoTx,proto3,oneof"`
}

func (*RawTransaction_NormalTx) isRawTransaction_Tx() {}

func (*RawTransaction_UtxoTx) isRawTransaction_Tx() {}

type CompactBlockBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// transaction hash of UnverifiedTransaction or UnverifyedUtxoTransaction.
	TxHashes [][]byte `protobuf:"bytes,1,rep,name=tx_hashes,json=txHashes,proto3" json:"tx_hashes,omitempty"`
}

func (x *CompactBlockBody) Reset() {
	*x = CompactBlockBody{}
	mi := &file_blockchain_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CompactBlockBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompactBlockBody) ProtoMessage() {}

func (x *CompactBlockBody) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompactBlockBody.ProtoReflect.Descriptor instead.
func (*CompactBlockBody) Descriptor() ([]byte, []int) {
	return file_blockchain_proto_rawDescGZIP(), []int{8}
}

func (x *CompactBlockBody) GetTxHashes() [][]byte {
	if x != nil {
		return x.TxHashes
	}
	return nil
}

type CompactBlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version uint32            `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Header  *BlockHeader      `protobuf:"bytes,2,opt,name=header,proto3" json:"header,omitempty"`
	Body    *CompactBlockBody `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *CompactBlock) Reset() {
	*x = CompactBlock{}
	mi := &file_blockchain_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CompactBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompactBlock) ProtoMessage() {}

func (x *CompactBlock) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompactBlock.ProtoReflect.Descriptor instead.
func (*CompactBlock) Descriptor() ([]byte, []int) {
	return file_blockchain_proto_rawDescGZIP(), []int{9}
}

func (x *CompactBlock) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *CompactBlock) GetHeader() *BlockHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *CompactBlock) GetBody() *CompactBlockBody {
	if x != nil {
		return x.Body
	}
	return nil
}

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version   uint32           `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Header    *BlockHeader     `protobuf:"bytes,2,opt,name=header,proto3" json:"header,omitempty"`
	Body      *RawTransactions `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	Proof     []byte           `protobuf:"bytes,4,opt,name=proof,proto3" json:"proof,omitempty"`
	StateRoot []byte           `protobuf:"bytes,5,opt,name=state_root,json=stateRoot,proto3" json:"state_root,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	mi := &file_blockchain_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_blockchain_proto_rawDescGZIP(), []int{10}
}

func (x *Block) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Block) GetHeader() *BlockHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Block) GetBody() *RawTransactions {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *Block) GetProof() []byte {
	if x != nil {
		return x.Proof
	}
	return nil
}

func (x *Block) GetStateRoot() []byte {
	if x != nil {
		return x.StateRoot
	}
	return nil
}

var File_blockchain_proto protoreflect.FileDescriptor

var file_blockchain_proto_rawDesc = []byte{
	0x0a, 0x10, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x22, 0xa8,
	0x01, 0x0a, 0x0b, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x72, 0x65, 0x76, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x08, 0x70, 0x72, 0x65, 0x76, 0x68, 0x61, 0x73, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x12, 0x2b, 0x0a, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x10, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x22, 0xd4, 0x01, 0x0a, 0x0b, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x02, 0x74, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x6f,
	0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x12,
	0x2a, 0x0a, 0x11, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x75, 0x6e, 0x74, 0x69, 0x6c, 0x5f, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x55, 0x6e, 0x74, 0x69, 0x6c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64,
	0x22, 0x3f, 0x0a, 0x07, 0x57, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x22, 0xac, 0x01, 0x0a, 0x15, 0x55, 0x6e, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0b, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x73,
	0x68, 0x12, 0x2d, 0x0a, 0x07, 0x77, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e,
	0x57, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x07, 0x77, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73,
	0x22, 0x7c, 0x0a, 0x0f, 0x55, 0x74, 0x78, 0x6f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a,
	0x0b, 0x70, 0x72, 0x65, 0x5f, 0x74, 0x78, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x09, 0x70, 0x72, 0x65, 0x54, 0x78, 0x48, 0x61, 0x73, 0x68, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x6f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x22, 0xb8,
	0x01, 0x0a, 0x19, 0x55, 0x6e, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x55, 0x74, 0x78,
	0x6f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3d, 0x0a, 0x0b,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x55,
	0x74, 0x78, 0x6f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x12, 0x31, 0x0a, 0x09, 0x77, 0x69, 0x74, 0x6e, 0x65, 0x73,
	0x73, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x57, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x09,
	0x77, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x65, 0x73, 0x22, 0x41, 0x0a, 0x0f, 0x52, 0x61, 0x77,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2e, 0x0a, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x61, 0x77, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x9a, 0x01, 0x0a,
	0x0e, 0x52, 0x61, 0x77, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x40, 0x0a, 0x09, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x5f, 0x74, 0x78, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e,
	0x55, 0x6e, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x08, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x54,
	0x78, 0x12, 0x40, 0x0a, 0x07, 0x75, 0x74, 0x78, 0x6f, 0x5f, 0x74, 0x78, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e,
	0x55, 0x6e, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x55, 0x74, 0x78, 0x6f, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x06, 0x75, 0x74, 0x78,
	0x6f, 0x54, 0x78, 0x42, 0x04, 0x0a, 0x02, 0x74, 0x78, 0x22, 0x2f, 0x0a, 0x10, 0x43, 0x6f, 0x6d,
	0x70, 0x61, 0x63, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x1b, 0x0a,
	0x09, 0x74, 0x78, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c,
	0x52, 0x08, 0x74, 0x78, 0x48, 0x61, 0x73, 0x68, 0x65, 0x73, 0x22, 0x8b, 0x01, 0x0a, 0x0c, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x42, 0x6f,
	0x64, 0x79, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0xb8, 0x01, 0x0a, 0x05, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x06,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x2f, 0x0a,
	0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x61, 0x77, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x70,
	0x72, 0x6f, 0x6f, 0x66, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x6f,
	0x6f, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x6f, 0x6f, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blockchain_proto_rawDescOnce sync.Once
	file_blockchain_proto_rawDescData = file_blockchain_proto_rawDesc
)

func file_blockchain_proto_rawDescGZIP() []byte {
	file_blockchain_proto_rawDescOnce.Do(func() {
		file_blockchain_proto_rawDescData = protoimpl.X.CompressGZIP(file_blockchain_proto_rawDescData)
	})
	return file_blockchain_proto_rawDescData
}

var file_blockchain_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_blockchain_proto_goTypes = []any{
	(*BlockHeader)(nil),               // 0: blockchain.BlockHeader
	(*Transaction)(nil),               // 1: blockchain.Transaction
	(*Witness)(nil),                   // 2: blockchain.Witness
	(*UnverifiedTransaction)(nil),     // 3: blockchain.UnverifiedTransaction
	(*UtxoTransaction)(nil),           // 4: blockchain.UtxoTransaction
	(*UnverifiedUtxoTransaction)(nil), // 5: blockchain.UnverifiedUtxoTransaction
	(*RawTransactions)(nil),           // 6: blockchain.RawTransactions
	(*RawTransaction)(nil),            // 7: blockchain.RawTransaction
	(*CompactBlockBody)(nil),          // 8: blockchain.CompactBlockBody
	(*CompactBlock)(nil),              // 9: blockchain.CompactBlock
	(*Block)(nil),                     // 10: blockchain.Block
}
var file_blockchain_proto_depIdxs = []int32{
	1,  // 0: blockchain.UnverifiedTransaction.transaction:type_name -> blockchain.Transaction
	2,  // 1: blockchain.UnverifiedTransaction.witness:type_name -> blockchain.Witness
	4,  // 2: blockchain.UnverifiedUtxoTransaction.transaction:type_name -> blockchain.UtxoTransaction
	2,  // 3: blockchain.UnverifiedUtxoTransaction.witnesses:type_name -> blockchain.Witness
	7,  // 4: blockchain.RawTransactions.body:type_name -> blockchain.RawTransaction
	3,  // 5: blockchain.RawTransaction.normal_tx:type_name -> blockchain.UnverifiedTransaction
	5,  // 6: blockchain.RawTransaction.utxo_tx:type_name -> blockchain.UnverifiedUtxoTransaction
	0,  // 7: blockchain.CompactBlock.header:type_name -> blockchain.BlockHeader
	8,  // 8: blockchain.CompactBlock.body:type_name -> blockchain.CompactBlockBody
	0,  // 9: blockchain.Block.header:type_name -> blockchain.BlockHeader
	6,  // 10: blockchain.Block.body:type_name -> blockchain.RawTransactions
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_blockchain_proto_init() }
func file_blockchain_proto_init() {
	if File_blockchain_proto != nil {
		return
	}
	file_blockchain_proto_msgTypes[7].OneofWrappers = []any{
		(*RawTransaction_NormalTx)(nil),
		(*RawTransaction_UtxoTx)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_blockchain_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_blockchain_proto_goTypes,
		DependencyIndexes: file_blockchain_proto_depIdxs,
		MessageInfos:      file_blockchain_proto_msgTypes,
	}.Build()
	File_blockchain_proto = out.File
	file_blockchain_proto_rawDesc = nil
	file_blockchain_proto_goTypes = nil
	file_blockchain_proto_depIdxs = nil
}
