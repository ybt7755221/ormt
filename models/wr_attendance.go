package models

import "time"

type WrAttendance struct {
  Id  int  `json:"id"`
  User_id  int  `json:"user_id"`
  Overtime  time.Time  `json:"overtime"`
  Hours  int  `json:"hours"`
  Used  int  `json:"used"`
  Status  int  `json:"status"`
  Created  time.Time  `json:"created"`
  Updated  time.Time  `json:"updated"`

}

func (m *WrAttendance) TableName() string {
	return "wr_attendance"
}