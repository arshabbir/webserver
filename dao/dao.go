package dao

import (
	"context"
	"database/sql"
	"time"

	"fmt"
	"os"
	"webservermod/model"

	"github.com/jackc/pgx/v5"
)

type userDB struct {
	conn *pgx.Conn
}

type UserDB interface {
	CreateUser(model.User) error
	GetUserByEmail(email string) (*model.User, error)
	UpdateUser(Id string, user model.User) error
	DeleteUser(Id string) error
	Close() error
}

func NewUserDB(connStr string) UserDB {
	conn, err := openDB(connStr)
	if err != nil {
		return nil
	}
	return &userDB{conn: conn}
}

func (u *userDB) CreateUser(user model.User) error {
	// Implement the creation logic

	_, err := u.conn.Exec(context.Background(), "insert into users (id, first_name, last_name, email, password, is_admin, created_at, updated_at) values($1, $2, $3, $4, $5, $6, $7, $8)", user.Id, user.First_name, user.Last_name, user.Email, user.Password, user.Is_admin, time.Now(), time.Now())
	fmt.Println(err)
	return err
}

func (u *userDB) Close() error {
	return u.conn.Close(context.Background())
}

func (u *userDB) GetUserByEmail(email string) (*model.User, error) {
	query := "SELECT id, first_name, last_name, is_admin FROM users WHERE email = $1"

	row := u.conn.QueryRow(context.Background(), query, email)

	var user model.User
	err := row.Scan(&user.Id, &user.First_name, &user.Last_name, &user.Is_admin)
	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		return nil, err
	}

	fmt.Println(user)
	return &user, nil
}

func (u *userDB) UpdateUser(Id string, user model.User) error {
	return nil
}

func (u *userDB) DeleteUser(Id string) error {
	return nil
}

func openDB(connstr string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), connstr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}

	if err := conn.Ping(context.Background()); err != nil {
		return nil, err
	}
	//defer conn.Close(context.Background())
	return conn, nil

}
