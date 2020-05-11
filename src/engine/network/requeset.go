package network

import serverface "server/src/engine/interface"

type Request struct {
	conn serverface.IConnection
	data []byte
}

func (r *Request) GetConnection() serverface.IConnection {
	return r.conn
}
 
func (r *Request) Getdata() []byte {
	return r.data
}
