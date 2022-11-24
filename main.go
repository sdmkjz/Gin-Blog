package main

import (
	"gin-blog/conf"
	"gin-blog/routes"
)

func main() {
	conf.Init()
	r := routes.NewRouter()
	r.Run(conf.HttpPort)
}
