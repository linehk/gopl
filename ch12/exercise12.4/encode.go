package sexpr

import (
	"bytes"
	"fmt"
	"reflect"
)

// Marshal encodes a Go value in S-expression form.
func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v), 0); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// encode writes to buf an S-expression representation of v.
func encode(buf *bytes.Buffer, v reflect.Value, indent int) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("nil")

	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(buf, "%d", v.Uint())

	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())

	case reflect.Ptr:
		return encode(buf, v.Elem(), indent)

	case reflect.Array, reflect.Slice: // (value ...)
		indent++
		buf.WriteByte('(')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				fmt.Fprintf(buf, "\n%*s", indent, "")
			}
			if err := encode(buf, v.Index(i), indent); err != nil {
				return err
			}
		}
		buf.WriteByte(')')

	case reflect.Struct: // ((name value) ...)
		indent++
		buf.WriteByte('(')
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				fmt.Fprintf(buf, "\n%*s", indent, "")
			}
			s := buf.Len()
			fmt.Fprintf(buf, "(%s ", v.Type().Field(i).Name)
			if err := encode(buf, v.Field(i), indent+buf.Len()-s); err != nil {
				return err
			}
			buf.WriteByte(')')
		}
		buf.WriteByte(')')

	case reflect.Map: // ((key value) ...)
		indent++
		buf.WriteByte('(')
		for i, key := range v.MapKeys() {
			if i > 0 {
				fmt.Fprintf(buf, "\n%*s", indent, "")
			}
			s := buf.Len()
			buf.WriteByte('(')
			if err := encode(buf, key, 0); err != nil {
				return err
			}
			buf.WriteByte(' ')
			if err := encode(buf, v.MapIndex(key), indent+buf.Len()-s); err != nil {
				return err
			}
			buf.WriteByte(')')
		}
		buf.WriteByte(')')

	default: // float, complex, bool, chan, func, interface
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}
