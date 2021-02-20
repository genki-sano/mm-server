up:
	docker-compose build && docker-compose up -d
down:
	docker-compose down
logs:
	docker-compose logs -f app
test:
	docker-compose exec app go test ./tests/...

install_wire:
	@type wire > /dev/null 2>&1 || go get github.com/google/wire/cmd/wire
wire: install_wire
	cd ./cmd/http/di && wire
