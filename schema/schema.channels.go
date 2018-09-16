package schema

type Channel struct {
  Schema
  Name         string             `gorm:"column:name;type=varchar(255);not null;unique" json:"name"`
  State        int8               `gorm:"column:state;type=tinyint;not null" json:"state"`
}
