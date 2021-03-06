package network

import (
	"errors"
	"fmt"
	"net"
	conf2 "server/src/engine/conf"
	"server/src/engine/coreface"
)

type Server struct {
	Name   string
	IPV    string
	IP     string
	Port   int
	Router coreface.IRouter
}

func CallBack(conn *net.TCPConn, data []byte, len int) error {

	if _, err := conn.Write(data[:len]); err != nil {

		return errors.New("connection write error")
	}
	return nil
}

func (s *Server) Start() {
	go func() {
		addr, err := net.ResolveTCPAddr(s.IPV, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			return
		}
		listener, err := net.ListenTCP(s.IPV, addr)
		if err != nil {
			return
		}
		var cid uint32
		cid = 0
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				continue
			}
			connect := NewConnection(conn, cid, s.Router)
			cid++
			go connect.Start()
		}
	}()
}

func (s *Server) Stop() {

}

func (s *Server) Serve() {
	s.Start()
	select {}
}
func (s *Server) AddRouter(router coreface.IRouter) {
	s.Router = router
}

func NewServer() coreface.IServer {
	s := &Server{
		Name:   conf2.G_serverConfig.NAME,
		IPV:    "tcp4",
		IP:     conf2.G_serverConfig.HOST,
		Port:   conf2.G_serverConfig.TCP_PORT,
		Router: nil,
	}
	return s
}
