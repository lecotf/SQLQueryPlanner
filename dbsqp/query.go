package dbsqp

import (
	"database/sql"
	"fmt"
	logsqp "sqlqp/log"
)

func ExecQuery(query string, db *sql.DB) {
	println("Exec : " + query)
	rows, err := db.Query(query)
	if err != nil {
		logsqp.Print(err.Error(), logsqp.SQLERROR)
		return
	}
	println("Result : ")
	println(rows)

	// get the column names from the query
	cols, err := rows.Columns()
	if err != nil {
		logsqp.Print(err.Error(), logsqp.SQLERROR)
		return
	}

	data := make(map[string]string)

	for rows.Next() {
		columns := make([]string, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}

		rows.Scan(columnPointers...)

		for i, colName := range cols {
			data[colName] = columns[i]
			if data[colName] == "" {
				fmt.Printf("%s=NULL ; ", colName)
			} else {
				fmt.Printf("%s=%s ; ", colName, data[colName])
			}
		}
		println()
	}

	for k, v := range data {
		fmt.Println(k, "value is", v)
	}

	/*
		// Remember to check err afterwards
		vals := make([]interface{}, len(cols))
		for i, _ := range cols {
			vals[i] = new(sql.RawBytes)
		}
		for rows.Next() {
			err = rows.Scan(vals...)
			// Now you can check each element of vals for nil-ness,
			// and you can use type introspection and type assertions
			// to fetch the column into a typed variable.
		}
		for i, _ := range cols {
			fmt.Println(vals[i])
		}
	*/
}

func RunQueryList() {
	fmt.Println("...query done")
}
