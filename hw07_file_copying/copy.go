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

func Copy(fromPath, toPath string, limit, offset int64) error {

	inFile, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	defer inFile.Close()

	inFileStat, err := inFile.Stat()
	if err != nil {
		return err
	}
	if offset > inFileStat.Size() {
		return ErrOffsetExceedsFileSize
	}
	_, err = inFile.Seek(offset, 0)
	if err != nil {
		return err
	}

	outFile, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	trigger := false

	bufSize := 1024
	if limit != 0 && limit < int64(bufSize) {
		bufSize = int(limit)
		trigger = true
	}

	buf := make([]byte, bufSize)
	for {

		r, err := inFile.Read(buf)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		_, wErr := outFile.Write(buf[:r])
		if wErr != nil {
			return wErr
		}
		if err == io.EOF {
			break
		}
		if trigger {
			break
		}
	}

	return nil
}
