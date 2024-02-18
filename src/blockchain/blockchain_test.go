package blockchain_test

import (
	"blockchain/src/blockchain"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testBlockChain *blockchain.Blockchain

func Init() {
	bc := blockchain.CreateBlockchain("foobar", "1")
	testBlockChain = bc
	UTXOSet := blockchain.UTXOSet{bc}
	UTXOSet.Reindex()
	bc.CloseDB()
}
func TestNewBlockchain(t *testing.T) {
	tests := []struct {
		name   string
		nodeID string
	}{
		{
			"Pass",
			"1",
		},
	}

	for _, tt := range tests {
		newBlockchain := blockchain.NewBlockchain(tt.nodeID)
		assert.NotNil(t, newBlockchain)
	}
}

func TestCreateBlockchain(t *testing.T) {
	tests := []struct {
		name    string
		nodeID  string
		address string
	}{
		{
			"Pass",
			"2",
			"testAddress",
		},
	}

	for _, tt := range tests {
		bc := blockchain.CreateBlockchain(tt.address, tt.nodeID)
		UTXOSet := blockchain.UTXOSet{bc}
		UTXOSet.Reindex()
		bc.CloseDB()
		assert.NotNil(t, bc)
	}
}

func TestMineBlock(t *testing.T) {
	tests := []struct {
		name         string
		transactions []*blockchain.Transaction
	}{
		{
			"Pass",
			[]*blockchain.Transaction{
				&blockchain.Transaction{
					[]byte("1"),
					[]blockchain.TXInput{
						{
							[]byte("1"),
							1,
							[]byte("1"),
							[]byte("1"),
						},
					},
					[]blockchain.TXOutput{
						{
							1,
							[]byte("1"),
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		block := testBlockChain.MineBlock(tt.transactions)

		assert.NotNil(t, block)
	}
}
