package bleve

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/blevesearch/bleve"
	"github.com/yourhe/gojieba"
)

func TestJiebaCutForSearch(t *testing.T) {
	Example()
}
func Example() {
	INDEX_DIR := "gojieba.bleve"
	messages := []struct {
		Id   string
		Body string
	}{
		{
			Id:   "1",
			Body: "你好",
		},
		{
			Id:   "2",
			Body: "世界",
		},
		{
			Id:   "3",
			Body: "亲口",
		},
		{
			Id:   "4",
			Body: "交代",
		},
		{
			Id:   "5",
			Body: "长江大桥",
		},
	}

	indexMapping := bleve.NewIndexMapping()
	os.RemoveAll(INDEX_DIR)
	// clean index when example finished
	defer os.RemoveAll(INDEX_DIR)

	err := indexMapping.AddCustomTokenizer("gojieba",
		map[string]interface{}{
			"dictpath":     gojieba.DICT_PATH,
			"hmmpath":      gojieba.HMM_PATH,
			"userdictpath": gojieba.USER_DICT_PATH,
			"mode":         1,
			"type":         "gojieba",
		},
	)
	if err != nil {
		panic(err)
	}
	err = indexMapping.AddCustomAnalyzer("gojieba",
		map[string]interface{}{
			"type":      "gojieba",
			"tokenizer": "gojieba",
		},
	)
	if err != nil {
		panic(err)
	}
	indexMapping.DefaultAnalyzer = "gojieba"

	index, err := bleve.New(INDEX_DIR, indexMapping)
	if err != nil {
		panic(err)
	}
	for _, msg := range messages {
		if err := index.Index(msg.Id, msg); err != nil {
			panic(err)
		}
	}

	querys := []string{
		"你好世界",
		"亲口交代",
		"长江大桥",
	}

	for _, q := range querys {
		req := bleve.NewSearchRequest(bleve.NewQueryStringQuery(q))
		req.Highlight = bleve.NewHighlight()
		res, err := index.Search(req)
		if err != nil {
			panic(err)
		}
		fmt.Println(prettify(res))
	}

	// Output:
	// [{"id":"2","score":0.47907267476955906},{"id":"1","score":0.47907267476955906}]
	// [{"id":"4","score":0.47907267476955906},{"id":"3","score":0.47907267476955906}]
	// [{"id":"5","score":0.9581453659370776}]
}

func prettify(res *bleve.SearchResult) string {
	type Result struct {
		Id    string  `json:"id"`
		Score float64 `json:"score"`
	}
	results := []Result{}
	for _, item := range res.Hits {
		results = append(results, Result{item.ID, item.Score})
	}
	b, err := json.Marshal(results)
	if err != nil {
		panic(err)
	}
	return string(b)
}
