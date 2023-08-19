# Concurrent Go

Code build in live-coding exercises for "Concurrent Programming in Go", taught via O'Reilly Media.

## Usage

Comments in the code explain the purpose of each function. To run the code, simply go run main.go to see concurrent patterns in action.

```go

package main

func main() {
	fanin.RunWorker()
	fanout.RunWorker()
	bufferedChannels.RunWorker()
}


```

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)