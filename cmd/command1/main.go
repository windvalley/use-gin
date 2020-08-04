package main

import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/spf13/pflag"

	"use-gin/config"
	"use-gin/logger"
	"use-gin/utils"
)

func main() {
	cfg := pflag.StringP("config", "c", "", "Specify your configuration file")
	pflag.Parse()
	if *cfg == "" {
		binName := filepath.Base(os.Args[0])
		fmt.Printf("missing parameter\nUsage of %s:\n  -c, --config string"+
			"   Specify your configuration file\n", binName)
		os.Exit(2)
	}

	config.ParseConfig(*cfg)
	logger.InitCMDLog()

	// process lock
	lock, lockFile, err := utils.ProcessLock()
	if err != nil {
		logger.Log.Fatal(err)
	}
	defer os.Remove(lockFile)
	defer lock.Close()

	//ctx, cancel := context.WithCancel(context.Background())
	//go xxx(ctx)

	sigC := make(chan os.Signal)
	signal.Notify(sigC)
	for sig := range sigC {
		switch sig {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT,
			syscall.SIGKILL:
			logger.Log.Infof("%v signal captured, quit.", sig)
			//cancel()
			os.Remove(lockFile)
			os.Exit(1)
		}
	}
}
