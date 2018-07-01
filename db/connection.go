package db

import (
	"database/sql"
	"fmt"
	"os"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

func NewConnection() *Connection  {
	return &Connection{}
}

type Connection struct {
	Con *sql.DB
}

func (c *Connection) Open(user string, password string, host string, port int, database string) {
	var dsn = user + ":" + password +"@tcp(" + host + ":" + strconv.Itoa(port) + ")/" + database
	connection, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err);
		os.Exit(1);
	}
	c.Con = connection
	//time.Sleep(10000000000)
}

func (c *Connection) Query(sql string) []map[string]interface{}{
	rows, err := c.Con.Query(sql)
	if err != nil {
		os.Exit(1)
		fmt.Println(err)
	}
	columns,_ := rows.Columns();
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
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
	return tableData
}

func (c *Connection) QueryOne(sql string) map[string]interface{} {
	tableData := c.Query(sql)
	return tableData[0]
}
