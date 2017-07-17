package main

import (
    "fmt"
    "os"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

var dsn = "work:work@(localhost:3306)/mysql"

type MysqlUser struct {
    Host sql.NullString
    Uname sql.NullString
    Passwd sql.NullString
}

func main() {
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        fmt.Printf("connect to mysql[%s] error: %s\n", dsn, err.Error())
        os.Exit(1)
    }
    defer db.Close()
    fmt.Printf("Connect to [%s] successed\n", dsn)

    sql := "SELECT Host, User, Password FROM mysql.user LIMIT 10"
    rows, err := db.Query(sql)
    if err != nil {
        fmt.Printf("Query failed[%s], SQL: %s\n", err.Error(), sql)
        os.Exit(1)
    }
    defer rows.Close()
    fmt.Printf("Query successed. SQL: %s\n", sql)

    rowDatas := make([]MysqlUser, 10)
    for i := 0; i < 4 && rows.Next(); i ++ {
        err = rows.Scan(&rowDatas[i].Host, &rowDatas[i].Uname, &rowDatas[i].Passwd)
        if err != nil {
            fmt.Printf("Scan failed[%s]\n", err.Error())
            os.Exit(1)
        }
    }
    fmt.Printf("Scan done\n")

    for idx, row := range rowDatas {
        fmt.Printf("row[%d]: Host = %s, Uname = %s, Passwd = %s\n", idx, row.Host.String, row.Uname.String, row.Passwd.String)
    }
}