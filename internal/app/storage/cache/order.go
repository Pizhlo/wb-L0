package cache

import (
	"fmt"
	"sync"

	"github.com/Pizhlo/wb-L0/internal/app/errs"
	models "github.com/Pizhlo/wb-L0/internal/model"
	"github.com/google/uuid"
)

type OrderCacher interface {
	Get(id uuid.UUID) (any, error)
	Save(id uuid.UUID, order models.Order)
}

type OrderCache struct {
	sync.Map
}

func NewOrder() *OrderCache {
	return &OrderCache{}
}

func (c *OrderCache) Get(id uuid.UUID) (any, error) {
	val, ok := c.Load(id)
	if !ok {
		return nil, errs.NotFound
	}
	return val, nil
}

func (c *OrderCache) Save(id uuid.UUID, order models.Order) {
	fmt.Println("cache save order: ", order)
	c.Store(id, order)
}
