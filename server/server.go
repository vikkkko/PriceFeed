package server

import (
	"ism/pkg/config"
	"ism/pkg/utils"
)

type Server struct {
	privateKey  string
	listenAddr  string
	queueTicker int64
	taskTicker  int64

	ethApi                 string
	swapApi                string
	assets				   []string
	contractOracleAddr 	   string
}

func NewServer(cfg *config.Config) (*Server, error) {
	s := &Server{
		privateKey:				cfg.PrivateKey,
		listenAddr:             cfg.ListenAddr,
		queueTicker:            cfg.QueueTicker,
		taskTicker:             cfg.TaskTicker,
		ethApi:                 cfg.EthApi,
		swapApi:                cfg.SwapApi,
		assets:					cfg.Assets,
		contractOracleAddr:		cfg.ContractOracleAddr,
	}
	return s, nil
}

func (svr *Server) Start() {
	utils.SafeGoWithRestart(svr.Task)
}
