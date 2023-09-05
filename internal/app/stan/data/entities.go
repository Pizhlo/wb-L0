package data

import (
	"fmt"

	model "github.com/Pizhlo/wb-L0/internal/model"
)

func RandomOrder() model.Order {
	entryEnum := []string{"WBIL"}

	trackNumber := randomString(10)

	items := randomItemsArr(randomInt(1, 5), trackNumber)

	var a *string

	order := model.Order{
		OrderUIID:         randomUIID(),
		TrackNumber:       trackNumber,
		Entry:             randomChoise(entryEnum),
		Delivery:          randomDelivery(),
		Payment:           randomPayment(),
		Items:             items,
		Locale:            "en",
		InternalSignature: a,
		CustomerID:        randomString(3),
		DeliveryService:   "meest",
		ShardKey:          randomString(1),
		SmID:              randomInt(1, 100),
		DateCreated:       randomTime(),
		OofShard:          randomString(1),
	}

	fmt.Println("generated random order id:", order.OrderUIID)

	return order
}

func randomDelivery() model.Delivery {
	delivery := model.Delivery{
		ID:      randomInt(1, 20),
		Name:    randomString(10),
		Phone:   randomPhone(),
		Zip:     randomString(6),
		City:    randomString(6),
		Address: randomString(10),
		Region:  randomString(11),
		Email:   randomEmail(5),
	}

	return delivery
}

func randomItem(trackNumber string) model.Item {
	item := model.Item{
		ChrtId:      randomInt(100, 10000),
		TrackNumber: trackNumber,
		Price:       randomInt(1000, 5000),
		RID:         randomUIID(),
		Name:        randomString(10),
		Sale:        randomInt(0, 51),
		Size:        randomString(3),
		TotalPrice:  randomInt(250, 1600),
		NmID:        randomInt(100000, 1500000),
		Brand:       randomString(10),
		Status:      randomInt(100, 200),
	}

	return item
}

func randomItemsArr(n int, trackNumber string) []model.Item {
	items := []model.Item{}

	for i := 0; i < n; i++ {
		items = append(items, randomItem(trackNumber))
	}

	return items
}

func randomPayment() model.Payment {
	currencyEnum := []string{"USD", "RUB"}
	providerEnum := []string{"wbpay"}

	var a *string
	payment := model.Payment{
		ID:           randomInt(1, 20),
		Transaction:  randomUIID(),
		RequestID:    a,
		Currency:     randomChoise(currencyEnum),
		Provider:     randomChoise(providerEnum),
		Amount:       randomInt(100, 2500),
		PaymentDate:  randomTimeISO(),
		Bank:         "alpha",
		DeliveryCost: randomInt(700, 1500),
		GoodsTotal:   randomInt(200, 1000),
		CustomFee:    0,
	}

	return payment
}
