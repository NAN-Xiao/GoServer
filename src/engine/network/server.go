package network

import (
	"fmt"
	"net"
	serverinterface "server/src/engine/interface"
)

type Server struct {
	Name string
	IPV  string
	IP   string
	Port int
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

		for {
			connect, err := listener.AcceptTCP()
			if err != nil {
				continue
			}
			go func() {
				for {
					buf := make([]byte, 512)
					cnt, err := connect.Read(buf)
					if err != nil {
						continue
					}
					if _, err := connect.Write(buf[:cnt]); err != nil {
						continue
					}
				}

			}()
		}
	}()
}

func (s *Server) Stop() {

}

func (s *Server) Serve() {
	s.Start()
	select {

	}
}

func NewServer(name string) serverinterface.IServer {
	s := &Server{
		Name: name,
		IPV:  "tcp4",
		IP:   "0.0.0.0",
		Port: 8999,
	}
	return s
}
