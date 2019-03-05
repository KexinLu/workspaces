package cmd

import (
	"github.com/spf13/cobra"
	"workspaces/logging"
	"github.com/manifoldco/promptui"
	. "workspaces/config_model"
	"github.com/pkg/errors"
	"os"
	//"bufio"
)

var (
	moveLogger logging.LoggableEntity

	cdCmd = &cobra.Command{
		Use: "cd",
		Long: `change directory to project`,
		Run: func(cmd *cobra.Command, args []string) {},
		Args: func(cmd *cobra.Command, args []string) error {
			initConfig()
			if len(args) > 0 {
				if v, ok := cfg.Projects[args[0]]; !ok {
					err := errors.New("project not managed")
					moveLogger.Fatal(err.Error())
					return err
				} else {
					moveTo(v)
				}
			} else {
				prompt := buildProjsPrompt()
				if _, name, err := prompt.Run(); err != nil {
					err := errors.New("failed to pick project name")
					moveLogger.Fatal(err.Error())
				} else if name == "" {
					moveLogger.Info("No project was selected, exiting")
					os.Exit(0)
				} else{
					moveLogger.Debug("moving")
					moveTo(cfg.Projects[name])
				}
			}
			return nil
		},
	}
)

func init() {
	moveLogger = logging.NewLoggableEntity("move", logging.Fields{ "module": "move" })
	pickLogger.Debug("initializing move logger")
}

func buildProjsPrompt() promptui.Select {
	keys := make([]string, len(cfg.Projects))
	for k := range cfg.Projects {
		keys = append(keys, k)
	}
	return promptui.Select{
		Label: "Select project",
		Items: keys,
	}
}

func moveTo(p Project) {
	moveLogger.Debug("moving to " + p.Path)
	//command := fmt.Sprintf("cd %s", p.Path)
	//fmt.Fprint(os.Stdout, command)
	os.Setenv("PWD", p.Path)
	//w := bufio.NewWriter(os.Stdout)
	//w.WriteString(command)
	//w.WriteString("\n")
	//w.Flush()
}

