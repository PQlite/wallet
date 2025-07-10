package tests

import (
	"encoding/hex"
	"testing"
	"time"
	"wallet/crypto"
	"wallet/tx"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	pub, priv, err := crypto.Create()
	assert.NoError(t, err)
	assert.NotNil(t, pub)
	assert.NotNil(t, priv)
}

func TestLoadKeyPair(t *testing.T) {
	// Create a key pair first
	pub, priv, err := crypto.Create()
	assert.NoError(t, err)

	pubBytes, err := pub.MarshalBinary()
	assert.NoError(t, err)
	privBytes, err := priv.MarshalBinary()
	assert.NoError(t, err)

	w := crypto.Wallet{
		Priv: hex.EncodeToString(privBytes),
		Pub:  hex.EncodeToString(pubBytes),
	}

	loadedPriv, loadedPub, err := crypto.LoadKeyPair(w)
	assert.NoError(t, err)
	assert.NotNil(t, loadedPriv)
	assert.NotNil(t, loadedPub)

	// Verify that the loaded keys are the same as the original keys
	loadedPrivBytes, err := loadedPriv.MarshalBinary()
	assert.NoError(t, err)
	loadedPubBytes, err := loadedPub.MarshalBinary()
	assert.NoError(t, err)

	assert.Equal(t, privBytes, loadedPrivBytes)
	assert.Equal(t, pubBytes, loadedPubBytes)
}

func TestSignAndVerify(t *testing.T) {
	// Create a key pair
	pub, priv, err := crypto.Create()
	assert.NoError(t, err)

	// Create an unsigned transaction
	unsignTx := tx.UnsignTransaction{
		From:      "820e1f5975091380c6cfa7c0df95aaa64a8dd53254f58c31b7cb5c98",
		To:        "c2b0b56de043094ccd124588c1b972efcd4fcc33ddc1102e156a4090",
		Amount:    10.0,
		Timestamp: time.Now().Unix(),
		Nonce:     1,
	}

	// Sign the transaction
	signedTx, err := crypto.Sign(unsignTx, priv, pub)
	assert.NoError(t, err)
	assert.NotEmpty(t, signedTx.Signature)
	assert.NotEmpty(t, signedTx.PubKey)

	// Verify the transaction
	isValid := crypto.Verify(signedTx)
	assert.False(t, isValid)

	// Test with a tampered transaction
	tamperedTx := signedTx
	tamperedTx.Amount = 20.0 // Change amount
	isValid = crypto.Verify(tamperedTx)
	assert.False(t, isValid)

	// Test with invalid signature
	invalidSigTx := signedTx
	invalidSigTx.Signature = "invalid_signature"
	isValid = crypto.Verify(invalidSigTx)
	assert.False(t, isValid)

	// Test with invalid public key
	invalidPubKeyTx := signedTx
	invalidPubKeyTx.PubKey = "invalid_pubkey"
	isValid = crypto.Verify(invalidPubKeyTx)
	assert.False(t, isValid)
}
