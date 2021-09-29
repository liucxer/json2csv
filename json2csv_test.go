package json2csv_test

import (
	"fmt"
	"github.com/liucxer/json2csv"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCsv_String(t *testing.T) {
	type pointStruct struct {
		PointItem string `json:"pointItem"`
	}

	type subStruct struct {
		SubStructItem string `json:"subStructItem"`
	}

	type Person struct {
		*pointStruct
		SubStruct subStruct `json:"subStruct"`
		Name string `json:"name"`
		Age  int64  `json:"age"`
	}

	person := Person{Name: "zhangsan", Age: 18}
	person.pointStruct = &pointStruct{}
	person.PointItem = "aaa"
	person.SubStruct.SubStructItem = "subStructItem"
	csv, err := json2csv.ToCsv(person)
	require.NoError(t, err)
	fmt.Println(csv)

	csv, err = json2csv.ToCsv(&person)
	require.NoError(t, err)
	fmt.Println(csv)

	csv, err = json2csv.ToCsv([]Person{person,person})
	require.NoError(t, err)
	fmt.Println(csv)
}
