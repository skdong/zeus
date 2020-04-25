package main

import (
	"flag"

	"github.com/skdong/zeus/pkg/conf"
	"github.com/skdong/zeus/pkg/storage"
)

func init() {
	flags := conf.FlagConfig{}

	// db init or rebuild
	flags.DbInit = flag.Bool("db", false, "init db")
	flags.DbInitForce = flag.Bool("f", false, "force init db first drop db then rebuild it")

	flag.Parse()

	conf.Init()
	storage.Init(flags)

}
