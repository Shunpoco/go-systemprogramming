package main

import (
	"context"
	"context-value-tips/fuga"
	"context-value-tips/hoge"
)

func main() {
	ctx := context.Background()

	ctx = hoge.SetValue(ctx)
	ctx = fuga.SetValue(ctx)

	hoge.GetValueFromHoge(ctx)
	fuga.GetValueFromFuga(ctx)
}
