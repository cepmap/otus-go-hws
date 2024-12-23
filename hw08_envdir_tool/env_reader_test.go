package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	// Place your code here
	tempDir := tmpDir(t)
	defer os.RemoveAll(tempDir)
	files := []struct {
		name  string
		value string
	}{
		{"HELLO", "hello"},
		{"FOO", "FOO"},
		{"EMPTY", " "},
		{"UNSET", ""},
		{"MUlTILINE", "1\n2\n3"},
	}
	for _, f := range files {
		createEnvFile(t, tempDir, f.name, f.value)
	}
	env, err := ReadDir(tempDir)
	if err != nil {
		t.Fatalf("Func ReadDir failed %v", err)
	}
	expectedEnv := Environment{
		"HELLO": EnvValue{
			Value:      "hello",
			NeedRemove: false,
		},
		"FOO": EnvValue{
			Value:      "FOO",
			NeedRemove: false,
		},
		"EMPTY": EnvValue{
			Value:      "",
			NeedRemove: false,
		},
		"UNSET": EnvValue{
			Value:      "",
			NeedRemove: true,
		},
		"MUlTILINE": EnvValue{
			Value:      "1",
			NeedRemove: false,
		},
	}

	require.Equal(t, expectedEnv, env)
}

func tmpDir(t *testing.T) string {
	t.Helper()
	tmpDir, err := os.MkdirTemp("", "test_read_dir")
	if err != nil {
		t.Fatalf("Failed to create temp test_read_dir: %v", err)
	}
	return tmpDir
}

func createEnvFile(t *testing.T, dir, name, value string) {
	t.Helper()
	filepath := dir + "/" + name
	err := os.WriteFile(filepath, []byte(value), 0o644)
	if err != nil {
		t.Fatalf("Failde to create envFile: %v", err)
	}
}
