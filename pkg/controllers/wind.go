package controllers

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/skdong/zeus/pkg/services"
)

type WindController struct {
	beego.Controller
}

func (c *WindController) Get() {
	var startTime, endTime *time.Time
	deviceId := c.GetString("device_id", "Q")
	startAt := c.GetString("start_at", "")
	endAt := c.GetString("end_at", "")

	if tmp, err := time.ParseInLocation("2006-01-02 15:04:05",
		endAt,
		time.Local); err == nil {
		endTime = &tmp
	}
	if tmp, err := time.ParseInLocation("2006-01-02 15:04:05",
		startAt,
		time.Local); err == nil {
		startTime = &tmp
	}

	start, err := c.GetInt("start", 0)
	if err != nil {
		logs.Warn("start is not valid", start)
		return
	}

	limit, err := c.GetInt("limit", 1000)
	if err != nil {
		logs.Warn(err)
		return
	}
	s := services.Wind{}
	winds, err := s.List(
		deviceId,
		startTime, endTime,
		limit, start)

	if err != nil {
		logs.Warn(err)
		return
	}
	c.Data["json"] = winds
	c.ServeJSON()
}
