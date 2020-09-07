package presenters

import "github.com/maykonlf/scaffold/internal/domain/entities"

type Template interface {
	TemplatesTable(templates []*entities.Template) *TemplatesTable
}

func NewTemplate() Template {
	return &template{}
}

type template struct{}

func (t template) TemplatesTable(templates []*entities.Template) *TemplatesTable {
	rows := make([][]string, len(templates))

	for i := range templates {
		rows[i] = []string{
			templates[i].Name,
			templates[i].Source,
		}
	}

	return &TemplatesTable{
		Headers: []string{"name", "source"},
		Rows:    rows,
	}
}

type TemplatesTable struct {
	Headers []string
	Rows    [][]string
}
