package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "Start an interactive wallet shell",
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Wallet Shell. Type 'h' for commands, 'exit' to quit.")

		for {
			fmt.Print("? ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			if input == "exit" {
				fmt.Println("Exiting shell.")
				break
			}

			if input == "h" {
				fmt.Println("Available commands:")
				fmt.Println("  gen - Generate a new key pair and save to wallet.json")
				fmt.Println("  load - Load and display all key pairs from wallet.json")
				fmt.Println("  exit - Exit the shell")
				fmt.Println("  send")
				continue
			}

			parts := strings.Fields(input)
			if len(parts) == 0 {
				continue
			}

			command := parts[0]
			cmdArgs := []string{}
			if len(parts) > 1 {
				cmdArgs = parts[1:]
			}

			switch command {
			case "gen":
				GenCmd.Run(GenCmd, cmdArgs)
			case "load":
				LoadCmd.Run(LoadCmd, cmdArgs)
			case "send":
				SendCmd.Run(SendCmd, cmdArgs)
			default:
				fmt.Printf("Невідома команда: %s. Введіть 'h' для доступних команд.", command)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(shellCmd)
}
