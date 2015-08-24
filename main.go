package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	go func() {
		for {
			fmt.Fprint(os.Stdout, ">>> stdout\n")
			time.Sleep(1 * time.Second)
			fmt.Fprint(os.Stderr, ">>> stderr\n")
			time.Sleep(1 * time.Second)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Printf("<<< %q\n", scanner.Text())
	}
}
