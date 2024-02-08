package blockchain

type BlockChain struct {
	Blocks []*Block
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func (chain *BlockChain) AddBlock(data string) {
	PrevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, PrevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}
