package html_builder

func Text(text string) Node {
	return TextNode(text)
}

func Br() Node {
	return RawNode("<br/>")
}
func Hr() Node {
	return RawNode("<hr/>")
}

func A(attributes Attrs, children ...Node) Node {
	return Element{"a", attributes, children}
}
func Abbr(attributes Attrs, children ...Node) Node {
	return Element{"abbr", attributes, children}
}
func Acronym(attributes Attrs, children ...Node) Node {
	return Element{"acronym", attributes, children}
}
func Address(attributes Attrs, children ...Node) Node {
	return Element{"address", attributes, children}
}
func Applet(attributes Attrs, children ...Node) Node {
	return Element{"applet", attributes, children}
}
func Area(attributes Attrs, children ...Node) Node {
	return Element{"area", attributes, children}
}
func Article(attributes Attrs, children ...Node) Node {
	return Element{"article", attributes, children}
}
func Aside(attributes Attrs, children ...Node) Node {
	return Element{"aside", attributes, children}
}
func Audio(attributes Attrs, children ...Node) Node {
	return Element{"audio", attributes, children}
}
func B(attributes Attrs, children ...Node) Node {
	return Element{"b", attributes, children}
}
func Base(attributes Attrs, children ...Node) Node {
	return Element{"base", attributes, children}
}
func Basefont(attributes Attrs, children ...Node) Node {
	return Element{"basefont", attributes, children}
}
func Bdi(attributes Attrs, children ...Node) Node {
	return Element{"bdi", attributes, children}
}
func Bdo(attributes Attrs, children ...Node) Node {
	return Element{"bdo", attributes, children}
}
func Big(attributes Attrs, children ...Node) Node {
	return Element{"big", attributes, children}
}
func Blockquote(attributes Attrs, children ...Node) Node {
	return Element{"blockquote", attributes, children}
}
func Body(attributes Attrs, children ...Node) Node {
	return Element{"body", attributes, children}
}
func Button(attributes Attrs, children ...Node) Node {
	return Element{"button", attributes, children}
}
func Canvas(attributes Attrs, children ...Node) Node {
	return Element{"canvas", attributes, children}
}
func Caption(attributes Attrs, children ...Node) Node {
	return Element{"caption", attributes, children}
}
func Center(attributes Attrs, children ...Node) Node {
	return Element{"center", attributes, children}
}
func Cite(attributes Attrs, children ...Node) Node {
	return Element{"cite", attributes, children}
}
func Code(attributes Attrs, children ...Node) Node {
	return Element{"code", attributes, children}
}
func Col(attributes Attrs, children ...Node) Node {
	return Element{"col", attributes, children}
}
func Colgroup(attributes Attrs, children ...Node) Node {
	return Element{"colgroup", attributes, children}
}
func Data(attributes Attrs, children ...Node) Node {
	return Element{"data", attributes, children}
}
func Datalist(attributes Attrs, children ...Node) Node {
	return Element{"datalist", attributes, children}
}
func Dd(attributes Attrs, children ...Node) Node {
	return Element{"dd", attributes, children}
}
func Del(attributes Attrs, children ...Node) Node {
	return Element{"del", attributes, children}
}
func Details(attributes Attrs, children ...Node) Node {
	return Element{"details", attributes, children}
}
func Dfn(attributes Attrs, children ...Node) Node {
	return Element{"dfn", attributes, children}
}
func Dialog(attributes Attrs, children ...Node) Node {
	return Element{"dialog", attributes, children}
}
func Dir(attributes Attrs, children ...Node) Node {
	return Element{"dir", attributes, children}
}
func Div(attributes Attrs, children ...Node) Node {
	return Element{"div", attributes, children}
}
func Dl(attributes Attrs, children ...Node) Node {
	return Element{"dl", attributes, children}
}
func Dt(attributes Attrs, children ...Node) Node {
	return Element{"dt", attributes, children}
}
func Em(attributes Attrs, children ...Node) Node {
	return Element{"em", attributes, children}
}
func Embed(attributes Attrs, children ...Node) Node {
	return Element{"embed", attributes, children}
}
func Fieldset(attributes Attrs, children ...Node) Node {
	return Element{"fieldset", attributes, children}
}
func Figcaption(attributes Attrs, children ...Node) Node {
	return Element{"figcaption", attributes, children}
}
func Figure(attributes Attrs, children ...Node) Node {
	return Element{"figure", attributes, children}
}
func Font(attributes Attrs, children ...Node) Node {
	return Element{"font", attributes, children}
}
func Footer(attributes Attrs, children ...Node) Node {
	return Element{"footer", attributes, children}
}
func Form(attributes Attrs, children ...Node) Node {
	return Element{"form", attributes, children}
}
func Frame(attributes Attrs, children ...Node) Node {
	return Element{"frame", attributes, children}
}
func Frameset(attributes Attrs, children ...Node) Node {
	return Element{"frameset", attributes, children}
}
func H1(attributes Attrs, children ...Node) Node {
	return Element{"h1", attributes, children}
}
func H2(attributes Attrs, children ...Node) Node {
	return Element{"h2", attributes, children}
}
func H3(attributes Attrs, children ...Node) Node {
	return Element{"h3", attributes, children}
}
func H4(attributes Attrs, children ...Node) Node {
	return Element{"h4", attributes, children}
}
func H5(attributes Attrs, children ...Node) Node {
	return Element{"h5", attributes, children}
}
func H6(attributes Attrs, children ...Node) Node {
	return Element{"h6", attributes, children}
}
func Head(attributes Attrs, children ...Node) Node {
	return Element{"head", attributes, children}
}
func Header(attributes Attrs, children ...Node) Node {
	return Element{"header", attributes, children}
}
func Html(attributes Attrs, children ...Node) Node {
	return Element{"html", attributes, children}
}
func I(attributes Attrs, children ...Node) Node {
	return Element{"i", attributes, children}
}
func Iframe(attributes Attrs, children ...Node) Node {
	return Element{"iframe", attributes, children}
}
func Img(attributes Attrs, children ...Node) Node {
	return Element{"img", attributes, children}
}
func Input(attributes Attrs, children ...Node) Node {
	return Element{"input", attributes, children}
}
func Ins(attributes Attrs, children ...Node) Node {
	return Element{"ins", attributes, children}
}
func Kbd(attributes Attrs, children ...Node) Node {
	return Element{"kbd", attributes, children}
}
func Label(attributes Attrs, children ...Node) Node {
	return Element{"label", attributes, children}
}
func Legend(attributes Attrs, children ...Node) Node {
	return Element{"legend", attributes, children}
}
func Li(attributes Attrs, children ...Node) Node {
	return Element{"li", attributes, children}
}
func Link(attributes Attrs, children ...Node) Node {
	return Element{"link", attributes, children}
}
func Main(attributes Attrs, children ...Node) Node {
	return Element{"main", attributes, children}
}
func Map(attributes Attrs, children ...Node) Node {
	return Element{"map", attributes, children}
}
func Mark(attributes Attrs, children ...Node) Node {
	return Element{"mark", attributes, children}
}
func Meta(attributes Attrs, children ...Node) Node {
	return Element{"meta", attributes, children}
}
func Meter(attributes Attrs, children ...Node) Node {
	return Element{"meter", attributes, children}
}
func Nav(attributes Attrs, children ...Node) Node {
	return Element{"nav", attributes, children}
}
func Noframes(attributes Attrs, children ...Node) Node {
	return Element{"noframes", attributes, children}
}
func Noscript(attributes Attrs, children ...Node) Node {
	return Element{"noscript", attributes, children}
}
func Object(attributes Attrs, children ...Node) Node {
	return Element{"object", attributes, children}
}
func Ol(attributes Attrs, children ...Node) Node {
	return Element{"ol", attributes, children}
}
func Optgroup(attributes Attrs, children ...Node) Node {
	return Element{"optgroup", attributes, children}
}
func Option(attributes Attrs, children ...Node) Node {
	return Element{"option", attributes, children}
}
func Output(attributes Attrs, children ...Node) Node {
	return Element{"output", attributes, children}
}
func P(attributes Attrs, children ...Node) Node {
	return Element{"p", attributes, children}
}
func Param(attributes Attrs, children ...Node) Node {
	return Element{"param", attributes, children}
}
func Picture(attributes Attrs, children ...Node) Node {
	return Element{"picture", attributes, children}
}
func Pre(attributes Attrs, children ...Node) Node {
	return Element{"pre", attributes, children}
}
func Progress(attributes Attrs, children ...Node) Node {
	return Element{"progress", attributes, children}
}
func Q(attributes Attrs, children ...Node) Node {
	return Element{"q", attributes, children}
}
func Rp(attributes Attrs, children ...Node) Node {
	return Element{"rp", attributes, children}
}
func Rt(attributes Attrs, children ...Node) Node {
	return Element{"rt", attributes, children}
}
func Ruby(attributes Attrs, children ...Node) Node {
	return Element{"ruby", attributes, children}
}
func S(attributes Attrs, children ...Node) Node {
	return Element{"s", attributes, children}
}
func Samp(attributes Attrs, children ...Node) Node {
	return Element{"samp", attributes, children}
}
func Script(attributes Attrs, children ...Node) Node {
	return Element{"script", attributes, children}
}
func Section(attributes Attrs, children ...Node) Node {
	return Element{"section", attributes, children}
}
func Select(attributes Attrs, children ...Node) Node {
	return Element{"select", attributes, children}
}
func Small(attributes Attrs, children ...Node) Node {
	return Element{"small", attributes, children}
}
func Source(attributes Attrs, children ...Node) Node {
	return Element{"source", attributes, children}
}
func Span(attributes Attrs, children ...Node) Node {
	return Element{"span", attributes, children}
}
func Strike(attributes Attrs, children ...Node) Node {
	return Element{"strike", attributes, children}
}
func Strong(attributes Attrs, children ...Node) Node {
	return Element{"strong", attributes, children}
}
func Style(attributes Attrs, children ...Node) Node {
	return Element{"style", attributes, children}
}
func Sub(attributes Attrs, children ...Node) Node {
	return Element{"sub", attributes, children}
}
func Summary(attributes Attrs, children ...Node) Node {
	return Element{"summary", attributes, children}
}
func Sup(attributes Attrs, children ...Node) Node {
	return Element{"sup", attributes, children}
}
func Svg(attributes Attrs, children ...Node) Node {
	return Element{"svg", attributes, children}
}
func Table(attributes Attrs, children ...Node) Node {
	return Element{"table", attributes, children}
}
func Tbody(attributes Attrs, children ...Node) Node {
	return Element{"tbody", attributes, children}
}
func Td(attributes Attrs, children ...Node) Node {
	return Element{"td", attributes, children}
}
func Template(attributes Attrs, children ...Node) Node {
	return Element{"template", attributes, children}
}
func Textarea(attributes Attrs, children ...Node) Node {
	return Element{"textarea", attributes, children}
}
func Tfoot(attributes Attrs, children ...Node) Node {
	return Element{"tfoot", attributes, children}
}
func Th(attributes Attrs, children ...Node) Node {
	return Element{"th", attributes, children}
}
func Thead(attributes Attrs, children ...Node) Node {
	return Element{"thead", attributes, children}
}
func Time(attributes Attrs, children ...Node) Node {
	return Element{"time", attributes, children}
}
func Title(attributes Attrs, children ...Node) Node {
	return Element{"title", attributes, children}
}
func Tr(attributes Attrs, children ...Node) Node {
	return Element{"tr", attributes, children}
}
func Track(attributes Attrs, children ...Node) Node {
	return Element{"track", attributes, children}
}
func Tt(attributes Attrs, children ...Node) Node {
	return Element{"tt", attributes, children}
}
func U(attributes Attrs, children ...Node) Node {
	return Element{"u", attributes, children}
}
func Ul(attributes Attrs, children ...Node) Node {
	return Element{"ul", attributes, children}
}
func Var(attributes Attrs, children ...Node) Node {
	return Element{"var", attributes, children}
}
func Video(attributes Attrs, children ...Node) Node {
	return Element{"video", attributes, children}
}
func Wbr(attributes Attrs, children ...Node) Node {
	return Element{"wbr", attributes, children}
}
