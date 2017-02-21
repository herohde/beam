package reflectx

import (
	"reflect"
	"strings"
)

const (
	KeyTag    string = "key"
	ValueTag  string = "value"
	ValuesTag string = "values"

	// Context binding tags
	DataTag string = "data"
	TypeTag string = "type"
)

// FindTaggedField returns the field tagged with any of the given values, if
// any. The tags are all under the "beam" StructTag key.
func FindTaggedField(t reflect.Type, values ...string) (reflect.StructField, bool) {
	if t == nil || t.Kind() != reflect.Struct {
		return reflect.StructField{}, false
	}

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		list := strings.Split(f.Tag.Get("beam"), ",")
		for _, elm := range list {
			for _, value := range values {
				if elm == value {
					return f, true
				}
			}
		}
	}
	return reflect.StructField{}, false
}