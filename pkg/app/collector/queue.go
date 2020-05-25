package collector

import (
	"github.com/astaxie/beego/logs"
	"github.com/skdong/zeus/pkg/models"
	"github.com/skdong/zeus/pkg/websocket"
)

type EntityQueue struct {
	entity  chan *models.Wind
	handler *Handler
	started bool
}

func NewEntityQueue(handler *Handler) *EntityQueue {
	q := new(EntityQueue)
	q.handler = handler
	q.entity = make(chan *models.Wind)
	q.started = false
	return q
}

func (q *EntityQueue) Start() {
	q.started = true
	logs.Info("Start Queue")
	go q.loop()
}

func (q *EntityQueue) Stop() {
	logs.Info("Stop Queue")
	q.started = false
	q.entity <- nil
}

func (q *EntityQueue) loop() {
	for q.started {
		wind := <-q.entity
		if wind != nil {
			q.handle(wind)
		}
	}
	logs.Info("Out loop")
}

func (q *EntityQueue) Add(entities ...*models.Wind) {
	for _, e := range entities {
		q.entity <- e
	}
}

func (q *EntityQueue) handle(w *models.Wind) {
	if wind, err := w.Insert(); err == nil {
		if _, ok := q.handler.Devices[wind.DeviceId]; !ok {
			q.handler.Devices[wind.DeviceId] = NewDeviceBuffer()
		}
		buffer := q.handler.Devices[wind.DeviceId]
		buffer.AddData(wind)
		websocket.WebSocketManager.BroadCast(buffer.GetData())
	} else {
		logs.Warn(err)
	}
}
