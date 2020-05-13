package restaurant

import (
	"fmt"
	"time"
)

type Waitress struct {
}

func NewWaitress() *Waitress {
	return &Waitress{}
}

func (waitress *Waitress) Handle(foodChan <-chan string, doneChannel chan bool) {
	go func() {
		for food := range foodChan {
			// time needed by waitress to deliver a food to customer
			time.Sleep(1 * time.Second)
			fmt.Println("deliver", food)
		}

		// all foods delivered
		doneChannel <- true

		close(doneChannel)
	}()
}
