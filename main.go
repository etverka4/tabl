package main
import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
)
 
func main() { 
 
    connStr := "user=mark password=******** dbname=core sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        panic(err)
    } 
    defer db.Close()
     
    result, err := db.Exec("insert into users (email, password, firstname, lastname), values('redrgdash@gmail.com'), values('1q2w3e4r'), values('Zif'), values('Dodaf')", 
        SELECT generateUUIDv4())
    if err != nil{
        panic(err)
    }
}