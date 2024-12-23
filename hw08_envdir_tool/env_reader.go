package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

const (
	prohibitedFilenameRunes = "="
	trimRunes               = " \t"
)

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	envs := make(Environment)

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if !file.Type().IsRegular() {
			continue
		}
		if strings.ContainsAny(file.Name(), prohibitedFilenameRunes) {
			continue
		}
		filePath := filepath.Join(dir, file.Name())
		if err != nil {
			return nil, err
		}
		fileStat, err := os.Stat(filePath)
		if err != nil {
			return nil, err
		}
		fileData, err := os.ReadFile(filePath)
		if err != nil {
			return nil, err
		}
		envs[strings.TrimRight(file.Name(), "=")] = EnvValue{
			Value:      processValue(fileData),
			NeedRemove: fileStat.Size() == 0,
		}
	}
	return envs, nil
}

func processValue(data []byte) string {
	value := strings.Split(string(data), "\n")[0]
	value = strings.TrimRight(value, trimRunes)
	value = string(bytes.ReplaceAll([]byte(value), []byte("\x00"), []byte("\n")))
	return value
}
