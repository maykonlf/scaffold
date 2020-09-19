package configs

import (
	"io/ioutil"
	"os"

	"github.com/maykonlf/scaffold/internal/domain/entities"
	"gopkg.in/yaml.v2"
)

type Service interface {
	GetTemplates() ([]*entities.Template, error)
	AddTemplate(template *entities.Template) error
}

func NewService(configPath string) Service {
	return &service{
		configPath:     configPath,
		configFilePath: configPath + "/config",
	}
}

type service struct {
	configPath     string
	configFilePath string
}

func (s *service) GetTemplates() ([]*entities.Template, error) {
	config, err := s.readConfig()
	if err != nil {
		return nil, err
	}

	templates := make([]*entities.Template, len(config.Templates))
	for i := range templates {
		templates[i] = config.Templates[i].ToEntity()
	}

	return templates, err
}

func (s *service) AddTemplate(template *entities.Template) error {
	config, err := s.readConfig()
	if err != nil {
		return err
	}

	config.Templates = append(config.Templates, &Template{
		Name:   template.Name,
		Source: template.Source,
	})

	return s.saveConfig(config)
}

func (s *service) readConfig() (*Config, error) {
	data, err := s.readFile()
	if err != nil {
		return nil, err
	}

	var config Config
	return &config, yaml.Unmarshal(data, &config)
}

func (s *service) readFile() ([]byte, error) {
	if err := s.createConfigFileIfNotExists(); err != nil {
		return nil, err
	}

	return ioutil.ReadFile(s.configPath + "/config")
}

func (s *service) createConfigFileIfNotExists() error {
	_ = os.Mkdir(s.configPath, 0770)
	_, err := os.Stat(s.configFilePath)
	if os.IsNotExist(err) {
		file, err := os.Create(s.configFilePath)
		if err != nil {
			return err
		}
		file.Close()
	}

	return nil
}

func (s *service) saveConfig(config *Config) error {
	data, _ := yaml.Marshal(config)
	return ioutil.WriteFile(s.configFilePath, data, 0600)
}
