package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/sijms/go-ora/v2"
)

// Hardcoded credentials
const (
	username = "system"
	password = "ab2003an2004"
	host     = "localhost"
	port     = 1521
	service  = "freepdb1"
)

func main() {
	connStr := fmt.Sprintf("oracle://%s:%s@%s:%d/%s", username, password, host, port, service)
	db, err := sql.Open("oracle", connStr)
	if err != nil {
		log.Fatalf("Error opening DB: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Cannot connect: %v", err)
	}
	fmt.Println("Connected to Oracle successfully!")

	// Step 1: Get schema/database name
	var schema string
	fmt.Print("Enter schema/database name: ")
	fmt.Scanln(&schema)

	// Step 2: List tables
	listTables(db, schema)

	// Step 3: Choose table
	var table string
	fmt.Print("Enter table name: ")
	fmt.Scanln(&table)
	listColumns(db, schema, table)

	// Step 4: Choose columns (array)
	fmt.Print("Enter column names (comma separated): ")
	var colInput string
	fmt.Scanln(&colInput)
	cols := parseColumns(colInput)

	// Step 5: Fetch metadata + data
	fetchColumnMetadata(db, schema, table, cols)
	data := fetchSelectedColumns(db, schema, table, cols)

	// Step 6: Export to JSON
	saveToJSON(data, "selected_data.json")
	fmt.Println("\n Data exported to selected_data.json")
}

func listTables(db *sql.DB, schema string) {
	query := `
		SELECT table_name 
		FROM all_tables 
		WHERE owner = :1
		ORDER BY table_name`
	rows, err := db.Query(query, schema)
	if err != nil {
		log.Fatalf("Error listing tables: %v", err)
	}
	defer rows.Close()

	fmt.Println("\n Tables:")
	for rows.Next() {
		var table string
		rows.Scan(&table)
		fmt.Println(" -", table)
	}
}

func listColumns(db *sql.DB, schema, table string) {
	query := `
		SELECT column_name, data_type 
		FROM all_tab_columns 
		WHERE owner = :1 AND table_name = :2`
	rows, err := db.Query(query, schema, table)
	if err != nil {
		log.Fatalf("Error listing columns: %v", err)
	}
	defer rows.Close()

	fmt.Println("\n Columns:")
	for rows.Next() {
		var name, dtype string
		rows.Scan(&name, &dtype)
		fmt.Printf(" - %-20s (%s)\n", name, dtype)
	}
}

func parseColumns(input string) []string {
	parts := strings.Split(input, ",")
	cols := []string{}
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			cols = append(cols, strings.ToUpper(p))
		}
	}
	return cols
}

func fetchColumnMetadata(db *sql.DB, schema, table string, cols []string) {
	fmt.Println("\n Column Metadata:")
	for _, col := range cols {
		query := `
			SELECT column_name, data_type, data_length, nullable, data_default
			FROM all_tab_columns
			WHERE owner = :1 AND table_name = :2 AND column_name = :3`
		row := db.QueryRow(query, schema, table, col)

		var name, dtype, nullable, defVal sql.NullString
		var length sql.NullInt64
		err := row.Scan(&name, &dtype, &length, &nullable, &defVal)
		if err != nil {
			fmt.Printf(" Column %s not found or error fetching metadata: %v\n", col, err)
			continue
		}
		fmt.Printf(" - %-15s | Type: %-10s | Length: %-5d | Nullable: %-3s | Default: %s\n",
			name.String, dtype.String, length.Int64, nullable.String, defVal.String)
	}
}

func fetchSelectedColumns(db *sql.DB, schema, table string, cols []string) []map[string]interface{} {
	query := fmt.Sprintf("SELECT %s FROM %s.%s", strings.Join(cols, ","), schema, table)
	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Error fetching selected columns: %v", err)
	}
	defer rows.Close()

	// Create slice of map[string]interface{}
	results := []map[string]interface{}{}
	colNames := cols
	values := make([]interface{}, len(colNames))
	valuePtrs := make([]interface{}, len(colNames))

	for rows.Next() {
		for i := range colNames {
			valuePtrs[i] = &values[i]
		}

		if err := rows.Scan(valuePtrs...); err != nil {
			log.Fatalf("Error scanning row: %v", err)
		}

		rowMap := make(map[string]interface{})
		for i, col := range colNames {
			val := values[i]
			if b, ok := val.([]byte); ok {
				rowMap[col] = string(b)
			} else {
				rowMap[col] = val
			}
		}
		results = append(results, rowMap)
	}
	fmt.Printf("\n Retrieved %d rows from %s\n", len(results), table)
	return results
}

func saveToJSON(data []map[string]interface{}, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Error creating JSON file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(data); err != nil {
		log.Fatalf("Error writing JSON: %v", err)
	}
}
