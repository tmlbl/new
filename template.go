package new

import (
	"io/ioutil"
	"path/filepath"

	"github.com/tmlbl/new/ui"
	"gopkg.in/yaml.v2"
)

type Template struct {
	Version string        `yaml:"version"`
	Name    string        `yaml:"name"`
	Vars    []TemplateVar `yaml:"vars"`
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
		if len(v.Choices) > 0 {
			ix, err := ui.NewChoicesPrompt(v.Prompt, v.Choices)
			if err != nil {
				return err
			}
			i.values[v.Name] = v.Choices[ix]
		}
	}
	return nil
}
