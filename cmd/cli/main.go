package main

import (
	"context"
	"fmt"
	"os"

	app "github.com/dugajean/goke/internal"
)

func main() {
	opts := app.NewCliOptions()

	cfg, err := app.ReadYamlConfig()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Wrappers
	fs := app.LocalFileSystem{}
	proc := app.ShellProcess{}

	// Main components
	p := app.NewParser(cfg, &opts, &fs)
	p.Bootstrap()

	// Global options
	handleGlobalOptions(&opts, &p)

	l := app.NewLockfile(p.GetFilePaths(), &opts, &fs)
	l.Bootstrap()

	ctx := context.Background()
	e := app.NewExecutor(&p, &l, &opts, &proc, &fs, &ctx)
	e.Start(opts.TaskName)
}
