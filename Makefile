include .env

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

.PHONY: live/server
live/server:
	air

.PHONY: live/tailwind
live/tailwind:
	tailwindcss -i ui/css/input.css -o ui/public/css/styles.css --watch

.PHONY: run/app
run/app:
	@./bin/app -env='development' -port='4001'

# ==================================================================================== #
# BUILD
# ==================================================================================== #

# go build -ldflags='-s' -C cmd/web -o ../../bin/app
.PHONY: build/app
build/app:
	@echo 'Building cmd/app...'
	tailwindcss -i ui/css/input.css -o ui/public/css/styles.css --minify
	go build -ldflags='-s' -o=./bin/app ./cmd/web
	GOOS=linux GOARCH=amd64 go build -ldflags='-s' -o=./bin/linux_amd64/torkeldev-app ./cmd/web


## Quality control
.PHONY: audit
audit: vendor
	@echo "Formatting code..."
	go fmt ./...
	@echo "Vetting code..."
	go vet ./...
	staticcheck ./...
	@echo "Running tests..."
	go test -race -vet=off ./...

.PHONY: vendor
vendor:
	@echo "Tidying and verifying module dependecies..."
	go mod tidy
	go mod verify
	@echo "Vendoring dependecies..."
	go mod vendor

# ==================================================================================== #
# PRODUCTION
# ==================================================================================== #

## production/connect: connect to the production server
.PHONY: production/connect
production/connect:
	ssh torkeldev@${PRODUCTION_HOST_IP}

## production/deploy/api: deploy the api to production
.PHONY: production/deploy/app
production/deploy/app:
	rsync -P ./bin/linux_amd64/torkeldev-app torkeldev@${PRODUCTION_HOST_IP}:~
	rsync -P ./remote/production/torkeldev.service torkeldev@${PRODUCTION_HOST_IP}:~
	rsync -P ./remote/production/Caddyfile torkeldev@${PRODUCTION_HOST_IP}:~
	ssh -t torkeldev@${PRODUCTION_HOST_IP} '\
		sudo mv ~/torkeldev.service /etc/systemd/system/ \
		&& sudo systemctl enable torkeldev \
		&& sudo systemctl restart torkeldev \
		&& sudo mv ~/Caddyfile /etc/caddy/ \
		&& sudo systemctl reload caddy \
	'
	@echo "deployment complete..."
