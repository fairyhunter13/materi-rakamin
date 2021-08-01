package changoroutine

type adder interface {
	AddCounter(int64)
}

func TryChanGoroutine(intChan chan<- int, add adder) {
	// Add the logic to complete the tests in here
}
