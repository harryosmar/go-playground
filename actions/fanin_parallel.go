package actions

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"sync"
	"time"
)

type FaninParallel struct {
	mux sync.Mutex
}

func NewFaninParallel() *FaninParallel {
	return &FaninParallel{}
}

func (faninParallel *FaninParallel) Index(c echo.Context) error {
	joe := faninParallel.boring("Joe")
	ann := faninParallel.boring("Ann")
	robin := faninParallel.boring("Robin")
	roni := faninParallel.boring("Roni")
	harry := faninParallel.boring("Harry")
	done := make(chan bool)
	faninChan := faninParallel.fanin(joe, ann, robin, roni, harry)
	var messages []string

	go func() {
		for i := 0; i < 5; i++ {
			msg := <-faninChan
			fmt.Println(msg)

			faninParallel.mux.Lock()
			messages = append(messages, msg)
			faninParallel.mux.Unlock()
		}

		done <- true
	}()

	select {
	case <-done:
		return c.JSON(http.StatusOK, messages)
	}
}

func (faninParallel *FaninParallel) fanin(cs ...<-chan string) <-chan string {
	faninChan := make(chan string, len(cs))

	for _, c := range cs {
		c := c
		go func() {
			faninChan <- <-c
		}()
	}

	return faninChan
}

func (faninParallel *FaninParallel) boring(msg string) <-chan string {
	c := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c <- fmt.Sprintf("Hi %s", msg)

		close(c)
	}()

	return c
}
