package actions

import (
	"bitbucket.org/wowbid/go-playground/restaurant"
	"github.com/labstack/echo/v4"
	"net/http"
)

type RestaurantAction struct {
	restaurant *restaurant.RestaurantV2
}

func NewRestaurantAction() *RestaurantAction {
	return &RestaurantAction{restaurant: restaurant.NewRestaurantV2()}
}

func (action *RestaurantAction) Index(c echo.Context) error {
	customers := []string{"customer 1", "customer 2", "customer 3", "customer 4", "customer 5"}
	doneChannel := make(chan bool)

	action.restaurant.Handle(customers, doneChannel)

	select {
	case <-doneChannel:
		return c.JSON(http.StatusOK, "DONE")
	}
}
