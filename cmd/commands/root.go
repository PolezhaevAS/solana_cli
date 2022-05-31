package cmd

import "github.com/spf13/cobra"

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "solana",
		Short: "Solana cli tool",
	}

	rootCmd = SetRootFlgas(rootCmd)

	rootCmd.CompletionOptions.DisableDefaultCmd = true

	rootCmd.AddCommand(NewSaberCmd())
	rootCmd.AddCommand(NewAirdropCmd())
	rootCmd.AddCommand(NewBalanceCmd())
	rootCmd.AddCommand(NewWalletCmd())

	return rootCmd
}
