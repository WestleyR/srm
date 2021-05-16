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

GO = go
GOFLAGS = -ldflags -w

SRC = $(shell find $(SOURCEDIR) ./ -name '*.go')

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

