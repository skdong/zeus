package collector

import (
	"net"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"

	"github.com/skdong/zeus/pkg/parser"
	"github.com/skdong/zeus/pkg/websocket"
)

func Run() {
	addr := beego.AppConfig.DefaultString(
		"collector::url",
		"0.0.0.0:8081")
	go service(addr)
	logs.Info("start service", addr)
}

func service(addr string) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		logs.Warn(err)
		return
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		logs.Warn(err)
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			logs.Warn(err)
		}
		go handlerConnection(conn)

	}

}

func handlerConnection(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte,
		beego.AppConfig.DefaultInt("collector::buffer_len",
			512))
	c := NewHandler()
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			logs.Warn(conn.RemoteAddr().String(), err)
			return
		}
		data := string(buffer)
		logs.Debug("get data", n, data)
		c.HandlerData(data)
		logs.Debug("handler over")
	}
}

type Handler struct {
	Parser *parser.Parser
	Buffer []interface{}
}

func NewHandler() *Handler {
	c := new(Handler)
	c.Parser = parser.NewParser()
	return c
}

func (c *Handler) HandlerData(data string) error {
	c.Parser.AddData(data)
	ws, err := c.Parser.GetAll()
	if err != nil {
		logs.Warn("get all winds", err)
	}
	for _, w := range ws {
		logs.Debug(w.ToString())
		if wind, err := w.Insert(); err == nil {
			logs.Info(wind)
			websocket.WebSocketManager.BroadCast(w)
		} else {
			logs.Warn(err)
		}
	}
	return nil
}
