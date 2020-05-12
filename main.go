package main

import (
	"bitbucket.org/wowbid/order-service/actions"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"net/http"
	"regexp"
	"sync"
	"time"
)

func init() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool("DEBUG") {
		fmt.Println("Service RUN on DEBUG mode")
	}
}

func routes(_echo *echo.Echo) {
	apiGroup := _echo.Group("api")


	apiGroup.GET("/test2", func(context echo.Context) error {
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

	apiGroup.GET("/routine/simple", actions.NewRestaurant().Index)
	apiGroup.GET("/routine/parallel", actions.NewParallelRestaurant().Index)
	apiGroup.GET("/routine/pipeline", actions.NewPipelineSample().Index)
	apiGroup.GET("/routine/fan/in/yes", actions.NewFaninSample().WithFanIn)
	apiGroup.GET("/routine/fan/in/no", actions.NewFaninSample().WithoutFanIn)
	apiGroup.GET("/routine/fan/in/parallel", actions.NewFaninParallel().Index)
}

func middlewares(_echo *echo.Echo) {
	if viper.GetString("APP_ENV") == "production" {
		_echo.Use(middleware.GzipWithConfig(middleware.GzipConfig{
			Level: 5,
		}))
	}

	_promotheus := prometheus.NewPrometheus("echo", func(c echo.Context) bool {
		if matched, _ := regexp.MatchString("^(/health|/metrics)$", c.Path()); matched {
			return true
		}

		return false
	})
	_promotheus.Use(_echo)
}

func main() {
	_echo := echo.New()
	middlewares(_echo)
	routes(_echo)

	_echo.Logger.Fatal(_echo.Start(viper.GetString("SERVER_ADDRESS")))
}
