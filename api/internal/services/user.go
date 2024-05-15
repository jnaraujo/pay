package services

import (
	"time"

	"github.com/google/uuid"
	"github.com/jnaraujo/pay/internal/db"
	"github.com/jnaraujo/pay/internal/models"
)

func CreateUser(user models.User) (*models.User, error) {
	user.Id = uuid.New()
	user.CreatedAt = time.Now()

	_, err := db.DB.Exec("insert into users (id, name, created_at) values (?, ?, ?);", user.Id, user.Name, user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
