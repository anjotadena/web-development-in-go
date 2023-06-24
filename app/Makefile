BINARY_NAME=centauri
# BINARY_NAME=centauri.exe


build:
	@go mod vendor
	@echo "Building Centauri..."
	@go build -o tmp/${BINARY_NAME} .
	@echo "Centauri started!"

run: build
	@echo "Starting Centauri..."
	@./tmp/${BINARY_NAME} &
	@echo "Centauri started!"

clean:
	@echo "Cleaning..."
	@go clean
	@rm tmp/${BINARY_NAME}
	@echo "Cleaned!"

test:
	@echo "Testing..."
	@go test ./...
	@echo "Done!"

start: run

stop:
	@echo "Stopping Centauri"
	@-pkill -SIGTERM -f "./tmp/${BINARY_NAME}"
	@echo "Stopped Centauri!"

restart: stop start

