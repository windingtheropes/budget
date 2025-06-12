package based

import (
	"github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"log"
)

var db *sql.DB

func DB() *sql.DB {
    if db == nil {
        log.Fatal("DB uninitialized.")
    }
    return db
}

func InitDB(cfg mysql.Config)  {
    // Get a database handle.
    var err error
    db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")
}

