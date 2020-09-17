package console

import (
	"fmt"
	"strings"
)

type Service interface {
	WriteTable(rows [][]string) error
}

func NewService(writer Writer) Service {
	return &service{
		writer: writer,
	}
}

type service struct {
	writer Writer
}

func (s *service) WriteTable(rows [][]string) error {
	for _, row := range rows {
		_, err := fmt.Fprintln(s.writer, strings.Join(row, "\t"))
		if err != nil {
			return err
		}
	}

	return s.writer.Flush()
}
