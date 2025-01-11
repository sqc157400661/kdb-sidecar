package stop

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/sqc157400661/kdb-sidecar/cmd/mysqlctl/root"
)

var resume = &cobra.Command{
	Use:   "resume",
	Short: "This is resume",
	Long:  "Command1 does something interesting in this CLI.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Executing stop")
	},
}

func init() {
	root.AddChildCommand(resume)
}
