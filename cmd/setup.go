package cmd

import (
	"github.com/spf13/cobra"
	"workspaces/logging"
	"github.com/spf13/viper"
	"workspaces/util"
	"os"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"encoding/json"
	"bufio"
	"errors"
	"github.com/spf13/afero"
)

var (
	toStdout bool
	force bool
	setupLogger logging.LoggableEntity
	baseDirPath string

	setupCmd = &cobra.Command{
		Use: "setup",
		Short: "initialize ~/.workspaces folder and ~/.workspaces/config",
		Long: `Set up your workspace with one click on a brand new work station, or navigate between your projects with ease`,
		Run: func(cmd *cobra.Command, args []string) {
			cfg := createSampleConfig()
			if toStdout {
				printConfigToStdout(cfg)
			} else {
				createBaseDirIfNotExist()
				storeConfigFile(cfg)
			}
		},
	}
)

func init() {
	setupLogger = logging.NewLoggableEntity( "init", logging.Fields{ "module": "init" })
	setupCmd.Flags().BoolVarP(&toStdout, "stdout", "s", false, "output example config to stdout")
	setupCmd.Flags().BoolVarP(&force, "force", "f", false, "force replace of config file")
	if dir, err := homedir.Dir(); err != nil {
		setupLogger.Fatal(err.Error(), "failed to located home dir")
	} else {
		baseDirPath = fmt.Sprintf(`%s/%s`, dir, BASE_DIRECTORY_PATH)
	}
}

// generate
func createBaseDirIfNotExist() {
	hdPath, err := homedir.Dir()
	if err != nil {
		setupLogger.Error(err, "Failed to locate home directory")
	}

	baseDirPath := fmt.Sprintf(`%s/%s`, hdPath, BASE_DIRECTORY_PATH)
	if err := util.EnsureDirExistSilent(baseDirPath, AppFs); err != nil {
		if err.Error() == util.FOLDER_NOT_FOUND {
			setupLogger.Debug("creating base dir")
			if err := AppFs.Mkdir(baseDirPath, 0755); err != nil  {
				setupLogger.Error(err, "Failed to create base directory")
			}
		} else {
			setupLogger.Error(err, "Failed to create base directory, init failed.")
			os.Exit(1)
		}
	}

	workspacePath := fmt.Sprintf(`%s/%s`, hdPath, "workspaces")
	if err := util.EnsureDirExistSilent(workspacePath, AppFs); err != nil {
		if err.Error() == util.FOLDER_NOT_FOUND {
			setupLogger.Debug("creating workspace dir")
			if err := AppFs.Mkdir(baseDirPath, 0755); err != nil  {
				setupLogger.Error(err, "Failed to create workspace directory")
			}
		} else {
			setupLogger.Error(err, "Failed to create workspace directory, init failed.")
			os.Exit(1)
		}
	}
}

var defaultCfg = map[string]interface{}{
	"projects": make(map[string]interface{}),
}

func createSampleConfig() *viper.Viper {
	cfg := viper.New()
	if dir, err := homedir.Dir(); err != nil {
		setupLogger.Error(err, "Failed to find home dir" )
	} else {
		cfg.Set("workspace_dir", dir+"/workspaces")
	}

	for k,v := range defaultCfg {
		cfg.Set(k, v)
	}

	return cfg
}

// print the value of viper config to stdout
func printConfigToStdout(cfg *viper.Viper) error {
	setupLogger.Debug("Printing config to stdout")
	c := cfg.AllSettings()
	if bs, err := json.MarshalIndent(c, "", "  "); err != nil {
		setupLogger.Fatal("Failed to marshal config")
		return err
	} else {
		w := bufio.NewWriter(os.Stdout)
		if n, err := w.Write(bs); err != nil {
			setupLogger.Fatal("Failed to write config to stdout")
			return err
		} else if n < len(bs) {
			err := errors.New("failed to write all config to stdout")
			setupLogger.Fatal(err.Error())
			return err
		}
		// new line
		w.Write([]byte("\n"))
		w.Flush()
	}
	return nil
}

func storeConfigFile(cfg *viper.Viper) {
	baseConfigFile := fmt.Sprintf(`%s/%s`,baseDirPath, "config.json" )
	if exist, err := afero.Exists(AppFs, baseConfigFile); err != nil {
		setupLogger.Error(err,"Failed to check if base config exist")
	} else if exist && !force {
		fmt.Println("config file already exist, cowardly refuse to replace it, use -f or --force=true to froce replace")
		os.Exit(1)
	}

	if err := util.EnsureFileExistSilent(baseConfigFile, AppFs); err != nil {
		setupLogger.Debug("Config file does not exist")
		if _, err := AppFs.Create(baseConfigFile); err != nil {
			setupLogger.Error(err, "Failed to create sample config file")
		}
	}
	cfg.SetConfigFile(baseConfigFile)
	if err := cfg.WriteConfig(); err != nil {
		setupLogger.Error(err, "Failed to write to sample config file")
	}
	setupLogger.Debug("Finish writing config")
}
