package schema

import "time"

type Schema struct {
	ID          int          `gorm:"column:id;primary_key" json:"id"`
	CreatedAt   time.Time    `json:"createdAt"`
	UpdatedAt   time.Time    `json:"updatedAt"`
	DeletedAt   *time.Time   `sql:"index" json:"deletedAt"`      
}
