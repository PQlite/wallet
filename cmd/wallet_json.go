package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/PQlite/crypto"

	"golang.org/x/term"
)

const walletFileName = "wallet.json"

// getPassword запитує у користувача пароль
func getPassword() (string, error) {
	fmt.Print("Enter password: ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", err
	}
	password := string(bytePassword)
	fmt.Println()
	return strings.TrimSpace(password), nil
}

// LoadWallets читає гаманці з файлу wallet.json
func LoadWallets() ([]crypto.Wallet, error) {
	var wallets []crypto.Wallet

	if _, err := os.Stat(walletFileName); os.IsNotExist(err) {
		return wallets, nil // Повертаємо порожній список, якщо файл не існує
	}

	encryptedWalletJSON, err := os.ReadFile(walletFileName)
	if err != nil {
		return nil, fmt.Errorf("Failed to read %s: %w", walletFileName, err)
	}

	if len(encryptedWalletJSON) == 0 {
		return wallets, nil // Повертаємо порожній список, якщо файл порожній
	}

	password, err := getPassword()
	if err != nil {
		return nil, err
	}

	walletJSON, err := crypto.Decrypt(encryptedWalletJSON, password)
	if err != nil {
		return nil, fmt.Errorf("Failed to decrypt %s: %w", walletFileName, err)
	}

	err = json.Unmarshal(walletJSON, &wallets)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal %s: %w", walletFileName, err)
	}

	return wallets, nil
}

// SaveWallets зберігає гаманці у файл wallet.json
func SaveWallets(wallets []crypto.Wallet) error {
	updatedWalletJSON, err := json.MarshalIndent(wallets, "", "  ")
	if err != nil {
		return fmt.Errorf("Failed to marshal wallets to JSON: %w", err)
	}

	password, err := getPassword()
	if err != nil {
		return err
	}

	encryptedData, err := crypto.Encrypt(updatedWalletJSON, password)
	if err != nil {
		return fmt.Errorf("Failed to encrypt wallets: %w", err)
	}

	err = os.WriteFile(walletFileName, encryptedData, 0644)
	if err != nil {
		return fmt.Errorf("Failed to write %s: %w", walletFileName, err)
	}

	return nil
}
