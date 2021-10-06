api-doc:
	docker-compose exec app /bin/sh -c "swag init -g=src/router/swagger.go #--parseDependency --parseInternal"

test-coverage:
	docker-compose exec app /bin/sh -c "go test -covermode=count -coverprofile=coverage.out ./src/..."

test:
	docker-compose exec app /bin/sh -c "go test ./src/..."

migrate:
	docker-compose exec app /bin/sh -c "go run ./src/console/migrate.go"

dev:
	docker-compose -f docker-compose-dev.yml up