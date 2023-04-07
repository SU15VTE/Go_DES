package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var key string
var text string
var rootCmd = &cobra.Command{
	Use:   "des",
	Short: "A brief description of your application",
	Long: `GO_DES is a program for DES encryption and decryption. You can use it like this:
	.\des.exe encrypt --key "SU15VTE!" --text "Hello World!"
	.\des.exe decrypt --key "SU15VTE!" --text "1c3568382f24163e7368342104541010".`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&text, "text", "t", "", "A Text")
	rootCmd.PersistentFlags().StringVarP(&key, "key", "k", "", "The key to use for encryption/decryption")
}
