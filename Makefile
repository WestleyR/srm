#
#  Makefile
#  srm - https://github.com/WestleyR/srm
#
# Created by WestleyR on July 28, 2020
# Source code: https://github.com/WestleyR/srm
#
# Copyright (c) 2020-2021 WestleyR. All rights reserved.
# This software is licensed under a BSD 3-Clause Clear License.
# Consult the LICENSE file that came with this software regarding
# your rights to distribute this software.
#

PREFIX = /usr/local

TARGET = srm

VERSION = $(shell git describe --tags --always)

GO = go
GOFLAGS = -ldflags "-w -X github.com/WestleyR/srm/cmd/srm.Version=$(VERSION)"

SRC = $(shell find $(SOURCEDIR) ./ -name '*.go')

.PHONY:
all: $(TARGET)

.PHONY:
$(TARGET): $(SRC)
	$(GO) build $(GOFLAGS)

test: $(TARGET)
	@bash ./run-tests

install: $(TARGET)
	mkdir -p $(PREFIX)/bin
	cp -f $(TARGET) $(PREFIX)/bin

package: $(TARGET)
	TARGET_VERSION=$(shell ./srm --version) ./package.sh

clean:
	rm -f $(TARGET)
	rm -rf binaries

uninstall: $(PREFIX)/$(TARGET)
	rm -f $(PREFIX)/$(TARGET)

