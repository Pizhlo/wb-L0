package postgres

import (
	"context"
	"time"

	"github.com/Pizhlo/wb-L0/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage interface {
	CreateTable() error
	GetUserByID(ctx context.Context, id int) (models.User, error)
}

type DB struct {
	*pgxpool.Pool
}

func New(conn *pgxpool.Pool) (*DB, error) {
	db := &DB{conn}

	return db, db.CreateTable()

}

func (db *DB) CreateTable() error {
	ctx, cancel := context.WithTimeout(context.TODO(), 3*time.Second)
	defer cancel()

	q := `CREATE TABLE IF NOT EXISTS "user"(id serial NOT NULL,
		login text NOT NULL,
		password text NOT NULL,
		created TIMESTAMP NOT NULL
	);
	
CREATE UNIQUE INDEX IF NOT EXISTS ON "user"(login);`

	txOptions := pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	}

	tx, err := db.BeginTx(ctx, txOptions)
	if err != nil {
		return err
	}

	tx.Exec(ctx, q)

	defer tx.Rollback(ctx)

	return tx.Commit(ctx)
}

func (db *DB) GetUserByID(ctx context.Context, id int) (models.User, error) {
	user := models.User{}

	q := `SELECT login, password, created FROM "user" WHERE id = $1`

	err := db.QueryRow(ctx, q, id).Scan(&user.Login, &user.Password, &user.Created)
	if err != nil {
		return user, err
	}

	return user, nil

}
