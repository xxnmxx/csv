package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/xxnmxx/csv/util"
)

var (
	del  string
	lazy bool
)

func argsToInt(args []string) []int {
	intArgs := make([]int, len(args))
	for i, arg := range args {
		intArg, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "args must be int. %v\n", err)
			return nil
		}
		intArgs[i] = intArg
	}
	return intArgs
}

func main() {
	flag.StringVar(&del, "d", ",", "delimiter")
	flag.BoolVar(&lazy, "l", false, "lazy quates")
	flag.Parse()
	args := flag.Args()
	d := []rune(del)
	t := util.ImportCsv(os.Stdin, d[0], lazy)
	c := argsToInt(args)
	nt := t.Cut(c...)
	nt.ExportCsv(os.Stdout)
}
