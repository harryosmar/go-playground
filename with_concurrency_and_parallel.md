[Home/](https://github.com/harryosmar/go-playground/blob/master/concurrency.md) [Prev/](https://github.com/harryosmar/go-playground/blob/master/with_concurrency.md) [Next](https://github.com/harryosmar/go-playground/blob/master/pipeline.md)

![restaurant_illustration_with_concurrency and parallel](https://github.com/harryosmar/go-playground/blob/master/resources/restaurant_illustration_with_concurrency_and_parallel.png)

There is a restaurant with : 
- 5 workers as `cashier` : a `cashier` need 1 second to handle all customer, 
- 5 workers `chef` : a `chef` need 1 second to cook all the food, and 
- 5 workers as `waitress` : a `waitress` need 1 second to deliver the food to all customer. 


To handle 5 customers, all the workers **should be** only took 1 second. But **IS NOT** because there is still **waiting time** 2 seconds when the process started :
- waiting 1 second for the 1st order ready
- waiting 1 second for the 1st food ready

So the total time should be `parallel time + waiting time = 1 + 2 = 3 seconds`


Source codes :
```
curl --location --request GET 'http://localhost:9091/api/routine/parallel/v2'
```

- [Action Code](https://github.com/harryosmar/go-playground/blob/master/actions/parallel_routine_v2.go)
- [Restaurant Code](https://github.com/harryosmar/go-playground/blob/master/restaurant_parallel/restaurant.go)
- [Cashier Code](https://github.com/harryosmar/go-playground/blob/master/restaurant_parallel/cashier.go)
- [Chef Code](https://github.com/harryosmar/go-playground/blob/master/restaurant_parallel/chef.go)
- [Waitress Code](https://github.com/harryosmar/go-playground/blob/master/restaurant_parallel/waitress.go)

- Time : 3.01 Secons
- Console Output
    ```
    cashier is handling customer 4
    cashier is handling customer 1
    cashier is handling customer 2
    cashier is handling customer 5
    cashier is handling customer 3
    chef cook order from customer 4
    chef cook order from customer 5
    chef cook order from customer 1
    chef cook order from customer 2
    chef cook order from customer 3
    deliver food for order from customer 4
    deliver food for order from customer 5
    deliver food for order from customer 3
    deliver food for order from customer 2
    deliver food for order from customer 1
    ```
- Http Response
    ```json
      [
          "food for order from customer 4 is COMPLETED",
          "food for order from customer 5 is COMPLETED",
          "food for order from customer 3 is COMPLETED",
          "food for order from customer 2 is COMPLETED",
          "food for order from customer 1 is COMPLETED"
      ]
    ```
  
## links
- [sync WaitGroups : To wait for multiple goroutines to finish](https://gobyexample.com/waitgroups)
- [sync.Mutex : only one goroutine can access a variable at a time to avoid conflicts](https://tour.golang.org/concurrency/9)