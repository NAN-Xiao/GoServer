package coreface

type IRequest interface {
	GetConnection() IConnection
	GetData() []byte
}
