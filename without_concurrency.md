![restaurant_illustration_without_concurrency](https://github.com/harryosmar/go-playground/blob/master/resources/restaurant_illustration_without_concurrency.png)

There is a restaurant which only has 1 worker which responsibilities as : 
- `cashier` : need 1 second to handle a customer, 
- `chef` : need 1 second to cook the food, and 
- `waitress` : : need 1 second to deliver the food to customer. 

To handle 1 customer the worker need 3 seconds. 5 customer `5 * 3 = 15` seconds