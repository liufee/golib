Go Library
===============================

对go基础库、常用三方库的二次封装，使用起来更方便。如导出Excel文件
```go
    excel.NewExcel("a.xlsx", "sheet1").SetColumns(columns).SetData(data).Save()
```



安装
---------------
```bash
    go get github.com/liufee/go-lib
    go get github.com/tealeg/xlsx
    go get github.com/go-sql-driver/mysql
```


api
-------
* compress zip解压缩
* db 数据库sql查询，自动返回select字段的array map
* excel 导出excel
* file/text 文本文件读写
* utils crypt加密、json(反)序列化


示例
------
[main.go](main.go)