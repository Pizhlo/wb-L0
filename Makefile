DB_URL=postgresql://root:secret@localhost:8081/wb_db?sslmode=disable

migrateup:
	migrate -path migrations -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path migrations -database "$(DB_URL)" -verbose down

start:
	cd cmd/app
	go run .

infra_up: 
	docker compose up

infra_down:
	docker compose stop

lint:
	golangci-lint run ./...

mock: 
	cd internal/app/storage/postgres
	mockgen -source=store.go -destination=mocks/mocks.go

.PHONY: migrateup migratedown start infra_up infra_down lint mock