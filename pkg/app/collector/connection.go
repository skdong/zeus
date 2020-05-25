package collector

import (
	"net"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type Connection struct {
	conn    net.Conn
	handler *Handler
	buffer  []byte
}

func NewConnection(conn net.Conn) *Connection {
	c := new(Connection)
	c.conn = conn
	name := c.conn.RemoteAddr().String()
	c.handler = NewHandler(name)

	return c
}

func (c *Connection) HandlerConnection() {
	go c.handlerConnection()
}

func (c *Connection) newBuffer() {
	c.buffer = make([]byte,
		beego.AppConfig.
			DefaultInt("collector::buffer_len",
				512))
}

func (c *Connection) handlerConnection() {
	defer c.conn.Close()

	c.handler.Start()
	defer c.handler.Stop()

	for {
		c.newBuffer()
		n, err := c.conn.Read(c.buffer)
		if err != nil {
			logs.Warn(c.conn.RemoteAddr().String(), err)
			return
		}
		data := string(c.buffer)
		logs.Debug("get data", n, data)
		c.handler.HandlerData(data)
		logs.Debug("handler over")
	}
}
