# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
CMD_PATH=./cmd/shipping/
BINARY_PATH=./build/
BINARY_NAME=shipping
BINARY_UNIX=$(BINARY_NAME)_unix

all: test build-linux

test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run:
	@$(GORUN) $(CMD_PATH)*.go
deps:
	$(GOGET) -u ./...


# Cross compilation
build-linux:
	CGO_ENABLED=0 \
		go build \
			-installsuffix cgo \
			-o $(BINARY_PATH)$(BINARY_NAME) \
			-ldflags "-X main.Version=$(APP_VERSION)" \
			$(CMD_PATH)*.go

