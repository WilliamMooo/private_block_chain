package core

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// Block 区块
type Block struct {
	Timestamp     int64  //区块时间戳
	Data          []byte //区块包含的数据
	PreBlockHash  []byte //前一个区块的hash
	Hash          []byte //区块自身的hash
}

// NewBlock 创建新区块
func NewBlock(data string, preBlockHash []byte) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		PreBlockHash: preBlockHash,
		Hash:          []byte{},
		Data:          []byte(data)}

	block.SetHash()
	return block
}

// SetHash 设置区块的hash
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PreBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

// NewGenesisBlock 创世纪区块
func NewGenesisBlock() *Block {
	return NewBlock("创世纪区块", []byte{})
}
