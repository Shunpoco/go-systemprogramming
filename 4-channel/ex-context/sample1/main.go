package main

import (
	"context"
	"fmt"
)

func main() {
	a := 1
	ctx, _ := context.WithCancel(context.Background())
	fmt.Println(ctx.Deadline())
	fmt.Println(ctx.Value(a))
	fmt.Println(ctx.Err())
	fmt.Println(ctx.Done())
	fmt.Println(ctx.Err())
}
