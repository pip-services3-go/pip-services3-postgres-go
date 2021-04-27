package persistence

import (
	"context"
	"reflect"
	"strconv"

	cconv "github.com/pip-services3-go/pip-services3-commons-go/convert"
	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
	cmpersist "github.com/pip-services3-go/pip-services3-data-go/persistence"
)

/*
Abstract persistence component that stores data in PostgreSQL
and implements a number of CRUD operations over data items with unique ids.
The data items must implement IIdentifiable interface.

In basic scenarios child classes shall only override getPageByFilter,
getListByFilter or deleteByFilter operations with specific filter function.
All other operations can be used out of the box.

In complex scenarios child classes can implement additional operations by
accessing c._collection and c._model properties.

### Configuration parameters ###

- collection:                  (optional) PostgreSQL collection name
- connection(s):
   - discovery_key:             (optional) a key to retrieve the connection from IDiscovery
   - host:                      host name or IP address
   - port:                      port number (default: 27017)
   - uri:                       resource URI or connection string with all parameters in it
- credential(s):
   - store_key:                 (optional) a key to retrieve the credentials from ICredentialStore
   - username:                  (optional) user name
   - password:                  (optional) user password
- options:
   - connect_timeout:      (optional) number of milliseconds to wait before timing out when connecting a new client (default: 0)
   - idle_timeout:         (optional) number of milliseconds a client must sit idle in the pool and not be checked out (default: 10000)
   - max_pool_size:        (optional) maximum number of clients the pool should contain (default: 10)

### References ###

- \*:logger:\*:\*:1.0           (optional) ILogger components to pass log messages components to pass log messages
- \*:discovery:\*:\*:1.0        (optional) IDiscovery services
- \*:credential-store:\*:\*:1.0 (optional) Credential stores to resolve credentials
 *
### Example ###

    class MyPostgresPersistence extends IdentifiablePostgresPersistence<MyData, string> {

    public constructor() {
        base("mydata", new MyDataPostgresSchema());
    }

    private composeFilter(filter: FilterParams): any {
        filter = filter || new FilterParams();
        let criteria = [];
        let name = filter.getAsNullableString('name');
        if (name != null)
            criteria.push({ name: name });
        return criteria.length > 0 ? { $and: criteria } : null;
    }

    public getPageByFilter(correlationId: string, filter: FilterParams, paging: PagingParams,
        callback: (err: any, page: DataPage<MyData>) => void): void {
        base.getPageByFilter(correlationId, c.composeFilter(filter), paging, null, null, callback);
    }

    }

    let persistence = new MyPostgresPersistence();
    persistence.configure(ConfigParams.fromTuples(
        "host", "localhost",
        "port", 27017
    ));

    persitence.open("123", (err) => {
        ...
    });

    persistence.create("123", { id: "1", name: "ABC" }, (err, item) => {
        persistence.getPageByFilter(
            "123",
            FilterParams.fromTuples("name", "ABC"),
            null,
            (err, page) => {
                console.log(page.data);          // Result: { id: "1", name: "ABC" }

                persistence.deleteById("123", "1", (err, item) => {
                   ...
                });
            }
        )
    });
*/
type IdentifiablePostgresPersistence struct {
	*PostgresPersistence
}

// Creates a new instance of the persistence component.
//   - overrides References to override virtual methods
//   - tableName    (optional) a table name.
func InheritIdentifiablePostgresPersistence(overrides IPostgresPersistenceOverrides, proto reflect.Type, tableName string) *IdentifiablePostgresPersistence {
	if tableName == "" {
		panic("Table name could not be empty")
	}

	c := &IdentifiablePostgresPersistence{}
	c.PostgresPersistence = InheritPostgresPersistence(overrides, proto, tableName)

	return c
}

// Gets a list of data items retrieved by given unique ids.
//   - correlationId     (optional) transaction id to trace execution through call chain.
//   - ids               ids of data items to be retrieved
// Returns          a data list or error.
func (c *IdentifiablePostgresPersistence) GetListByIds(correlationId string, ids []interface{}) (items []interface{}, err error) {
	params := c.GenerateParameters(ids)
	query := "SELECT * FROM " + c.QuoteTableNameWithSchema() + " WHERE \"id\" IN(" + params + ")"

	qResult, qErr := c.Client.Query(context.TODO(), query, ids...)
	if qErr != nil {
		return nil, qErr
	}
	defer qResult.Close()
	items = make([]interface{}, 0, 0)
	for qResult.Next() {

		item := c.Overrides.ConvertToPublic(qResult)
		items = append(items, item)
	}

	if items != nil {
		c.Logger.Trace(correlationId, "Retrieved %d from %s", len(items), c.TableName)
	}

	return items, qResult.Err()
}

