package main

import (
	"log"
	"os"
	"flag"
	"path/filepath"
	"strings"

	"github.com/mitchellh/cli"

	"github.com/albatiqy/gopoh/pkg/lib/fs"
	"github.com/albatiqy/gopoh/pkg/lib/env"

	"github.com/albatiqy/gopoh-app-dbimport/internal"
)

var (
	cfg *internal.Config
	app = cli.NewCLI("app", "1.0.0")
	cmds = map[string]cli.CommandFactory{}
	workingDir string
	ui = &cli.BasicUi{Writer: os.Stdout}
)

func main() {
	var (
		pathEnv string
		envDir  string
	)

	workingDir = fs.WorkingDir()

	flag.StringVar(&pathEnv, "e", "", "lokasi .env, default direktori aktif")
	flag.Parse()

	if pathEnv != "" {
		if strings.HasPrefix(pathEnv, ".") {
			pathEnv = filepath.Join(workingDir, pathEnv)
		}
		if fs.FileInfo(pathEnv) == nil {
			log.Fatal(`file ".env" tidak ditemukan`)
		}
		env.Load(pathEnv)
		envDir = filepath.Dir(pathEnv)
	} else {
		env.Load()
		envDir = workingDir
	}

	var err error
	if cfg, err = internal.NewConfig(envDir); err != nil {
		log.Fatalf(err.Error())
	}

	app.Args = flag.Args()
	app.Commands = cmds
	// app.Autocomplete = true

	exitStatus, err := app.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}