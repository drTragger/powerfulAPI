package storage

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type Storage struct {
	config            *Config
	db                *sql.DB
	userRepository    *UserRepository
	articleRepository *ArticleRepository
}

func New(config *Config) *Storage {
	return &Storage{
		config: config,
	}
}

func (storage *Storage) Open() error {
	db, err := sql.Open("postgres", storage.config.DatabaseURI)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	storage.db = db
	log.Println("Database connection has been created successfully")
	return nil
}

func (storage *Storage) Close() {
	if err := storage.db.Close(); err != nil {
		log.Println("Error during closing DB connection:", err)
	}
}

// User Public Repo for models.User
func (storage *Storage) User() *UserRepository {
	if storage.userRepository != nil {
		return storage.userRepository
	}
	storage.userRepository = &UserRepository{
		storage: storage,
	}
	return storage.userRepository
}

// Article Public Repo for models.Article
func (storage *Storage) Article() *ArticleRepository {
	if storage.articleRepository != nil {
		return storage.articleRepository
	}
	storage.articleRepository = &ArticleRepository{
		storage: storage,
	}
	return storage.articleRepository
}
