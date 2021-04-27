package models

import "time"

type WrWorks struct {
	Id         int       `json:"id"`
	User_id    int       `json:"user_id"`
	Project_id int       `json:"project_id"`
	Title      string    `json:"title"`
	Url        string    `json:"url"`
	Progress   int       `json:"progress"`
	Work_type  int       `json:"work_type"`
	Backup     string    `json:"backup"`
	Created    time.Time `json:"created"`
	Updated    time.Time `json:"updated"`
}

func (m *WrWorks) TableName() string {
	return "wr_works"
}
