package article

func GetAllArticleQuery() string {
	return `SELECT * FROM article`
}

func GetOneArticleQuery() string {
	return `SELECT * FROM article WHERE article_id = $1`
}
