package help

import (
	"github.com/yuk7/wsldl/backup"
	"github.com/yuk7/wsldl/clean"
	"github.com/yuk7/wsldl/config"
	"github.com/yuk7/wsldl/get"
	"github.com/yuk7/wsldl/run"
	"github.com/yuk7/wsllib-go"
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
