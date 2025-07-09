package main

import (
	"compose-parser-component/internal/compose-parser/parser/parse"
	"context"
	"log"

	"github.com/compose-spec/compose-go/v2/cli"
)

func init() {
	parse.Exports.Parse = func() {
		println("Hello from go wasm component")
		composeFilePath := "docker-compose.yml"
		projectName := "testproject"
		ctx := context.Background()

		options, err := cli.NewProjectOptions(
			[]string{composeFilePath},
			cli.WithName(projectName),
		)
		if err != nil {
			log.Fatal(err)
		}

		project, err := options.LoadProject(ctx)
		if err != nil {
			log.Fatal(err)
		}

		projectJSON, err := project.MarshalJSON()
		if err != nil {
			log.Fatal(err)
		}

		println(string(projectJSON))
	}
}

// main is required for the `wasi` target, even if it isn't used.
func main() {}
