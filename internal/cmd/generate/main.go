package main

import (
	"github.com/mporracindie/go-elasticsearch/v6/internal/cmd/generate/commands"
	_ "github.com/mporracindie/go-elasticsearch/v6/internal/cmd/generate/commands/gensource"
	_ "github.com/mporracindie/go-elasticsearch/v6/internal/cmd/generate/commands/genstruct"
	_ "github.com/mporracindie/go-elasticsearch/v6/internal/cmd/generate/commands/gentests"
)

func main() {
	commands.Execute()
}
