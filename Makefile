# Created by: WestleyR
# Email: westleyr@nym.hush.com
# Url: https://github.com/WestleyR/srm
# Last modified date: 2020-07-29
#
# This file is licensed under the terms of
#
# The Clear BSD License
#
# Copyright (c) 2020 WestleyR
# All rights reserved.
#
# This software is licensed under a Clear BSD License.
#

GO = go
TARGET = srm
TARGET_VERSION = 2.0.0.a1

PREFIX = /usr/local

MODDED = $(shell if command -v git > /dev/null ; then (git diff --exit-code --quiet && echo \"[No changes]\") || echo \"[With uncommited changes]\" ; else echo \"[unknown]\" ; fi)

COMMIT = "$(shell git log -1 --oneline --decorate=short --no-color || ( echo 'ERROR: unable to get commit hash' >&2 ; echo unknown ) )"

CFLAGS += -DCOMMIT_HASH=\"$(COMMIT)\"
CFLAGS += -DUNCOMMITED_CHANGES=\"$(MODDED)\"

ifeq ($(DEBUG), true)
	CFLAGS += -DDEBUG
endif

SRC = $(wildcard ./*.go)

.PHONY:
all: $(TARGET)

.PHONY:
$(TARGET): $(SRC)
	$(GO) build
	
test: $(TARGET)
	#@bash ./run-tests

install: $(TARGET)
	mkdir -p $(PREFIX)/bin
	cp -f $(TARGET) $(PREFIX)/bin

package: $(SRC)
ifeq (, $(shell which gox))
$(error "gox not installed; run: go get github.com...")
endif
	TARGET_VERSION=$(TARGET_VERSION) ./package.sh

clean:
	rm -f $(TARGET)
	rm -rf binaries

uninstall: $(PREFIX)/$(TARGET)
	rm -f $(PREFIX)/$(TARGET)

