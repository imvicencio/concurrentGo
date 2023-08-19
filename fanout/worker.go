package fanout

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func RunWorker() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1000*time.Millisecond)
	ctx, deadline := context.WithDeadline(ctx, time.Now().Add(1*time.Minute))
	defer cancel()
	defer deadline()

	var wg sync.WaitGroup
	n := 10

	workQueue := make(chan int)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(msgId int) {
			defer wg.Done()
			for {
				select {
				case msg, ok := <-workQueue:
					if !ok {
						fmt.Println("worker", msgId, " finished")
						return
					}
					doProcess(ctx, msgId, msg)
				case <-ctx.Done():
					fmt.Println("context was cancelled in worker", msgId)
					return
				}

			}
		}(i)
	}

loop:
	for i := 0; i < 1000; i++ {
		select {
		case workQueue <- i:
		case <-ctx.Done():
			fmt.Println("context was cancelled")
			break loop
		}
	}

	close(workQueue)
	wg.Wait()
	fmt.Println("All messages were sent")
}

func doProcess(ctx context.Context, workedId int, msgId int) {
	fmt.Println("sending message ", msgId, " from worker: ", workedId)
	time.Sleep(100 * time.Millisecond)

}
