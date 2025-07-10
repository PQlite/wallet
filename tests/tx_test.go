package tests

import (
	"testing"
	"time"

	"wallet/tx"

	"github.com/stretchr/testify/assert"
)

func TestGetUnsignTransaction(t *testing.T) {
	tx := tx.Transaction{
		From:      "sender",
		To:        "receiver",
		Amount:    10.0,
		Timestamp: time.Now().Unix(),
		Nonce:     1,
		Signature: "some_signature",
		PubKey:    "some_pubkey",
	}

	unsignTx := tx.GetUnsignTransaction()

	assert.Equal(t, tx.From, unsignTx.From)
	assert.Equal(t, tx.To, unsignTx.To)
	assert.Equal(t, tx.Amount, unsignTx.Amount)
	assert.Equal(t, tx.Timestamp, unsignTx.Timestamp)
	assert.Equal(t, tx.Nonce, unsignTx.Nonce)
}
