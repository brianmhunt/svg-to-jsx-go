package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const commentRex = "<!--.*-->"

// Read the SVG, stripping any HTML/SVG <!-- ... --> comments
func pureSvg(svgPath string) string {
	r := regexp.MustCompile(commentRex)
	body, err := ioutil.ReadFile(svgPath)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	s := r.ReplaceAllString(string(body), "")
	s = strings.ReplaceAll(s, "{", "{'{")
	s = strings.ReplaceAll(s, "}", "}'}")
	return s
}

func svgToJsx(dir, outdir, filename string, verbose bool) {
	basename := strings.TrimSuffix(filename, filepath.Ext(filename))
	outFilename := basename + ".tsx"
	if verbose {
		fmt.Printf("[%s] %s => %s/%s\n", dir, filename, outdir, outFilename)
	}

	tf, err := os.Create(filepath.Join(outdir, outFilename))
	if err != nil {
		log.Fatal(err)
	}
	defer tf.Close()

	tsx := `export default ` + pureSvg(filepath.Join(dir, filename))

	_, err = tf.WriteString(tsx)
	if err != nil {
		log.Fatal(err)
	}
}

func convertSvgsInDir(dir, outdir string, verbose bool) {
	files, err := ioutil.ReadDir(dir)
	os.MkdirAll(outdir, 0772)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".svg" {
			svgToJsx(dir, outdir, file.Name(), verbose)
		}
	}
}

func main() {
	outdir := flag.String("o", ".", "Output directory")
	verbose := flag.Bool("v", false, "Verbose operation.")

	flag.Parse()

	// log.Printf("Writing to %s (verbose: %v)", *outdir, *verbose)
	args := flag.Args()

	convertSvgsInDir(args[0], *outdir, *verbose)
}
