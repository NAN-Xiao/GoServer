package coreface

type IRequest interface {
	GetConnection() *IConnection
	Getdata() []byte
}
