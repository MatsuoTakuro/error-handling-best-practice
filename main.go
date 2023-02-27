package main

import (
	"fmt"
	"os"
	"strconv"

	"golang.org/x/xerrors"
)

// ref to https://zenn.dev/malt03/articles/cd0365608a26c4
func main() {
	if err := parse(os.Args[1]); err != nil {
		fmt.Printf("%+v\n", xerrors.Errorf(": %w", err))
	}
}

func parse(s string) error {
	n, err := strconv.Atoi(s)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}
	if n < 10 {
		return xerrors.Errorf("n is %d", n)
	}

	fmt.Printf("n is %d\n", n)
	return nil
}
