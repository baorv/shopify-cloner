GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
BINARY_NAME=bin/shopify
MAIN_GO=cmd/shopify/main.go

build:
	$(GOBUILD) -o $(BINARY_NAME) $(MAIN_GO)
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)