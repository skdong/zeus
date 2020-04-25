package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

var TABLENAME = "winds"

type Wind struct {
	Id        int64  `pk:"auto;column(id)"`
	DeviceId  string `orm:"size(64)"`
	Direction int
	Speed     float64
	Unit      string    `orm:"size(32)"`
	CreateAt  time.Time `orm:"auto_now;type(datetime)"`
}

func NewWind(deviceId string, direction int, speed float64, unit string) *Wind {
	w := &Wind{
		DeviceId:  deviceId,
		Direction: direction,
		Speed:     speed,
		Unit:      unit,
	}
	return w
}

func (w *Wind) TableName() string {
	return TABLENAME
}

func (w *Wind) Insert() error {
	if _, e := orm.NewOrm().Insert(w); e != nil {
		return e
	}
	return nil
}

func (w *Wind) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(w)
}

func (w *Wind) Equal(o *Wind) bool {
	if w.Id != o.Id ||
		w.Direction != o.Direction ||
		w.Speed != o.Speed ||
		w.Unit != o.Unit {
		return false
	}
	return true
}

func (w *Wind) ToString() string {
	return fmt.Sprintf("DeviceId: %v,Direction: %v,Speed:%v,Unit:%v,CreateAt:%v",
		w.DeviceId, w.Direction,
		w.Speed, w.Unit, w.CreateAt)
}
