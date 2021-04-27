package models

import "time"

type WrUsers struct {
  Id  int  `json:"id"`
  Username  string  `json:"username"`
  Mobile  string  `json:"mobile"`
  Password  string  `json:"password"`
  Email  string  `json:"email"`
  Created  time.Time  `json:"created"`
  Updated  time.Time  `json:"updated"`

}

func (m *WrUsers) TableName() string {
	return "wr_users"
}