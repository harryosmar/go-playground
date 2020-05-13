package restaurant_parallel

import (
	"fmt"
	"sync"
	"time"
)

type Waitress struct {
}

func NewWaitress() *Waitress {
	return &Waitress{}
}

func (waitress *Waitress) Handle(foodChan <-chan string, completedChan chan []string) {
	go func() {
		var wg sync.WaitGroup
		var mux sync.Mutex
		var completed []string

		for food := range foodChan {
			wg.Add(1)

			food := food
			go func() {
				defer wg.Done()

				// time needed by waitress to deliver a food to customer
				time.Sleep(1 * time.Second)
				fmt.Println("deliver", food)

				mux.Lock()
				completed = append(completed, fmt.Sprintf("%s is COMPLETED", food))
				mux.Unlock()
			}()
		}

		defer func() {
			wg.Wait()
			// all foods delivered
			completedChan <- completed

			close(completedChan)
		}()
	}()
}
