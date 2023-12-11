include .env 

build: 
	if [ -f "${BINARY_FILE_NAME}"]; then \
		rm ${BINARY_FILE_NAME}; \
		echo "deleted ${BINARY_FILE_NAME}"; \
	fi 
	@echo "Building binary..."
	go build -o ${BINARY_FILE_NAME} cmd/blog/*.go

run: build
	./${BINARY_FILE_NAME}

docker_container:
	docker-compose up -d
