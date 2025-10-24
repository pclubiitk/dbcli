package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// ListTables connects to the database and prints all table names
func ListTables(username, password, dbName string) {
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", username, password, dbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Failed to connect to MySQL:", err)
	}

	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Tables in the database:")
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			log.Fatal(err)
		}
		fmt.Println("-", tableName)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

// DescribeTable connects to the database and prints column details for a given table
func DescribeTable(username, password, dbName, tableName string) {
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", username, password, dbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	query := fmt.Sprintf("DESCRIBE %s", tableName)
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Printf("Description of table `%s`:\n", tableName)
	fmt.Println("Field\tType\tNull\tKey\tDefault\tExtra")

	for rows.Next() {
		var field, typ, null, key, extra string
		var defaultVal sql.NullString
		if err := rows.Scan(&field, &typ, &null, &key, &defaultVal, &extra); err != nil {
			log.Fatal(err)
		}
		def := "NULL"
		if defaultVal.Valid {
			def = defaultVal.String
		}
		fmt.Printf("%s\t%s\t%s\t%s\t%s\t%s\n", field, typ, null, key, def, extra)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Example usage:
	username := "root"
	password := "my_password"
	database := "college"

	// List all tables
	fmt.Print("Do you want to see all existing tables? (y/n): ")
	var response string
	fmt.Scanln(&response)
	if response == "y" {
		ListTables(username, password, database)
	}

	fmt.Println()

	// Describe a specific table
	fmt.Print("Do you want to describe a specific table? (y/n): ")
	fmt.Scanln(&response)
	if response != "y" {
		return
	}

	fmt.Print("Enter the table name to describe:")
	var table string
	fmt.Scanln(&table)
	DescribeTable(username, password, database, table)
}
