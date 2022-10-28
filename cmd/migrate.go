package cmd

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
	"github.com/vodeacloud/hr-api/config"
	"github.com/vodeacloud/hr-api/pkg/logger"
	"os"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:       "migrate",
	Short:     "Migrate sql command",
	ValidArgs: []string{"version"},
	Args:      cobra.OnlyValidArgs,
	Run:       runMigrate,
}

var (
	migrateStep int
)

const (
	migrationPath = "db/migrations"
)

func init() {
	rootCmd.AddCommand(migrateCmd)
	migrateCmd.Flags().IntVarP(&migrateStep, "step", "s", 0, "Step migrations")
}

func runMigrate(_ *cobra.Command, args []string) {
	db, err := sql.Open("mysql", config.GetDatabaseDSN())
	if err != nil {
		logger.Fatalf("failed to open connection: %v", err)
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		logger.Fatalf("failed to create mysql instance: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", migrationPath),
		"mysql",
		driver,
	)
	if err != nil {
		logger.Fatalf("failed to construct migrate instance: %v", err)
	}

	var state string
	if len(args) == 1 {
		state = args[0]
	}

	if state == "version" {
		err = runMigrateVersion(m)
	} else if migrateStep != 0 {
		err = m.Steps(migrateStep)
	} else {
		err = m.Up()
	}

	if err != nil {
		logger.Fatalf("failed to run migrate: %v", err)
		os.Exit(1)
	}
}

func runMigrateVersion(m *migrate.Migrate) error {
	v, dirty, err := m.Version()
	if err != nil {
		return err
	}
	status := "dirty"
	if !dirty {
		status = "clean"
	}
	logger.Printf("migration version %v status '%s'", v, status)
	return nil
}
