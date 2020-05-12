package network

import (
	"server/src/engine/coreface"
)

type BaseRouter struct {
}

func (br *BaseRouter) PreHandle(request coreface.IRequest)  {}
func (br *BaseRouter) Handle(request coreface.IRequest)     {}
func (br *BaseRouter) PostHandle(request coreface.IRequest) {}
