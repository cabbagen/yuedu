package model

import (
	"yuedu/schema"
	"yuedu/database"
	"github.com/jinzhu/gorm"
	"strings"
	"time"
)

type ArticleModel struct {
	database        *gorm.DB
}

func NewArticleModel() ArticleModel {
	return ArticleModel { database.GetDataBase() }
}

// 通过文章 Id 获取文章详细信息
type ArticleInfo struct {
	Id                       int             `json:"id"`
	ChannelInfo              schema.Channel  `json:"channelInfo"`
	Title                    string          `json:"title"`
	Author                   string          `json:"author"`
	AnchorInfo               schema.User     `json:"anchorInfo"`
	During                   int             `json:"during"`
	PlayNumber               int             `json:"playNumber"`
	CoverImg                 string          `json:"coverImg"`
	Audio                    string          `json:"audio"`
	Tags                     []schema.Tag    `json:"tags"`
	ContentText              string          `json:"contentText"`
	CreatedAt                time.Time       `json:"createdAt"`
	UpdatedAt                time.Time       `json:"updatedAt"`
	DeletedAt                *time.Time      `json:"deletedAt"`
}

func (am ArticleModel) GetArticleInfoById(articleId int) ArticleInfo {
	var article schema.Article
	var articleInfo ArticleInfo

	am.database.Where("id = ?", articleId).First(&article).Scan(&articleInfo)

	am.database.Table("yd_channels").Where("id = ?", article.ChannelId).Scan(&articleInfo.ChannelInfo)

	am.database.Table("yd_users").Where("id = ?", article.Anchor).First(&articleInfo.AnchorInfo)

	am.database.Table("yd_tags").Where("state = 1 and id in (?)", strings.Split(article.TagIds, ",")).Find(&articleInfo.Tags)

	return articleInfo
}



type FullArticleInfo struct {
	schema.Article
	Supports            int         `json="supports"`
	Collections         int         `json="collections"`
	TagNames            []string    `json="tagNames"`
}

func (am ArticleModel) GetFullArticleInfo(articleId int, fullArticleInfo *FullArticleInfo) {
	var articleInfo schema.Article
	var tagNames []string

	if articleId > 0 {
		am.database.First(&articleInfo, articleId)
	} else {
		am.database.Where(map[string]interface{} {"channel_id": 1}).Last(&articleInfo)
	}

	(*fullArticleInfo).Article = articleInfo

	am.database.Table("yd_tags").
		Select("name").
		Where("state = 1 and id in (?)", strings.Split(articleInfo.TagIds, ",")).
		Pluck("name", &tagNames)

	(*fullArticleInfo).TagNames = tagNames

	am.database.Table("yd_articles").
		Select("count(yd_collections.id) as collections, count(yd_supports.id) as supports").
		Where("yd_articles.id = ?", articleInfo.ID).
		Joins("left join yd_collections on yd_collections.article_id = yd_articles.id").
		Joins("left join yd_supports on yd_supports.article_id = yd_articles.id").
		Group("yd_collections.article_id, yd_supports.article_id").
		Scan(fullArticleInfo)

}

