package article

import (
	"errors"
)

type (
	ArticleDataRequest struct {
		Title   string `json:"title", db:"article_title"`
		Content string `json:"content", db:"content_article"`
	}
)

func (p *pkgArticle) PostArticle() error {
	err := errors.New("Fake Error")
	return err
}
