package main

import(
	"fmt"
)
type block struct{
	//版本
	Version int32
	//前区块的哈希值
	PrevBlockHash []byte
	//当前区块的哈希值
	Hash []byte
	//梅克尔根
	MerKerRoot []byte
	//时间戳
	TimeStamp int64
	//难度值
	Bits int64
	//随机值
	Nonce int64

	//交易信息
	Data []byte
}

func NewBlock(data string.prevBlockHash []byte) *Block{
	//
	var block Block
	block = Block{
		Version:1,
		PrevBlockHash:prevBlockHash,
		//Hash
		MerKerRoot:[]byte,
		TimeStamp:time.Now().Unix(),
		Bits:1,
		Nonce:1,
		Data:[]byte(data)}
	block.SetHash()
	return &block
}

func (block* Block)SetHash(){

	tmp := [][]byte{
		IntToByte(block.Version),
		block.PrevBlockHash,
		block.MerKerRoot,
		IntToByte(block.TimeStamp),
		IntToByte(block.Bits),
		IntToByte(block.Nonce),
		block.Data}
	byte.join(tmp,[]byte{})
	hash := sha256.Sum256(data)
	block.Hash = hash[:]
}


func NewGenesisBlock() {
	NewBlock("Genesis Block",[]byte{})
}