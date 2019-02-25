package cmd

import (
	. "github.com/onsi/ginkgo"

	"github.com/spf13/afero"
	"github.com/sirupsen/logrus"
	. "github.com/onsi/gomega"
	"github.com/spf13/viper"
)

var _ = Describe("Root Command", func() {
	Describe("InitLog", func() {
		When("Initialized with verbose value to be true", func() {
			It("Should set logrus level to debug", func() {
				verbose = true
				initLog()
				Expect(logrus.GetLevel(), logrus.DebugLevel)
			})
		})
		Describe("InitLog", func() {
			When("Initialized with verbose value to be false", func() {
				It("Should set logrus level to debug", func() {
					verbose = false
					initLog()
					Expect(logrus.GetLevel(), logrus.DebugLevel)
				})
			})
		})
	})

	Describe("InitConfig", func() {
		BeforeEach(func() {
			// clear config path
			cfgPath = ""
		})
		When("config path is provided by flag", func() {
			It("should set viper config path to the provided path", func() {
				cfgPath = "/some_path.yaml"
				fs := afero.NewMemMapFs()
				// use a new viper to write a config to the path
				v := viper.New()
				v.SetFs(fs)
				v.Set("some_key", "some_value")
				v.WriteConfigAs(cfgPath)

				// point viper to the correct file system
				viper.SetFs(fs)

				// should init with cfgPath
				initConfig()

				Expect(viper.Get("some_key"), "some_value")
			})
		})

		/**
		TODO: make the function testable with homedir
		When("config path is not provided by flag", func() {
			It("should set viper config path to default path: ~/.workspaces/config.yaml", func() {
				fs := afero.NewMemMapFs()
				// use a new viper to write a config to the path
				v := viper.New()
				v.SetFs(fs)
				v.Set("some_key", "some_value")
				v.WriteConfigAs(cfgPath)

				v2 := viper.New()
				v2.SetFs(fs)
				v2.Set("some_key", "some_other_value")
				v2.WriteConfigAs("~/.workspaces/config.yaml")

				// point viper to the correct file system
				viper.SetFs(fs)
				viper.GetViper()

				// should init with cfgPath
				initConfig()

				Expect(viper.Get("some_key"), "some_other_value")
				fmt.Println(viper.Get("some_key"))
			})
		})
		**/
	})
})
