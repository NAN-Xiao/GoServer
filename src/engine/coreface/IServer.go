package coreface

type IServer interface {
	Start()
	Stop()
	Serve()
}
