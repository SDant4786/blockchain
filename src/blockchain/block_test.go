package blockchain_test

import (
	"blockchain/src/blockchain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBlock(t *testing.T) {
	tests := []struct {
		name          string
		transactions  []*blockchain.Transaction
		prevBlockHash []byte
		height        int
	}{
		{
			"Pass",
			[]*blockchain.Transaction{
				{
					[]byte(""),
					[]blockchain.TXInput{
						{
							[]byte(""),
							-1,
							[]byte(""),
							[]byte(""),
						},
					},
					[]blockchain.TXOutput{
						{
							10,
							[]byte(""),
						},
					},
				},
			},
			[]byte(""),
			0,
		},
	}

	for _, tt := range tests {
		newBlock := blockchain.NewBlock(tt.transactions, tt.prevBlockHash, tt.height)
		assert.NotNil(t, newBlock)
	}
}
func TestNewGenesisBlock(t *testing.T) {
	tests := []struct {
		name         string
		transactions blockchain.Transaction
	}{
		{
			"Pass",
			blockchain.Transaction{
				[]byte(""),
				[]blockchain.TXInput{
					{
						[]byte(""),
						-1,
						[]byte(""),
						[]byte(""),
					},
				},
				[]blockchain.TXOutput{
					{
						10,
						[]byte(""),
					},
				},
			},
		},
	}

	for _, tt := range tests {
		newBlock := blockchain.NewGenesisBlock(&tt.transactions)
		assert.NotNil(t, newBlock)
	}
}

func TestSerialize(t *testing.T) {
	tests := []struct {
		name         string
		transactions blockchain.Transaction
	}{
		{
			"Pass",
			blockchain.Transaction{
				[]byte(""),
				[]blockchain.TXInput{
					{
						[]byte(""),
						-1,
						[]byte(""),
						[]byte(""),
					},
				},
				[]blockchain.TXOutput{
					{
						10,
						[]byte(""),
					},
				},
			},
		},
	}

	for _, tt := range tests {
		block := blockchain.NewGenesisBlock(&tt.transactions)
		output := block.Serialize()
		assert.NotNil(t, output)
	}
}

func TestHashTransactions(t *testing.T) {
	tests := []struct {
		name         string
		transactions blockchain.Transaction
	}{
		{
			"Pass",
			blockchain.Transaction{
				[]byte(""),
				[]blockchain.TXInput{
					{
						[]byte(""),
						-1,
						[]byte(""),
						[]byte(""),
					},
				},
				[]blockchain.TXOutput{
					{
						10,
						[]byte(""),
					},
				},
			},
		},
	}

	for _, tt := range tests {
		block := blockchain.NewGenesisBlock(&tt.transactions)
		output := block.HashTransactions()
		assert.NotNil(t, output)
	}
}
