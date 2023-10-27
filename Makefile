BINARY_NAME=monte-carlo-simulation

build:
	@echo "building binary"
	@go build -o bin/${BINARY_NAME}
	@echo "Finished"

run:
	@echo "running binary"
	@./bin/${BINARY_NAME}

clean:
	@echo "cleaning binaries"
	@go clean
	@rm bin/${BINARY_NAME}
	@echo "binaries removed"

clean-logs:
	@echo "removing logs"
	@rm data/logs/*.log
	@echo "logs removed"

go-get:
	@go get
