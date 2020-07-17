package g

import (
	"log"
	"runtime"
)

var (
	CurrDaySync bool
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
