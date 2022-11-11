package main

import (
	"sqlqp/config"
	"sqlqp/dbsqp"
	"sqlqp/scheduler"
	"time"
)

func main() {
	config.LoadConfig()
	mdb, sdb := dbsqp.ConnectDB()
	dbsqp.ExecQuery("SELECT NUMVERSIONCLEUNIK, DATE_MAJ, NUMARCHIVE, MODEUPDATE, NUMVERSION  FROM VERSIONS", mdb)
	dbsqp.ExecQuery("SELECT NUMVERSIONCLEUNIK, DATE_MAJ, NUMARCHIVE, MODEUPDATE, NUMVERSION  FROM DB_VERSIONS", sdb)
	dbsqp.ExecQuery("SELECT * FROM BONJOUR", mdb)

	go scheduler.QueryScheduler()
	for {
		dbsqp.RunQueryList()
		time.Sleep(10 * time.Second)
	}
}
