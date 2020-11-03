package test

import (
	"os"
	"testing"

	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
	tf "github.com/pip-services3-go/pip-services3-postgres-go/test/fixtures"
)

func TestDummyMapPostgresPersistence(t *testing.T) {

	var persistence *DummyMapPostgresPersistence
	var fixture tf.DummyMapPersistenceFixture

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

	persistence = NewDummyMapPostgresPersistence()
	persistence.Configure(dbConfig)

	fixture = *tf.NewDummyMapPersistenceFixture(persistence)

	opnErr := persistence.Open("")
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

	t.Run("DummyMapPostgresPersistence:CRUD", fixture.TestCrudOperations)

	opnErr = persistence.Clear("")
	if opnErr != nil {
		t.Error("Error cleaned persistence", opnErr)
		return
	}

	t.Run("DummyMapPostgresPersistence:Batch", fixture.TestBatchOperations)

}
