package models

import "time"

type WrProjects struct {
  Id  int  `json:"id"`
  Project_name  string  `json:"project_name"`
  Test_time  time.Time  `json:"test_time"`
  Publish_time  time.Time  `json:"publish_time"`
  Lft  int  `json:"lft"`
  Rgt  int  `json:"rgt"`
  Level  int  `json:"level"`
  Fid  int  `json:"fid"`
  Created  time.Time  `json:"created"`
  Updated  time.Time  `json:"updated"`

}

func (m *WrProjects) TableName() string {
	return "wr_projects"
}