package bundb

import (
	"context"
	"crypto/tls"
	"database/sql"
	"sync"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
	"github.com/uptrace/bun/schema"

	"github.com/safciplak/capila/src/apm"
	helpers "github.com/safciplak/capila/src/helpers/environment"
	"github.com/safciplak/capila/src/logger"
)

var l sync.Mutex // used for locking the process

const (
	readTimeout  = 30 * time.Second
	writeTimeout = 5 * time.Second
)

// Connection contains the possible database connections
type Connection struct {
	Read  InterfaceDB
	Write InterfaceDB
}

// user contains the environment variables.
type user struct {
	environmentName, environmentPassword string
}

// NewDatabase instantiates a DB connection
func NewDatabase(ctx context.Context, log logger.InterfaceLogger) *Connection {
	var (
		db            *Connection
		connectionErr error
		currentTry    = 0
		maxRetries    = 10
		retryInterval = time.Duration(2)
	)

	for {
		db, connectionErr = GenerateConnections(ctx)
		if connectionErr == nil {
			break
		}

		if currentTry >= maxRetries {
			log.Log(ctx).Error("Broken DB: fatal retries exceeded: " + connectionErr.Error())

			break
		}

		log.Log(ctx).Error("Broken DB: connection error, retrying in 2 seconds: " + connectionErr.Error())

		currentTry++
		time.Sleep(time.Second * retryInterval)
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

	databaseConnection.Read = adapterDB{readDatabase}
	databaseConnection.Write = adapterDB{writeDatabase}

	return databaseConnection, nil
}

// createConnection generates the connection for the given user
func createConnection(ctx context.Context, user user) (*bun.DB, error) {
	l.Lock() // lock the process
	defer l.Unlock()

	var (
		environmentHelper = new(helpers.EnvironmentHelper)
		connector         *pgdriver.Connector
		sqlDB             *sql.DB
		db                *bun.DB
		err               error
	)

	schema.SetTableNameInflector(tableNameInflector)

	connector, err = getConnector(user)
	if err != nil {
		return nil, apm.TraceError(ctx, err)
	}

	sqlDB = sql.OpenDB(connector)
	db = bun.NewDB(sqlDB, pgdialect.New(), bun.WithDiscardUnknownColumns())
	db.AddQueryHook(NewQueryHook(&Options{
		User:     environmentHelper.Get(user.environmentName),
		Database: environmentHelper.Get("DB_DATABASE"),
	}))

	// Conditionally add debug logging
	isDebugMode, _ := environmentHelper.GetBoolean("DB_DEBUG")
	if isDebugMode {
		db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	}

	// If the connection happens to fail we should clear it and return an error
	err = db.Ping()
	if err != nil {
		return nil, apm.TraceError(ctx, err)
	}

	return db, nil
}

// tableNameInflector prefixes the table names with the correct schema
func tableNameInflector(tableName string) string {
	// default tableNameInflector = inflection.Plural
	var (
		environmentHelper = new(helpers.EnvironmentHelper)
	)

	return environmentHelper.Get("DB_SCHEMA") + "." + tableName
}

// getConnector defines a new connector for the given user
func getConnector(user user) (*pgdriver.Connector, error) {
	var (
		environmentHelper = new(helpers.EnvironmentHelper)
		tlsConfig         *tls.Config
	)

	useSSL, _ := environmentHelper.GetBoolean("DB_SSL_MODE")
	if useSSL {
		tlsConfig = &tls.Config{
			//nolint:gosec // @TODO: figure out TLS
			InsecureSkipVerify: true,
		}
	}

	conn := pgdriver.NewConnector(
		pgdriver.WithNetwork("tcp"),
		pgdriver.WithAddr(environmentHelper.Get("DB_HOST")+":"+environmentHelper.Get("DB_PORT")),
		pgdriver.WithTLSConfig(tlsConfig),
		pgdriver.WithUser(environmentHelper.Get(user.environmentName)),
		pgdriver.WithPassword(environmentHelper.Get(user.environmentPassword)),
		pgdriver.WithDatabase(environmentHelper.Get("DB_DATABASE")),
		pgdriver.WithApplicationName(environmentHelper.Get("APPLICATION_NAME")),
		pgdriver.WithTimeout(readTimeout),
		pgdriver.WithDialTimeout(writeTimeout),
		pgdriver.WithReadTimeout(readTimeout),
		pgdriver.WithWriteTimeout(writeTimeout),
		pgdriver.WithConnParams(map[string]interface{}{
			"search_path": environmentHelper.Get("DB_SCHEMA"),
		}),
	)

	return conn, environmentHelper.Err
}
