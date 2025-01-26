package list

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/sqc157400661/kdb-sidecar/cmd/mysqlctl/root"
	"github.com/sqc157400661/kdb-sidecar/internal/biz"
	"github.com/sqc157400661/kdb-sidecar/internal/types"
	"github.com/sqc157400661/kdb-sidecar/pkg/output"
	"github.com/sqc157400661/util"
)

type Option struct {
	types.InstancesReq
	// output format
	Format string
}

func NewListCommand() *cobra.Command {
	var opt Option
	var cmd = &cobra.Command{
		Use:   "list",
		Short: "list information of instances",
		Long:  "view the list information of instances",
		Run: func(cmd *cobra.Command, args []string) {
			data, err := biz.ListInstance(opt.InstancesReq)
			if err != nil {
				util.PrintFatalError(errors.Wrap(err, "list instances error"))
			}
			output.FormatOutToStdout(data, opt.Format)
		},
	}
	cmd.Flags().StringVarP(&opt.Format, "format", "f", "", "set output format,support json and table format")
	cmd.Flags().StringVarP(&opt.Status, "status", "s", "", "filter output based on status")
	return cmd
}

func init() {
	root.AddChildCommand(NewListCommand())
}
