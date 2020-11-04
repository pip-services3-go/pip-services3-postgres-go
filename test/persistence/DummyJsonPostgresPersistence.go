package test

import (
	"reflect"

	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
	ppersist "github.com/pip-services3-go/pip-services3-postgres-go/persistence"
	tf "github.com/pip-services3-go/pip-services3-postgres-go/test/fixtures"
)

type DummyJsonPostgresPersistence struct {
	ppersist.IdentifiableJsonPostgresPersistence
}

func NewDummyJsonPostgresPersistence() *DummyJsonPostgresPersistence {
	proto := reflect.TypeOf(tf.Dummy{})
	c := &DummyJsonPostgresPersistence{
		IdentifiableJsonPostgresPersistence: *ppersist.NewIdentifiableJsonPostgresPersistence(proto, "dummies_json"),
	}

	c.EnsureTable("", "")
	c.EnsureIndex("dummies_json_key", map[string]string{"(data->>'key')": "1"}, map[string]string{"unique": "true"})
	return c
}

func (c *DummyJsonPostgresPersistence) GetPageByFilter(correlationId string, filter *cdata.FilterParams, paging *cdata.PagingParams) (page *tf.DummyPage, err error) {

	if &filter == nil {
		filter = cdata.NewEmptyFilterParams()
	}

	key := filter.GetAsNullableString("Key")
	filterObj := ""
	if key != nil && *key != "" {
		filterObj += "data->key='" + *key + "'"
	}

	tempPage, err := c.IdentifiablePostgresPersistence.GetPageByFilter(correlationId,
		filterObj, paging,
		nil, nil)
	// Convert to DummyPage
	dataLen := int64(len(tempPage.Data)) // For full release tempPage and delete this by GC
	data := make([]tf.Dummy, dataLen)
	for i, v := range tempPage.Data {
		data[i] = v.(tf.Dummy)
	}
	page = tf.NewDummyPage(&dataLen, data)
	return page, err
}

func (c *DummyJsonPostgresPersistence) GetCountByFilter(correlationId string, filter *cdata.FilterParams) (count int64, err error) {

	if &filter == nil {
		filter = cdata.NewEmptyFilterParams()
	}

	key := filter.GetAsNullableString("Key")
	filterObj := ""

	if key != nil && *key != "" {
		filterObj += "data->key='" + *key + "'"
	}

	return c.IdentifiablePostgresPersistence.GetCountByFilter(correlationId, filterObj)
}

func (c *DummyJsonPostgresPersistence) Create(correlationId string, item tf.Dummy) (result tf.Dummy, err error) {
	value, err := c.IdentifiablePostgresPersistence.Create(correlationId, item)

	if value != nil {
		val, _ := value.(tf.Dummy)
		result = val
	}
	return result, err
}

func (c *DummyJsonPostgresPersistence) GetListByIds(correlationId string, ids []string) (items []tf.Dummy, err error) {
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

func (c *DummyJsonPostgresPersistence) GetOneById(correlationId string, id string) (item tf.Dummy, err error) {
	result, err := c.IdentifiablePostgresPersistence.GetOneById(correlationId, id)
	if result != nil {
		val, _ := result.(tf.Dummy)
		item = val
	}
	return item, err
}

func (c *DummyJsonPostgresPersistence) Update(correlationId string, item tf.Dummy) (result tf.Dummy, err error) {
	value, err := c.IdentifiablePostgresPersistence.Update(correlationId, item)
	if value != nil {
		val, _ := value.(tf.Dummy)
		result = val
	}
	return result, err
}

func (c *DummyJsonPostgresPersistence) UpdatePartially(correlationId string, id string, data *cdata.AnyValueMap) (item tf.Dummy, err error) {
	// In json persistence this method must call from IdentifiableJsonPostgresPersistence
	result, err := c.IdentifiableJsonPostgresPersistence.UpdatePartially(correlationId, id, data)

	if result != nil {
		val, _ := result.(tf.Dummy)
		item = val
	}
	return item, err
}

func (c *DummyJsonPostgresPersistence) DeleteById(correlationId string, id string) (item tf.Dummy, err error) {
	result, err := c.IdentifiablePostgresPersistence.DeleteById(correlationId, id)
	if result != nil {
		val, _ := result.(tf.Dummy)
		item = val
	}
	return item, err
}

func (c *DummyJsonPostgresPersistence) DeleteByIds(correlationId string, ids []string) (err error) {
	convIds := make([]interface{}, len(ids))
	for i, v := range ids {
		convIds[i] = v
	}
	return c.IdentifiablePostgresPersistence.DeleteByIds(correlationId, convIds)
}
