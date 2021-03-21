[English](README.md) | [日本語](README_ja.md)

# 🍡 dango

[![test](https://github.com/ebc-2in2crc/dango/actions/workflows/test.yml/badge.svg)](https://github.com/ebc-2in2crc/dango/actions/workflows/test.yml)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/ebc-2in2crc/dango)](https://goreportcard.com/report/github.com/ebc-2in2crc/dango)
[![Go Version](https://img.shields.io/github/go-mod/go-version/ebc-2in2crc/dango)](https://img.shields.io/github/go-mod/go-version/ebc-2in2crc/dango)
[![Version](https://img.shields.io/github/release/ebc-2in2crc/dango.svg?label=version)](https://img.shields.io/github/release/ebc-2in2crc/dango.svg?label=version)

🍡 `dango` は標準入力を連結したり分割するプログラムです。

## Description

`dango` は標準入力を連結したり分割します。

`dango` は標準入力をバイト、文字、単語あるいは改行によって連結したり分割します。
`dango` は標準入力を読み取って、読み取った要素が指定した数になるごとにそれを1行にまとめて、標準出力に出力します。

## Usage

```
$ cat << EOF | dango
━
🟠
🟡
🟢
━
EOF
━🟠🟡🟢━
```

バイト数で連結したり分割する。

```
$ echo 'dango' | dango -b -n 3
dan
go

$ echo 'dango' | dango -b -n 3
dan
go
```

文字数で連結したり分割する。

```
$ echo '━🟠🟡🟢━' | dango -c -n 1
━
🟠
🟡
🟢
━

$ echo '━🟠🟡🟢━' | dango -c -n 3
━🟠🟡
🟢━
```

単語数で連結したり分割する。

```
$ echo 'Dango is a program that concatenates and splits standard input.' | dango -w -n 2 -d ' '
Dango is
a program
that concatenates
and splits
standard input.
```

改行の数で連結したり分割する。

```
$ seq 1 4 | dango -l -n 2
12
34

$ seq 1 4 | dango -l
1234
```

デリミターを指定して連結したり分割する。

```
$ echo 'dango' | dango -c -d '-'
d-a-n-g-o

$ seq 1 4 | dango -l -d ' '
1 2 3 4
```

ヘルプ。

```
$ dango -help
# ...
```

Docker を使うこともできます。

```
$ cat << EOF | docker container run --rm -i ebc2in2crc/dango
━
🟠
🟡
🟢
━
EOF
━🟠🟡🟢━
```

## Installation

### Developer

Go 1.16 以降。

```
$ go install github.com/ebc-2in2crc/dango@latest
```

Go 1.15。

```
$ go get github.com/ebc-2in2crc/dango/...
```

### User

次の URL からダウンロードします。

- [https://github.com/ebc-2in2crc/dango/releases](https://github.com/ebc-2in2crc/dango/releases)

Homebrew を使うこともできます (Mac のみ)

```
$ brew tap ebc-2in2crc/tap
$ brew install dango
```

Docker を使うこともできます。

```
$ docker image pull ebc2in2crc/dango
```

## Contribution

1. Fork this repository
2. Create your feature branch (git checkout -b my-new-feature)
3. Commit your changes (git commit -am 'Add some feature')
4. Rebase your local changes against the main branch
5. Run test suite with the `make test` command and confirm that it passes
6. Run `make fmt`
7. Create new Pull Request

## License

[MIT](https://github.com/ebc-2in2crc/dango/blob/main/LICENSE)

## Author

[ebc-2in2crc](https://github.com/ebc-2in2crc)
