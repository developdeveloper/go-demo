package tag

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"reflect"
	"testing"
)

func Test_recordDump(t *testing.T) {
	rec := record{"zhangsan", 20, 98.5}
	buf, err := json.Marshal(rec)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(buf)) // {"Name":"zhangsan","Age":20,"Score":98.5}
}

func Test_recordLoad(t *testing.T) {
	rec := record{}
	// str := `{"nickname":"zhangsan","age":20,"score":98.5}`
	str := `{
		"nickname":"zhangsan",
		"age":20,
		"score":98.5
	}`
	if err := json.Unmarshal([]byte(str), &rec); err != nil {
		t.Fatal(err)
	}

	fmt.Println(rec) // {zhangsan 20 98.5}
}

func Test_xmlDump(t *testing.T) {
	rec := record{"zhangsan", 20, 98.5}
	buf, _ := xml.Marshal(rec)
	fmt.Println(string(buf)) // <record nickname="zhangsan" age="20" score="98.5"></record>

	newRec := record{}
	_ = xml.Unmarshal([]byte(buf), &newRec)
	if reflect.DeepEqual(rec, newRec) {
		fmt.Println("dump and load ok")
	}
}
