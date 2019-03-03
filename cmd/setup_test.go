package cmd

import (
	. "github.com/onsi/gomega"
	. "github.com/onsi/ginkgo"

	"github.com/spf13/afero"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"workspaces/util"
	"workspaces/_vendor-20190219161801/github.com/mitchellh/go-homedir"
	"fmt"
)

var _ = Describe("Init Command", func() {
	Describe("createDirIfNotExist", func() {
		BeforeEach(func() {
			AppFs = afero.NewMemMapFs()
			logrus.SetOutput(ioutil.Discard)
		})
		When("base directory does not exist", func() {
			It("should create the directory", func() {
				home, _ := homedir.Dir()
				basePath := fmt.Sprintf("%s/%s", home, BASE_DIRECTORY_PATH)
				if exist, err := afero.DirExists(AppFs, basePath); exist  {
					Fail("Folder should not exist")
				} else if err != nil {
					Fail("Failed to detect if directory exist")
				}
				createBaseDirIfNotExist()
				if exist, err := afero.DirExists(AppFs, basePath); err != nil {
					Fail("Failed to detect if directory exist")
				} else if !exist {
					Fail("Failed to create base directory")
				}
			})
		})
	})

	Describe("createSampleConfig", func() {
		BeforeEach(func() {
			AppFs = afero.NewMemMapFs()
			logrus.SetOutput(ioutil.Discard)
		})
		When("called", func() {
			It("should create a config with default values", func() {
				cfg := createSampleConfig()
				Expect(cfg.Get("workspace_dir")).To(Equal("~/workspaces"))
				Expect(cfg.Get("projects")).To(Equal(make(map[string]interface{})))
			})
		})
	})

	Describe("printConfigToStdout", func() {
		Context("create sample config", func() {
			It("should print to stdout", func() {
				soWatcher := util.StdoutWatcher{}
				soWatcher.Start()
				cfg := createSampleConfig()
				printConfigToStdout(cfg)
				Expect(soWatcher.Stop()).To(ContainSubstring(`"workspace_dir": "~/workspaces"`))
			})
		})
	})
})
