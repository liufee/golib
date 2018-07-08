package db

import (
	"database/sql"
	"fmt"
	"os"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

func NewConnection() *connection {
	return &connection{}
}

type connection struct {
	Db     *sql.DB
	Rows *sql.Rows
	columns []string
	data []map[string]interface{}
}

func (c *connection) Open(user string, password string, host string, port int, database string) *connection {
	var dsn = user + ":" + password +"@tcp(" + host + ":" + strconv.Itoa(port) + ")/" + database
	connection, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err);
		os.Exit(1);
	}
	c.Db = connection
	return c
}

func (c *connection) Query(sql string) *connection{
	rows, err := c.Db.Query(sql)
	if err != nil {
		os.Exit(1)
		fmt.Println(err)
	}
	c.Rows = rows
	return c
}

func (c *connection) FetchAll() []map[string]interface{} {
	c.parse()
	return c.data
}

func (c *connection) Fetch() map[string]interface{} {
	c.parse()
	return c.data[0]
}

func (c *connection) GetColumns() []string {
	c.parse()
	return c.columns
}

func (c *connection) parse() {
	if( c.columns != nil ){
		return
	}
	if( c.Rows == nil ){
		fmt.Print("必须先query")
		os.Exit(1)
	}
	columns,_ := c.Rows.Columns();
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for c.Rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		c.Rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	c.data = tableData
	c.columns = columns
}
