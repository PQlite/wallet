package cmd

import (
	"encoding/hex"
	"log"
	"wallet/crypto"

	"github.com/spf13/cobra"
)

var GenCmd = &cobra.Command{
	Use:   "gen",
	Short: "generate keypair",
	Run: func(cmd *cobra.Command, args []string) {
		pub, priv, err := crypto.Create()
		if err != nil {
			log.Println(err)
			return
		}
		privBytes, err := priv.MarshalBinary()
		if err != nil {
			log.Println(err)
			return
		}
		pubBytes, err := pub.MarshalBinary()
		if err != nil {
			log.Println(err)
			return
		}

		newWallet := crypto.Wallet{
			Priv: hex.EncodeToString(privBytes),
			Pub:  hex.EncodeToString(pubBytes),
		}

		wallets, err := LoadWallets()
		if err != nil {
			log.Println(err)
			return
		}

		// Append new wallet
		wallets = append(wallets, newWallet)

		err = SaveWallets(wallets)
		if err != nil {
			log.Println(err)
			return
		}

		log.Println("Нові ключі згенеровано та додано до wallet.json")
	},
}

func init() {
	rootCmd.AddCommand(GenCmd)
}
