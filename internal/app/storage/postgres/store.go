package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/Pizhlo/wb-L0/errs"
	"github.com/Pizhlo/wb-L0/models"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage interface {
	GetOrderByID(ctx context.Context, id uuid.UUID) (models.Order, error)
	SaveOrder(ctx context.Context, order models.Order) error
	SaveItem(ctx context.Context, items []models.Item) error
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

func (db *DB) SaveOrder(ctx context.Context, order models.Order) error {
	fmt.Println("starting saving order...")
	deliveryID, err := db.saveDelivery(ctx, order.Delivery)
	if err != nil {
		return err
	}

	paymentID, err := db.savePayment(ctx, order.Payment)
	if err != nil {
		return err
	}

	q := `INSERT INTO orders (order_id, track_number, entry, delivery_id, payment_id, locale, internal_signature, customer_id, delivery_service, shard_key, sm_id, date_created, oof_shard)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`

	_, err = db.Exec(ctx, q, order.OrderUIID, order.TrackNumber, order.Entry, deliveryID, paymentID, order.Locale, order.InternalSignature, order.CustomerID, order.DeliveryService,
		order.ShardKey, order.SmID, order.DateCreated, order.OofShard)
	if err != nil {
		fmt.Println("an error accured while saving order: ", err)
		return err
	}

	err = db.SaveItem(ctx, order.Items)
	if err != nil {
		return err
	}

	fmt.Println("order saved successfully")
	return nil
}

// works
func (db *DB) saveDelivery(ctx context.Context, delivery models.Delivery) (int, error) {
	fmt.Println("starting saving delivery...")
	q := `INSERT INTO delivery (name, phone, zip, city, address, region, email) VALUES($1, $2, $3, $4, $5, $6, $7)
	RETURNING id`

	row := db.QueryRow(ctx, q, delivery.Name, delivery.Phone, delivery.Zip, delivery.City, delivery.Address, delivery.Region, delivery.Email)
	var id int

	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	fmt.Println("delivery saved successfully")

	return id, nil
}

// works
func (db *DB) savePayment(ctx context.Context, payment models.Payment) (int, error) {
	fmt.Println("starting saving payment...")

	q := `INSERT INTO payments (transaction, request_id, currency, provider, amount, payment_date, bank, delivery_cost, goods_total, custom_fee) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`

	var id int

	row := db.QueryRow(ctx, q, payment.Transaction, payment.RequestID, payment.Currency, payment.Provider, payment.Amount, payment.PaymentDate, payment.Bank, payment.DeliveryCost, payment.GoodsTotal, payment.CustomFee)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	fmt.Println("payment saved successfully")
	return id, nil
}

func (db *DB) SaveItem(ctx context.Context, items []models.Item) error {
	fmt.Println("starting saving items...")

	//fmt.Printf("got these items: %+v\n", items)

	q := `INSERT INTO items (chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	for _, item := range items {
		_, err := db.Exec(ctx, q, item.ChrtId, item.TrackNumber, item.Price, item.RID, item.Name, item.Sale, item.Size, item.TotalPrice, item.NmID, item.Brand, item.Status)
		if err != nil {
			fmt.Println("an error accured while saving item: ", err)
			return err
		}
	}

	fmt.Println("items saved successfully")
	return nil
}
