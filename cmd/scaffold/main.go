package main

import (
	"github.com/maykonlf/scaffold/internal/config"
)

func main() {
	container := config.NewContainer()
	rootCmd := container.GetRootCommand()
	rootCmd.Execute()
}
