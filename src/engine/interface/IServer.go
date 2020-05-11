package serverface

type IServer interface {
	Start()
	Stop()
	Serve()
}

