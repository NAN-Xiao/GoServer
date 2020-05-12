package main

import (
	"server/src/engine/coreface"
	"server/src/engine/network"
)

type TestRouter struct {
	coreface.IRouter
}

func (tr *TestRouter) PreHandle(request coreface.IRequest) {

}
func (tr *TestRouter) Handle((request coreface.IRequest) {

}
func (tr *TestRouter) PostHandle((request coreface.IRequest) {

}

func main() {
	server := network.NewServer("v1")
	server.Serve()
}
