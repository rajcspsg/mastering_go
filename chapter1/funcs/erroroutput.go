package funcs

import (
	"io"
	"os"
)

func StdErrorDemo() {
	myString := ""
	arguments := os.Args

	if len(arguments) == 1 {
		myString = "Please provide atleast single argument"
	} else {
		myString = arguments[1]
	}

	io.WriteString(os.Stdout, "This is standard output")
	io.WriteString(os.Stderr, myString)
	io.WriteString(os.Stderr, "\n")
}
