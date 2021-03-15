package excel

import (
	"github.com/tealeg/xlsx"
)

// 读取excel
func ReadExcel(path string, sheetIndex int) ([][]string, error) {
	data, err := xlsx.FileToSlice(path)
	if err != nil {
		return nil, err
	}
	sheetData := data[sheetIndex]
	//for rowIndex, row := range sheetData {
	//	for cellIndex, cell := range row {
	//		fmt.Println(fmt.Sprintf("第%d行，第%d个单元格：%s", rowIndex+1, cellIndex+1, cell))
	//	}
	//}
	return sheetData, nil
}
