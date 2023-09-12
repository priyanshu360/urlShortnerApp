.PHONY: all build run stop clean

# Default target
all: build run

# Build the containers
build:
	docker-compose build

# Run the containers in the background
run:
	docker-compose up -d

# Stop and remove the containers
stop:
	docker-compose down

# Clean up all stopped containers and volumes
clean:
	docker-compose down -v
