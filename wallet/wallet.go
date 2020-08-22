package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
)

type Wallet struct {
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
	address    string
}

func NewWallet() *Wallet {
	// Creating ECDSA PrivateKey (32 bytes) key and PublicKey (64 bytes)
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	publicKey := privateKey.PublicKey
	address := generateWalletAddress(publicKey)
	return &Wallet{privateKey, &publicKey, address}
}

// The private key
func (w Wallet) PrivateKey() *ecdsa.PrivateKey {
	return w.privateKey
}

// The private key as a string
func (w Wallet) PrivateKeyString() string {
	return fmt.Sprintf("%x", w.privateKey.D.Bytes())
}

// The public key
func (w Wallet) PublicKey() *ecdsa.PublicKey {
	return w.publicKey
}

// The public key as a string
func (w Wallet) PublicKeyString() string {
	return fmt.Sprintf("%x%x", w.publicKey.X, w.publicKey.Y)
}

// The blockchain address
func (w Wallet) Address() string {
	return w.address
}

func (w Wallet) SendTo(recipientAddress string, value float32) *Transaction {
	return NewTransaction(w.PrivateKey(), w.PublicKey(), w.Address(), recipientAddress, value)
}

func generateWalletAddress(publicKey ecdsa.PublicKey) string {
	// SHA256 hashing on the PublicKey (32 bytes)
	h2 := sha256.New()
	h2.Write(publicKey.X.Bytes())
	h2.Write(publicKey.Y.Bytes())
	digest2 := h2.Sum(nil)
	// RIPEMD160 hashing (20 bytes)
	h3 := ripemd160.New()
	h3.Write(digest2)
	digest3 := h3.Sum(nil)
	// Add version byte
	vd4 := make([]byte, 21)
	vd4[0] = 0x00
	copy(vd4[1:], digest3[:])
	// SHA256 hashing
	h5 := sha256.New()
	h5.Write(vd4)
	digest5 := h5.Sum(nil)
	// Hashing again
	h6 := sha256.New()
	h6.Write(digest5)
	digest6 := h6.Sum(nil)
	// First 4 bytes for checksum
	checksum := digest6[:4]
	// Add the checksum at the end
	dc8 := make([]byte, 25)
	copy(dc8[:21], vd4[:])
	copy(dc8[21:], checksum)
	// Convert from bytes to base 58
	address := base58.Encode(dc8)
	return address
}
