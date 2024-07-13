package models

type EmailCode struct {
	Email string `json:"email"`
	Code  string `json:"code"`
	Type  string `json:"type"`
}

func (EmailCode) TableName() string {
	return "email_code"
}
