package cmd

import (
	"encoding/hex"
	"log"
	"strconv"
	"time"

	"golang.org/x/crypto/sha3"

	"github.com/PQlite/crypto"
	"github.com/spf13/cobra"
)

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

		var senderWallet crypto.Wallet
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

		senderPrivKey, senderPubKey, err := crypto.LoadKeyPair(senderWallet)
		if err != nil {
			log.Println("Failed to load key pair: ", err)
			return
		}

		unsignTransaction := crypto.UnsignTransaction{
			From:      args[0],
			To:        args[1],
			Amount:    float32(f64),
			Timestamp: time.Now().Unix(),
			Nonce:     1,
		}

		signedTx, err := crypto.Sign(unsignTransaction, senderPrivKey, senderPubKey)
		if err != nil {
			log.Println("Failed to sign transaction: ", err)
			return
		}

		log.Printf("Signed Transaction: %+v\n", signedTx)
		isValid := crypto.Verify(signedTx)
		log.Println("верифікація транзакції повертає: ", isValid)
	},
}
