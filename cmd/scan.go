package cmd

import (
	"github.com/spf13/cobra"
	"workspaces/logging"
	"github.com/spf13/afero"
	"github.com/manifoldco/promptui"
	"fmt"
	"workspaces/config_model"
	"encoding/json"
	"github.com/pkg/errors"
	"os"
)

var (
	scanLogger logging.LoggableEntity

	scanCmd = &cobra.Command{
		Use: "scan",
		Short: "Show all projects managed by workspaces",
		Long: `Show all projects managed by workspaces in the config file`,
		Run: func(cmd *cobra.Command, args []string) {
			initConfig()
			scanWorkDir()
		},
	}
)

func init() {
	scanLogger = logging.NewLoggableEntity("scan", logging.Fields{ "module": "scan" })
	scanLogger.Debug("initializing scan logger")
}

func scanWorkDir() {
	fs := afero.NewOsFs()
	scanLogger.DebugWithFields(logging.Fields{"dir": cfg.BaseDir}, "scanning working dir")
	if fis, err := afero.ReadDir(fs, cfg.BaseDir); err != nil {
		scanLogger.Fatal("Failed to scan base dir: ", err.Error())
	} else {
		for _, fi := range fis {
			 p := config_model.Project{
				Name: fi.Name(),
				Path: fmt.Sprintf(`%s/%s`, cfg.BaseDir, fi.Name()),
				IsGit: false,
			}
			//bs, _ := json.MarshalIndent(p, "", "  ")
			prompt := promptui.Prompt{
				Label:    "Add project? (y/n) " + fi.Name(),
				Validate: func(s string) error {
					allowed := map[string]interface{}{"y": 0, "n":0, "Y":0, "N":0}
					if _, ok := allowed[s]; !ok {
						return errors.New("only y/n allowed")
					}
					return nil
				},
				Default:  "n",
			}

			if result, err := prompt.Run(); err != nil {
				scanLogger.Error(err, "Failed to get a response")
			} else {
				if result == "y" || result == "Y" {
					cfg.Projects = append(cfg.Projects, p)
				}
			}
		}
		if cfgBytes, err := json.MarshalIndent(cfg, "", "  "); err != nil {
			scanLogger.Fatal(err.Error(), "Failed to marshal config")
			os.Exit(1)
		} else {
			afero.WriteFile(AppFs, cfgPath, cfgBytes, 0755)
		}
	}
}
