package model

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestRecords(t *testing.T) {
	file, err := os.Open("/Users/yorhe/Downloads/发给陈煜熙做平台20160401-1/zl/zl.xml")
	if err != nil {
		t.FailNow()
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		t.FailNow()
	}
	v := Records{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("%s\n", err)
		t.FailNow()
	}
	fmt.Print(len(v.Records))
	for _, record := range v.Records {
		switch record.Keys()["F_Class"].(type) {
		case []string:
			fmt.Printf("%s\n\n\n", len(record.Keys()["F_Class"].([]string)))
		default:
			// fmt.Printf("%s\n\n\n", record.Keys()["F_Class"])

		}
		fmt.Printf("%s\n", record.keys)
		// for _, value := range record.Fields {
		//
		// 	fmt.Printf("Key: %d  Value: %s\n", i, value.Value)
		// }
	}

	// fmt.Printf("%s", v.Records[0].Fields)
}
