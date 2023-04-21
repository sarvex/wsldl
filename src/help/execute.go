package help

import (
	"github.com/sarvex/wsldl/backup"
	"github.com/sarvex/wsldl/clean"
	"github.com/sarvex/wsldl/config"
	"github.com/sarvex/wsldl/get"
	"github.com/sarvex/wsldl/run"
	"github.com/sarvex/wsllib-go"
)

//Execute is default install entrypoint
func Execute(name string, args []string) {
	if len(args) == 0 {
		ShowHelpAll(wsllib.WslIsDistributionRegistered(name))
	} else {
		switch args[0] {
		case "run", "-c", "/c", "runp", "-p", "/p":
			run.ShowHelp(true)
		case "config", "set":
			config.ShowHelp(true)
		case "get":
			get.ShowHelp(true)
		case "backup":
			backup.ShowHelp(true)
		case "clean":
			clean.ShowHelp(true)
		case "help":
			ShowHelp(true)
		default:
			ShowHelpAll(wsllib.WslIsDistributionRegistered(name))
		}
	}

}
