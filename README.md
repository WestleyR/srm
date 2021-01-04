# Safe Re-Move (`rm`) command with cache/undo

[![Go Report Card](https://goreportcard.com/badge/github.com/WestleyR/srm)](https://goreportcard.com/report/github.com/WestleyR/srm)
[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![GitHub release](https://img.shields.io/github/release/WestleyR/srm.svg)](https://GitHub.com/WestleyR/srm/releases/)
[![Github all releases](https://img.shields.io/github/downloads/WestleyR/srm/total.svg)](https://GitHub.com/WestleyR/srm/releases/)

This is a `rm` command imitation, but without actually removing anything, only
moving it into cache (`~/.cache/srm`). By doing this, you can recover
accidentally-removed files.

_undo/list command still WIP..._

## Install

**Please see the [release page](https://github.com/WestleyR/srm/releases) for the latest pre-release.**

Install via package manager ([gpack](https://github.com/WestleyR/gpack)):

```
gpack install WestleyR/srm
```

Or install the Go dev code: _beta_

```
git clone https://github.com/WestleyR/srm
cd srm/
make
sudo make install  # Or without root: 'make install PREFIX=${HOME}/.local'
```

Optional, but recommended to add this to your `~/.bashrc` or `~/.bash_profile`:

```
alias rm="srm"
```

### Examples

Demo coming soon. However, this should be equivalent to the `rm` command.

## License

This project is licensed under the terms of The Clear BSD License. See the
[LICENSE file](./LICENSE) for more details.

