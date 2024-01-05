package database

import (
	"context"
	"strings"
	"sync"

	"github.com/go-pg/pg/extra/pgdebug/v10"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/pkg/errors"

	"github.com/safciplak/capila/src/apm"
	helpers "github.com/safciplak/capila/src/helpers/environment"
	"github.com/safciplak/capila/src/logger"
)

//nolint:gochecknoglobals // this global is needed for go-pg
var l sync.Mutex // used for locking the process

// Connection contains the possible database connections
type Connection struct {
	Read  InterfacePGDB
	Write InterfacePGDB
}

// user contains the environment variables.
type user struct {
	environmentName, environmentPassword string
}

// initialize initializes the database connection
func initialize(connection *pg.Options) (*pg.DB, error) {
	l.Lock() // lock the process
	defer l.Unlock()

	// Fix pluralization and set correct schema.
	orm.SetTableNameInflector(tableNameInflector)

	db := pg.Connect(connection)
	environmentHelper := new(helpers.EnvironmentHelper)

	// Conditionally add debug logging
	isDebugMode, _ := environmentHelper.GetBoolean("DB_DEBUG")
	if isDebugMode {
		db.AddQueryHook(&pgdebug.DebugHook{
			// Print all queries.
			Verbose: true,
		})
	}

	// Register the APM instrumentation
	db.AddQueryHook(queryHook{})

	// If the connection happens to fail we should clear it and return an error
	err := db.Ping(db.Context())
	if err != nil {
		return nil, err
	}

	return db, err
}

// tableNameInflector prefixes the table names with the correct schema
func tableNameInflector(tableName string) string {
	var environmentHelper = helpers.EnvironmentHelper{}
	return environmentHelper.Get("DB_SCHEMA") + "." + strings.ReplaceAll(tableName, "_", "")
}

// NewDatabase instantiates a DB connection
func NewDatabase(ctx context.Context, log logger.InterfaceLogger) *Connection {
	db, err := GenerateConnections(ctx)
	if err != nil {
		log.Log(ctx).Error("Start: broken DB")
	}

	return db
}

// GenerateConnections generates the read and write connections.
func GenerateConnections(ctx context.Context) (*Connection, error) {
	var databaseConnection = new(Connection)

	readDatabase, err := createConnection(ctx,
		user{"DB_READ_USER", "DB_READ_PASSWORD"})

	if err != nil {
		return databaseConnection, apm.TraceError(ctx, err)
	}

	writeDatabase, err := createConnection(ctx,
		user{"DB_WRITE_USER", "DB_WRITE_PASSWORD"})

	if err != nil {
		return databaseConnection, apm.TraceError(ctx, err)
	}

	databaseConnection.Read = adapterPgDB{readDatabase}
	databaseConnection.Write = adapterPgDB{writeDatabase}

	return databaseConnection, nil
}

// createConnection generates the connection for the given user
func createConnection(ctx context.Context, user user) (*pg.DB, error) {
	var connection *pg.DB

	config, err := getOptions(user)

	if err != nil {
		_ = apm.TraceError(ctx, err)

		return connection, errors.New("config generation went wrong")
	}

	return initialize(&config)
}

// getOptions generates the database options for the given user
func getOptions(user user) (pg.Options, error) {
	var (
		environmentHelper = new(helpers.EnvironmentHelper)
	)

	databaseConnector := pg.Options{
		ApplicationName: environmentHelper.Get("APPLICATION_NAME"),
		Addr:            environmentHelper.Get("DB_HOST") + ":" + environmentHelper.Get("DB_PORT"),
		User:            environmentHelper.Get(user.environmentName),
		Password:        environmentHelper.Get(user.environmentPassword),
		Database:        environmentHelper.Get("DB_DATABASE"),
		TLSConfig:       nil, // @NOTE: at the moment, our database is only accessible within it's own cluster. No SSL needed for now.
	}

	return databaseConnector, environmentHelper.Err
}
