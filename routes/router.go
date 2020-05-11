package routes

import (
	"bitbucket.org/wowbid/order-service/app/cart/actions"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"sync"
	"time"
)

type Router struct {
}

func NewRouter() *Router {
	return &Router{}
}

func (_router *Router) Guest(guestGroup *echo.Group) *Router {
	guestGroup.GET("/health", func(context echo.Context) error {
		return context.JSON(http.StatusOK, "server ready")
	})

	guestGroup.GET("/test2", func(context echo.Context) error {
		theMine := [5]string{"rock", "ore", "ore", "rock", "ore"}
		itemChan := make(chan string, 5)
		theMineUpdatedChan := make(chan []string)

		go func(mine [5]string) {
			var wg sync.WaitGroup

			for _, item := range mine {
				wg.Add(1)

				item := item
				go func(wg *sync.WaitGroup) {
					defer wg.Done()

					fmt.Println("sent ", item)
					time.Sleep(1 * time.Second)
					itemChan <- item
				}(&wg)
			}

			defer func() {
				wg.Wait()
				close(itemChan)
			}()
		}(theMine)

		go func() {
			var theMineUpdated []string
			var wg sync.WaitGroup
			var mux sync.Mutex

			for foundOre := range itemChan {
				wg.Add(1)

				foundOre := foundOre
				go func(wg *sync.WaitGroup) {
					defer wg.Done()

					mux.Lock()
					if foundOre == "ore" {
						theMineUpdated = append(theMineUpdated, fmt.Sprintf("valuable %s", foundOre))
					} else {
						theMineUpdated = append(theMineUpdated, fmt.Sprintf("beautiful %s", foundOre))
					}
					mux.Unlock()

					time.Sleep(1 * time.Second)
					fmt.Println("process ", foundOre)
				}(&wg)
			}

			defer func() {
				wg.Wait()
				theMineUpdatedChan <- theMineUpdated
			}()
		}()

		select {
		case theMineUpdated := <-theMineUpdatedChan:
			return context.JSON(http.StatusOK, theMineUpdated)
		}
	})

	guestGroup.GET("/test", func(context echo.Context) error {
		theMine := [5]string{"rock", "ore", "ore", "rock", "ore"}

		oreChannel := make(chan string, 3)
		minedChannel := make(chan string, 3)
		doneChannel := make(chan bool)

		go func(mine [5]string) {
			for _, item := range mine {
				if item == "ore" {
					time.Sleep(1 * time.Second)
					fmt.Println("found", item)
					oreChannel <- item
				}
			}
			close(oreChannel)
		}(theMine)

		go func() {
			for foundOre := range oreChannel {
				time.Sleep(1 * time.Second)
				fmt.Println("mine", foundOre)
				minedChannel <- foundOre
			}

			//for i := 0; i < 3; i++ {
			//	foundOre := <-oreChannel
			//	time.Sleep(1 * time.Second)
			//	fmt.Println("mine", foundOre)
			//	minedChannel <- foundOre
			//}
			close(minedChannel)
		}()

		go func() {
			for minedOre := range minedChannel {
				time.Sleep(1 * time.Second)
				fmt.Println("smelter", minedOre)
			}
			//for i := 0; i < 3; i++ {
			//	minedOre := <-minedChannel
			//	time.Sleep(1 * time.Second)
			//	fmt.Println("smelter", minedOre)
			//}
			doneChannel <- true
		}()

		select {
		case <-doneChannel:
			return context.JSON(http.StatusOK, "server ready")
		}
	})

	return _router
}

func (_router *Router) Auth(authGroup *echo.Group, cartAction *actions.CartAction) *Router {
	//authGroup.Use(middlewares.NewJWTToken(map[string]string{
	//	"store-service": "./storage/keys/jwtRS256.key.pub",
	//}).Check)

	authGroup.GET("/cart", cartAction.Index).Name = "cart"

	return _router
}
