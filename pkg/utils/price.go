package utils

import (
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"io/ioutil"
	"net/http"
)

var hpyPriceUrl = "http://47.88.158.38//hpymarket/api/all_coin_info?coins=%s"

func GetPriceAndPercentFromCoinw(coins string) (float64, float64, error) {
	resp, err := http.Get(fmt.Sprintf(hpyPriceUrl, coins))
	if err != nil {
		return 0, 0, fmt.Errorf("get hpyPriceUrl err %s", err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		return 0, 0, fmt.Errorf("get price Url statue %d", resp.StatusCode)
	}

	bts, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, 0, fmt.Errorf("ioutil.ReadAll err %s", err.Error())
	}
	rspPrice := RspPrice{}
	json.Unmarshal(bts, &rspPrice)

	if rspPrice.Status != "200" || len(rspPrice.Data) == 0 {
		return 0, 0, fmt.Errorf("get hpyPriceUrl status %s,data len %d", rspPrice.Status, len(rspPrice.Data))
	}
	var priceFloat float64
	switch rspPrice.Data[0].Price.(type) {
	case string:
		priceDeci, err := decimal.NewFromString(rspPrice.Data[0].Price.(string))
		if err != nil {
			return 0, 0, err
		}
		priceFloat, _ = priceDeci.Float64()
	case float64:
		priceFloat = rspPrice.Data[0].Price.(float64)
	}

	percent := rspPrice.Data[0].Percent * 100

	return priceFloat, percent, nil
}

type RspPrice struct {
	Status string         `json:"status"`
	Data   []RspPriceData `json:"data"`
	Msg    string         `json:"msg"`
}

type RspPriceData struct {
	ID        string      `json:"id"`
	Coin      string      `json:"coin"`
	Cn        string      `json:"cn"`
	En        string      `json:"en"`
	Price     interface{} `json:"price"`
	T         string      `json:"t"`
	Volume24  string      `json:"volume24"`
	Logo      string      `json:"logo"`
	Circulate string      `json:"circulate"`
	Supply    string      `json:"supply"`
	Percent   float64     `json:"percent"`
	Max24     string      `json:"max24"`
	Min24     string      `json:"min24"`
}
