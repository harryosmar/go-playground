package actions

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type Restaurant struct {
}

func NewRestaurant() *Restaurant {
	return &Restaurant{}
}

func (restaurant *Restaurant) Index(c echo.Context) error {
	customers := [5]string{"customer 1", "customer 2", "customer 3", "customer 4", "customer 5"}
	orderChan := make(chan string, 5)
	foodChan := make(chan string, 5)
	doneChannel := make(chan bool)

	go restaurant.handleCustomer(customers, orderChan)
	go restaurant.handleOrder(orderChan, foodChan)
	go restaurant.handleFood(foodChan, doneChannel)

	select {
	case <-doneChannel:
		return c.JSON(http.StatusOK, "DONE")
	}
}

func (restaurant *Restaurant) handleCustomer(customers [5]string, orderChan chan string) {
	for _, customer := range customers {
		// time needed by cashier to handle a customer
		time.Sleep(1 * time.Second)
		fmt.Println("cashier is handling", customer)

		// cashier received an order from customer
		orderChan <- fmt.Sprintf("order from %s", customer)
	}

	// close a channel to indicate that no more values will be sent.
	close(orderChan)
}

func (restaurant *Restaurant) handleOrder(orderChan chan string, foodChan chan string) {
	for order := range orderChan {
		// time needed by chef to prepare a order : cooking, etc
		time.Sleep(1 * time.Second)
		fmt.Println("chef cook", order)

		// result : chef completed an order
		foodChan <- fmt.Sprintf("food for %s", order)
	}

	// close a channel to indicate that no more values will be sent.
	close(foodChan)
}

func (restaurant *Restaurant) handleFood(foodChan chan string, doneChannel chan bool) {
	for food := range foodChan {
		// time needed by waitress to deliver a food to customer
		time.Sleep(1 * time.Second)
		fmt.Println("deliver", food)
	}

	// all foods delivered
	doneChannel <- true

	close(doneChannel)
}
