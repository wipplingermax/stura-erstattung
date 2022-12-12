package models

import (
	"fmt"

	loggable "github.com/daqingshu/gorm-loggable"
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"gorm.io/gorm"

	"server/utils"
)

type Request struct {
	UUIDPkey
	loggable.LoggableModel
	FirstName           string    `gorm:"not null" json:"firstname"`
	LastName            string    `gorm:"not null" json:"lastname"`
	MatriculationNumber string    `gorm:"not null" json:"matriculationnumber"`
	UniID               string    `gorm:"not null" json:"uniid"`
	Email               string    `gorm:"not null" json:"email"`
	Phone               string    `json:"phone"`
	IBAN                string    `gorm:"not null" json:"iban"`
	BIC                 string    `json:"bic"`
	AccountOwner        string    `gorm:"not null" json:"accountowner"`
	Verified            bool      `gorm:"not null"`
	StatusCode          uint      `gorm:"not null" json:"statuscode"`
	Status              string    `json:"status"`
	RefundID            uuid.UUID `gorm:"type:uuid"`
	Refund              Refund
}

// StatusCodes
const (
	UNDEFINED int = 0
	ACCEPTED      = 1
	WAITING       = 2
	DECLINED      = 3
)

func ParseRequest(r *Request) (err error) {

	if perr := utils.ParseFirstName(&r.FirstName); perr != nil {
		err = fmt.Errorf("%w%w\n", perr, err)
	}

	if perr := utils.ParseLastName(&r.LastName); perr != nil {
		err = fmt.Errorf("%w%w\n", perr, err)
	}

	if perr := utils.ParseMatriculationNumber(&r.MatriculationNumber); perr != nil {
		err = fmt.Errorf("%w%w\n", perr, err)
	}

	if perr := utils.ParseUniID(&r.UniID); perr != nil {
		err = fmt.Errorf("%w%w\n", perr, err)
	}

	if perr := utils.ParseEmail(&r.Email); perr != nil {
		err = fmt.Errorf("%w%w\n", perr, err)
	}

	// optional
	if r.Phone != "" {
		if perr := utils.ParsePhone(&r.Phone); perr != nil {
			err = fmt.Errorf("%w%w\n", perr, err)
		}
	}

	if perr := utils.ParseIBAN(&r.IBAN); perr != nil {
		err = fmt.Errorf("%w%w\n", perr, err)
	}

	// optional
	if r.BIC != "" {
		if perr := utils.ParseBIC(&r.BIC); perr != nil {
			err = fmt.Errorf("%w%w\n", perr, err)
		}
	}

	if perr := utils.ParseAccountOwner(&r.AccountOwner); perr != nil {
		err = fmt.Errorf("%w%w\n", perr, err)
	}

	if err != nil {
		return err
	}
	return nil
}

func IsValidRequest(r *Request, db *gorm.DB) (err error) {

	// check if uniID exists
	return nil

	// parses

	// isVerfied

}

func (r *Request) CreateRefund() (err error) {
	return nil
}
