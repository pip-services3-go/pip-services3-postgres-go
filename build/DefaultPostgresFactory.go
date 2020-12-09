package build

import (
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	cbuild "github.com/pip-services3-go/pip-services3-components-go/build"
	ppersist "github.com/pip-services3-go/pip-services3-postgres-go/persistence"
)

// Creates Postgres components by their descriptors.
// See Factory
// See PostgresConnection
type DefaultPostgresFactory struct {
	cbuild.Factory
	Descriptor                   *cref.Descriptor
	PostgresConnectionDescriptor *cref.Descriptor
}

//	Create a new instance of the factory.
func NewDefaultPostgresFactory() *DefaultPostgresFactory {

	c := &DefaultPostgresFactory{

		Descriptor:                   cref.NewDescriptor("pip-services", "factory", "postgres", "default", "1.0"),
		PostgresConnectionDescriptor: cref.NewDescriptor("pip-services", "connection", "postgres", "*", "1.0"),
	}
	c.RegisterType(c.PostgresConnectionDescriptor, ppersist.NewPostgresConnection)
	return c
}
