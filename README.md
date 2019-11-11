# gofmtmd

<img src="https://img.shields.io/badge/go-v1.13-blue.svg"/> [![GoDoc](https://godoc.org/github.com/po3rin/gofmtmd?status.svg)](https://godoc.org/github.com/po3rin/gofmtmd) [![CircleCI](https://circleci.com/gh/po3rin/gofmtmd.svg?style=shield)](https://circleci.com/gh/po3rin/gofmtmd) [![codecov](https://codecov.io/gh/po3rin/gofmtmd/branch/master/graph/badge.svg)](https://codecov.io/gh/po3rin/gofmtmd) [![GolangCI](https://golangci.com/badges/github.com/po3rin/gofmtmd.svg)](https://golangci.com)

<img src="image/cover.png" width="640px"/>

gofmtmd formats go source code block in Markdown. detects fenced code & formats code using gofmt.

## Usage

```bash
# replace Go code with formated code
$ gofmtmd testdata/testdata.md -r

# write result to file instead of stdout
$ gofmtmd testdata/testdata.md -w formatted.md
```

## Help

```bash
$ gofmtmd -h
This CLI formats Go Code in Markdown.

Usage:
  gofmtmd [flags]

Flags:
  -h, --help           help for gofmtmd
  -r, --replace        replace Go code with formated code
      --version        version for gofmtmd
  -w, --write string   write result to file instead of stdout
```
