package services

import (
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/skdong/zeus/pkg/models"
)

type WindList struct {
	Num          int
	SpeedAvg     float32
	DirectionAvg float32
	Winds        *[]orm.Params `json:"winds"`
}

type Wind struct {
}

func (s *Wind) List(deviceId string,
	startAt, endAt *time.Time,
	limt, start int) (list WindList, err error) {

	w := new(models.Wind)
	q := w.Query()
	winds := []orm.Params{}
	q = q.Filter("DeviceId", deviceId)
	if startAt != nil {
		logs.Info("start:", startAt)
		q = q.Filter("CreateAt__gte", *startAt)
	}
	if endAt != nil {
		logs.Info("end:", endAt)
		q = q.Filter("CreateAt__lte", *endAt)
	}
	q.OrderBy("-CreateAt").Limit(limt, start).Values(&winds)
	var speedSum, directionSum float64
	for _, w := range winds {
		if w["Speed"] == nil || w["Direction"] == nil {
			logs.Warn("data not valid", w)
			continue
		}
		speedSum += float64(w["Speed"].(float64))
		directionSum += float64(w["Direction"].(int64))
	}
	list.Num = len(winds)
	list.SpeedAvg = float32(speedSum / float64(list.Num))
	list.DirectionAvg = float32(directionSum / float64(list.Num))
	list.Winds = &winds
	return list, nil
}
