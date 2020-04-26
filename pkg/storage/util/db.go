/*
	Copyright 2017 by rabbit author: gdccmcm14@live.com.
	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at
		http://www.apache.org/licenses/LICENSE-2.0
	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License
*/
package util

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/skdong/zeus/pkg/conf"
)

func init() {
	orm.DefaultTimeLoc = time.UTC
}

func Syncdb(force bool) {
	beego.Trace("db, sync db start")

	Createdb(force)
	Connect()
	CreateConfig()

	beego.Trace("sync db end, please reopen app again")
}

func CreateConfig() {
	name := "default" // database alias name
	force := true     // drop table force
	verbose := true   // print log
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		beego.Error("database config set to force error:" + err.Error())
	}
}

//创建数据库
func Createdb(force bool) {
	beego.Trace("create database start")
	var url, createdbsql, dropdbsql string

	switch conf.DbType {
	case "mysql":
		url = fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8", conf.DbUser, conf.DbPass, conf.DbHost, conf.DbPort)
		createdbsql = fmt.Sprintf("CREATE DATABASE if not exists `%s` CHARSET utf8 COLLATE utf8_general_ci", conf.DbName)
		dropdbsql = fmt.Sprintf("DROP DATABASE IF EXISTS `%s`", conf.DbName)
		if force {
			fmt.Println(dropdbsql)
		}
		fmt.Println(createdbsql)
		break
	default:
		beego.Critical("db driver not support:", conf.DbType)
		return
	}
	db, err := sql.Open(conf.DbType, url)
	if err != nil {
		panic(err.Error())
	}
	if force {
		_, err = db.Exec(dropdbsql)
	}
	_, err1 := db.Exec(createdbsql)
	if err != nil || err1 != nil {
		beego.Error("db exec error：", err, err1)
		panic(err.Error())
	} else {
		beego.Trace("database ", conf.DbName, " created")
	}
	defer db.Close()
	beego.Trace("create database end")
}

func Connect() error {
	var url string
	switch conf.DbType {
	case "mysql":
		orm.RegisterDriver("mysql", orm.DRMySQL)
		url = conf.MYSQLURL
		break
	default:
		beego.Critical("db driver not support:", conf.DbType)
		panic("db driver not support")
	}

	beego.Trace("database start to connect", url)
	err := orm.RegisterDataBase("default", conf.DbType, url)
	if err != nil {
		beego.Error("register data:" + err.Error())
		return err
	}

	if conf.DbLog == "open" {
		beego.Trace("develop mode，debug database: db.log")
		orm.Debug = true
		w, e := os.OpenFile("log/db.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if e != nil {
			beego.Error(e.Error())
		}
		orm.DebugLog = orm.NewLog(w)
	}

	RegisterDBModel() // must register
	return nil
}
