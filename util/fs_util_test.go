package util

import (
	. "github.com/onsi/ginkgo"
	//. "github.com/onsi/gomega"

	"github.com/spf13/afero"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"fmt"
)

var _ = Describe("FileSystem", func() {
	var fs afero.Fs
	BeforeEach(func() {
		logrus.SetOutput(ioutil.Discard)
		fs = afero.NewMemMapFs()
	})

	Describe("EnsureDirExist", func() {
		Context("directory exist", func() {
			It("should not return error", func() {
				path := "~/.path"
				fs.Mkdir(path, 0777)
				if err := EnsureDirExist(path, fs); err != nil {
					Fail("Should not return error")
				}
			})
		})

		Context("directory does not exist", func() {
			It("should return error", func() {
				path := "~/.path"
				if err := EnsureDirExist(path, fs); err == nil {
					Fail("Should return error")
				}
			})
		})
	})

	Describe("EnsureFileExist", func() {
		When("file exist", func() {
			It("should not return error", func() {
				path := "~/.some_path"
				fs.Create(path)
				if err := EnsureFileExist(path, fs); err != nil {
					Fail(fmt.Sprintf("Should not return error: %s", err.Error()))
				}
			})
		})

		When("file does not exist", func() {
			It("should return error", func() {
				path := "~/.path"
				if err := EnsureFileExist(path, fs); err == nil {
					Fail(fmt.Sprintf("Should return error"))
				}
			})
		})
	})
})
