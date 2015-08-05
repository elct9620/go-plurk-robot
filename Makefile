PKG=github.com/elct9620/go-plurk-robot
MAIN_PKG=$(PKG)/plurk
CMD_PKG=$(PKG)/cmd/go-plurk-robot
ROBOT_PKG=$(PKG)/robot
DB_PKG=$(PKG)/db
PKGS=$(MAIN_PKG)

all: test

build:
	go build $(CMD_PKG)

test:
	go test ./...

# Merge profile
coverage:
	go test -covermode count -coverprofile go-plurk-robot.cov -coverpkg $(PKGS) $(MAIN_PKG)
	go test -covermode count -coverprofile profile.tmp $(CMD_PKG) && cat profile.tmp | tail -n +2 >> go-plurk-robot.cov
	go test -covermode count -coverprofile profile.tmp $(ROBOT_PKG) && cat profile.tmp | tail -n +2 >> go-plurk-robot.cov
	go test -covermode count -coverprofile profile.tmp $(DB_PKG) && cat profile.tmp | tail -n +2 >> go-plurk-robot.cov
	rm -f profile.tmp

cov-annoate: coverage
	gocov convert go-plurk-robot.cov | gocov annotate -

report: coverage
	gocov convert go-plurk-robot.cov | gocov-html > report.html

clean:
	rm go-plurk-robot.cov
	rm report.html
