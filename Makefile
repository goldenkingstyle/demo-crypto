BINARY_NAME=crypto.exe

run: compile
	@./build/${BINARY_NAME}

compile: 
	@go build -o build/${BINARY_NAME} cmd/main.go