
.PHONy:confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

.PHONY: build/app
build/app:
	@go build -ldflags='-s' -C cmd/web -o ../../bin/app

.PHONY: run/app
run/app: build
	@./bin/app

.PHONY: live/server
live/server:
	air

.PHONY: live/tailwind
live/tailwind:
	tailwind -i ui/css/input.css -o ui/public/styles.css --watch


## Quality control
.PHONY: audit
audit: vendor
	@echo "Formatting code..."
	go fmt ./..
	@echo "Vetting code..."
	go vet ./..
	staticcheck ./..
	@echo "Running tests..."
	go test -race -vet=off ./..

.PHONY: vendor
vendor:
	@echo "Tidying and verifying module dependecies..."
	go mod tidy
	go mod verify
	@echo "Vendoring dependecies..."
	go mod vendor


