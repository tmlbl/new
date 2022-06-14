package new

import (
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/tmlbl/new/ui"
	"gopkg.in/yaml.v2"
)

type Template struct {
	Version     string        `yaml:"version"`
	Name        string        `yaml:"name"`
	Description string        `yaml:"description"`
	Vars        []TemplateVar `yaml:"vars"`
}

func NewBlankTemplate() Template {
	return Template{
		Version:     "0.0.1",
		Name:        "my-new-template",
		Description: "A blank template for you to use",
	}
}

type TemplateVar struct {
	Name    string   `yaml:"name"`
	Type    string   `yaml:"type"`
	Prompt  string   `yaml:"prompt"`
	Choices []string `yaml:"choices"`
}

// Instance represents the contextual information for rendering a template
type Instance struct {
	spec      Template
	values    map[string]string
	sourceDir string
}

// NewInstance parses the template.yaml in the source directory and initializes
// an Instance
func NewInstance(sourceDir string) (*Instance, error) {
	data, err := ioutil.ReadFile(filepath.Join(sourceDir, "template.yaml"))
	if err != nil {
		return nil, err
	}
	spec := Template{}
	err = yaml.Unmarshal(data, &spec)
	if err != nil {
		return nil, err
	}
	return &Instance{
		spec:      spec,
		values:    make(map[string]string),
		sourceDir: sourceDir,
	}, nil
}

// Prompt displays the interactive prompts in the terminal and gathers
// user-supplied values
func (i *Instance) Prompt() error {
	for _, v := range i.spec.Vars {
		// Display choice prompt if there are choices
		if len(v.Choices) > 0 {
			ix, err := ui.NewChoicesPrompt(v.Prompt, v.Choices)
			if err != nil {
				return err
			}
			i.values[v.Name] = v.Choices[ix]
		} else {
			input, err := ui.NewInputPrompt(v.Prompt)
			if err != nil {
				return err
			}
			i.values[v.Name] = input
		}
	}
	return nil
}

// Execute creates the templated files in the target directory
func (i *Instance) Execute(dest string) error {
	return filepath.Walk(i.sourceDir,
		func(path string, info fs.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			if err != nil {
				return err
			}
			tpl, err := template.ParseFiles(path)
			if err != nil {
				return err
			}
			targetPath := strings.TrimPrefix(path, i.sourceDir)
			targetDir := filepath.Dir(targetPath)
			err = os.MkdirAll(targetDir, 0755)
			if err != nil {
				return err
			}
			targetFile, err := os.Create(targetPath)
			if err != nil {
				return err
			}
			return tpl.Execute(targetFile, i.values)
		})
}
