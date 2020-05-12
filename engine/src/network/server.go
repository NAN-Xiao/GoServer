package network

import (
	"errors"
	"fmt"
	"net"
	"server/engine/src/coreface"
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

func NewServer(name string) coreface.IServer {
	s := &Server{
		Name:   name,
		IPV:    "tcp4",
		IP:     "0.0.0.0",
		Port:   8999,
		Router: nil,
	}
	return s
}
