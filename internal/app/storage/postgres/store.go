package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/Pizhlo/wb-L0/internal/app/errs"
	models "github.com/Pizhlo/wb-L0/internal/model"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	pkg_errs "github.com/pkg/errors"
)

type Repo interface {
	GetOrderByID(ctx context.Context, id uuid.UUID) (*models.Order, error)
	SaveOrder(ctx context.Context, order models.Order) (*models.Order, error)
	GetPaymentByOrderID(ctx context.Context, orderId uuid.UUID) (*models.Payment, error)
	SavePayment(ctx context.Context, payment models.Payment) (int, error)
	GetDeliveryByOrderID(ctx context.Context, orderId uuid.UUID) (*models.Delivery, error)
	SaveDelivery(ctx context.Context, delivery models.Delivery) (int, error)
	SaveItems(ctx context.Context, items []models.Item) error
	GetItemByTrackNumber(ctx context.Context, trackNumber string) ([]models.Item, error)
	GetAll(ctx context.Context) ([]models.Order, error)
}

type Store struct {
	*pgxpool.Pool
}

func New(conn *pgxpool.Pool) *Store {
	db := &Store{conn}
	return db
}

func (db *Store) Close() {
	db.Pool.Close()
}

func (db *Store) GetOrderByID(ctx context.Context, id uuid.UUID) (*models.Order, error) {
	fmt.Println("getting  order from db")
	order := &models.Order{}

	q := `SELECT * FROM orders WHERE order_id = $1;`

	row := db.QueryRow(ctx, q, id)
	err := row.Scan(&order.OrderUIID, &order.TrackNumber, &order.Entry, &order.Delivery.ID, &order.Payment.ID,
		&order.Locale, &order.InternalSignature, &order.CustomerID, &order.DeliveryService, &order.ShardKey,
		&order.SmID, &order.DateCreated, &order.OofShard)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			fmt.Println("not found in db")
			return order, pkg_errs.Wrap(errs.NotFound, "not found order")
		}
		return order, err
	}

	delivery, err := db.GetDeliveryByOrderID(ctx, id)
	if err != nil {
		return order, pkg_errs.Wrap(err, "err while getting delivery from db")
	}

	order.Delivery = *delivery

	payment, err := db.GetPaymentByOrderID(ctx, id)
	if err != nil {
		return order, pkg_errs.Wrap(err, "err while getting payment from db")
	}

	order.Payment = *payment

	items, err := db.GetItemByTrackNumber(ctx, order.TrackNumber)
	if err != nil {
		return order, pkg_errs.Wrap(err, "err while getting items from db")
	}

	order.Items = items

	return order, nil
}

func (db *Store) GetPaymentByOrderID(ctx context.Context, orderId uuid.UUID) (*models.Payment, error) {
	payment := &models.Payment{}

	q := `SELECT * FROM payments WHERE id = (SELECT payment_id FROM orders WHERE order_id = $1);`

	row := db.QueryRow(ctx, q, orderId)
	err := row.Scan(&payment.ID, &payment.Transaction, &payment.RequestID, &payment.Currency, &payment.Provider, &payment.Amount,
		&payment.PaymentDate, &payment.Bank, &payment.DeliveryCost, &payment.GoodsTotal, &payment.CustomFee)
	if err != nil {
		return payment, pkg_errs.Wrap(err, "err while getting payment by id from db")
	}

	return payment, nil
}

func (db *Store) SaveOrder(ctx context.Context, order models.Order) (*models.Order, error) {
	fmt.Println("starting saving order...")
	deliveryID, err := db.SaveDelivery(ctx, order.Delivery)
	if err != nil {
		return &order, err
	}

	paymentID, err := db.SavePayment(ctx, order.Payment)
	if err != nil {
		return &order, err
	}

	q := `INSERT INTO orders (order_id, track_number, entry, delivery_id, payment_id, locale, internal_signature, customer_id, delivery_service, shard_key, sm_id, date_created, oof_shard)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`

	_, err = db.Exec(ctx, q, order.OrderUIID, order.TrackNumber, order.Entry, deliveryID, paymentID, order.Locale, order.InternalSignature, order.CustomerID, order.DeliveryService,
		order.ShardKey, order.SmID, order.DateCreated, order.OofShard)
	if err != nil {
		return &order, err
	}

	err = db.SaveItems(ctx, order.Items)
	if err != nil {
		return &order, err
	}

	order.Payment.ID = paymentID
	order.Delivery.ID = deliveryID

	fmt.Println("order saved successfully")
	return &order, nil
}

func (db *Store) SavePayment(ctx context.Context, payment models.Payment) (int, error) {
	fmt.Println("starting saving payment...")

	q := `INSERT INTO payments (transaction, request_id, currency, provider, amount, payment_date, bank, delivery_cost, goods_total, custom_fee) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`

	var id int

	row := db.QueryRow(ctx, q, payment.Transaction, payment.RequestID, payment.Currency, payment.Provider, payment.Amount, payment.PaymentDate, payment.Bank, payment.DeliveryCost, payment.GoodsTotal, payment.CustomFee)
	err := row.Scan(&id)
	if err != nil {
		return 0, pkg_errs.Wrap(err, "error while saving payment to db")
	}

	fmt.Println("payment saved successfully")
	return id, nil
}

