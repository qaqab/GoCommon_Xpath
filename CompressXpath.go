package GoCommon_Xpath

import (
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

type Xpth struct {
	X *html.Node
}

func (x *Xpth) XpathList(s string) []Xpth {
	nodes := htmlquery.Find(x.X, s)
	result := []Xpth{}
	if x.X == nil {
		return result
	}
	for _, node := range nodes {
		result = append(result, Xpth{X: node})
	}
	return result
}

func (x *Xpth) Xpath(s string) Xpth {
	node := htmlquery.FindOne(x.X, s)
	return Xpth{X: node}
}

func (x *Xpth) Extract(s string) []string {
	result := []string{}
	if x.X == nil {
		return result
	}
	nodes := htmlquery.Find(x.X, s)
	for _, node := range nodes {
		result = append(result, htmlquery.InnerText(node))
	}
	return result
}

func (x *Xpth) ExtractFirst(s string) string {
	var ef string
	if x.X == nil {
		return ef
	}
	node := htmlquery.FindOne(x.X, s)
	if node != nil {
		return htmlquery.InnerText(node)
	} else {
		return ef
	}
}
