package GoCommon_Xpath

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/antchfx/htmlquery"

	"github.com/qaqab/GoCommon_Xpath"
)

func TestPythonIn(t *testing.T) {
	f, _ := os.Open("./testHtml/test.html")
	defer f.Close()
	docStr, _ := io.ReadAll(f)
	node, _ := htmlquery.Parse(bytes.NewReader([]byte(docStr)))
	doc := GoCommon_Xpath.Xpth{X: node}
	fmt.Println(doc.Extract("//span[@id='subtitle']/@data-typed-text"))
	fmt.Println(doc.ExtractFirst("//h1[@id='seo-header']/text()"))
	fmt.Println(doc.Xpath("//span[@id='subtitle']/@data-typed-text"))
	fmt.Println(doc.XpathList("//span[@id='subtitle']"))

}
