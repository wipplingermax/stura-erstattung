package controller

import (
	models "server/models"

	"github.com/gin-gonic/gin"
)

type Request struct {
	ID     uint32 `json:"id"`
	Status string `json:"status"`
	Refund Refund `json:"refund"`
}

func GETRequest(c *gin.Context) {
	id := c.Params.ByName("id")

	var request models.Request

	err := models.Request.GETRequest(&request, id)

	if err != nil {

	}

}
