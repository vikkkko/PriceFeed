package server

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
	"ism/task"
	"time"
)

func (svr *Server)  Task() {
	ticker := time.NewTicker(time.Duration(svr.taskTicker) * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			logrus.Infoln("task updatePrice start -----------")
			_,assets,prices:=task.GetTokenPrices(svr.swapApi,svr.assets)
			task.UpdateTokenPrice(svr.privateKey,svr.ethApi,common.HexToAddress(svr.contractOracleAddr),assets,prices)
			logrus.Infoln("task updatePrice end -----------")
		}
	}
}
