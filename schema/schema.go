package schema

import "time"

type Schema struct {
  ID          uint         `gorm:"primary_key" json:"id"`
  CreatedAt   time.Time    `json:"createdAt"`
  UpdatedAt   time.Time		 `json:"updatedAt"`
  DeletedAt   *time.Time   `sql:"index" json:"deletedAt"`      
}
