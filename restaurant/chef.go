package restaurant

import (
	"fmt"
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
		for order := range orderChan {
			// time needed by chef to prepare a order : cooking, etc
			time.Sleep(1 * time.Second)
			fmt.Println("chef cook", order)

			// result : chef completed an order
			foodChan <- fmt.Sprintf("food for %s", order)
		}

		// close a channel to indicate that no more values will be sent.
		close(foodChan)
	}()

	return foodChan
}
