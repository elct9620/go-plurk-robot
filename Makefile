PKG=github.com/elct9620/go-plurk-robot
MAIN_PKG=$(PKG)/plurk
CMD_PKG=$(PKG)/cmd/go-plurk-robot

all: test

build:
	go build $(CMD_PKG)

test:
	go test $(MAIN_PKG)

coverage:
	go test -cover -coverprofile go-plurk-robot.cov $(MAIN_PKG)
