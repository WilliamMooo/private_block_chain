package main

import (
	"core"
	"fmt"
)

func main() {
	bc := core.NewBlockchain() // 初始化区块链

	bc.AddBlock("这是第一个区块")
	bc.AddBlock("这是第二个区块")

	for _, block := range bc.Blocks {
		fmt.Printf("上一个区块的hash: %x\n", block.PreBlockHash)
		fmt.Printf("区块上的信息: %s\n", block.Data)
		fmt.Printf("本区块的hash: %x\n", block.Hash)
		fmt.Println()
	}
}
