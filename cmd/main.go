package main

import (
	"server_explorer/pkg/configs"
	"server_explorer/pkg/server"
	"server_explorer/pkg/sqlstore"
)

func main() {
	cfg := configs.NewCfg()

	cfg.Load()
	sqlstore.New(cfg)
	server.NewHttpServer(cfg.ServerPort)
}
