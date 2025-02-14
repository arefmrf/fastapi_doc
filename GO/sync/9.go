package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.WithValue(context.Background(), "language", "Go")

	fmt.Println(manager(ctx, "language"))
}

func manager(ctx context.Context, key string) string {
	if v := ctx.Value(key); v != nil {

		fmt.Printf("%T\n", v)
		return v.(string)
	}
	return "not found value"
}
