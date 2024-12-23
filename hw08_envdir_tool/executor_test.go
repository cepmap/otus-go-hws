package main

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	cmd := []string{"printenv", "FOO", "BAR"}
	env := Environment{
		"FOO": EnvValue{
			Value:      "FOO",
			NeedRemove: false,
		},
		"BAR": EnvValue{
			Value:      "BAR",
			NeedRemove: false,
		},
	}
	output := getOutput(t, func() {
		RunCmd(cmd, env)
	})
	require.Equal(t, "FOO\nBAR\n", output)
}

func TestIsEmpty(t *testing.T) {
	cmd := []string{"printenv", "$PATH"}
	env := Environment{
		"PATH": EnvValue{
			Value:      "",
			NeedRemove: false,
		},
	}

	actualOutput := getOutput(t, func() {
		RunCmd(cmd, env)
	})

	require.Equal(t, "", actualOutput)
}

func getOutput(t *testing.T, f func()) string {
	t.Helper()
	reader, writer, _ := os.Pipe()
	savedStdout := os.Stdout
	os.Stdout = writer
	out := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, reader)
		out <- buf.String()
	}()

	f()
	writer.Close()
	os.Stdout = savedStdout

	return <-out
}
