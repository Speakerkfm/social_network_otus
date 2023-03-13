DB_POSTGRES_DATABASE_NAME?=social_network_otus
DB_POSTGRES_HOST?=localhost:5432
DB_POSTGRES_MASTER_HOST?=localhost
DB_POSTGRES_MASTER_PORT?=5432
DB_POSTGRES_USER?=social-network-user
DB_POSTGRES_PASS?=social-network-password

install:
	go mod tidy
	go install github.com/rakyll/statik@latest
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
	go install github.com/pressly/goose/v3/cmd/goose@latest

generate:
	cp ./api/swagger.json ./swaggerui/swagger.json
	statik -src=./swaggerui --dest=./internal/app/admin
	oapi-codegen --config=oapi-config.server.yaml ./api/swagger.json

start-postgres:
	@echo "recreating postgres container..."
	docker stop social_network_pg || true
	docker rm social_network_pg || true
	docker run --name social_network_pg \
    		-e POSTGRES_USER=$(DB_POSTGRES_USER) \
     		-e POSTGRES_PASSWORD=$(DB_POSTGRES_PASS) \
     		-e POSTGRES_DB=$(DB_POSTGRES_DATABASE_NAME) -p 5432:5432 -d postgres
	@echo "postgres container recreated"
	echo "starting postgres container..."
	sleep 5

migration:
	goose -dir migrations postgres \
				"host=$(DB_POSTGRES_MASTER_HOST) \
				port=$(DB_POSTGRES_MASTER_PORT) \
				dbname=$(DB_POSTGRES_DATABASE_NAME) \
				user=$(DB_POSTGRES_USER) \
				password=$(DB_POSTGRES_PASS) \
				sslmode=disable" up

run:
	go run cmd/social_network_server/main.go