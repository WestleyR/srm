#
#  package.sh
#  srm - https://github.com/WestleyR/srm
#
# Created by WestleyR on July 29, 2020
# Source code: https://github.com/WestleyR/srm
#
# Copyright (c) 2020-2022 WestleyR. All rights reserved.
# This software is licensed under a BSD 3-Clause Clear License.
# Consult the LICENSE file that came with this software regarding
# your rights to distribute this software.
#

set -e

GO_CMD="go"
GO_FLAGS="-ldflags -w"

echo "Target version: ${TARGET_VERSION}"
echo "Go compiler: ${GO_CMD} with build flags: ${GO_FLAGS}"

mkdir -p binaries
cd binaries
mkdir -p x86_64_linux/srm/bin
mkdir -p macos/srm/bin
mkdir -p armv6l/srm/bin

# Compile for a couple common os/arch
GOOS=linux GOARCH=amd64 $GO_CMD build $GO_FLAGS -o srm_linux_amd64 ../cmd/srm/main.go
GOOS=darwin GOARCH=amd64 $GO_CMD build $GO_FLAGS -o srm_darwin_amd64 ../cmd/srm/main.go
GOOS=linux GOARCH=arm GOARM=5 $GO_CMD build $GO_FLAGS -o srm_linux_arm ../cmd/srm/main.go

# Copy the compiled binaries to the tarball dirs
mv srm_linux_amd64 x86_64_linux/srm/bin/srm
mv srm_linux_arm armv6l/srm/bin/srm
mv srm_darwin_amd64 macos/srm/bin/srm

# Copy the LICENSE file and README to the tarball
cp ../LICENSE x86_64_linux/srm
cp ../LICENSE macos/srm
cp ../LICENSE armv6l/srm
cp ../README.md x86_64_linux/srm
cp ../README.md macos/srm
cp ../README.md armv6l/srm

# Copy the third-party licenses to the tarball
cp -r ../THIRD_PARTY_LICENSES x86_64_linux/srm/
cp -r ../THIRD_PARTY_LICENSES macos/srm/
cp -r ../THIRD_PARTY_LICENSES armv6l/srm/

cd x86_64_linux
tar -czf srm-${TARGET_VERSION}-x86_64_linux.tar.gz srm
cd ../macos
tar -czf srm-${TARGET_VERSION}-macos.tar.gz srm
cd ../armv6l
tar -czf srm-${TARGET_VERSION}-armv6l.tar.gz srm
cd ../

mkdir -p dist

mv x86_64_linux/srm-${TARGET_VERSION}-x86_64_linux.tar.gz dist
mv macos/srm-${TARGET_VERSION}-macos.tar.gz dist
mv armv6l/srm-${TARGET_VERSION}-armv6l.tar.gz dist

cd dist

echo
echo "Binary packages are at: binaries/dist/*"

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
echo "sha1sum"
for p in `ls` ; do
	sha1sum ${p} || echo "No sha1sum command found"
done

echo
echo "sha256sum"
for p in `ls` ; do
	sha256sum ${p} || echo "No sha256sum command found"
done


