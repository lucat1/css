package main

import (
	"fmt"

	"github.com/lucat1/css"
	"github.com/lucat1/helmet"
	"github.com/lucat1/randr"
)

func example(ctx randr.Context) string {
	cx := css.Style(ctx, map[string]string{
		"color":     "red",
		"font-size": "2rem",
	})

	return randr.HTML(`
		<h1 class={cx}>Hello world! {cx}</h1>
	`)
}

func main() {
	res, ctx := randr.Render(example, nil)

	fmt.Printf("Result:\n%s\nHead:\n%s\n", res, helmet.ExtractHead(ctx))
}
