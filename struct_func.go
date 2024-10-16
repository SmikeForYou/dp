package dp

import (
	"fmt"
	"reflect"
)

// ToMap converts a struct to a map using the struct's tags.
//
// ToMap uses tags on struct fields to decide which fields to add to the
// returned map.
func StructToMap(in any, tag string) (map[string]any, error) {
	out := make(map[string]any)

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// we only accept structs
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("ToMap only accepts structs; got %T", v)
	}

	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		// gets us a StructField
		fi := typ.Field(i)
		if tagv := fi.Tag.Get(tag); tagv != "" {
			// set key of map to value in struct field
			out[tagv] = v.Field(i).Interface()
		}
	}
	return out, nil
}

// ToArr converts a struct to a slice of values of struct field.
//
// ToArr uses tags on struct fields to decide which fields to add to the
// returned slice.
func StructToArr(in any, tag string) ([]any, error) {
	out := make([]any, 0)
	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("ToArr only accepts structs; got %T", v)
	}
	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		// gets us a StructField
		fi := typ.Field(i)
		if tagv := fi.Tag.Get(tag); tagv != "" {
			// set key of map to value in struct field
			out = append(out, v.Field(i).Interface())
		}
	}
	return out, nil
}

// GetFieldByTag returns the struct field and value of the field with the given tag.
// If tagValue is empty, the first field with the given tag is returned.
func GetFieldByTag(in any, tag string, tagValue string) (reflect.StructField, reflect.Value, error) {
	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return reflect.StructField{}, reflect.Value{}, fmt.Errorf("GetFieldByTag only accepts structs; got %T", v)
	}
	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fi := typ.Field(i)
		if tagv := fi.Tag.Get(tag); tagv != "" {
			if tagValue != "" {
				if tagv == tagValue {
					return fi, v.Field(i), nil
				}
			} else {
				return fi, v.Field(i), nil
			}
		}
	}
	return reflect.StructField{}, reflect.Value{}, fmt.Errorf("GetFieldByTag: tag not found")
}

// DerefArray converts array of pointers to array of values
func DerefArray[T any](arr []*T) []T {
	res := make([]T, 0)
	for _, i := range arr {
		res = append(res, *i)
	}
	return res
}

// RefArray converts array of values to array of pointers
func RefArray[T any](arr []T) []*T {
	res := make([]*T, 0)
	for _, i := range arr {
		var v = i
		res = append(res, &v)
	}
	return res
}
