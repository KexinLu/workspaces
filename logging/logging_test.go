package logging_test

import (
	. "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
	"io/ioutil"

	"workspaces/logging"
	_ "github.com/sirupsen/logrus"
	_ "github.com/sirupsen/logrus/hooks/test"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/sirupsen/logrus"
	"errors"
)

var _ = Describe("Loggable Entity", func() {
	var myEntity logging.LoggableEntity

	BeforeEach(func() {
		myEntity = logging.NewLoggableEntity(
			"myName",
			logging.Fields{
				"position": "here",
				"location": "there",
			},
		)

		logrus.SetOutput(ioutil.Discard)
	})

	When("Creating New Loggable Entity", func() {

		It("should has the provided name", func() {
			Expect(myEntity.GetName()).To(Equal("myName"))
		})

		It("should has the provided fields", func() {
			Expect(myEntity.GetFields()).To(Equal(
				logging.Fields{
					"position": "here",
					"location": "there",
				},
			))
		})
	})

	When("Logging with Loggable Entity", func() {
		hook := test.NewGlobal()
		logrus.SetOutput(ioutil.Discard)
		BeforeEach(func() {
			logrus.SetLevel(logrus.DebugLevel)
			hook.Reset()
		})

		It("should log with correct message", func() {
			myEntity.Warn("some message")
			Expect(hook.LastEntry().Message).To(Equal("some message"))
		})

		It("should log with correct fields", func() {
			myEntity.Warn("some message")
			Expect(hook.LastEntry().Data).To(Equal(
				(logrus.Fields)(logging.Fields{
					"position": "here",
					"location": "there",
				}),
			))
		})

		It("should info with info level", func() {
			myEntity.Info("txt")
			Expect(hook.LastEntry().Level).To(Equal(logrus.InfoLevel))
		})

		It("should infof with info level", func() {
			myEntity.Infof(logging.Fields{},"txt %s", "final")
			Expect(hook.LastEntry().Level).To(Equal(logrus.InfoLevel))
		})

		It("should warn with warn level", func() {
			myEntity.Warn("txt")
			Expect(hook.LastEntry().Level).To(Equal(logrus.WarnLevel))
		})

		It("should warnf with warn level", func() {
			myEntity.Warnf(logging.Fields{},"txt %s", "final")
			Expect(hook.LastEntry().Level).To(Equal(logrus.WarnLevel))
		})

		It("should error with error level", func() {
			myEntity.Error(errors.New("some error"), "txt")
			Expect(hook.LastEntry().Level).To(Equal(logrus.ErrorLevel))
		})

		It("should errorf with error level", func() {
			myEntity.Errorf(logging.Fields{}, errors.New("some error"), "txt %s", "final")
			Expect(hook.LastEntry().Level).To(Equal(logrus.ErrorLevel))
		})

		It("should debug with debug level", func() {
			myEntity.Debug("txt")
			Expect(hook.LastEntry().Level).To(Equal(logrus.DebugLevel))
		})

		It("should debugf with debug level", func() {
			myEntity.Debugf(logging.Fields{},"txt %s", "final")
			Expect(hook.LastEntry().Level).To(Equal(logrus.DebugLevel))
		})

		It("should fatal with fatal level", func() {
			myEntity.Fatal("txt")
			Expect(hook.LastEntry().Level).To(Equal(logrus.FatalLevel))
		})

		It("should fatalf with fatal level", func() {
			myEntity.Fatalf(logging.Fields{},"txt %s", errors.New("some er"), "final")
			Expect(hook.LastEntry().Level).To(Equal(logrus.FatalLevel))
		})

		It("should panic with panic level", func() {
			myEntity.Panic(errors.New("an panic error"), "txt")
			Expect(hook.LastEntry().Level).To(Equal(logrus.PanicLevel))
		})

		It("should panicf with panic level", func() {
			myEntity.Panicf(logging.Fields{}, errors.New("an panic error"), "txt %s", "final")
			Expect(hook.LastEntry().Level).To(Equal(logrus.PanicLevel))
		})
	})
})
