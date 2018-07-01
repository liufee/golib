package excel

import (
	"github.com/tealeg/xlsx"
	"fmt"
	"os"
)

func NewExcel(filename string, sheetName string) *Excel{
	var file *xlsx.File
	var errs error
	var sheet *xlsx.Sheet

	file = xlsx.NewFile()
	sheet, errs = file.AddSheet(sheetName)
	if errs!= nil {
		fmt.Printf(errs.Error())
		os.Exit(1)
	}

	return &Excel{
		filename:filename,
		file:file,
		sheet:sheet,
	}
}

type Excel struct {
	filename string
	file *xlsx.File
	sheet *xlsx.Sheet
}

func (t *Excel) SetColumns(columns []string) *Excel {

	row := t.sheet.AddRow()
	for _,value := range columns {
		cell := row.AddCell()
		cell.Value = value
	}
	return t
}

func (t *Excel) SetData(data [][]string) *Excel {
	for _,value := range data {
		row := t.sheet.AddRow()
		for _,v := range value {
			cell := row.AddCell()
			cell.Value = v
		}
	}
	return t
}

func (t *Excel) Save() string {
	t.file.Save( t.filename )
	return t.filename
}
