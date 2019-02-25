package cmd

import (
	. "github.com/onsi/gomega"
	. "github.com/onsi/ginkgo"

	"github.com/spf13/afero"
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

var _ = Describe("Init Command", func() {
	Describe("createDirIfNotExist", func() {
		BeforeEach(func() {
			fs = afero.NewMemMapFs()
			logrus.SetOutput(ioutil.Discard)
		})
		When("base directory does not exist", func() {
			It("should create the directory", func() {
				if exist, err := afero.DirExists(fs, BASE_DIRECTORY_PATH); exist  {
					Fail("Folder should not exist")
				} else if err != nil {
					Fail("Failed to detect if directory exist")
				}
				createBaseDirIfNotExist()
				if exist, err := afero.DirExists(fs, BASE_DIRECTORY_PATH); err != nil {
					Fail("Failed to detect if directory exist")
				} else if !exist {
					Fail("Failed to create base directory")
				}
			})
		})
	})
	Describe("createSampleConfig", func() {
		BeforeEach(func() {
			fs = afero.NewMemMapFs()
			logrus.SetOutput(ioutil.Discard)
		})
		When("called", func() {
			It("should create a config with default values", func() {
				cfg := createSampleConfig()
				Expect(cfg.Get("workspace_dir")).To(Equal("~/workspaces"))
			})
		})
	})
})




