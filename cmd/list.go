package cmd

import (
	"github.com/spf13/cobra"
	"workspaces/logging"
	. "workspaces/config_model"
	"fmt"
	"encoding/json"
	"github.com/spf13/viper"
	"github.com/fatih/color"
)

var (
	listLogger logging.LoggableEntity
	detail bool

	listCmd = &cobra.Command{
		Use: "list",
		Aliases: []string{"ls"},
		Short: "Show all projects in registry",
		Long: `Show all projects in registry`,
		Run: func(cmd *cobra.Command, args []string) {
			initConfig()
			printAllProjects(&cfg)
		},
	}
)

func init() {
	listLogger.Debug("initializing list logger")
	listLogger = logging.NewLoggableEntity("list", logging.Fields{ "module": "list" })

	listCmd.Flags().BoolVarP(&detail, "detail","d", false, "show list with detail")
	viper.BindPFlag("with_detail", listCmd.Flags().Lookup("with_detail"))
}

func printAllProjects(c *Config)  {
	listLogger.Debug("Printing all projects")
	listLogger.Debug(fmt.Sprintf("We have %v projects", len(c.Projects)))
	if len(c.Projects) == 0 {
		fmt.Println("No projects found in registry, use `workpsaces scan /path/to/dir` or `workspaces pick /path/to/project` to add project ")
	}
	for _, p := range c.Projects {
		listLogger.Debugf(logging.Fields{"name": p.Name, "path": p.Path }, "project")
		if detail {
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
		} else {
			highlight := color.New(color.FgHiBlack, color.BgHiCyan).SprintFunc()
			fmt.Printf("  %s(%s) : %s\n", highlight(p.Name), p.Alias, p.Path)
		}
	}
}

func hydrateConfig(vc *viper.Viper, c *Config) error {
	listLogger.Debug("hydrating")
	listLogger.Debug((*vc).AllSettings())
	if err := (*vc).Unmarshal(c); err != nil {
		listLogger.Error(err, "failed to unmarshal config")
	}
	return nil
}
