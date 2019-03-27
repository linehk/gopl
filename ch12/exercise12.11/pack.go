package pack

import (
	"fmt"
	"reflect"
	"strings"
)

func Pack(data interface{}) string {
	v := reflect.ValueOf(data)
	url := "?"
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		tag := field.Tag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(field.Name)
		}
		url += name + "=" + convert(v.Field(i)) + "&"
	}
	url = strings.TrimSuffix(url, "&")
	return url
}

func convert(v reflect.Value) string {
	switch v.Kind() {
	case reflect.String:
		return v.String()
	case reflect.Int:
		return fmt.Sprintf("%d", v.Int())
	case reflect.Bool:
		return fmt.Sprintf("%t", v.Bool())
	default:
		return ""
	}
}
