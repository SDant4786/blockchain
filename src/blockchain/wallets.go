package blockchain

import (
	"bytes"
	"crypto/elliptic"
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

type Wallets struct {
	Wallets map[string]*Wallet
}

type walletIntermediary struct {
	PrivateKey []byte
	PublicKey  []byte
}

func NewWallets(nodeID string) (*Wallets, error) {
	wallets := Wallets{}
	wallets.Wallets = make(map[string]*Wallet)

	err := wallets.LoadFromFile(nodeID)

	return &wallets, err
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

	walletsIntermediary := map[string]walletIntermediary{}
	gob.Register(elliptic.P256())
	decoder := gob.NewDecoder(bytes.NewReader(fileContent))
	err = decoder.Decode(&walletsIntermediary)
	if err != nil {
		log.Panic(err)
	}
	wallets := Wallets{Wallets: map[string]*Wallet{}}

	for name, w := range walletsIntermediary {
		privateKey := decode(w.PrivateKey)
		w := Wallet{PrivateKey: *privateKey, PublicKey: w.PublicKey}
		wallets.Wallets[name] = &w
	}
	ws.Wallets = wallets.Wallets

	return nil
}

func (ws Wallets) SaveToFile(nodeID string) {
	var content bytes.Buffer
	walletsIntermediary := map[string]walletIntermediary{}
	walletFile := fmt.Sprintf(walletFile, nodeID)

	for name, wallet := range ws.Wallets {
		privateKey := encode(&wallet.PrivateKey)
		walletsIntermediary[name] = walletIntermediary{
			PrivateKey: privateKey,
			PublicKey:  wallet.PublicKey,
		}
	}
	gob.Register(elliptic.P256())

	encoder := gob.NewEncoder(&content)
	err := encoder.Encode(walletsIntermediary)
	if err != nil {
		log.Panic(err)
	}

	err = os.WriteFile(walletFile, content.Bytes(), 0644)
	if err != nil {
		log.Panic(err)
	}
}
