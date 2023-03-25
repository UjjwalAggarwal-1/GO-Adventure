package main

import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func main(){
    
    store := map[string]string{
        "user":"root",
        "password": "root",
        "database": "testdb",
    }
    
    con, err := sql.Open("mysql", store["user"]+":"+store["password"]+"@/"+store["database"])
    if err != nil{
        fmt.Println(err)
    }
    defer con.Close()
    
    row := con.QueryRow("select mdpr, x, y, z from sometable where id=?", 1)
    // cb := new(SomeThing)
    // err := row.Scan(&cb.Mdpr, &cb.X, &cb.Y, &cb.Z)
    fmt.Println(row)
}