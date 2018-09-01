package schema

type User struct {
  Schema
  UserName         string       `gorm:"column:username;type=varchar(255);not null" json:"username"`
  PassWord         string       `gorm:"column:passworde;type=varchar(255);not null" json:"password"`
  Avator           string       `gorm:"column:avator;type=varchar(255);not null;default:\"\"" json:"avator"`
  Gender           int8         `gorm:"column:gender;type=tinyint;not null" json:"gender"`
  Mobile           string       `gorm:"column:mobile;unique;type=varchar(20);not null" json:"mobile"`
  Email            string       `gorm:"column:email;type=varchar(100);not null" json:"email"`
  Extra            string       `gorm:"column:extra;type=varchar(255);not null;default:\"\"" json:"extra"`
}
