package model

import (
	"encoding/xml"
	"fmt"
)

type Records struct {
	XMLName xml.Name `xml:"Records"`
	Records []Record `xml:"Record"`
}

type Record struct {
	XMLName  xml.Name `xml:"Record"`
	Database string   `xml:"database,attr"`
	Fields   []Field  `xml:"Field"`
	keys     map[string]interface{}
}

func (r *Record) Keys() map[string]interface{} {
	if r.keys == nil {
		r.keys = make(map[string]interface{})
		for _, field := range r.Fields {
			if v, found := r.keys[field.Name]; !found {
				r.keys[field.Name] = field.Value
			} else {
				switch v.(type) {
				case string:
					r.keys[field.Name] = []string{v.(string), field.Value}
				case []string:
					r.keys[field.Name] = append(v.([]string), field.Value)
				default:

				}
			}
		}
	}

	return r.keys
}

type Field struct {
	XMLName xml.Name `xml:"Field"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:",innerxml"`
}

type KV map[string]interface{}

// StringMap marshals into XML.
func (this KV) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// v := KV[start.Name]
	// if v != nil {
	// 	append(v)
	// } else {
	//
	// }
	if this == nil {
		this = map[string]interface{}{}
	}
	this[start.Attr[0].Value] = "int{1}"
	fmt.Println(this)

	d.Skip()
	return nil
	for t, err := d.Token(); err == nil; t, err = d.Token() {
		// fmt.Println(start.Attr[0].Value)
		switch t.(type) {
		default:

		}
		// fmt.Println(t)
		// switch token := t.(type) {
		// // 处理元素开始（标签）
		// case xml.StartElement:
		// 	name := token.Name.Local
		// 	fmt.Printf("Token name: %s\n", name)
		// 	for _, attr := range token.Attr {
		// 		attrName := attr.Name.Local
		// 		attrValue := attr.Value
		// 		fmt.Printf("An attribute is: %s %s\n", attrName, attrValue)
		// 	}
		// // 处理元素结束（标签）
		// case xml.EndElement:
		// 	fmt.Printf("Token of '%s' end\n", token.Name.Local)
		// // 处理字符数据（这里就是元素的文本）
		// case xml.CharData:
		// 	content := string([]byte(token))
		// 	fmt.Printf("This is the content: %v\n", content)
		// default:
		// 	// ...
		// }
	}

	// start
	return nil
}
func (s KV) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	// fmt.Print(start)
	// tokens := []xml.Token{start}
	// for key, value := range s {
	//
	// 	t := xml.StartElement{Name: xml.Name{"", key}}
	// 	tokens = append(tokens, t, xml.CharData(value.(string)), xml.EndElement{t.Name})
	// }
	//
	// tokens = append(tokens, xml.EndElement{start.Name})
	//
	// for _, t := range tokens {
	// 	err := e.EncodeToken(t)
	// 	if err != nil {
	// 		return err
	// 	}
	// }
	//
	// // flush to ensure tokens are written
	// err := e.Flush()
	// if err != nil {
	// 	return err
	// }

	return nil
}
