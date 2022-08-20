package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

const bufsize = 8192

func main() {
	var outname string

	if len(os.Args) > 1 {
		outname = os.Args[1]
	}

	if err := realMain(os.Stdin, outname); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func realMain(in io.Reader, outname string) error {
	buf, err := read(in)
	if err != nil {
		return err
	}

	var out io.Writer

	if outname == "" {
		out = os.Stdout
	} else {
		fp, err := os.Create(outname)
		if err != nil {
			return fmt.Errorf("opening output file: %w", err)
		}
		defer fp.Close()

		out = fp
	}

	_, err = out.Write(buf)
	if err != nil {
		return fmt.Errorf("writing buffer to output file: %w", err)
	}

	return nil
}

func read(in io.Reader) ([]byte, error) {
	var buf []byte

	var total int

	for {
		b := make([]byte, bufsize)
		n, err := in.Read(b)
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return nil, fmt.Errorf("reading from stdin: %w", err)
		}

		total += n

		buf = append(buf, b...)
	}

	return buf[:total], nil
}
