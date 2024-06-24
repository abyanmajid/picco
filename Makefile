BROKER_BINARY=broker
USER_BINARY=user
COMPILER_BINARY=compiler
JUDGE_BINARY=judge

# up: starts all containers in the background without forcing build
up:
	@echo "[codemore.io] Starting Docker images..."
	docker-compose up -d
	@echo "[codemore.io] Containers has successfully been started!"

# down: stopping docker containers
down:
	@echo "[codemore.io] Stopping containers..."
	docker-compose down
	@echo "[codemore.io] Containers has successfully been stopped!"

# build: stops docker-compose (if running), builds all projects and starts docker compose
build: build-broker build-user build-compiler build-judge
	@echo "[codemore.io] Stopping docker images (if running...)"
	docker-compose down
	@echo "[codemore.io] Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "[codemore.io] Docker images built and started!"

# build-broker: build linux executable for broker service
build-broker:
	@echo "[codemore.io] Building broker..."
	cd ./broker && env GOOS=linux CGO_ENABLED=0 go build -o ${BROKER_BINARY} ./cmd/api
	@echo "[codemore.io] Broker has successfully been built!"

# build-user: build linux executable for user service
build-user:
	@echo "[codemore.io] Building user..."
	cd ./user && env GOOS=linux CGO_ENABLED=0 go build -o ${USER_BINARY} ./cmd/api
	@echo "[codemore.io] User has successfully been built!"

# build-compiler: build linux executable for compiler service
build-compiler:
	@echo "[codemore.io] Building compiler..."
	cd ./compiler && env GOOS=linux CGO_ENABLED=0 go build -o ${COMPILER_BINARY} ./cmd/api
	@echo "[codemore.io] Compiler has successfully been built!"

# build-judge: build linux executable for judge service
build-judge:
	@echo "[codemore.io] Building judge..."
	cd ./judge && env GOOS=linux CGO_ENABLED=0 go build -o ${JUDGE_BINARY} ./cmd/api
	@echo "[codemore.io] Judge has successfully been built!"

# users-migrate-up: run goose migrate up for users database
users-migrate-up:
	@echo "[codemore.io] Running goose up migration on users database..."
	goose -dir ./user/sql/migrations postgres postgresql://postgres:postgres@localhost:5432/users up
	@echo "[codemore.io] Successfully ran goose up migration!"

# users-migrate-down: run goose migrate down for users database
users-migrate-down:
	@echo "[codemore.io] Running goose down migration on users database..."
	goose -dir ./user/sql/migrations postgres postgresql://postgres:postgres@localhost:5432/users down
	@echo "[codemore.io] Successfully ran goose down migration!"

# tv-user: vendor dependencies in user microservice
tv-user:
	@echo "[codemore.io] Vendoring dependencies for user microservice..."
	cd user && go mod tidy && go mod vendor && cd ..
	@echo "[codemore.io] Successfully vendored dependencies for user microservice..."

# ui: start client
ui:
	@echo "[codemore.io] Starting client..."
	cd ./web/ui && npm run dev

docs:
	@echo "[codemore.io] Starting docs client..."
	cd ./web/docs && npm run dev