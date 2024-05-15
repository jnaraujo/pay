package services

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jnaraujo/pay/internal/db"
	"github.com/jnaraujo/pay/internal/models"
)

func DoesGameExists(id uuid.UUID) bool {
	res, err := db.DB.Exec("select (id) from games where id = ?", id)
	if err != nil {
		log.Panic(err)
		return false
	}

	affect, err := res.RowsAffected()
	if err != nil {
		log.Panic(err)
		return false
	}

	return affect > 0
}

func CreateGame(owner_id uuid.UUID) (*models.Game, error) {
	game := models.Game{
		Id:        uuid.New(),
		OwnerId:   owner_id,
		CreatedAt: time.Now(),
	}
	_, err := db.DB.Exec("insert into games (id, owner_id, created_at) values (?, ?, ?)", game.Id, game.OwnerId, game.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &game, nil
}

func FindGameById(id uuid.UUID) *models.Game {
	rows, err := db.DB.Query("select (id, created_at, owner_id) from games where id = ? limit 1", id)
	if err != nil {
		log.Panic(err)
		return nil
	}
	defer rows.Close()

	var game models.Game
	err = rows.Scan(&game.Id, &game.CreatedAt, &game.OwnerId)
	if err != nil {
		log.Panic(err)
		return nil
	}

	return &game
}
