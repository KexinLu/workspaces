package cmd

import (
	"github.com/spf13/cobra"
	"workspaces/logging"
	"workspaces/_vendor-20190219161801/github.com/spf13/viper"
	"workspaces/util"
	"os"
	"fmt"
	"workspaces/_vendor-20190219161801/github.com/mitchellh/go-homedir"
	"encoding/json"
	"bufio"
	"errors"
	"github.com/spf13/afero"
)

var (
	toStdout bool
	force bool
	logger logging.LoggableEntity
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
	logger = logging.NewLoggableEntity( "init", logging.Fields{ "module": "init" })
	setupCmd.Flags().BoolVarP(&toStdout, "stdout", "s", false, "output example config to stdout")
	setupCmd.Flags().BoolVarP(&force, "force", "f", false, "force replace of config file")
	if dir, err := homedir.Dir(); err != nil {
		logger.Fatal(err.Error(), "failed to located home dir")
	} else {
		baseDirPath = fmt.Sprintf(`%s/%s`, dir, BASE_DIRECTORY_PATH)
	}
}

// generate
func createBaseDirIfNotExist() {
	hdPath, err := homedir.Dir()
	if err != nil {
		logger.Error(err, "Failed to locate home directory")
	}

	baseDirPath := fmt.Sprintf(`%s/%s`, hdPath, BASE_DIRECTORY_PATH)
	if err := util.EnsureDirExistSilent(baseDirPath, AppFs); err != nil {
		if err.Error() == util.FOLDER_NOT_FOUND {
			logger.Debug("creating base dir")
			if err := AppFs.Mkdir(baseDirPath, 0744); err != nil  {
				logger.Error(err, "Failed to create base directory")
			}
		} else {
			logger.Error(err, "Failed to create base directory, init failed.")
			os.Exit(1)
		}
	}
}

var defaultCfg = map[string]interface{}{
	"workspace_dir": "~/workspaces",
	"projects": make(map[string]interface{}),
}

func createSampleConfig() *viper.Viper {
	cfg := viper.New()
	for k,v := range defaultCfg {
		cfg.Set(k, v)
	}

	return cfg
}

// print the value of viper config to stdout
func printConfigToStdout(cfg *viper.Viper) error {
	logger.Debug("Printing config to stdout")
	c := cfg.AllSettings()
	if bs, err := json.MarshalIndent(c, "", "  "); err != nil {
		logger.Fatal("Failed to marshal config")
		return err
	} else {
		w := bufio.NewWriter(os.Stdout)
		if n, err := w.Write(bs); err != nil {
			logger.Fatal("Failed to write config to stdout")
			return err
		} else if n < len(bs) {
			err := errors.New("failed to write all config to stdout")
			logger.Fatal(err.Error())
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
		logger.Error(err,"Failed to check if base config exist")
	} else if exist && !force {
		fmt.Println("config file already exist, cowardly refuse to replace it, use -f or --force=true to froce replace")
		os.Exit(1)
	}

	if err := util.EnsureFileExistSilent(baseConfigFile, AppFs); err != nil {
		logger.Debug("Config file does not exist")
		if _, err := AppFs.Create(baseConfigFile); err != nil {
			logger.Error(err, "Failed to create sample config file")
		}
	}
	cfg.SetConfigFile(baseConfigFile)
	if err := cfg.WriteConfig(); err != nil {
		logger.Error(err, "Failed to write to sample config file")
	}
	logger.Debug("Finish writing config")
}
