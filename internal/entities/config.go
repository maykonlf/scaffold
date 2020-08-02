package entities

type Config struct {
	Templates []*TemplateSource `yaml:"templates"`
}

type TemplateSource struct {
	Name   string `yaml:"name"`
	Source string `yaml:"source"`
}
