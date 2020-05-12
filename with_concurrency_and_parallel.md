[Home/](https://github.com/harryosmar/go-playground/blob/master/concurrency.md) [Prev/](https://github.com/harryosmar/go-playground/blob/master/with_concurrency.md)

![restaurant_illustration_with_concurrency and parallel](https://github.com/harryosmar/go-playground/blob/master/resources/restaurant_illustration_with_concurrency_and_parallel.png)

There is a restaurant with : 
- 5 workers as `cashier` : a `cashier` need 1 second to handle a customer, 
- 5 workers `chef` : a `chef` need 1 second to cook the food, and 
- 5 workers as `waitress` : a `waitress` need 1 second to deliver the food to customer. 

To handle 5 customers all the workers only need 1 second.