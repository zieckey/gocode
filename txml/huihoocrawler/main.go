package main

import (
	"fmt"
	"github.com/Unknwon/com"
	"github.com/zieckey/gocom/txml"
	"net/http"
	"os"
	"io/ioutil"
	"strings"
	"strconv"
	"sync"
)

func main() {
	url := "http://docs.huihoo.com/rsaconference/china-2012/"
	thread := 1
	if len(os.Args) > 1 {
		if os.Args[1] == "-h" || os.Args[1] == "--help" {
			fmt.Printf("Usage : %v <the-url-directory> <thread-count>\n", os.Args[0])
			fmt.Printf("Example : %v http://docs.huihoo.com/rsaconference/china-2012/ 1\n", os.Args[0])
			return
		}
		
		url = os.Args[1]
		if len(os.Args) > 2 {
			thread, _ = strconv.Atoi(os.Args[2])
			if thread < 1 {
				fmt.Printf("thread argument ERROR : %v\n", os.Args[2])
			}
		}
	}

	p, err := com.HttpGetBytes(&http.Client{}, url, nil)
	if err != nil {
		fmt.Printf("ERROR HTTP GET %v [%v]\n", url, err.Error())
		return
	}

	html := string(p)
	doc := txml.New()
	err = doc.ParseHTML(strings.NewReader(html))
	if err != nil {
		fmt.Printf("ERROR xml parsing [%v]\n", err.Error())
		return
	}

	//ea := doc.Root.FindAll("a")
	ea := doc.Root.FindAll("table tr td a")
	prefixurl := url
	if !strings.HasSuffix(prefixurl, "/") {
		prefixurl = url + "/"
	}
	var wg sync.WaitGroup
	running := 0
	for _, e := range ea {
		href, ok := e.Attrs["href"]
		if !ok {
			continue
		}
		if !strings.HasSuffix(href, ".pdf") {
			continue
		}

		pdfurl := prefixurl + href
		running++
		wg.Add(1)
		go DownPDF(strings.Trim(href, "/"), pdfurl, &wg)
		if running >= thread {
			wg.Wait()
			running = 0
		}
	}
	fmt.Printf("Done.\n");
}

func DownPDF(pdf, url string, wg * sync.WaitGroup) {
	defer wg.Done()
	p, err := com.HttpGetBytes(&http.Client{}, url, nil)
	if err != nil {
		fmt.Printf("download pdf [%v] failed : %v\n", url, err.Error())
		return
	}
	
	ioutil.WriteFile(pdf, p, 0644)
	fmt.Printf("download pdf [%v] OK\n", url)
}
