[Home/](https://github.com/harryosmar/go-playground/blob/master/concurrency.md) [Prev](https://github.com/harryosmar/go-playground/blob/master/without_concurrency.md)

![restaurant_illustration_with_concurrency](https://github.com/harryosmar/go-playground/blob/master/resources/restaurant_illustration_with_concurrency.png)

There is a restaurant with 3 workers :
- 1 worker as `cashier` : need 1 second to handle a customer, 
- 1 worker as `chef` : need 1 second to cook the food, and 
- the last worker as `waitress` : need 1 second to deliver the food to customer. 

To handle 1 customer the workers need 1 seconds. **That's not correct**
Because there is still waiting time between each workers when the process started
- `Chef` waiting for an `order` from `cashier` : 1 second
- `waitress` waiting for the `food` from `chef` : 1 second
- `waiting time = count(worker) - 1`
- The total time should be : `waiting time + count(customer) = (3 - 1 + 5) = 7 seconds`



Source codes :

- [Restaurant Concurrency](https://github.com/harryosmar/go-playground/blob/master/actions/simple_routine.go)
    ```go
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
    ```
- Time : 7.05 Secons
- Output
    ```
    cashier is handling customer 1
    chef cook order from customer 1
    cashier is handling customer 2
    cashier is handling customer 3
    chef cook order from customer 2
    deliver food for order from customer 1
    deliver food for order from customer 2
    chef cook order from customer 3
    cashier is handling customer 4
    cashier is handling customer 5
    deliver food for order from customer 3
    chef cook order from customer 4
    chef cook order from customer 5
    deliver food for order from customer 4
    deliver food for order from customer 5
    ```
  
## links
- [Learning Go’s Concurrency Through Illustrations](https://medium.com/@trevor4e/learning-gos-concurrency-through-illustrations-8c4aff603b3)
- [Goroutines](https://tour.golang.org/concurrency/1)
- [Channels](https://tour.golang.org/concurrency/2)
- [Buffered Channels](https://tour.golang.org/concurrency/3)
- [Range and Close a Channel](https://tour.golang.org/concurrency/4)
- [Select](https://tour.golang.org/concurrency/5)