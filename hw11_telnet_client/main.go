package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

type WrapReader struct {
	io.Reader
}

func (w *WrapReader) Close() error {
	return nil
}

var timeout time.Duration

func init() {
	flag.DurationVar(&timeout, "timeout", 10*time.Second, "connection timeout")
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 2 {
		fmt.Println("Usage: go-telnet --timeout=10s host port")
		os.Exit(1)
	}
	host := args[0]
	port := args[1]

	if intPort, err := strconv.Atoi(port); err != nil {
		fmt.Printf("Invalid port: %s\n", port)
		os.Exit(1)
	} else if intPort <= 1023 {
		fmt.Printf("WARNING! System reserved port 1-1023: %s\n", port)
	} else if intPort > 65535 {
		fmt.Printf("Invalid port, out of range(1-65535) : %s\n", port)
		os.Exit(1)
	}

	address := net.JoinHostPort(host, port)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, os.Interrupt)

	tc := NewTelnetClient(address, timeout, os.Stdin, os.Stdout)
	if err := tc.Connect(ctx); err != nil {
		fmt.Printf("Connection error: %v. %v \n\r", address, err)
		os.Exit(1)
	}
	defer tc.Close()

	go func() {
		if err := tc.Receive(); err != nil {
			log.Fatal(err)
		}
		stop()
	}()

	go func() {
		if err := tc.Send(); err != nil {
			log.Fatal(err)
		}
		stop()
	}()

	<-ctx.Done()
	stop()
}
