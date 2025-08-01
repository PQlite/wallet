package cmd

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/PQlite/crypto"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/sha3"
)

type Transaction struct {
	From      string  `json:"from"`
	To        string  `json:"to"`
	Amount    float32 `json:"amount"`
	Timestamp int64   `json:"timestamp"`
	Nonce     int     `json:"nonce"`
	PubKey    []byte  `json:"pubkey"`
	Signature []byte  `json:"signature"`
}

var SendCmd = &cobra.Command{
	Use:   "send",
	Short: "send `from` `to` `amount`",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 3 {
			log.Println("Використання: send <публічний_ключ_відправника> <адреса_отримувача> <сума>")
			return
		}

		f64, err := strconv.ParseFloat(args[2], 32)
		if err != nil {
			log.Println("помилка розпаковки amount")
			return
		}

		wallets, err := LoadWallets()
		if err != nil {
			log.Println("Failed to load wallets: ", err)
			return
		}

		var senderWallet Wallet
		found := false

		for _, w := range wallets {
			pubBytes, err := hex.DecodeString(w.Pub)
			if err != nil {
				log.Println("помилка конвертації pubkey в байти", err)
				return
			}
			pubHash := sha3.Sum224(pubBytes)
			if hex.EncodeToString(pubHash[:]) == args[0] {
				senderWallet = w
				found = true
				break
			}
		}

		if !found {
			log.Println("Sender public key not found in wallet.json")
			return
		}

		senderPrivKeyBytes, err := hex.DecodeString(senderWallet.Priv)
		if err != nil {
			log.Println("Failed to decode private key: ", err)
			return
		}

		senderPubKeyBytes, err := hex.DecodeString(senderWallet.Pub)
		if err != nil {
			log.Println("Failed to decode public key: ", err)
			return
		}

		tx := &Transaction{
			From:      args[0],
			To:        args[1],
			Amount:    float32(f64),
			Timestamp: time.Now().UnixMilli(),
			Nonce:     1, // TODO: nonce
			PubKey:    senderPubKeyBytes,
		}

		unsignedTxBytes, err := json.Marshal(tx)
		if err != nil {
			log.Println("помилка marshal unsignedTx")
			return
		}

		signature, err := crypto.Sign(senderPrivKeyBytes, unsignedTxBytes)
		if err != nil {
			log.Println("Failed to sign transaction: ", err)
			return
		}
		tx.Signature = signature

		body, err := json.Marshal(tx)
		if err != nil {
			log.Println("помилка marshal signedTx")
			return
		}

		r, err := http.Post("http://localhost:8081/tx", "application/json", bytes.NewBuffer(body))
		if err != nil {
			log.Println("помилка відправки транзакції")
			return
		}
		log.Println(r.StatusCode)
	},
}
