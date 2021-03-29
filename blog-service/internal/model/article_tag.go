package model

type ArticleTag struct {
	*Model
	TagID     uint32 `json:"tag_id"`
	ArticleTD uint32 `json:"article_td"`
}

func (a ArticleTag) TableName() string {
	return "blog_article_tag"
}
