package cmd

import (
	"github.com/spf13/cobra"
	"workspaces/logging"
	"github.com/spf13/afero"
	"github.com/manifoldco/promptui"
	"fmt"
	. "workspaces/config_model"
	"github.com/pkg/errors"
	"github.com/KexinLu/goisgit"
	"github.com/spf13/viper"
	"os"
)

var (
	scanLogger logging.LoggableEntity

	scanCmd = &cobra.Command{
		Use: "scan",
		Short: "Show all projects managed by workspaces",
		Long: `Show all projects managed by workspaces in the config file`,
		Args: func(cmd *cobra.Command, args []string) error {
			initConfig()
			if len(args) == 0 {
				scanBaseDir()
			} else if len(args) > 0{
				scanDir(args[0:])
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {},
	}
	forceScan bool
)

func init() {
	scanLogger = logging.NewLoggableEntity("scan", logging.Fields{ "module": "scan" })
	scanLogger.Debug("initializing scan logger")

	scanCmd.Flags().BoolVarP(&forceScan, "force","f", false, "force replace existing projects")
	viper.BindPFlag("forceScan", rootCmd.PersistentFlags().Lookup(VERBOSE))
}

func scanBaseDir() {
	scanDir([]string{cfg.BaseDir})
}

func scanDir(dirs []string) {
	for _, dir := range dirs {
		scanLogger.DebugWithFields(logging.Fields{"dir": dir}, "scanning working dir")
		if fis, err := afero.ReadDir(AppFs, dir); err != nil {
			scanLogger.Fatal("Failed to scan base dir: ", cfg.BaseDir,  err.Error())
		} else {
			for _, fi := range fis {
				if fi.IsDir() {
					path := fmt.Sprintf(`%s/%s`, dir, fi.Name())
					if p, err := buildProject(path); err != nil {
						scanLogger.Fatal("Failed to build project: ", path,  err.Error())
						os.Exit(1)
					} else {
						prompt := buildPromptP(p)
						if result, err := prompt.Run(); err != nil {
							scanLogger.Error(err, "Failed to get a response")
							os.Exit(1)
						} else {
							if result == "y" || result == "Y" {
								appendToProjMap(p, &cfg, forceScan)
							}
						}
					}
				}
			}
			printToConfig()
		}
	}
}

func buildPromptP(p Project) promptui.Prompt {
	return promptui.Prompt{
		Label:    "Add project? (y/n) " + p.Name,
		Validate: func(s string) error {
			allowed := map[string]interface{}{"y": 0, "n":0, "Y":0, "N":0}
			if _, ok := allowed[s]; !ok {
				return errors.New("only y/n allowed")
			}
			return nil
		},
		Default:  "n",
	}
}

func appendToProjMap(proj Project, cfg *Config, force bool) {
	if _, exist := cfg.Projects[proj.Name]; exist && !force {
		scanLogger.Error(errors.New("project_exist_in_config") , "Refuse to add the same project")
	}

	if isGit, err := is_git.IsGitDir(proj.Path); err != nil {
		scanLogger.Error(errors.New("failed_to_confirm_if_is_git") , err.Error())
	} else if isGit {
		proj.IsGit = true
	}


	cfg.Projects[proj.Name] = proj
}
