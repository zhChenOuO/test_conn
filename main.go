package main

import (
	"fmt"
	"os"

	"test_conn/cmd"

	cobra "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "server client"}

func main() {
	rootCmd.AddCommand(cmd.ServerCmd, cmd.ClientCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
