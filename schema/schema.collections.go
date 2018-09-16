package schema

type Collection struct {
  Schema
  State        int8               `gorm:"column:state;type=tinyint;not null" json:"state"`
  UserId       int                `gorm:"column:user_id;type=int;not null" json:"userId"`
  ArticleId    int                `gorm:"column:article_id;type=int;not null" json:"articleId"`
}
