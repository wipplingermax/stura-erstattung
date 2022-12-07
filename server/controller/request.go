package controller

import (
	"net/http"
	"server/models"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func santizeRequest(r *models.Request) {
	r.CreatedAt = time.Time{}
	r.UpdatedAt = time.Time{}
	r.DeltedAt = gorm.DeletedAt{}
}

func (c *Controller) CreateRequest(ctx *gin.Context) {

	// Binding Request Body to data-struct
	var request models.Request

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	santizeRequest(&request)

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

	id := ctx.Params.ByName("id")

	// Binding Request Body to data-struct
	var request models.Request

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update data
	// result only represents the status
	result := c.db.Model(&request).Where("id = ?", id).Select("FirstName", "LastName", "MatriculationNumber", "UniID",
		"Email", "Phone", "IBAN", "BIC", "AccountOwner").Updates(&request)

	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	//update data-structure for output
	result = c.db.Model(&request).Where("id = ?", id).First(&request)

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
