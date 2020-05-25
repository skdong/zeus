package collector

import (
	"net"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func Run() {
	addr := beego.AppConfig.DefaultString(
		"collector::url",
		"0.0.0.0:8081")
	go service(addr)
	logs.Info("start service", addr)
}

func service(addr string) {
	conns := []*Connection{}
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
		c := NewConnection(conn)
		c.HandlerConnection()
		conns = append(conns, c)
		logs.Info("Connections", len(conns))

	}

}
