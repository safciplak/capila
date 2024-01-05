package database

import (
	"context"
	"os"
	"testing"

	"github.com/go-pg/pg/v10"
	"github.com/stretchr/testify/assert"
)

func setup() {
	_ = os.Setenv("APPLICATION_NAME", "capila")
	_ = os.Setenv("DB_CONNECTION", "postgress")
	_ = os.Setenv("DB_HOST", "postgress")
	_ = os.Setenv("DB_PORT", "5432")
	_ = os.Setenv("DB_DATABASE", "mydatabasethatdoesntreallyexist")
	_ = os.Setenv("DB_SCHEMA", "testschema")
	_ = os.Setenv("DB_SSL_MODE", "false")

	_ = os.Setenv("DB_READ_USER", "reader")
	_ = os.Setenv("DB_READ_PASSWORD", "1234d312")
	_ = os.Setenv("DB_WRITE_USER", "sa_supah_usah")
	_ = os.Setenv("DB_WRITE_PASSWORD", "sa_supah_password")
}

func teardown() {
	os.Clearenv()
}

// TestTableNameInflector tests the tableNameInflector func
func TestTableNameInflector(t *testing.T) {
	setup()

	tableName := tableNameInflector("reservation")

	assert.Equal(t, "testschema.reservation", tableName)

	teardown()
}

// TestDebugHandler tests that nothing errors when debug is set to true
func TestDebugHandler(t *testing.T) {
	setup()

	ctx := context.Background()

	_ = os.Setenv("DB_DEBUG", "true")

	var connection, err = GenerateConnections(ctx)

	// We can only test that enabling the variable doesn't break anything, actual side effects are handled in go-pg
	assert.Error(t, err)
	assert.Nil(t, connection.Read)
	assert.Nil(t, connection.Write)

	teardown()
}

// TestGenerateConnectionString tests the connection string
func TestGenerateReadConnectionString(t *testing.T) {
	setup()

	connection, err := getOptions(user{"DB_READ_USER", "DB_READ_PASSWORD"})

	assert.Nil(t, err)
	assert.Equal(t, "reader", connection.User)
	assert.Equal(t, "mydatabasethatdoesntreallyexist", connection.Database)

	teardown()
}

// TestGenerateConnections tests the creation of the DB connections
func TestGenerateConnections(t *testing.T) {
	setup()

	ctx := context.Background()
	connection, err := GenerateConnections(ctx)

	// Since we don't actually want to connect to a DB in our tests the following asserts are not valid
	assert.NotNil(t, err)
	assert.Nil(t, connection.Read)
	assert.Nil(t, connection.Write)

	teardown()
}

// TestGenerateConnectionsErrorReadUser tests the creation of the DB connections when a Read user properties are missing
func TestGenerateConnectionsErrorReadUser(t *testing.T) {
	setup()

	ctx := context.Background()

	_ = os.Unsetenv("DB_READ_USER")
	_ = os.Unsetenv("DB_READ_PASSWORD")

	connection, err := GenerateConnections(ctx)

	// Since we don't actually want to connect to a DB in our tests the following asserts are not valid
	assert.NotNil(t, err)
	assert.Nil(t, connection.Read)
	assert.Nil(t, connection.Write)

	teardown()
}

// TestGenerateConnectionsErrorWriteUser tests the creation of the DB connections when a Write user properties are missing
func TestGenerateConnectionsErrorWriteUser(t *testing.T) {
	setup()

	ctx := context.Background()

	_ = os.Unsetenv("DB_WRITE_USER")
	_ = os.Unsetenv("DB_WRITE_PASSWORD")

	connection, err := GenerateConnections(ctx)

	// Since we don't actually want to connect to a DB in our tests the following asserts are not valid
	assert.NotNil(t, err)
	assert.Nil(t, connection.Read)
	assert.Nil(t, connection.Write)

	teardown()
}

// TestInitWithEmtpyConnectionDetails tests the init of the DB connection with empty connection details
func TestInitWithEmtpyConnectionDetails(t *testing.T) {
	setup()

	connectionDetails := &pg.Options{}
	connection, err := initialize(connectionDetails)

	// Since we don't actually want to connect to a DB in our tests the following asserts are not valid
	assert.NotNil(t, err)
	assert.Nil(t, connection)

	teardown()
}

// TestInitWithWrongDriver tests the init of the DB connection with wrong driver details
func TestInitWithWrongDriver(t *testing.T) {
	setup()

	connection, err := getOptions(user{"DB_READ_USER", "DB_READ_PASSWORD"})

	assert.Nil(t, err)

	_, err2 := initialize(&connection)

	// Since we don't actually want to connect to a DB in our tests the following asserts are not valid
	assert.NotNil(t, err2)

	teardown()
}
