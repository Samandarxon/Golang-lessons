

migration-down:
	migrate -path migrations/postgres/ -database "postgresql://asadbek:3066586@localhost:5432/travel?sslmode=disable" -verbose down

migration-up:
	migrate -path migrations/postgres/ -database "postgresql://asadbek:3066586@localhost:5432/travel?sslmode=disable" -verbose up

gen-swag:
	swag init -g ./api/api.go -o ./api/docs

