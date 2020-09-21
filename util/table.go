package util

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
)

type Table [][]string

func NewTable(r, c int) Table {
	t := make(Table, r)
	for i := range t {
		t[i] = make([]string, c)
	}
	return t
}

func (t Table) Shape() (r, c int) {
	return len(t), len(t[0])
}

func ImportCsv(in io.Reader, d rune, l bool) Table {
	recs := make(Table, 0)
	r := csv.NewReader(in)
	r.Comma = d
	r.LazyQuotes = l
	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("import error: %v", err)
		}
		recs = append(recs, rec)
	}
	return recs
}

func (t Table) ExportCsv(out io.Writer) {
	w := csv.NewWriter(out)
	for _, rec := range t {
		if err := w.Write(rec); err != nil {
			log.Fatalf("writing error: %v", err)
		}
	}
	w.Flush()
	if err := w.Error(); err != nil {
		log.Fatalf("error on export: %v", err)
	}
}

func (t Table) Cut(c ...int) Table {
	nt := NewTable(len(t), 0)
	ruler := t.ruler(c...)
	for i := range t {
		for j, ok := range ruler {
			if ok {
				continue
			} else {
				nt[i] = append(nt[i], t[i][j])
			}
		}
	}
	return nt
}

func (t Table) Select(c ...int) Table {
	nt := NewTable(len(t), 0)
	ruler := t.ruler(c...)
	for i := range t {
		for j, ok := range ruler {
			if !ok {
				continue
			} else {
				nt[i] = append(nt[i], t[i][j])
			}
		}
	}
	return nt
}

func (t Table) ruler(c ...int) []bool {
	ruler := make([]bool, len(t[0]))
	for _, v := range c {
		ruler[v] = true
	}
	return ruler
}

func (t Table) Info() {
	fmt.Fprintf(os.Stdout, "rows: %v\n", len(t))
	fmt.Fprintf(os.Stdout, "columns: %v\n", len(t[0]))
	fmt.Fprintf(os.Stdout, "headers: %v\n", t.ListOfHeader())
}

func (t Table) Print(out io.Writer) {
	w := tabwriter.NewWriter(out, 0, 2, '\t', 0,0)
	for _, v := range t {
		fmt.Fprint(w, v)
	}
	w.Flush()
}

func strSliceToStr(s []string) string {
	var b strings.Builder
	for _, v := range s {
		b.WriteString(v + " ")
	}
	return b.String()
}

func (t Table) Header() []string {
	h := make([]string, len(t[0]))
	for i, v := range t[0] {
		h[i] = v
	}
	return h
}

func (t Table) ListOfHeader() string {
	h := ""
	for i, v := range t[0] {
		idx := strconv.Itoa(i)
		h += "[" + idx + "]" + v + " "
	}
	return h
}

//func (t Table) Unique(c int) []string {
//	m := make(map[string]int)
//	for i, v := range t {
//
//	}
//}