// Gets a data item by its unique id.
//   - correlationId     (optional) transaction id to trace execution through call chain.
//   - id                an id of data item to be retrieved.
// Returns           data item or error.
func (c *IdentifiablePostgresPersistence) GetOneById(correlationId string, id interface{}) (item interface{}, err error) {

	query := "SELECT * FROM " + c.QuoteTableNameWithSchema() + " WHERE \"id\"=$1"

	qResult, qErr := c.Client.Query(context.TODO(), query, id)
	if qErr != nil {
		return nil, qErr
	}
	defer qResult.Close()
	if !qResult.Next() {
		return nil, qResult.Err()
	}
	rows, vErr := qResult.Values()
	if vErr == nil && len(rows) > 0 {
		result := c.Overrides.ConvertToPublic(qResult)
		if result == nil {
			c.Logger.Trace(correlationId, "Nothing found from %s with id = %s", c.TableName, id)
		} else {
			c.Logger.Trace(correlationId, "Retrieved from %s with id = %s", c.TableName, id)
		}
		return result, nil
	}
	return nil, vErr
}

// Creates a data item.
//   - correlation_id    (optional) transaction id to trace execution through call chain.
//   - item              an item to be created.
// Returns          (optional)  created item or error.
func (c *IdentifiablePostgresPersistence) Create(correlationId string, item interface{}) (result interface{}, err error) {
	if item == nil {
		return nil, nil
	}
	// Assign unique id
	var newItem interface{}
	newItem = cmpersist.CloneObject(item, c.Prototype)
	cmpersist.GenerateObjectId(&newItem)

	return c.PostgresPersistence.Create(correlationId, newItem)
}

// Sets a data item. If the data item exists it updates it,
// otherwise it create a new data item.
//   - correlation_id    (optional) transaction id to trace execution through call chain.
//   - item              a item to be set.
// Returns          (optional)  updated item or error.
func (c *IdentifiablePostgresPersistence) Set(correlationId string, item interface{}) (result interface{}, err error) {

	if item == nil {
		return nil, nil
	}

	// Assign unique id
	var newItem interface{}
	newItem = cmpersist.CloneObject(item, c.Prototype)
	cmpersist.GenerateObjectId(&newItem)

	row := c.Overrides.ConvertFromPublic(item)
	params := c.GenerateParameters(row)
	setParams, columns := c.GenerateSetParameters(row)
	values := c.GenerateValues(columns, row)
	id := cmpersist.GetObjectId(newItem)

	query := "INSERT INTO " + c.QuoteTableNameWithSchema() + " (" + columns + ")" +
		" VALUES (" + params + ")" +
		" ON CONFLICT (\"id\") DO UPDATE SET " + setParams + " RETURNING *"

	qResult, qErr := c.Client.Query(context.TODO(), query, values...)
	if qErr != nil {
		return nil, qErr
	}
	defer qResult.Close()

	if !qResult.Next() {
		return nil, qResult.Err()
	}
	rows, vErr := qResult.Values()
	if vErr == nil && len(rows) > 0 {
		result := c.Overrides.ConvertToPublic(qResult)
		c.Logger.Trace(correlationId, "Set in %s with id = %s", c.TableName, id)
		return result, nil
	}
	return nil, vErr

}

