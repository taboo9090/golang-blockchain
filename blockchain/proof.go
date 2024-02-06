package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

// Take the data from the block.

// Create the counter (nonce) which starts at 0.

// Create a hash of the data + the counter.

// Check the hash to see if it meets a set of requirements.

// Requirements:
// The First few bytes must contain 0s.

const Difficulty = 17

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

// NewProof creates a new instance of the ProofOfWork struct
func NewProof(b *Block) *ProofOfWork {
	// target is initialized to 1 and then shifted left by 256 - Difficulty
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))

	pow := &ProofOfWork{b, target}

	return pow
}

// InitData creates the initial data for the proof of work algorithm
func (pow *ProofOfWork) InitData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevHash,
			pow.Block.Data,
			ToHex(int64(nonce)),
			ToHex(int64(Difficulty)),
		},
		[]byte{})

	return data
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var intHash big.Int
	var hash [32]byte

	nonce := 0

	for nonce < math.MaxInt64 {
		data := pow.InitData(nonce)
		hash = sha256.Sum256(data)

		intHash.SetBytes(hash[:])
		fmt.Printf("\r%x", hash)

		if intHash.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}
	}

	fmt.Println()

	return nonce, hash[:]
}

// Validate checks if the proof of work is valid
func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int

	data := pow.InitData(pow.Block.Nonce)

	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])

	return intHash.Cmp(pow.Target) == -1
}

// ToHex converts an integer to a slice of bytes in big endian format.
func ToHex(num int64) []byte {
	// create a new buffer to write the bytes to
	buff := new(bytes.Buffer)

	// write the integer to the buffer in big endian format
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}
