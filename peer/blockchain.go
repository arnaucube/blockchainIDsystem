package main

import (
	"errors"
	"time"
)

type Block struct {
	Hash         string    `json:"hash"`
	Height       int64     `json:"height"`
	Date         time.Time `json:"date"`
	PreviousHash string    `json:"previoushash"`
	NextHash     string    `json:"nexthash"`
	Data         []Address `json:"data"`
	Emitter      string    `json:"emitter"` //the ID of the peer that has emmited the block
}

type Blockchain struct {
	GenesisBlock string    `json:"genesisblock"`
	LastUpdate   time.Time `json:"lastupdate"`
	Blocks       []Block   `json:"blocks"`
}

var blockchain Blockchain

func (bc *Blockchain) getBlockByHash(hash string) (Block, error) {
	for _, block := range bc.Blocks {
		if block.Hash == hash {
			return block, nil
		}
	}
	var b Block
	return b, errors.New("Block Hash not found")
}

func (bc *Blockchain) createBlock(address Address) Block {
	var b Block
	b.Height = int64(len(bc.Blocks))
	if len(bc.Blocks) == 0 {
		b.Height = 0
	} else {
		b.PreviousHash = bc.Blocks[len(bc.Blocks)-1].Hash
	}
	b.Date = time.Now()
	b.Data = append(b.Data, address)
	b.Emitter = runningPeer.ID

	b.Hash = hashBlock(b)
	return b
}

func (bc *Blockchain) blockExists(block Block) bool {
	for _, b := range bc.Blocks {
		if b.Hash == block.Hash {
			return true
		}
	}
	return false
}
func (bc *Blockchain) addBlock(block Block) error {
	if len(bc.Blocks) > 0 {
		bc.Blocks[len(bc.Blocks)-1].NextHash = block.Hash
	} else {
		bc.GenesisBlock = block.Hash
	}
	bc.Blocks = append(bc.Blocks, block)

	return nil
}
