package restaurant_parallel

import (
	"fmt"
	"sync"
	"time"
)

type Chef struct {
}

func NewChef() *Chef {
	return &Chef{}
}

func (chef *Chef) Handle(orderChan <-chan string) <-chan string {
	foodChan := make(chan string)

	go func() {
		var wg sync.WaitGroup

		for order := range orderChan {
			wg.Add(1)

			order := order
			go func() {
				defer wg.Done()

				// time needed by chef to prepare a order : cooking, etc
				time.Sleep(1 * time.Second)
				fmt.Println("chef cook", order)

				// result : chef completed an order
				foodChan <- fmt.Sprintf("food for %s", order)
			}()
		}

		defer func() {
			wg.Wait()
			// close a channel to indicate that no more values will be sent.
			close(foodChan)
		}()
	}()

	return foodChan
}
