package cmd

import (
	"log"
	"solana/pkg/client"
	"strconv"

	"github.com/gagliardetto/solana-go"
	"github.com/spf13/cobra"
)

func NewAirdropCmd() *cobra.Command {
	airdropCmd := &cobra.Command{
		Use:   "airdrop [public key] [amount]",
		Short: "Request airdrop to account",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cluster, err := ClusterFromFlag(cmd)
			if err != nil {
				return err
			}

			publicKey, err := solana.PublicKeyFromBase58(args[0])
			if err != nil {
				return err
			}

			amount, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			client, err := client.NewClient(cmd.Context(), cluster)
			if err != nil {
				return err
			}
			defer client.Close()

			sig, err := client.Airdrop(cmd.Context(), publicKey, amount)
			if err != nil {
				return err
			}

			log.Println(sig.String())
			return nil
		},
	}

	return airdropCmd
}

func NewBalanceCmd() *cobra.Command {
	balanceCmd := &cobra.Command{
		Use:   "balance [public key]",
		Short: "Get balance",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cluster, err := ClusterFromFlag(cmd)
			if err != nil {
				return err
			}

			publicKey, err := solana.PublicKeyFromBase58(args[0])
			if err != nil {
				return err
			}

			client, err := client.NewClient(cmd.Context(), cluster)
			if err != nil {
				return err
			}
			defer client.Close()

			out, err := client.Balance(cmd.Context(), publicKey)
			if err != nil {
				return err
			}

			log.Printf("Balance lamports: %d", out)
			log.Printf("Balance sol: %d,%d SOL", out/solana.LAMPORTS_PER_SOL, out%solana.LAMPORTS_PER_SOL)
			return nil
		},
	}

	return balanceCmd
}

func NewWalletCmd() *cobra.Command {
	walletCmd := &cobra.Command{
		Use:   "wallet",
		Short: "Create new wallet",
		Run: func(cmd *cobra.Command, args []string) {
			wallet := solana.NewWallet()

			log.Println("Private key: ", wallet.PrivateKey.String())
			log.Println("Public key: ", wallet.PublicKey().String())
		},
	}

	return walletCmd
}
