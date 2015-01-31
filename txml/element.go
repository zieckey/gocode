package txml

import (
	"bytes"
	//"encoding/xml"
	//"fmt"
	"io"
	//"strings"
)

type ElementArray []*Element

type Element struct {
	Name   string
	Value  string
	Attrs  map[string]string
	Childs map[string]ElementArray
	Parent *Element
	//Root *Element
}

func NewElement() *Element {
	el := &Element{
		Attrs:  make(map[string]string),
		Childs: make(map[string]ElementArray),
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
	if len(e.Childs) > 0 {
		for _, cl := range e.Childs {
			for _, c := range cl {
				c.Write(w)
			}
		}
	}

	w.Write([]byte("</"))
	w.Write([]byte(e.Name))
	w.Write([]byte(">"))
}
