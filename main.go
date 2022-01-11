package main

import (
	"fmt"
	"io"
	"os"
	"text/template"
)

func main() {
	if len(os.Args) < 3 {
		Fatalf("Input file and output file should be specified.")
	}

	in, err := os.Open(os.Args[1])
	if err != nil {
		Fatalf("Opening %s: %v", os.Args[1], err)
	}
	defer in.Close()

	out, err := os.Create(os.Args[2])
	if err != nil {
		Fatalf("Creating %s: %v", os.Args[2], err)
	}
	defer out.Close()

	if err := expand(out, in); err != nil {
		Fatalf("Expanding: %v", err)
	}
}

var functions = template.FuncMap{
	"makeSlice": func(args ...string) []string {
		return args
	},
}

func expand(w io.Writer, r io.Reader) error {
	bs, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	tpl := template.New("foo")
	tpl.Funcs(functions)

	tpl, err = tpl.Parse(string(bs))
	if err != nil {
		return err
	}

	if err := tpl.Execute(w, nil); err != nil {
		return err
	}

	return nil
}

func Fatalf(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(1)
}
