package conv

import (
	"errors"
	"reflect"
)

var ErrNotArray = errors.New("only support array")

// ToSlice conv an array-interface to []interface{}
// will error if the type is not slice
func ToSlice(array interface{}) ([]interface{}, error) {
	v := reflect.ValueOf(array)
	if v.Kind() != reflect.Slice {
		return nil, ErrNotArray
	}
	l := v.Len()
	ret := make([]interface{}, l)
	for i := 0; i < l; i++ {
		ret[i] = v.Index(i).Interface()
	}
	return ret, nil
}