// Updates a data item.
//   - correlation_id    (optional) transaction id to trace execution through call chain.
//   - item              an item to be updated.
// Returns          (optional)  updated item or error.
func (c *IdentifiablePostgresPersistence) Update(correlationId string, item interface{}) (result interface{}, err error) {

	if item == nil {
		return nil, nil
	}
	var newItem interface{}
	newItem = cmpersist.CloneObject(item, c.Prototype)
	id := cmpersist.GetObjectId(newItem)

	row := c.Overrides.ConvertFromPublic(newItem)
	params, col := c.GenerateSetParameters(row)
	values := c.GenerateValues(col, row)
	values = append(values, id)

	query := "UPDATE " + c.QuoteTableNameWithSchema() +
		" SET " + params + " WHERE \"id\"=$" + strconv.FormatInt((int64)(len(values)), 16) + " RETURNING *"

	qResult, qErr := c.Client.Query(context.TODO(), query, values...)

	if qErr != nil {
		return nil, qErr
	}
	defer qResult.Close()
	if !qResult.Next() {
		return nil, qResult.Err()
	}
	rows, vErr := qResult.Values()
	if vErr == nil && len(rows) > 0 {
		result := c.Overrides.ConvertToPublic(qResult)
		c.Logger.Trace(correlationId, "Updated in %s with id = %s", c.TableName, id)
		return result, nil
	}
	return nil, vErr
}

// Updates only few selected fields in a data item.
//   - correlation_id    (optional) transaction id to trace execution through call chain.
//   - id                an id of data item to be updated.
//   - data              a map with fields to be updated.
// Returns           updated item or error.
func (c *IdentifiablePostgresPersistence) UpdatePartially(correlationId string, id interface{}, data *cdata.AnyValueMap) (result interface{}, err error) {

	if id == nil {
		return nil, nil
	}

	row := c.Overrides.ConvertFromPublicPartial(data.Value())
	params, col := c.GenerateSetParameters(row)
	values := c.GenerateValues(col, row)
	values = append(values, id)

	query := "UPDATE " + c.QuoteTableNameWithSchema() +
		" SET " + params + " WHERE \"id\"=$" + strconv.FormatInt((int64)(len(values)), 16) + " RETURNING *"

	qResult, qErr := c.Client.Query(context.TODO(), query, values...)

	if qErr != nil {
		return nil, qErr
	}
	defer qResult.Close()
	if !qResult.Next() {
		return nil, qResult.Err()
	}
	rows, vErr := qResult.Values()
	if vErr == nil && len(rows) > 0 {
		result := c.Overrides.ConvertToPublic(qResult)
		c.Logger.Trace(correlationId, "Updated partially in %s with id = %s", c.TableName, id)
		return result, nil
	}
	return nil, vErr
}

// Deleted a data item by it's unique id.
//   - correlation_id    (optional) transaction id to trace execution through call chain.
//   - id                an id of the item to be deleted
// Returns          (optional)  deleted item or error.
func (c *IdentifiablePostgresPersistence) DeleteById(correlationId string, id interface{}) (result interface{}, err error) {

	query := "DELETE FROM " + c.QuoteTableNameWithSchema() + " WHERE \"id\"=$1 RETURNING *"

	qResult, qErr := c.Client.Query(context.TODO(), query, id)

	if qErr != nil {
		return nil, qErr
	}
	defer qResult.Close()
	if !qResult.Next() {
		return nil, qResult.Err()
	}
	rows, vErr := qResult.Values()
	if vErr == nil && len(rows) > 0 {
		result = c.Overrides.ConvertToPublic(qResult)
		c.Logger.Trace(correlationId, "Deleted from %s with id = %s", c.TableName, id)
		return result, nil
	}
	return nil, vErr
}

// Deletes multiple data items by their unique ids.
//   - correlationId     (optional) transaction id to trace execution through call chain.
//   - ids               ids of data items to be deleted.
// Returns          (optional)  error or null for success.
func (c *IdentifiablePostgresPersistence) DeleteByIds(correlationId string, ids []interface{}) error {

	params := c.GenerateParameters(ids)
	query := "DELETE FROM " + c.QuoteTableNameWithSchema() + " WHERE \"id\" IN(" + params + ")"

	qResult, qErr := c.Client.Query(context.TODO(), query, ids...)

	if qErr != nil {
		return qErr
	}
	defer qResult.Close()
	if !qResult.Next() {
		return qResult.Err()
	}
	var count int64 = 0
	rows, vErr := qResult.Values()
	if vErr == nil && len(rows) == 1 {
		count = cconv.LongConverter.ToLong(rows[0])
		if count != 0 {
			c.Logger.Trace(correlationId, "Deleted %d items from %s", count, c.TableName)
		}
	}
	return vErr
}
