install:
	test -f .env || cp .env.example .env
	test -f database/db.db || cp database/db_start.db database/db.db
test:
	go test -v ./...