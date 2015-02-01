package txml

import (
    "testing"
    "strings"
    "github.com/bmizerany/assert"
)

//func TestXYZ(t *testing.T) {
//    input := `<!DOCTYPE html>
//<html>
//	<head>
//		<title>
//		the title of the page
//		</title>
//	</head>
//	<body>
//		<div class="hey" custom_attr="wow"><h2>Title here</h2></div>
//		<span><h2>Yoyoyo</h2></span>
//		<div id="x">
//			<span>
//				span content<a href="xxx"><div><li>1st div content</li></div></a>
//			</span>
//		</div>
//		<div class="yo hey">
//			<a href="xyz"><div class="cow sheep bunny"><h8>h8 content</h8></div></a>
//		</div>
//	</body>
//</html>
//`
//}

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
    doc := New()
    err := doc.ParseString(input)
    root := doc.Root
    assert.Equal(t, err, nil)
    assert.NotEqual(t, doc, nil)
    assert.Equal(t, root.Name, "note")
    assert.Equal(t, strings.TrimSpace(root.Value), "")
    assert.Equal(t, len(root.Children), 4)
    assert.Equal(t, len(root.Attrs), 0)
    
    println(doc.ToString())
}