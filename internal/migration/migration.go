package migration

import (
	"database/sql"
)

type Migration struct {
	DB *sql.DB
}

func New(db *sql.DB) *Migration {
	// Create new migration instance
	return &Migration{DB: db}
}

func (m *Migration) Run(migrationDir string) error {
	// Run migrations from the specified directory
	return nil
}

// func (m *Migration) executeMigration(filePath string) error {
// Execute a single migration file
// return errors.New("not implemented")
// }
