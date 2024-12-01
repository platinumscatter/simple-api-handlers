DB_DSN := "postgres://postgres:yourpassword@localhost:5432/?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

gen:
	oapi-codegen -config openapi/.openapi -include-tags tasks -package tasks openapi/openapi.yaml > ./internal/web/tasks/api.gen.go

gen-user:
	oapi-codegen -config openapi/.openapi -include-tags users -package users openapi/openapi.yaml > ./internal/web/users/api.gen.go


migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}

migrate:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down

lint:
	golangci-lint run --out-format=colored-line-number
	
run:
	go run cmd/app/main.go 