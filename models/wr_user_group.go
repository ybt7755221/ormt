package models

import "time"

type WrUserGroup struct {
  Id  int  `json:"id"`
  User_id  int  `json:"user_id"`
  Project_id  int  `json:"project_id"`
  Group_id  int  `json:"group_id"`
  Created  time.Time  `json:"created"`
  Updated  time.Time  `json:"updated"`

}

func (m *WrUserGroup) TableName() string {
	return "wr_user_group"
}