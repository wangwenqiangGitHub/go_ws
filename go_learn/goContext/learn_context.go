package  main

import (
	"context"
	"fmt"
	"time"
)
// go chanal select代替
func main(){
	ctx, cancel := context.WithCancel(context.Background())

	go work(ctx, "node01")
	go work(ctx,"node02")
	go work(ctx,"node03")

	time.Sleep(5 * time.Second)
	fmt.Println("stop the gorutine")

	cancel()

}

func work(ctx context.Context, name string) {
	go func() {
		for{
			select {
			case<-ctx.Done():
				fmt.Println(name, "got the stop channel")
				return
			default:
				fmt.Println(name, "still working")
				time.Sleep(1 * time.Second)
			}
		}
	}()

}

