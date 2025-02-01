.PHONY: create_migration migrateup migratedown run

create_migration:
	migrate create -ext sql -dir database/migrations $(filter-out $@,$(MAKECMDGOALS));

migrate-up:
	go run database/migrate.go --envFile .env

migrate-down:
	go run database/migrate.go --envFile .env --down=true

run:
	go run cmd/server/* --envFile=.env