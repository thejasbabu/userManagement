PACKAGE  = github.com/thejasbabu/userManagement/
BASE     = $(GOPATH)/src/$(PACKAGE)

.PHONY: build, clean
build: | $(BASE)
	cd $(BASE) && go build main.go

linux-build:
	cd $(BASE) && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go

clean:
	@echo "Cleaning"
	go clean
