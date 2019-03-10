package schema

type Support struct {
	Schema
	Type         int8               `gorm:"column:type;type=tinyint;not null" json:"type"`
	State        int8               `gorm:"column:state;type=tinyint;not null" json:"state"`
	UserId       int                `gorm:"column:user_id;type=int;not null" json:"userId"`
	ArticleId    int                `gorm:"column:article_id;type=int" json:"articled"`
	CommentId    int                `gorm:"column:comment_id;type=int" json:"commentId"`
}
