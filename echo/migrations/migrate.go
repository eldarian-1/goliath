package migrations

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"goliath/repositories"
)

const (
	migrationsDir = "/Users/eldarian-13/goliath/echo/migrations/postgres"
)

func Migrate(ctx context.Context) error {
	files, err := os.ReadDir(migrationsDir)
	if err != nil {
		fmt.Printf("failed to read migrations directory: %s", err.Error())
		return nil
	}

	var sqlFiles []string
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".sql" {
			sqlFiles = append(sqlFiles, file.Name())
		}
	}
	sort.Strings(sqlFiles)

	for _, filename := range sqlFiles {
		filePath := filepath.Join(migrationsDir, filename)

		content, err := os.ReadFile(filePath)
		if err != nil {
			return fmt.Errorf("failed to read migration file %s: %w", filename, err)
		}

		if _, err := repositories.Exec(ctx, string(content)); err != nil {
			return fmt.Errorf("failed to execute migration %s: %w", filename, err)
		}

		fmt.Printf("Successfully executed migration: %s\n", filename)
	}

	fmt.Println("All migrations completed successfully")
	return nil
}
