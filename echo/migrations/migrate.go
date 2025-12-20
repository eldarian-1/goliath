package migrations

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	_ "github.com/lib/pq"
)

const (
	ConnectionStr = "host=postgresql port=5432 user=user password=password dbname=goliath sslmode=disable"
	DriverName    = "postgres"
)

func Migrate() error {
	// Connect to PostgreSQL
	db, err := sql.Open(DriverName, ConnectionStr)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	defer db.Close()

	// Test connection
	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	// Get migrations directory path
	migrationsDir := "./postgres"

	// Read all SQL files from migrations directory
	files, err := os.ReadDir(migrationsDir)
	if err != nil {
		return fmt.Errorf("failed to read migrations directory: %w", err)
	}

	// Filter and sort SQL files
	var sqlFiles []string
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".sql" {
			sqlFiles = append(sqlFiles, file.Name())
		}
	}
	sort.Strings(sqlFiles)

	// Execute each migration file
	for _, filename := range sqlFiles {
		filePath := filepath.Join(migrationsDir, filename)
		
		// Read SQL file content
		content, err := os.ReadFile(filePath)
		if err != nil {
			return fmt.Errorf("failed to read migration file %s: %w", filename, err)
		}

		// Execute SQL
		if _, err := db.Exec(string(content)); err != nil {
			return fmt.Errorf("failed to execute migration %s: %w", filename, err)
		}

		fmt.Printf("Successfully executed migration: %s\n", filename)
	}

	fmt.Println("All migrations completed successfully")
	return nil
}
