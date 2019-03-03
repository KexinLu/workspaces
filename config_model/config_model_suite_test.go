package config_model_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestConfigModel(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ConfigModel Suite")
}
