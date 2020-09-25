package main

import "github.com/tiennv147/http-echo/config"

func main() {
	flags := config.ParseFlags()
	server.Run(flags)
}
