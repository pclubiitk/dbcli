package DB

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/godror/godror"
)

var OracleDB *sql.DB

func ConnectOracle(host, port, service, user, password string) {
    dsn := fmt.Sprintf("%s/%s@%s:%s/%s", user, password, host, port, service)

    db, err := sql.Open("godror", dsn)
    if err != nil {
        log.Fatalf("Failed to connect to Oracle: %v", err)
    }

    if err := db.Ping(); err != nil {
        log.Fatalf("Oracle ping failed: %v", err)
    }

    OracleDB = db
    log.Println("✅ Connected to Oracle database successfully")
}
