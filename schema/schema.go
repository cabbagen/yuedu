package schema

import "time"

type Schema struct {
	ID          int          `gorm:"column:id;primary_key" json:"id"`
	CreatedAt   time.Time    `gorm:"column:created_at"  json:"createdAt"`
	UpdatedAt   time.Time    `gorm:"column:updated_at"  json:"updatedAt"`
	DeletedAt   *time.Time   `gorm:"column:deleted_at"  sql:"index" json:"deletedAt"`
}
