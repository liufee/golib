package main

import (
	"./db"
	"fmt"
	"./excel"
	"./utils/json"
	"./utils/crypt"
	textFile "./file/text"
	"./compress"
)

func main()  {
	excelTry()
}

func dbTry()  {
	c := db.NewConnection()
	c.Open("root", "123456", "127.0.0.1", 3306, "cms")//"root:123456@tcp(127.0.0.1:3306)/cms"
	data := c.QueryOne("select id,name,value from options");
	for i,value := range data{
		fmt.Print(i,value)
		fmt.Print("\n\n")
	}
}

func excelTry(){

	data := [][]string{{"1","11","111"}, {"2","22","222"}}

	ex := excel.NewExcel("a.xlsx", "sheet1")

	columns := []string{"id", "username", "password"}

	ex.SetColumns(columns).SetData(data).Save()
}

func utilTry()  {
	data := map[string]interface{}{"username":"fff","familyname":"liu"}
	var d string
	d = json.JsonEncode(data)

	d = crypt.Md5("123456")
	fmt.Print(d,"\n\n\n")
}

func fileTry(){
	f := textFile.NewFile("sss.txt")
	//f.Write("fucj youyou \n sdg", 0777)
	//f.Write("sssssssss\n ccc", 0777)
	s := f.Read()
	fmt.Print(s)
}

func compressTry() {
	compress.Unzip("feehicms.zip", "fee")
}