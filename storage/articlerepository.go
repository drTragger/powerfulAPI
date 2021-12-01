package storage

import (
	"database/sql"
	"fmt"
	"github.com/drTragger/powerfulAPI/internal/app/models"
	"log"
)

// ArticleRepository instance of Article repository (model interface)
type ArticleRepository struct {
	storage *Storage
}

var (
	tableArticle = "articles"
)

func (ar *ArticleRepository) Create(a *models.Article) (*models.Article, error) {
	query := fmt.Sprintf("INSERT INTO %s (title, author, content) VALUES ($1, $2, $3) RETURNING id", tableArticle)
	if err := ar.storage.db.QueryRow(query, a.Title, a.Author, a.Content).Scan(&a.ID); err != nil {
		return nil, err
	}
	return a, nil
}

func (ar *ArticleRepository) DeleteById(id int) (*models.Article, error) {
	article, ok, err := ar.FindArticleById(id)
	if err != nil {
		return nil, err
	}
	if ok {
		query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", tableArticle)
		_, err := ar.storage.db.Exec(query, id)
		if err != nil {
			return nil, err
		}
	}
	return article, nil
}

func (ar *ArticleRepository) FindArticleById(id int) (*models.Article, bool, error) {
	articles, err := ar.SelectAll()
	var found bool
	if err != nil {
		return nil, found, err
	}
	var articleFound *models.Article
	for _, a := range articles {
		if a.ID == id {
			articleFound = a
			found = true
			break
		}
	}
	return articleFound, found, nil
}

func (ar *ArticleRepository) SelectAll() ([]*models.Article, error) {
	query := fmt.Sprintf("SELECT * FROM %s", tableArticle)
	rows, err := ar.storage.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println("Error happened during closing row connection:", err)
		}
	}(rows)
	articles := make([]*models.Article, 0)
	for rows.Next() {
		a := models.Article{}
		err := rows.Scan(&a.ID, &a.Title, &a.Author, &a.Content)
		if err != nil {
			log.Println(err)
			continue
		}
		articles = append(articles, &a)
	}
	return articles, nil
}
