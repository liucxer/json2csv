package json2csv_test

import (
	"fmt"
	"github.com/liucxer/json2csv"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCsv_String(t *testing.T) {
	type pointStruct struct {
		PointItem string `json:"pointItem" title:"选项"`
	}

	type subStruct struct {
		SubStructItem string `json:"subStructItem" title:"子选项"`
	}

	type Person struct {
		*pointStruct
		SubStruct subStruct `json:"subStruct" title:"子结构体"`
		Name      string    `json:"name" title:"姓名"`
		Age       int64     `json:"age" title:"年龄"`
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

	csv, err = json2csv.ToCsv([]Person{person, person})
	require.NoError(t, err)
	fmt.Println(csv)
}
