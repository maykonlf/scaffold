package console

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

type Service interface {
	WriteTable(rows [][]string)
}

func NewService() Service {
	return &service{
		writer: tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight),
	}
}

type service struct {
	writer *tabwriter.Writer
}

func (s *service) WriteTable(rows [][]string) {
	for _, row := range rows {
		fmt.Fprintln(s.writer, strings.Join(row, "\t"))
	}

	s.writer.Flush()
}
