package controller

type Refund struct {
	ID                  uint32 `json:"id"`
	FirstName           string `json:"firstname"`
	LastName            string `json:"lastname"`
	MatriculationNumber string `json:"matriculationNumber"`
	UniID               string `json:"uniID"`
	Email               string `json:"email"`
	Phone               string `json:"phone"`
	IBAN                string `json:"iban"`
	BIC                 string `json:"bic"`
	AccountOwner        string `json:"accountOwner"`
}
