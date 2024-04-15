package cpfp

import (
	"encoding/json"
	"fmt"
	"github.com/tristan-club/kit/log"
	"io/ioutil"
	"math"
	"net/http"
)

type Ancestor struct {
	Txid   string `json:"txid"`
	Fee    int    `json:"fee"`
	Weight int    `json:"weight"`
}

type ResponseData struct {
	Error                string               `json:"error"`
	Ancestors            []*TransactionDetail `json:"ancestors"`
	Descendants          []Ancestor           `json:"descendants"` // Assuming descendants have the same structure
	EffectiveFeePerVsize float64              `json:"effectiveFeePerVsize"`
	AdjustedVsize        float64              `json:"adjustedVsize"`
}

func getCpfpHistory(txId string) (*ResponseData, error) {

	endpoint := fmt.Sprintf("https://mempool.space/api/v1/cpfp/%s", txId)

	resp, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data ResponseData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

type TransactionInput struct {
	Txid string `json:"txid"`
	Vout int    `json:"vout"`
}

type TransactionDetail struct {
	Txid   string             `json:"txid"`
	Vin    []TransactionInput `json:"vin"`
	Weight int                `json:"weight"`
	Fee    int                `json:"fee"`
	Status Status             `json:"status"`
}

type Status struct {
	Confirmed bool `json:"confirmed"`
}

func getTransactionDetails(txid string) (*TransactionDetail, error) {
	url := fmt.Sprintf("https://mempool.space/api/tx/%s", txid)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var details TransactionDetail
	if err := json.NewDecoder(resp.Body).Decode(&details); err != nil {
		return nil, err
	}

	return &details, nil
}

// vinGetCpfpHistory recursively fetches all unconfirmed ancestors of a given transaction
func vinGetCpfpHistory(txId string, visited map[string]bool) ([]*TransactionDetail, error) {
	if visited == nil {
		visited = make(map[string]bool)
	}

	// Check if we have already visited this transaction to avoid cycles
	if visited[txId] {
		return nil, nil
	}

	tx, err := getTransactionDetails(txId)
	if err != nil {
		return nil, fmt.Errorf("error getting transaction details: %s", err)
	}

	visited[txId] = true

	var unconfirmedAncestors []*TransactionDetail
	if !tx.Status.Confirmed {
		unconfirmedAncestors = append(unconfirmedAncestors, tx)

		// Recursively fetch all unconfirmed ancestors of each input
		for _, vin := range tx.Vin {
			parentAncestors, err := vinGetCpfpHistory(vin.Txid, visited)
			if err != nil {
				log.Error().Msgf("Failed to fetch ancestors of transaction %s: %s", vin.Txid, err)
				continue
			}
			unconfirmedAncestors = append(unconfirmedAncestors, parentAncestors...)
		}
	}

	return unconfirmedAncestors, nil
}

type CalCpfpData struct {
	DecentTxId     string
	VinFeeRate     float64
	MemPoolFeeRate float64
}

// CalCpfp 计算CPFP
// 必须参数 txID, desiredFeeRate 分别为目标txId以及目标的综合gas费
// 可选参数 childTxSizeVb 下一个交易的预估vb(虚拟大小=weight/4)
// 当 childTxSizeVb=0 时，此时操作视为rbf替换txID 对应的交易，所以计算的是当前txId对应的交易的新的feeRate
// 当 childTxSizeVb!=0时，此时操作视为为一个新的交易
func CalCpfp(txID string, childTxSizeVb float64, desiredFeeRate float64) (*CalCpfpData, error) {

	// 获取交易详情
	details, err := getTransactionDetails(txID)
	if err != nil {
		return nil, fmt.Errorf("error getting transaction details: %s", err)
	}

	log.Info().Fields(map[string]interface{}{"action": "get tx details", "data": details, "txId": txID}).Send()

	if details.Status.Confirmed {
		return nil, fmt.Errorf("the transaction is already confirmed. ")
	}

	memPoolCpfpData, err := getCpfpHistory(txID)
	if err != nil {
		fmt.Println("Error calling API:", err)
		return nil, fmt.Errorf("get cpfp history error: %s", err.Error())
	} else if memPoolCpfpData.Error != "" {
		return nil, fmt.Errorf("invalid cpfp response: %s", memPoolCpfpData.Error)
	}

	log.Info().Fields(map[string]interface{}{"action": "get cpfp history", "data": memPoolCpfpData}).Send()

	//if len(memPoolCpfpData.Ancestors) == 0 && len(memPoolCpfpData.Descendants) == 0 {
	//	return nil, fmt.Errorf("The transaction has already been confirmed or has no unconfirmed ancestors. ")
	//}

	var decentTxId string
	if len(memPoolCpfpData.Descendants) > 0 {
		log.Info().Fields(map[string]interface{}{"action": "tx has decent", "tx": memPoolCpfpData.Descendants}).Send()
		for _, d := range memPoolCpfpData.Descendants {
			decentTxId = d.Txid
		}
	}

	var vinFeeRate float64
	_vinCpfpData, err := vinGetCpfpHistory(txID, nil)
	if err != nil {
		log.Error().Fields(map[string]interface{}{"action": "get cpfp from vin error", "error": err.Error()}).Send()
	} else {

		// 创建一个新的切片来存储除了txID之外的所有交易
		var vinCpfpData []*TransactionDetail
		for _, tx := range _vinCpfpData {
			if tx.Txid != txID {
				vinCpfpData = append(vinCpfpData, tx)
			}
		}

		if len(vinCpfpData) == 0 && childTxSizeVb == 0 {
			return nil, fmt.Errorf("The transaction has already been confirmed or has no unconfirmed ancestors. ")
		}

		vinFeeRate = calCpfp(details, vinCpfpData, childTxSizeVb, desiredFeeRate)
	}

	memPoolFeeRate := calCpfp(details, memPoolCpfpData.Ancestors, childTxSizeVb, desiredFeeRate)

	log.Info().Msgf("Recommended fee rate for new transaction vin %v mempool %v", vinFeeRate, memPoolFeeRate)

	resp := &CalCpfpData{
		DecentTxId:     decentTxId,
		VinFeeRate:     vinFeeRate,
		MemPoolFeeRate: memPoolFeeRate,
	}

	//// Calculate recommended fee rate for new transaction D
	//var accumulativeFees int
	//var accumulativeSize float64
	//
	//if childTxSizeVb != 0 {
	//	accumulativeFees = details.Fee
	//	accumulativeSize = childTxSizeVb + float64(details.Weight)/4
	//} else {
	//	childTxSizeVb = float64(details.Weight) / 4
	//	accumulativeSize = childTxSizeVb
	//}
	//
	//for _, ancestor := range data.Ancestors {
	//	accumulativeFees += ancestor.Fee
	//	accumulativeSize += float64(ancestor.Weight) / 4
	//}
	//additionalFeeNeeded := desiredFeeRate*accumulativeSize - float64(accumulativeFees)
	//// 向上取整再加2让feeRate稍大于计算值，作为冗余
	//recommendedFeeRateForD := math.Ceil(additionalFeeNeeded/childTxSizeVb) + 2

	return resp, nil
}

func calCpfp(details *TransactionDetail, data []*TransactionDetail, childTxSizeVb float64, desiredFeeRate float64) float64 {
	var accumulativeFees int
	var accumulativeSize float64

	if childTxSizeVb != 0 {
		accumulativeFees = details.Fee
		accumulativeSize = childTxSizeVb + float64(details.Weight)/4
	} else {
		childTxSizeVb = float64(details.Weight) / 4
		accumulativeSize = childTxSizeVb
	}

	for _, ancestor := range data {
		accumulativeFees += ancestor.Fee
		accumulativeSize += float64(ancestor.Weight) / 4
	}
	additionalFeeNeeded := desiredFeeRate*accumulativeSize - float64(accumulativeFees)
	// 向上取整再加2让feeRate稍大于计算值，作为冗余
	recommendedFeeRateForD := math.Ceil(additionalFeeNeeded/childTxSizeVb) + 2

	return recommendedFeeRateForD
}
