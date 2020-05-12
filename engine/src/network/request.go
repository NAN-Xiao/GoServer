package network

import (
	"server/engine/src/coreface"
)

type Request struct {
	conn coreface.IConnection
	data []byte
}

func (r *Request) GetConnection() coreface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.data
}
