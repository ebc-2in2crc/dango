package main

import (
	"bytes"
	"strings"
	"testing"
)

type param struct {
	args     string
	stdin    string
	exitCode int
	stdout   string
	stderr   string
}

func TestSplitLines(t *testing.T) {
	params := []param{
		{args: "-l -n 1", stdin: "first\nsecond\n", exitCode: 0, stdout: "first\nsecond\n"},
		{args: "-l -n 2", stdin: "first\nsecond\n", exitCode: 0, stdout: "firstsecond\n"},
		{args: "-l -n 2", stdin: "first\n\nsecond\n", exitCode: 0, stdout: "first\nsecond\n"},
		{args: "-l -n 2 -d _", stdin: "first\nsecond\n", exitCode: 0, stdout: "first_second\n"},
		{args: "-l -n 3", stdin: "first\nsecond\nthird\n", exitCode: 0, stdout: "firstsecondthird\n"},
		{args: "-l -n 3", stdin: "first\n\nsecond\n\nthird\n", exitCode: 0, stdout: "firstsecond\nthird\n"},
	}

	testRun(t, params)
}

func testRun(t *testing.T, params []param) {
	for _, p := range params {
		c, outStream, errStream := newCLI(p.stdin)
		args := strings.Split(" "+p.args, " ")

		exitCode := c.run(args)

		if exitCode != p.exitCode {
			t.Errorf("run(%s): Output = %q; want %q", args, exitCode, p.exitCode)
		}
		if outStream.String() != p.stdout {
			t.Errorf("run(%s): Output = %q; want %q", args, outStream.String(), p.stdout)
		}
		if errStream.String() != p.stderr {
			t.Errorf("run(%s): Err output = %q; want %q", args, errStream.String(), p.stderr)
		}
	}
}

func newCLI(stdin string) (c cli, outStream, errStream *bytes.Buffer) {
	inStream := bytes.NewBufferString(stdin)
	outStream = bytes.NewBuffer(nil)
	errStream = bytes.NewBuffer(nil)
	c = cli{inStream: inStream, outStream: outStream, errStream: errStream}
	return
}

func TestSplitWords(t *testing.T) {
	params := []param{
		{args: "-w -n 1", stdin: "first\nsecond\n", exitCode: 0, stdout: "first\nsecond\n"},
		{args: "-w -n 1", stdin: "first second\n", exitCode: 0, stdout: "first\nsecond\n"},
		{args: "-w -n 2", stdin: "first\nsecond\n", exitCode: 0, stdout: "firstsecond\n"},
		{args: "-w -n 2", stdin: "first\n\nsecond\n", exitCode: 0, stdout: "firstsecond\n"},
		{args: "-w -n 2 -d _", stdin: "first\nsecond\n", exitCode: 0, stdout: "first_second\n"},
		{args: "-w -n 3", stdin: "first\nsecond\nthird\n", exitCode: 0, stdout: "firstsecondthird\n"},
		{args: "-w -n 3", stdin: "first\n\nsecond\n\nthird\n", exitCode: 0, stdout: "firstsecondthird\n"},
	}

	testRun(t, params)
}

func TestSplitCharacters(t *testing.T) {
	params := []param{
		{args: "-c -n 1", stdin: "12\n", exitCode: 0, stdout: "1\n2\n\n\n"},
		{args: "-c -n 2", stdin: "1\n2\n", exitCode: 0, stdout: "1\n\n2\n\n"},
		{args: "-c -n 2", stdin: "1234\n", exitCode: 0, stdout: "12\n34\n\n\n"},
		{args: "-c -n 2", stdin: "12\n\n34\n", exitCode: 0, stdout: "12\n\n\n\n34\n\n\n"},
		{args: "-c -n 2 -d _", stdin: "1234\n", exitCode: 0, stdout: "1_2\n3_4\n\n\n"},
		{args: "-c -n 2", stdin: "１\n２\n", exitCode: 0, stdout: "１\n\n２\n\n"},
		{args: "-c -n 3", stdin: "１\n２\n", exitCode: 0, stdout: "１\n２\n\n\n"},
	}

	testRun(t, params)
}

func TestSplitBytes(t *testing.T) {
	params := []param{
		{args: "-b -n 1", stdin: "12\n", exitCode: 0, stdout: "1\n2\n\n\n"},
		{args: "-b -n 2", stdin: "12\n", exitCode: 0, stdout: "12\n\n\n"},
		{args: "-b -n 2", stdin: "1\n2\n", exitCode: 0, stdout: "1\n\n2\n\n"},
		{args: "-b -n 2", stdin: "1234\n", exitCode: 0, stdout: "12\n34\n\n\n"},
		{args: "-b -n 2", stdin: "12\n\n34\n", exitCode: 0, stdout: "12\n\n\n\n34\n\n\n"},
		{args: "-b -n 2 -d _", stdin: "1234\n", exitCode: 0, stdout: "1_2\n3_4\n\n\n"},
	}

	testRun(t, params)
}

func TestBufferSize(t *testing.T) {
	params := []param{
		{args: "-l -n 1 -B 1", stdin: "a\n", exitCode: 1, stdout: "", stderr: "failed to scan stdin: bufio.Scanner: token too long\n"},
		{args: "-l -n 1 -B 2", stdin: "a\n", exitCode: 0, stdout: "a\n", stderr: ""},
	}

	testRun(t, params)
}
