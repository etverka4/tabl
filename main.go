package main
import (
    "database/sql"
    "fmt"
    "github.com/jackc/pgx/v4"
)

 
func main() { 
    db, err := sql.Open("pgx", "user=mark password=******** host=localhost port=5432 database=core sslmode=disable")
    if err != nil {
    }
    defer db.Close()
    res, err := db.Exec("insert into users (email, password, firstname, lastname), values('redrgdash@gmail.com'), values('1q2w3e4r'), values('Zif'), values('Dodaf')")
    if err != nil{
        panic(err)
    }
    fmt.Println(res.LastInsertId()) 
}
