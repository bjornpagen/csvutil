package main

// csvcat.go

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// a quick little script to read csv from stdin and extract columns
// outputs another csv to stdout with the selected columns
func usage() {
	s := fmt.Sprintf("usage: %s <rows>", os.Args[0])
	fmt.Fprint(os.Stderr, s)
}

// argv[1] are the rows to be echoed out, comma separated
func main() {
	if len(os.Args) != 2 {
		usage()
		os.Exit(1)
	}

	rows := strings.Split(os.Args[1], ",")

	r := csv.NewReader(os.Stdin)
	w := csv.NewWriter(os.Stdout)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		var out []string
		for _, row := range rows {
			i, err := strconv.Atoi(row)
			if err != nil {
				log.Fatal(err)
			}
			out = append(out, record[i])
		}

		if err := w.Write(out); err != nil {
			log.Fatal(err)
		}
	}

	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
