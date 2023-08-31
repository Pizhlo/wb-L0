package cache

type Cache struct {
	Order OrderCacher
}

func New(order OrderCacher) *Cache {
	return &Cache{
		Order: order,
	}
}
