package conv

import (
	"fmt"
	"strconv"
)

// ToInt64 casts a interface to an int64
func ToInt64(i interface{}) (int64, error) {
	switch s := i.(type) {
	case int:
		return int64(s), nil
	case int64:
		return s, nil
	case string:
		v, err := strconv.ParseInt(s, 0, 64)
		if err == nil {
			return v, nil
		}
		return 0, fmt.Errorf("unable to cast %#v to int64, %w", i, err)
	case float64:
		return int64(s), nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v to int64, unsupported type", i)
	}
}
