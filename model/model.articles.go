package model

import (
	"strings"
	"yuedu/schema"
	"yuedu/database"
	"github.com/jinzhu/gorm"
)

type ArticleModel struct {
	database        *gorm.DB
}

func NewArticleModel() ArticleModel {
	return ArticleModel { database.GetDataBase() }
}

// 文章详情实体
type ArticleInfo struct {
	Id                       int             `json:"id"`
	ChannelInfo              schema.Channel  `json:"channelInfo"`
	Title                    string          `json:"title"`
	Author                   string          `json:"author"`
	AnchorInfo               UserInfo        `json:"anchorInfo"`
	During                   int             `json:"during"`
	PlayNumber               int             `json:"playNumber"`
	CoverImg                 string          `json:"coverImg"`
	Audio                    string          `json:"audio"`
	Tags                     []schema.Tag    `json:"tags"`
	ContentText              string          `json:"contentText"`
	Supports                 int             `json:"supports"`
	Collections              int             `json:"collections"`
}

// 通过文章 Id 获取文章详细信息
func (am ArticleModel) GetArticleInfoById(articleId int) ArticleInfo {
	var article schema.Article

	var articleInfo ArticleInfo

	am.database.Where("id = ?", articleId).First(&article).Scan(&articleInfo)

	am.database.Table("yd_channels").Where("id = ?", article.ChannelId).Scan(&articleInfo.ChannelInfo)

	am.database.Table("yd_tags").Where("state = 1 and id in (?)", strings.Split(article.TagIds, ",")).Find(&articleInfo.Tags)

	am.database.Table("yd_supports").Where("article_id = ? and state = 1 and type = 1", articleId).Count(&articleInfo.Supports)

	am.database.Table("yd_collections").Where("article_id = ? and state = 1", articleId).Count(&articleInfo.Collections)

	articleInfo.AnchorInfo = NewUserModel().GetUserInfo(article.Anchor)

	return articleInfo
}


// 文章列表实体
type SimpleArticleInfo struct {
	Id                       int             `json:"id"`
	Title                    string          `json:"title"`
	Author                   string          `json:"author"`
	AnchorName               string          `json:"anchorName"`
	During                   int             `json:"during"`
	PlayNumber               int             `json:"playNumber"`
	CoverImg                 string          `json:"coverImg"`
	Audio                    string          `json:"audio"`
	ContentText              string          `json:"contentText"`
}

// 获取指定文章相关的 n 条文章
func (am ArticleModel) GetReleasedArticlesByArticleId(articleId int, limit int) []SimpleArticleInfo {
	var releasedArticleIds = am.GetReleaseArticleIdsByArticleId(articleId, limit)
	var releasedArticles []SimpleArticleInfo

	am.database.Table("yd_articles").Select("yd_articles.*, yd_users.username as anchorName").
		Where("yd_articles.id in (?)", releasedArticleIds).
		Joins("inner join yd_users on yd_articles.anchor = yd_users.id").
		Scan(&releasedArticles)

	return releasedArticles
}

func (am ArticleModel) GetReleaseArticleIdsByArticleId(articleId int, limit int) []int {
	var articleTagIds string

	am.database.Table("yd_articles").Where("id = ?", articleId).Pluck("tag_ids", &articleTagIds)

	var mainTagId string = strings.Split(articleTagIds, ",")[0]

	var articleIds []int

	am.database.Table("yd_articles").Where("tag_ids LIKE ?", "%" + mainTagId + "%").Limit(limit).Pluck("id", &articleIds)

	return articleIds
}


// 获取指定文章其他类型的最新文章列表
func (am ArticleModel) GetOtherChannelLastArticlesByArticleId(articleId int) []SimpleArticleInfo {
	var channelId int

	var otherArticles []SimpleArticleInfo

	am.database.Table("yd_articles").Where("id = ?", articleId).Pluck("channel_id", &channelId)

	am.database.Table("yd_articles").Select("yd_articles.*, yd_users.username as anchorName, max(yd_articles.id)").
		Where("yd_articles.channel_id != ?", channelId).
		Joins("inner join yd_users on yd_articles.anchor = yd_users.id").
		Group("channel_id").
		Scan(&otherArticles)

	return otherArticles;
}

