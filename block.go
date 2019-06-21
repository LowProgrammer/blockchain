package main

type block struct{
	//版本
	Version int32
	//前区块的哈希值
	PrevBlockHash []byte
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