package chain_info

import (
	"fmt"
	"os"
	"time"

	"github.com/tristan-club/kit/log"
)

const (
	ChainTypeBsc      = 1
	ChainTypeMetis    = 2
	ChainTypePolygon  = 3
	ChainTypeKlaytn   = 4
	ChainTypeOkc      = 5
	ChainTypeEvmos    = 6
	ChainTypeCronos   = 7
	ChainTypeAurora   = 8
	ChainTypeArbitrum = 9
	ChainTypeConflux  = 10
	ChainTypeEthereum = 100
)

var supportChainTypeList = []uint32{
	ChainTypeBsc,
	ChainTypeMetis,
	ChainTypePolygon,
	ChainTypeKlaytn,
	ChainTypeOkc,
	ChainTypeEvmos,
	ChainTypeCronos,
	ChainTypeAurora,
	ChainTypeArbitrum,
	ChainTypeConflux}

func GetSupportChainTypeList() []uint32 {
	return supportChainTypeList
}

type Net struct {
	ChainType        uint32 `json:"chain_type"`
	ChainId          uint64 `json:"chain_id"`
	NetworkName      string `json:"network_name"`
	Symbol           string `json:"symbol"`
	Decimals         uint8  `json:"decimals"`
	RpcUrl           string `json:"rpc_url"`
	OriginRpcUrl     string `json:"origin_rpc_url"`
	WssUrl           string `json:"-"`
	BlockExplorer    string `json:"block_explorer"`
	ExplorerApiUrl   string `json:"-"`
	ExplorerApiKey   string `json:"-"`
	Type             uint8  `json:"type"`
	PollingInterval  int64  `json:"polling_interval"`
	AverageBlockTime int64  `json:"average_block_time"`
}

func (c *Net) GetPollingInterval() int64 {
	if c.PollingInterval != 0 {
		return c.PollingInterval
	}
	return 5000
}

func (c *Net) GetAverageBlockTime() int64 {
	if c.AverageBlockTime != 0 {
		return c.AverageBlockTime
	}
	return int64(10 * time.Second)
}

func (c *Net) UseWssSubscribe() bool {
	return c.WssUrl != ""
}

func (c *Net) IsAvailable() bool {
	return c.ChainId == 0
}

type Chain struct {
	ChainType    uint32 `json:"chain_type"`
	Name         string `json:"name"`
	Remark       string `json:"remark"`
	Icon         string `json:"icon"`
	Symbol       string `json:"coin_symbol"`
	CoinDecimals uint8  `json:"coin_decimals"`
	Type         uint8  `json:"type"`
}

var supportChainList = []*Chain{
	{
		ChainType:    ChainTypeBsc,
		Name:         "BNB Chain",
		Symbol:       "BNB",
		CoinDecimals: 18,
		Type:         NetworkTypeMainNet,
	},
	{

		ChainType:    ChainTypeMetis,
		Name:         "Metis",
		Symbol:       "Metis",
		CoinDecimals: 18,
		Type:         NetworkTypeMainNet,
	},
	{
		ChainType:    ChainTypePolygon,
		Name:         "Polygon",
		Symbol:       "Matic",
		CoinDecimals: 18,
		Type:         NetworkTypeMainNet,
	},
	{
		ChainType:    ChainTypeKlaytn,
		Name:         "Klaytn",
		Symbol:       "KLAY",
		CoinDecimals: 18,
		Type:         NetworkTypeMainNet,
	},
	{
		ChainType:    ChainTypeOkc,
		Name:         "OKC",
		Symbol:       "OKC",
		CoinDecimals: 18,
		Type:         NetworkTypeMainNet,
	},
	{
		ChainType:    ChainTypeEvmos,
		Name:         "Evmos",
		Symbol:       "EVMOS",
		CoinDecimals: 18,
		Type:         NetworkTypeMainNet,
	},
	{
		ChainType:    ChainTypeCronos,
		Name:         "Cronos",
		Symbol:       "CRO",
		CoinDecimals: 18,
		Type:         NetworkTypeMainNet,
	},
	{
		ChainType:    ChainTypeAurora,
		Name:         "Aurora",
		Symbol:       "ETH",
		CoinDecimals: 18,
		Type:         NetworkTypeMainNet,
	},
	{
		ChainType:    ChainTypeArbitrum,
		Name:         "Arbitrum",
		Symbol:       "ETH",
		CoinDecimals: 18,
		Type:         NetworkTypeMainNet,
	},
	{
		ChainType:    ChainTypeConflux,
		Name:         "Conflux",
		Symbol:       "CFX",
		CoinDecimals: 18,
		Type:         NetworkTypeMainNet,
	},
}

