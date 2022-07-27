package chain_info

import (
	"fmt"
	"os"
	"time"
)

const (
	ChainTypeBsc      = 1
	ChainTypeMetis    = 2
	ChainTypePolygon  = 3
	ChainTypeKlaytn   = 4
	ChainTypeOkc      = 5
	ChainTypeEthereum = 10
)

var supportChainTypeList = []uint32{ChainTypeBsc, ChainTypeMetis, ChainTypePolygon, ChainTypeOkc}

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
	WssUrl           string `json:"wss_url"`
	BlockExplorer    string `json:"block_explorer"`
	ExplorerApiUrl   string `json:"explorer_api_url"`
	ExplorerApiKey   string `json:"explorer_api_key"`
	Type             uint8  `json:"type"`
	PollingInterval  int64  `json:"polling_interval"`
	AverageBlockTime int64  `json:"average_block_time"`
}

func (c *Net) GetPollingInterval() int64 {
	if c.PollingInterval != 0 {
		return c.PollingInterval
	}
	return int64(10 * time.Second)
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
		Symbol:       "BNB",
		CoinDecimals: 18,
		Type:         NetworkTypeMainNet,
	},
	{
		ChainType:    ChainTypeMetis,
		Symbol:       "Metis",
		CoinDecimals: 18,
		Type:         NetworkTypeMainNet,
	},
	{
		ChainType:    ChainTypePolygon,
		Symbol:       "Matic",
		CoinDecimals: 18,
		Type:         NetworkTypeMainNet,
	},
	{
		ChainType:    ChainTypeKlaytn,
		Symbol:       "KLAY",
		CoinDecimals: 18,
		Type:         NetworkTypeMainNet,
		Name:         "Klaytn",
	},
	{
		ChainType:    ChainTypeOkc,
		Symbol:       "OKC",
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
			AverageBlockTime: 10000,
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
			AverageBlockTime: 10000,
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
			PollingInterval:  750,
			AverageBlockTime: 2000,
			BlockExplorer:    "https://mumbai.polygonscan.com/",
			//WssUrl:           "wss://rpc-mumbai.matic.today",
		},
		{
			ChainType:     ChainTypeKlaytn,
			ChainId:       8217,
			Symbol:        "KLAY",
			Decimals:      18,
			Type:          NetworkTypeMainNet,
			NetworkName:   "Klaytn Cypress",
			RpcUrl:        "https://public-node-api.klaytnapi.com/v1/cypress",
			BlockExplorer: "https://scope.klaytn.com/",
		},
		{
			ChainType:     ChainTypeKlaytn,
			ChainId:       1001,
			Symbol:        "KLAY",
			Decimals:      18,
			Type:          NetworkTypeTestNet,
			NetworkName:   "Klaytn Baobab",
			RpcUrl:        "https://api.baobab.klaytn.net:8651/",
			BlockExplorer: "https://baobab.scope.klaytn.com/",
		},
		{
			ChainType:     ChainTypeOkc,
			ChainId:       66,
			Symbol:        "OKT",
			Decimals:      18,
			Type:          NetworkTypeMainNet,
			NetworkName:   "OKC Mainnet",
			RpcUrl:        "https://exchainrpc.okex.org",
			BlockExplorer: "https://www.oklink.com/okc/",
		},
		{
			ChainType:     ChainTypeOkc,
			ChainId:       65,
			Symbol:        "OKT",
			Decimals:      18,
			Type:          NetworkTypeTestNet,
			NetworkName:   "OKC Testnet",
			RpcUrl:        "https://exchaintestrpc.okex.org",
			BlockExplorer: "https://www.oklink.com/okc-test/",
		},
	}

	if os.Getenv("ENV") != "dev" {
		for k, _ := range supportChainNetList {
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
			}
		}
	}

	for k, _ := range supportChainNetList {
		if supportChainNetList[k].ChainId == 137 && os.Getenv("EXPLORER_API_KEY_POLYGON") != "" {
			supportChainNetList[k].ExplorerApiKey = os.Getenv("EXPLORER_API_KEY_POLYGON")
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
	return nil
}
