package schema

type Address struct {
  Schema
  Name         string             `gorm:"column:name;type=varchar(255);not null" json:"name"`
  Code         string             `gorm:"column:code;type=varchar(120);not null;unique" json:"code"`
  Level        int8               `gorm:"column:level;type=tinyint;not null" json:"level"`
  ParentCode   string             `gorm:"column:parent_code;type=varchar(120);not null" json:"parentCode"`
}
