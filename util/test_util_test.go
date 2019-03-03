package util

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"bufio"
	"os"
)

var _ = Describe("StdoutWatcher", func(){
	Context("watch is called, and has printed to stdout", func() {
		When("close is called", func() {
			It("should text from stdout", func() {
				watcher := StdoutWatcher{}

				watcher.Start()
				f := bufio.NewWriter(os.Stdout)
				f.Write([]byte("some important message\n"))
				f.Flush()
				f.Write([]byte("another"))
				f.Flush()
				Expect(watcher.Stop()).To(ContainSubstring("some important message"))
				Expect(watcher.Stop()).To(ContainSubstring("another"))
			})
		})
	})
})
