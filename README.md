[English](README.md) | [日本語](README_ja.md)

# 🍡 dango

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

Concatenate or split by number of characters.

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

Concatenate or split by word count.

```
$ echo 'Hello World!' | dango -w -n 1
Hello
World!

$ echo 'Hello World!' | dango -w
HelloWorld!
```

Concatenate or split by the number of line breaks.

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

Specify a delimiter to concatenate or split.

```
$ echo 'abcdefg' | dango -c -d ' '
a b c d e f g 

$ cat << EOF | dango -d ' '
Hello
World!
EOF
Hello World!
```

Show help.

```
$ dango -help
# ...
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
