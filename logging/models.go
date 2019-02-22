package logging

import "github.com/sirupsen/logrus"

type Loggable interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})

	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
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

func (l LoggableEntity) InfoWithFields(fields Fields, args ...interface{}) {
	logrus.WithFields((logrus.Fields)(combineFields(fields, l.logFields))).Info(args...)
}

func (l LoggableEntity) Info(args ...interface{}) {
	l.InfoWithFields(Fields{}, args...)
}

func (l LoggableEntity) WarnWithFields(fields Fields, args ...interface{}) {
	logrus.WithFields((logrus.Fields)(combineFields(fields, l.logFields))).Warn(args...)
}

func (l LoggableEntity) Warn(args ...interface{}) {
	l.WarnWithFields(Fields{}, args...)
}

func (l LoggableEntity) ErrorWithFields(fields Fields, err error, args ...interface{}) {
	fs := fields
	fs["error"] = err
	logrus.WithFields((logrus.Fields)(combineFields(fs, l.logFields))).Error(args...)
}

func (l LoggableEntity) Error(err error, args ...interface{}) {
	l.ErrorWithFields(Fields{}, err, args...)
}

func (l LoggableEntity) DebugWithFields(fields Fields, args ...interface{}) {
	logrus.WithFields((logrus.Fields)(combineFields(fields, l.logFields))).Debug(args...)
}

func (l LoggableEntity) Debug(args ...interface{}) {
	l.DebugWithFields(Fields{}, args...)
}

func (l LoggableEntity) FatalWithFields(fields Fields, args ...interface{}) {
	logrus.WithFields((logrus.Fields)(combineFields(fields, l.logFields))).Fatal(args...)
}

func (l LoggableEntity) Fatal(args ...interface{}) {
	l.FatalWithFields(Fields{}, args...)
}

func (l LoggableEntity) PanicWithFields(fields Fields, err error, args ...interface{}) {
	logrus.WithFields((logrus.Fields)(combineFields(fields, l.logFields))).Panic(args...)
}

func (l LoggableEntity) Panic(err error, args ...interface{}) {
	l.PanicWithFields(Fields{}, err, args...)
}

func (l LoggableEntity) Debugf(fields Fields, format string, args ...interface{}) {
	logrus.WithFields((logrus.Fields)(combineFields(fields, l.logFields))).Debugf(format, args...)
}

func (l LoggableEntity) Infof(fields Fields, format string, args ...interface{}) {
	logrus.WithFields((logrus.Fields)(combineFields(fields, l.logFields))).Infof(format, args...)
}

func (l LoggableEntity) Warnf(fields Fields, format string, args ...interface{}) {
	logrus.WithFields((logrus.Fields)(combineFields(fields, l.logFields))).Warnf(format, args...)
}

func (l LoggableEntity) Errorf(fields Fields, err error, format string, args ...interface{}) {
	fs := fields
	fs["error"] = err
	logrus.WithFields((logrus.Fields)(combineFields(fs, l.logFields))).Errorf(format, args...)
}

// Combine two fields, first argument's key value pair will override second
func combineFields(fs1 Fields, fs2 Fields) Fields {
	result := make(map[string]interface{})
	for k,v := range fs2 {
		result[k] = v
	}
	for k,v := range fs1 {
		result[k] = v
	}

	return result
}

func (l LoggableEntity) Fatalf(fields Fields, format string, err error, args ...interface{}) {
	logrus.WithFields((logrus.Fields)(combineFields(fields, l.logFields))).Fatalf(format, args...)
}

func (l LoggableEntity) Panicf(fields Fields, err error, format string, args ...interface{}) {
	fs := fields
	fs["error"] = err
	logrus.WithFields((logrus.Fields)(combineFields(fs, l.logFields))).Panicf(format, args...)
}
