package excel

import (
	"fmt"
	"testing"
)

func TestReadExcel(t *testing.T) {
	path := "地区工商股东类型汇总.xlsx"
	sheetIndex := 0
	data, err := ReadExcel(path, sheetIndex)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, row := range data {
		fmt.Println(fmt.Sprintf("u\"%s\",", row[0]))
	}
}
