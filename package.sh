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

mkdir -p binaries
cd binaries
mkdir -p x86_64_linux/${TARGET_VERSION}/bin
mkdir -p macos/${TARGET_VERSION}/bin
mkdir -p armv6l/${TARGET_VERSION}/bin

gox -osarch="linux/amd64 darwin/amd64 linux/arm" ../

mv srm_linux_amd64 x86_64_linux/${TARGET_VERSION}/bin
mv srm_linux_arm armv6l/${TARGET_VERSION}/bin
mv srm_darwin_amd64 macos/${TARGET_VERSION}/bin

tar -czf srm-v${TARGET_VERSION}-x86_64_linux.tar.gz x86_64_linux
tar -czf srm-v${TARGET_VERSION}-macos.tar.gz macos
tar -czf srm-v${TARGET_VERSION}-armv6l.tar.gz armv6l

mkdir -p dist

mv srm-v${TARGET_VERSION}-x86_64_linux.tar.gz dist
mv srm-v${TARGET_VERSION}-macos.tar.gz dist
mv srm-v${TARGET_VERSION}-armv6l.tar.gz dist

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


