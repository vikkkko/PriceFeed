package utils_test

import (
	"ism/pkg/utils"
	"strconv"
	"testing"
	"time"
)

func TestGetSwapHash(t *testing.T) {
	timeNow := time.Now().UnixNano()
	t.Log(timeNow)
	t.Log(strconv.FormatInt(timeNow, 10))
	t.Log(utils.GetSwapHash("swap", "swap.Sender", time.Now().Unix()))
}

func TestGetTokenPrice(t *testing.T) {
	p, err := utils.GetTokenPrice(utils.GraphqlApiHecoMainNet, "0x348CCc5A616aBAE8A639457FC469917B03d938c3")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(p)
}

func TestGetTokenPrice2(t *testing.T) {
	p, err := utils.GetTokenPrice2(utils.GraphqlApiHecoMainNet, "0xd625416c4932854de0ce29de6b25bbab852d86c6")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(p)
}

func TestGetTokenPriceOnBlock(t *testing.T) {
	price, err := utils.GetTokenPriceOnBlock(utils.GraphqlApiHecoMainNet, "0xDF99182976fb14cE5ebe06e3DEF538338ffdDCCf", 1734537)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(price)
}

func TestGetTokenPriceChangeOneDay(t *testing.T) {
	price, err := utils.GetTokenPriceChangeOneDay(utils.GraphqlApiHecoMainNet, "0x64ff637fb478863b7468bc97d30a5bf3a428a1fd", 2419036)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(price)
}
func TestGetTokenPriceListOneDay(t *testing.T) {
	priceList, err := utils.GetTokenPriceListOneDay(utils.GraphqlApiHecoMainNet, "0x64ff637fb478863b7468bc97d30a5bf3a428a1fd", 1705630, time.Now().Unix())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(priceList)
}

func TestGetPairTokens(t *testing.T) {
	pairs, err := utils.GetPairTokens(utils.GraphqlApiHecoMainNet, "0x78c90d3f8a64474982417cdb490e840c01e516d4")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(pairs)
}
func TestGetTokenPriceFromPair(t *testing.T) {
	price, err := utils.GetTokenPriceFromPair(utils.GraphqlApiHecoMainNet, "0x348CCc5A616aBAE8A639457FC469917B03d938c3", "0x55136fdf4e9a3b13b4c0bc4218484f722b31bc60")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(price)
}

func TestGetPriceAndPercentFromCoinw(t *testing.T) {
	price, percent, err := utils.GetPriceAndPercentFromCoinw("aaa")
	t.Log(price, percent, err)
}

func TestGetLpPrice(t *testing.T) {
	p,err:=utils.GetLpPrice(utils.GraphqlApiHecoMainNet,"0xc6fce394010713c351a467325171a335087fc60b")
	if err!=nil{
		t.Fatal(err)
	}
	t.Log(p)
}
