package postgres

import (
	"context"
	"errors"

	"github.com/Pizhlo/wb-L0/errs"
	"github.com/Pizhlo/wb-L0/models"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage interface {
	CreateTable() error
	GetOrderByID(ctx context.Context, id int) (models.Order, error)
}

type DB struct {
	*pgxpool.Pool
}

// postgresql://root:secret@localhost:8081/wb_db?sslmode=disable

func New(conn *pgxpool.Pool, dbAddress string, migratePath string) *DB {
	db := &DB{conn}
	return db
}

func (db *DB) Close() {
	db.Pool.Close()
}

// func runMigrations(dsn string, migratePath string) error {
// 	m, err := migrate.New(fmt.Sprintf("file:///%s", migratePath), dsn)
// 	if err != nil {
// 		return err
// 	}

// 	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
// 		return err
// 	}

// 	return nil
// }

// func (db *DB) CreateTable() error {
// 	ctx, cancel := context.WithTimeout(context.TODO(), 3*time.Second)
// 	defer cancel()

// 	q := `CREATE TABLE IF NOT EXISTS "user"(id serial NOT NULL,
// 		login text NOT NULL,
// 		password text NOT NULL,
// 		created TIMESTAMP NOT NULL DEFAULT 'now()'
// 	);

// CREATE UNIQUE INDEX IF NOT EXISTS "user_login" ON "user" ("login");`

// 	txOptions := pgx.TxOptions{
// 		IsoLevel: pgx.ReadCommitted,
// 	}

// 	tx, err := db.BeginTx(ctx, txOptions)
// 	if err != nil {
// 		return err
// 	}

// 	_, err = tx.Exec(ctx, q)
// 	if err != nil {
// 		return err
// 	}

// 	defer tx.Rollback(ctx)

// 	return tx.Commit(ctx)
// }

func (db *DB) GetOrderByID(ctx context.Context, id uuid.UUID) (models.Order, error) {
	order := models.Order{
		OrderUIID: id,
	}

	q := `SELECT * FROM orders WHERE order_id = $1;`

	row := db.QueryRow(ctx, q, id)
	err := row.Scan(&order.ID, &order.OrderUIID, &order.TrackNumber, &order.Entry, &order.Delivery.ID, &order.Payment.ID,
		&order.Locale, &order.InternalSignature, &order.CustomerID, &order.DeliveryService, &order.ShardKey,
		&order.SmID, &order.DateCreated, &order.OofShard)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return order, errs.NotFound
		}
		return order, err
	}

	delivery, err := db.getDeliveryByOrderID(ctx, id)
	if err != nil {
		return order, err
	}

	order.Delivery = delivery

	payment, err := db.getPaymentByOrderID(ctx, id)
	if err != nil {
		return order, err
	}

	order.Payment = payment

	return order, nil
}

func (db *DB) getDeliveryByOrderID(ctx context.Context, orderId uuid.UUID) (models.Delivery, error) {
	delivery := models.Delivery{}

	q := `SELECT * FROM delivery WHERE id = (SELECT delivery_id FROM orders WHERE order_id = $1);`

	row := db.QueryRow(ctx, q, orderId)
	err := row.Scan(&delivery.ID, &delivery.Name, &delivery.Phone, &delivery.Zip, &delivery.City, &delivery.Address, &delivery.Region,
		&delivery.Email)

	if err != nil {
		return delivery, err
	}

	return delivery, nil
}

func (db *DB) getPaymentByOrderID(ctx context.Context, orderId uuid.UUID) (models.Payment, error) {
	payment := models.Payment{}

	q := `SELECT * FROM payments WHERE id = (SELECT payment_id FROM orders WHERE order_id = $1);`

	row := db.QueryRow(ctx, q, orderId)
	err := row.Scan(&payment.ID, &payment.Transaction, &payment.RequestID, &payment.Currency, &payment.Provider, &payment.Amount,
		&payment.PaymentDate, &payment.Bank, &payment.DeliveryCost, &payment.GoodsTotal, &payment.CustomFee)
	if err != nil {
		return payment, err
	}

	return payment, nil
}
