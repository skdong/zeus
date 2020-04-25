package websocket

import (
	"net/http"

	"github.com/astaxie/beego/logs"
	"github.com/gorilla/websocket"
	"github.com/skdong/zeus/pkg/models"
	"github.com/skdong/zeus/pkg/services"
)

var WebSocketManager *Manager

type Request struct {
	Limit   int      `json:"limit"`
	Devices []string `json:"devices"`
	Enable  bool     `json:"enable"`
	Close   bool     `json:"close"`
}

type Handler struct {
	manager *Manager
	Conn    *websocket.Conn
	Devices []string
	Enable  bool
}

func (h *Handler) Close() {
	h.Conn.Close()
	h.manager.RemoveHandler(h)
}

func (h *Handler) HandlerMessage() {
	var req Request
	for {
		if err := h.Conn.ReadJSON(&req); err != nil {
			logs.Warn(h.Conn, "error")
			h.Close()
			break
		}
		if req.Close {
			h.Close()
			break
		}
		h.handlerRequest(&req)
	}
	logs.Info(h.Conn, "overed")
}

func (h *Handler) handlerRequest(req *Request) {
	if req.Enable {
		h.Devices = req.Devices
		if req.Limit > 0 {
			h.responseEntities(req.Limit)
		}
	} else {
		h.Devices = []string{}
	}
}

func (h *Handler) responseEntities(limit int) {
	s := services.Wind{}
	for _, id := range h.Devices {
		list, err := s.List(id, nil, nil, limit, 0)
		if err != nil {
			logs.Warn(err)
			return
		}
		h.Conn.WriteJSON(list)
	}
}

func (h *Handler) SendEntity(wind *models.Wind) {
	for _, id := range h.Devices {
		if id == wind.DeviceId {
			h.Conn.WriteJSON(wind)
		}
	}
}

type Manager struct {
	Handlers map[*Handler]bool
	Cast     chan *models.Wind
	upgrader websocket.Upgrader
}

func NewManager() *Manager {
	m := new(Manager)
	m.Handlers = make(map[*Handler]bool)
	m.Cast = make(chan *models.Wind)
	m.upgrader = websocket.Upgrader{}
	return m
}

func (m *Manager) NewWebSocketHander(
	w http.ResponseWriter,
	r *http.Request,
	header http.Header,
) (*Handler, error) {

	conn, error := m.upgrader.Upgrade(w, r, header)
	handler := &Handler{
		Conn:    conn,
		manager: m}
	m.Handlers[handler] = true
	return handler, error
}

func (m *Manager) RemoveHandler(handler *Handler) {
	delete(m.Handlers, handler)
}

func (m *Manager) BroadCast(wind *models.Wind) {
	m.Cast <- wind
}

func (m *Manager) HandleCast(wind *models.Wind) {
	logs.Info("Websocket Handler:", len(m.Handlers))
	for handler := range m.Handlers {
		handler.SendEntity(wind)
	}
}

func (m *Manager) run() {
	for {
		e := <-m.Cast
		m.HandleCast(e)
	}
}

func (m *Manager) Run() {
	go m.run()
}

func init() {
	WebSocketManager = NewManager()
	WebSocketManager.Run()
}
