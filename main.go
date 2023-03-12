package main

import (
	"demo/kernel"
	"demo/services/config"
	"demo/services/mdb"
	"demo/services/zlog"
)

func main() {
	if !config.Init() {
		return
	}
	zlog.Start()
	defer zlog.Stop()
	zlog.Warn("program", "start ...")
	defer zlog.Warn("program", "stop ...")
	mdb.Start()
	defer mdb.Stop()
	kernel.Run()
}
