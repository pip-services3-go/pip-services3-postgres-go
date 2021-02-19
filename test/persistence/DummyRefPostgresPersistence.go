package test

import (
	"reflect"

	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
	ppersist "github.com/pip-services3-go/pip-services3-postgres-go/persistence"
	tf "github.com/pip-services3-go/pip-services3-postgres-go/test/fixtures"
)

// extends IdentifiablePostgresPersistence<Dummy, string>
// implements IDummyPersistence {
type DummyRefPostgresPersistence struct {
	ppersist.IdentifiablePostgresPersistence
}

func NewDummyRefPostgresPersistence() *DummyRefPostgresPersistence {
	proto := reflect.TypeOf(&tf.Dummy{})
	return &DummyRefPostgresPersistence{*ppersist.NewIdentifiablePostgresPersistence(proto, "dummies")}
}

func (c *DummyRefPostgresPersistence) Create(correlationId string, item *tf.Dummy) (result *tf.Dummy, err error) {
	value, err := c.IdentifiablePostgresPersistence.Create(correlationId, item)

	if value != nil {
		val, _ := value.(*tf.Dummy)
		result = val
	}
	return result, err
}

func (c *DummyRefPostgresPersistence) GetListByIds(correlationId string, ids []string) (items []*tf.Dummy, err error) {
	convIds := make([]interface{}, len(ids))
	for i, v := range ids {
		convIds[i] = v
	}
	result, err := c.IdentifiablePostgresPersistence.GetListByIds(correlationId, convIds)
	items = make([]*tf.Dummy, len(result))
	for i, v := range result {
		val, _ := v.(*tf.Dummy)
		items[i] = val
	}
	return items, err
}

func (c *DummyRefPostgresPersistence) GetOneById(correlationId string, id string) (item *tf.Dummy, err error) {
	result, err := c.IdentifiablePostgresPersistence.GetOneById(correlationId, id)
	if result != nil {
		val, _ := result.(*tf.Dummy)
		item = val
	}
	return item, err
}

func (c *DummyRefPostgresPersistence) Update(correlationId string, item *tf.Dummy) (result *tf.Dummy, err error) {
	value, err := c.IdentifiablePostgresPersistence.Update(correlationId, item)
	if value != nil {
		val, _ := value.(*tf.Dummy)
		result = val
	}
	return result, err
}

func (c *DummyRefPostgresPersistence) UpdatePartially(correlationId string, id string, data *cdata.AnyValueMap) (item *tf.Dummy, err error) {
	result, err := c.IdentifiablePostgresPersistence.UpdatePartially(correlationId, id, data)

	if result != nil {
		val, _ := result.(*tf.Dummy)
		item = val
	}
	return item, err
}

func (c *DummyRefPostgresPersistence) DeleteById(correlationId string, id string) (item *tf.Dummy, err error) {
	result, err := c.IdentifiablePostgresPersistence.DeleteById(correlationId, id)
	if result != nil {
		val, _ := result.(*tf.Dummy)
		item = val
	}
	return item, err
}

func (c *DummyRefPostgresPersistence) DeleteByIds(correlationId string, ids []string) (err error) {
	convIds := make([]interface{}, len(ids))
	for i, v := range ids {
		convIds[i] = v
	}
	return c.IdentifiablePostgresPersistence.DeleteByIds(correlationId, convIds)
}

func (c *DummyRefPostgresPersistence) GetPageByFilter(correlationId string, filter *cdata.FilterParams, paging *cdata.PagingParams) (page *tf.DummyRefPage, err error) {

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
	// Convert to DummyRefPage
	dataLen := int64(len(tempPage.Data)) // For full release tempPage and delete this by GC
	data := make([]*tf.Dummy, dataLen)
	for i := range tempPage.Data {
		temp := tempPage.Data[i].(*tf.Dummy)
		data[i] = temp
	}
	page = tf.NewDummyRefPage(&dataLen, data)
	return page, err
}

func (c *DummyRefPostgresPersistence) GetCountByFilter(correlationId string, filter *cdata.FilterParams) (count int64, err error) {

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
