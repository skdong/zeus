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
package conf

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type FlagConfig struct {
	DbInit      *bool
	DbInitForce *bool
}

var (
	LogLevel string
)

var (
	DbType   string
	DbHost   string
	DbPort   string
	DbUser   string
	DbPass   string
	DbName   string
	DbLog    string
	MYSQLURL string
)

func Init() {
	// version
	DbType = beego.AppConfig.String("db_type")
	DbHost = beego.AppConfig.String("db_host")
	DbPort = beego.AppConfig.String("db_port")
	DbUser = beego.AppConfig.String("db_user")
	DbPass = beego.AppConfig.String("db_pass")
	DbName = beego.AppConfig.String("db_name")

	LogLevel = beego.AppConfig.String("log:log_level")

	logs.Info("DbType", ":", DbType)
	logs.Info(DbType)

	MYSQLURL = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", DbUser, DbPass, DbHost, DbPort, DbName)

}
