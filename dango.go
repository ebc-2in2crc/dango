package main

import (
	"bufio"
	"fmt"
	"io"

	"golang.org/x/xerrors"
)

const (
	startBufSize = 4096 // Size of initial allocation for buffer.
)

type dango struct {
	reader    io.Reader
	writer    io.Writer
	delimiter []byte
}

func (d *dango) run() error {
	switch {
	case options.bytes:
		return d.splitBytes()
	case options.lines:
		return d.splitLines()
	case options.words:
		return d.splitWords()
	case options.characters:
		return d.splitCharacters()
	default:
		return d.splitLines()
	}
}

func (d *dango) splitBytes() error {
	return d.joinTokens(bufio.ScanBytes)
}

func (d *dango) splitWords() error {
	return d.joinTokens(bufio.ScanWords)
}

func (d *dango) joinTokens(splitFunc bufio.SplitFunc) error {
	scanner := bufio.NewScanner(d.reader)
	scanner.Split(splitFunc)
	if options.maxBufferSize > 0 {
		size := startBufSize
		if size > options.maxBufferSize {
			size = options.maxBufferSize
		}
		b := make([]byte, size)
		scanner.Buffer(b, options.maxBufferSize)
	}

	limit := options.size
	count := 0
	for scanner.Scan() {
		t := scanner.Bytes()
		count++

		if err := d.addDelimiter(count, t); err != nil {
			return fmt.Errorf("failed to add delimiter: %w", err)
		}

		_, err := d.writer.Write(t)
		if err != nil {
			return fmt.Errorf("failed to write to stdout: %w", err)
		}
		if limit > 0 && count >= limit {
			_, err := d.writer.Write([]byte("\n"))
			if err != nil {
				return xerrors.Errorf("failed write LF: %w", err)
			}
			count = 0
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("failed to scan stdin: %w", err)
	}

	if count != 0 && count != limit {
		_, err := d.writer.Write([]byte("\n"))
		if err != nil {
			return xerrors.Errorf("failed write LF: %w", err)
		}
	}

	return nil
}

func (d *dango) addDelimiter(count int, t []byte) error {
	if count <= 1 {
		return nil
	}
	if len(t) == 1 && t[0] == '\n' {
		return nil
	}

	_, err := d.writer.Write(d.delimiter)
	if err != nil {
		return fmt.Errorf("failed to write delimiter to stdout: %w", err)
	}
	return nil
}

func (d *dango) splitLines() error {
	return d.joinTokens(bufio.ScanLines)
}

func (d *dango) splitCharacters() error {
	return d.joinTokens(bufio.ScanRunes)
}
