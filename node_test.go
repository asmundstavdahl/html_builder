package elmy_html

import "testing"
import "fmt"

func TestNode(t *testing.T) {
	html := Html(Attrs{},
		Body(Attrs{},
			Div(Attrs{
				"onclick": "location.href = 'https://github.com/asmundstavdahl/elmy-html-go'",
				"style":   "background-color: lightblue;"},
				Br(),
				Br(),
				Br())))
	fmt.Println(html.Html(HtmlConfig{Pretty: true, Indent: "    "}))
}
