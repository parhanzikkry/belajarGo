package article

import (
	"belajarGo/connection/database"
	"context"
	"database/sql"
	"time"
)

type (
	ArticleDataResponse struct {
		Id           int64     `json:"id", db:"article_id"`
		Title        string    `json:"title", db:"article_title"`
		Content      string    `json:"content", db:"content_article"`
		Created_at   time.Time `json:"created_at", db:"created_at"`
		Updated_date time.Time `json;"updated_at", db:"updated_at"`
	}
)

var temp ArticleDataResponse
var result []ArticleDataResponse

func (article *pkgArticle) GetAllArticle() ([]ArticleDataResponse, error) {
	strConn, err := database.SetConnection()
	if err != nil {
		return result, err
	}

	dbConn, err := sql.Open(cfg.DatabaseConfig.DriverName, strConn)
	defer dbConn.Close()
	if err != nil {
		return result, err
	}

	query := GetAllArticleQuery()

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	datas, err := dbConn.QueryContext(ctx, query)
	if err != nil {
		return result, err
	}

	for datas.Next() {
		if err := datas.Scan(
			&temp.Id,
			&temp.Title,
			&temp.Content,
			&temp.Created_at,
			&temp.Updated_date,
		); err != nil {
			return result, err
		}
		result = append(result, temp)
	}

	return result, err
}

func (article *pkgArticle) GetArticleById(id int64) (ArticleDataResponse, error) {
	result := temp
	strConn, err := database.SetConnection()
	if err != nil {
		return result, err
	}

	dbConn, err := sql.Open(cfg.DatabaseConfig.DriverName, strConn)
	defer dbConn.Close()
	if err != nil {
		return result, err
	}

	query := GetOneArticleQuery()

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	datas, err := dbConn.QueryContext(ctx, query, id)
	if err != nil {
		return result, err
	}

	if err := datas.Scan(
		&result.Id,
		&result.Title,
		&result.Content,
		&result.Created_at,
		&result.Updated_date,
	); err != nil {
		return result, err
	}

	return result, err
}
