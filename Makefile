PKG=github.com/elct9620/go-plurk-robot
MAIN_PKG=$(PKG)/plurk
CMD_PKG=$(PKG)/cmd/go-plurk-robot
PKGS=$(MAIN_PKG)

all: test

build:
	go build $(CMD_PKG)

test:
	go test ./...

coverage:
	go test -covermode count -coverprofile go-plurk-robot.cov -coverpkg $(PKGS) $(MAIN_PKG)

cov-annoate: coverage
	gocov convert go-plurk-robot.cov | gocov annotate -

report: coverage
	gocov convert go-plurk-robot.cov | gocov-html > report.html

clean:
	rm go-plurk-robot.cov
	rm report.html
