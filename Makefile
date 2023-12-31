FRONT_END_BINARY=frontApp
BROKER_BINARY=brokerApp
AUTH_BINARY=authApp
LOGGER_BINARY=loggerApp
MAIL_BINARY=mailApp
LISTENER_BINARY=listenerApp

## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build:
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

## build_broker: builds the broker binary as a linux executable
build_broker:
	@echo "Building broker binary..."
	cd broker-service && env GOOS=linux CGO_ENABLED=0 go build -o ${BROKER_BINARY} ./api
	chmod +x ./broker-service/${BROKER_BINARY}
	@echo "Done!"

build_auth:
	@echo "Building auth binary..."
	cd auth-service && env GOOS=linux CGO_ENABLED=0 go build -o ${AUTH_BINARY} ./api
	chmod +x ./auth-service/${AUTH_BINARY}
	@echo "Done!"

build_logger:
	@echo "Building logger binary..."
	cd logger-service && env GOOS=linux CGO_ENABLED=0 go build -o ${LOGGER_BINARY} ./api
	chmod +x ./logger-service/${LOGGER_BINARY}
	@echo "Done!"

build_mail:
	@echo "Building mail binary..."
	cd mail-service && env GOOS=linux CGO_ENABLED=0 go build -o ${MAIL_BINARY} ./api
	chmod +x ./mail-service/${MAIL_BINARY}
	@echo "Done!"

build_listener:
	@echo "Building listener binary..."
	cd listener-service && env GOOS=linux CGO_ENABLED=0 go build -o ${LISTENER_BINARY} ./api
	chmod +x ./listener-service/${LISTENER_BINARY}
	@echo "Done!"

## build_front: builds the frone end binary
build_front:
	@echo "Building front end binary..."
	cd front-end && env CGO_ENABLED=0 go build -o ${FRONT_END_BINARY} ./web
	chmod +x ./front-end/${FRONT_END_BINARY}
	@echo "Done!"

## start: starts the front end
start: build_front
	@echo "Starting front end"
	cd front-end && sudo ./${FRONT_END_BINARY} &

## stop: stop the front end
stop:
	@echo "Stopping front end..."
	@-pkill -SIGTERM -f "./${FRONT_END_BINARY}"
	@echo "Stopped front end!"