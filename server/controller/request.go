package controller

import (
	"net/http"
	"server/models"

	"github.com/gin-gonic/gin"
)

type requestJSON struct {
	ID                  uint   `json:"id"`
	FirstName           string `json:"firstname"`
	LastName            string `json:"lastname"`
	MatriculationNumber string `json:"matriculationnumber"`
	UniID               string `json:"uniid"`
	Email               string `json:"email"`
	Phone               string `json:"phone"`
	IBAN                string `json:"iban"`
	BIC                 string `json:"bic"`
	AccountOwner        string `json:"accountowner"`
}

func CreateRequest(c *gin.Context) {

	// Binding Request Body to data-struct
	var data requestJSON

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// convert request-data to database-model
	var request models.Request

	{
		request.FirstName = data.FirstName
		request.LastName = data.LastName
		request.MatriculationNumber = data.MatriculationNumber
		request.UniID = data.UniID
		request.Email = data.Email
		request.Phone = data.Phone
		request.IBAN = data.IBAN
		request.BIC = data.BIC
		request.AccountOwner = data.AccountOwner
	}

	// query database

	result := DB.Create(&request)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	// convert database-model to response-data
	var response requestJSON

	{
		response.ID = request.ID
		response.FirstName = request.FirstName
		response.LastName = request.LastName
		response.MatriculationNumber = request.MatriculationNumber
		response.UniID = request.UniID
		response.Email = request.Email
		response.Phone = request.Phone
		response.IBAN = request.IBAN
		response.BIC = request.BIC
		response.AccountOwner = request.AccountOwner
	}

	c.JSON(http.StatusCreated, response)
}
