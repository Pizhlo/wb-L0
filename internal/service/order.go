package service

import (
	"context"
	"errors"

	"github.com/Pizhlo/wb-L0/internal/app/errs"
	models "github.com/Pizhlo/wb-L0/internal/model"
	"github.com/google/uuid"
)

func (s *Service) GetOrderByID(ctx context.Context, id uuid.UUID) (*models.Order, error) {
	var order *models.Order
	var err error

	order, err = s.GetOrderByIDFromCache(id)
	if err == nil { // if everything is good and order is in cache
		return order, nil
	}

	if !errors.Is(err, errs.NilOrder) {
		return order, err
	}

	// else if there's no order in cache
	order, err = s.Repo.GetOrderByID(ctx, id)
	if err != nil {
		return order, err
	}

	return order, nil
}

func (s *Service) GetOrderByIDFromCache(id uuid.UUID) (*models.Order, error) {
	order := models.Order{}

	// get order
	orderAny, err := s.Cache.Order.Get(id)
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

func (s *Service) SaveOrder(ctx context.Context, order models.Order) error {
	s.saveOrderInCache(order)
	err := s.saveOrderInDB(ctx, order)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) saveOrderInCache(order models.Order) {
	s.Cache.Order.Save(order.OrderUIID, order)
}

func (s *Service) saveOrderInDB(ctx context.Context, order models.Order) error {
	err := s.Repo.SaveOrder(ctx, order)
	if err != nil {
		return err
	}

	return nil
}
