package main

import (
	"flag"

	"github.com/astaxie/beego/logs"
	"github.com/skdong/zeus/pkg/conf"
	"github.com/skdong/zeus/pkg/storage"
)

var LogLevelMaps = map[string]int{
	"INFO":  logs.LevelInfo,
	"DEBUG": logs.LevelDebug,
	"ERROR": logs.LevelError,
}

func InitLog() {
	if level, ok := LogLevelMaps[conf.LogLevel]; ok {
		logs.SetLevel(level)
	} else {

		logs.SetLevel(logs.LevelInfo)
	}
}

func init() {
	flags := conf.FlagConfig{}

	// db init or rebuild
	flags.DbInit = flag.Bool("db", false, "init db")
	flags.DbInitForce = flag.Bool("f", false, "force init db first drop db then rebuild it")

	flag.Parse()
	conf.Init()

	InitLog()
	storage.Init(flags)

}
