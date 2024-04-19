package main

import (
	"github.com/MarkTBSS/079_Filtering/config"
	"github.com/MarkTBSS/079_Filtering/databases"
	"github.com/MarkTBSS/079_Filtering/server"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(conf.Database)
	server := server.NewEchoServer(conf, db)

	server.Start()
}
