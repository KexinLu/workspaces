package cmd

import (
	"github.com/spf13/cobra"
	"workspaces/logging"
	"github.com/manifoldco/promptui"
	"github.com/pkg/errors"
	"os"
	"fmt"
)

var (
	aliasLogger logging.LoggableEntity

	aliasCmd = &cobra.Command{
		Use: "alias",
		Short: "Set alias to project",
		Long: `Set alias to project`,
		Args: func(cmd *cobra.Command, args []string) error {
			aliasLogger.Debug("aliasing")
			initConfig()

			aliasLogger.Debug(fmt.Sprintf("got %v arguments", len(args)))
			if len(args) == 2 {
				cfg.AliasProject(args[0], args[1])
			} else if len(args) == 1 {
				aliasProject(args[0])
				return nil
			} else {
				pName := selectProject()
				aliasProject(pName)
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {},
	}
)

func init() {
	aliasLogger = logging.NewLoggableEntity("alias", logging.Fields{ "module": "alias" })
	aliasLogger.Debug("initializing alias logger")
}

func aliasProject(name string) {
	aliasLogger.DebugWithFields(logging.Fields{"proj": name}, "aliasing project")
	if !cfg.HasProject(name) {
		aliasLogger.Error(errors.New("project does not exist"), "can not find project with name " + name)
		os.Exit(1)
	}

	pa := promptui.Prompt{
		Label:    "Enter the alias",
		Validate: func(s string) error {
			if len(s) < 1 {
				return errors.New("empty alias")
			} else if cfg.HasProject(s) {
				return errors.New("alias occupied by name or alias")
			}

			return nil
		},
	}

	if a, err := pa.Run(); err != nil {
		aliasLogger.Error(err, "Failed to alias the project")
		os.Exit(1)
	} else {
		if err := cfg.AliasProject(name, a); err != nil {
			aliasLogger.Error(err, "Failed to alias project")
			os.Exit(1)
		}
	}

	printToConfig()
}
