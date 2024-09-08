.PHONY: db-migrations

db-migrations:
	cd db_migrations; \
	go get; \
	go run .;