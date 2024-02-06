package blockchain

type BlockChain struct {
	Blocks []*Block
}

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

// CreateBlock creates a new block with the given data and previous block hash
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{
		[]byte{}, // set the block's hash to an empty byte slice
		[]byte(data),
		prevHash,
		0,
	}

	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// AddBlock adds a new block to the blockchain with the given data
func (chain *BlockChain) AddBlock(data string) {
	PrevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, PrevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

// Genesis creates the first block in the blockchain
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// InitBlockChain creates a new instance of BlockChain with the first block as the genesis block
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
