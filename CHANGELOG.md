# CHANGELOG for srm

_**NOTES:**_
 1. Versions containing "alpha", "a", "beta", "b", or "rc" are pre-releases, and
subject to change.

## v2.1.1 - 2021-12-12

### FIXED
 - Fixed panic issue


## v2.1.0 - ??

## v2.0.1 - Jan 28, 2021

### FIXED
 - Fixed the package version (duh!)


## v2.0.0 - Jan 28, 2021

### ADDED
 - Added `--list-cache` option
 - Added `--recover` option to recover deleted files

### CHANGED
 - Changed code from C to Go

### FIXED
 - Now checks all files in a dir for write-protected files


## v1.1.1 - Dec 22, 2019

### CHANGES
 - Fixed the issue when removing a file in a path


## v1.1.0 - Dec 18, 2019

### CHANGES
 - Fixed the `-f` when removing a directory (should fail)

### ADDED
 - Added `-r` flag to remove a directory


## v1.0.0 - Dec 17, 2019

First release of `srm`.

