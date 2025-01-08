package list

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/sqc157400661/kdb-sidecar/cmd/mysqlctl/root"
)

// command1Cmd 是命令1
var list = &cobra.Command{
	Use:   "list",
	Short: "This is list",
	Long:  "Command1 does something interesting in this CLI.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Executing list!")
	},
}

func init() {
	root.AddChildCommand(list)
}
