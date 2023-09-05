package cache

import (
	"fmt"
	"sync"

	"github.com/Pizhlo/wb-L0/internal/app/errs"
	models "github.com/Pizhlo/wb-L0/internal/model"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type OrderCacher interface {
	Get(id uuid.UUID) (any, error)
	Save(id uuid.UUID, order models.Order)
	SaveAll(orders []models.Order)
}

type OrderCache struct {
	data sync.Map
}

func NewOrder() *OrderCache {
	return &OrderCache{}
}

func (c *OrderCache) Get(id uuid.UUID) (any, error) {
	fmt.Println("getting order from cache id = ", id)
	val, ok := c.data.Load(id)
	if !ok {
		fmt.Println("not found in cache")
		return nil, errors.Wrap(errs.NotFound, "not found order in cache")
	}
	fmt.Println("found in cache")
	return val, nil
}

func (c *OrderCache) Save(id uuid.UUID, order models.Order) {
	fmt.Println("cache save order:", order.OrderUIID)
	c.data.Store(id, order)
}

// восстановить данные из БД в случае падения сервиса
func (c *OrderCache) SaveAll(orders []models.Order) {
	count := len(orders)
	var wg sync.WaitGroup
	wg.Add(count)

	for i := 0; i < count; i++ {
		go func(j int) {
			defer wg.Done()
			c.Save(orders[j].OrderUIID, orders[j])
		}(i)
	}

	wg.Wait()

	fmt.Printf("totally recovered %d orders\n", count)
}
