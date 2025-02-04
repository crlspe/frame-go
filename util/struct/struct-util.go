package Struct

import (
	"reflect"
	"strings"

	Array "github.com/crlspe/frame-go/util/array"
)

func ToMap(structType any, readTag string, excludeKeyTag string, includeExcludedKeys ...string) map[string]any {
	var structMap = make(map[string]any)

	var _struct = reflect.ValueOf(structType).Elem()
	for i := 0; i < _struct.NumField(); i++ {
		var key = _struct.Type().Field(i)
		var value = _struct.Field(i)

		var keyTags = key.Tag.Get(readTag)
		if keyTags == "" {
			continue
		}
		if strings.Contains(keyTags, excludeKeyTag) && !Array.Contains(includeExcludedKeys, key.Name) {
			continue
		}

		structMap[key.Name] = value.Interface()
	}
	return structMap
}
