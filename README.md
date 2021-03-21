[English](README.md) | [日本語](README_ja.md)

# 🍡 dango

[![test](https://github.com/ebc-2in2crc/dango/actions/workflows/test.yml/badge.svg)](https://github.com/ebc-2in2crc/dango/actions/workflows/test.yml)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/ebc-2in2crc/dango)](https://goreportcard.com/report/github.com/ebc-2in2crc/dango)
[![Go Version](https://img.shields.io/github/go-mod/go-version/ebc-2in2crc/dango)](https://img.shields.io/github/go-mod/go-version/ebc-2in2crc/dango)
[![Version](https://img.shields.io/github/release/ebc-2in2crc/dango.svg?label=version)](https://img.shields.io/github/release/ebc-2in2crc/dango.svg?label=version)

🍡`dango` is a program that concatenates and splits standard input.

## Description

`dango` concatenates and splits the standard input.

`dango` concatenates and splits standard input by bytes, characters, words, or line breaks.
`dango` reads the standard input and whenever it reads a specified number of elements, it puts them on a line and prints them to the standard output.

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

Concatenate or split by number of bytes.

```
$ echo 'dango' | dango -b -n 1
d
a
n
g
o

$ echo 'dango' | dango -b -n 3
dan
go
```

Concatenate or split by number of characters.

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

Concatenate or split by word count.

```
$ echo 'Dango is a program that concatenates and splits standard input.' | dango -w -n 2 -d ' '
Dango is
a program
that concatenates
and splits
standard input.
```

Concatenate or split by the number of line breaks.

```
$ seq 1 4 | dango -l -n 2
12
34

$ seq 1 4 | dango -l
1234
```

Specify a delimiter to concatenate or split.

```
$ echo 'dango' | dango -c -d '-'
d-a-n-g-o

$ seq 1 4 | dango -l -d ' '
1 2 3 4
```

Show help.

```
$ dango -help
# ...
```

Or, you can use Docker.

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

Go 1.16 or later.

```
$ go install github.com/ebc-2in2crc/dango@latest
```

Go 1.15.

```
$ go get github.com/ebc-2in2crc/dango/...
```

### User

Download from the following url.

- [https://github.com/ebc-2in2crc/dango/releases](https://github.com/ebc-2in2crc/dango/releases)

Or, you can use Homebrew (Only macOS).

```
$ brew tap ebc-2in2crc/tap
$ brew install dango
```

Or, you can use Docker.

```
$ docker image pull ebc2in2crc/dango
```

## Contribution

1. Fork this repository
2. Create your feature branch (git checkout -b my-new-feature)
3. Commit your changes (git commit -am 'Add some feature')
4. Rebase your local changes against the master branch
5. Run test suite with the `make test` command and confirm that it passes
6. Run `make fmt`
7. Create new Pull Request

## License

[MIT](https://github.com/ebc-2in2crc/wareki/blob/master/LICENSE)

## Author

[ebc-2in2crc](https://github.com/ebc-2in2crc)
