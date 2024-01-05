package cli

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

// setup the env
func setup() {
	os.Setenv("APPLICATION_NAME", "capila")

	os.Setenv("DB_CONNECTION", "postgres")
	os.Setenv("DB_HOST", "postgres")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_DATABASE", "mydatabasethatdoesntreallyexist")
	os.Setenv("DB_SCHEMA", "testschema")
	os.Setenv("DB_MIGRATION_USER", "sa_supah_usah")
	os.Setenv("DB_MIGRATION_PASSWORD", "sa_supah_password")
}

// teardown the evn
func teardown() {
	os.Clearenv()
}

func Test_Run(t *testing.T) {
	setup()

	var arguments []string

	arguments = append(arguments, "--help")

	Run(arguments)
	teardown()
}

// Test_setInfo
func Test_SetInfo(t *testing.T) {
	setup()

	var capila CapilaCLI

	capila.cli = cli.NewApp()

	capila.setInfo()

	assert.NotEqual(t, "", capila.cli.Name)
	assert.NotEqual(t, "", capila.cli.Usage)
	assert.NotEqual(t, "", capila.cli.Version)

	teardown()
}

// Test_registerMigrations
func Test_registerMigrations(t *testing.T) {
	setup()

	var capila CapilaCLI

	capila.cli = cli.NewApp()

	capila.registerMigrations()

	migrationCommand := capila.cli.Command("migrate")

	// Test 1 parent command
	assert.Equal(t, 1, len(capila.cli.Commands))
	assert.NotNil(t, migrationCommand)
	assert.Greater(t, len(migrationCommand.Subcommands), 0)

	for _, cmd := range migrationCommand.Subcommands {
		assert.NotNil(t, cmd.Name)

		err := capila.cli.Run([]string{"capila", "migrate", cmd.Name})

		assert.Nil(t, err)
	}

	teardown()
}

func Test_GetConnectionString(t *testing.T) {
	setup()

	connectionString := getConnectionString()

	assert.Equal(t, "postgres://sa_supah_usah:sa_supah_password@postgres:5432/mydatabasethatdoesntreallyexist?sslmode=disable&search_path=testschema", connectionString)

	teardown()
}

func Test_GetOrFail(t *testing.T) {
	setup()

	environmentVariable := getOrFail("DB_CONNECTION")

	assert.Equal(t, "postgres", environmentVariable)

	setup()
}

func Test_GenerateMigrationCommand(t *testing.T) {
	setup()

	emptySecondCommand := generateMigrationCommand("runDmc", "")
	assert.NotNil(t, emptySecondCommand)

	withSecondCommand := generateMigrationCommand("runDmc", "yes")
	assert.NotNil(t, withSecondCommand)

	teardown()
}
