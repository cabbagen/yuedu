package model

import (
	"yuedu/database"
	"github.com/jinzhu/gorm"
	"yuedu/schema"
	"strconv"
	"time"
)

type CommentlModel struct {
	database        *gorm.DB
}

func NewCommentlModel() CommentlModel {
	return CommentlModel { database.GetDataBase() }
}

// 评论信息实体
type CommentInfo struct {
	Id                  int             `json:"id"`
	UserId              int             `json:"userId"`
	Username            string          `json:"username"`
	UserAvatar          string          `json:"userAvatar"`
	CommentPath         string          `json:"commentPath"`
	CommentContent      string          `json:"commentContent"`
	CommentTime         string          `json:"commentTime"`
	Supports            int             `json:"supports"`
	IsSupported         bool            `json:"isSupported"`
	Children            []CommentInfo   `json:"children"`
}

func(cm CommentlModel) GetArticleComments(userId, articleId int) ([]CommentInfo, error) {

	var comments []CommentInfo

	rows, error := cm.database.Table("yd_comments").
		Select("yd_comments.id, yd_comments.user_id, yd_users.username, yd_users.avatar, yd_comments.path, yd_comments.content_text, yd_comments.created_at").
		Joins("inner join yd_users on yd_users.id = yd_comments.user_id").
		Where("yd_comments.article_id = ?", articleId).
		Rows()

	if error != nil {
		return comments, error
	}

	for rows.Next() {
		var comment CommentInfo = CommentInfo{}

		if error := rows.Scan(&comment.Id, &comment.UserId, &comment.Username, &comment.UserAvatar, &comment.CommentPath, &comment.CommentContent, &comment.CommentTime); error != nil {
			return comments, error
		}

		supports, error := NewSupportModel().GetSupportCountForComment(comment.Id)

		if error != nil {
			return comments, error
		}

		isSupported, error := NewSupportModel().IsSupportComment(userId, comment.Id)

		if error != nil {
			return comments, error
		}

		subComments, error := cm.GetSubComments(userId, comment.Id)

		if error != nil {
			return comments, error
		}

		comment.Supports, comment.IsSupported, comment.Children = supports, isSupported, subComments

		comments = append(comments, comment)
	}

	return comments, nil
}

// 获取评论的评论
func (cm CommentlModel) GetSubComments(userId, commentId int) ([]CommentInfo, error) {

	var subComments []CommentInfo

	rows, error := cm.database.Table("yd_comments").
		Select("yd_comments.id, yd_comments.user_id, yd_users.username, yd_users.avatar, yd_comments.path, yd_comments.content_text, yd_comments.created_at").
		Joins("inner join yd_users on yd_users.id = yd_comments.user_id").
		Where("yd_comments.path = ?", "0," + strconv.Itoa(commentId)).
		Rows()

	if error != nil {
		return subComments, error
	}

	for rows.Next() {
		var comment CommentInfo = CommentInfo{}

		if error := rows.Scan(&comment.Id, &comment.UserId, &comment.Username, &comment.UserAvatar, &comment.CommentPath, &comment.CommentContent, &comment.CommentTime); error != nil {
			return subComments, error
		}

		supports, error := NewSupportModel().GetSupportCountForComment(comment.Id)

		if error != nil {
			return subComments, error
		}

		isSupported, error := NewSupportModel().IsSupportComment(userId, comment.Id)

		if error != nil {
			return subComments, error
		}

		comment.Supports, comment.IsSupported = supports, isSupported

		subComments = append(subComments, comment)
	}

	return subComments, nil
}


// 创建文章评论
func (cm CommentlModel) CreateArticleComment(userId, articleId int, content string) error {
	comment := schema.Comment {
		State: 1,
		Path: "0",
		UserId: userId,
		ArticleId: articleId,
		ContentText: content,
	}

	comment.CreatedAt = time.Now()
	comment.UpdatedAt = time.Now()

	result := cm.database.Table("yd_comments").Create(&comment)

	return result.Error
}

// 新建评论的评论
func (cm CommentlModel) CreateCommentComment(userId, commentId int, content string) error {
	comment := schema.Comment {
		State: 1,
		Path: "0," + strconv.Itoa(commentId),
		UserId: userId,
		ArticleId: 0,
		ContentText: content,
	}

	comment.CreatedAt = time.Now()
	comment.UpdatedAt = time.Now()

	result := cm.database.Table("yd_comments").Create(&comment)

	return result.Error
}

// 删除评论
func (cm CommentlModel) DeleteComment(commendId int) error {

	var subComments []schema.Comment

	queryResult := cm.database.Table("yd_comments").Where("path = ?", "0," + strconv.Itoa(commendId)).Find(&subComments)

	if queryResult.Error != nil && !queryResult.RecordNotFound() {
		return queryResult.Error
	}

	commentIds := []int{commendId}

	for _, subComment := range subComments {
		commentIds = append(commentIds, subComment.ID)
	}

	deleteResult := cm.database.Table("yd_comments").Where("id in (?)", commentIds).Delete(&schema.Comment{})

	if deleteResult.Error != nil {
		return deleteResult.Error
	}

	return NewSupportModel().DeleteSupportCommentByCommentIds(commentIds)
}
