package util

import (
	. "github.com/onsi/ginkgo"
	//. "github.com/onsi/gomega"
	"bufio"
	"os"
//	"fmt"
	"fmt"
)

var _ = Describe("StdoutWatcher", func(){
	Context("watch is called, and has printed to stdout", func() {
		When("close is called", func() {
			It("should text from stdout", func() {
				watcher := GetNewStdoutWatcher()
				watcher.Watch()

				f := bufio.NewWriter(os.Stdout)
				f.Write([]byte("some important message"))
				f.Flush()

				watcher.Close()
				fmt.Println(watcher.buf.String())
				fmt.Println(watcher.buf.String())
				fmt.Println(watcher.buf.String())
				fmt.Println(watcher.buf.String())
				//fmt.Println(result)
			})
		})
	})
})
