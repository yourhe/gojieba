package bleve

import (
	"errors"

	"github.com/blevesearch/bleve/analysis"
	"github.com/blevesearch/bleve/registry"
	"github.com/yourhe/gojieba"
)

// JiebaTokenizer struct
type JiebaTokenizer struct {
	handle *gojieba.Jieba
	mode   gojieba.TokenizeMode
}

// NewJiebaTokenizer should return JiebaTokenizer
func NewJiebaTokenizer(dictpath, hmmpath, userdictpath string, mode gojieba.TokenizeMode) *JiebaTokenizer {
	x := gojieba.NewJieba(dictpath, hmmpath, userdictpath)
	return &JiebaTokenizer{x, mode}
}

// Free relsese JiebaTokenizer
func (x *JiebaTokenizer) Free() {
	x.handle.Free()
}

// Tokenize 分词->返回 token stream
func (x *JiebaTokenizer) Tokenize(sentence []byte) analysis.TokenStream {
	result := make(analysis.TokenStream, 0)
	pos := 1
	words := x.handle.Tokenize(string(sentence), x.mode, false)
	// fmt.Printf("%s\n", string(sentence))
	for _, word := range words {
		token := analysis.Token{
			Term:     []byte(word.Str),
			Start:    word.Start,
			End:      word.End,
			Position: pos,
			Type:     analysis.Ideographic,
		}
		result = append(result, &token)
		pos++
	}
	return result
}

// func (x *JiebaTokenizer) Tokenize(sentence []byte) analysis.TokenStream {
// 	// result := make(analysis.TokenStream, 0)
// 	// pos := 1
// 	tokenzie := character.NewCharacterTokenizer(unicode.IsLetter)
// 	words := tokenzie.Tokenize(sentence)
// 	gg := fmt.Sprintf("\n\n\n\n%s\n", string(sentence))
// for _, word := range words {
// 	gg +=fmt.Sprintf("[%s]", word.Term)
// }
// 	gg +=
// 	fmt.Println(fmt.Sprintf("\n\n\n\n%s\n", string(sentence))+words)
// 	return words
// 	// words := x.handle.Tokenize(string(sentence), x.mode, false)
// 	// sg := fmt.Sprintf("\n\n\n\n%s\n", sentence)
// 	// for _, word := range words {
// 	// 	token := analysis.Token{
// 	// 		Term:     []byte(word.Str),
// 	// 		Start:    word.Start,
// 	// 		End:      word.End,
// 	// 		Position: pos,
// 	// 		Type:     analysis.Ideographic,
// 	// 	}
// 	// 	sg += fmt.Sprintf("[%s]", word.Str)
// 	// 	result = append(result, &token)
// 	// 	pos++
// 	// }
// 	// fmt.Println(sg)
// 	// return result
// }

func tokenizerConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.Tokenizer, error) {
	dictpath, ok := config["dictpath"].(string)
	if !ok {
		return nil, errors.New("config dictpath not found")
	}
	hmmpath, ok := config["hmmpath"].(string)
	if !ok {
		return nil, errors.New("config hmmpath not found")
	}
	userdictpath, ok := config["userdictpath"].(string)
	if !ok {
		return nil, errors.New("config userdictpath not found")
	}
	mode, ok := config["mode"].(int)
	if ok {
		if mode == 0 {
			return NewJiebaTokenizer(dictpath, hmmpath, userdictpath, gojieba.DefaultMode), nil
		}
	}

	return NewJiebaTokenizer(dictpath, hmmpath, userdictpath, gojieba.SearchMode), nil
}

func init() {
	registry.RegisterTokenizer("gojieba", tokenizerConstructor)
}
