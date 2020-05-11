package network

import (
	"net"
	serverface "server/src/engine/interface"
)

type Connection struct {
	_conn     *net.TCPConn
	_id       uint32
	_isClose  bool
	_handle   serverface.HandleFunc
	_exitChan chan bool
}

func NewConnection(conn *net.TCPConn, id uint32, callback serverface.HandleFunc) *Connection {
	c := &Connection{
		_conn:     conn,
		_id:       id,
		_isClose:  false,
		_handle:   callback,
		_exitChan: make(chan bool, 1),
	}
	return c
}

func (c *Connection) Read() {

}

func (c *Connection) Start() {
	go c.Read()
	defer c.Stop()
	for {
		buf := make([]byte, 512)
		cnt, err := c._conn.Read(buf)
		if err != nil {
			continue
		}

		if	err := c._handle(c._conn, buf, cnt);err!=nil{
			break
		}

	}
}
func (c *Connection) Stop() {
	if c._isClose == true {
		return
	}
	c._isClose = true
	c._conn.Close()
	close(c._exitChan)

}
func (c *Connection) GetConnectiong() *net.TCPConn {
	return c._conn
}
func (c *Connection) GetID() uint32 {
	return c._id
}
func (c *Connection) GetRemotDir() net.Addr {
	return c._conn.RemoteAddr()
}
func (c *Connection) Sent(data []byte) error {
	return nil
}
