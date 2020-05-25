package collector

import (
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/skdong/zeus/pkg/parser"
)

type Handler struct {
	name    string
	Parser  *parser.Parser
	Devices DevicesBuffer
	Queue   *EntityQueue
}

func NewHandler(name string) *Handler {
	h := new(Handler)
	h.name = name
	h.Parser = parser.NewParser()
	h.Devices = make(DevicesBuffer)
	h.Queue = NewEntityQueue(h)
	return h
}

func (h *Handler) Start() {
	h.Queue.Start()
}

func (h *Handler) Stop() {
	h.Queue.Stop()
}

func (h *Handler) HandlerData(data string) error {
	logs.Debug(h.name, time.Now(), data)
	h.Parser.AddData(data)
	ws, err := h.Parser.GetAll()
	if err != nil {
		logs.Warn("get all winds", err)
	} else {
		for _, w := range ws {
			logs.Info(h.name, w)
			h.Queue.Add(w)
		}
	}
	return nil
}
