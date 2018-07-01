package json

import (
	"encoding/json"
)

func JsonEncode(data map[string]interface{}) string  {
	bytes,_ := json.Marshal(data)
	str:= string(bytes)
	return str
}

func JsonDecode(str string) {

}
