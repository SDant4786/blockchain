package blockchain_test

import (
	"blockchain/src/blockchain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBase58Encode(t *testing.T) {
	var tests = []struct {
		name   string
		input  []byte
		output []byte
	}{
		{"Pass", []byte("foo"), []byte{0x31, 0x62, 0x51, 0x62, 0x70}},
	}

	for _, tt := range tests {
		out := blockchain.Base58Encode(tt.input)
		assert.Equal(t, out, tt.output)
	}
}

func TestBase58Decode(t *testing.T) {
	var tests = []struct {
		name   string
		input  []byte
		output []byte
	}{
		{"Pass", []byte{0x31, 0x62, 0x51, 0x62, 0x70}, []byte{0x0, 0x66, 0x6f, 0x6f}},
	}

	for _, tt := range tests {
		out := blockchain.Base58Decode(tt.input)
		assert.Equal(t, out, tt.output)
	}
}
