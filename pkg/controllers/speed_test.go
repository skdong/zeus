package controllers

import (
	"testing"
	"time"

	"github.com/astaxie/beego"

	"github.com/skdong/zeus/pkg/conf"
	"github.com/skdong/zeus/pkg/storage"
)

func TestSpeed(t *testing.T) {
	setUp()
	period := 195
	interval := 2

	end := time.Now()
	start := end.Add(time.Duration(-1*period) * time.Hour)
	end = start.Add(time.Duration(4) * time.Hour)

	list := GetWinds(start, end)
	data := AggrateSpeed(start, end, interval, list)
	t.Log(start, end, interval)
	t.Log(data)
	t.Fatal("aa")
}

func setUp() {

	beego.AppConfig.Set("db_type", "mysql")
	beego.AppConfig.Set("db_host", "localhost")
	beego.AppConfig.Set("db_port", "3306")
	beego.AppConfig.Set("db_user", "root")
	beego.AppConfig.Set("db_pass", "root")
	beego.AppConfig.Set("db_name", "zeus")
	conf.Init()

	storage.ConnectDB()
}
