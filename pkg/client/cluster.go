package client

import (
	"fmt"

	"github.com/gagliardetto/solana-go/rpc"
)

func ClusterFromString(cluster string) (rpc.Cluster, error) {
	switch cluster {
	case "main":
		return rpc.MainNetBeta, nil
	case "dev":
		return rpc.DevNet, nil
	case "test":
		return rpc.TestNet, nil
	case "local":
		return rpc.LocalNet, nil
	default:
		return rpc.Cluster{}, fmt.Errorf("cann't parse cluster flag - %s", cluster)
	}
}

func SwapUrlFromCluster(cluster rpc.Cluster) (string, error) {
	switch cluster {
	case rpc.MainNetBeta:
		return "https://registry.saber.so/data/pools-info.mainnet.json", nil
	case rpc.DevNet:
		return "https://registry.saber.so/data/pools-info.devnet.json", nil
	default:
		return "", fmt.Errorf("cann't find url on cluster %s", cluster.Name)
	}
}
