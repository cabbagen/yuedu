package schema

type User struct {
  Schema
  UserName         string       `gorm:"column:username;type=varchar(255);not null;unique" json:"username"`
  PassWord         string       `gorm:"column:password;type=varchar(255);not null" json:"password"`
  Gender           int8         `gorm:"column:gender;type=tinyint;not null" json:"gender"`
  Email            string       `gorm:"column:email;type=varchar(255);not null;default:\"\";unique" json:"email"`
  Address          int          `gorm:"column:address;type=int;not null" json:"address"`
  HomePage         string       `gorm:"column:homepage;type=varchar(255);not null;default:\"\"" json:"homepage"`
  Avator           string       `gorm:"column:avator;type=varchar(255);not null;default:\"\"" json:"avator"`
  Backdrop         string       `gorm:"column:backdrop;type=varchar(255);not null;default:\"\"" json:"backdrop"`
  Extra            string       `gorm:"column:extra;type=varchar(255);not null;default:\"\"" json:"extra"`
}
