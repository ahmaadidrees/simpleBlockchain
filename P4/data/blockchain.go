package data

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
	"sync"
)
type Blockchain struct {
	Chain map[int32][]Block
	Length int32

}

func NewBlockchain() Blockchain {
	return Blockchain{
		Chain: make(map[int32][]Block),
		Length: 0,
	}
}



//var BlockChain map[int32][]Block
//var Length int32 = int32(len(BlockChain))
type  syncblockchain struct{
	Bc Blockchain
	Mux sync.RWMutex
}
func NewSBC() syncblockchain{
	b := NewBlockchain()
	return syncblockchain{
		Bc:  b,
		Mux: sync.RWMutex{},
	}
}
func (blockchain *Blockchain)GetLatestBlock() []Block{
	return blockchain.Chain[blockchain.Length]//BlockChain[Length]
}
func (sbc *syncblockchain)GetLatestBlock() []Block{
	sbc.Mux.Lock()
	defer sbc.Mux.Unlock()
	return sbc.Bc.GetLatestBlock()
}
func (blockchain *Blockchain)GetParentBlock(block Block) Block{
	//parentIndex := 0
	for height , Blocks := range blockchain.Chain {
		for index ,_ := range Blocks{
			if Blocks[index] == block && index == 0{
				lengthOfParentChain := len(blockchain.Chain[height-1])
				ParentChain := blockchain.Chain[height-1]
				return ParentChain[lengthOfParentChain]

			} else if Blocks[index] == block && index != 0{
				return Blocks[index - 1]
			}
		}
	}
	return Block{}
}
func (sbc *syncblockchain)GetParentBlock(block Block) Block{
	sbc.Mux.Lock()
	defer sbc.Mux.Unlock()
	return sbc.Bc.GetParentBlock(block)
}
func (sbc *syncblockchain) Get(heightKey int32) []Block {
	sbc.Mux.Lock()
	defer sbc.Mux.Unlock()
	return sbc.Bc.Get(heightKey)
}

func (blockchain *Blockchain)Get(heightKey int32) []Block{


	_, ok := blockchain.Chain[heightKey]
	if ok{
		return blockchain.Chain[heightKey]
	}

	return nil
}
func (sbc *syncblockchain)Insert(block Block){
	sbc.Mux.Lock()
	defer sbc.Mux.Unlock()
	sbc.Bc.Insert(block)
}
func (blockchain *Blockchain)Insert(block Block) {

	blockchain.Chain[block.HeaderData.Height] = append(blockchain.Chain[block.HeaderData.Height], block)

}

func EncodesToJSON(blockchain map[int32][]Block) string{
	finalString := ""
	for _,blocs := range blockchain{
		for _,bloc := range blocs{
			finalString += EncodeToJSON(bloc)
		}
	}

	return finalString
}



func (blockchain *Blockchain)DecodesFromJSON(JSONstring string){
	splitedString := strings.Split(JSONstring, "}")
	for _,value := range splitedString{

		bloc := DecodeFromJSON(value)
		blockchain.Insert(bloc)
	}
}


func (blockchain *Blockchain) Show() string {
	rs := ""
	var idList []int
	for id := range blockchain.Chain {
		idList = append(idList, int(id))
	}

	sort.Ints(idList)
	for _, id := range idList {
		var hashs []string
		for _, block := range blockchain.Chain[int32(id)] {
			hashs = append(hashs, block.HeaderData.Hash+"<="+block.HeaderData.ParentHash)
		}

		sort.Strings(hashs)
		rs += fmt.Sprintf("%v: ", id)
		for _, h := range hashs {
			rs += fmt.Sprintf("%s, ", h)
		}
		rs += "\n"
	}

	sum := sha256.Sum256([]byte(rs))
	rs = fmt.Sprintf("This is the BlockChain: %s\n", hex.EncodeToString(sum[:])) + rs

	return rs
}
