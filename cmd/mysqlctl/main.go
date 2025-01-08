package main

import (
	"fmt"
	root "github.com/sqc157400661/kdb-sidecar/cmd/mysqlctl/root"
	"os"
)

func main() {
	// 初始化根命令
	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
