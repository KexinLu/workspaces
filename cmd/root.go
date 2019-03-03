package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"workspaces/logging"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/afero"
	"workspaces/util"
)

var (
	cfgPath string
	verbose bool
	logDir string
	vipCfg *viper.Viper
	AppFs = afero.NewOsFs()

	rootCmd = &cobra.Command{
		Use: "workspaces",
		Short: "workspaces is a workspace management tool",
		Long: `Set up your workspace with one click on a brand new work station, or navigate between your projects with ease`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 || args[0] != "init" {
				initConfig()
				return nil
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	rootLogger = logging.NewLoggableEntity(
		"root",
		logging.Fields{
			"module": "root",
		},
	)
)

// Execute root command
func Execute() {
	rootLogger.Debug("Root Command execute")
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

// init root command
func init() {
	cobra.OnInitialize(initLog)

	rootCmd.PersistentFlags().BoolVarP(&verbose,VERBOSE, "v", false, "-v or --verbose for debug information")
	rootCmd.PersistentFlags().StringVarP(&cfgPath, CONFIG, "c", "", "config file (default is $HOME/.workspaces/config)")
	rootCmd.PersistentFlags().StringVar(&logDir, LOGDIR,"", "log directory, default to ~/.workspaces/log")
	viper.BindPFlag(VERBOSE, rootCmd.PersistentFlags().Lookup(VERBOSE))
	viper.BindPFlag(CONFIG, rootCmd.PersistentFlags().Lookup(CONFIG))
	viper.BindPFlag(LOGDIR, rootCmd.PersistentFlags().Lookup(LOGDIR))

	rootCmd.AddCommand(setupCmd)
	rootCmd.AddCommand(listCmd)
}

// config file sequence
// 1. flag
// 2. default at ~/.workspaces
func initConfig() {
	vipCfg = viper.New()
	vipCfg.SetConfigType("json")
	if cfgPath != "" {
		rootLogger.Debugf(logging.Fields{"path": cfgPath}, "setting config path from flag")
		vipCfg.SetConfigFile(cfgPath)
	} else {
		rootLogger.Debug("Config path not provided, using default $HOME/.workspaces/config")
		// Get home dir
		if homeDir, err := homedir.Dir(); err != nil {
			rootLogger.Panic(err,"failed to get home directory", )
			os.Exit(1)
		} else {
			path := homeDir + `/.workspaces`
			rootLogger.Debug("Using path " + path)
			if err := util.EnsureDirExist(path, AppFs); err != nil {
				rootLogger.ErrorWithFields(logging.Fields{ "path": path }, err, "Failed to ensure base directory")
				os.Exit(1)
			}

			vipCfg.AddConfigPath(path)
			vipCfg.SetConfigName("config")
		}
	}

	rootLogger.Debug("Config Automatic Env")
	viper.AutomaticEnv()

	if err := vipCfg.ReadInConfig(); err != nil {
		rootLogger.Error(err, "Cannot load config")
		rootLogger.Info( "No config file provided")
		rootLogger.Info("create ~/.workspaces")
		rootLogger.Info("or use --config=/path/to/config")
		rootLogger.Info("to specify config file")
	}
}

// initialize logging.
// When verbose is true, set log level to debug and info otherwise
func initLog() {
	if verbose {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}
}
