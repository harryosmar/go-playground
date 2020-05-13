package actions

import (
	"bitbucket.org/wowbid/go-playground/restaurant_parallel"
	"github.com/labstack/echo/v4"
	"net/http"
)

type RestaurantParallelAction struct {
	restaurant *restaurant_parallel.RestaurantV2
}

func NewRestaurantParallelAction() *RestaurantParallelAction {
	return &RestaurantParallelAction{restaurant: restaurant_parallel.NewRestaurantV2()}
}

func (action *RestaurantParallelAction) Index(c echo.Context) error {
	customers := []string{"customer 1", "customer 2", "customer 3", "customer 4", "customer 5"}
	completedChan := make(chan []string)

	action.restaurant.Handle(customers, completedChan)

	select {
	case completed := <-completedChan:
		return c.JSON(http.StatusOK, completed)
	}
}
