package logging

import "github.com/sirupsen/logrus"

type Loggable interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Printf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Warningf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})

	Debug(args ...interface{})
	Info(args ...interface{})
	Print(args ...interface{})
	Warn(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})
}

type Fields logrus.Fields

type LoggableEntity struct {
	name      string
	logFields Fields
}

func NewLoggableEntity(name string, logFields Fields) LoggableEntity {
	return LoggableEntity{
		name,
		logFields,
	}
}

func (l LoggableEntity) GetName() string {
	return l.name
}

func (l LoggableEntity) GetFields() Fields {
	return l.logFields
}

func (l LoggableEntity) Info(args ...interface{}) {
	logrus.WithFields((logrus.Fields)(l.logFields)).Info(args...)
}

func (l LoggableEntity) Warning(args ...interface{}) {
	logrus.WithFields((logrus.Fields)(l.logFields)).Warn(args...)
}

func (l LoggableEntity) Warn(args ...interface{}) {
	logrus.WithFields((logrus.Fields)(l.logFields)).Warn(args...)
}

func (l LoggableEntity) Error(args ...interface{}) {
	logrus.WithFields((logrus.Fields)(l.logFields)).Error(args...)
}

func (l LoggableEntity) Debug(args ...interface{}) {
	logrus.WithFields((logrus.Fields)(l.logFields)).Debug(args...)
}

func (l LoggableEntity) Print(args ...interface{}){
	logrus.WithFields((logrus.Fields)(l.logFields)).Print(args...)
}

func (l LoggableEntity) Fatal(args ...interface{}) {
	logrus.WithFields((logrus.Fields)(l.logFields)).Fatal(args...)
}

func (l LoggableEntity) Panic(args ...interface{}) {
	logrus.WithFields((logrus.Fields)(l.logFields)).Panic(args...)
}

func (l LoggableEntity) Debugf(format string, args ...interface{}) {
	logrus.WithFields((logrus.Fields)(l.logFields)).Debugf(format, args...)
}

func (l LoggableEntity) Infof(format string, args ...interface{}) {
	logrus.WithFields((logrus.Fields)(l.logFields)).Infof(format, args...)
}

func (l LoggableEntity) Printf(format string, args ...interface{}) {
	logrus.WithFields((logrus.Fields)(l.logFields)).Printf(format, args...)
}

func (l LoggableEntity) Warnf(format string, args ...interface{}) {
	logrus.WithFields((logrus.Fields)(l.logFields)).Warnf(format, args...)
}

func (l LoggableEntity) Warningf(format string, args ...interface{}) {
	logrus.WithFields((logrus.Fields)(l.logFields)).Warningf(format, args...)
}

func (l LoggableEntity) Errorf(format string, args ...interface{}) {
	logrus.WithFields((logrus.Fields)(l.logFields)).Errorf(format, args...)
}

func (l LoggableEntity) Fatalf(format string, args ...interface{}) {
	logrus.WithFields((logrus.Fields)(l.logFields)).Fatalf(format, args...)
}

func (l LoggableEntity) Panicf(format string, args ...interface{}) {
	logrus.WithFields((logrus.Fields)(l.logFields)).Panicf(format, args...)
}
