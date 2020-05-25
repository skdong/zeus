package collector

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/skdong/zeus/pkg/models"
	"github.com/skdong/zeus/pkg/websocket"
)

type DevicesBuffer map[string]*DeviceBuffer

type DeviceBuffer struct {
	Duration time.Duration
	Datas    []*models.Wind
}

func NewDeviceBuffer() *DeviceBuffer {
	b := new(DeviceBuffer)
	b.Duration, _ = time.ParseDuration(
		beego.AppConfig.DefaultString("duration", "5m"))
	return b
}

func (b *DeviceBuffer) AddData(d *models.Wind) {
	b.Datas = append(b.Datas, d)
}

func (b *DeviceBuffer) GetData() *websocket.Wind {
	b.clear()
	return b.getData()
}

func (b *DeviceBuffer) getData() *websocket.Wind {
	if dataNum := len(b.Datas); dataNum > 0 {
		w := websocket.NewWind(b.Datas[dataNum-1])
		var sumSpeed, sumDirection float64
		for _, d := range b.Datas {
			sumSpeed += float64(d.Speed)
			sumDirection += float64(d.Direction)

			if w.MaxDirection < d.Direction {
				w.MaxDirection = d.Direction
			}
			if w.MaxSpeed < d.Speed {
				w.MaxSpeed = d.Speed
			}
			if w.MinDirection > d.Direction {
				w.MinDirection = d.Direction
			}
			if w.MinSpeed > d.Speed {
				w.MinSpeed = d.Speed
			}
		}
		w.AvgDirection = sumDirection / float64(dataNum)
		w.AvgSpeed = sumSpeed / float64(dataNum)
		return w

	} else {
		return nil
	}
}

func (b *DeviceBuffer) clear() {
	datas := []*models.Wind{}
	now := time.Now()
	for _, d := range b.Datas {
		if !d.CreateAt.Add(b.Duration).Before(now) {
			datas = append(datas, d)
		} else {
			logs.Debug("remove ", d)
		}
	}
	b.Datas = datas
}
