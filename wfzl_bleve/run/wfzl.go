package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/blevesearch/bleve"
	"github.com/yourhe/gojieba"
	_ "github.com/yourhe/gojieba/bleve"
)

var indexMapping *bleve.IndexMapping

func main() {

	// fmt.Print()

}

func init() {
	// 定义bleve分词 －> gojieba(中文分词)
	indexMapping = bleve.NewIndexMapping()
	err := indexMapping.AddCustomTokenizer("gojieba",
		map[string]interface{}{
			"dictpath":     gojieba.DICT_PATH,
			"hmmpath":      gojieba.HMM_PATH,
			"userdictpath": gojieba.USER_DICT_PATH,
			"type":         "gojieba",
		},
	)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	err = indexMapping.AddCustomAnalyzer("gojieba",
		map[string]interface{}{
			"type":      "gojieba",
			"tokenizer": "gojieba",
		},
	)
	if err != nil {
		fmt.Println(err)

		panic(err)
	}
	indexMapping.DefaultAnalyzer = "gojieba"
}

func batch(i bleve.Index) error {
	file, err := os.Open("/Users/yorhe/Downloads/发给陈煜熙做平台20160401-1/zl/zl.xml")
	if err != nil {
		fmt.Printf("%s\n", err)
		return err
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("%s\n", err)
		return err
	}
	v := Records{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("%s\n", err)
		return err
	}

	for _, reocrd := range v.Records {
		fmt.Printf("%s", record)
	}
	return nil
}
