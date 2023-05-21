package main

import (
	"gin-template/v2/server"
	_ "gin-template/v2/service/ping"
)


func main() {
	server.Init()
	server.Start()
}
