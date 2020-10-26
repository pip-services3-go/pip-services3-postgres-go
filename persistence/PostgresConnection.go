package persistence

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
	cerr "github.com/pip-services3-go/pip-services3-commons-go/errors"
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	clog "github.com/pip-services3-go/pip-services3-components-go/log"
	pcon "github.com/pip-services3-go/pip-services3-postgres-go/connect"
)

/**
 * PostgreSQL connection using plain driver.
 *
 * By defining a connection and sharing it through multiple persistence components
 * you can reduce number of used database connections.
 *
 * ### Configuration parameters ###
 *
 * - connection(s):
 *   - discovery_key:             (optional) a key to retrieve the connection from [[IDiscovery]]
 *   - host:                      host name or IP address
 *   - port:                      port number (default: 27017)
 *   - uri:                       resource URI or connection string with all parameters in it
 * - credential(s):
 *   - store_key:                 (optional) a key to retrieve the credentials from [[ICredentialStore]]
 *   - username:                  user name
 *   - password:                  user password
 * - options:
 *   - connect_timeout:      (optional) number of milliseconds to wait before timing out when connecting a new client (default: 0)
 *   - idle_timeout:         (optional) number of milliseconds a client must sit idle in the pool and not be checked out (default: 10000)
 *   - max_pool_size:        (optional) maximum number of clients the pool should contain (default: 10)
 *
 * ### References ###
 *
 * - \*:logger:\*:\*:1.0           (optional) [[ILogger]] components to pass log messages
 * - \*:discovery:\*:\*:1.0        (optional) [[ IDiscovery]] services
 * - \*:credential-store:\*:\*:1.0 (optional) Credential stores to resolve credentials
 *
 */
type PostgresConnection struct {
	defaultConfig *cconf.ConfigParams
	// The logger.
	Logger *clog.CompositeLogger
	// The connection resolver.
	ConnectionResolver *pcon.PostgresConnectionResolver
	// The configuration options.
	Options *cconf.ConfigParams
	// The PostgreSQL connection pool object.
	Connection *pgxpool.Pool
	// The PostgreSQL database name.
	DatabaseName string
}

// NewPostgresConnection creates a new instance of the connection component.
func NewPostgresConnection() *PostgresConnection {
	c := &PostgresConnection{
		defaultConfig: cconf.NewConfigParamsFromTuples(

			"options.connect_timeout", 0,
			"options.idle_timeout", 10000,
			"options.max_pool_size", 3,
		),
		Logger:             clog.NewCompositeLogger(),
		ConnectionResolver: pcon.NewPostgresConnectionResolver(),
		Options:            cconf.NewEmptyConfigParams(),
	}
	return c
}

// Configures component by passing configuration parameters.
//   - config    configuration parameters to be set.
func (c *PostgresConnection) Configure(config *cconf.ConfigParams) {
	config = config.SetDefaults(c.defaultConfig)

	c.ConnectionResolver.Configure(config)

	c.Options = c.Options.Override(config.GetSection("options"))
}

// Sets references to dependent components.
//  - references 	references to locate the component dependencies.
func (c *PostgresConnection) SetReferences(references cref.IReferences) {
	c.Logger.SetReferences(references)
	c.ConnectionResolver.SetReferences(references)
}

// Checks if the component is opened.
// Returns true if the component has been opened and false otherwise.
func (c *PostgresConnection) IsOpen() bool {
	return c.Connection != nil
}

// Opens the component.
//  - correlationId 	(optional) transaction id to trace execution through call chain.
//  - Return 			error or nil no errors occured.
func (c *PostgresConnection) Open(correlationId string) error {

	config, err := c.ConnectionResolver.Resolve(correlationId)

	if err != nil {
		c.Logger.Error(correlationId, err, "Failed to resolve Postgres connection")
		return nil
	}

	maxPoolSize := c.Options.GetAsNullableInteger("max_pool_size")
	idleTimeoutMS := c.Options.GetAsNullableInteger("idle_timeout")
	connectTimeoutMS := c.Options.GetAsNullableInteger("connect_timeout")

	if connectTimeoutMS != nil && *connectTimeoutMS != 0 {
		config.ConnectTimeout = time.Duration((int64)(*connectTimeoutMS)) * time.Millisecond
	}

	c.Logger.Debug(correlationId, "Connecting to postgres")

	options := &pgxpool.Config{
		ConnConfig: config,
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), options)

	if err != nil || pool == nil {
		err = cerr.NewConnectionError(correlationId, "CONNECT_FAILED", "Connection to postgres failed").WithCause(err)
	} else {
		if idleTimeoutMS != nil && *idleTimeoutMS != 0 {
			pool.Config().MaxConnIdleTime = time.Duration((int64)(*idleTimeoutMS)) * time.Millisecond
		}
		if maxPoolSize != nil && *maxPoolSize != 0 {
			pool.Config().MaxConns = (int32)(*maxPoolSize)
		}
		c.Connection = pool
		c.DatabaseName = config.Database
	}
	return err
}

// Closes component and frees used resources.
//  - correlationId 	(optional) transaction id to trace execution through call chain.
// Return			 error or nil no errors occured
func (c *PostgresConnection) Close(correlationId string) error {
	if c.Connection == nil {
		return nil
	}
	c.Connection.Close()
	c.Logger.Debug(correlationId, "Disconnected from postgres database %s", c.DatabaseName)
	c.Connection = nil
	c.DatabaseName = ""
	return nil
}

func (c *PostgresConnection) GetConnection() *pgxpool.Pool {
	return c.Connection
}

func (c *PostgresConnection) GetDatabaseName() string {
	return c.DatabaseName
}
