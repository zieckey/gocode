package txml

import (
	"bytes"
	//"encoding/xml"
	//"fmt"
	"io"
	"strings"
	//"sync"
)

type ElementArray []*Element

// Element is a Node of the XML document
type Element struct {
	// XML exmaple : <Name Attr1="Attr1 Value" Attr2="Attr2 Value">Value</Name>

	Name     string
	Value    string
	Attrs    map[string]string       // The attributes of this Node
	Children map[string]ElementArray // The children of this Node
	Parent   *Element
	//Root *Element
}

func NewElement() *Element {
	el := &Element{
		Attrs:    make(map[string]string),
		Children: make(map[string]ElementArray),
	}
	return el
}

func (e *Element) ToString() string {
	var buf bytes.Buffer
	e.Write(&buf, 0)
	return buf.String()
}

func (e *Element) ToPrettyString() string {
	var buf bytes.Buffer
	e.Write(&buf, 4)
	return buf.String()
}

func (e *Element) Write(w io.Writer, indent int) {
	if indent > 0 {
		w.Write(bytes.Repeat([]byte(" "), indent))
	}
	w.Write([]byte("<"))
	w.Write([]byte(e.Name))
	if len(e.Attrs) > 0 {
		for n, v := range e.Attrs {
			w.Write([]byte(" "))
			w.Write([]byte(n))
			w.Write([]byte("=\""))
			w.Write([]byte(v))
			w.Write([]byte("\""))
		}
	}
	w.Write([]byte(">"))
	if indent > 0 {
		w.Write([]byte("\n"))
		w.Write(bytes.Repeat([]byte(" "), indent+4))
	}
	w.Write([]byte(e.Value))
	if len(e.Children) > 0 {
		if indent > 0 {
			w.Write([]byte("\n"))
		}
		for _, cl := range e.Children {
			for _, c := range cl {
				i := indent
				if i > 0 {
					i = i + 4
				}
				c.Write(w, i)
				//if index+1 < len(cl) && indent > 0 {
				w.Write([]byte("\n"))
				//}
			}
		}
	}
	if indent > 0 {
		w.Write([]byte("\n"))
		w.Write(bytes.Repeat([]byte(" "), indent))
	}
	w.Write([]byte("</"))
	w.Write([]byte(e.Name))
	w.Write([]byte(">"))

}

// SetAttr sets the attribute of the Element with name and value
// It will override the original value if the Element already has the
// attribute with the same name
//func (e *Element) SetAttr(name, value string) {
//	// TODO add a mutex Locker?
//	e.Attrs[name] = value
//}

// FindFirst finds the fist Element with the selector
// selector can be splitted by " " or "/"
func (e *Element) FindFirst(selector string) *Element {
	names := strings.FieldsFunc(selector, selectorSatisfy)
	es := e.findAll(names)
	if len(es) > 0 {
		return es[0]
	}

	return nil
}

// FindAll recursively finds all the Element with the selector
// selector can be splitted by " " or "/"
func (e *Element) FindAll(selector string) ElementArray {
	names := strings.FieldsFunc(selector, selectorSatisfy)
	return e.findAll(names)
}

func (e *Element) findAll(names []string) ElementArray {
	es := make(ElementArray, 0)
	if len(names) == 0 {
		return es
	}

	if len(names) == 1 {
		if e.Name == names[0] {
			es = append(es, e)
			return es
		}

		goto FindInChildrenElement
	}

	if e.Name == names[0] {
		if arr, ok := e.Children[names[1]]; ok {
			for _, c := range arr {
				r := c.findAll(names[1:])
				if len(r) > 0 {
					es = append(es, r...)
				}
			}
			return es
		}
	}

FindInChildrenElement:
	for _, arr := range e.Children {
		for _, c := range arr {
			r := c.findAll(names[0:])
			if len(r) > 0 {
				es = append(es, r...)
			}
		}
	}
	return es
}

func selectorSatisfy(r rune) bool {
	if r == ' ' || r == '/' {
		return true
	}
	return false
}
