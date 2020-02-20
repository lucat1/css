package css

import "github.com/lucat1/randr"

// CSSContext is the key used to access the internal context of this library
const CSSContext randr.ContextKey = "__css__"

// ExtractStyles extracts a valid css string generated
// by the rendered components which called the `Style` function
func ExtractStyles(ctx randr.Context) string {
	if ctx.Data[CSSContext] == nil {
		return ""
	}
	data := ctx.Data[CSSContext].(Data).Data
	res := ""
	for rule, class := range data {
		res += "." + class + "{" + rule + "}"
	}

	return res
}
