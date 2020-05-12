[Home/](https://github.com/harryosmar/go-playground/blob/master/concurrency.md) [Prev/](https://github.com/harryosmar/go-playground/blob/master/with_concurrency.md)

![restaurant_illustration_with_concurrency and parallel](https://github.com/harryosmar/go-playground/blob/master/resources/restaurant_illustration_with_concurrency_and_parallel.png)

There is a restaurant with : 
- 5 workers as `cashier` : a `cashier` need 1 second to handle all customer, 
- 5 workers `chef` : a `chef` need 1 second to cook all the food, and 
- 5 workers as `waitress` : a `waitress` need 1 second to deliver the food to all customer. 


To handle 5 customers all the workers only need 3 seconds.

Source codes :

- [Restaurant Concurrency Parallel](https://github.com/harryosmar/go-playground/blob/master/actions/parallel_routine.go)
    ```go
    foodChan := make(chan string, 5)
    completedOrdersChan := make(chan []string)
  
    // example of go routine waitress when delivering the food to customer  
    go func(foodChan chan string, completedOrdersChan chan []string) {
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
    }(foodChan, completedOrdersChan)
  
    select {
    	case completedOrders := <-completedOrdersChan:
    		return c.JSON(http.StatusOK, completedOrders)
    }
    ```
- Time : 3.01 Secons
- Output
    ```
    cashier is handling customer 1
    cashier is handling customer 4
    cashier is handling customer 2
    cashier is handling customer 5
    cashier is handling customer 3
    chef cook order from customer 4
    chef cook order from customer 5
    chef cook order from customer 3
    chef cook order from customer 2
    chef cook order from customer 1
    deliver food for order from customer 2
    deliver food for order from customer 3
    deliver food for order from customer 5
    deliver food for order from customer 1
    deliver food for order from customer 4
    ```
  
## links
- [WaitGroups : To wait for multiple goroutines to finish](https://gobyexample.com/waitgroups)
- [sync.Mutex : only one goroutine can access a variable at a time to avoid conflicts](https://tour.golang.org/concurrency/9)