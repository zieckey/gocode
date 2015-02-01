package txml

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"strings"
)

type Document struct {
	ProcInst xml.ProcInst
	Root     *Element

	TrimSpace bool
}

func New() *Document {
	return &Document{
		TrimSpace: true,
	}
}

func (doc *Document) Parse(r io.Reader) (err error) {
	var current *Element
	decoder := xml.NewDecoder(r)
	for {
		t, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		switch token := t.(type) {
		case xml.StartElement:
			el := NewElement()
			el.Name = token.Name.Local
			for _, a := range token.Attr {
				el.Attrs[a.Name.Local] = a.Value
			}

			if doc.Root == nil {
				doc.Root = el
			} else {
				current.Children[el.Name] = append(current.Children[el.Name], el)
				el.Parent = current
			}
			current = el
		case xml.EndElement:
			current = current.Parent
		case xml.CharData:
			/*
			   <Person>xxx
			       <FirstName>Zieckey</FirstName>
			       <LastName>Wei</LastName>
			   </Person>

			   The "Person" element's CharData will return 3 times.
			   So we only accept the non empty string as its value.
			*/
			if current != nil {
				if len(current.Value) == 0 {
					current.Value = strings.TrimSpace(string(token))
				} else {
					//TODO how to process this case :
					/*
					   <Person>
					       xxx
					       <FirstName>Zieckey</FirstName>
					       zzz
					   </Person>
					*/
				}
			}
		case xml.Comment:
		case xml.ProcInst:
			doc.ProcInst = token
		case xml.Directive:
		default:
			return fmt.Errorf("unknown token : parse xml fail!")
		}
	}
	return nil
}

func (doc *Document) ParseString(xmlstr string) (err error) {
	r := strings.NewReader(xmlstr)
	return doc.Parse(r)
}

func (doc *Document) ToString() string {
	return doc.toPrettyString(0)
}

func (doc *Document) ToPrettyString() string {
	return doc.toPrettyString(4)
}

func (doc *Document) toPrettyString(indent int) string {	
	var buf bytes.Buffer
	buf.Write([]byte("<?"))
	if len(doc.ProcInst.Target) > 0 {
		buf.Write([]byte(doc.ProcInst.Target))
	} else {
		buf.Write([]byte("xml"))
	}
	buf.Write([]byte(" "))
	if len(doc.ProcInst.Inst) > 0 {
		buf.Write(doc.ProcInst.Inst)
	} else {
		buf.Write([]byte("version=\"1.0\" encoding=\"utf-8\""))
	}
	buf.Write([]byte("?>\n"))
	doc.Root.Write(&buf, indent)
	return buf.String()
}
