up:
	docker-compose up -d --build
down:
	docker-compose down
logs:
	docker-compose logs -f app

install_wire:
	@type wire > /dev/null 2>&1 || go get github.com/google/wire/cmd/wire
wire: install_wire
	cd ./cmd/http/di && wire

gotest: wire
	docker-compose exec app go test ./tests/... -p=1 -count=1
