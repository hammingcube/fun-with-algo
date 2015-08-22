package main

import (
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	cmd := exec.Command("diff", "/dev/fd/3", "/dev/fd/4")

	foo, err := readerFile(strings.NewReader("foo\n"))
	if err != nil {
		log.Fatal(err)
	}
	defer foo.Close()

	bar, err := readerFile(strings.NewReader("bar\n"))
	if err != nil {
		log.Fatal(err)
	}
	defer bar.Close()

	cmd.ExtraFiles = []*os.File{foo, bar}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func readerFile(r io.Reader) (*os.File, error) {
	reader, writer, err := os.Pipe()

	if err != nil {
		return nil, err
	}

	go func() {
		io.Copy(writer, r)
		writer.Close()
	}()

	return reader, nil
}
