package actions

import (
	"bitbucket.org/wowbid/order-service/utils"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type PipelineSample struct {
}

func NewPipelineSample() *PipelineSample {
	return &PipelineSample{}
}

func (pipelineSample *PipelineSample) Index(c echo.Context) error {
	for n := range utils.Sq(utils.Gen(2, 3, 4, 5, 6)) {
		fmt.Println(n)
	}

	return c.JSON(http.StatusOK, "DONE")
}