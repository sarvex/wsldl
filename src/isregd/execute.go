package isregd

import (
	"os"

	"github.com/sarvex/wsllib-go"
)

//Execute is default isregd entrypoint. Exits with registerd status
func Execute(name string) {

	if wsllib.WslIsDistributionRegistered(name) {
		os.Exit(0)
	}
	os.Exit(1)
}
