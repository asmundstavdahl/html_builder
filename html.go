package html_builder

import (
	"fmt"
	"html"
	"strings"
)

type Attrs map[string]string

type Node interface {
	Html(HtmlConfig) string
}

type TextNode string
type RawNode string
type Element struct {
	TagName    string
	Attributes Attrs
	Children   []Node
}

type HtmlConfig struct {
	Pretty      bool
	IndentWith  string
	IndentLevel int
}

func (node TextNode) Html(conf HtmlConfig) string {
	return html.EscapeString(string(node))
}

func (node RawNode) Html(conf HtmlConfig) string {
	return string(node)
}

func (el Element) Html(conf HtmlConfig) string {
	children := ""

	if len(el.Children) > 0 {
		childConf := conf
		childConf.IndentLevel += 1
		children = el.compileChildren(childConf)
	}

	prettyBeforeChildren := ""
	prettyAfterChildren := ""
	if conf.Pretty {
		if len(children) > 0 {
			prettyBeforeChildren = strcat("\n", strings.Repeat(conf.IndentWith, conf.IndentLevel+1), "")
			prettyAfterChildren = strcat("\n", strings.Repeat(conf.IndentWith, conf.IndentLevel+0), "")
		} else {
			// No whitespace in innerHTML
		}
	}

	return strcat(
		"<", el.TagName, el.Attributes.String(), ">",
		prettyBeforeChildren, children, prettyAfterChildren,
		"</", el.TagName, ">",
		"")
}

func (attrs Attrs) String() string {
	attrStrings := make([]string, len(attrs))

	i := 0
	for name, value := range map[string]string(attrs) {
		attrStrings[i] = fmt.Sprintf(" %s=\"%s\"", name, value)
		i += 1
	}

	return strings.Join(attrStrings, "")
}

func (el Element) compileChildren(conf HtmlConfig) string {
	childHtmls := make([]string, len(el.Children))

	for i := 0; i < len(el.Children); i++ {
		childHtmls[i] = el.Children[i].Html(conf)
	}

	glue := ""
	if conf.Pretty {
		glue = strcat(
			"\n", strings.Repeat(conf.IndentWith, conf.IndentLevel),
			"")
	}

	return strings.Join(childHtmls, glue)
}

func strcat(strs ...string) string {
	return strings.Join(strs[:len(strs)-1], strs[len(strs)-1])
}
