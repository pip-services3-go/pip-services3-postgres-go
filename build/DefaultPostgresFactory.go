package build

import (
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	cbuild "github.com/pip-services3-go/pip-services3-components-go/build"
	conn "github.com/pip-services3-go/pip-services3-postgres-go/connect"
)

// Creates Postgres components by their descriptors.
// See Factory
// See PostgresConnection
type DefaultPostgresFactory struct {
	cbuild.Factory
}

//	Create a new instance of the factory.
func NewDefaultPostgresFactory() *DefaultPostgresFactory {

	c := &DefaultPostgresFactory{}

	postgresConnectionDescriptor := cref.NewDescriptor("pip-services", "connection", "postgres", "*", "1.0")

	c.RegisterType(postgresConnectionDescriptor, conn.NewPostgresConnection)

	return c
}
