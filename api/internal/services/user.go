package services

import (
	"time"

	"github.com/google/uuid"
	"github.com/jnaraujo/pay/internal/db"
	"github.com/jnaraujo/pay/internal/errs"
	"github.com/jnaraujo/pay/internal/models"
)

func CreateUser(name string) (*models.User, error) {
	user := models.User{
		Id:        uuid.New(),
		Name:      name,
		CreatedAt: time.Now(),
	}

	_, err := db.DB.Exec("insert into users (id, name, created_at) values (?, ?, ?)", user.Id, user.Name, user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func FindUserById(id uuid.UUID) (*models.User, error) {
	user := models.User{}
	rows, err := db.DB.Query("select id, name, created_at, balance from users where id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rows.Next()
	err = rows.Scan(&user.Id, &user.Name, &user.CreatedAt, &user.Balance)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func TransferPayment(from uuid.UUID, to uuid.UUID, amount int) error {
	tx, err := db.DB.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}
	var sender models.User
	var receiver models.User

	rows, err := tx.Query("select balance from users where id = ?", from)
	if err != nil {
		tx.Rollback()
		return err
	}

	rows.Next()
	err = rows.Scan(&sender.Balance)
	if err != nil {
		tx.Rollback()
		return err
	}

	rows, err = tx.Query("select balance from users where id = ?", to)
	if err != nil {
		tx.Rollback()
		return err
	}

	rows.Next()
	err = rows.Scan(&receiver.Balance)
	if err != nil {
		tx.Rollback()
		return err
	}

	if sender.Balance < amount {
		tx.Rollback()
		return errs.ErrInsufficientBalance
	}

	_, err = tx.Exec("update users set balance = balance - ? where id = ?", amount, from)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("update users set balance = balance + ? where id = ?", amount, to)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func BankTransfer(to uuid.UUID, amount int) error {
	tx, err := db.DB.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}

	var receiver models.User

	rows, err := tx.Query("select balance from users where id = ?", to)
	if err != nil {
		tx.Rollback()
		return err
	}
	rows.Scan(&receiver.Balance)

	_, err = tx.Exec("update users set balance = balance + ? where id = ?", amount, to)
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
