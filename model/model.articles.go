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

// 文章收藏实体
type SmallArticleInfo struct {
	Id                       int             `json:"id"`
	Title                    string          `json:"title"`
	Author                   string          `json:"author"`
	AnchorName               string          `json:"anchorName"`
	CoverImg                 string          `json:"coverImg"`
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
func (am ArticleModel) GetArticleInfoById(articleId int) (ArticleInfo, error) {
	var article schema.Article

	var articleInfo ArticleInfo

	if result := am.database.Where("id = ?", articleId).First(&article).Scan(&articleInfo); result.Error != nil {
		return articleInfo, result.Error
	}

	if result := am.database.Table("yd_channels").Where("id = ?", article.ChannelId).Scan(&articleInfo.ChannelInfo); result.Error != nil {
		return articleInfo, result.Error
	}

	if result := am.database.Table("yd_tags").Where("state = 1 and id in (?)", strings.Split(article.TagIds, ",")).Find(&articleInfo.Tags); result.Error != nil {
		return articleInfo, result.Error
	}

	if result := am.database.Table("yd_supports").Where("article_id = ? and state = 1 and type = 1", articleId).Count(&articleInfo.Supports); result.Error != nil {
		return articleInfo, result.Error
	}

	if result := am.database.Table("yd_collections").Where("article_id = ? and state = 1", articleId).Count(&articleInfo.Collections); result.Error != nil {
		return articleInfo, result.Error
	}

	if archorInfo, error := NewUserModel().GetUserInfo(article.Anchor); error != nil {
		return articleInfo, error
	} else {
		articleInfo.AnchorInfo = archorInfo
	}

	return articleInfo, nil
}

// 获取指定文章相关的 n 条文章
func (am ArticleModel) GetReleasedArticlesByArticleId(articleId int, limit int) ([]SimpleArticleInfo, error) {
	var releasedArticles []SimpleArticleInfo

	releasedArticleIds, error := am.GetReleaseArticleIdsByArticleId(articleId, limit)

	if error != nil {
		return releasedArticles, error
	}

	rows, error := am.database.Table("yd_articles").
		Select("yd_articles.id, title, author, yd_users.username, during, play_number, cover_img, audio, content_text").
		Where("yd_articles.id in (?)", releasedArticleIds).
		Joins("inner join yd_users on yd_articles.anchor = yd_users.id").
		Rows()

	if error != nil {
		return releasedArticles, error
	}

	for rows.Next() {
		var article SimpleArticleInfo = SimpleArticleInfo{}

		if error := rows.Scan(&article.Id, &article.Title, &article.Author, &article.AnchorName, &article.During, &article.PlayNumber, &article.CoverImg, &article.Audio, &article.ContentText); error != nil {
			return releasedArticles, error
		}

		releasedArticles = append(releasedArticles, article)
	}

	return releasedArticles, nil
}

func (am ArticleModel) GetReleaseArticleIdsByArticleId(articleId int, limit int) ([]int, error) {
	var articleTagIds []string

	if result := am.database.Table("yd_articles").Where("id = ?", articleId).Pluck("tag_ids", &articleTagIds); result.Error != nil {
		return []int{}, result.Error
	}

	var mainTagId string = strings.Split(articleTagIds[0], ",")[0]

	var articleIds []int

	if result := am.database.Table("yd_articles").Where("tag_ids LIKE ?", "%" + mainTagId + "%").Limit(limit).Pluck("id", &articleIds); result.Error != nil {
		return articleIds, result.Error
	}

	return articleIds, nil
}

// 获取指定文章其他类型的最新文章列表
func (am ArticleModel) GetOtherChannelLastArticlesByArticleId(articleId int) ([]SimpleArticleInfo, error) {
	var channelId []int
	var otherArticles []SimpleArticleInfo
	var query string = "yd_articles.id, title, author, yd_users.username, during, play_number, cover_img, audio, content_text, max(yd_articles.id)"

	if result := am.database.Table("yd_articles").Where("id = ?", articleId).Pluck("channel_id", &channelId); result.Error != nil {
		return otherArticles, result.Error
	}

	rows, error := am.database.Table("yd_articles").Select(query).Where("yd_articles.channel_id != ?", channelId[0]).
		Joins("inner join yd_users on yd_articles.anchor = yd_users.id").
		Group("channel_id").
		Rows()

	if error != nil {
		return otherArticles, error
	}

	for rows.Next() {
		var article SimpleArticleInfo = SimpleArticleInfo{}
		var maxId int

		if error := rows.Scan(&article.Id, &article.Title, &article.Author, &article.AnchorName, &article.During, &article.PlayNumber, &article.CoverImg, &article.Audio, &article.ContentText, &maxId); error != nil {
			return otherArticles, error
		}

		otherArticles = append(otherArticles, article)
	}

	return otherArticles, nil
}

// 获取指定频道的文章列表
func (am ArticleModel) GetArticlesByChannelId(channelId, page, size int) ([]SimpleArticleInfo, error) {
	var articles []SimpleArticleInfo

	rows, error := am.database.Table("yd_articles").
		Select("yd_articles.id, title, author, yd_users.username, during, play_number, cover_img, audio, content_text").
		Where("channel_id = ?", channelId).
		Joins("inner join yd_users on yd_users.id = yd_articles.anchor").
		Limit(size).
		Offset(page * size).
		Rows()

	if error != nil {
		return articles, error
	}

	for rows.Next() {
		var article SimpleArticleInfo = SimpleArticleInfo{}

		if error := rows.Scan(&article.Id, &article.Title, &article.Author, &article.AnchorName, &article.During, &article.PlayNumber, &article.CoverImg, &article.Audio, &article.ContentText); error != nil {
			return articles, error
		}

		articles = append(articles, article)
	}

	return articles, nil
}

// 获取指定频道文章的总条数
func (am ArticleModel) GetArticleCountByChannelId(channelId int) (int, error) {
	var count int = 0

	if result := am.database.Table("yd_articles").Where("channel_id = ?", channelId).Count(&count); result.Error != nil {
		return count, result.Error
	}

	return count, nil
}

// 根据点赞数目查询指定类目最受欢迎的文章
func (am ArticleModel) GetTopArticles(channelId, numbers int) ([]SimpleArticleInfo, error) {
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
		return articles, error
	}

	for rows.Next() {
		var article SimpleArticleInfo = SimpleArticleInfo{}
		var count int

		if error := rows.Scan(&count, &article.Id, &article.Title, &article.Author, &article.AnchorName, &article.During, &article.PlayNumber, &article.CoverImg, &article.Audio, &article.ContentText); error != nil {
			return articles, error
		}

		articles = append(articles, article)
	}

	// 如果小于指定的条数，取最新的文章补充上去
	if len(articles) < numbers {
		if lastNewArticles, error := am.GetLastNewArticles(channelId, numbers - len(articles)); error != nil {
			return articles, error
		} else {
			articles = append(articles, lastNewArticles...)
		}
	}

	return articles, nil
}

// 获取指定类目最新的文章
func (am ArticleModel) GetLastNewArticles(channelId, numbers int) ([]SimpleArticleInfo, error) {
	var articles []SimpleArticleInfo

	rows, error := am.database.Table("yd_articles").
		Select("yd_articles.id, title, author, yd_users.username, during, play_number, cover_img, audio, content_text").
		Joins("inner join yd_users on yd_articles.anchor = yd_users.id").
		Where("channel_id = ?", channelId).
		Limit(numbers).
		Order("yd_articles.created_at desc").
		Rows()

	if error != nil {
		return articles, error
	}

	for rows.Next() {
		var article SimpleArticleInfo = SimpleArticleInfo{}

		if error := rows.Scan(&article.Id, &article.Title, &article.Author, &article.AnchorName, &article.During, &article.PlayNumber, &article.CoverImg, &article.Audio, &article.ContentText); error != nil {
			return articles, error
		}

		articles = append(articles, article)
	}

	return articles, nil
}
