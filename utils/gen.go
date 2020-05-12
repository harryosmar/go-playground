package utils

import "time"

// converts a list of integers to a channel
func Gen(nums ...int) <-chan int {
	out := make(chan int, len(nums))
	go func() {
		for _, n := range nums {
			time.Sleep(1 * time.Second)
			// goroutine that sends the integers on the channel
			out <- n
		}

		// closes the channel when all the values have been sent
		close(out)
	}()
	return out
}
