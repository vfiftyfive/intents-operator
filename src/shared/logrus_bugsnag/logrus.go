package logrus_bugsnag

import (
	gerrors "errors"
	"github.com/bugsnag/bugsnag-go/v2"
	bugsnagerrors "github.com/bugsnag/bugsnag-go/v2/errors"
	"github.com/otterize/intents-operator/src/shared/errors"
	"github.com/sirupsen/logrus"
)

// Based on github.com/Shopify/logrus-bugsnag, but adapted to use the v2 version of bugsnag-go
type bugsnagHook struct{}

// ErrBugsnagUnconfigured is returned if NewBugsnagHook is called before
// bugsnag.Configure. Bugsnag must be configured before the hook.
var ErrBugsnagUnconfigured = gerrors.New("bugsnag must be configured before installing this logrus hook")

// ErrBugsnagSendFailed indicates that the hook failed to submit an error to
// bugsnag. The error was successfully generated, but `bugsnag.Notify()`
// failed.
type ErrBugsnagSendFailed struct {
	err error
}

func (e ErrBugsnagSendFailed) Error() string {
	return "failed to send error to Bugsnag: " + e.err.Error()
}

// NewBugsnagHook initializes a logrus hook which sends exceptions to an
// exception-tracking service compatible with the Bugsnag API. Before using
// this hook, you must call bugsnag.Configure(). The returned object should be
// registered with a log via `AddHook()`
//
// Entries that trigger an Error, Fatal or Panic should now include an "error"
// field to send to Bugsnag.
func NewBugsnagHook() (*bugsnagHook, error) {
	if bugsnag.Config.APIKey == "" {
		return nil, ErrBugsnagUnconfigured
	}
	return &bugsnagHook{}, nil
}

// skipStackFrames skips logrus stack frames before logging to Bugsnag.
const skipStackFrames = 3
const errorLogKey = "error"

type ErrorList interface {
	Unwrap() []error
}

// Fire forwards an error to Bugsnag. Given a logrus.Entry, it extracts the
// "error" field (or the Message if the error isn't present) and sends it off.
func (hook *bugsnagHook) Fire(entry *logrus.Entry) error {
	return SendToBugsnag(entry)
}

func getErrors(entry *logrus.Entry) []error {
	errFromMessage := bugsnagerrors.New(entry.Message, 1).Err
	errFromLog, ok := entry.Data[errorLogKey].(error)
	if !ok {
		return []error{errFromMessage}
	}
	// don't delete error key from entry as it's a standard logrus field and other hooks may make use of it

	bugsnagErr, ok := errFromLog.(*bugsnagerrors.Error)
	if !ok {
		return []error{errFromLog}
	}

	// Check if the underlying error field contains a list, usually created by calling to go stdlib errors.Join
	errList, ok := bugsnagErr.Err.(ErrorList)
	if !ok {
		return []error{bugsnagErr}
	}
	return errList.Unwrap()
}

func SendToBugsnag(entry *logrus.Entry, rawData ...any) error {
	bugsnagRawData := make([]any, 0)
	notifyErrs := getErrors(entry)

	metadata := bugsnag.MetaData{}
	metadata["metadata"] = make(map[string]interface{})
	for key, val := range entry.Data {
		if key != "error" {
			metadata["metadata"][key] = val
		}
	}

	bugsnagRawData = append(bugsnagRawData, metadata)
	bugsnagRawData = append(bugsnagRawData, rawData...)

	for _, notifyErr := range notifyErrs {
		errWithStack := errors.WrapWithSkip(notifyErr, skipStackFrames)
		bugsnagErr := bugsnag.Notify(errWithStack, bugsnagRawData...)
		if bugsnagErr != nil {
			return ErrBugsnagSendFailed{bugsnagErr}
		}
	}

	return nil
}

// Levels enumerates the log levels on which the error should be forwarded to
// bugsnag: everything at or above the "Error" level.
func (hook *bugsnagHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.ErrorLevel,
		logrus.FatalLevel,
		logrus.PanicLevel,
	}
}
