package main

import (
	"fmt"
	"server/engine/src/coreface"
	"server/engine/src/network"
)

type TestRouter struct {
	network.BaseRouter
}

func (tr *TestRouter) PreHandle(request coreface.IRequest) {

}
func (tr *TestRouter) Handle(request coreface.IRequest) {

}
func (tr *TestRouter) PostHandle(request coreface.IRequest) {

}

func main() {
	fmt.Println("start server")
	server := network.NewServer("v1")
	server.AddRouter(&TestRouter{})
	server.Serve()
}