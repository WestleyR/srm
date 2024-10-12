# Safe Remove (`rm`) command with cache/undo

This is a `rm` command imitation, but without actually removing anything, only
moving it into cache (`~/.cache/srm3`). By doing this, you can recover
accidentally-removed files.

## Install

If you have go installed, then you can run:

```
GOPROXY=direct go install github.com/WestleyR/srm@main
```

Or via clone:

```
git clone https://github.com/WestleyR/srm
cd srm/
make
sudo make install # Or without root: 'make install PREFIX=${HOME}/.local'
```

### Aliasing

Optional to add this to your `~/.bashrc` or `~/.bash_profile`:

```
alias rm="srm"
```

### Linking

Instead or addition to aliasing, you can symlink srm -> rm in a first-search
path directory.

```
$ echo $PATH
/usr/local/sbin:/usr/local/bin ...

# ln -s /usr/local/bin/srm /usr/local/sbin/rm
```

This way, `rm` will always run `srm`, even for other users. You can always run
the normal `rm` by calling `/bin/rm ...`.

## Disclaimer

This is not designed to be a foolproof rm command, you still need to be careful.

## License

This project is licensed under the terms of The Clear BSD License. See the
[LICENSE file](./LICENSE) for more details.

