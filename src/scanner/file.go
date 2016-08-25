package scanner

import (
	"bufio"
	"os"
	"io"
)

type  File struct {
	Lines []string
	currentLn int
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
		f.Lines = append(f.Lines, scanner.Text())
	}

	f.currentLn = 0
	f.maxLen = len(f.Lines)
	return f, nil
}

func (f *File) NextLine() (string, error) {
	if f.currentLn >= f.maxLen {
		return "", io.EOF
	}
	ret:= f.Lines[f.currentLn]
	f.currentLn += 1
	return ret, nil
}