func GetChainInfo(chainType uint32) *Chain {
	for _, v := range supportChainList {
		if v.ChainType == chainType {
			return v
		}
	}
	log.Error().Msgf("get chain info invalid chain type param %d", chainType)

	return &Chain{
		ChainType:    0,
		Symbol:       "-1",
		CoinDecimals: 0,
		Type:         0,
	}
}

const (
	NetworkTypeMainNet = iota + 1
	NetworkTypeTestNet
)

var supportChainNetList = []*Net{}

func init() {

	supportChainNetList = []*Net{
		{
			ChainType:   ChainTypeBsc,
			ChainId:     56,
			Symbol:      "BNB",
			Decimals:    18,
			Type:        NetworkTypeMainNet,
			NetworkName: "BSC Mainnet",
			RpcUrl:      "https://bsc-dataseed1.binance.org/",
			//RpcUrl:           fmt.Sprintf("https://bsc.getblock.io/mainnet/?api_key=%s", blockIOProvider),
			BlockExplorer:    "https://bscscan.com/",
			PollingInterval:  1400,
			AverageBlockTime: 3000,
		},
		{
			ChainType:        ChainTypeBsc,
			ChainId:          97,
			Symbol:           "BNB",
			Decimals:         18,
			Type:             NetworkTypeTestNet,
			NetworkName:      "BSC Testnet",
			RpcUrl:           "https://data-seed-prebsc-1-s1.binance.org:8545/",
			BlockExplorer:    "https://testnet.bscscan.com/",
			PollingInterval:  1400,
			AverageBlockTime: 3000,
		},
		{
			ChainType:   ChainTypeMetis,
			ChainId:     1088,
			Symbol:      "Metis",
			Decimals:    18,
			Type:        NetworkTypeMainNet,
			NetworkName: "Metis Mainnet",
			RpcUrl:      "https://andromeda.metis.io/?owner=1088",
			//WssUrl:      "wss://andromeda-ws.metis.io",
			BlockExplorer:    "https://andromeda-explorer.metis.io/",
			PollingInterval:  1000,
			AverageBlockTime: 5000,
		},
		{
			ChainType:     ChainTypeMetis,
			ChainId:       588,
			Symbol:        "Metis",
			Decimals:      18,
			Type:          NetworkTypeTestNet,
			NetworkName:   "Metis TestNet",
			RpcUrl:        "https://stardust.metis.io/?owner=588",
			BlockExplorer: "https://stardust-explorer.metis.io/",
			//WssUrl:      "wss://stardust-ws.metis.io/",
			PollingInterval:  1000,
			AverageBlockTime: 5000,
		},
		{
			ChainType:        ChainTypePolygon,
			ChainId:          137,
			Symbol:           "Matic",
			Decimals:         18,
			Type:             NetworkTypeMainNet,
			NetworkName:      "Polygon Mainnet",
			RpcUrl:           "https://polygon-rpc.com/",
			BlockExplorer:    "https://polygonscan.com/",
			ExplorerApiUrl:   "https://api.polygonscan.com/api",
			PollingInterval:  750,
			AverageBlockTime: 2000,
			//WssUrl:           "wss://rpc-mainnet.matic.network/",
		},
		{
			ChainType:        ChainTypePolygon,
			ChainId:          80001,
			Symbol:           "Matic",
			Decimals:         18,
			Type:             NetworkTypeTestNet,
			NetworkName:      "Polygon TestNet",
			RpcUrl:           "https://matic-mumbai.chainstacklabs.com",
			PollingInterval:  1000,
			AverageBlockTime: 2000,
			BlockExplorer:    "https://mumbai.polygonscan.com/",
			//WssUrl:           "wss://rpc-mumbai.matic.today",
		},
		{
			ChainType:        ChainTypeKlaytn,
			ChainId:          8217,
			Symbol:           "KLAY",
			Decimals:         18,
			Type:             NetworkTypeMainNet,
			NetworkName:      "Klaytn Cypress",
			RpcUrl:           "https://public-node-api.klaytnapi.com/v1/cypress",
			BlockExplorer:    "https://scope.klaytn.com/",
			PollingInterval:  1000,
			AverageBlockTime: 1000,
		},
		{
			ChainType:        ChainTypeKlaytn,
			ChainId:          1001,
			Symbol:           "KLAY",
			Decimals:         18,
			Type:             NetworkTypeTestNet,
			NetworkName:      "Klaytn Baobab",
			RpcUrl:           "https://api.baobab.klaytn.net:8651/",
			BlockExplorer:    "https://baobab.scope.klaytn.com/",
			PollingInterval:  750,
			AverageBlockTime: 1000,
		},
		{
			ChainType:        ChainTypeOkc,
			ChainId:          66,
			Symbol:           "OKT",
			Decimals:         18,
			Type:             NetworkTypeMainNet,
			NetworkName:      "OKC Mainnet",
			RpcUrl:           "https://exchainrpc.okex.org",
			BlockExplorer:    "https://www.oklink.com/okc/",
			PollingInterval:  2000,
			AverageBlockTime: 4000,
		},
		{
			ChainType:        ChainTypeOkc,
			ChainId:          65,
			Symbol:           "OKT",
			Decimals:         18,
			Type:             NetworkTypeTestNet,
			NetworkName:      "OKC Testnet",
			RpcUrl:           "https://exchaintestrpc.okex.org",
			BlockExplorer:    "https://www.oklink.com/okc-test/",
			PollingInterval:  2000,
			AverageBlockTime: 4000,
		},
		{
			ChainType:        ChainTypeEvmos,
			ChainId:          9001,
			Symbol:           "EVMOS",
			Decimals:         18,
			Type:             NetworkTypeMainNet,
			NetworkName:      "Evmos Mainnet",
			RpcUrl:           "https://eth.bd.evmos.org:8545",
			BlockExplorer:    "https://evm.evmos.org/",
			PollingInterval:  1000,
			AverageBlockTime: 1900,
		},
		{
			ChainType:        ChainTypeEvmos,
			ChainId:          9000,
			Symbol:           "EVMOS",
			Decimals:         18,
			Type:             NetworkTypeTestNet,
			NetworkName:      "Evmos Testnet",
			RpcUrl:           "https://eth.bd.evmos.dev:8545",
			BlockExplorer:    "https://evm.evmos.dev/",
			PollingInterval:  1000,
			AverageBlockTime: 1900,
		},
		{
			ChainType:        ChainTypeCronos,
			ChainId:          25,
			Symbol:           "CRO",
			Decimals:         18,
			Type:             NetworkTypeMainNet,
			NetworkName:      "Cronos Mainnet",
			RpcUrl:           "https://evm.cronos.org",
			BlockExplorer:    "https://cronoscan.com/",
			PollingInterval:  2000,
			AverageBlockTime: 4000,
		},
		{
			ChainType:        ChainTypeCronos,
			ChainId:          338,
			Symbol:           "CRO",
			Decimals:         18,
			Type:             NetworkTypeTestNet,
			NetworkName:      "Cronos Testnet",
			RpcUrl:           "https://evm-t3.cronos.org",
			BlockExplorer:    "https://testnet.cronoscan.com/",
			PollingInterval:  2000,
			AverageBlockTime: 4000,
		},
		{
			ChainType:        ChainTypeAurora,
			ChainId:          1313161554,
			Symbol:           "ETH",
			Decimals:         18,
			Type:             NetworkTypeMainNet,
			NetworkName:      "Aurora Mainnet",
			RpcUrl:           "https://mainnet.aurora.dev",
			BlockExplorer:    "https://aurorascan.dev/",
			PollingInterval:  1000,
			AverageBlockTime: 1000,
		},
		{
			ChainType:        ChainTypeAurora,
			ChainId:          1313161555,
			Symbol:           "ETH",
			Decimals:         18,
			Type:             NetworkTypeTestNet,
			NetworkName:      "Aurora Testnet",
			RpcUrl:           "https://testnet.aurora.dev",
			BlockExplorer:    "https://testnet.aurorascan.dev/",
			PollingInterval:  1000,
			AverageBlockTime: 1000,
		},
		{
			ChainType:        ChainTypeArbitrum,
			ChainId:          42161,
			Symbol:           "ETH",
			Decimals:         18,
			Type:             NetworkTypeMainNet,
			NetworkName:      "Arbitrum Mainnet",
			RpcUrl:           "https://arb1.arbitrum.io/rpc",
			BlockExplorer:    "https://arbiscan.io/",
			PollingInterval:  1000,
			AverageBlockTime: 1000,
		},
		{
			ChainType:        ChainTypeConflux,
			ChainId:          1030,
			Symbol:           "CFX",
			Decimals:         18,
			Type:             NetworkTypeMainNet,
			NetworkName:      "Conflux Mainnet",
			RpcUrl:           "https://evm.confluxrpc.com",
			BlockExplorer:    "https://evm.confluxscan.io/",
			PollingInterval:  1000,
			AverageBlockTime: 1000,
		},
	}

	if os.Getenv("IGNORE_NODE_BALANCE") != "1" {
		for k, _ := range supportChainNetList {
			supportChainNetList[k].OriginRpcUrl = supportChainNetList[k].RpcUrl
			supportChainNetList[k].RpcUrl = fmt.Sprintf("http://node-balance-svc/%d", supportChainNetList[k].ChainId)
		}
	} else {
		blockIOProvider := os.Getenv("BLOCK_IO_PROVIDER")
		if blockIOProvider != "" {
			for k, _ := range supportChainNetList {
				if supportChainNetList[k].ChainId == 56 {
					supportChainNetList[k].RpcUrl = fmt.Sprintf("https://bsc.getblock.io/mainnet/?api_key=%s", blockIOProvider)
				}
				if supportChainNetList[k].ChainId == 137 {
					supportChainNetList[k].RpcUrl = fmt.Sprintf("https://matic.getblock.io/mainnet/?api_key=%s", blockIOProvider)
				}
				if supportChainNetList[k].ChainId == 25 {
					supportChainNetList[k].RpcUrl = fmt.Sprintf("https://cro.getblock.io/mainnet/?api_key=%s", blockIOProvider)
				}
			}
		}
	}

	for k, _ := range supportChainNetList {
		if supportChainNetList[k].ChainId == 137 && os.Getenv("EXPLORER_API_KEY_POLYGON") != "" {
			supportChainNetList[k].ExplorerApiKey = os.Getenv("EXPLORER_API_KEY_POLYGON")
		}

		if url := os.Getenv(fmt.Sprintf("RPC_URL_%d", supportChainNetList[k].ChainId)); url != "" {
			supportChainNetList[k].RpcUrl = url
		}
	}
}

func GetSupportChainList() []*Chain {
	return supportChainList
}

func GetSupportNetList() []*Net {
	return supportChainNetList
}

func GetNetByChainId(chainId uint64) *Net {
	for _, v := range supportChainNetList {
		if v.ChainId == chainId {
			return v
		}
	}
	log.Info().Msgf("get net invalid chain id param %d", chainId)
	return &Net{ChainType: 0, ChainId: 0}
}

func GetNetByChainType(chainType uint32) *Net {
	netType := NetworkTypeMainNet
	if os.Getenv("IS_TEST_NET") == "1" {
		netType = NetworkTypeTestNet
	}

	for _, net := range supportChainNetList {
		if net.ChainType == chainType && net.Type == uint8(netType) {
			return net
		}
	}
	log.Info().Msgf("get net invalid chain type param %d", chainType)
	return &Net{ChainType: chainType}
}
