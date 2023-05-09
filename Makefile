GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BINARY_NAME=simple-http-server

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

run:
	$(GOBUILD) -o $(BINARY_NAME) -v
	./$(BINARY_NAME)

test:
	$(GOTEST) -v ./...

