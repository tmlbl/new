package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/tmlbl/new"
	"github.com/tmlbl/new/ui"
	"gopkg.in/yaml.v2"
)

var templateCmd = &cobra.Command{
	Use:   "template",
	Short: "Create a new template project",
	RunE: func(cmd *cobra.Command, args []string) error {
		dir, err := os.Getwd()
		if err != nil {
			return err
		}
		current, err := ui.NewYesNoPrompt("Use the current directory?",
			false)
		if err != nil {
			return err
		}
		if !current {
			// Create a directory for the project
			name, err := ui.NewInputPrompt(
				"Enter a name for the project")
			if err != nil {
				return err
			}
			err = os.MkdirAll(name, 0755)
			if err != nil {
				return err
			}
			dir = filepath.Join(dir, name)
		}
		path := filepath.Join(dir, "template.yaml")
		data, err := yaml.Marshal(new.NewBlankTemplate())
		if err != nil {
			return err
		}
		fmt.Println("Writing", path)
		return ioutil.WriteFile(path, data, 0644)
	},
}
