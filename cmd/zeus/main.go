package main

import (
	"github.com/astaxie/beego"
	"github.com/skdong/zeus/pkg/app/collector"
	_ "github.com/skdong/zeus/pkg/routers"
)

func main() {
	collector.Run()
	beego.Run()
}
