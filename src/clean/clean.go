package clean

import (
	"fmt"
	"os"

	"github.com/sarvex/wsldl/lib/utils"
	"github.com/sarvex/wsllib-go"
)

//Clean cleans distribution
func Clean(name string, showProgress bool) {
	if showProgress {
		fmt.Println("Unregistering...")
	}
	err := wsllib.WslUnregisterDistribution(name)
	if err != nil {
		utils.ErrorExit(err, showProgress, true, false)
	}
	os.Exit(0)
}
