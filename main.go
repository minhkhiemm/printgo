package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "missing image")
		os.Exit(2)
	}
	for _, path := range os.Args[1:] {
		if err := cat(path); err != nil {
			fmt.Fprintf(os.Stderr, "could not cat the image")
		}

	}
}

func cat(path string) error {
	f, err := os.Open(path)
	if err != nil {
		fmt.Errorf("could not open path")
	}

	defer f.Close()
	//ESC ] 1337 ; File = [optional arguments] : base-64 encoded file contents ^G
	fmt.Printf("\033]1337;File=inline=1:")
	wc := base64.NewEncoder(base64.StdEncoding, os.Stdout)
	_, err = io.Copy(wc, f)
	if err != nil {
		fmt.Errorf("could not copy to Stdout")
	}
	err = wc.Close()
	if err != nil {
		fmt.Errorf("could not close writter")
	}
	fmt.Printf("\a\n")

	return nil
}
