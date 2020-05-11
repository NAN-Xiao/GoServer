package main

import "server/src/engine/network"

func  main()  {
	server:=network.NewServer("v1")
	server.Serve()
}
