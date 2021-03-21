package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

type cli struct {
	inStream             io.Reader
	outStream, errStream io.Writer
}

type dangoOptions = struct {
	size          int
	delimiter     string
	maxBufferSize int
	lines         bool
	words         bool
	characters    bool
	bytes         bool
	version       bool
	help          bool
}

var options dangoOptions

func init() {
	flag.IntVar(&options.maxBufferSize, "B", 0, "Maximum size used to buffer a token.")
	flag.BoolVar(&options.bytes, "b", false, "Concatenates or split standard input by bytes.")
	flag.BoolVar(&options.characters, "c", false, "Concatenates or split standard input by characters.")
	flag.StringVar(&options.delimiter, "d", "", `Element delimiter (default: "")`)
	flag.BoolVar(&options.lines, "l", false, "Concatenates or split standard input by lines.")
	flag.IntVar(&options.size, "n", 0, "Number of elements in each line (default: 0)")
	flag.BoolVar(&options.version, "version", false, "Show version.")
	flag.BoolVar(&options.help, "help", false, "Show this help message.")
	flag.BoolVar(&options.words, "w", false, "Concatenates or split standard input by words.")

	flag.Usage = func() {
		fmt.Printf(`NAME:
  dango - concatenates or split standard input

USAGE:
  dango [options]

DESCRIPTION:
  Concatenates or split standard input.

  e.g.
    $ seq 1 4 | dango -l -n 2
    12
    34
    
    $ seq 1 4 | dango -l
    1234

OPTIONS:
`)
		flag.PrintDefaults()
	}
}

func (c *cli) run(args []string) int {
	c.parseFlag(args)

	if options.version {
		_, _ = fmt.Fprintf(c.outStream, "dango version %s\n", version)
		return 0
	}
	if options.help {
		flag.Usage()
		return 0
	}

	d := dango{
		reader:    c.inStream,
		writer:    c.outStream,
		delimiter: []byte(options.delimiter),
	}
	err := d.run()
	if err != nil {
		msg := fmt.Sprintf("%+v\n", err)
		_, _ = c.errStream.Write([]byte(msg))
		return 1
	}

	return 0
}

func (c *cli) parseFlag(args []string) {
	os.Args = args
	options = dangoOptions{}
	flag.Parse()
}
