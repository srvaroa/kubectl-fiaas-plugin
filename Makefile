VERSION=0.0.1
# MAKEFLAGS += --silent

PREFIX=kubectl-fiaas-
OBJS:=$(shell ls cmd)

BIN_DIR:=bin
RELEASE_DIR:=release

.PHONY: all
all: $(OBJS)
	echo "Plugins generated at $(BIN_DIR)"

.PHONY: clean
clean:
	[ -d $(BIN_DIR) ] && rm -rf $(BIN_DIR)
	[ -d $(RELEASE_DIR) ] && rm -rf $(RELEASE_DIR)

.PHONY: release
release:
	mkdir -p $(RELEASE_DIR)
	for O in $(OBJS) ; do \
		GOOS=linux GOARCH=amd64 go build -o $(RELEASE_DIR)/$(PREFIX)$$O-$(VERSION)-linux-amd64 cmd/$$O/*; \
		GOOS=darwin GOARCH=amd64 go build -o $(RELEASE_DIR)/$(PREFIX)$$O-$(VERSION)-darwin-amd64 cmd/$$O/*; \
		GOOS=windows GOARCH=amd64 go build -o $(RELEASE_DIR)/$(PREFIX)$$O-$(VERSION)-windows-amd64.exe cmd/$$O/*; \
	done

%:
	go build -o $(BIN_DIR)/$(PREFIX)$@ cmd/$@/*
