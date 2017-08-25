package diagnostic

import (
	"log"
	"time"

	"github.com/influxdata/kapacitor/keyvalue"
	"github.com/influxdata/kapacitor/services/slack"
	"github.com/influxdata/kapacitor/services/victorops"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// HTTPD handler

type HTTPDHandler struct {
	l *zap.Logger
}

func (h *HTTPDHandler) NewHTTPServerErrorLogger() *log.Logger {
	// TODO: implement
	//panic("not implemented")
	return nil
}

func (h *HTTPDHandler) StartingService() {
	h.l.Info("starting HTTP service")
}

func (h *HTTPDHandler) StoppedService() {
	h.l.Info("closed HTTP service")
}

func (h *HTTPDHandler) ShutdownTimeout() {
	h.l.Error("shutdown timedout, forcefully closing all remaining connections")
}

func (h *HTTPDHandler) AuthenticationEnabled(enabled bool) {
	h.l.Info("authentication", zap.Bool("enabled", enabled))
}

func (h *HTTPDHandler) ListeningOn(addr string, proto string) {
	h.l.Info("listening on", zap.String("addr", addr), zap.String("protocol", proto))
}

func (h *HTTPDHandler) WriteBodyReceived(body string) {
	h.l.Debug("write body received by handler: %s", zap.String("body", body))
}

func (h *HTTPDHandler) HTTP(
	host string,
	username string,
	start time.Time,
	method string,
	uri string,
	proto string,
	status int,
	referer string,
	userAgent string,
	reqID string,
	duration time.Duration,
) {
	// TODO: what is the message?
	h.l.Info("???",
		zap.String("host", host),
		zap.String("username", username),
		zap.Time("start", start),
		zap.String("method", method),
		zap.String("uri", uri),
		zap.String("protocol", proto),
		zap.Int("status", status),
		zap.String("referer", referer),
		zap.String("user-agent", userAgent),
		zap.String("request-id", reqID),
		zap.Duration("duration", duration),
	)
}

func (h *HTTPDHandler) RecoveryError(
	msg string,
	err string,
	host string,
	username string,
	start time.Time,
	method string,
	uri string,
	proto string,
	status int,
	referer string,
	userAgent string,
	reqID string,
	duration time.Duration,
) {
	h.l.Error(
		msg,
		zap.String("err", err),
		zap.String("host", host),
		zap.String("username", username),
		zap.Time("start", start),
		zap.String("method", method),
		zap.String("uri", uri),
		zap.String("protocol", proto),
		zap.Int("status", status),
		zap.String("referer", referer),
		zap.String("user-agent", userAgent),
		zap.String("request-id", reqID),
		zap.Duration("duration", duration),
	)
}

func (h *HTTPDHandler) Error(msg string, err error) {
	h.l.Error(msg, zap.Error(err))
}

// Reporting handler
type ReportingHandler struct {
	l *zap.Logger
}

func (h *ReportingHandler) Error(msg string, err error) {
	h.l.Error(msg, zap.Error(err))
}

// Slack Handler

type SlackHandler struct {
	l *zap.Logger
}

func (h *SlackHandler) InsecureSkipVerify() {
	h.l.Warn("service is configured to skip ssl verification")
}

func (h *SlackHandler) Error(msg string, err error) {
	h.l.Error(msg, zap.Error(err))
}

func (h *SlackHandler) WithContext(ctx ...keyvalue.T) slack.Diagnostic {
	fields := []zapcore.Field{}
	for _, kv := range ctx {
		fields = append(fields, zap.String(kv.Key, kv.Value))
	}

	return &SlackHandler{
		l: h.l.With(fields...),
	}
}

// Storage Handler

type StorageHandler struct {
	l *zap.Logger
}

func (h *StorageHandler) Error(msg string, err error) {
	h.l.Error(msg, zap.Error(err))
}

// TaskStore Handler

type TaskStoreHandler struct {
	l *zap.Logger
}

func (h *TaskStoreHandler) StartingTask(taskID string) {
	h.l.Debug("starting enabled task on startup", zap.String("task", taskID))
}

func (h *TaskStoreHandler) StartedTask(taskID string) {
	h.l.Debug("started task during startup", zap.String("task", taskID))
}

func (h *TaskStoreHandler) FinishedTask(taskID string) {
	h.l.Debug("task finished", zap.String("task", taskID))
}

func (h *TaskStoreHandler) Debug(msg string) {
	h.l.Debug(msg)
}

func (h *TaskStoreHandler) Error(msg string, err error, ctx ...keyvalue.T) {
	// Special case the three ways that the function is actually used
	// to avoid allocations
	if len(ctx) == 0 {
		h.l.Error(msg, zap.Error(err))
		return
	}

	if len(ctx) == 1 {
		el := ctx[0]
		h.l.Error(msg, zap.Error(err), zap.String(el.Key, el.Value))
		return
	}

	if len(ctx) == 2 {
		x := ctx[0]
		y := ctx[1]
		h.l.Error(msg, zap.Error(err), zap.String(x.Key, x.Value), zap.String(y.Key, y.Value))
		return
	}

	// This isn't great wrt to allocation, but should not ever actually occur
	fields := make([]zapcore.Field, len(ctx)+1) // +1 for error
	fields[0] = zap.Error(err)
	for i := 1; i < len(fields); i++ {
		kv := ctx[i-1]
		fields[i] = zap.String(kv.Key, kv.Value)
	}

	h.l.Error(msg, fields...)
}

func (h *TaskStoreHandler) AlreadyMigrated(entity, id string) {
	h.l.Debug("entity has already been migrated skipping", zap.String(entity, id))
}

func (h *TaskStoreHandler) Migrated(entity, id string) {
	h.l.Debug("entity was migrated to new storage service", zap.String(entity, id))
}

// VictorOps Handler

type VictorOpsHandler struct {
	l *zap.Logger
}

func (h *VictorOpsHandler) Error(msg string, err error) {
	h.l.Error(msg, zap.Error(err))
}

func (h *VictorOpsHandler) WithContext(ctx ...keyvalue.T) victorops.Diagnostic {
	fields := []zapcore.Field{}
	for _, kv := range ctx {
		fields = append(fields, zap.String(kv.Key, kv.Value))
	}

	return &VictorOpsHandler{
		l: h.l.With(fields...),
	}
}