package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	bgCtx := context.Background()
	ctx := context.WithValue(bgCtx, "language", "Go")
	fmt.Println(manager(ctx, "language"))

	fmt.Println("===============================================")
	cancelCtx, cancelFunc := context.WithCancel(bgCtx)
	go task(cancelCtx)
	time.Sleep(time.Second * 3)
	cancelFunc()
	time.Sleep(time.Second * 1)

	fmt.Println("===============================================")
	cancelCtx, cancel := context.WithTimeout(bgCtx, time.Second*3)
	defer cancel()
	go task1(cancelCtx)
	time.Sleep(time.Second * 4)

	fmt.Println("===============================================")
	cancelCtx2, cancel2 := context.WithDeadline(bgCtx, time.Now().Add(time.Second*5))
	defer cancel2()
	go task2(cancelCtx2)
	time.Sleep(time.Second * 6)
}

func manager(ctx context.Context, key string) string {
	if v := ctx.Value(key); v != nil {

		fmt.Printf("%T\n", v)
		vString, ok := v.(string)
		if ok {
			return vString
		}
		return ""
	}
	return "not found value"
}

func task(ctx context.Context) {
	i := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Gracefully exit")
			fmt.Println(ctx.Err())
			return
		default:
			fmt.Println(i)
			time.Sleep(time.Second * 1)
			i++
		}
	}
}

func task1(ctx context.Context) {
	i := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Gracefully exit")
			fmt.Println(ctx.Err())
			return
		default:
			fmt.Println(i)
			time.Sleep(time.Second * 1)
			i++
		}
	}
}

func task2(ctx context.Context) {
	i := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Gracefully exit")
			fmt.Println(ctx.Err())
			return
		default:
			fmt.Println(i)
			time.Sleep(time.Second * 1)
			i++
		}
	}
}
