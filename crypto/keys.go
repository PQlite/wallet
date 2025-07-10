// Package crypto відповідає за створення та використання ключів
package crypto

import (
	"encoding/hex"
	"log"

	"github.com/cloudflare/circl/sign"
	"github.com/cloudflare/circl/sign/schemes"
)

type Wallet struct {
	Priv string `json:"priv"`
	Pub  string `json:"pub"`
}

func Create() (sign.PublicKey, sign.PrivateKey, error) {
	scheme := schemes.ByName("Dilithium3")
	pub, priv, err := scheme.GenerateKey()
	if err != nil {
		log.Println("помилка створення пари priv/pub ", err)
		return nil, nil, err
	}

	return pub, priv, err
}

// LoadKeyPair loads private and public key from Wallet object
func LoadKeyPair(w Wallet) (sign.PrivateKey, sign.PublicKey, error) {
	privBytes, err := hex.DecodeString(w.Priv)
	if err != nil {
		return nil, nil, err
	}
	pubBytes, err := hex.DecodeString(w.Pub)
	if err != nil {
		return nil, nil, err
	}

	scheme := schemes.ByName("Dilithium3")
	privKey, err := scheme.UnmarshalBinaryPrivateKey(privBytes)
	if err != nil {
		return nil, nil, err
	}
	pubKey, err := scheme.UnmarshalBinaryPublicKey(pubBytes)
	if err != nil {
		return nil, nil, err
	}

	return privKey, pubKey, nil
}
