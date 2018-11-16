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
	serviceName = "itqa"
)

/*
1.install gopkg.in/goracle.v2 -> ＄go get gopkg.in/goracle.v2
2.ref https://www.oracle.com/technetwork/topics/intel-macsoft-096467.html
3.Download the desired Instant Client ZIP files.
4.Add links to ~/lib or /usr/local/lib to enable applications to find the libraries.
	-> ＄ln -s ~/instantclient_12_2/libclntsh.dylib /usr/local/lib
	-> ＄ln -s ~/instantclient_12_2/libclntsh.dylib.12.1 /usr/local/lib
5.If you intend to co-locate optional Oracle configuration files such as tnsnames.ora, sqlnet.ora, ldap.ora, or oraaccess.xml
with Instant Client, then create a network/admin subdirectory.
	-> ＄mkdir -p ~/instantclient_12_2/network/admin
6.dbusername/dbpassword@serviceName
*/
func ConnectDB() *sql.DB {
	glog.Info("Connect DB")
	//connStr := fmt.Sprintf("%s/%s@%s", dbusername, dbpassword, serviceName)
	db, err := sql.Open("goracle", fmt.Sprintf("%s/%s@%s", dbusername, dbpassword, serviceName))
	if err != nil {
		log.Fatal(err)
	}
	glog.Info("Connect to DB")
	return db
}
