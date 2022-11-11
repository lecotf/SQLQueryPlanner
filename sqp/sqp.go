package sqp

type Cfg struct {
	ProjectName      string
	ProjectShortName string
	DataFolder       string `json:"DataFolder"`
	Debug            int    `json:"Debug"`
	SQLDebug         int    `json:"SQLDebug"`
	MonitoredDb      string `json:"MonitoredDb"`
	StorageDb        string `json:"StorageDb"`
	Aliases          []struct {
		AliasName   string `json:"AliasName"`
		Driver      string `json:"Driver"`
		Host        string `json:"Host"`
		Port        int    `json:"Port"`
		Instance    string `json:"Instance"`
		Tns         string `json:"Tns"`
		Database    string `json:"Database"`
		User        string `json:"User"`
		PassCrypted string `json:"PassCrypted"`
	} `json:"Aliases"`
	Queries []struct {
		QID             int    `json:"QId"`
		QTitle          string `json:"QTitle"`
		QShortName      string `json:"QShortName"`
		QTimeFirstExec  string `json:"QTimeFirstExec"`
		QTimeLastExec   string
		QIntervalTime   string `json:"QIntervalTime"`
		QManualPlanning string `json:"QManualPlanning"`
		QUsedAsFilter   bool   `json:"QUsedAsFilter"`
		QSQLText        struct {
			MSSQL      string `json:"MSSQL"`
			Oracle     string `json:"Oracle"`
			MySQL      string `json:"MySQL"`
			PostgreSQL string `json:"PostgreSQL"`
		} `json:"QSqlText"`
	} `json:"Queries"`
}

var Config Cfg
