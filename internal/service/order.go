package service

import (
	"context"
	"errors"

	"github.com/Pizhlo/wb-L0/internal/app/errs"
	"github.com/Pizhlo/wb-L0/internal/app/storage/cache"
	"github.com/Pizhlo/wb-L0/internal/app/storage/postgres"
	models "github.com/Pizhlo/wb-L0/internal/model"
	"github.com/google/uuid"
)

type Order struct {
	Repo       postgres.Repo
	OrderCache cache.OrderCacher
}

func New(repo postgres.Repo, cache cache.OrderCacher) *Order {
	return &Order{
		Repo:       repo,
		OrderCache: cache,
	}
}

func (s *Order) GetOrderByID(ctx context.Context, id uuid.UUID) (*models.Order, error) {
	var order *models.Order
	var err error

	order, err = s.getOrderByIDFromCache(id)
	if err == nil { // if everything is good and order is in cache
		return order, nil
	}

	if !errors.Is(err, errs.NotFound) {
		return order, err
	}

	// else if there's no order in cache
	order, err = s.Repo.GetOrderByID(ctx, id)
	if err != nil {
		return order, err
	}

	return order, nil
}

func (s *Order) getOrderByIDFromCache(id uuid.UUID) (*models.Order, error) {
	order := models.Order{}

	// get order
	orderAny, err := s.OrderCache.Get(id)
	if err != nil {
		return &order, err
	}

	if orderAny == nil {
		return &order, errs.NilOrder
	}

	order, ok := orderAny.(models.Order)
	if !ok {
		return &order, errs.UnableConvertOrder
	}

	return &order, nil

}

func (s *Order) SaveOrder(ctx context.Context, order models.Order) error {
	orderWithIDs, err := s.saveOrderInDB(ctx, order)
	if err != nil {
		return err
	}

	s.saveOrderInCache(*orderWithIDs)

	return nil
}

func (s *Order) saveOrderInCache(order models.Order) {
	s.OrderCache.Save(order.OrderUIID, order)
}

func (s *Order) saveOrderInDB(ctx context.Context, order models.Order) (*models.Order, error) {
	orderWithIDs, err := s.Repo.SaveOrder(ctx, order)
	if err != nil {
		return orderWithIDs, err
	}

	return orderWithIDs, nil
}

func (s *Order) Recover(ctx context.Context) error {
	orders, err := s.Repo.GetAll(ctx)
	if err != nil {
		return err
	}

	s.OrderCache.SaveAll(orders)
	return nil
}

func (s *Order) GetAllOrders(ctx context.Context) ([]models.Order, error) {
	orders, err := s.Repo.GetAll(ctx)
	if err != nil {
		return orders, err
	}

	return orders, nil
}
