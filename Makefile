DB_NAME=postgres
DB_URL=postgresql://postgres:postgres@localhost:5432/${DB_NAME}?sslmode=disable

swagger-generate:
	swag init -g api/server.go -o api/docs

migrateup:
	migrate -database ${DB_URL} -path migrations -verbose up

migratedown:
	migrate -database ${DB_URL} -path migrations -verbose down

test:
	go test -v -cover ./...

mock:
	mockgen -package mockdb -destination storage/mock/db_repository.go github.com/TemurMannonov/query_analyzer/api DBRepositoryI

local:
	docker compose up -d

local-down:
	docker compose down

.PHONY: swagger-generate migrateup migratedown mock test local local-down
