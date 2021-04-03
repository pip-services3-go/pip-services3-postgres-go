package test_connect

import (
	"os"
	"testing"

	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
	conn "github.com/pip-services3-go/pip-services3-postgres-go/connect"
	"github.com/stretchr/testify/assert"
)

func TestPostgresConnection(t *testing.T) {
	var connection *conn.PostgresConnection

	postgresUri := os.Getenv("POSTGRES_URI")
	postgresHost := os.Getenv("POSTGRES_HOST")
	if postgresHost == "" {
		postgresHost = "localhost"
	}
	postgresPort := os.Getenv("POSTGRES_PORT")
	if postgresPort == "" {
		postgresPort = "5432"
	}
	postgresDatabase := os.Getenv("POSTGRES_DB")
	if postgresDatabase == "" {
		postgresDatabase = "test"
	}
	postgresUser := os.Getenv("POSTGRES_USER")
	if postgresUser == "" {
		postgresUser = "postgres"
	}
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	if postgresPassword == "" {
		postgresPassword = "postgres"
	}

	if postgresUri == "" && postgresHost == "" {
		panic("Connection params losse")
	}

	dbConfig := cconf.NewConfigParamsFromTuples(
		"connection.uri", postgresUri,
		"connection.host", postgresHost,
		"connection.port", postgresPort,
		"connection.database", postgresDatabase,
		"credential.username", postgresUser,
		"credential.password", postgresPassword,
	)

	connection = conn.NewPostgresConnection()
	connection.Configure(dbConfig)
	err := connection.Open("")
	assert.Nil(t, err)
	defer connection.Close("")

	assert.NotNil(t, connection.GetConnection())
	assert.NotNil(t, connection.GetDatabaseName())
	assert.NotEqual(t, "", connection.GetDatabaseName())
}
