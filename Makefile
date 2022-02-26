include .env
export $(shell sed 's/=.*//' .env)

LDFLAGS="-w -s -X github.com/hinupurthakur/grep/cli/constants.version=${CTL_VERSION}"
make:
	go build -ldflags=${LDFLAGS} -o grep cmd/main.go