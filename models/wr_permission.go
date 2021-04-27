package models

import "time"

type WrPermission struct {
  Id  int  `json:"id"`
  Name  string  `json:"name"`
  Action  string  `json:"action"`
  Type  int  `json:"type"`
  Created  time.Time  `json:"created"`
  Updated  time.Time  `json:"updated"`

}

func (m *WrPermission) TableName() string {
	return "wr_permission"
}