package test

import (
	"os"
	"testing"

	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	ppersist "github.com/pip-services3-go/pip-services3-postgres-go/persistence"
	tf "github.com/pip-services3-go/pip-services3-postgres-go/test/fixtures"
	"github.com/stretchr/testify/assert"
)

func TestDummyPostgresConnection(t *testing.T) {

	var persistence *DummyPostgresPersistence
	var fixture tf.DummyPersistenceFixture
	var connection *ppersist.PostgresConnection

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
		panic("Connection params not set")
	}

	dbConfig := cconf.NewConfigParamsFromTuples(
		"connection.uri", postgresUri,
		"connection.host", postgresHost,
		"connection.port", postgresPort,
		"connection.database", postgresDatabase,
		"credential.username", postgresUser,
		"credential.password", postgresPassword,
	)

	connection = ppersist.NewPostgresConnection()
	connection.Configure(dbConfig)

	persistence = NewDummyPostgresPersistence()
	descr := cref.NewDescriptor("pip-services", "connection", "postgres", "default", "1.0")
	ref := cref.NewReferencesFromTuples(descr, connection)
	persistence.SetReferences(ref)

	fixture = *tf.NewDummyPersistenceFixture(persistence)

	opnErr := connection.Open("")
	if opnErr != nil {
		t.Error("Error opened connection", opnErr)
		return
	}
	defer connection.Close("")

	opnErr = persistence.Open("")
	if opnErr != nil {
		t.Error("Error opened persistence", opnErr)
		return
	}
	defer persistence.Close("")

	opnErr = persistence.Clear("")
	if opnErr != nil {
		t.Error("Error cleaned persistence", opnErr)
		return
	}

	t.Run("Connection", func(t *testing.T) {
		assert.NotNil(t, connection.GetConnection())
		assert.NotNil(t, connection.GetDatabaseName())
		assert.NotEqual(t, "", connection.GetDatabaseName())
	})

	t.Run("DummyPostgresConnection:CRUD", fixture.TestCrudOperations)

	opnErr = persistence.Clear("")
	if opnErr != nil {
		t.Error("Error cleaned persistence", opnErr)
		return
	}

	t.Run("DummyPostgresConnection:Batch", fixture.TestBatchOperations)

}
