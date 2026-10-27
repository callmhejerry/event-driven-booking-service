include .env

.PHONY: test-config


test-config:
	@echo "The database url is $(DB_URL)"

migrate-create:
	@if [ -z "${name} " ]; then \
			echo "❌ Error: name argument is required."\
			echo "👉🏽 Usuage: make migrate-create name=some_name" \
			exit 1;\
	fi

	@migrate create -dir sql/migrations -ext sql -seq ${name}

migrate-up:
	@migrate -path ./sql/migrations -database ${DB_URL} up

migrate-down:
	@migrate -path ./sql/migrations -database ${DB_URL} down

migrate-version:
	@migrate -path ./sql/migrations -database ${DB_URL} version

migrate-force:
	@migrate -path ./sql/migrations -database ${DB_URL} force ${force}