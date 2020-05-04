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
package storage

import (
	"flag"
	"os"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/skdong/zeus/pkg/conf"
	"github.com/skdong/zeus/pkg/storage/util"
)

func ConnectDB() {

	for util.Connect() != nil {
		logs.Warn("Connect DB error !!!")
		logs.Info("retry 1 sencond")
		time.Sleep(time.Second)
	}
}

func Init(config conf.FlagConfig) {
	beego.Trace("database start to run")
	initDb(config)
	//util.Connect()
	go ConnectDB()
}

func initDb(config conf.FlagConfig) {
	if !flag.Parsed() {
		flag.Parse()
	}
	if *config.DbInit {

		util.Syncdb(*config.DbInitForce)
		os.Exit(0)
	}
}
