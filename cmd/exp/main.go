package main

import (
	"fmt"
	"os"

	"github.com/xxnmxx/csv/util"
)

func main() {
	t := util.ImportCsv(os.Stdin, ',', true)
	//t.ExportCsv(os.Stdout)
	t.Print(os.Stdout)
	fmt.Println(t)
}
