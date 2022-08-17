package chain_info

import "fmt"

type ExplorerTargetType int

const (
	ExplorerTargetTransaction ExplorerTargetType = 1
	ExplorerTargetAddress     ExplorerTargetType = 2
)

func GetExplorerTargetUrl(chainId uint64, txHash string, targetType ExplorerTargetType) string {
	net := GetNetByChainId(chainId)
	if net == nil {
		return fmt.Sprintf("invalid chain id %d", chainId)
	}

	var target string

	switch targetType {
	case ExplorerTargetTransaction:
		target = "tx"
	case ExplorerTargetAddress:
		target = "address"
	}

	return fmt.Sprintf("%s%s/%s", net.BlockExplorer, target, txHash)
}
