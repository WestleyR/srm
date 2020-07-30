#!/bin/bash
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

set -e

echo "Target version: ${TARGET_VERSION}"

mkdir -p binaries
cd binaries
mkdir -p x86_64_linux/srm/${TARGET_VERSION}/bin
mkdir -p macos/srm/${TARGET_VERSION}/bin
mkdir -p armv6l/srm/${TARGET_VERSION}/bin

gox -osarch="linux/amd64 darwin/amd64" ../

# Now compile for arm since there was some issues with gox
GOOS=linux GOARCH=arm GOARM=5 go build -o srm_linux_arm ../

mv srm_linux_amd64 x86_64_linux/srm/${TARGET_VERSION}/bin/srm
mv srm_linux_arm armv6l/srm/${TARGET_VERSION}/bin/srm
mv srm_darwin_amd64 macos/srm/${TARGET_VERSION}/bin/srm

cd x86_64_linux
tar -czf srm-v${TARGET_VERSION}-x86_64_linux.tar.gz srm
cd ../macos
tar -czf srm-v${TARGET_VERSION}-macos.tar.gz srm
cd ../armv6l
tar -czf srm-v${TARGET_VERSION}-armv6l.tar.gz srm
cd ../

mkdir -p dist

mv x86_64_linux/srm-v${TARGET_VERSION}-x86_64_linux.tar.gz dist
mv macos/srm-v${TARGET_VERSION}-macos.tar.gz dist
mv armv6l/srm-v${TARGET_VERSION}-armv6l.tar.gz dist

cd dist

echo
echo "#########"
echo "Checksums"
echo "#########"
echo

echo "md5sum"
for p in `ls` ; do
	md5sum ${p} || echo "No md5sum command found"
done

echo
echo "ssum"
for p in `ls` ; do
	ssum ${p} || echo "No ssum command found"
done

echo
echo "sha256sum"
for p in `ls` ; do
	sha256sum ${p} || echo "No sha256sum command found"
done


