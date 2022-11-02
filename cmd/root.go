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

var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Manage template repositories",
}

var repoAddCmd = &cobra.Command{
	Use:   "add [uri]",
	Short: "Add a template repository to the local database",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("provide at least one argument")
		}
		for _, a := range args {
			err := new.AddRepo(a)
			if err != nil {
				return err
			}
		}
		return nil
	},
}

func Execute() {
	repoCmd.AddCommand(repoAddCmd)
	rootCmd.AddCommand(repoCmd)

	rootCmd.PersistentFlags().String("path", "", "The path...")

	rootCmd.AddCommand(templateCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
