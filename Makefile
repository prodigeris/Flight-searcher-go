install:
	test -f .env || cp .env.example .env
test:
	go test -v ./...