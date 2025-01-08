package stop

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/sqc157400661/kdb-sidecar/cmd/mysqlctl/root"
)

var stop = &cobra.Command{
	Use:   "stop",
	Short: "This is stop",
	Long:  "Command1 does something interesting in this CLI.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Executing stop")
	},
}

func init() {
	root.AddChildCommand(stop)
}
