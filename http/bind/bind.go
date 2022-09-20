package bind

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
)

func Struct(r *http.Request, i interface{}) error {
	v := reflect.ValueOf(i).Elem()
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)
		tag := fieldType.Tag

		// json获取
		jsonName := tag.Get("json")
		s := r.FormValue(jsonName)

		// 正则校验
		pattern := tag.Get("pattern")
		if pattern != "" {
			match, err := regexp.Match(pattern, []byte(s))
			if err != nil {
				return err
			}
			if !match {
				return errors.New(fmt.Sprint(fieldType.Name, " param is regexp match ", pattern))
			}
		}

		err := setValue(s, field)
		if err != nil {
			return err
		}
	}

	return nil
}

var bitMap = map[reflect.Kind]int{
	reflect.Int:     32,
	reflect.Int16:   16,
	reflect.Int32:   32,
	reflect.Int64:   64,
	reflect.Int8:    8,
	reflect.Uint:    32,
	reflect.Uint16:  16,
	reflect.Uint32:  32,
	reflect.Uint64:  64,
	reflect.Uint8:   8,
	reflect.Float32: 32,
	reflect.Float64: 64,
}

func setValue(data string, v reflect.Value) error {
	kind := v.Kind()
	switch kind {
	case reflect.Int64, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int:
		var d int64
		d, err := strconv.ParseInt(data, 10, bitMap[kind])
		if err != nil {
			return err
		}
		v.SetInt(d)
	case reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint8:
		var d uint64
		d, err := strconv.ParseUint(data, 10, bitMap[kind])
		if err != nil {
			return err
		}
		v.SetUint(d)
	case reflect.Float64, reflect.Float32:
		var d float64
		d, err := strconv.ParseFloat(data, bitMap[kind])
		if err != nil {
			return err
		}
		v.SetFloat(d)
	case reflect.String:
		if data == "" {
			return nil
		}
		v.SetString(data)
		return nil
	}
	return nil
}
