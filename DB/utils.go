package DB

import "gorm.io/gorm"

func FetchTables(db *gorm.DB) ([]string, error) {
	var tables []string
	err := db.Raw("SHOW TABLES").Scan(&tables).Error
	return tables, err
}

func FetchColumns(db *gorm.DB, table string) ([]string, error) {
	var cols []string
	err := db.Raw("SHOW COLUMNS FROM " + table).Scan(&cols).Error
	return cols, err
}
