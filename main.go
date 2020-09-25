package main

import (
	"github.com/tiennv147/http-echo/config"
	"github.com/tiennv147/http-echo/server"
)

func main() {
	flags := config.ParseFlags()
	server.Run(flags)
}
