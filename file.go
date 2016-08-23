package gocc

import (
	"bufio"
	"os"

	"io"
)

type  File struct {
	Lines []string
	index int
	maxLen int
}



func NewFile(src string) (*File, error) {
	file, err := os.Open(src)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	f := new(File)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		append(f.Lines, scanner.Text())
	}

	f.index = -1
	f.maxLen = len(f.Lines)
	return f
}

func (f *File) nextLine() (string, error) {
	if f.index >= f.maxLen {
		return nil, io.EOF
	}
	f.index += 1
	return f.Lines[f.index], nil
}