
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
run/app: build
	@./bin/app

# ==================================================================================== #
# BUILD
# ==================================================================================== #

# go build -ldflags='-s' -C cmd/web -o ../../bin/app
.PHONY: build/app
build/app:
	@echo 'Building cmd/app...'
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

production_host_ip = '10.1.1.23'

## production/connect: connect to the production server
.PHONY: production/connect
production/connect:
	ssh torkeldev@${production_host_ip}

## production/deploy/api: deploy the api to production
.PHONY: production/deploy/app
production/deploy/app:
	rsync -P ./bin/linux_amd64/torkeldev-app torkeldev@${production_host_ip}:~