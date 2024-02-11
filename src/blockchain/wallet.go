package blockchain

import (
	"crypto/ecdh"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const (
	version            = byte(0x00)
	walletFile         = "wallet_%s.dat"
	addressChecksumLen = 4
)

type Wallet struct {
	PrivateKey ecdh.PrivateKey
	PublicKey  []byte
}

func (w Wallet) GetAddress() []byte {
	pubKeyHash := HashPubKey(w.PublicKey)

	versionedPayload := append([]byte{version}, pubKeyHash...)
	checksum := checksum(versionedPayload)

	fullPayload := append(versionedPayload, checksum...)
	address := Base58Encode(fullPayload)

	return address
}

type Wallets struct {
	Wallets map[string]*Wallet
}

func (ws *Wallets) CreateWallet() string {
	wallet := NewWallet()
	address := fmt.Sprintf("%s", wallet.GetAddress())

	ws.Wallets[address] = wallet

	return address
}

func (ws *Wallets) GetAddresses() []string {
	var addresses []string

	for address := range ws.Wallets {
		addresses = append(addresses, address)
	}

	return addresses
}

func (ws Wallets) GetWallet(address string) Wallet {
	return *ws.Wallets[address]
}

func (ws *Wallets) LoadFromFile(nodeID string) error {
	walletFile := fmt.Sprintf(walletFile, nodeID)
	if _, err := os.Stat(walletFile); os.IsNotExist(err) {
		return err
	}

	fileContent, err := os.ReadFile(walletFile)
	if err != nil {
		log.Panic(err)
	}

	var wallets Wallets
	err = json.Unmarshal(fileContent, &wallets)
	if err != nil {
		log.Panic(err)
	}

	ws.Wallets = wallets.Wallets

	return nil
}

func (ws Wallets) SaveToFile(nodeID string) {
	walletFile := fmt.Sprintf(walletFile, nodeID)

	jsonData, err := json.Marshal(ws)
	fmt.Println(string(jsonData))
	if err != nil {
		log.Panic(err)
	}

	err = os.WriteFile(walletFile, jsonData, 0644)
	if err != nil {
		log.Panic(err)
	}
}
