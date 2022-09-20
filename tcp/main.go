package main

import (
	"database/sql"
	"entrytask/tcp/config"
	"entrytask/tcp/db"
)

func main() {
	_ = db.Db // MySQL
	_ = db.R  // Redis
	config.RPCServer.Run()
	defer func(Db *sql.DB) {
		_ = Db.Close()
	}(db.Db)
}
