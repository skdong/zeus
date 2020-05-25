package websocket

import (
	"github.com/skdong/zeus/pkg/models"
)

type Wind struct {
	models.Wind
	MaxSpeed     float64 `json:"maxSpeed"`
	MinSpeed     float64 `json:"minSpeed"`
	AvgSpeed     float64 `json:"avgSpeed"`
	MaxDirection int     `json:"maxDirection"`
	MinDirection int     `json:"minDirection"`
	AvgDirection float64 `json:"avgDirection"`
}

func NewWind(w *models.Wind) *Wind {
	d := new(Wind)
	d.Id = w.Id
	d.DeviceId = w.DeviceId
	d.CreateAt = w.CreateAt
	d.Speed = w.Speed
	d.MaxSpeed = w.Speed
	d.MinSpeed = w.Speed
	d.Direction = w.Direction
	d.MaxDirection = w.Direction
	d.MinDirection = w.Direction
	d.Unit = w.Unit
	return d
}
