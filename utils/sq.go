package utils

import "time"

// receives integers from a channel and returns a channel that emits the square of each received integer
func Sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			time.Sleep(1 * time.Second)
			// goroutine that sends the square of each received integer on the channel
			out <- n * n
		}

		// After the inbound channel is closed and this stage has sent all the values downstream, it closes the outbound channel
		close(out)
	}()
	return out
}
