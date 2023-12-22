package internal

import (
	"fmt"
	"strconv"
	"time"

	"github.com/jetrails/proposal-nginx/pkg/vhost"
	"github.com/spf13/cobra"
)

var rollbackCmd = &cobra.Command{
	Use:   "rollback SITE_NAME REVISION",
	Short: "rollback to a previous checkpoint",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		siteName := args[0]
		revision := args[1]

		revisionInt, errRevision := strconv.Atoi(revision)
		if errRevision != nil {
			ExitWithError(1, fmt.Sprintf("revision %q is not an integer", revision))
		}
		if !vhost.SiteExists(siteName) {
			ExitWithError(2, fmt.Sprintf("site %q does not exist", siteName))
		}

		checkPoint, errCheckPoint := vhost.GetCheckPoint(siteName, revisionInt)
		if errCheckPoint != nil {
			ExitWithError(3, errCheckPoint.Error())
		}

		latestCheckPoint, errLatest := vhost.GetLatestCheckPoint(siteName)
		if errLatest != nil {
			ExitWithError(4, errLatest.Error())
		}

		if errDelete := latestCheckPoint.Output.DeleteFiles(true); errDelete != nil {
			ExitWithError(5, errDelete.Error())
		}

		checkPoint.Revision = latestCheckPoint.Revision + 1
		checkPoint.Description = fmt.Sprintf("rollback to revision %d", revisionInt)
		checkPoint.Timestamp = time.Now()

		errOutputSave := checkPoint.Output.Save()
		if errOutputSave != nil {
			ExitWithError(6, errOutputSave.Error())
		}

		errSave := checkPoint.Save()
		if errSave != nil {
			ExitWithError(7, errSave.Error())
		}
	},
}

func init() {
	RootCmd.AddCommand(rollbackCmd)
}