package service

import (
	"context"

	"github.com/google/uuid"
)

func (s *Service) GetOrderByID(ctx context.Context, id uuid.UUID) (model.Order, error) {
	delivery, err := s.Repo.GetDeliveryByOrderID(ctx, id)
	if err != nil {
		return err
	}

	payment, err := s.Repo.GetPaymentByOrderID(ctx, id)
	if err != nil {
		return err
	}

	order, err := s.Repo.GetOrderByID(ctx, id)
	if err != nil {
		return err
	}

	order.Delivery = delivery
	order.Payment = payment

	return nil
}
