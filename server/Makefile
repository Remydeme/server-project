CMD="main.go"
PATH_MAIN="./main.go"
CONFIG_ENV_SCRIPT_PATH="config/configure.ENV.sh"
CONFIG_FILE="./config/config.ENV.json"
CONFIG_TEST_FILE="config/config.TEST.json"
PATH_TO_DIR=$(pwd)

all: test build 


build:
	CONFIG_PATH=${CONFIG_FILE} go run ${PATH_MAIN}

test:
	CONFIG_PATH=${PATH_TO_DIR}${CONFIG_TEST_FILE}; go test ./...

