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
		"connection.ssl", true,
		"credential.username", "postgres",
		"credential.password", "postgres",
	)

	resolver := pcon.NewPostgresConnectionResolver()
	resolver.Configure(dbConfig)

	config, err := resolver.Resolve("")
	assert.Nil(t, err)

	assert.NotNil(t, config)
	assert.Equal(t, "localhost", config.Host)
	assert.Equal(t, (uint16)(5432), config.Port)
	assert.Equal(t, "test", config.Database)
	assert.Equal(t, "postgres", config.User)
	assert.Equal(t, "postgres", config.Password)
}
