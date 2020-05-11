![restaurant_illustration_with_concurrency](https://github.com/harryosmar/go-playground/blob/master/resources/restaurant_illustration_with_concurrency.png)

There is a restaurant with 3 workers :
- 1 worker as `cashier` : need 1 second to handle a customer, 
- 1 worker as `chef` : need 1 second to cook the food, and 
- the last worker as `waitress` : need 1 second to deliver the food to customer. 

To handle 1 customer the workers need 1 seconds. **That's not correct**
Because there is still waiting time between each workers.
- `Chef` waiting for an `order` from `cashier`
- `waitress` waiting for the `food` from `chef`
- But it will be faster compare to only 1 worker.


Source codes :

- [Restaurant Concurrency](https://github.com/harryosmar/go-playground/blob/master/actions/simple_routine.go)
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