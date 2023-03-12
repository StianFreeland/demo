package kernel

import (
	"demo/routers"
	"demo/services/zlog"
	"net/http"
	"strconv"
	"sync"
)

const moduleName = "kernel"

var wg *sync.WaitGroup
var srv *http.Server

func Run() {
	zlog.Warn(moduleName, "run begin")
	loadConfig()
	wg = &sync.WaitGroup{}
	addr := ":" + strconv.FormatInt(port, 10)
	srv = &http.Server{Addr: addr, Handler: routers.GetEngine()}
	wg.Add(1)
	go handleServer()
	wg.Add(1)
	go handleSignal()
	wg.Wait()
	zlog.Warn(moduleName, "run end")
}
