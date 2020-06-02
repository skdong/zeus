package controllers

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/skdong/zeus/pkg/domain/entity"
	"github.com/skdong/zeus/pkg/services"
)

var LIMIT = -1

type SpeedController struct {
	beego.Controller
}

func GetWinds(deviceId string, startAt, endAt time.Time) services.WindList {

	start := 0

	s := services.Wind{}
	winds, err := s.List(
		deviceId,
		&startAt,
		&endAt,
		LIMIT,
		start)

	if err != nil {
		logs.Warn(err)
		return services.WindList{}
	}
	return winds
}

func CreateSpeeds(start, end time.Time, interval int) entity.Speeds {
	period := end.Sub(start).Minutes()
	speedsNum := int(period / float64(interval))
	if int(period)%interval != 0 {
		speedsNum++
	}
	speeds := make(entity.Speeds, speedsNum)
	return speeds
}

func addSpeed(speeds entity.Speeds, start time.Time, interval int, device map[string]interface{}) {
	createAt := device["CreateAt"].(time.Time)
	speed := device["Speed"].(float64)
	i := int(createAt.Sub(start).Minutes() / float64(interval))
	if speed > speeds[i].Speed || start.After(speeds[i].CreateAt) {
		speeds[i].CreateAt = createAt
		speeds[i].Speed = speed
	}
}

func trimSpeeds(speeds entity.Speeds, start time.Time) entity.Speeds {
	ss := entity.Speeds{}
	for _, s := range speeds {
		if !start.After(s.CreateAt) {
			ss = append(ss, s)
		}
	}
	return ss
}

func AggrateSpeed(start, end time.Time, interval int, winds services.WindList) entity.DevicesSpeeds {
	devicesSpeeds := entity.DevicesSpeeds{}
	for _, d := range *winds.Winds {
		deviceId := d["DeviceId"].(string)
		if speeds, ok := devicesSpeeds[deviceId]; ok {
			addSpeed(speeds, start, interval, d)
		} else {
			speeds = CreateSpeeds(start, end, interval)
			devicesSpeeds[deviceId] = speeds
			addSpeed(speeds, start, interval, d)
		}
	}
	for k, v := range devicesSpeeds {
		devicesSpeeds[k] = trimSpeeds(v, start)
	}
	return devicesSpeeds
}

func (c *SpeedController) Get() {
	interval, _ := c.GetInt("interval", 2)
	period, _ := c.GetInt("period", 3)

	deviceId := c.GetString("device_id", "")

	end := time.Now()
	start := end.Add(time.Duration(-1*period) * time.Hour)

	startAt := c.GetString("start_at", "")
	endAt := c.GetString("end_at", "")
	if tmp, err := time.ParseInLocation("2006-01-02 15:04:05",
		endAt,
		time.Local); err == nil {
		end = tmp
	}
	if tmp, err := time.ParseInLocation("2006-01-02 15:04:05",
		startAt,
		time.Local); err == nil {
		start = tmp
	}

	winds := GetWinds(deviceId, start, end)
	c.Data["json"] = AggrateSpeed(start, end, interval, winds)
	c.ServeJSON()
}
