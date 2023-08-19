package main

import (
	bufferedChannels "concurrentGo/buffered-channels"
	fanin "concurrentGo/fanIn"
	"concurrentGo/fanout"
)

func main() {
	fanin.RunWorker()
	fanout.RunWorker()
	bufferedChannels.RunWorker()
}
