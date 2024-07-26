package main

import (
	"context"
	"fmt"

	"contextparams/parameters"
)

func main() {
	// Imagine this is a handler and it has decoded a bunch of variables from
	// the buffalo context...
	p := parameters.NewParameters(map[string]any{
		"bird":  "budgerigar",
		"fruit": "pineapple",
		"year":  1999,
	})
	ctx := context.Background()
	// In reality we'd extract a clean context.Context from inside the
	// buffalo.Context by doing:
	// 	 dc := c.(*buffalo.DefaultContext)
	//	 ctx := dc.Context
	// then use that to store the parameters. That way we'd get cancellation
	// support, but ensure none of the buffalo context methods were accessible
	// even by type-asserting or casting.
	ctx = parameters.PutParameters(ctx, p)
	restOfCode(ctx)
}

func restOfCode(ctx context.Context) {
	p := parameters.GetParameters(ctx)
	if v, err := p.GetString("bird"); err != nil {
		fmt.Printf("error getting bird value: %s\n", err)
	} else {
		fmt.Printf("bird value: %s\n", v)
	}
	if v, err := p.GetString("cat"); err != nil {
		fmt.Printf("error getting cat value: %s\n", err)
	} else {
		fmt.Printf("cat value: %s\n", v)
	}
	if v, err := p.GetString("year"); err != nil {
		fmt.Printf("error getting year value: %s\n", err)
	} else {
		fmt.Printf("year value: %s\n", v)
	}
	// ...but no access to write to p and use it as a bag of globals, which
	// also means it should be safe for concurrent use.
}
