# HTML builder for Go

Exmaple:
```go
package main

import (
	"fmt"
	. "github.com/asmundstavdahl/html_builder"
)

func main() {
	topNode := Html(Attrs{},
		Body(Attrs{},
			Div(Attrs{"style": "background-color: lightblue;"},
				Text("Testing testing…"),
				Hr(),
				A(Attrs{"href": "location.href = 'https://github.com/asmundstavdahl/html_builder'"},
					Img(Attrs{"src": "https://avatars2.githubusercontent.com/u/1936058?s=40&v=4"})),
				Br())))
	output := topNode.Html(
		HtmlConfig{
			Pretty:     true,
			IndentWith: "    "})

	fmt.Println(output)
}

/* $ go run usage_example.go
<html>
    <body>
        <div style="background-color: lightblue;">
            Testing testing…
            <hr/>
            <a href="location.href = 'https://github.com/asmundstavdahl/html_builder'">
                <img src="https://avatars2.githubusercontent.com/u/1936058?s=40&v=4"></img>
            </a>
            <br/>
        </div>
    </body>
</html>
*/
```
