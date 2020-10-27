package test_connect

import (
	"testing"

	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
	pcon "github.com/pip-services3-go/pip-services3-postgres-go/connect"
	"github.com/stretchr/testify/assert"
)

func TestPostgresConnectionResolver(t *testing.T) {

	dbConfig := cconf.NewConfigParamsFromTuples(
		"connection.host", "localhost",
		"connection.port", 5432,
		"connection.database", "test",
		"connection.sslmode", "verify-ca",
		"credential.username", "postgres",
		"credential.password", "postgres",
	)

	resolver := pcon.NewPostgresConnectionResolver()
	resolver.Configure(dbConfig)

	config, err := resolver.Resolve("")
	assert.Nil(t, err)

	assert.NotNil(t, config)
	assert.Equal(t, "postgres://postgres:postgres@localhost:5432/test?sslmode=verify-ca", config)

}
