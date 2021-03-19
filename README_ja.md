[English](README.md) | [日本語](README_ja.md)

# 🍡 dango

[![test](https://github.com/ebc-2in2crc/dango/actions/workflows/test.yml/badge.svg)](https://github.com/ebc-2in2crc/dango/actions/workflows/test.yml)

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
$ echo 'Hello World!' | dango -b -n 1
H
e
l
l
o

W
o
r
l
d
!

$ echo 'Hello World!' | dango -b -n 5
Hello
 Worl
d!
```

文字数で連結したり分割する。

```
$ echo '花より団子' | dango -c -n 1
花
よ
り
団
子

$ echo '花より団子' | dango -c -n 3
花より
団子
```

単語数で連結したり分割する。

```
$ echo 'Hello World!' | dango -w -n 1
Hello
World!

$ echo 'Hello World!' | dango -w
HelloWorld!
```

改行の数で連結したり分割する。

```
$ cat << EOF | dango -l -n 1
Hello
World!
EOF
Hello
World!

$ cat << EOF | dango -l
Hello

World

!
EOF
HelloWorld!
```

デリミターを指定して連結したり分割する。

```
$ echo 'abcdefg' | dango -c -d ' '
a b c d e f g 

$ cat << EOF | dango -d ' '
Hello
World!
EOF
Hello World!
```

ヘルプ。

```
$ dango -help
# ...
```

Docker を使うこともできます。

```
$ cat << EOF | docker container run --rm --interactive ebc2in2crc/dango
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
