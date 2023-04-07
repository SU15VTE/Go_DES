package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "Decryption program",
	Long:  `nil`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(key) != 8 {
			fmt.Println("invalid entry, key must be 8 bits (64 bytes)")
		} else {
			fmt.Printf("You key :	%s\n", key)
			fmt.Printf("You Text:	%s\n", text)
			fmt.Printf("Decrypt result:	%s\n", Decrypt(key, text))
		}
	},
}

func init() {
	rootCmd.AddCommand(decryptCmd)
}
