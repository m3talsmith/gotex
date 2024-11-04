package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"text/template"
)

var (
	name       string
	outputPath string
)

func init() {
	flag.StringVar(&name, "name", "world", "Your name")
	flag.StringVar(&outputPath, "out", ".", "The path to output your pdf at")
	flag.Parse()
}

func main() {
	var buff bytes.Buffer
	tmpl, _ := template.New("doc").Parse(string(tmplFile))
	err := tmpl.Execute(&buff, struct{ Name string }{Name: name})
	if err != nil {
		panic(err)
	}
	latexFile := fmt.Sprintf("%s/doc.tex", os.TempDir())
	os.WriteFile(latexFile, buff.Bytes(), os.ModePerm)
	outputFile := fmt.Sprintf("%s/doc.pdf", outputPath)
	cmd := exec.Command("pandoc", "-s", "-f", "latex", "-t", "pdf", "-o", outputFile, latexFile)
	err = cmd.Run()
	if err != nil {
		fmt.Printf("runtime error: %s", cmd.Stderr)
		panic(err)
	}
}
