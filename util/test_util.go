package util

import (
	"os"
	"io"
	"bytes"
	"github.com/sirupsen/logrus"
)

type StdoutWatcher struct {
	r *os.File
	w *os.File
	old *os.File
}

func (s *StdoutWatcher) Start() {
	s.old = os.Stdout
	r, w, _ := os.Pipe()
	s.r  = r
	s.w = w
	os.Stdout = s.w
}

func (s *StdoutWatcher) Stop() string {
	defer func() {
		s.r.Close()
		os.Stdout = s.old
	}()
	var buf bytes.Buffer
	s.w.Close()
	_, err := io.Copy(&buf, s.r)
	if err != nil {
		logrus.Fatal(err)
	}

	return buf.String()
}
