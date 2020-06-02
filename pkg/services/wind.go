package services

import (
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/skdong/zeus/pkg/models"
)

type WindList struct {
	Num   int
	Winds *[]orm.Params `json:"winds"`
}

type Wind struct {
}

func (s *Wind) List(deviceId string,
	startAt, endAt *time.Time,
	limit, start int) (list WindList, err error) {

	w := new(models.Wind)
	q := w.Query()
	winds := []orm.Params{}
	if len(deviceId) != 0 {
		q = q.Filter("DeviceId", deviceId)
	}
	if startAt != nil {
		logs.Info("startAt:", startAt)
		q = q.Filter("CreateAt__gte", *startAt)
	}
	if endAt != nil {
		logs.Info("endAt:", endAt)
		q = q.Filter("CreateAt__lte", *endAt)
	}
	q = q.OrderBy("CreateAt")

	if limit != -1 {
		q = q.Limit(limit, start)
	}
	q.Values(&winds)
	list.Num = len(winds)
	list.Winds = &winds
	return list, nil
}
