package config_model

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"workspaces/_vendor-20190219161801/github.com/spf13/viper"
	"fmt"
	"bytes"
)

var _ = Describe("Config", func() {
	Context("Parsing", func() {
		It("Should be able to handle parse", func() {
			js := []byte(`{
              "workspace_dir": "~/some_dir",
              "projects": [
                 {
                   "name": "proj1",
                   "path": "some_path",
                   "git": false
                 },
                 {
                   "name": "proj2",
                   "path": "some_path2",
                   "git": true
                 }
              ]
            }`)
			v := viper.New()
			v.SetConfigType("json")
			if err := v.ReadConfig(bytes.NewReader(js)); err != nil {
				Fail(fmt.Sprintf("Failed to read in config %s", err.Error()))
			}
			var config Config
			if err := v.Unmarshal(&config); err != nil {
				Fail(fmt.Sprintf("Failed to Unmarshal setting: %s", err.Error()))
			}

			Expect(config.BaseDir).To(Equal("~/some_dir"))
			Expect(config.Projects[0]).To(Equal(Project{"proj1", "some_path", false}))
			Expect(config.Projects[1]).To(Equal(Project{"proj2", "some_path2", true}))
		})
	})
})
