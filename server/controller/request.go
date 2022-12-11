package controller

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"server/models"
)

func santizeRequest(r *models.Request) (err error) {

	// catches manipulation for create requests
	r.CreatedAt = time.Time{}
	r.UpdatedAt = time.Time{}
	r.DeltedAt = gorm.DeletedAt{}
	r.StatusCode = 0
	r.Status = ""
	r.RefundID = 0
	r.Verified = false

	// checks and format request data
	if perr := models.ParseRequest(r); perr != nil {
		return perr
	}
	return nil
}

func (c *Controller) CreateRequest(ctx *gin.Context) {

	// Binding Request Body to data-struct
	var request models.Request

	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Anfrage fehlgeschlagen"})
		return
	}

	// sanitize Request
	if err := santizeRequest(&request); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Anfrage fehlgeschlagen"})
		return
	}

	// query database
	// result only represents the status of the query
	// the data will be stored in request!
	res := c.db.Create(&request)

	if res.Error != nil {
		log.Println(res.Error.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Anfrage fehlgeschlagen"})
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

	// sanitize Request
	if err := santizeRequest(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update data
	// result only represents the status
	result := c.db.Model(&request).Where("id = ?", id).
		Select(
			"FirstName",
			"LastName",
			"MatriculationNumber",
			"UniID",
			"Email",
			"Phone",
			"IBAN",
			"BIC",
			"AccountOwner").
		Updates(&request)

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
