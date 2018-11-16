package main

import (
	"database/sql"
	"flag"
	"github.com/golang/glog"
)

var db *sql.DB

func main() {
	// Parse CLI flags
	flag.Parse()
	glog.Info("Starting Up")
	glog.Info("Connecting to database")
	db = ConnectDB()

	//Query database
	if err := db.Ping(); err != nil {
		glog.Fatal("Error:", err)
	} else {
		glog.Info("Ping database success!")
	}
	row := db.QueryRow("SELECT 1 FROM DUAL")
	var dummy string
	if err := row.Scan(&dummy); err != nil {
		glog.Fatal(err)
	}
	if dummy != "1" {
		glog.Fatal("Failed to connect to database")
	} else {
		glog.Info("Success to connect to database")
	}
	//var code string
	//var description string
	//row := db.QueryRow("SELECT CODE, DESCRIPTION FROM AC_REGION WHERE CODE = :1 AND DESCRIPTION = :2", "CN", "China")
	//if err := row.Scan(&code, &description); err != nil {
	//	glog.Fatal(err)
	//}
	//glog.Info("Code:", code)
	//glog.Info("Description:", description)
	//glog.Info("Shutting Down")
	defer db.Close()
	glog.Flush()
}
