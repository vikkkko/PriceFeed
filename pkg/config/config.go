package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	PrivateKey			   string
	QueueTicker            int64 // queue duration 毫秒
	TaskTicker             int64 // task duration
	EthApi                 string
	SwapApi                string
	ContractOracleAddr     string
	ListenAddr             string
	LogFilePath            string
	Assets				   []string
}

func Load() (*Config, error) {
	configFilePath := flag.String("C", "conf.toml", "Config file path")
	flag.Parse()

	var cfg = Config{}
	if err := loadSysConfig(*configFilePath, &cfg); err != nil {
		return nil, err
	}
	if cfg.LogFilePath == "" {
		cfg.LogFilePath = "./log_data"
	}
	return &cfg, nil
}

func loadSysConfig(path string, config *Config) error {
	_, err := os.Open(path)
	if err != nil {
		return err
	}
	if _, err := toml.DecodeFile(path, config); err != nil {
		return err
	}
	fmt.Println("load sysConfig success")
	return nil
}
