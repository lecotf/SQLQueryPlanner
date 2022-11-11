package config

import (
	"encoding/json"
	"fmt"
	"os"
	logsqp "sqlqp/log"
	"sqlqp/sqp"
	"time"
)

func LoadConfig() {
	sqp.Config.ProjectName = "SQLQueryPlanner"
	sqp.Config.ProjectShortName = "SQP"
	sqp.Config.DataFolder = "DATA/Logs"
	sqp.Config.Debug = 0
	sqp.Config.SQLDebug = 0

	logsqp.Print("Starting "+sqp.Config.ProjectName+" at "+time.Now().String(), logsqp.LOG)
	b, err := os.ReadFile(sqp.Config.ProjectName + ".json")
	if err != nil {
		logsqp.Print(err.Error(), logsqp.FATALERROR)
	}

	if json.Valid(b) == false {
		logsqp.Print(sqp.Config.ProjectName+" is invalid. Please check the configuration file.\r\n"+sqp.Config.ProjectName+".json = \r\n"+string(b), logsqp.FATALERROR)
	}
	err = json.Unmarshal(b, &sqp.Config)
	if err != nil {
		logsqp.Print(err.Error(), logsqp.FATALERROR)
	}
	logsqp.Print("Result of unmarshalling "+sqp.Config.ProjectName+".json :\r\n"+fmt.Sprintf("%#v", sqp.Config), logsqp.DEBUG)
	logsqp.Print("Configuration loaded", logsqp.LOG)
}
