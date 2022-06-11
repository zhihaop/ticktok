package repository

import (
	"gorm.io/gorm"
)

// CheckOrCreateTable will check if the specific table exists.
// If the table not exist, this function will create the table.
func CheckOrCreateTable(db *gorm.DB, entry any) error {
	if db.Migrator().HasTable(entry) {
		return nil
	}

	err := db.Migrator().CreateTable(entry)
	if err != nil {
		return err
	}
	return nil
}
