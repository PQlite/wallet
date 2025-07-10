package crypto

import (
	"encoding/hex"
	"encoding/json"
	"log"
	"wallet/tx"

	"github.com/cloudflare/circl/sign"
	"github.com/cloudflare/circl/sign/schemes"
	"golang.org/x/crypto/sha3"
)

// Sign підписує транзакцію і повертаю повноцінну Transaction з підписом і pubkey
func Sign(unsignTransaction tx.UnsignTransaction, priv sign.PrivateKey, pub sign.PublicKey) (tx.Transaction, error) {
	scheme := schemes.ByName("Dilithium3")

	// TODO: SerializeTransaction
	rawData, err := json.Marshal(unsignTransaction)
	if err != nil {
		log.Println("помилка json.Marshal: ", err)
		return tx.Transaction{}, err
	}

	rawDataHash := sha3.Sum224(rawData)

	sig := scheme.Sign(priv, rawDataHash[:], nil)
	pubBytes, err := pub.MarshalBinary()
	if err != nil {
		log.Println("помилка pub.MarshalBinary: ", err)
		return tx.Transaction{}, err
	}

	return tx.Transaction{
		From: unsignTransaction.From,
		To:   unsignTransaction.To, Amount: unsignTransaction.Amount,
		Timestamp: unsignTransaction.Timestamp,
		Nonce:     unsignTransaction.Nonce,
		Signature: hex.EncodeToString(sig),
		PubKey:    hex.EncodeToString(pubBytes),
	}, nil
}

// Verify робить перевірку підпису транзакції і власника гаманця з поля From
func Verify(transaction tx.Transaction) bool {
	scheme := schemes.ByName("Dilithium3")

	pubKeyBytes, err := hex.DecodeString(transaction.PubKey)
	if err != nil {
		log.Println("помилка hex.DecodeString (PubKey): ", err)
		return false
	}

	pubKey, err := scheme.UnmarshalBinaryPublicKey(pubKeyBytes)
	if err != nil {
		log.Println("помилка UnmarshalBinaryPublicKey: ", err)
		return false
	}

	signatureBytes, err := hex.DecodeString(transaction.Signature)
	if err != nil {
		log.Println("помилка hex.DecodeString (Signature): ", err)
		return false
	}

	unsignTransaction := transaction.GetUnsignTransaction()
	rawData, err := json.Marshal(unsignTransaction)
	if err != nil {
		log.Println("помилка json.Marshal: ", err)
		return false
	}

	rawDataHash := sha3.Sum224(rawData)

	// check that 'from' belongs to the transaction creator
	pubKeySum := sha3.Sum224(pubKeyBytes)
	if hex.EncodeToString(pubKeySum[:]) != transaction.From {
		return false
	}

	return scheme.Verify(pubKey, rawDataHash[:], signatureBytes, nil)
}
