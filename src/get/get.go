package get

import (
	"github.com/sarvex/wsldl/lib/utils"
	"github.com/sarvex/wsllib-go"
)

//WslGetConfig is getter of distribution configuration
func WslGetConfig(distributionName string) (uid uint64, flags uint32) {
	var err error
	_, uid, flags, err = wsllib.WslGetDistributionConfiguration(distributionName)
	if err != nil {
		utils.ErrorRedPrintln("ERR: Failed to GetDistributionConfiguration")
		utils.ErrorExit(err, true, true, false)
	}
	return
}
