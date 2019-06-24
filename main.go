package main

import "fmt"
func main(){
	bc :=NewBlockChain()
	bc.AddBlock("A send B 1BTC")
	bc.AddBlock("B send C 1BTC")

	for _,block := range bc.block{
		fmt.Print("Vesion",block.Version)
	}
}