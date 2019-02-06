package model

import (
	"strings"
	"yuedu/schema"
	"yuedu/database"
	"github.com/jinzhu/gorm"
	"log"
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

	rows, error := am.database.Table("yd_articles").
		Select("yd_articles.id, title, author, yd_users.username, during, play_number, cover_img, audio, content_text").
		Where("yd_articles.id in (?)", releasedArticleIds).
		Joins("inner join yd_users on yd_articles.anchor = yd_users.id").
		Rows()

	if error != nil {
		log.Println(error)
	}

	for rows.Next() {
		var article SimpleArticleInfo = SimpleArticleInfo{}

		if error := rows.Scan(&article.Id, &article.Title, &article.Author, &article.AnchorName, &article.During, &article.PlayNumber, &article.CoverImg, &article.Audio, &article.ContentText); error != nil {
			log.Println(error)
		}

		releasedArticles = append(releasedArticles, article)
	}

	return releasedArticles
}

func (am ArticleModel) GetReleaseArticleIdsByArticleId(articleId int, limit int) []int {
	var articleTagIds []string

	am.database.Table("yd_articles").Where("id = ?", articleId).Pluck("tag_ids", &articleTagIds)

	var mainTagId string = strings.Split(articleTagIds[0], ",")[0]

	var articleIds []int

	am.database.Table("yd_articles").Where("tag_ids LIKE ?", "%" + mainTagId + "%").Limit(limit).Pluck("id", &articleIds)

	return articleIds
}


// 获取指定文章其他类型的最新文章列表
func (am ArticleModel) GetOtherChannelLastArticlesByArticleId(articleId int) []SimpleArticleInfo {
	var channelId []int
	var otherArticles []SimpleArticleInfo
	var query string = "yd_articles.id, title, author, yd_users.username, during, play_number, cover_img, audio, content_text, max(yd_articles.id)"

	am.database.Table("yd_articles").Where("id = ?", articleId).Pluck("channel_id", &channelId)

	rows, error := am.database.Table("yd_articles").Select(query).Where("yd_articles.channel_id != ?", channelId[0]).
		Joins("inner join yd_users on yd_articles.anchor = yd_users.id").
		Group("channel_id").
		Rows()

	if error != nil {
		log.Println(error)
	}

	for rows.Next() {
		var article SimpleArticleInfo = SimpleArticleInfo{}
		var maxId int

		if error := rows.Scan(&article.Id, &article.Title, &article.Author, &article.AnchorName, &article.During, &article.PlayNumber, &article.CoverImg, &article.Audio, &article.ContentText, &maxId); error != nil {
			log.Println(error)
		}

		otherArticles = append(otherArticles, article)
	}

	return otherArticles;
}


// 获取指定频道的文章列表
func (am ArticleModel) GetArticlesByChannelId(channelId, page, size int) []SimpleArticleInfo {
	var articles []SimpleArticleInfo

	rows, error := am.database.Table("yd_articles").
		Select("yd_articles.id, title, author, yd_users.username, during, play_number, cover_img, audio, content_text").
		Where("channel_id = ?", channelId).
		Joins("inner join yd_users on yd_users.id = yd_articles.anchor").
		Limit(size).
		Offset(page * size).
		Rows()

	if error != nil {
		log.Println(error)
		return articles
	}

	for rows.Next() {
		var article SimpleArticleInfo = SimpleArticleInfo{}

		if error := rows.Scan(&article.Id, &article.Title, &article.Author, &article.AnchorName, &article.During, &article.PlayNumber, &article.CoverImg, &article.Audio, &article.ContentText); error != nil {
			log.Println(error)
			return articles
		}

		articles = append(articles, article)
	}

	return articles
}


// 获取指定频道文章的总条数
func (am ArticleModel) GetArticleCountByChannelId(channelId int) int {
	var count int = 0

	am.database.Table("yd_articles").Where("channel_id = ?", channelId).Count(&count)

	return count
}


// 根据点赞数目查询指定类目最受欢迎的文章
func (am ArticleModel) GetTopArticles(channelId, numbers int) []SimpleArticleInfo {
	var articles []SimpleArticleInfo

	rows, error := am.database.Table("yd_supports").
		Select("count(*) as number, yd_articles.id, title, author, yd_users.username, during, play_number, cover_img, audio, content_text").
		Joins("inner join yd_articles on yd_articles.id = yd_supports.article_id and channel_id = ?", channelId).
		Joins("inner join yd_users on yd_users.id = yd_articles.anchor").
		Group("article_id").
		Order("number desc").
		Limit(numbers).
		Rows()

	if error != nil {
		log.Println(error)
		return articles
	}

	for rows.Next() {
		var article SimpleArticleInfo = SimpleArticleInfo{}
		var count int

		if error := rows.Scan(&count, &article.Id, &article.Title, &article.Author, &article.AnchorName, &article.During, &article.PlayNumber, &article.CoverImg, &article.Audio, &article.ContentText); error != nil {
			log.Println(error)
			return articles
		}

		articles = append(articles, article)
	}

	if len(articles) < numbers {
		articles = append(articles, am.GetLastNewArticles(channelId, numbers - len(articles))...)
	}

	return articles
}

// 获取指定类目最新的文章
func (am ArticleModel) GetLastNewArticles(channelId, numbers int) []SimpleArticleInfo {
	var articles []SimpleArticleInfo

	rows, error := am.database.Table("yd_articles").
		Select("yd_articles.id, title, author, yd_users.username, during, play_number, cover_img, audio, content_text").
		Joins("inner join yd_users on yd_articles.anchor = yd_users.id").
		Where("channel_id = ?", channelId).
		Limit(numbers).
		Order("yd_articles.created_at desc").
		Rows()

	if error != nil {
		log.Println(error)
		return  articles
	}

	for rows.Next() {
		var article SimpleArticleInfo = SimpleArticleInfo{}

		if error := rows.Scan(&article.Id, &article.Title, &article.Author, &article.AnchorName, &article.During, &article.PlayNumber, &article.CoverImg, &article.Audio, &article.ContentText); error != nil {
			log.Println(error)
			return articles
		}

		articles = append(articles, article)
	}

	return articles
}
