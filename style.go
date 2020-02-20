package css

import (
	"strings"

	"github.com/lucat1/randr"
)

// Data is the struct holding this
// library's data during rendering
type Data struct {
	Data  map[string]string
	Count int
}

// Style returns a classname for the associated stylesheet required
func Style(ctx randr.Context, styles map[string]string) string {
	if ctx.Data[CSSContext] == nil {
		ctx.Data[CSSContext] = Data{
			Data: map[string]string{},
		}
	}
	data := ctx.Data[CSSContext].(Data)

	classes := []string{}
	for k, v := range styles {
		rule := k + ":" + v
		if data.Data[rule] == "" {
			// Generate the classname
			val := Next(data.Count)
			data.Data[rule] = val
			data.Count++
		}

		classes = append(classes, data.Data[rule])
	}
	return strings.Join(classes, " ")
}
