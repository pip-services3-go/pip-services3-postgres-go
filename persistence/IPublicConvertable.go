package persistence

type IPublicConvertable interface {
	//  Convert object value from public to internal format.
	ConvertToPublic(value interface{}) interface{}
	//  Updates only few selected fields in a data item.
	ConvertFromPublic(value interface{}) interface{}
	// Converts the given object from the public partial format.
	ConvertFromPublicPartial(value interface{}) interface{}
}
