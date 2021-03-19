package task_test

import (
	"ism/task"
	"testing"
	"ism/pkg/utils"
)

func TestGetTokenPrice(t *testing.T) {
	addrs := []string{"0x0000000000000000000000000000000000000000"}
	_,assets,prices := task.GetTokenPrices( utils.GraphqlApiHecoMainNet,addrs)
	t.Log(assets)
	t.Log(prices)
}