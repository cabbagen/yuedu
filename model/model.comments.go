package model

import (
	"yuedu/database"
	"github.com/jinzhu/gorm"
	"log"
)

type CommentlModel struct {
	database        *gorm.DB
}

func NewCommentlModel() CommentlModel {
	return CommentlModel { database.GetDataBase() }
}

// 评论信息实体
type CommentInfo struct {
	Id                  int       `json:"id"`
	UserId              int       `json:"userId"`
	Username            string    `json:"username"`
	UserAvatar          string    `json:"userAvatar"`
	CommentPath         string    `json:"commentPath"`
	CommentContent      string    `json:"commentContent"`
	CommentTime         string    `json:"commentTime"`
}

func(cm CommentlModel) GetArticleCommentInfos(articleId int) []CommentInfo {

	var commentInfos []CommentInfo

	rows, rowsErr := cm.database.Table("yd_comments").
		Select("yd_comments.id, yd_comments.path, yd_comments.content_text, yd_comments.created_at, yd_users.id, yd_users.username, yd_users.avatar").
		Joins("inner join yd_users on yd_users.id = yd_comments.user_id").
		Where("yd_comments.article_id = ?", articleId).
		Rows()

	if rowsErr != nil {
		log.Fatalln(rowsErr)
	}

	for rows.Next() {
		var commentInfo CommentInfo = CommentInfo{}

		if err := rows.Scan(&commentInfo.Id, &commentInfo.CommentPath, &commentInfo.CommentContent, &commentInfo.CommentTime, &commentInfo.UserId, &commentInfo.Username, &commentInfo.UserAvatar); err != nil {
			log.Fatalln(err)
		}

		commentInfos = append(commentInfos, commentInfo)
	}

	return commentInfos
}

