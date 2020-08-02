package usecase

import (
	"fmt"
	"io/ioutil"
	"os"
	"text/tabwriter"

	"github.com/maykonlf/scaffold/internal/entities"
	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v2"
)

func NewTemplaceUseCase() (TemplateUseCaseI, error) {
	dir, err := homedir.Dir()
	if err != nil {
		return nil, err
	}

	configPath := fmt.Sprintf("%s/.scaffold", dir)

	u := &TemplateUseCase{
		configPath: configPath,
	}

	return u, u.init()
}

type TemplateUseCaseI interface {
	List()
	Add(name, source string) error
}

type TemplateUseCase struct {
	configPath string
	config     entities.Config
}

func (u *TemplateUseCase) init() error {
	_ = os.Mkdir(u.configPath, 0770)

	data, err := ioutil.ReadFile(u.configPath + "/config")
	if err != nil {
		return err
	}

	return yaml.Unmarshal(data, &u.config)
}

func (u *TemplateUseCase) List() {
	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)

	fmt.Fprintln(writer, "name\tsource")
	for _, v := range u.config.Templates {
		fmt.Fprintf(writer, "%s\t%s\n", v.Name, v.Source)
	}

	writer.Flush()
}

func (u *TemplateUseCase) Add(name, source string) error {
	u.config.Templates = append(u.config.Templates, &entities.TemplateSource{
		Name:   name,
		Source: source,
	})

	data, err := yaml.Marshal(u.config)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(u.configPath+"/config", data, 0600)
}
