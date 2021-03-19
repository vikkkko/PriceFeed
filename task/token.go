package task

import (
	"github.com/ethereum/go-ethereum/common"
	"ism/contract/Oracle"
	"ism/pkg/utils"
	"math/big"
	"strings"
)

const wht = "0x5545153ccfca01fbd7dd11c0b23ba694d9509a6f"

func GetTokenPrices(swapApi string, assets []string) (error,[]common.Address,[]*big.Int) {
	prices := make([]*big.Int,0,0)
	_assets := make([]common.Address,0,0)
	for _,asset:=range assets{
		//判断是否为基础链代币地址
		if strings.ToLower(asset) == "0x0000000000000000000000000000000000000000" {
			asset = wht
		}
		//从 mdex 获取 claim 代币的价格,尝试两种方式
		price, err := utils.GetTokenPrice(swapApi, asset)
		if err != nil {
			panic(err)
		}
		prices = append(prices, big.NewInt(int64(price)))
		_assets = append(_assets, common.HexToAddress(asset))
	}
	return nil,_assets,prices
}

//更新token价格
func UpdateTokenPrice(hexKey,ethApi string,contractAddr common.Address ,assets []common.Address, prices []*big.Int) error {
	//
	// 在这里处理一些价格逻辑
	//
	Oracle.SetDirectPrice(hexKey,ethApi,contractAddr,assets,prices)
	return nil
}
