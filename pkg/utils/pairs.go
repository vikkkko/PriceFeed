package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var GraphqlApiUniSwapV2MainNet = "https://api.thegraph.com/subgraphs/name/uniswap/uniswap-v2"
var GraphqlApiHecoMainNet = "https://graph.mdex.cc/subgraphs/name/mdex/swap"

var transactionGraphqlMainNetMiniV2 = "https://api.thegraph.com/subgraphs/name/noberk/chapter3"
var transactionGraphqlTestNetMiniV2 = "https://api.thegraph.com/subgraphs/name/noberk/chapter4"

const ethOneDayBlock = 5760
const hecoOneDayBlock = 5760 * 5
const ettTime = 15
const hecoTime = 3

func GetLpPrice(api, id string) (float64, error) {
	id = strings.ToLower(id)
	type rspData struct {
		Data struct {
			Pair struct {
				ID          string `json:"id"`
				ReserveUSD  string `json:"reserveUSD"`
				TotalSupply string `json:"totalSupply"`
			} `json:"pair"`
		} `json:"data"`
	}
	queryStr := `
            { 
                pair(id:"%s") {
   				 	id
    				totalSupply
    				reserveUSD
  				} 
            }
        `
	jsonData := map[string]string{
		"query": fmt.Sprintf(queryStr, id),
	}
	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		return 0, err
	}
	rsp := rspData{}
	err = getGraphqlData(api, jsonValue, &rsp)
	if err != nil {
		return 0, err
	}

	reserveUsdDeci, err := decimal.NewFromString(rsp.Data.Pair.ReserveUSD)
	if err != nil {
		return 0, err
	}
	totalSupplyDeci, err := decimal.NewFromString(rsp.Data.Pair.TotalSupply)
	if err != nil {
		return 0, err
	}
	priceDeci, _ := reserveUsdDeci.Div(totalSupplyDeci).Float64()

	return priceDeci * 2, nil
}

type PairTokens struct {
	Token0Symbol  string
	Token0Address string
	Token1Symbol  string
	Token1Address string
}

func GetPairTokens(api, id string) (*PairTokens, error) {
	id = strings.ToLower(id)
	type rspData struct {
		Data struct {
			Pair struct {
				Token0 struct {
					ID     string `json:"id"`
					Symbol string `json:"symbol"`
				} `json:"token0"`
				Token1 struct {
					ID     string `json:"id"`
					Symbol string `json:"symbol"`
				} `json:"token1"`
			} `json:"pair"`
		} `json:"data"`
	}
	queryStr := `
            { 
                pair(id:"%s") {
					token0{
      					id
      					symbol
    				}
    				token1{
    					id
      					symbol
    				}
				}
            }
        `
	jsonData := map[string]string{
		"query": fmt.Sprintf(queryStr, id),
	}
	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	rsp := rspData{}
	err = getGraphqlData(api, jsonValue, &rsp)
	if err != nil {
		return nil, err
	}

	pairTokens := PairTokens{}
	pairTokens.Token0Address = rsp.Data.Pair.Token0.ID
	pairTokens.Token0Symbol = rsp.Data.Pair.Token0.Symbol
	pairTokens.Token1Address = rsp.Data.Pair.Token1.ID
	pairTokens.Token1Symbol = rsp.Data.Pair.Token1.Symbol

	return &pairTokens, nil
}

func GetTokenPrice(api, id string) (float64, error) {
	id = strings.ToLower(id)
	type rspData struct {
		Data struct {
			TokenDayDatas []struct {
				ID       string `json:"id"`
				PriceUSD string `json:"priceUSD"`
			} `json:"tokenDayDatas"`
		} `json:"data"`
	}

	queryStr := `
            { 
                tokenDayDatas(first:1,orderBy: date, orderDirection: desc,where:{token:"%s"}) {
   				 	id
    				priceUSD
  				} 
            }
        `
	jsonData := map[string]string{
		"query": fmt.Sprintf(queryStr, id),
	}
	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		return 0, err
	}

	rsp := rspData{}
	getGraphqlData(api, jsonValue, &rsp)

	if len(rsp.Data.TokenDayDatas) == 0 {
		return 0, errors.New("tokenDayDatas len is 0")
	}
	priceUsdDeci, err := decimal.NewFromString(rsp.Data.TokenDayDatas[0].PriceUSD)
	if err != nil {
		return 0, err
	}
	price, _ := priceUsdDeci.Float64()
	return price, nil
}

func GetTokenPriceChangeOneDay(api, id string, nowBlock uint64) (float64, error) {
	id = strings.ToLower(id)
	nowBlock -= 10
	priceNow, err := GetTokenPriceOnBlock(api, id, nowBlock)
	if err != nil {
		return 0, err
	}

	time.Sleep(100 * time.Millisecond)
	priceOneDay, err := GetTokenPriceOnBlock(api, id, nowBlock-hecoOneDayBlock)
	if err != nil {
		return 0, err
	}

	if priceOneDay <= 0 {
		return 0, errors.New(fmt.Sprintf("priceOneDay err %f", priceOneDay))
	}

	return ((priceNow - priceOneDay) * 100) / priceOneDay, nil
}

