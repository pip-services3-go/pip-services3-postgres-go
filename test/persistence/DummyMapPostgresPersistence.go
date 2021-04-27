package test

import (
	"reflect"

	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
	persist "github.com/pip-services3-go/pip-services3-postgres-go/persistence"
	tf "github.com/pip-services3-go/pip-services3-postgres-go/test/fixtures"
)

type DummyMapPostgresPersistence struct {
	persist.IdentifiablePostgresPersistence
}

func NewDummyMapPostgresPersistence() *DummyMapPostgresPersistence {
	var t map[string]interface{}
	proto := reflect.TypeOf(t)
	c := &DummyMapPostgresPersistence{}
	c.IdentifiablePostgresPersistence = *persist.InheritIdentifiablePostgresPersistence(c, proto, "dummies")
	return c
}

func (c *DummyMapPostgresPersistence) DefineSchema() {
	c.ClearSchema()
	c.IdentifiablePostgresPersistence.DefineSchema()
	c.EnsureSchema("CREATE TABLE \"" + c.TableName + "\" (\"id\" TEXT PRIMARY KEY, \"key\" TEXT, \"content\" TEXT)")
	c.EnsureIndex(c.TableName+"_key", map[string]string{"key": "1"}, map[string]string{"unique": "true"})
}

func (c *DummyMapPostgresPersistence) Create(correlationId string, item map[string]interface{}) (result map[string]interface{}, err error) {
	value, err := c.IdentifiablePostgresPersistence.Create(correlationId, item)
	if value != nil {
		val, _ := value.(map[string]interface{})
		result = val
	}
	return result, err
}

func (c *DummyMapPostgresPersistence) GetListByIds(correlationId string, ids []string) (items []map[string]interface{}, err error) {
	convIds := make([]interface{}, len(ids))
	for i, v := range ids {
		convIds[i] = v
	}
	result, err := c.IdentifiablePostgresPersistence.GetListByIds(correlationId, convIds)
	items = make([]map[string]interface{}, len(result))
	for i, v := range result {
		val, _ := v.(map[string]interface{})
		items[i] = val
	}
	return items, err
}

func (c *DummyMapPostgresPersistence) GetOneById(correlationId string, id string) (item map[string]interface{}, err error) {
	result, err := c.IdentifiablePostgresPersistence.GetOneById(correlationId, id)

	if result != nil {
		val, _ := result.(map[string]interface{})
		item = val
	}
	return item, err
}

func (c *DummyMapPostgresPersistence) Update(correlationId string, item map[string]interface{}) (result map[string]interface{}, err error) {
	value, err := c.IdentifiablePostgresPersistence.Update(correlationId, item)

	if value != nil {
		val, _ := value.(map[string]interface{})
		result = val
	}
	return result, err
}

func (c *DummyMapPostgresPersistence) Set(correlationId string, item map[string]interface{}) (result map[string]interface{}, err error) {
	value, err := c.IdentifiablePostgresPersistence.Set(correlationId, item)

	if value != nil {
		val, _ := value.(map[string]interface{})
		result = val
	}
	return result, err
}

func (c *DummyMapPostgresPersistence) UpdatePartially(correlationId string, id string, data *cdata.AnyValueMap) (item map[string]interface{}, err error) {
	result, err := c.IdentifiablePostgresPersistence.UpdatePartially(correlationId, id, data)

	if result != nil {
		val, _ := result.(map[string]interface{})
		item = val
	}
	return item, err
}

func (c *DummyMapPostgresPersistence) DeleteById(correlationId string, id string) (item map[string]interface{}, err error) {
	result, err := c.IdentifiablePostgresPersistence.DeleteById(correlationId, id)

	if result != nil {
		val, _ := result.(map[string]interface{})
		item = val
	}
	return item, err
}

func (c *DummyMapPostgresPersistence) DeleteByIds(correlationId string, ids []string) (err error) {
	convIds := make([]interface{}, len(ids))
	for i, v := range ids {
		convIds[i] = v
	}
	return c.IdentifiablePostgresPersistence.DeleteByIds(correlationId, convIds)
}

func (c *DummyMapPostgresPersistence) GetPageByFilter(correlationId string, filter *cdata.FilterParams, paging *cdata.PagingParams) (page *tf.MapPage, err error) {

	if &filter == nil {
		filter = cdata.NewEmptyFilterParams()
	}

	key := filter.GetAsNullableString("Key")
	filterObj := ""
	if key != nil && *key != "" {
		filterObj += "key='" + *key + "'"
	}
	sorting := ""

	tempPage, err := c.IdentifiablePostgresPersistence.GetPageByFilter(correlationId, filterObj, paging,
		sorting, nil)
	dataLen := int64(len(tempPage.Data))
	data := make([]map[string]interface{}, dataLen)
	for i, v := range tempPage.Data {
		data[i] = v.(map[string]interface{})
	}
	dataPage := tf.NewMapPage(&dataLen, data)
	return dataPage, err
}

func (c *DummyMapPostgresPersistence) GetCountByFilter(correlationId string, filter *cdata.FilterParams) (count int64, err error) {

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
