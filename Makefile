api-doc:
	swag init -g=src/router/swagger.go #--parseDependency --parseInternal

migrate:
	go run ./src/console/migrate.go