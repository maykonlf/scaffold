package main

import (
	"github.com/maykonlf/scaffold/internal/command"
	"github.com/maykonlf/scaffold/internal/usecase"
)

func main() {
	templateUseCase, err := usecase.NewTemplaceUseCase()
	if err != nil {
		panic(err)
	}

	rootCmd := command.NewRootCommand(command.NewTemplateCommand(templateUseCase))
	rootCmd.Execute()
}
