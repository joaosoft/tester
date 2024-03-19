package mapper

import (
	"fmt"
	"reflect"
	"strings"
)

// String ...
func (mapper *Mapper) String(value interface{}) (string, error) {
	var spaces int
	var print string

	if err := convertToString(value, "", spaces, "\n", &print); err != nil {
		return "", err
	}

	return print, nil
}

func convertToString(obj interface{}, path string, spaces int, breaker string, print *string) error {
	types := reflect.TypeOf(obj)
	value := reflect.ValueOf(obj)

	if !value.CanInterface() {
		return nil
	}

	if value.Kind() == reflect.Ptr && !value.IsNil() {
		value = value.Elem()

		if value.IsValid() {
			types = value.Type()
		} else {
			return nil
		}
	}

	switch value.Kind() {
	case reflect.Struct:
		for i := 0; i < types.NumField(); i++ {
			nextValue := value.Field(i)

			if !nextValue.CanInterface() {
				continue
			}

			newPath := fmt.Sprintf("%s%s", path, types.Field(i).Name)
			*print += fmt.Sprintf("%s%s%s", breaker, strings.Repeat(" ", spaces), types.Field(i).Name)
			convertToString(nextValue.Interface(), newPath, spaces+2, breaker, print)
		}

	case reflect.Array, reflect.Slice:
		for i := 0; i < value.Len(); i++ {
			nextValue := value.Index(i)

			if !nextValue.CanInterface() {
				continue
			}

			newPath := fmt.Sprintf("[%d]", i)
			*print += fmt.Sprintf("%s%s%s", breaker, strings.Repeat(" ", spaces), newPath)
			convertToString(nextValue.Interface(), newPath, spaces+2, breaker, print)
		}

	case reflect.Map:
		for _, key := range value.MapKeys() {
			var keyValue string
			nextValue := value.MapIndex(key)

			if !nextValue.CanInterface() {
				continue
			}

			convertToString(key.Interface(), "", 0, " ", &keyValue)
			newPath := fmt.Sprintf("{%s}", strings.Trim(keyValue, " "))
			*print += fmt.Sprintf("%s%s%s", breaker, strings.Repeat(" ", spaces), newPath)
			convertToString(nextValue.Interface(), newPath, spaces+2, breaker, print)
		}

	default:
		var rtnValue interface{}
		if value.IsValid() {
			if value.CanInterface() {
				rtnValue = value.Interface()
			} else {
				rtnValue = value
			}

			if path != "" {
				*print += ": "
			}

			newPath := fmt.Sprintf("%s=%+v", path, rtnValue)
			*print += fmt.Sprintf("%+v", rtnValue)

			log.Debugf(fmt.Sprintf("%s", newPath))
		}
	}
	return nil
}
