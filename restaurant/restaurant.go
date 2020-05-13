package restaurant

type RestaurantV2 struct {
	cashier  *Cashier
	chef     *Chef
	waitress *Waitress
}

func NewRestaurantV2() *RestaurantV2 {
	return &RestaurantV2{cashier: NewCashier(), chef: NewChef(), waitress: NewWaitress()}
}

func (restaurant *RestaurantV2) Handle(customers []string, doneChannel chan bool) {
	restaurant.waitress.Handle(restaurant.chef.Handle(restaurant.cashier.Handle(customers)), doneChannel)
}
