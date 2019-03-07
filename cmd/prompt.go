package cmd

import (
	"github.com/manifoldco/promptui"
	. "workspaces/config_model"
	"errors"
	"workspaces/logging"
	"os"
)

var pLog logging.LoggableEntity
func init() {
	pLog = logging.NewLoggableEntity("prompt", logging.Fields{"module": "prompt"})
}

func buildProjsPrompt() (promptui.Select, *[]Project) {
	pLog.Debug("building project selection prompt")
	var opts []Project
	for _, v := range cfg.Projects {
		opts = append(opts, v)
	}
	spTemplate := promptui.SelectTemplates{
		Active:   ` {{ .Name | green | bold }} ({{ .Alias | red | bold }})`,
		Inactive: ` {{ .Name | cyan }} ({{ .Alias | cyan }})`,
		Selected: ` {{ "✔" | green | bold }} {{ .Name | cyan }}`,
		Details:  ` 
Detail: 
Project: {{ .Name }}
Alias: {{.Alias}}
Path: {{.Path}}
IsGit: {{.IsGit}}
`,
	}
	sp := promptui.Select{
		Label: "Select project",
		Items: opts,
		Templates: &spTemplate,
	}
	pLog.Debug("finished building project selection prompt")

	return sp, &opts
}

func getYesNoResponse(question string) (bool, error) {
	p := promptui.Prompt{
		Label: question,
		Validate: func(s string) error {
			allowed := map[string]interface{}{"y": 0, "n": 0, "Y": 0, "N": 0}
			if _, ok := allowed[s]; !ok {
				return errors.New("only y/n allowed")
			}
			return nil
		},
		Default: "n",
	}
	if result, err := p.Run(); err != nil {
		return false, err
	} else {
		if result == "y" || result == "Y" {
			return true, nil
		}
		return false, nil
	}
}

func selectProject() string {
	s, opts := buildProjsPrompt()
	if i, _, err := s.Run(); err != nil {
		pLog.Error(err,"failed to pick project")
		os.Exit(1)
	} else {
		return (*opts)[i].Name
	}
	return ""
}

