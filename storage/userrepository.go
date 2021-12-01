package storage

import (
	"database/sql"
	"fmt"
	"github.com/drTragger/powerfulAPI/internal/app/models"
	"log"
)

// UserRepository instance of User repository (model interface)
type UserRepository struct {
	storage *Storage
}

var (
	tableUser = "users"
)

// Create new user in DB
func (ur *UserRepository) Create(u *models.User) (*models.User, error) {
	query := fmt.Sprintf("INSERT INTO %s (login, password) VALUES ($1, $2) RETURNING id", tableUser)
	if err := ur.storage.db.QueryRow(query, u.Login, u.Password).Scan(&u.ID); err != nil {
		return nil, err
	}
	return u, nil
}

func (ur *UserRepository) FindByLogin(login string) (*models.User, bool, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE login=$1", tableUser)
	user := models.User{}
	row := ur.storage.db.QueryRow(query, login)
	switch err := row.Scan(&user.ID, &user.Login, &user.Password); err {
	case sql.ErrNoRows:
		return nil, false, nil
	case nil:
		return &user, true, nil
	default:
		return nil, false, err
	}
}

func (ur *UserRepository) SelectAll() ([]*models.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s", tableUser)
	rows, err := ur.storage.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println("Error during closing connection:", err)
		}
	}(rows)

	users := make([]*models.User, 0)
	for rows.Next() {
		u := models.User{}
		err := rows.Scan(&u.ID, &u.Login, &u.Password)
		if err != nil {
			log.Println(err)
			continue
		}
		users = append(users, &u)
	}
	return users, nil
}
