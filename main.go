//Ahmad Naeem
//20I-1810
//Section A
//Assignment 1

package main

import (
	"crypto/sha256" //Importing the SHA-256 hashing algorithm package
	"encoding/hex"  //Importing the package for hexadecimal encoding
	"fmt"           //Importing the package for formatted I/O operations
)

type Block struct {
	Transaction  string //Transaction data stored in the block
	Nonce        int    //Nonce value used in mining
	PreviousHash string //Hash of the previous block in the blockchain
	Hash         string //Hash of the current block
}

var Blockchain []Block //Blockchain variable to store the list of blocks

func NewBlock(transaction string, nonce int, previousHash string) *Block {
	block := &Block{
		Transaction:  transaction,  //Set the transaction data
		Nonce:        nonce,        //Set the nonce value
		PreviousHash: previousHash, //Set the previous block's hash
	}
	block.Hash = block.CreateHash()         //Calculate and set the current block's hash
	Blockchain = append(Blockchain, *block) //Append the block to the blockchain
	return block                            //Return the newly created block
}

func (b *Block) CreateHash() string {
	data := fmt.Sprintf("%s%d%s", b.Transaction, b.Nonce, b.PreviousHash) //Concatenate transaction, nonce, and previous hash
	hash := sha256.New()                                                  //Create a new SHA-256 hash instance
	hash.Write([]byte(data))                                              //Write the concatenated data to the hash
	return hex.EncodeToString(hash.Sum(nil))                              //Return the hexadecimal representation of the hash
}

func DisplayBlocks() {
	for _, block := range Blockchain { //For loop to iterate through each block in the blockchain
		fmt.Printf("Transaction: %s\nNonce: %d\nPrevious Hash: %s\nCurrent Hash: %s\n\n", block.Transaction, block.Nonce, block.PreviousHash, block.Hash) //Print block details
	}
}

func ChangeBlock(blockIndex int, newTransaction string) {
	if blockIndex >= 0 && blockIndex < len(Blockchain) { //Check if the index is valid
		Blockchain[blockIndex].Transaction = newTransaction               //Update the transaction of the block
		Blockchain[blockIndex].Hash = Blockchain[blockIndex].CreateHash() //Recalculate the block's hash
	}
}

func VerifyChain() bool {
	for i := 1; i < len(Blockchain); i++ { //For loop to iterate through the blockchain starting from the second block
		currentBlock := Blockchain[i]    //Get the current block
		previousBlock := Blockchain[i-1] //Get the previous block

		if currentBlock.Hash != currentBlock.CreateHash() { //Check if the stored hash matches the calculated hash
			return false //If not, the blockchain is invalid
		}

		if currentBlock.PreviousHash != previousBlock.Hash { //Check if the previous hash matches the hash of the previous block
			return false //If not, the blockchain is invalid
		}
	}
	return true //If all blocks are valid, the blockchain is verified
}

func main() {
	// Creating new blocks
	NewBlock("ahmad to ali", 123, "0")
	NewBlock("ali to babar", 456, Blockchain[0].Hash)

	// Displaying blocks
	fmt.Println("Blockchain:")
	DisplayBlocks()

	// Changing the transaction of a specific block
	ChangeBlock(1, "babar to ahmad")

	// Displaying blocks after changing a block's transaction
	fmt.Println("Blockchain after changing a block:")
	DisplayBlocks()

	// Verifying the integrity of the blockchain
	isValid := VerifyChain()
	fmt.Println("Blockchain Validity:", isValid)
}
