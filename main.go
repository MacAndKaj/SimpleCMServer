package main

import (
	"SimpleCMServer/config"
	"SimpleCMServer/server"
	"SimpleCMServer/server/cmshandlers"
	"log"
	"os"
)

func main() {
	l := log.New(os.Stdout, "[SimpleCMServer]", log.LstdFlags)

	serverArguments := config.ParseArguments(os.Args)

	cms := server.NewServer(serverArguments.Port, l, cmshandlers.BuildHandlers(l))
	cms.GoServer()
}
