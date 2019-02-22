package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"workspaces/logging"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/afero"
	"github.com/pkg/errors"
)

var (
	cfgPath string
	verbose bool
	logDir string
	AppFs = afero.NewOsFs()

	rootCmd = &cobra.Command{
		Use: "workspaces",
		Short: "workspaces is a workspace management tool",
		Long: `Set up your workspace with one click on a brand new work station, or navigate between your projects with ease`,
		Run: func(cmd *cobra.Command, args []string) {
			//
		},
	}
	rootLoggable = logging.NewLoggableEntity(
		"root",
		logging.Fields{
			"module": "root",
		},
	)
)

func Execute() {
	rootLoggable.Debug("Root Command execute")
	if err := rootCmd.Execute(); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}

const (
	VERBOSE = "verbose"
	CONFIG = "config"
	LOGDIR = "log_dir"
)

func init() {
	cobra.OnInitialize(initLog)
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().BoolVarP(&verbose,VERBOSE, "v", false, "-v or --verbose for debug information")
	rootCmd.PersistentFlags().StringVarP(&cfgPath, CONFIG, "f", "", "config file (default is $HOME/.workspaces/config)")
	rootCmd.PersistentFlags().StringVar(&logDir, LOGDIR,"", "log directory, default to ~/.workspaces/log")
	viper.BindPFlag(VERBOSE, rootCmd.PersistentFlags().Lookup(VERBOSE))
	viper.BindPFlag(CONFIG, rootCmd.PersistentFlags().Lookup(CONFIG))
	viper.BindPFlag(LOGDIR, rootCmd.PersistentFlags().Lookup(LOGDIR))

	//rootCmd.AddCommand(addCmd)
	//rootCmd.AddCommand(initCmd)
}

// config file sequence
// 1. flag
// 2. default at ~/.workspaces
func initConfig() {
	if cfgPath != "" {
		rootLoggable.Debugf(logging.Fields{"path": cfgPath}, "setting config path from flag")
		viper.SetConfigFile(cfgPath)
	} else {
		rootLoggable.Debug("Config path not provided, using default $HOME/.workspaces/config")
		// Get home dir
		if homeDir, err := homedir.Dir(); err != nil {
			rootLoggable.Panic(err,"failed to get home directory", )
		} else {
			path := homeDir + `/.workspaces`
			rootLoggable.Debug("Using path " + path)
			ensureBaseDir(path)

			viper.SetConfigType("json")
			viper.AddConfigPath(path)
			viper.SetConfigName("config")
		}
	}

	rootLoggable.Debug("Config Automatic Env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		rootLoggable.Error(err, "Cannot load config")
		rootLoggable.Info( "No config file provided")
		rootLoggable.Info("create ~/.workspaces")
		rootLoggable.Info("or use --config=/path/to/config")
		rootLoggable.Info("to specify config file")
	}
}

func ensureBaseDir(path string) {
	rootLoggable.Debug("Trying to confirm directory exist")
	if exist, err := afero.DirExists(AppFs, path); err != nil {
		rootLoggable.Errorf(
			logging.Fields{
				"path": path,
			},
			err,
			"failed to check if base directory exist",
		)
		os.Exit(1)
	} else {
		if !exist {
			rootLoggable.Error(
				errors.New(FOLDER_NOT_FOUND),
				"base folder ~/.workspaces not found",
			)
			os.Exit(1)
		}
	}
	rootLoggable.Debug("Found base directory")
}

func initLog() {
	if verbose {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}
}

