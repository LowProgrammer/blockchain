package main

type BlockChain struct{
	blocks []*Block
}

func NewBlockChain() *BlockChain{
	block := NewGenesisBlock()
	return BlockChain{blocks:[]*Block{block}}
}

func (bc *BlockChain)AddBlock(data string){
	preBlockHash := bc.blocks[len(bc.block)-1].Hash
	block :=NewBlock(data,preBlockHash)
	bc.blocks=append(bc.blocks,blocks)
}