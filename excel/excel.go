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
	columns []string
}

func (t *Excel) SetColumns(columns []string) *Excel {

	row := t.sheet.AddRow()
	for _,value := range columns {
		cell := row.AddCell()
		cell.Value = value
	}
	t.columns = columns
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

func (t *Excel) SetDataFromDb(results []map[string]interface{}) *Excel{
	data := make([][]string, len(results))
	i := 0
	for _,v := range results{
		j := 0
		rows := make([]string, len(v))
		for _,column := range t.columns{
			rows[j] = ( v[column] ).(string)
			j++;
		}
		data[i] = rows
		i++
	}
	return t
}

func (t *Excel) Save() string {
	t.file.Save( t.filename )
	return t.filename
}
