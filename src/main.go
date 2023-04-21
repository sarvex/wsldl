package main

import (
	"os"
	"path/filepath"

	"github.com/sarvex/wsldl/backup"
	"github.com/sarvex/wsldl/clean"
	"github.com/sarvex/wsldl/config"
	"github.com/sarvex/wsldl/get"
	"github.com/sarvex/wsldl/help"
	"github.com/sarvex/wsldl/install"
	"github.com/sarvex/wsldl/isregd"
	"github.com/sarvex/wsldl/lib/utils"
	"github.com/sarvex/wsldl/run"
	"github.com/sarvex/wsldl/version"
	"github.com/sarvex/wsllib-go"
)

func main() {
	efPath, _ := os.Executable()
	name := filepath.Base(efPath[:len(efPath)-len(filepath.Ext(efPath))])

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "version", "-v", "--version":
			version.Execute()

		case "isregd":
			isregd.Execute(name)

		case "install":
			install.Execute(name, os.Args[2:])

		case "run", "-c", "/c":
			run.Execute(name, os.Args[2:])

		case "runp", "-p", "/p":
			run.ExecuteP(name, os.Args[2:])

		case "config", "set":
			config.Execute(name, os.Args[2:])

		case "get":
			get.Execute(name, os.Args[2:])

		case "backup":
			backup.Execute(name, os.Args[2:])

		case "clean":
			clean.Execute(name, os.Args[2:])

		case "help", "-h", "--help", "/?":
			help.Execute(name, os.Args[2:])

		default:
			utils.ErrorExit(os.ErrInvalid, true, true, false)
		}
	} else {
		if !wsllib.WslIsDistributionRegistered(name) {
			install.Execute(name, nil)
		} else {
			run.ExecuteNoArgs(name)
		}
	}
}
