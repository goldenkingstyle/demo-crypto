ifeq ($(OS),Windows_NT)
	BINARY_NAME=crypto.exe
else 
	BINARY_NAME=crypto
endif

BUILD_PATH=build/${BINARY_NAME}
ARGS?=""

init: ARGS=init
profile: ARGS=profile
set: ARGS=set
price: ARGS=price
buy: ARGS=buy

init profile set price buy run: compile
	@./${BUILD_PATH} ${ARGS}


compile: 
	@go build -o ${BUILD_PATH} cmd/main.go