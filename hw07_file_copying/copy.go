package main

import (
	"errors"
	"io"
	"log"
	"os"
)

//go run main.go copy.go -from testdata/input.txt -limit 0 -offset 0 -to testdata/output.txt

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	infile, err := os.Open(fromPath)
	if err != nil {
		return ErrUnsupportedFile
	}
	defer func(ifile *os.File) {
		err := ifile.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(infile)

	nfile, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer nfile.Close()

	infileStat, err := infile.Stat()
	if err != nil {
		return err
	}

	if offset > infileStat.Size() {
		return ErrOffsetExceedsFileSize
	}

	if infileStat.Size() < limit {
		limit = infileStat.Size()
	}

	_, err = infile.Seek(offset, io.SeekStart)
	if err != nil {
		return err
	}

	return nil
}
