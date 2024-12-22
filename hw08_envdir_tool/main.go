package main

import (
	"errors"
	"log"
	"os"
)

var ErrNoArgsProvided = errors.New("no args provided")

func main() {
	// В условии явно не указано, ну или я не вижу. Считаем что аргумент с путем до папки и командой обязательны.
	if len(os.Args) <= 3 {
		log.Fatal(ErrNoArgsProvided)
	}
	parsedEnvs, err := ReadDir(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	returnCode := RunCmd(os.Args[2:], parsedEnvs)
	os.Exit(returnCode)
}
