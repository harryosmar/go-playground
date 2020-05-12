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
    customers := [5]string{"customer 1", "customer 2", "customer 3", "customer 4", "customer 5"}
    orderChan := make(chan string, 5)
  
    // example of go routine cashier when handling the customer
    go handleCustomer(customers, orderChan)
    
    func handleCustomer(customers [5]string, orderChan chan string) {
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