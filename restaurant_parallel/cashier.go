package restaurant_parallel

import (
	"fmt"
	"sync"
	"time"
)

type Cashier struct {
}

func NewCashier() *Cashier {
	return &Cashier{}
}

func (cashier *Cashier) Handle(customers []string) <-chan string {
	orderChan := make(chan string, len(customers))

	go func() {
		var wg sync.WaitGroup

		for _, customer := range customers {
			wg.Add(1)

			customer := customer
			go func() {
				defer wg.Done()

				// time needed by cashier to handle a customer
				time.Sleep(1 * time.Second)
				fmt.Println("cashier is handling", customer)

				// cashier received an order from customer
				orderChan <- fmt.Sprintf("order from %s", customer)
			}()
		}

		// all foods delivered
		defer func() {
			wg.Wait()
			// close a channel to indicate that no more values will be sent.
			close(orderChan)
		}()
	}()

	return orderChan
}
