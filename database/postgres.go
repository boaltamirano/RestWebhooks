package database

import (
	"context"
	"database/sql"
	"log"

	"github.com/RestWebkooks/models"
	_ "github.com/lib/pq"
)

type PostgresRespository struct {
	db *sql.DB
}

// Constructor par crear la instancia de struct PostgresRespository
func NewPostgresRepository(url string) (*PostgresRespository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	return &PostgresRespository{db}, nil
}

// Reciberfunction para definir los metodos del struct PostgresRespository
// Metodo InsetUser -> es la funcion que se inplementa en el UserRepository para insertar usuarios
func (repo *PostgresRespository) InsetUser(ctx context.Context, user *models.User) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO users (id, email, password) VALUES ($1, $2, $3)", user.Id, user.Email, user.Password)

	return err
}

func (repo *PostgresRespository) GetUserById(ctx context.Context, id string) (*models.User, error) {
	rowsUser, err := repo.db.QueryContext(ctx, "SELECT id, email FROM users WHERE id = $1", id)

	defer func() {
		err := rowsUser.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var user = models.User{}
	for rowsUser.Next() {
		if err = rowsUser.Scan(&user.Id, &user.Email); err == nil {
			return &user, nil
		}
	}

	if err = rowsUser.Err(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *PostgresRespository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	rowsUser, err := repo.db.QueryContext(ctx, "SELECT id, email, password FROM users WHERE email = $1", email)

	defer func() {
		err := rowsUser.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var user = models.User{}
	for rowsUser.Next() {
		if err = rowsUser.Scan(&user.Id, &user.Email); err == nil {
			return &user, nil
		}
	}

	if err = rowsUser.Err(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *PostgresRespository) Close() error {
	return repo.db.Close()
}
