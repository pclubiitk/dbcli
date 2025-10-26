package DB

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/cengsin/oracle"

	"gorm.io/gorm"
)

var OracleDB *gorm.DB

func ConnectOracle(host, port, password, dbName, user string) {
	// Data Source Name (DSN) format for Oracle
	// Example: user:password@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbName)

	database, err := gorm.Open(oracle.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Fatal("Failed to connect to Oracle database: ", err)
	}

	OracleDB = database

	logrus.Info("✅ Connected to Oracle database successfully")
}
