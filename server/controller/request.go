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

	var request models.Request

	id := c.Params.ByName("id")

	models.GETRequest(&request, id)

}
