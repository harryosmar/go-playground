package actions

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"sync"
	"time"
)

type ParallelRestaurant struct {
}

func NewParallelRestaurant() *ParallelRestaurant {
	return &ParallelRestaurant{}
}

func (parallelRestaurant *ParallelRestaurant) Index(c echo.Context) error {
	customers := [5]string{"customer 1", "customer 2", "customer 3", "customer 4", "customer 5"}
	orderChan := make(chan string, 5)
	foodChan := make(chan string, 5)
	completedOrdersChan := make(chan []string)

	go parallelRestaurant.handleCustomer(customers, orderChan)
	go parallelRestaurant.handleOrder(orderChan, foodChan)
	go parallelRestaurant.handleFood(foodChan, completedOrdersChan)

	select {
	case completedOrders := <-completedOrdersChan:
		return c.JSON(http.StatusOK, completedOrders)
	}
}

func (parallelRestaurant *ParallelRestaurant) handleCustomer(customers [5]string, orderChan chan string) {
	var wg sync.WaitGroup

	for _, customer := range customers {
		wg.Add(1)
		customer := customer

		go func(wg *sync.WaitGroup) {
			defer wg.Done()

			// time needed by cashier to handle a customer
			time.Sleep(1 * time.Second)
			fmt.Println("cashier is handling", customer)

			// cashier received an order from customer
			orderChan <- fmt.Sprintf("order from %s", customer)
		}(&wg)
	}

	defer func() {
		wg.Wait()
		close(orderChan)
	}()
}

func (parallelRestaurant *ParallelRestaurant) handleOrder(orderChan chan string, foodChan chan string) {
	var wg sync.WaitGroup

	for order := range orderChan {
		wg.Add(1)
		order := order

		go func(wg *sync.WaitGroup) {
			defer wg.Done()

			// time needed by chef to prepare a order : cooking, etc
			time.Sleep(1 * time.Second)
			fmt.Println("chef cook", order)

			// result : chef completed an order
			foodChan <- fmt.Sprintf("food for %s", order)
		}(&wg)
	}

	defer func() {
		wg.Wait()
		close(foodChan)
	}()
}

func (parallelRestaurant *ParallelRestaurant) handleFood(foodChan chan string, completedOrdersChan chan []string) {
	var wg sync.WaitGroup
	var completedOrders []string
	var mux sync.Mutex

	for food := range foodChan {
		wg.Add(1)
		food := food

		go func(wg *sync.WaitGroup) {
			defer wg.Done()

			// time needed by waitress to deliver a food to customer
			time.Sleep(1 * time.Second)
			fmt.Println("deliver", food)

			mux.Lock()
			// Lock so only one goroutine at a time can access the completedOrders
			completedOrders = append(completedOrders, fmt.Sprintf("%s is COMPLETED", food))
			mux.Unlock()
		}(&wg)
	}

	// all foods delivered
	defer func() {
		wg.Wait()
		completedOrdersChan <- completedOrders
	}()
}
