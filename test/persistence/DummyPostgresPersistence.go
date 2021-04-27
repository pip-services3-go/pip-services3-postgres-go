package test

import (
	"reflect"

	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
	persist "github.com/pip-services3-go/pip-services3-postgres-go/persistence"
	tf "github.com/pip-services3-go/pip-services3-postgres-go/test/fixtures"
)

type DummyPostgresPersistence struct {
	persist.IdentifiablePostgresPersistence
}

func NewDummyPostgresPersistence() *DummyPostgresPersistence {
	proto := reflect.TypeOf(tf.Dummy{})
	c := &DummyPostgresPersistence{}
	c.IdentifiablePostgresPersistence = *persist.InheritIdentifiablePostgresPersistence(c, proto, "dummies")
	return c
}

func (c *DummyPostgresPersistence) DefineSchema() {
	c.ClearSchema()
	c.IdentifiablePostgresPersistence.DefineSchema()
	// Row name must be in double quotes for properly case!!!
	c.EnsureSchema("CREATE TABLE " + c.QuoteTableNameWithSchema() + " (\"id\" TEXT PRIMARY KEY, \"key\" TEXT, \"content\" TEXT)")
	c.EnsureIndex(c.TableName+"_key", map[string]string{"key": "1"}, map[string]string{"unique": "true"})
}

func (c *DummyPostgresPersistence) Create(correlationId string, item tf.Dummy) (result tf.Dummy, err error) {
	value, err := c.IdentifiablePostgresPersistence.Create(correlationId, item)

	if value != nil {
		val, _ := value.(tf.Dummy)
		result = val
	}
	return result, err
}

func (c *DummyPostgresPersistence) GetListByIds(correlationId string, ids []string) (items []tf.Dummy, err error) {
	convIds := make([]interface{}, len(ids))
	for i, v := range ids {
		convIds[i] = v
	}
	result, err := c.IdentifiablePostgresPersistence.GetListByIds(correlationId, convIds)
	items = make([]tf.Dummy, len(result))
	for i, v := range result {
		val, _ := v.(tf.Dummy)
		items[i] = val
	}
	return items, err
}

func (c *DummyPostgresPersistence) GetOneById(correlationId string, id string) (item tf.Dummy, err error) {
	result, err := c.IdentifiablePostgresPersistence.GetOneById(correlationId, id)
	if result != nil {
		val, _ := result.(tf.Dummy)
		item = val
	}
	return item, err
}

func (c *DummyPostgresPersistence) Update(correlationId string, item tf.Dummy) (result tf.Dummy, err error) {
	value, err := c.IdentifiablePostgresPersistence.Update(correlationId, item)
	if value != nil {
		val, _ := value.(tf.Dummy)
		result = val
	}
	return result, err
}

func (c *DummyPostgresPersistence) Set(correlationId string, item tf.Dummy) (result tf.Dummy, err error) {
	value, err := c.IdentifiablePostgresPersistence.Set(correlationId, item)
	if value != nil {
		val, _ := value.(tf.Dummy)
		result = val
	}
	return result, err
}

func (c *DummyPostgresPersistence) UpdatePartially(correlationId string, id string, data *cdata.AnyValueMap) (item tf.Dummy, err error) {
	result, err := c.IdentifiablePostgresPersistence.UpdatePartially(correlationId, id, data)

	if result != nil {
		val, _ := result.(tf.Dummy)
		item = val
	}
	return item, err
}

func (c *DummyPostgresPersistence) DeleteById(correlationId string, id string) (item tf.Dummy, err error) {
	result, err := c.IdentifiablePostgresPersistence.DeleteById(correlationId, id)
	if result != nil {
		val, _ := result.(tf.Dummy)
		item = val
	}
	return item, err
}

func (c *DummyPostgresPersistence) DeleteByIds(correlationId string, ids []string) (err error) {
	convIds := make([]interface{}, len(ids))
	for i, v := range ids {
		convIds[i] = v
	}
	return c.IdentifiablePostgresPersistence.DeleteByIds(correlationId, convIds)
}

func (c *DummyPostgresPersistence) GetPageByFilter(correlationId string, filter *cdata.FilterParams, paging *cdata.PagingParams) (page *tf.DummyPage, err error) {

	if &filter == nil {
		filter = cdata.NewEmptyFilterParams()
	}

	key := filter.GetAsNullableString("Key")
	filterObj := ""
	if key != nil && *key != "" {
		filterObj += "key='" + *key + "'"
	}
	sorting := ""

	tempPage, err := c.IdentifiablePostgresPersistence.GetPageByFilter(correlationId,
		filterObj, paging,
		sorting, nil)
	// Convert to DummyPage
	dataLen := int64(len(tempPage.Data)) // For full release tempPage and delete this by GC
	data := make([]tf.Dummy, dataLen)
	for i, v := range tempPage.Data {
		data[i] = v.(tf.Dummy)
	}
	page = tf.NewDummyPage(&dataLen, data)
	return page, err
}

func (c *DummyPostgresPersistence) GetCountByFilter(correlationId string, filter *cdata.FilterParams) (count int64, err error) {

	if &filter == nil {
		filter = cdata.NewEmptyFilterParams()
	}

	key := filter.GetAsNullableString("Key")
	filterObj := ""
	if key != nil && *key != "" {
		filterObj += "key='" + *key + "'"
	}
	return c.IdentifiablePostgresPersistence.GetCountByFilter(correlationId, filterObj)
}
