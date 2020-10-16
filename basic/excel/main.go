package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func main() {
	f := excelize.NewFile()
	// create a new sheet
	index := f.NewSheet("Sheet1")
	// set values of a cell
	f.SetCellValue("sheet1", "A1", 10)
	f.SetCellValue("sheet1", "B2", 20)
	f.SetActiveSheet(index)
	// save xlsx file byte given path
	if err := f.SaveAs("S1.xlsx"); err != nil {
		fmt.Println(err)
	}
}
