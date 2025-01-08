package root

import (
	"fmt"
	"github.com/spf13/cobra"
)

// rootCmd 是根命令
var rootCmd = &cobra.Command{
	Use:   "mycli",
	Short: "My CLI application",
	Long:  "This is a CLI application for various tasks",
	Run: func(cmd *cobra.Command, args []string) {
		// 当没有指定子命令时，执行此命令
		fmt.Println("Welcome to mycli!")
	},
}

// Execute 运行根命令
func Execute() error {
	return rootCmd.Execute()
}

func AddChildCommand(cmd *cobra.Command) {
	rootCmd.AddCommand(cmd)
}
