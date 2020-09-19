.DEFAULT_GOAL := generate-proto

 ## build: Generate proto file.
generate-proto:
	@echo "Building Proto file..."
    $(shell ./proto.sh $<)

## build: Build Dockerfile.
build-service-user:
	@echo "Building Go Binary..."
	docker build --no-cache -f Dockerfile -t user .
build-service-product:
	@echo "Building Go Binary..."
	docker build --no-cache -f Dockerfile -t product .

## start-server: Start in development mode. Gets reloaded automatically when code changes.
start-server:
	@echo "Running Server..."
	docker-compose up -d

## stop-server: Stop development mode.
stop-server:
	@echo "Stopping Server..."
	docker-compose down

## watch-logs: Display logs in the console.
watch-logs:
	docker-compose logs -f

## clean: Remove all unused locale Volumes and remove all stopped containers.
clean:
	docker system prune -f
	docker volume prune -f

help: Makefile
	@echo
	@echo "Available commands:"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo