package main

import (
	"database/sql"
	"fmt"
	"github.com/golang/glog"
	_ "gopkg.in/goracle.v2"
	"log"
)

var (
	dbusername  = "AGBS"
	dbpassword  = "agbsTWitqa1"
	tnsName = "itqa"
)

func ConnectDB() *sql.DB {
	glog.Info("Connect DB")
	connStr := fmt.Sprintf("%s/%s@%s", dbusername, dbpassword, tnsName)
	db, err := sql.Open("goracle", connStr)
	if err != nil {
		log.Fatal(err)
	}
	glog.Info("Connect to DB")
	return db
}
