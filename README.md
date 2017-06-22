# goed

**Go edit that Go file!**

This is a basic wrapper for my editor session.
It runs the go formatter, linter, and whatnot.
Optionally, it can also run the build or install tool.

## Get it

```bash
go install github.com/ldfritz/goed
```

## Basic usage

```bash
goed filename.go
```

## TODO

* check for required extra commands
* detect missing files
* add testing flags
* add verbose output
