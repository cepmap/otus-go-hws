package main

import (
	"errors"
	"io"
	"log"
	"os"
	"syscall"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile          = errors.New("unsupported file")
	ErrOffsetExceedsFileSize    = errors.New("offset exceeds file size")
	ErrSameFile                 = errors.New("input and output could not be the same")
	ErrFileIsDir                = errors.New("file is a directory")
	ErrNoLimitedDeviceOperation = errors.New("no limited device operation")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	if fromPath == toPath {
		return ErrSameFile
	}

	inFile, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	defer inFile.Close()

	outFile, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	if err := checkFiles(inFile, outFile, limit); err != nil {
		return err
	}

	inFileStat, err := inFile.Stat()
	if err != nil {
		return err
	}
	inLen := inFileStat.Size()

	log.Println(offset, inLen)
	if offset > inLen {
		return ErrOffsetExceedsFileSize
	}

	_, err = inFile.Seek(offset, 0)
	if err != nil {
		return err
	}

	var bSize int64
	if inLen == 0 || (limit != 0 && inLen-offset > limit) {
		bSize = limit
	} else {
		bSize = inLen - offset
	}

	pBar := pb.StartNew(int(bSize))
	pBar.Set(pb.Bytes, true)

	bytesCount := 4 * 1024
	writeBuffer := make([]byte, bytesCount)
	readCounter := 0
	breakTrigger := false

	for {
		readActual := bytesCount
		if limit != 0 && ((readCounter + bytesCount) > int(limit)) {
			readActual = int(limit) - readCounter
			breakTrigger = true
		}

		r, err := inFile.Read(writeBuffer[:readActual])
		if err != nil && err != io.EOF {
			return err
		}
		readCounter += r
		if r < readActual {
			readActual = r
		}

		if err == io.EOF {
			break
		}

		outFile.Write(writeBuffer[:readActual])
		pBar.Add(readActual)
		if breakTrigger {
			break
		}
	}
	pBar.Finish()
	return nil
}

func checkFiles(inFile, outFile *os.File, limit int64) error {
	isDevice := false
	files := []*os.File{inFile, outFile}
	for _, file := range files {
		fileInfo, err := file.Stat()
		if err != nil {
			return err
		}

		if fileInfo.IsDir() {
			return ErrFileIsDir
		}

		linkInfo, err := os.Lstat(file.Name())
		if err != nil {
			return err
		}
		if linkInfo.Mode()&os.ModeSymlink != 0 {
			return ErrUnsupportedFile
		}

		if fileInfo.Sys().(*syscall.Stat_t).Mode&syscall.S_IFBLK != 0 {
			isDevice = true
		}
	}

	if isDevice && limit == 0 {
		return ErrNoLimitedDeviceOperation
	}

	return nil
}
