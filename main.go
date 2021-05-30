package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"text/template"

	"golang.org/x/tools/imports"
)

func main() {
	in := flag.String("in", "", "input file, e.g. in.tpl")
	out := flag.String("out", "", "output file, e.g. out.go")
	js := flag.String("json", "", "data file, e.g. data.json")
	dryRun := flag.Bool("dry-run", false, "print to stdout without creating files")
	flag.Parse()

	if err := exec(*in, *out, *js, *dryRun); err != nil {
		panic(err)
	}
	if !*dryRun {
		fmt.Printf("wrote file to %s\n", *out)
	}
}

func exec(in, out, js string, dryRun bool) error {
	// Get the current working directory of the caller.
	dirpath := filepath.Dir(in)

	in = filepath.Join(dirpath, in)
	out = filepath.Join(dirpath, out)
	js = filepath.Join(dirpath, js)

	b, err := os.ReadFile(in)
	if err != nil {
		return fmt.Errorf("failed to read in file: %w", err)
	}

	var data map[string]interface{}
	{
		b, err := os.ReadFile(js)
		if err != nil {

			return fmt.Errorf("failed to read json data: %w", err)
		}
		if err := json.Unmarshal(b, &data); err != nil {
			return fmt.Errorf("failed to unmarshal json data: %w", err)
		}
	}

	t := template.Must(template.New("").Parse(string(b)))

	var bb bytes.Buffer
	if err := t.Execute(&bb, data); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}
	result := bb.Bytes()

	// Additional formatting as well as imports for .go files.
	if filepath.Ext(out) == ".go" {
		result, err = format(result)
		if err != nil {
			return fmt.Errorf("failed to format go file: %w", err)
		}
	}

	if dryRun {
		// Print to stdout.
		fmt.Println(string(result))
	} else {
		// Create the output directory.
		if err := os.MkdirAll(path.Dir(out), os.ModePerm); err != nil {
			return fmt.Errorf("failed to create output directory: %w", err)
		}

		// Create the output file. This will overwrite existing file content.
		f, err := os.OpenFile(out, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			return fmt.Errorf("failed to open out file: %w", err)
		}
		defer f.Close()

		// Write to the output file.
		_, err = f.Write(result)
		if err != nil {
			return fmt.Errorf("failed to write: %w", err)
		}
	}
	return nil
}

// format is gofmt with addition of removing any unused imports.
func format(source []byte) ([]byte, error) {
	return imports.Process("", source, &imports.Options{
		AllErrors: true, Comments: true, TabIndent: true, TabWidth: 8,
	})
}
