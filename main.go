package main

import (
	flag "github.com/spf13/pflag"
	"fmt"
	"os"
)

var(
	note string
)

func main() {
	flag.Parse()

	if flag.NFlag() == 0 {
		fmt.Printf("Usage: %s [options]\n", os.Args[0])
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	fmt.Printf("Your note is %s\n", note)
}

func init() {
	flag.StringVarP(&note, "jot", "j", "", "WriteNoteDown")
}
