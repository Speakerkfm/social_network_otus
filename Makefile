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
	go install go.k6.io/xk6/cmd/xk6@latest

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

load-data:
	PGPASSWORD=$(DB_POSTGRES_PASS) psql -h $(DB_POSTGRES_MASTER_HOST) \
			-p $(DB_POSTGRES_MASTER_PORT) \
			-U $(DB_POSTGRES_USER) \
			-d $(DB_POSTGRES_DATABASE_NAME) \
			-c "CREATE TABLE IF NOT EXISTS tmp_table( \
                   	name      text                    NOT NULL, \
                    age             integer                 NOT NULL, \
                    city            text                     NOT NULL)"
	PGPASSWORD=$(DB_POSTGRES_PASS) psql -h $(DB_POSTGRES_MASTER_HOST) \
    			-p $(DB_POSTGRES_MASTER_PORT) \
    			-U $(DB_POSTGRES_USER) \
    			-d $(DB_POSTGRES_DATABASE_NAME) \
    			-c "\copy tmp_table FROM './load_test/people.csv' WITH(FORMAT csv);"
	PGPASSWORD=$(DB_POSTGRES_PASS) psql -h $(DB_POSTGRES_MASTER_HOST) \
    			-p $(DB_POSTGRES_MASTER_PORT) \
    			-U $(DB_POSTGRES_USER) \
    			-d $(DB_POSTGRES_DATABASE_NAME) \
    			-c "INSERT INTO social_user(id, first_name, second_name, age, city, sex, biography, hashed_password) \
                SELECT gen_random_uuid(), split_part(name, ' ', 1), split_part(name, ' ', 2), age, city, 1, 'biography', 'hashed_password' \
                FROM tmp_table \
                ON CONFLICT DO NOTHING; \
                DROP TABLE tmp_table;"

run:
	go run cmd/social_network_server/main.go