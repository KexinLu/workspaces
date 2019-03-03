package cmd

import (
	"github.com/spf13/cobra"
	"workspaces/logging"
	. "workspaces/config_model"
	"fmt"
	"encoding/json"
)

var (
	listLogger logging.LoggableEntity

	listCmd = &cobra.Command{
		Use: "list",
		Short: "Show all projects managed by workspaces",
		Long: `Show all projects managed by workspaces in the config file`,
		Run: func(cmd *cobra.Command, args []string) {
			conf := Config{}
			if err := hydrateConfig(&conf); err != nil {
				fmt.Println("failed to hydrate config")
			} else {
				printAllProjects(&conf)
			}
		},
	}
)

func init() {
	cobra.OnInitialize(initConfig)
	listLogger.Debug("initializing list logger")
	listLogger = logging.NewLoggableEntity("list", logging.Fields{ "module": "list" })
}

func printAllProjects(c *Config)  {
	listLogger.Debug("Printing all projects")
	listLogger.Debug(fmt.Sprintf("We have %v projects", len(c.Projects)))
	for _, p := range c.Projects {
		listLogger.Debugf(logging.Fields{"name": p.Name, "path": p.Path }, "project")
		if jsb, err := json.MarshalIndent(p, "", "  "); err != nil {
			listLogger.ErrorWithFields(
				logging.Fields{
					"error": err.Error(),
					"name": p.Name,
				},
				err,"failed to marshal project")
		} else {
			fmt.Println(string(jsb))
		}
	}
}

func hydrateConfig(c *Config) error {
	listLogger.Debug("Hydrating config")
	if err := vipCfg.Unmarshal(c); err != nil {
		listLogger.Error(err, "failed to unmarshal config")
	}
	return nil
}
