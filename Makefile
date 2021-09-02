GO_PROJECT_NAME := Redirect API

go_prep_build:
	@echo "\n....Preparing installation environment for $(GO_PROJECT_NAME)...."
	go get github.com/cespare/reflex

go_dep_install:
	@echo "\n....Installing dependencies for $(GO_PROJECT_NAME)...."
	go mod download

go_build:
	@echo "\n....Building $(GO_PROJECT_NAME)...."
	go build -o bin/server server.go

go_run:
	@echo "\n....Running $(GO_PROJECT_NAME)...."
	./bin/server

go_test:
	go test -g=server.go -race -covermode=atomic -coverprofile=coverage.out

# Project rules
install:
	$(MAKE) go_prep_build
	$(MAKE) go_dep_install
	$(MAKE) go_build

run:
	reflex --start-service -r '(\.go$$|\.tmpl$$)' make restart

restart:
	@$(MAKE) go_test
	@$(MAKE) go_build
	@$(MAKE) go_run

.PHONY: go_prep_build go_dep_install go_build go_run install run restart reflex

api-doc:
	docker-compose exec app /bin/sh -c "swag init -g=src/router/swagger.go #--parseDependency --parseInternal"

test:
	docker-compose exec app /bin/sh -c "go test -covermode=count -coverprofile=coverage.out ./src/..."

migrate:
	docker-compose exec app /bin/sh -c "go run ./src/console/migrate.go"