package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Code-Hex/takopi"
)

func main() {
	if err := run(context.Background()); err != nil {
		fmt.Fprintf(os.Stderr, "err: %+v", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	msg, err := getMessage()
	if err != nil {
		return fmt.Errorf("failed to get message: %w", err)
	}
	fmt.Println(takopi.Do(msg))
	return nil
}

func getMessage() (string, error) {
	if len(os.Args[1:]) > 0 {
		return os.Args[1], nil
	}
	msg, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return "", fmt.Errorf("read error: %w", err)
	}
	return string(msg), nil
}
