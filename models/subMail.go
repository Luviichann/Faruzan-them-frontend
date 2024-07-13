package models

type SubMail struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
}

func (SubMail) TableName() string {
	return "sub_mail"
}
