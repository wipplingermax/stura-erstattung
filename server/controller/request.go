package controller

import (
	"net/http"
	"server/models"

	"github.com/gin-gonic/gin"
)

func (c *Controller) CreateRequest(ctx *gin.Context) {

	// Binding Request Body to data-struct
	var request models.Request

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// query database
	// result only represents the status of the query
	// the data will be stored in request!
	result := c.db.Create(&request)

	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	// Set gin.Context
	ctx.JSON(http.StatusCreated, request)
}

func (c *Controller) ReadRequest(ctx *gin.Context) {

	var request models.Request

	// Get ID from Context
	id := ctx.Params.ByName("id")

	// query database
	// result only represents the status of the query
	// the data will be stored in request!
	result := c.db.First(&request, id)

	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	// Set gin.Context
	ctx.JSON(http.StatusCreated, request)
}

func (c *Controller) ReadRequests(ctx *gin.Context) {

	// query database
	// result only represents the status of the query
	// the data will be stored in request!
	var requests []models.Request

	result := c.db.Find(&requests)

	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	// Set gin.Context
	ctx.JSON(http.StatusCreated, requests)
}

func (c *Controller) UpdateRequest(ctx *gin.Context) {

	// Binding Request Body to data-struct
	var request models.Request

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// query database
	// result only represents the status of the query
	// the data will be stored in request!
	result := c.db.Save(&request)

	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	// Set gin.Context
	ctx.JSON(http.StatusCreated, request)
}

func (c *Controller) DeleteRequest(ctx *gin.Context) {

	var request models.Request

	// Get ID from Context
	id := ctx.Params.ByName("id")

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// query database
	// result only represents the status of the query
	// the data will be stored in request!
	result := c.db.Delete(&request, id)

	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	// Set gin.Context
	ctx.JSON(http.StatusCreated, request)
}
