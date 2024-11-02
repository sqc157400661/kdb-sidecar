package main

import (
	"flag"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/sqc157400661/kdb-sidecar/cmd/sidecar/mysql"
	"k8s.io/klog/v2"
)

func init() {
	fs := flag.FlagSet{}
	klog.InitFlags(&fs)
	fs.Set("v", "0")
	fs.Set("logtostderr", "false")
	fs.Set("log_file", "/tmp/cli.log")
	fs.Set("stderrthreshold", "3")

	klog.Info("start cli")

}
func main() {
	defer klog.Flush()
	rootCmd := cobra.Command{
		Use:              "sidecar",
		Short:            "sidecar service",
		Long:             `sidecar service`,
		SilenceUsage:     true,
		Args:             cobra.NoArgs,
		TraverseChildren: true,
		Version:          fmt.Sprintf("%#v", "v0.0.1"),
	}
	// 把这两条命令加入到根命令里面
	rootCmd.AddCommand(mysql.NewMySQLSidecarServerCmd())
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Could not run command")
	}
}
