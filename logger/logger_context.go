package logger

import (
	"context"
)

// DefaultContext show detail of log
func DefaultContext(ctx context.Context, args ...interface{}) {
	setFieldSeverity(DEFAULT)
	log.WithFields(*fields).WithContext(ctx).Info(args...)
}

// NoticeContext show detail of log
func NoticeContext(ctx context.Context, args ...interface{}) {
	setFieldSeverity(NOTICE)
	log.WithFields(*fields).WithContext(ctx).Info(args...)
}

// WarningContext show detail of log
func WarningContext(ctx context.Context, args ...interface{}) {
	setFieldSeverity(WARNING)
	log.WithFields(*fields).WithContext(ctx).Warning(args...)
}

// AlertContext show detail of errors
func AlertContext(ctx context.Context, args ...interface{}) {
	setFieldSeverity(ALERT)
	log.WithFields(*fields).WithContext(ctx).Error(args...)
}

// CriticalContext show detail of log and generate a fatal
func CriticalContext(ctx context.Context, args ...interface{}) {
	setFieldSeverity(CRITICAL)
	log.WithFields(*fields).WithContext(ctx).Fatal(args...)
}

// EmergencyContext show detail of log and generate a panic
func EmergencyContext(ctx context.Context, args ...interface{}) {
	setFieldSeverity(EMERGENCY)
	log.WithFields(*fields).WithContext(ctx).Panic(args...)
}

// ErrorContext exibe detalhes do erro com o contexto
func ErrorContext(ctx context.Context, args ...interface{}) {
	setFieldSeverity(ERROR)
	log.WithFields(*fields).WithContext(ctx).Error(args...)
}

// InfoContext exibe detalhes do log info com o contexto
func InfoContext(ctx context.Context, args ...interface{}) {
	setFieldSeverity(INFO)
	log.WithFields(*fields).WithContext(ctx).Info(args...)
}

// DebugContext exibe detalhes do log debug com o contexto
func DebugContext(ctx context.Context, args ...interface{}) {
	setFieldSeverity(DEBUG)
	log.WithFields(*fields).WithContext(ctx).Debug(args...)
}

// TraceContext exibe detalhes do log trace com o contexto
func TraceContext(ctx context.Context, args ...interface{}) {
	setFieldSeverity(DEBUG)
	log.WithFields(*fields).WithContext(ctx).Debug(args...)
}
