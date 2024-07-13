package models

type LeaveMessage struct {
	Id      int    `json:"id"`
	Email   string `json:"email"`
	Ip      string `json:"ip"`
	Address string `json:"address"`
	Browser string `json:"browser"`
	Os      string `json:"os"`
	Time    int    `json:"time"` //时间戳
	Content string `json:"content"`
	Examine string `json:"examine"`
}

func (LeaveMessage) TableName() string {
	return "leave_message"
}

func (lm LeaveMessage) DeleteData(id int) {
	DB.Where("id = ?", id).Delete(&lm)
}

func (lm LeaveMessage) PassExamine(id int) {
	DB.Where("id = ?", id).First(&lm)
	lm.Examine = "yes"
	DB.Save(&lm)
}
