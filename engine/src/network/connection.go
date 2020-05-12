package network

import (
	"net"
	"server/demo/engine/coreface"
)

type Connection struct {
	_conn     *net.TCPConn
	_id       uint32
	_isClose  bool
	_exitChan chan bool
	_router   coreface.IRouter
}

func (c *Connection) Read() {

}

func (c *Connection) Start() {
	go c.Read()
	defer c.Stop()

	for {
		buf := make([]byte, 512)
		_, err := c._conn.Read(buf)
		if err != nil {
			continue
		}
		req := Request{
			conn: c,
			data: buf,
		}
		//router
		go func(request coreface.IRequest) {
			c._router.PreHandle(request)
			c._router.Handle(request)
			c._router.PostHandle(request)
		}(&req)

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

func NewConnection(conn *net.TCPConn, id uint32, router coreface.IRouter) coreface.IConnection {
	c := &Connection{
		_conn:     conn,
		_id:       id,
		_isClose:  false,
		_exitChan: make(chan bool, 1),
		_router:   router,
	}
	return c
}
