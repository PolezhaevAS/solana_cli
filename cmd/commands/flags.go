package cmd

import (
	"solana/pkg/client"

	"github.com/gagliardetto/solana-go/rpc"
	"github.com/spf13/cobra"
)

func SetRootFlgas(rootCmd *cobra.Command) *cobra.Command {
	rootCmd.PersistentFlags().StringP("cluster", "c", "dev", "RPC cluster. Mainnet - main, Devnet - dev")
	return rootCmd
}

func ClusterFromFlag(cmd *cobra.Command) (rpc.Cluster, error) {
	flags := cmd.InheritedFlags()
	clusterFlag, err := flags.GetString("cluster")
	if err != nil {
		return rpc.Cluster{}, err
	}

	cluster, err := client.ClusterFromString(clusterFlag)
	if err != nil {
		return rpc.Cluster{}, err
	}

	return cluster, nil
}
