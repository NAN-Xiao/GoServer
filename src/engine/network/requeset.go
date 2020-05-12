package network

import (
	"server/src/engine/coreface"
)

type Request struct {
	conn coreface.IConnection
	data []byte
}

func (r *Request) GetConnection() coreface.IConnection {
	return r.conn
}

func (r *Request) Getdata() []byte {
	return r.data
}
