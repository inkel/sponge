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
	var buf []byte

	var total int

	for {
		b := make([]byte, bufsize)
		n, err := in.Read(b)
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return fmt.Errorf("reading from stdin: %w", err)
		}

		total += n

		buf = append(buf, b...)
	}

	buf = buf[:total]

	if outname == "" {
		_, err := os.Stdout.Write(buf)
		return err
	}

	fp, err := os.Create(outname)
	if err != nil {
		return fmt.Errorf("opening output file: %w", err)
	}

	_, err = fp.Write(buf)
	if err != nil {
		return fmt.Errorf("writing buffer to output file: %w", err)
	}

	return nil
}
