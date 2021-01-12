# Created by: WestleyR
# Email: westleyr@nym.hush.com
# Url: https://github.com/WestleyR/srm
# Last modified date: 2021-01-11
#
# This file is licensed under the terms of
#
# The Clear BSD License
#
# Copyright (c) 2020-2021 WestleyR
# All rights reserved.
#
# This software is licensed under a Clear BSD License.
#

PREFIX = /usr/local

TARGET = srm
TARGET_VERSION = 2.0.0.a1

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

