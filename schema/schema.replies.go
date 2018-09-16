package schema

type Reply struct {
	Schema
	Type         int8               `gorm:"column:type;type=tinyint;not null" json:"type"`
	State        int8               `gorm:"column:state;type=tinyint;not null" json:"state"`
	UserId       int                `gorm:"column:user_id;type=int;not null" json:"userId"`
	ReplyId      int                `gorm:"column:reply_id;type=int" json:"replyId"`
	CommentId    int                `gorm:"column:comment_id;type=int" json:"commentId"`
	ReplyPath    string             `gorm:"column:reply_path;type=varchar(255)" json:"replyPath"`
	ContentText  string             `gorm:"column:content_text;type=varchar(300)" json:"contentText"`
}
