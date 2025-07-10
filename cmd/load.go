package cmd

import (
	"encoding/hex"
	"fmt"
	"log"

	"golang.org/x/crypto/sha3"

	"github.com/spf13/cobra"
)

var LoadCmd = &cobra.Command{
	Use:   "load",
	Short: "Load keys from wallet.json",
	Run: func(cmd *cobra.Command, args []string) {
		wallets, err := LoadWallets()
		if err != nil {
			log.Println(err)
			return
		}

		if len(wallets) == 0 {
			fmt.Println("No wallets found in wallet.json")
			return
		}

		fmt.Println("Loaded Wallets:")
		for i, wallet := range wallets {
			fmt.Printf("Wallet %d:\n", i+1)
			pubBytes, err := hex.DecodeString(wallet.Pub)
			if err != nil {
				log.Println(err)
				return
			}
			pubSum := sha3.Sum224(pubBytes)
			// fmt.Println("  Private Key:", wallet.Priv)
			fmt.Println("  Публічний ключ:", hex.EncodeToString(pubSum[:]))
		}
	},
}

func init() {
	rootCmd.AddCommand(LoadCmd)
}
