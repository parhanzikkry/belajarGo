package article

import (
	"belajarGo/config"
)

type (
	PkgArticle interface {
		GetAllArticle() ([]ArticleDataResponse, error)
		GetArticleById(int64) (ArticleDataResponse, error)
		PostArticle() error
	}
	pkgArticle struct{}
)

var cfg config.Config

func New() PkgArticle {
	cfg = config.GetConfig()
	return &pkgArticle{}
}