type TokenPrice struct {
	TimeStamp int64
	Price     float64
}

func GetTokenPriceListOneDay(api, id string, block uint64, now int64) ([]*TokenPrice, error) {
	id = strings.ToLower(id)
	retList := make([]*TokenPrice, 0)
	for i := 23; i >= 0; i-- {
		n := 0
		for {
			price, err := GetTokenPriceOnBlock(api, id, block-uint64(i)*3600/hecoTime)
			if err != nil {
				time.Sleep(100 * time.Millisecond)
				logrus.Errorf("GetTokenPriceOnBlock err %s", err)
				n++
				if n < 5 {
					continue
				} else {
					return nil, err
				}
			} else {
				tokenPrice := TokenPrice{now - 3600*int64(i), price}
				retList = append(retList, &tokenPrice)
				break
			}
		}
	}

	return retList, nil
}

func GetTokenPriceOnBlock(api, id string, block uint64) (float64, error) {
	id = strings.ToLower(id)
	type rspData struct {
		Data struct {
			Bundles []struct {
				EthPrice string `json:"ethPrice"`
			} `json:"bundles"`
			Tokens []struct {
				DerivedETH string `json:"derivedETH"`
			} `json:"tokens"`
		} `json:"data"`
	}

	queryStr := `
			{
  				bundles(block:{number:%d}){
    				ethPrice
  				}
  				tokens(where:{id:"%s"},block:{number:%d}){
    				derivedETH
  				}
			}`
	jsonData := map[string]string{
		"query": fmt.Sprintf(queryStr, block, id, block),
	}
	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		return 0, err
	}
	rsp := rspData{}
	err = getGraphqlData(api, jsonValue, &rsp)
	if err != nil {
		return 0, err
	}

	if len(rsp.Data.Tokens) == 0 || len(rsp.Data.Bundles) == 0 {
		return 0, errors.New("rsp data len is 0")
	}
	derivedEthDecimal, err := decimal.NewFromString(rsp.Data.Tokens[0].DerivedETH)
	if err != nil {
		return 0, err
	}
	ethPriceDecimal, err := decimal.NewFromString(rsp.Data.Bundles[0].EthPrice)
	if err != nil {
		return 0, err
	}

	retPrice, _ := derivedEthDecimal.Mul(ethPriceDecimal).Float64()
	return retPrice, nil
}

func GetTokenPrice2(api, id string) (float64, error) {
	id = strings.ToLower(id)
	var queryPriceStr = `
					{	pair(id:"%s"){
							id
							token1Price
						}
					}`

	jsonData := map[string]string{
		"query": fmt.Sprintf(queryPriceStr, id),
	}
	type Rsp struct {
		Data struct {
			Pair struct {
				Id          string `json:"id"`
				Token0Price string `json:"token1Price"`
			} `json:"pair"`
		} `json:"data"`
	}

	jsonValue, _ := json.Marshal(jsonData)
	var rsp Rsp

	err := getGraphqlData(api, jsonValue, &rsp)
	if err != nil {
		return 0, err
	}
	return StrToFloat(rsp.Data.Pair.Token0Price), nil
}

//适用于其中之一是稳定币
func GetTokenPriceFromPair(api, tokenAddr, pairAddr string) (float64, error) {
	tokenAddr = strings.ToLower(tokenAddr)
	pairAddr = strings.ToLower(pairAddr)
	var queryPriceStr = `
					{	
						pair(id:"%s"){
  							id
    						token0{
      							id
      							symbol
    						}
    						token1{
      							id
      							symbol
    						}
    						token0Price
    						token1Price
  						}
					}`

	jsonData := map[string]string{
		"query": fmt.Sprintf(queryPriceStr, pairAddr),
	}
	type Rsp struct {
		Data struct {
			Pair struct {
				ID     string `json:"id"`
				Token0 struct {
					ID     string `json:"id"`
					Symbol string `json:"symbol"`
				} `json:"token0"`
				Token0Price string `json:"token0Price"`
				Token1      struct {
					ID     string `json:"id"`
					Symbol string `json:"symbol"`
				} `json:"token1"`
				Token1Price string `json:"token1Price"`
			} `json:"pair"`
		} `json:"data"`
	}

	jsonValue, _ := json.Marshal(jsonData)
	var rsp Rsp

	err := getGraphqlData(api, jsonValue, &rsp)
	if err != nil {
		return 0, err
	}
	var priceStr string
	if tokenAddr == strings.ToLower(rsp.Data.Pair.Token0.ID) {
		priceStr = rsp.Data.Pair.Token1Price
	} else {
		priceStr = rsp.Data.Pair.Token0Price
	}

	return StrToFloat(priceStr), nil
}

func getGraphqlData(api string, jsonValue []byte, retValue interface{}) error {
	request, err := http.NewRequest("POST", api, bytes.NewBuffer(jsonValue))
	client := &http.Client{Timeout: time.Second * 3}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("ret status code %d", response.StatusCode))
	}

	rspData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(rspData, retValue)
}
