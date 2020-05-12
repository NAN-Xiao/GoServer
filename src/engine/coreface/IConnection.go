package coreface

import "net"

type IConnection interface {
	Start()
	Stop()
	GetConnectiong() *net.TCPConn
	GetID() uint32
	GetRemotDir() net.Addr
	Sent(data []byte) error
}

type HandleFunc func(*net.TCPConn, []byte, int) error
