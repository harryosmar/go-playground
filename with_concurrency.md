[Home/](https://github.com/harryosmar/go-playground/blob/master/concurrency.md) [Prev/](https://github.com/harryosmar/go-playground/blob/master/without_concurrency.md) [Next](https://github.com/harryosmar/go-playground/blob/master/with_concurrency_and_parallel.md)

![restaurant_illustration_with_concurrency](https://github.com/harryosmar/go-playground/blob/master/resources/restaurant_illustration_with_concurrency.png)

There is a restaurant with 3 workers :
- 1 worker as `cashier` : need 1 second to handle a customer, 
- 1 worker as `chef` : need 1 second to cook the food, and 
- the last worker as `waitress` : need 1 second to deliver the food to customer. 

"To handle 5 customer the workers need 5 seconds". **That's not correct**
Because there is still **waiting** time between each workers **when the process started**
- `waiting time = count(worker) - 1`
    - `Chef` waiting for 1st `order` from `cashier` : 1 second
    - `Waitress` waiting for 1st `food` from `chef` : 1 second
- The total time should be : `waiting time + count(customer) = (3 - 1 + 5) = 7 seconds`



Source codes :

```
curl --location --request GET 'http://localhost:9091/api/routine/simple/v2'
```

- [Action Code](https://github.com/harryosmar/go-playground/blob/master/actions/simple_routine_v2.go)
- [Restaurant Code](https://github.com/harryosmar/go-playground/blob/master/restaurant/restaurant.go)
- [Cashier Code](https://github.com/harryosmar/go-playground/blob/master/restaurant/cashier.go)
- [Chef Code](https://github.com/harryosmar/go-playground/blob/master/restaurant/chef.go)
- [Waitress Code](https://github.com/harryosmar/go-playground/blob/master/restaurant/waitress.go)
  
- Time : 7.05 Secons
- Output
    ```
    cashier is handling customer 1
    cashier is handling customer 2
    chef cook order from customer 1
    chef cook order from customer 2
    deliver food for order from customer 1
    cashier is handling customer 3
    deliver food for order from customer 2
    chef cook order from customer 3
    cashier is handling customer 4
    chef cook order from customer 4
    deliver food for order from customer 3
    cashier is handling customer 5
    deliver food for order from customer 4
    chef cook order from customer 5
    deliver food for order from customer 5
    ```
  
## links
- [Learning Go’s Concurrency Through Illustrations](https://medium.com/@trevor4e/learning-gos-concurrency-through-illustrations-8c4aff603b3)
- [Goroutines](https://tour.golang.org/concurrency/1)
- [Channels](https://tour.golang.org/concurrency/2)
- [Buffered Channels](https://tour.golang.org/concurrency/3)
- [Range and Close a Channel](https://tour.golang.org/concurrency/4)
- [Select](https://tour.golang.org/concurrency/5)
- [Pipelines](https://blog.golang.org/pipelines)