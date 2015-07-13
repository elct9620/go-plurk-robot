PKG=github.com/elct9620/go-plurk-robot
MAIN_PKG=$(PKG)/plurk
CMD_PKG=$(PKG)/cmd/go-plurk-robot
PKGS=$(MAIN_PKG),$(PKG)/logger

all: test

build:
	go build $(CMD_PKG)

test:
	go test ./...

coverage:
	go test -covermode count -coverprofile go-plurk-robot.cov -coverpkg $(PKGS) $(MAIN_PKG)
