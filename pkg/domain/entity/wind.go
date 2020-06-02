package entity

import "time"

type Wind struct {
	Id        int64  `pk:"auto;column(id)"`
	DeviceId  string `orm:"size(64);index"`
	Direction int
	Speed     float64
	Unit      string    `orm:"size(32)"`
	CreateAt  time.Time `orm:"auto_now;type(datetime);index"`
}

type WindList struct {
	Num   int
	Winds *[]interface{} `json:"winds"`
}
