package main

import (
	"fmt"
	"github.com/liufee/go-lib/compress"
	"github.com/liufee/go-lib/db"
	"github.com/liufee/go-lib/excel"
	textFile "github.com/liufee/go-lib/file/text"
	"github.com/liufee/go-lib/mail"
	"github.com/liufee/go-lib/utils/crypt"
	"github.com/liufee/go-lib/utils/json"
	"github.com/liufee/go-lib/utils/rand"
)

func main() {
	//dbTry()
	//mailTry()
}

func dbTry() {
	c := db.NewConnection()
	c.Open("root", "123456", "127.0.0.1", 3306, "cms")              //"root:123456@tcp(127.0.0.1:3306)/cms"
	data := c.Query("select id,name,value from options").FetchAll() //Fetch
	for i, value := range data {
		fmt.Print(i, value)
		fmt.Print("\n\n")
	}
}

func excelTry() {

	data := [][]string{{"1", "11", "111"}, {"2", "22", "222"}}

	ex := excel.NewExcel("a.xlsx", "sheet1")

	columns := []string{"id", "username", "password"}

	ex.SetColumns(columns).SetData(data).Save()
}

func utilTry() {
	data := map[string]interface{}{"username": "fff", "familyname": "liu"}
	var d string
	d = json.JsonEncode(data)

	d = crypt.Md5("123456")
	fmt.Print(d, "\n\n\n")
}

func fileTry() {
	f := textFile.NewFile("sss.txt")
	//f.Write("fucj youyou \n sdg", 0777)
	//f.Write("sssssssss\n ccc", 0777)
	s := f.Read()
	fmt.Print(s)
}

func compressTry() {
	compress.Unzip("feehicms.zip", "fee")
}

func mailTry() {
	ok, err := mail.SendMail("smtp.qiye.aliyun.com", 465, "postmaster@mail.feehi.com", "xxx", "postmaster@mail.feehi.com", "job@feehi.com", "subjecttt", "sssss111222", "text/html")
	if ok == false {
		fmt.Println(err)
	}
}

func randomString() {
	fmt.Println(rand.GenerateRandomString(6))
}
