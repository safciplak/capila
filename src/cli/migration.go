package cli

import (
	"fmt"
	"os/exec"

	"github.com/urfave/cli/v2"
)

const defaultMigrationPath = "database/migrations"
const defaultMigrationFormat = "20060102150405"

// getConnectionString generates the connection string that is required for the migrate package
func getConnectionString() string {
	connection := getOrFail("DB_CONNECTION")
	user := getOrFail("DB_MIGRATION_USER")
	password := getOrFail("DB_MIGRATION_PASSWORD")
	host := getOrFail("DB_HOST")
	port := getOrFail("DB_PORT")
	database := getOrFail("DB_DATABASE")
	schema := getOrFail("DB_SCHEMA")

	// @TODO: ENV variable for ssl?
	return fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s", connection, user, password, host, port, database, schema)
}

// createMigrationCommand returns the creation command
func createMigrationCommand(name string) *exec.Cmd {
	return exec.Command("migrate", "create", "-ext", "sql", "-dir", defaultMigrationPath, "-format", defaultMigrationFormat, name)
}

// generateMigrationCommand generates the default migration command with two parameters for the action per migration
func generateMigrationCommand(action, suffix string) *exec.Cmd {
	if suffix == "" {
		//nolint:gosec // migrate is allowed
		return exec.Command("migrate", "-verbose", "-path", defaultMigrationPath, "-database", getConnectionString(), action)
	}
	//nolint:gosec // migrate is allowed
	return exec.Command("migrate", "-verbose", "-path", defaultMigrationPath, "-database", getConnectionString(), action, suffix)
}

// migrateInfoCommand returns the migrate info command
func (capila *CapilaCLI) migrateInfoCommand() *cli.Command {
	return &cli.Command{
		Name:  "info",
		Usage: "Shows information about the migration package",
		Action: func(context *cli.Context) error {
			capila.run(context, generateMigrationCommand("", ""))
			return nil
		},
	}
}

// migrateCreateCommand returns the migrate create command
func (capila *CapilaCLI) migrateCreateCommand() *cli.Command {
	return &cli.Command{
		Name:  "create",
		Usage: "Supply a name and a new migration ( up & down ) will be created",
		Action: func(context *cli.Context) error {
			name := context.Args().First()

			if name == "" {
				fmt.Printf("%s", "Please provide a name \n")
				return nil
			}

			capila.run(context, createMigrationCommand(name))
			return nil
		},
	}
}

// migrateGotoCommand returns the migrate goto command
func (capila *CapilaCLI) migrateGotoCommand() *cli.Command {
	return &cli.Command{
		Name:  "goto",
		Usage: "Migrate to version V",
		Action: func(context *cli.Context) error {
			version := context.Args().First()

			if version == "" {
				fmt.Printf("%s", "Please provide a version \n")
				return nil
			}

			capila.run(context, generateMigrationCommand("goto", version))
			return nil
		},
	}
}

// migrateUpCommand returns the migrate up command
func (capila *CapilaCLI) migrateUpCommand() *cli.Command {
	return &cli.Command{
		Name:  "up",
		Usage: "Apply all or N up migrations",
		Action: func(context *cli.Context) error {
			amountOfMigrations := context.Args().First()
			capila.run(context, generateMigrationCommand("up", amountOfMigrations))
			return nil
		},
	}
}

// migrateDownCommand returns the migrate down command
func (capila *CapilaCLI) migrateDownCommand() *cli.Command {
	return &cli.Command{
		Name:  "down",
		Usage: "Apply all or N down migrations",
		Action: func(context *cli.Context) error {
			amountOfMigrations := context.Args().First()

			if amountOfMigrations == "" {
				amountOfMigrations = "-all"
			}

			capila.run(context, generateMigrationCommand("down", amountOfMigrations))
			return nil
		},
	}
}

// migrateDropCommand returns the migrate drop command
func (capila *CapilaCLI) migrateDropCommand() *cli.Command {
	return &cli.Command{
		Name:  "drop",
		Usage: "Drop everything inside database",
		Action: func(context *cli.Context) error {
			capila.run(context, generateMigrationCommand("drop", "-f"))
			return nil
		},
	}
}

// migrateFreshCommand returns the migrate fresh command
func (capila *CapilaCLI) migrateFreshCommand() *cli.Command {
	return &cli.Command{
		Name:  "fresh",
		Usage: "Create a fresh instance",
		Action: func(context *cli.Context) error {
			capila.run(context, generateMigrationCommand("drop", "-f"))
			capila.run(context, generateMigrationCommand("up", ""))
			return nil
		},
	}
}

// migrateFreshCommand returns the migrate force command
func (capila *CapilaCLI) migrateForceCommand() *cli.Command {
	return &cli.Command{
		Name:  "force",
		Usage: "Set version V but don't run migration (ignores dirty state)",
		Action: func(context *cli.Context) error {
			version := context.Args().First()
			capila.run(context, generateMigrationCommand("force", version))
			return nil
		},
	}
}

// migrateVersionCommand returns the migrate version command
func (capila *CapilaCLI) migrateVersionCommand() *cli.Command {
	return &cli.Command{
		Name:  "version",
		Usage: "Print current migration version",
		Action: func(context *cli.Context) error {
			command := generateMigrationCommand("version", "")
			capila.run(context, command)
			return nil
		},
	}
}

// registerMigrations registers migration commands
func (capila *CapilaCLI) registerMigrations() {
	cmd := &cli.Command{
		Name:  "migrate",
		Usage: "Run and check for migrations CLI",
		Subcommands: []*cli.Command{
			capila.migrateInfoCommand(),
			capila.migrateCreateCommand(),
			capila.migrateGotoCommand(),
			capila.migrateUpCommand(),
			capila.migrateDownCommand(),
			capila.migrateDropCommand(),
			capila.migrateFreshCommand(),
			capila.migrateForceCommand(),
			capila.migrateVersionCommand(),
		},
	}

	capila.cli.Commands = append(capila.cli.Commands, cmd)
}
