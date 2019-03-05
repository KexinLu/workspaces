package cmd

import (
	"github.com/spf13/cobra"
	"workspaces/logging"
	"github.com/spf13/afero"
	"github.com/manifoldco/promptui"
	"fmt"
	. "workspaces/config_model"
	"encoding/json"
	"github.com/pkg/errors"
	"os"
	"github.com/KexinLu/goisgit"
	"github.com/spf13/viper"
	"path/filepath"
)

var (
	pickLogger logging.LoggableEntity

	pickCmd = &cobra.Command{
		Use: "pick",
		Short: "Add directory to managed projects",
		Long: `Add directory to managed projects`,
		Args: func(cmd *cobra.Command, args []string) error {
			initConfig()
			pickDirs(args[0:])
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {},
	}
	forcePick bool
)

func init() {
	pickLogger = logging.NewLoggableEntity("pick", logging.Fields{ "module": "pick" })
	pickLogger.Debug("initializing pick logger")

	pickCmd.Flags().BoolVarP(&forcePick , "force","f", false, "force replace existing projects")
	viper.BindPFlag("forcePick", rootCmd.PersistentFlags().Lookup(VERBOSE))
}

func pickDirs(dirs []string) {
	for _, dir := range dirs {
		pickLogger.DebugWithFields(logging.Fields{"dir": dir}, "picking dir")
		p, err := buildProject(dir)
		if err != nil {
			pickLogger.Fatal("Failed to build project : ", dir,  err.Error())
		}
		prompt := buildPromptP(p)
		if result, err := prompt.Run(); err != nil {
			pickLogger.Error(err, "Failed to get a response")
		} else {
			if result == "y" || result == "Y" {
				appendToProjMap(p, &cfg, forcePick)
			}
		}
	}

	printToConfig()
}

func printToConfig() {
	if cfgBytes, err := json.MarshalIndent(cfg, "", "  "); err != nil {
		pickLogger.Fatal(err.Error(), "Failed to marshal config")
		os.Exit(1)
	} else {
		afero.WriteFile(AppFs, cfgPath, cfgBytes, 0755)
	}
}

func buildProject(path string) (Project, error) {
	abs, err := filepath.Abs(path)
	if err != nil {
		return Project{}, fmt.Errorf("can not parse absolute path of %s: %s", path, err.Error())
	}
	if err := ensureIsDir(abs); err != nil {
		return Project{}, err
	}
	isGit, err := is_git.IsGitDir(path)
	if err != nil {
		return Project{}, errors.New("failed_to_confirm_if_is_git" + err.Error())
	}
	return Project{
		Name: filepath.Base(abs),
		Path: abs,
		IsGit: isGit,
	}, nil
}

func buildPrompt(fi os.FileInfo) promptui.Prompt {
	return promptui.Prompt{
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
}

func ensureIsDir(path string) error {
	if isDir, err := afero.IsDir(AppFs, path); err != nil {
		return fmt.Errorf("failed to check if %s is a directory", path)
	} else if !isDir {
		return fmt.Errorf("%s is not a directory", path)
	}
	return nil
}

