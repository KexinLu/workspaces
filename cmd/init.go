package cmd

import (
	"github.com/spf13/cobra"
	"workspaces/logging"
	"workspaces/_vendor-20190219161801/github.com/spf13/viper"
	"github.com/spf13/afero"
	"workspaces/util"
	"io"
	"io/ioutil"
	"os"
)

var (
	initCmd = &cobra.Command{
		Use: "init",
		Short: "initialize ~/.workspaces folder and ~/.workspaces/config",
		Long: `Set up your workspace with one click on a brand new work station, or navigate between your projects with ease`,
		Run: func(cmd *cobra.Command, args []string) {
			createBaseDirIfNotExist()

		},
	}

	fs afero.Fs

	toStdout bool
	initModule logging.LoggableEntity
)

func init() {
	initModule = logging.NewLoggableEntity( "init", logging.Fields{ "module": "init" })
	initCmd.Flags().BoolVarP(&toStdout, "stdout", "s", false, "output example config to stdout")
	fs = afero.NewOsFs()
}

// generate
func createBaseDirIfNotExist() {
	if err := util.EnsureDirExist(BASE_DIRECTORY_PATH, fs); err != nil {
		if err.Error() == util.FOLDER_NOT_FOUND {
			fs.Mkdir(BASE_DIRECTORY_PATH, 0744)
		} else {
			initModule.Error(err, "Failed to create base directory, init failed.")
			os.Exit(1)
		}
	}
}

func produceConfigFile() {
	//defaultCfg := createSampleConfig()
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
		fileSystem := afero.NewMemMapFs()
		cfg.SetFs(fileSystem)
		cfg.WriteConfigAs("/tmp")
		if file, err := fileSystem.Open("/tmp"); err != nil {
			initModule.Error(err, "Failed to open temporary config")
			return err
		} else {
			io.Copy(os.Stdout, file)
			return nil
		}
}

func storeConfigFile() {

}



