package help

import (
	"github.com/sarvex/wsldl/backup"
	"github.com/sarvex/wsldl/clean"
	"github.com/sarvex/wsldl/config"
	"github.com/sarvex/wsldl/get"
	"github.com/sarvex/wsldl/install"
	"github.com/sarvex/wsldl/run"
)

// ShowHelpAll shows all help messages
func ShowHelpAll(registered bool) {
	println("Usage :")
	if registered {
		run.ShowHelp(false)
		println()
		config.ShowHelp(false)
		println()
		get.ShowHelp(false)
		println()
		backup.ShowHelp(false)
		println()
		clean.ShowHelp(false)
		println()
	} else {
		install.ShowHelp(false)
		println()
	}

	ShowHelp(false)
}

// ShowHelp shows help message
func ShowHelp(showTitle bool) {
	if showTitle {
		println("Usage:")
	}
	println("    help")
	println("      - Print this usage message.")
}
