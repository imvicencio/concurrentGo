package fanin

import (
	"fmt"
	"sync"
)

func RunWorker() {

	var producerGroup sync.WaitGroup
	var workerGroup sync.WaitGroup
	producer := 10

	produceQueue := make(chan int)

	for i := 0; i < producer; i++ {
		producerGroup.Add(1)
		go func(msgId int) {
			doProcess(msgId, produceQueue, &producerGroup)
		}(i)
	}

	workerGroup.Add(1)
	go func() {
		defer workerGroup.Done()
		count := 0

		for range produceQueue {
			count++
		}

		fmt.Println("Message was received: ", count)
	}()

	producerGroup.Wait()
	//close(produceQueue)
	workerGroup.Wait()

}

func doProcess(workedId int, produceQueue chan int, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		if workedId == 0 {
			wg.Wait()
			close(produceQueue)
		}
	}()
	for i := 0; i < 100; i++ {
		produceQueue <- i
	}
	fmt.Println("message was sent by worker", workedId)
}
