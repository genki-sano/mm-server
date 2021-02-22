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

install_sops:
	@type sops > /dev/null 2>&1 || go get go.mozilla.org/sops/v3/cmd/sops
sops: install_sops
	sops -e --gcp-kms projects/money-maneger-304709/locations/global/keyRings/sops/cryptoKeys/sops-key secret.yaml > secret.enc.yaml
