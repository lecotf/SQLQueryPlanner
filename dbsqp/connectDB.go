package dbsqp

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	logsqp "sqlqp/log"
	"sqlqp/sqp"
	"strings"
)

func ConnectDB() (mdb, sdb *sql.DB) {
	logsqp.Print("Connecting to database...", logsqp.LOG)
	mdbIndex, sdbIndex := findAliases()
	mdb = initConnection(mdbIndex)
	if mdbIndex == sdbIndex {
		return mdb, mdb
	}
	sdb = initConnection(sdbIndex)
	return mdb, sdb
}

func findAliases() (mdbIndex, sdbIndex int) {
	mdbIndex, sdbIndex = -1, -1
	if sqp.Config.MonitoredDb == "" || sqp.Config.StorageDb == "" {
		logsqp.Print("MonitoredDb and StorageDb fields can't be empty. Please check the configuration file.", logsqp.FATALERROR)
	}
	for index, _ := range sqp.Config.Aliases {
		if strings.ToUpper(sqp.Config.MonitoredDb) == strings.ToUpper(sqp.Config.Aliases[index].AliasName) {
			mdbIndex = index
		}
		if strings.ToUpper(sqp.Config.StorageDb) == strings.ToUpper(sqp.Config.Aliases[index].AliasName) {
			sdbIndex = index
		}
	}
	if mdbIndex == -1 {
		logsqp.Print("\""+sqp.Config.MonitoredDb+"\" not find in the list of aliases. Please check the configuration file.", logsqp.FATALERROR)
	}
	if sdbIndex == -1 {
		logsqp.Print("\""+sqp.Config.StorageDb+"\" not find in the list of aliases. Please check the configuration file.", logsqp.FATALERROR)
	}
	return mdbIndex, sdbIndex
}

func initConnection(dbIndex int) (db *sql.DB) {
	var connectString string
	switch strings.ToUpper(sqp.Config.Aliases[dbIndex].Driver) {
	case "MSSQL":
		connectString = getConnectStringMSSQL(dbIndex)
	default:
		logsqp.Print("Unknown driver \""+sqp.Config.Aliases[dbIndex].Driver+"\". Please check the configuration file.", logsqp.FATALERROR)
	}
	logsqp.Print("Connection properties : "+connectString, logsqp.DEBUG)

	db, err := sql.Open("mssql", connectString)
	if err != nil {
		logsqp.Print(err.Error(), logsqp.ERROR)
	}
	if err = db.Ping(); err != nil {
		logsqp.Print("Unable to reach \""+sqp.Config.Aliases[dbIndex].AliasName+"\" : "+err.Error(), logsqp.FATALERROR)
	}
	logsqp.Print("Connected to database.", logsqp.LOG)
	return db
}

func getConnectStringMSSQL(dbIndex int) (connectString string) {
	// URL connection string formats
	//	connectString := "sqlserver://user:password@host/instance?database=MyDatabase&connection+timeout=30"
	connectString = fmt.Sprintf("sqlserver://%s:%s@%s/%s?connection+timeout=%d", sqp.Config.Aliases[dbIndex].User, sqp.Config.Aliases[dbIndex].PassCrypted, sqp.Config.Aliases[dbIndex].Host, sqp.Config.Aliases[dbIndex].Instance, 0)
	if sqp.Config.Aliases[dbIndex].Database != "" {
		connectString += "&database=" + sqp.Config.Aliases[dbIndex].Database
	}
	return connectString
}
