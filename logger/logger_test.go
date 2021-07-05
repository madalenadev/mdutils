package logger

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getStdout(f func(buf *bytes.Buffer)) {
	buf := &bytes.Buffer{}
	log.Out = buf
	f(buf)
}

var hasSeverity = map[string]string{
	INFO:    `"severity":"Info"`,
	DEFAULT: `"severity":"Default"`,
	NOTICE:  `"severity":"Notice"`,
	WARNING: `"severity":"Warning"`,
	ERROR:   `"severity":"Error"`,
	ALERT:   `"severity":"Alert"`,
}

func callFunc(severity string) {
	msgLog := "Hello, world..."
	switch severity {
	case INFO:
		Info(msgLog)
		return
	case DEFAULT:
		Default(msgLog)
		return
	case NOTICE:
		Notice(msgLog)
		return
	case WARNING:
		Warning(msgLog)
		return
	case ERROR:
		Error(msgLog)
		return
	case ALERT:
		Alert(msgLog)
		return
	}
}
func callFuncContext(ctx context.Context, severity string) {
	msgLog := "Hello, world..."
	switch severity {
	case INFO:
		InfoContext(ctx, msgLog)
		return
	case DEFAULT:
		DefaultContext(ctx, msgLog)
		return
	case NOTICE:
		NoticeContext(ctx, msgLog)
		return
	case WARNING:
		WarningContext(ctx, msgLog)
		return
	case ERROR:
		ErrorContext(ctx, msgLog)
		return
	case ALERT:
		AlertContext(ctx, msgLog)
		return
	}
}

func TestAllFieldsSeverity(t *testing.T) {
	for level, severity := range hasSeverity {
		getStdout(func(buf *bytes.Buffer) {
			callFunc(level)
			assert.Contains(t, buf.String(), severity, "Info() not contains "+severity)
		})
	}
}

func TestAllFieldsSeverityWithContext(t *testing.T) {
	for level, severity := range hasSeverity {
		getStdout(func(buf *bytes.Buffer) {
			callFuncContext(context.TODO(), level)
			assert.Contains(t, buf.String(), severity, "Info() not contains "+severity)
		})
	}
}
