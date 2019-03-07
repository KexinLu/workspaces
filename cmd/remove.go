package cmd

import (
	"github.com/spf13/cobra"
	"workspaces/logging"
	"github.com/pkg/errors"
	"os"
)

var (
	removeLogger logging.LoggableEntity

	removeCmd = &cobra.Command{
		Use: "remove",
		Short: "remove project from registry",
		Long: `remove project from registry`,
		Args: func(cmd *cobra.Command, args []string) error {
			initConfig()
			if len(args) == 1 {
				removeProject(args[0])
			}
			if len(args) == 0 {
				pName := selectProject()
				removeProject(pName)
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {},
	}
)

func init() {
	removeLogger = logging.NewLoggableEntity("remove", logging.Fields{ "module": "remove" })
	removeLogger.Debug("initializing remove logger")
}

func removeProject(name string) {
	removeLogger.DebugWithFields(logging.Fields{"proj": name}, "removing project")
	if !cfg.HasProject(name) {
		removeLogger.Error(errors.New("project does not exist"), "can not find project with name " + name)
		os.Exit(1)
	}
	q :=  "Remove project? (y/n) " + name
	if rmv, err := getYesNoResponse(q); err != nil {
		removeLogger.Error(err, "Failed to get a y/n response")
		os.Exit(1)
	} else if rmv {
		if err := cfg.RemoveProject(name); err != nil {
			removeLogger.Error(err, "Failed to remove project")
			os.Exit(1)
		}
	}

	printToConfig()
}
