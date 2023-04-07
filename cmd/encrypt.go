package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var encryptCmd = &cobra.Command{
	Use: "encrypt",
	Short: "encryption program	",
	Long: `nil`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(key) != 8 {
			fmt.Println("invalid entry, key must be 8 bits (64 bytes)")
		} else {
			fmt.Printf("You key : %s\n", key)
			fmt.Printf("You Text: %s\n", text)
			fmt.Printf("Encrypt result:%s\n", Encrypt(key, text))
		}
	},
}

func init() {
	rootCmd.AddCommand(encryptCmd)
}
