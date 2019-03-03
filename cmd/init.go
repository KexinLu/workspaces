package cmd

import (
	"github.com/spf13/cobra"
	"workspaces/logging"
	"workspaces/_vendor-20190219161801/github.com/spf13/viper"
	"github.com/spf13/afero"
	"workspaces/util"
	"io"
	"os"
	"fmt"
	"workspaces/_vendor-20190219161801/github.com/mitchellh/go-homedir"
)

var (
	fs afero.Fs
	toStdout bool
	logger logging.LoggableEntity
	baseDirPath string

	initCmd = &cobra.Command{
		Use: "init",
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
	initCmd.Flags().BoolVarP(&toStdout, "stdout", "s", false, "output example config to stdout")
	fs = afero.NewOsFs()
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
	if err := util.EnsureDirExistSilent(baseDirPath, fs); err != nil {
		if err.Error() == util.FOLDER_NOT_FOUND {
			logger.Debug("creating base dir")
			if err := fs.Mkdir(baseDirPath, 0744); err != nil  {
				logger.Error(err, "Failed to create base directory")
			}
		} else {
			logger.Error(err, "Failed to create base directory, init failed.")
			os.Exit(1)
		}
	}
}

var defaultCfg = map[string]string{
	"workspace_dir": "~/workspaces",
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
	fs := afero.NewMemMapFs()
	fs.Create("/tmp.json")
	cfg.SetFs(fs)
	cfg.SetConfigFile("/tmp.json")
	if err := cfg.WriteConfig() ; err != nil {
		logger.Fatal("Failed to write config in memory: " + err.Error())
	}
	if file, err := fs.Open("/tmp.json"); err != nil {
		logger.Error(err, "Failed to open temporary config")
		return err
	} else {
		if _,err := io.Copy(os.Stdout, file); err != nil {
			logger.Fatal(err)
		}
		return nil
	}
}

func storeConfigFile(cfg *viper.Viper) {
	baseConfigFile := fmt.Sprintf(`%s/%s`,baseDirPath, "config.json" )
	if err := util.EnsureFileExistSilent(baseConfigFile, fs); err != nil {
		logger.Debug("Config file does not exist")
		if _, err := fs.Create(baseConfigFile); err != nil {
			logger.Error(err, "Failed to create sample config file")
		}
	}
	cfg.SetConfigFile(baseConfigFile)
	if err := cfg.WriteConfig(); err != nil {
		logger.Error(err, "Failed to write to sample config file")
	}
	logger.Debug("Finish writing config")
}
