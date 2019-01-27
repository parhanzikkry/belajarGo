package handler

import (
	pkgcarticle "belajarGo/controllers/article"
)

var carticle pkgcarticle.PkgArticle

func Init() {
	if carticle == nil {
		carticle = pkgcarticle.New()
	}
}
