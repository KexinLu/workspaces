package util

import (
	"os"
	"io"
	"sync"
	"bufio"
	"bytes"
)

type StdoutWatcher struct {
	wg sync.WaitGroup
	pipe io.ReadCloser
	buf []byte
}

func GetNewStdoutWatcher() StdoutWatcher {
	return StdoutWatcher{
		sync.WaitGroup{},
		io.ReadCloser(os.Stdout),
		make([]byte, 100, 1000),
	}
}

func (s *StdoutWatcher) Watch() {
	s.wg.Add(1)
	defer func() {
		n, _ := s.pipe.Read(s.buf)
		buffer := s.buf[0:n]

	}()
	pr := bufio.NewReader(os.Stdout)

	go func() {
		for {
			pr.ReadBytes('\n')
		}
	}()
	s.wg.Wait()
}

func (s *StdoutWatcher) Close() string {
}
