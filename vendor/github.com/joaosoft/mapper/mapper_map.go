package mapper

import (
	"fmt"
	"reflect"
	"strings"
)

// Map ...
func (mapper *Mapper) Map(value interface{}) (map[string]interface{}, error) {
	mapping := make(map[string]interface{})

	if _, err := convertToMap(value, "", mapping, true); err != nil {
		return nil, err
	}

	return mapping, nil
}

func convertToMap(obj interface{}, path string, mapping map[string]interface{}, add bool) (string, error) {
	types := reflect.TypeOf(obj)
	value := reflect.ValueOf(obj)

	if !value.CanInterface() {
		return "", nil
	}

	if value.Kind() == reflect.Ptr && !value.IsNil() {
		value = value.Elem()

		if value.IsValid() {
			types = value.Type()
		} else {
			return "", nil
		}
	}

	switch value.Kind() {
	case reflect.Struct:

		if !value.CanInterface() {
			return "", nil
		}

		path = addPoint(path)
		var innerPath string
		var len = value.NumField()
		for i := 0; i < types.NumField(); i++ {
			len--
			nextValue := value.Field(i)

			if !nextValue.CanInterface() {
				continue
			}

			jsonName, exists := types.Field(i).Tag.Lookup("json")

			var name string
			if exists {
				split := strings.SplitN(jsonName, ",", 2)
				name = split[0]
			} else {
				name = types.Field(i).Name
			}

			newPath := fmt.Sprintf("%s%s", path, name)
			tmp, _ := convertToMap(nextValue.Interface(), newPath, mapping, add)
			if len > 0 {
				tmp += ","
			}
			innerPath += fmt.Sprintf("%s", tmp)
		}

		if !add {
			return innerPath, nil
		}

	case reflect.Array, reflect.Slice:
		path = addPoint(path)
		var innerPath string
		var len = value.Len()
		for i := 0; i < value.Len(); i++ {
			len--
			nextValue := value.Index(i)

			if !nextValue.CanInterface() {
				continue
			}

			newPath := fmt.Sprintf("%s[%d]", path, i)
			tmp, _ := convertToMap(nextValue.Interface(), newPath, mapping, add)
			if len > 0 {
				tmp += ","
			}
			innerPath += fmt.Sprintf("%s", tmp)
		}

		if !add {
			return innerPath, nil
		}

	case reflect.Map:
		path = addPoint(path)
		var innerPath string
		var len = value.Len()
		for _, key := range value.MapKeys() {
			len--
			nextValue := value.MapIndex(key)

			if !nextValue.CanInterface() {
				continue
			}

			newPath := fmt.Sprintf("%s{", path)
			keyValue, _ := convertToMap(key.Interface(), "", mapping, false)
			newPath += fmt.Sprintf("%s}", keyValue)
			tmp, _ := convertToMap(nextValue.Interface(), newPath, mapping, add)
			if len > 0 {
				tmp += ","
			}
			innerPath += fmt.Sprintf("%s", tmp)
		}

		if !add {
			return innerPath, nil
		}

	default:
		if value.IsValid() {
			var rtnValue interface{}
			if value.CanInterface() {
				rtnValue = value.Interface()
				log.Debugf(fmt.Sprintf("%s=%+v", path, value.Interface()))
			} else {
				rtnValue = value
				log.Debugf(fmt.Sprintf("%s=%+v", path, value))
			}

			if add {
				mapping[path] = rtnValue
			} else {
				if path != "" {
					path += "="
				}
				return fmt.Sprintf("%s%+v", path, rtnValue), nil
			}
		}
	}
	return "", nil
}

func addPoint(path string) string {
	if path != "" {
		path += "."
	}
	return path
}
