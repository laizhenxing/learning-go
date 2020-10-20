package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"os"
)

func main() {
	write()
	read()

}

func write()  {
	f := excelize.NewFile()
	// create a new sheet
	index := f.NewSheet("sheet3")
	// set values of a cell
	f.SetCellValue("sheet3", "A2", "hello world")
	f.SetCellValue("sheet3", "A4", "1")
	f.SetCellValue("sheet3", "A3", "3333")
	f.SetCellValue("sheet3", "A5", "3332222")
	f.SetCellValue("sheet3", "A6", "777")
	f.SetCellValue("sheet3", "A7", "test")
	f.SetCellValue("sheet1", "B2", 20)
	f.SetActiveSheet(index-1)
	// save xlsx file byte given path
	if err := f.SaveAs("S1.xlsx"); err != nil {
		fmt.Println(err)
	}
}

func read()  {
	r, err := excelize.OpenFile("S1.xlsx")
	if err != nil {
		fmt.Println("OpenFile err: ", err)
		os.Exit(1)
	}

	cell, err := r.GetCellValue("sheet3", "A1")
	if err != nil {
		fmt.Println("Get Cell Value err: ", err)
		os.Exit(2)
	}
	fmt.Println("Cell Value: ", cell)

	// 获取sheet3上所有的单元
	rows, err := r.GetRows("sheet3")
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}
