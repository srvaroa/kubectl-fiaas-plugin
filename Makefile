VERSION=0.1
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
		GO111MODULE=on GOOS=linux GOARCH=amd64 go build -o $(RELEASE_DIR)/$(PREFIX)$$O-$(VERSION)-linux-amd64 cmd/$$O/*; \
		GO111MODULE=on GOOS=darwin GOARCH=amd64 go build -o $(RELEASE_DIR)/$(PREFIX)$$O-$(VERSION)-darwin-amd64 cmd/$$O/*; \
		GO111MODULE=on GOOS=windows GOARCH=amd64 go build -o $(RELEASE_DIR)/$(PREFIX)$$O-$(VERSION)-windows-amd64.exe cmd/$$O/*; \
	done

%:
	GO111MODULE=on go build -o $(BIN_DIR)/$(PREFIX)$@ cmd/$@/*
