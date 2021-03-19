package Oracle_test

import (
	"github.com/ethereum/go-ethereum/common"
	"ism/contract/Oracle"
	"math/big"
	"testing"
)

func TestSetDirectPrice(t *testing.T)  {
	addrs:=[]common.Address{common.HexToAddress("0x7ed03ad065fE473585DbC6028AcbB6b8DA8B4C15")}
	prices:= []*big.Int{big.NewInt(22222)}
	t.Log(Oracle.SetDirectPrice(
		"8d65d87a2de0fc95f860352250f73240baf413e21dac5e21e00b9e5a3dff16d5",
		"https://http-testnet.hecochain.com",
		common.HexToAddress("0xe55B57b9fbC7C09199b4B4234D7E77B8A0c82613"),
		addrs,
		prices,
		))
}