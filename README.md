# Safe Re-Move (`rm`) command with cache/undo

This is a `rm` command imitation, but without actrally removing anything, only
caching it. By doing this, you can recover accidentally-removed files.

<br>

## Install

Install by cloning the source code:

```
git clone https://github.com/WestleyR/srm
cd srm/
make
sudo make install  # Or without root: make install PREFIX=${HOME}/.local
```

Optional, but recommended

```
alias rm="srm"
```

<br>

### Examples

Nothing much to show, its just a rm command.

<br>

## License

```
The Clear BSD License

Copyright (c) 2019 WestleyR
All rights reserved.
```

See the [LICENSE file](LICENSE)
for more details.

<br>

