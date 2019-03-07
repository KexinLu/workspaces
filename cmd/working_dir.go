package cmd

import (
	"github.com/spf13/cobra"
	"workspaces/logging"
	. "workspaces/config_model"
	"github.com/pkg/errors"
	"os"
	"fmt"
)

var (
	wdLog logging.LoggableEntity

	wdCmd = &cobra.Command{
		Use: "wd",
		Short: `show project path`,
		Long: `show absolute path of project`,
		Run: func(cmd *cobra.Command, args []string) {},
		Args: func(cmd *cobra.Command, args []string) error {
			initConfig()
			if len(args) > 0 {
				noa := args[0]
				if !cfg.HasProject(noa) {
					err := errors.New("project not managed")
					wdLog.Fatal(err.Error())
					return err
				}
				p := cfg.GetProject(noa)
				wd(*p)
			} else {
				s, opts := buildProjsPrompt()
				if i, _, err := s.Run(); err != nil {
					err := errors.New("failed to pick project name")
					wdLog.Fatal(err.Error())
				} else if i < 0{
					wdLog.Info("No project was selected, exiting")
					os.Exit(0)
				} else{
					wdLog.Debug("moving")
					wd((*opts)[i])
				}
			}
			return nil
		},
	}
)

func init() {
	wdLog = logging.NewLoggableEntity("working_directory", logging.Fields{ "module": "working_directory" })
	pickLogger.Debug("initializing working directory logger")
}

func wd(p Project) {
	wdLog.Debug("showing path: " + p.Path)
	fmt.Fprint(os.Stdout, fmt.Sprintf("%s\n", p.Path))
}
