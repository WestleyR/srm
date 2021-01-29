# Created by WestleyR on July 28, 2020
# Source code: https://github.com/WestleyR/srm
# Last modified data: 2021-01-28
#
# This file is licensed under the terms of
#
# The Clear BSD License
#
# Copyright (c) 2020-2021 WestleyR
# All rights reserved.
#

PREFIX = /usr/local

TARGET = srm
TARGET_VERSION = 2.0.1

GO = go
GOFLAGS = -ldflags -w

SRC = $(wildcard ./*.go)

.PHONY:
all: $(TARGET)

.PHONY:
$(TARGET): $(SRC)
	$(GO) build $(GOFLAGS) ./cmd/srm
	
test: $(TARGET)
	@bash ./run-tests

install: $(TARGET)
	mkdir -p $(PREFIX)/bin
	cp -f $(TARGET) $(PREFIX)/bin

package: $(SRC)
	TARGET_VERSION=$(TARGET_VERSION) ./package.sh

clean:
	rm -f $(TARGET)
	rm -rf binaries

uninstall: $(PREFIX)/$(TARGET)
	rm -f $(PREFIX)/$(TARGET)

