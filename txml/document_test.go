package txml

import (
	"github.com/bmizerany/assert"
	"testing"
	//"io/ioutil"
)

func TestFindAll(t *testing.T) {
	input := `<!DOCTYPE html>
<html>
	<head>
		<title>
		the title of the page
		</title>
	</head>
	<body>
		<div class="hey"><h2>Title here</h2></div>
		<span><h2>Yoyoyo</h2></span>
		<div id="x">
			<span>
				span content<a href="xxx"><div><li><h2>1st div content</h2></li></div></a>
			</span>
		</div>
		<div class="yo hey">
			<a href="xyz"><div class="cow sheep bunny"><h8>h8 content</h8></div></a>
		</div>
	</body>
</html>
`
	doc := New()
	err := doc.ParseString(input)
	root := doc.Root
	assert.Equal(t, err, nil)
	assert.NotEqual(t, root, nil)

	//	s1 := doc.Root.ToString()
	//	d1 := New()
	//	err = d1.ParseString(s1)
	//	assert.Equal(t, err, nil)
	//
	//	s2 := d1.Root.ToString()
	//	d2 := New()
	//	err = d2.ParseString(s2)
	//	assert.Equal(t, err, nil)
	//
	//	s3 := d2.Root.ToString()
	//	d3 := New()
	//	err = d3.ParseString(s3)
	//
	//	ioutil.WriteFile("s1.txt", []byte(s1), 0644)
	//	ioutil.WriteFile("s2.txt", []byte(s2), 0644)
	//	ioutil.WriteFile("s3.txt", []byte(s3), 0644)
	//
	//	assert.Equal(t, s1, s2)
	//	assert.Equal(t, s2, s3)

	// Test FindAll
	es := root.FindAll("head")
	assert.Equal(t, len(es), 1)
	assert.Equal(t, es[0].Name, "head")
	assert.Equal(t, es[0].Value, "")
	println(es[0].ToString())

	es = root.FindAll("head title")
	assert.Equal(t, len(es), 1)
	assert.Equal(t, es[0].Name, "title")
	assert.Equal(t, es[0].Value, "the title of the page")
	println(es[0].ToString())

	es = root.FindAll("body/span/h2")
	assert.Equal(t, len(es), 1)
	assert.Equal(t, es[0].Name, "h2")
	assert.Equal(t, es[0].Value, "Yoyoyo")
	println(es[0].ToString())

	es = root.FindAll("div")
	assert.Equal(t, len(es), 3)
	assert.Equal(t, es[0].Name, "div")
	assert.Equal(t, es[0].Value, "")
	println(es[0].ToString())
	println(es[1].ToString())
	println(es[2].ToString())

	es = root.FindAll("h2")
	assert.Equal(t, len(es), 3)
	assert.Equal(t, es[0].Name, "h2")
	println(es[0].ToString())
	println(es[1].ToString())
	println(es[2].ToString())
}

func TestParse(t *testing.T) {
	input := `
	<!--   Copyright w3school.com.cn  -->
	<note>
		<to>George</to>
		<from>John</from>
		<heading>Reminder</heading>
		<body>Don't forget the meeting!</body>
	</note>
    `

	// Test parsing
	doc := New()
	err := doc.ParseString(input)
	root := doc.Root
	assert.Equal(t, err, nil)
	assert.NotEqual(t, doc, nil)
	assert.Equal(t, root.Name, "note")
	assert.Equal(t, root.Value, "")
	assert.Equal(t, len(root.Children), 4)
	assert.Equal(t, len(root.Attrs), 0)

	// Test FindAll
	es := root.FindAll("note")
	assert.Equal(t, len(es), 1)
	assert.Equal(t, es[0], root)
	println(es[0].ToString())

	es = root.FindAll("note to")
	assert.Equal(t, len(es), 1)
	assert.Equal(t, es[0].Name, "to")
	assert.Equal(t, es[0].Value, "George")
	println(es[0].ToString())

	es = root.FindAll("note/from")
	assert.Equal(t, len(es), 1)
	assert.Equal(t, es[0].Name, "from")
	assert.Equal(t, es[0].Value, "John")
	println(es[0].ToString())
}
