package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tmlbl/new"
)

var rootCmd = &cobra.Command{
	Use:   "new",
	Short: "A joyful project templating tool",
	RunE: func(cmd *cobra.Command, args []string) error {
		path := cmd.Flag("path").Value.String()
		wd, err := os.Getwd()
		if err != nil {
			return err
		}
		if path == "" {
			path = wd
		}
		inst, err := new.NewInstance(path)
		if err != nil {
			return err
		}
		err = inst.Prompt()
		if err != nil {
			return err
		}
		return inst.Execute(wd)
	},
}

func Execute() {
	rootCmd.PersistentFlags().String("path", "", "The path...")

	rootCmd.AddCommand(templateCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
