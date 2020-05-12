package actions

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type FaninSample struct {
}

func NewFaninSample() *FaninSample {
	return &FaninSample{}
}

func (faninSample *FaninSample) WithFanInParallel(c echo.Context) error {
	// Combine Joe and Ann Channels
	faninChan := faninSample.fanin(faninSample.boring("Joe"), faninSample.boring("Ann"))

	// because we only used 1 channel then double the loop values 2 * 5
	for i := 0; i < 10; i++ {
		// there is no block because only 1 channel used
		fmt.Println(<-faninChan)
	}

	return c.JSON(http.StatusOK, "DONE")
}

func (faninSample *FaninSample) WithFanIn(c echo.Context) error {
	// Combine Joe and Ann Channels
	faninChan := faninSample.fanin(faninSample.boring("Joe"), faninSample.boring("Ann"))

	// because we only used 1 channel then double the loop values 2 * 5
	for i := 0; i < 10; i++ {
		// there is no block because only 1 channel used
		fmt.Println(<-faninChan)
	}

	return c.JSON(http.StatusOK, "DONE")
}

func (faninSample *FaninSample) WithoutFanIn(c echo.Context) error {
	joe := faninSample.boring("Joe")
	ann := faninSample.boring("Ann")

	for i := 0; i < 5; i++ {
		// Joe still block Ann, because of the channel behavior block on receive
		fmt.Println(<-joe)
		fmt.Println(<-ann)
	}

	return c.JSON(http.StatusOK, "DONE")
}

func (faninSample *FaninSample) boring(msg string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			time.Sleep(1 * time.Second)
			c <- fmt.Sprintf("%s %d", msg, i)
		}
	}()

	return c
}

func (faninSample *FaninSample) fanin(cs ...<-chan string) <-chan string {
	faninChan := make(chan string)

	for _, c := range cs {
		c := c
		go func() {
			for {
				faninChan <- <-c
			}
		}()
	}

	return faninChan
}