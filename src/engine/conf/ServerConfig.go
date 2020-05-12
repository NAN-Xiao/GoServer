package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"server/src/engine/coreface"
)

type ServerConfig struct {
	SERVER   coreface.IServer
	HOST     string
	TCP_PORT int
	NAME     string

	VERSION        string
	MAX_CON       int
	MAX_PACKAGE_SIZE int
}

var G_serverConfig *ServerConfig

func init() {

	data, err := ioutil.ReadFile("/conf/serverconfig.json")
	if err != nil {
		fmt.Println("load json error")
	}
	err = json.Unmarshal(data, &ServerConfig{})
	if err != nil {
		fmt.Println("ser json error")
	}
	fmt.Println("load config success")
}
