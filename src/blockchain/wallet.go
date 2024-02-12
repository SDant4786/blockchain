package blockchain

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/x509"
	"encoding/gob"
	"encoding/pem"
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
	PrivateKey ecdsa.PrivateKey
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

type walletIntermediary struct {
	PrivateKey []byte
	PublicKey  []byte
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

func encode(privateKey *ecdsa.PrivateKey) []byte {
	x509Encoded, _ := x509.MarshalECPrivateKey(privateKey)
	pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: x509Encoded})

	return pemEncoded
}

func decode(pemEncoded []byte) *ecdsa.PrivateKey {
	block, _ := pem.Decode(pemEncoded)
	x509Encoded := block.Bytes
	privateKey, _ := x509.ParseECPrivateKey(x509Encoded)

	return privateKey
}
