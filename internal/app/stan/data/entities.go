package data

import (
	"database/sql"

	"github.com/Pizhlo/wb-L0/models"
)

func RandomOrder() models.Order {
	entryEnum := []string{"WBIL"}

	trackNumber := randomString(10)

	items := randomItemsArr(randomInt(1, 5), trackNumber)

	order := models.Order{
		ID:          randomInt(0, 50),
		OrderUIID:   randomUIID(),
		TrackNumber: trackNumber,
		Entry:       randomChoise(entryEnum),
		Delivery:    randomDelivery(),
		Payment:     randomPayment(),
		Items:       items,
		Locale:      "en",
		InternalSignature: sql.NullString{
			String: "",
			Valid:  false,
		},
		CustomerID:      randomString(3),
		DeliveryService: "meest",
		ShardKey:        randomString(1),
		SmID:            randomInt(1, 100),
		DateCreated:     randomTime(),
		OofShard:        randomString(1),
	}

	return order
}

func randomDelivery() models.Delivery {
	delivery := models.Delivery{
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

func randomItem(trackNumber string) models.Item {
	item := models.Item{
		ID:          randomInt(1, 50),
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

func randomItemsArr(n int, trackNumber string) []models.Item{
	items := []models.Item{}

	for i := 0; i < n; i++ {
		item := randomItem(trackNumber)
		items = append(items, item)
	}

	return items
}

func randomPayment() models.Payment {
	currencyEnum := []string{"USD", "RUB"}
	providerEnum := []string{"wbpay"}

	payment := models.Payment{
		ID:          randomInt(1, 20),
		Transaction: randomUIID(),
		RequestID: sql.NullString{
			Valid: false,
		},
		Currency:     randomChoise(currencyEnum),
		Provider:     randomChoise(providerEnum),
		PaymentDate:  randomTimeISO(),
		Bank:         "alpha",
		DeliveryCost: randomInt(700, 1500),
		GoodsTotal:   randomInt(200, 1000),
		CustomFee:    0,
	}

	return payment
}