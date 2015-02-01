package txml

import (
	"bytes"
	//"encoding/xml"
	//"fmt"
	"io"
	//"strings"
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
	e.Write(&buf)
	return buf.String()
}

func (e *Element) Write(w io.Writer) {
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

	w.Write([]byte(e.Value))
	if len(e.Children) > 0 {
		for _, cl := range e.Children {
			for _, c := range cl {
				c.Write(w)
			}
		}
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

// FindAll 
//func (e *Element) FindAll(selector string) ElementArray {
//	
//}