# gofmtmd

<img src="https://img.shields.io/badge/go-v1.13-blue.svg"/> [![GoDoc](https://godoc.org/github.com/po3rin/gofmtmd?status.svg)](https://godoc.org/github.com/po3rin/gofmtmd) [![CircleCI](https://circleci.com/gh/po3rin/gofmtmd.svg?style=shield)](https://circleci.com/gh/po3rin/gofmtmd) [![codecov](https://codecov.io/gh/po3rin/gofmtmd/branch/master/graph/badge.svg)](https://codecov.io/gh/po3rin/gofmtmd) [![GolangCI](https://golangci.com/badges/github.com/po3rin/gofmtmd.svg)](https://golangci.com)

<img src="image/cover.png" width="640px"/>

gofmtmd formats go source code block in Markdown. detects fenced code & formats code using gofmt.

## Installation

```
$ go get github.com/po3rin/gofmtmd/cmd/gofmtmd
```

## Usage

```bash
# replace Go code with formatted code
$ gofmtmd testdata/testdata.md -r

# write result to file instead of stdout
$ gofmtmd testdata/testdata.md -w formatted.md

# you can use stdndard input
$ gofmtmd < testdata/testdata.md
$ echo "#hello" | gofmtmd
```

## Help

```bash
$ gofmtmd -h
This CLI formats Go Code in Markdown.

Usage:
  gofmtmd [flags]

Flags:
  -h, --help           help for gofmtmd
  -r, --replace        replace Go code with formatted code
      --version        version for gofmtmd
  -w, --write string   write result to file instead of stdout
```

## Vim Plugin

Vim plugin version is here! this plugin lets you to run automatically when saved.

<a href="https://github.com/po3rin/vim-gofmtmd"><img src="https://github-link-card.s3.ap-northeast-1.amazonaws.com/po3rin/vim-gofmtmd.png" width="460px"></a>
