package cmd

import (
	"log"
	"solana/pkg/client"
	"solana/pkg/instructions"
	"solana/pkg/model"
	"strconv"

	"github.com/gagliardetto/solana-go"
	"github.com/spf13/cobra"
)

func NewSaberCmd() *cobra.Command {
	saberCmd := &cobra.Command{
		Use:   "saber",
		Short: "Work with solana saber dex",
	}

	saberCmd.AddCommand(newSaberSwapPoolsCmd())
	saberCmd.AddCommand(newSaberSwapCmd())

	return saberCmd
}

func newSaberSwapPoolsCmd() *cobra.Command {
	poolsInfoCmd := &cobra.Command{
		Use:              "pools",
		Short:            "Swap pools info",
		Long:             "Get swap pools info from url",
		TraverseChildren: true,
		RunE: func(cmd *cobra.Command, args []string) error {

			cluster, err := ClusterFromFlag(cmd)
			if err != nil {
				return err
			}

			url, err := client.SwapUrlFromCluster(cluster)
			if err != nil {
				return err
			}

			swapInfo, err := model.NewJsonSwapInfo(url)
			if err != nil {
				return err
			}

			swapInfo.ListPools()
			return nil
		},
	}

	return poolsInfoCmd
}

func newSaberSwapCmd() *cobra.Command {

	var programIdKey string
	var privateKey string
	var showAccounts bool

	saberSwapCmd := &cobra.Command{
		Use:   "swap [swap account] [amount] [token mint a] [amount] [token mint b]",
		Short: "Swap tokens",
		Args:  cobra.MinimumNArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			cluster, err := ClusterFromFlag(cmd)
			if err != nil {
				return err
			}

			amountTokenA, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			amountTokenB, err := strconv.ParseUint(args[3], 10, 64)
			if err != nil {
				return err
			}

			swapData := instructions.NewSwapData(amountTokenA, amountTokenB)

			swapAccount, err := solana.PublicKeyFromBase58(args[0])
			if err != nil {
				return err
			}

			tokenA, err := solana.PublicKeyFromBase58(args[2])
			if err != nil {
				return err
			}

			tokenB, err := solana.PublicKeyFromBase58(args[4])
			if err != nil {
				return err
			}

			programId, err := solana.PublicKeyFromBase58(programIdKey)
			if err != nil {
				return err
			}

			wallet, err := solana.WalletFromPrivateKeyBase58(privateKey)
			if err != nil {
				return err
			}

			client, err := client.NewClient(cmd.Context(), cluster)
			if err != nil {
				return err
			}
			defer client.Close()

			sig, err := client.Swap(cmd.Context(), programId, swapAccount, tokenA, tokenB, wallet, swapData, showAccounts)
			if err != nil {
				return err
			}

			log.Print(sig.String())

			return nil
		},
	}

	saberSwapCmd.PersistentFlags().StringVarP(&privateKey, "private", "p", "", "Private key")
	saberSwapCmd.Flags().StringVarP(&programIdKey, "program", "", "SSwpkEEcbUqx4vtoEByFjSkhKdCT862DNVb52nZg1UZ", "Stabe Swap Program Account")
	saberSwapCmd.Flags().BoolVarP(&showAccounts, "show", "s", false, "Show accounts in instruction (Don't send transaction)")
	return saberSwapCmd
}
