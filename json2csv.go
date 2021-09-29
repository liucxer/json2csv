package json2csv

import (
	"errors"
	"fmt"
	"reflect"
)

type CsvList struct {
	Title []string      `json:"title"`
	Value [][]interface{} `json:"value"`
}

type Csv struct {
	Title []string      `json:"title"`
	Value []interface{} `json:"value"`
}

func (c *CsvList) String() string {
	res := ""
	name := ""
	for _, v := range c.Title {
		name = name + v + ","
	}
	res = res + name + "\n"

	for _, values := range c.Value {
		resValue := ""
		for _, value := range values {
			resValue = resValue + fmt.Sprintf("%v", value) + ","
		}
		res += resValue + "\n"
	}
	return res
}

func (c *Csv) Append(csv *Csv) {
	c.Title = append(c.Title, csv.Title...)
	c.Value = append(c.Value, csv.Value...)
}

func ToCsv(object interface{}) (*CsvList, error) {
	var (
		csv CsvList
	)

	rv := reflect.ValueOf(object)
	if rv.Kind() != reflect.Struct &&
		rv.Kind() != reflect.Ptr &&
		rv.Kind() != reflect.Slice {
		return nil, errors.New("not support object type")
	}
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	if rv.Kind() == reflect.Struct {
		res, err := dumpStruct("", rv)
		if err != nil {
			return nil, err
		}
		csv.Title = res.Title
		csv.Value = append(csv.Value, res.Value)
	}

	if rv.Kind() == reflect.Slice {
		sliceLen := rv.Len()
		for i := 0; i < sliceLen; i++ {
			subrv := rv.Index(i)
			if subrv.Kind() == reflect.Ptr {
				subrv = subrv.Elem()
			}

			if subrv.Kind() == reflect.Struct {
				res, err := dumpStruct("", subrv)
				if err != nil {
					return nil, err
				}
				csv.Title = res.Title
				csv.Value = append(csv.Value, res.Value)
			}
		}
	}

	return &csv, nil
}

func IsFieldKind(k reflect.Kind) bool {
	fieldKind := []reflect.Kind{
		reflect.Bool,
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Uintptr,
		reflect.Float32,
		reflect.Float64,
		reflect.Complex64,
		reflect.Complex128,
		reflect.Array,
		reflect.String,
	}
	for _, item := range fieldKind {
		if item == k {
			return true
		}
	}
	return false
}

func dumpField(parentName, title string, rv reflect.Value) (*Csv, error) {
	var (
		csv Csv
	)
	if !IsFieldKind(rv.Kind()) {
		return nil, errors.New("only support fieldKind")
	}

	if parentName == "" {
		csv.Title = append(csv.Title, title)
	} else {
		csv.Title = append(csv.Title, parentName+"."+title)
	}

	csv.Value = append(csv.Value, rv.Interface())
	return &csv, nil
}

func dumpStruct(parentName string, rv reflect.Value) (*Csv, error) {
	var (
		csv Csv
	)
	if rv.Kind() != reflect.Struct {
		return nil, errors.New("only support struct")
	}

	numField := rv.NumField()
	rt := rv.Type()
	for i := 0; i < numField; i++ {
		name := rt.Field(i).Name
		frv := rv.Field(i)
		if frv.Kind() == reflect.Ptr {
			frv = frv.Elem()
		}
		if frv.Kind() == reflect.Struct {
			res, err := dumpStruct(name, frv)
			if err != nil {
				return nil, err
			}
			csv.Append(res)
		}

		if IsFieldKind(frv.Kind()) {
			res, err := dumpField(parentName, name, frv)
			if err != nil {
				return nil, err
			}
			csv.Append(res)
		}
	}
	return &csv, nil
}
