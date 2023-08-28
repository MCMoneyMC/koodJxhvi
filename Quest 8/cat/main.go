package main

import (
	"io"
	"os"
)

func main() {
	args := os.Args

	if len(args) == 1 {
		// Read from stdin and print to stdout
		_, err := io.Copy(os.Stdout, os.Stdin)
		if err != nil {
			os.Stderr.Write([]byte("ERROR: " + err.Error() + "\n"))
			os.Exit(1)
		}
	} else {
		// Read each file and print its contents to stdout
		for _, arg := range args[1:] {
			file, err := os.Open(arg)
			if err != nil {
				os.Stderr.Write([]byte("ERROR: " + err.Error() + "\n"))
				os.Exit(1)
			}
			_, err = io.Copy(os.Stdout, file)
			file.Close()
			if err != nil {
				os.Stderr.Write([]byte("ERROR: " + err.Error() + "\n"))
				os.Exit(1)
			}
		}
	}
}