func (db *Store) GetDeliveryByOrderID(ctx context.Context, orderId uuid.UUID) (*models.Delivery, error) {
	delivery := &models.Delivery{}

	q := `SELECT * FROM delivery WHERE id = (SELECT delivery_id FROM orders WHERE order_id = $1);`

	row := db.QueryRow(ctx, q, orderId)
	err := row.Scan(&delivery.ID, &delivery.Name, &delivery.Phone, &delivery.Zip, &delivery.City, &delivery.Address, &delivery.Region,
		&delivery.Email)

	if err != nil {
		return delivery, pkg_errs.Wrap(err, "error while getting delivery from db")
	}

	return delivery, nil
}

func (db *Store) SaveDelivery(ctx context.Context, delivery models.Delivery) (int, error) {
	fmt.Println("starting saving delivery...")
	q := `INSERT INTO delivery (name, phone, zip, city, address, region, email) VALUES($1, $2, $3, $4, $5, $6, $7)
	RETURNING id`

	row := db.QueryRow(ctx, q, delivery.Name, delivery.Phone, delivery.Zip, delivery.City, delivery.Address, delivery.Region, delivery.Email)
	var id int

	err := row.Scan(&id)
	if err != nil {
		return 0, pkg_errs.Wrap(err, "error while saving delivery to db")
	}

	fmt.Println("delivery saved successfully")

	return id, nil
}

func (db *Store) SaveItems(ctx context.Context, items []models.Item) error {
	fmt.Println("starting saving items...")

	q := `INSERT INTO items (chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	for _, item := range items {
		_, err := db.Exec(ctx, q, item.ChrtId, item.TrackNumber, item.Price, item.RID, item.Name, item.Sale, item.Size, item.TotalPrice, item.NmID, item.Brand, item.Status)
		if err != nil {
			return pkg_errs.Wrap(err, "error while saving items to db")
		}
	}

	fmt.Printf("%d items saved successfully\n", len(items))
	return nil
}

func (db *Store) GetItemByTrackNumber(ctx context.Context, trackNumber string) ([]models.Item, error) {
	items := []models.Item{}

	q := `SELECT * FROM items WHERE track_number = $1`

	rows, err := db.Query(ctx, q, trackNumber)
	if err != nil {
		return items, pkg_errs.Wrap(err, "error while getting items by track number from db")
	}

	for rows.Next() {
		item := models.Item{}
		err := rows.Scan(&item.ChrtId, &item.TrackNumber, &item.Price, &item.RID, &item.Name, &item.Sale, &item.Size, &item.TotalPrice, &item.NmID, &item.Brand, &item.Status)
		if err != nil {
			return items, pkg_errs.Wrap(err, "error while getting items by track number from db")
		}

		items = append(items, item)
	}

	return items, nil
}

// используется для восстановления в случае падения сервиса
func (db *Store) GetAll(ctx context.Context) ([]models.Order, error) {
	orders, err := db.getAllOrders(ctx)
	if err != nil {
		return orders, pkg_errs.Wrap(err, "error while recovering all orders from db")
	}

	for i := 0; i < len(orders); i++ {
		delivery, err := db.GetDeliveryByOrderID(ctx, orders[i].OrderUIID)
		if err != nil {
			return orders, pkg_errs.Wrap(err, "error while recovering all delivery from db")
		}
		fmt.Println("db delivery: ", delivery)
		orders[i].Delivery = *delivery

		items, err := db.GetItemByTrackNumber(ctx, orders[i].TrackNumber)
		if err != nil {
			return orders, pkg_errs.Wrap(err, "error while recovering all items from db")
		}
		orders[i].Items = items

		payment, err := db.GetPaymentByOrderID(ctx, orders[i].OrderUIID)
		if err != nil {
			return orders, pkg_errs.Wrap(err, "error while recovering all payments from db")
		}
		orders[i].Payment = *payment
	}

	fmt.Println("recovered orders 0 delivery = ", orders[0].Delivery)

	return orders, nil
}

func (db *Store) getAllOrders(ctx context.Context) ([]models.Order, error) {
	orders := []models.Order{}

	q := `SELECT * FROM orders`

	rows, err := db.Query(ctx, q)
	if err != nil {
		return orders, pkg_errs.Wrap(err, "error while getting all orders from db")
	}

	for rows.Next() {
		order := models.Order{}
		err := rows.Scan(&order.OrderUIID, &order.TrackNumber, &order.Entry, &order.Delivery.ID, &order.Payment.ID,
			&order.Locale, &order.InternalSignature, &order.CustomerID, &order.DeliveryService, &order.ShardKey,
			&order.SmID, &order.DateCreated, &order.OofShard)
		if err != nil {
			return orders, pkg_errs.Wrap(err, "error while scanning all orders from db")
		}

		orders = append(orders, order)
	}

	return orders, nil
}
