package main

import (
    "fmt"
    "database/sql"
    // "time"
    _ "github.com/go-sql-driver/mysql"
)

func main(){
    
    store := map[string]string{
        "user":"root2",
        "password": "root",
        "database": "testdb",
    }
    

    con, err := sql.Open("mysql", store["user"]+ ":" +store["password"]+ "@/" +store["database"])
    if err != nil{
        fmt.Println(err)
    }
    // con.SetConnMaxLifetime(time.Hour * 3)
    // con.SetMaxOpenConns(10)
    // con.SetMaxIdleConns(10)
    
    err = con.Ping()
    // Open doesn't open a connection. Validate DSN data:
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    defer con.Close()
    

    // stmtIns,err := con.Prepare("insert into test1 (name) values (?), (?);")
    // if err != nil {
    //     panic(err.Error())
    // }
    // defer stmtIns.Close()
    
    // _, err = stmtIns.Exec("test1", "test2")
    // if err != nil {
    //     panic(err.Error())
    // }

    stmtOut, err := con.Prepare("select name from test1 where id = ?;")
    if err != nil {
        panic(err.Error())
    }
    defer stmtOut.Close()

    var name string
    err = stmtOut.QueryRow(1).Scan(&name)
    if err != nil {
        panic(err.Error())
    }
    fmt.Println(name)
    
    err = stmtOut.QueryRow(4).Scan(&name)
    if err != nil {
        panic(err.Error())
    }
    fmt.Println(name)

    rows, err := con.Query("select * from test1;")
    if err != nil {
        panic(err.Error())
    }
    defer rows.Close()
    
    columns, err := rows.Columns()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
    fmt.Println(columns)
    values := make([]sql.RawBytes, len(columns))
    scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}


    for rows.Next() {
        err = rows.Scan(scanArgs...)
        if err != nil {
            break
        }
        for i, col := range values {
			if col == nil {
				name = "NULL"
			} else {
				name = string(col)
			}
			fmt.Println(columns[i], ": ", name)
		}
    }
    
    fmt.Println("Done")

}