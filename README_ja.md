[English](README.md) | [ๆฅๆฌ่ช](README_ja.md)

# ๐ก dango

[![test](https://github.com/ebc-2in2crc/dango/actions/workflows/test.yml/badge.svg)](https://github.com/ebc-2in2crc/dango/actions/workflows/test.yml)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/ebc-2in2crc/dango)](https://goreportcard.com/report/github.com/ebc-2in2crc/dango)
[![Go Version](https://img.shields.io/github/go-mod/go-version/ebc-2in2crc/dango)](https://img.shields.io/github/go-mod/go-version/ebc-2in2crc/dango)
[![Version](https://img.shields.io/github/release/ebc-2in2crc/dango.svg?label=version)](https://img.shields.io/github/release/ebc-2in2crc/dango.svg?label=version)

๐ก `dango` ใฏๆจๆบๅฅๅใ้ฃ็ตใใใๅๅฒใใใใญใฐใฉใ ใงใใ

## Description

`dango` ใฏๆจๆบๅฅๅใ้ฃ็ตใใใๅๅฒใใพใใ

`dango` ใฏๆจๆบๅฅๅใใใคใใๆๅญใๅ่ชใใใใฏๆน่กใซใใฃใฆ้ฃ็ตใใใๅๅฒใใพใใ
`dango` ใฏๆจๆบๅฅๅใ่ชญใฟๅใฃใฆใ่ชญใฟๅใฃใ่ฆ็ด ใๆๅฎใใๆฐใซใชใใใจใซใใใ1่กใซใพใจใใฆใๆจๆบๅบๅใซๅบๅใใพใใ

## Usage

```
$ cat << EOF | dango
โ
๐ 
๐ก
๐ข
โ
EOF
โ๐ ๐ก๐ขโ
```

ใใคใๆฐใง้ฃ็ตใใใๅๅฒใใใ

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

ๆๅญๆฐใง้ฃ็ตใใใๅๅฒใใใ

```
$ echo 'โ๐ ๐ก๐ขโ' | dango -c -n 1
โ
๐ 
๐ก
๐ข
โ

$ echo 'โ๐ ๐ก๐ขโ' | dango -c -n 3
โ๐ ๐ก
๐ขโ
```

ๅ่ชๆฐใง้ฃ็ตใใใๅๅฒใใใ

```
$ echo 'Dango is a program that concatenates and splits standard input.' | dango -w -n 2 -d ' '
Dango is
a program
that concatenates
and splits
standard input.
```

ๆน่กใฎๆฐใง้ฃ็ตใใใๅๅฒใใใ

```
$ seq 1 4 | dango -l -n 2
12
34

$ seq 1 4 | dango -l
1234
```

ใใชใใฟใผใๆๅฎใใฆ้ฃ็ตใใใๅๅฒใใใ

```
$ echo 'dango' | dango -c -d '-'
d-a-n-g-o

$ seq 1 4 | dango -l -d ' '
1 2 3 4
```

ใใซใใ

```
$ dango -help
# ...
```

Docker ใไฝฟใใใจใใงใใพใใ

```
$ cat << EOF | docker container run --rm -i ebc2in2crc/dango
โ
๐ 
๐ก
๐ข
โ
EOF
โ๐ ๐ก๐ขโ
```

## Installation

### Developer

Go 1.16 ไปฅ้ใ

```
$ go install github.com/ebc-2in2crc/dango@latest
```

Go 1.15ใ

```
$ go get github.com/ebc-2in2crc/dango/...
```

### User

ๆฌกใฎ URL ใใใใฆใณใญใผใใใพใใ

- [https://github.com/ebc-2in2crc/dango/releases](https://github.com/ebc-2in2crc/dango/releases)

Homebrew ใไฝฟใใใจใใงใใพใ (Mac ใฎใฟ)

```
$ brew tap ebc-2in2crc/tap
$ brew install dango
```

Docker ใไฝฟใใใจใใงใใพใใ

```
$ docker image pull ebc2in2crc/dango
```

## ๐กdango ใจไปใฎใใผใซใฎๆฏ่ผ

`dango -l` ใจ `tr`

```
$ seq 1 4 | dango -l
1234

$ seq 1 4 | tr -d '\n'
1234
```

`dango -w` ใจ `xargs`

```
$ seq -s ' ' 1 4
1 2 3 4

$ seq -s ' ' 1 4 | dango -w -n 2 -d ' '
1 2
3 4

$ seq -s ' ' 1 4 | xargs -n2
1 2
3 4
```

`dango -c` ใจ `fold`

```
$ seq -s '' 1 4 | dango -c -n 2
12
34

$ seq -s '' 1 4 | fold -w 2
12
34
```

`dango -b` ใจ `fold`

```
$ seq -s '\t' 1 4 | dango -b -n 3
1	2
	3
4

$ seq -s '\t' 1 4 | fold -b -w 3
1	2
	3
4
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
