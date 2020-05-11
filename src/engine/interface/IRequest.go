package serverface

type  IRequest interface {

	GetConnection() *IConnection
	Getdata() []byte

}