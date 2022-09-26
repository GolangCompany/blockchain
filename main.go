package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type BlockChain struct {
	blocks []*Cryptoblock
}

type Cryptoblock struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

func (c *Cryptoblock) BuildHash() {
	details := bytes.Join([][]byte{c.Data, c.PrevHash}, []byte{})
	hash := sha256.Sum256(details)
	c.Hash = hash[:]
}

func BuildBlock(data string, prevHash []byte) *Cryptoblock {
	block := &Cryptoblock{[]byte{}, []byte(data), prevHash}
	block.BuildHash()
	return block
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := BuildBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

func Inception() *Cryptoblock {
	return BuildBlock("Inception", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Cryptoblock{Inception()}}
}

func main() {
	chain := InitBlockChain()

	chain.AddBlock("First Block after Inception")
	chain.AddBlock("Second Block after Inception")
	chain.AddBlock("Third Block after Inception")

	for _, block := range chain.blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}
}
