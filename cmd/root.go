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
	"workspaces/config_model"
)

var (
	cfgPath string
	cfg config_model.Config
	verbose bool
	logDir string
	vipCfg *viper.Viper
	AppFs = afero.NewOsFs()

	rootCmd = &cobra.Command{
		Use: "workspaces",
		Short: "workspaces is a workspace management tool",
		Long: `Set up your workspace with one click on a brand new work station, or navigate between your projects with ease`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 || args[0] != "setup" {
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
	rootCmd.AddCommand(pickCmd)
	rootCmd.AddCommand(scanCmd)
	rootCmd.AddCommand(wdCmd)
}

// config file sequence
// 1. flag
// 2. default at ~/.workspaces
func initConfig() {
	vipCfg = viper.New()
	vipCfg.SetConfigType("json")
	rootLogger.Debug("trying to hydrate config")
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
			cfgDir := homeDir + `/.workspaces`
			cfgPath = cfgDir + `/config.json`
			rootLogger.Debug("Using path " + cfgPath)
			if err := util.EnsureDirExist(cfgDir, AppFs); err != nil {
				rootLogger.ErrorWithFields(logging.Fields{ "path": cfgDir }, err, "Failed to ensure base directory")
				os.Exit(1)
			}

			vipCfg.SetConfigFile(cfgPath)
		}
	}

	rootLogger.Debug("Config Automatic Env")
	viper.AutomaticEnv()

	rootLogger.Debug("Reading config from file")
	if err := vipCfg.ReadInConfig(); err != nil {
		rootLogger.Error(err, "Cannot load config")
		rootLogger.Info( "No config file provided")
		rootLogger.Info("run workspaces setup to setup")
		rootLogger.Info("or use --config=/path/to/config")
		rootLogger.Info("to specify config file")
	}

	if err := hydrateConfig(vipCfg, &cfg); err != nil {
		rootLogger.Fatal(err.Error(), "fail to hydrate config")
		os.Exit(1)
	}

	if cfg.Projects == nil {
		cfg.Projects = make(map[string]config_model.Project)
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
