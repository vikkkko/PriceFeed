package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"ism/pkg/config"
	"ism/pkg/log"
	"ism/server"
	"os"
	"runtime"
	"runtime/debug"
)

func _main() error {
	cfg, err := config.Load()
	if err != nil {
		fmt.Printf("loadConfig err: %s", err)
		return err
	}
	log.InitLogFile(cfg.LogFilePath)
	logrus.Infof("config info:%+v ", cfg)

	//interrupt signal
	ctx := server.ShutdownListener()

	//server
	server, err := server.NewServer(cfg)
	if err != nil {
		logrus.Errorf("new server err: %s", err)
		return err
	}
	server.Start()

	<-ctx.Done()
	return nil
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	debug.SetGCPercent(40)
	err := _main()
	if err != nil {
		os.Exit(1)
	}
}
