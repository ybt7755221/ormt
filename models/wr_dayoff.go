package models

import "time"

type WrDayoff struct {
	Id            int       `json:"id"`
	Attendance_id int       `json:"attendance_id"`
	Dayoff        time.Time `json:"dayoff"`
	Hours         int       `json:"hours"`
	Backup        string    `json:"backup"`
	Created       time.Time `json:"created"`
	Updated       time.Time `json:"updated"`
}

func (m *WrDayoff) TableName() string {
	return "wr_dayoff"
}
