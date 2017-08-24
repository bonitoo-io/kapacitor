package diagnostic

import (
	"github.com/influxdata/kapacitor/keyvalue"
	"github.com/influxdata/kapacitor/services/slack"
	"github.com/influxdata/kapacitor/services/victorops"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

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
