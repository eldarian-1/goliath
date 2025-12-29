package migrations

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"goliath/repositories"
	"goliath/utils"
)

const (
	defaultMigrationsDir = "/Users/eldarian-13/goliath/service/migrations/postgres"
)

var migrationsDir string

func init() {
	migrationsDir = utils.GetEnv("GOLIATH_MIGRATIONS", defaultMigrationsDir)
}

func Migrate(ctx context.Context) error {
	fmt.Println("Start migrations")

	dir := fmt.Sprintf("%s/postgres", migrationsDir)

	files, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("failed to read directory: %w", err)
	}

	var sqlFiles []string
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".sql" {
			sqlFiles = append(sqlFiles, file.Name())
		}
	}
	sort.Strings(sqlFiles)

	for _, filename := range sqlFiles {
		filePath := filepath.Join(dir, filename)

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
